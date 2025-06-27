package pack

import (
	"bytes"
	"context"
	"io"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/packutil"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/gotidy/ptr"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/minio/sha256-simd"
	"github.com/multiformats/go-multihash"
	"github.com/multiformats/go-varint"
	"github.com/rclone/rclone/fs"
)

var ErrFileModified = errors.New("file has been modified")

// Assembler assembles various objects and data streams into a coherent output stream.
// It uses the provided file ranges, and internal buffers to produce the desired output.
type Assembler struct {
	rootCID cid.Cid
	// objects represents a map of object IDs to their corresponding fs.Object representations.
	objects map[model.FileID]fs.Object
	// ctx provides the context for the assembler's operations.
	ctx context.Context
	// reader is the storage system reader used to read file data.
	reader storagesystem.Reader
	// fileRanges contains the ranges of the files that should be read and processed.
	fileRanges []model.FileRange
	// index is the current position in the fileRanges slice.
	index int
	// buffer is a reader that holds data that's ready to be read out.
	buffer io.Reader
	// fileReadCloser reads the actual content from files.
	fileReadCloser io.ReadCloser
	// buf is a buffer for temporarily holding data.
	buf []byte
	// fileOffset tracks the offset into the current file being read.
	fileOffset int64
	// carOffset tracks the offset within a CAR (Content Addressable Archive).
	carOffset int64
	// pendingLinks contains a list of links waiting to be assembled.
	pendingLinks []format.Link
	// carBlocks is a slice of CAR blocks that are used in the assembly process.
	carBlocks []model.CarBlock
	// assembleLinkFor is a pointer to the index in fileRanges for which links need to be assembled.
	assembleLinkFor       *int
	noInline              bool
	skipInaccessibleFiles bool
	fileLengthCorrection  map[model.FileID]int64
}

// NewAssembler initializes a new Assembler instance with the given parameters.
func NewAssembler(ctx context.Context, reader storagesystem.Reader,
	fileRanges []model.FileRange, noInline bool, skipInaccessibleFiles bool,
) *Assembler {
	return &Assembler{
		ctx:                   ctx,
		reader:                reader,
		fileRanges:            fileRanges,
		buf:                   make([]byte, packutil.ChunkSize),
		objects:               make(map[model.FileID]fs.Object),
		noInline:              noInline,
		skipInaccessibleFiles: skipInaccessibleFiles,
		fileLengthCorrection:  make(map[model.FileID]int64),
	}
}

// Close closes the assembler and all of its underlying readers
func (a *Assembler) Close() error {
	if a.fileReadCloser != nil {
		err := a.fileReadCloser.Close()
		if err != nil {
			return errors.WithStack(err)
		}
		a.fileReadCloser = nil
	}
	return nil
}

// readBuffer reads data from the internal buffer, handling buffer-related flags and states.
// It returns the number of bytes read and any errors encountered.
func (a *Assembler) readBuffer(p []byte) (int, error) {
	n, err := a.buffer.Read(p)

	//nolint:errorlint
	switch err {
	case io.EOF:
		a.buffer = nil
		return n, nil
	case nil:
		return n, nil
	default:
		return n, errors.WithStack(err)
	}
}

// populateBuffer sets the buffer to a MultiReader containing the provided CAR blocks.
func (a *Assembler) populateBuffer(carBlocks []model.CarBlock) error {
	var readers []io.Reader
	if a.carOffset == 0 {
		a.rootCID = packutil.EmptyFileCid
		if len(carBlocks) > 0 {
			a.rootCID = cid.Cid(carBlocks[0].CID)
		}
		header, err := util.GenerateCarHeader(a.rootCID)
		if err != nil {
			return errors.WithStack(err)
		}
		readers = append(readers, bytes.NewReader(header))
		a.carOffset += int64(len(header))
	}
	for i, carBlock := range carBlocks {
		readers = append(readers, bytes.NewReader(carBlock.Varint), bytes.NewReader(cid.Cid(carBlock.CID).Bytes()), bytes.NewReader(carBlock.RawBlock))
		carBlocks[i].CarOffset = a.carOffset
		carBlockLength := int32(len(carBlock.Varint) + cid.Cid(carBlock.CID).ByteLen() + len(carBlock.RawBlock))
		carBlocks[i].CarBlockLength = carBlockLength
		a.carOffset += int64(carBlockLength)
	}

	a.buffer = io.MultiReader(readers...)

	return nil
}

