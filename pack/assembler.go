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

type Assembler struct {
	objects         map[uint64]fs.Object
	ctx             context.Context
	reader          storagesystem.Reader
	encryptor       encryption.Encryptor
	fileRanges      []model.FileRange
	index           int
	buffer          io.Reader
	fileReader      io.Reader
	closers         []io.Closer
	buf             []byte
	maxSize         int64
	fileOffset      int64
	carOffset       int64
	pendingLinks    []format.Link
	carBlocks       []model.CarBlock
	hasBoundary     bool
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

// Next returns true if there is another chunk to read. This method should only be called after hitting io.EOF
func (a *Assembler) Next() bool {
	if a.index >= len(a.fileRanges) && a.buffer == nil {
		return false
	}

	return true
}

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

// populateBuffer will set the buffer to a MultiReader of the given content
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

	if a.buffer == nil {
		a.buffer = io.MultiReader(readers...)
	} else {
		readers = append([]io.Reader{a.buffer}, readers...)
		a.buffer = io.MultiReader(readers...)
	}

	if a.carOffset > a.maxSize {
		a.hasBoundary = true
		a.carOffset = 0
	}

	return nil
}

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

// prefetch reads the next chunk from the fileRanges and populates the buffer
// This method should only be used when the buffer is empty
// This method may return without populating the buffer when there is no more file to read
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
