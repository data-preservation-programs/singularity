package pack

import (
	"bytes"
	"context"
	"io"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/encryption"
	util2 "github.com/data-preservation-programs/singularity/pack/util"
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
// It uses the provided encryption, file ranges, and internal buffers to produce the desired output.
type Assembler struct {
	// objects represents a map of object IDs to their corresponding fs.Object representations.
	objects map[uint64]fs.Object
	// ctx provides the context for the assembler's operations.
	ctx context.Context
	// reader is the storage system reader used to read file data.
	reader storagesystem.Reader
	// encryptor handles the encryption of the data if needed.
	encryptor encryption.Encryptor
	// fileRanges contains the ranges of the files that should be read and processed.
	fileRanges []model.FileRange
	// index is the current position in the fileRanges slice.
	index int
	// buffer is a reader that holds data that's ready to be read out.
	buffer io.Reader
	// fileReader reads the actual content from files.
	fileReader io.Reader
	// closers contains a list of io.Closers that need to be closed.
	closers []io.Closer
	// buf is a buffer for temporarily holding data.
	buf []byte
	// maxSize defines the maximum size of the output.
	maxSize int64
	// fileOffset tracks the offset into the current file being read.
	fileOffset int64
	// carOffset tracks the offset within a CAR (Content Addressable Archive).
	carOffset int64
	// pendingLinks contains a list of links waiting to be assembled.
	pendingLinks []format.Link
	// carBlocks is a slice of CAR blocks that are used in the assembly process.
	carBlocks []model.CarBlock
	// hasBoundary indicates if a boundary exists in the CAR.
	hasBoundary bool
	// assembleLinkFor is a pointer to the index in fileRanges for which links need to be assembled.
	assembleLinkFor *int
}

// Close closes the assembler and all of its underlying readers
func (a *Assembler) Close() error {
	var errs []error
	for _, closer := range a.closers {
		err := closer.Close()
		if err != nil {
			errs = append(errs, err)
		}
	}
	a.closers = nil
	if len(errs) > 0 {
		return util.AggregateError{Errors: errs}
	}
	return nil
}

// NewAssembler initializes a new Assembler instance with the given parameters.
func NewAssembler(ctx context.Context, reader storagesystem.Reader, encryptor encryption.Encryptor,
	fileRanges []model.FileRange, maxSize int64) *Assembler {
	return &Assembler{
		ctx:        ctx,
		reader:     reader,
		encryptor:  encryptor,
		fileRanges: fileRanges,
		buf:        make([]byte, util2.ChunkSize),
		maxSize:    maxSize,
		objects:    make(map[uint64]fs.Object),
	}
}

// Next checks if there are more chunks to read from the fileRanges or buffer.
// This method should only be called after hitting an io.EOF.
func (a *Assembler) Next() bool {
	if a.index >= len(a.fileRanges) && a.buffer == nil {
		return false
	}

	return true
}

// readBuffer reads data from the internal buffer, handling buffer-related flags and states.
// It returns the number of bytes read and any errors encountered.
func (a *Assembler) readBuffer(p []byte) (int, error) {
	n, err := a.buffer.Read(p)

	switch err {
	case io.EOF:
		a.buffer = nil
		if !a.hasBoundary {
			err = nil
		} else {
			a.hasBoundary = false
		}
		return n, err
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
		rootCid := EmptyFileCid
		if len(carBlocks) > 0 {
			rootCid = cid.Cid(carBlocks[0].CID)
		}
		header, err := util.GenerateCarHeader(rootCid)
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

	if a.carOffset > a.maxSize {
		a.hasBoundary = true
		a.carOffset = 0
	}

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

	blks, rootNode, err := util2.AssembleFileFromLinks(a.pendingLinks)
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
	a.carBlocks = append(a.carBlocks, carBlocks...)
	a.pendingLinks = nil
	return nil
}

// prefetch reads the next chunk from fileRanges and fills the buffer.
// This method should only be called when the buffer is empty.
// It might return without populating the buffer if there's no more data to read.
func (a *Assembler) prefetch() error {
	firstChunk := false
	if a.fileReader == nil {
		fileRange := a.fileRanges[a.index]
		readCloser, obj, err := a.reader.Read(a.ctx, fileRange.File.Path, fileRange.Offset, fileRange.Length)
		if err != nil {
			return errors.WithStack(err)
		}
		same, detail := storagesystem.IsSameEntry(a.ctx, *fileRange.File, obj)
		if !same {
			return errors.Wrapf(ErrFileModified, "fileRange has been modified: %s, %s", fileRange.File.Path, detail)
		}
		a.objects[fileRange.File.ID] = obj
		a.fileReader = readCloser
		a.fileOffset = fileRange.Offset
		a.closers = append(a.closers, readCloser)
		if a.encryptor != nil {
			encryptStream, err := a.encryptor.Encrypt(readCloser)
			if err != nil {
				return errors.WithStack(err)
			}
			a.fileReader = encryptStream
			a.closers = append(a.closers, encryptStream)
		}
		firstChunk = true
		a.pendingLinks = nil
	}

	hasher := sha256.New()
	reader := io.TeeReader(a.fileReader, hasher)
	n, err := io.ReadFull(reader, a.buf)

	// Last empty chunk of a file
	if err == io.EOF && !firstChunk {
		a.assembleLinkFor = ptr.Of(a.index)
		a.fileReader = nil
		a.Close()
		a.index++
		return nil
	}

	// read more than 0 bytes, or the first block of an empty file
	if err == nil || err == io.ErrUnexpectedEOF || err == io.EOF {
		var cidValue cid.Cid
		var vint []byte
		if err == io.EOF {
			cidValue = EmptyFileCid
			vint = EmptyFileVarint
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
			CID:           model.CID(cidValue),
			RawBlock:      a.buf[:n],
			Varint:        vint,
			FileOffset:    a.fileOffset,
			FileEncrypted: a.encryptor != nil,
			FileID:        &a.fileRanges[a.index].FileID,
		}}
		err2 := a.populateBuffer(carBlocks)
		if err2 != nil {
			return errors.WithStack(err2)
		}
		carBlocks[0].RawBlock = nil
		a.carBlocks = append(a.carBlocks, carBlocks[0])
		a.pendingLinks = append(a.pendingLinks, format.Link{
			Cid:  cidValue,
			Size: uint64(n),
		})

		if err == nil {
			a.fileOffset += int64(n)
			return nil
		}

		a.assembleLinkFor = ptr.Of(a.index)
		a.Close()
		a.fileReader = nil
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