// assembleLinks assembles links from pendingLinks and populates the buffer with CAR blocks.
func (a *Assembler) assembleLinks() error {
	defer func() {
		a.assembleLinkFor = nil
	}()
	if len(a.pendingLinks) == 0 {
		return nil
	}
	if len(a.pendingLinks) == 1 {
		a.fileRanges[*a.assembleLinkFor].CID = model.CID(a.pendingLinks[0].Cid)
		return nil
	}

	blks, rootNode, err := packutil.AssembleFileFromLinks(a.pendingLinks)
	if err != nil {
		return errors.WithStack(err)
	}

	rootCid := rootNode.Cid()
	a.fileRanges[*a.assembleLinkFor].CID = model.CID(rootCid)
	carBlocks := make([]model.CarBlock, len(blks))
	for i, blk := range blks {
		vint := varint.ToUvarint(uint64(blk.Cid().ByteLen() + len(blk.RawData())))
		carBlocks[i] = model.CarBlock{
			CID:      model.CID(blk.Cid()),
			Varint:   vint,
			RawBlock: blk.RawData(),
		}
	}

	err = a.populateBuffer(carBlocks)
	if err != nil {
		return errors.WithStack(err)
	}
	if !a.noInline {
		a.carBlocks = append(a.carBlocks, carBlocks...)
	}
	a.pendingLinks = nil
	return nil
}

// prefetch reads the next chunk from fileRanges and fills the buffer.
// This method should only be called when the buffer is empty.
// It might return without populating the buffer if there's no more data to read.
func (a *Assembler) prefetch() error {
	firstChunk := false
	if a.fileReadCloser == nil {
		fileRange := a.fileRanges[a.index]
		readCloser, obj, err := a.reader.Read(a.ctx, fileRange.File.Path, fileRange.Offset, fileRange.Length)
		if err != nil {
			if a.skipInaccessibleFiles {
				logger.Warnf("skipping inaccessible file %s: %v", fileRange.File.Path, err)
				a.index++
				return nil
			} else {
				return errors.Wrapf(err, "failed to open file %s", fileRange.File.Path)
			}
		}
		same, detail := storagesystem.IsSameEntry(a.ctx, *fileRange.File, obj)
		if !same {
			return errors.Wrapf(ErrFileModified, "fileRange has been modified: %s, %s", fileRange.File.Path, detail)
		}
		a.objects[fileRange.File.ID] = obj
		a.fileReadCloser = readCloser
		a.fileOffset = fileRange.Offset
		firstChunk = true
		a.pendingLinks = nil
	}

	hasher := sha256.New()
	reader := io.TeeReader(a.fileReadCloser, hasher)
	n, err := io.ReadFull(reader, a.buf)

	// Last empty chunk of a file
	if err == io.EOF && !firstChunk {
		a.assembleLinkFor = ptr.Of(a.index)
		a.fileReadCloser = nil
		_ = a.Close()
		if a.fileRanges[a.index].Length < 0 {
			a.fileLengthCorrection[a.fileRanges[a.index].FileID] = a.fileOffset
		}
		a.index++
		return nil
	}

	// read more than 0 bytes, or the first block of an empty file
	// nolint:err113
	if err == nil || errors.Is(err, io.ErrUnexpectedEOF) || err == io.EOF {
		var cidValue cid.Cid
		var vint []byte
		if err == io.EOF {
			cidValue = packutil.EmptyFileCid
			vint = packutil.EmptyFileVarint
		} else {
			sum := hasher.Sum(nil)
			mh, err2 := multihash.Encode(sum, multihash.SHA2_256)
			if err2 != nil {
				return errors.WithStack(err2)
			}
			cidValue = cid.NewCidV1(cid.Raw, mh)
			vint = varint.ToUvarint(uint64(cidValue.ByteLen() + n))
		}
		carBlocks := []model.CarBlock{{
			CID:        model.CID(cidValue),
			RawBlock:   a.buf[:n],
			Varint:     vint,
			FileOffset: a.fileOffset,
			FileID:     &a.fileRanges[a.index].FileID,
		}}
		err2 := a.populateBuffer(carBlocks)
		if err2 != nil {
			return errors.WithStack(err2)
		}
		carBlocks[0].RawBlock = nil
		if !a.noInline {
			a.carBlocks = append(a.carBlocks, carBlocks...)
		}

		// Check for negative file size
		size := n
		if size < 0 {
			logger.Warnf("Encountered unknown size file (%s)", a.fileRanges[a.index].File.Path)
			size = 0
		}

		a.pendingLinks = append(a.pendingLinks, format.Link{
			Cid:  cidValue,
			Size: uint64(size), //nolint:gosec
		})

		if err == nil {
			a.fileOffset += int64(n)
			return nil
		}

		a.assembleLinkFor = ptr.Of(a.index)
		_ = a.Close()
		if a.fileRanges[a.index].Length < 0 {
			a.fileLengthCorrection[a.fileRanges[a.index].FileID] = a.fileOffset + int64(n)
		}
		a.fileReadCloser = nil
		a.index++

		return nil
	}

	return errors.WithStack(err)
}

// Read reads data from the buffer, or fetches the next chunk from fileRanges if the buffer is empty.
// It will assemble links if needed and respect the context's cancellation or deadline.
func (a *Assembler) Read(p []byte) (int, error) {
	if a.ctx.Err() != nil {
		return 0, a.ctx.Err()
	}

	if a.buffer != nil {
		return a.readBuffer(p)
	}

	if a.assembleLinkFor != nil {
		return 0, errors.WithStack(a.assembleLinks())
	}

	if a.index == len(a.fileRanges) {
		return 0, io.EOF
	}

	return 0, a.prefetch()
}
