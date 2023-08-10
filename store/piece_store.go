package store

import (
	"context"
	"io"
	"sort"

	"github.com/data-preservation-programs/singularity/pack"
	"github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-varint"
	"github.com/rclone/rclone/fs"

	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-cid"
	"github.com/pkg/errors"
)

var logger = log.Logger("piece_store")
var ErrNoCarBlocks = errors.New("no Blocks provided")
var ErrSourceMismatch = errors.New("item source does not match source")
var ErrInvalidStartOffset = errors.New("first block must start at car Header")
var ErrInvalidEndOffset = errors.New("last block must end at car end")
var ErrIncontiguousBlocks = errors.New("Blocks must be contiguous")
var ErrInvalidVarintLength = errors.New("varint read does not match varint length")
var ErrVarintDoesNotMatchBlockLength = errors.New("varint does not match block length")
var ErrItemNotProvided = errors.New("item not provided")
var ErrInvalidWhence = errors.New("invalid whence")
var ErrNegativeOffset = errors.New("negative offset")
var ErrOffsetOutOfRange = errors.New("position past end of file")
var ErrTruncated = errors.New("original file has been truncated")

type FileHasChangedError struct {
	Message string
}

func (e *FileHasChangedError) Error() string {
	return e.Message
}
func (e *FileHasChangedError) Is(target error) bool {
	var errFileHasChanged *FileHasChangedError
	ok := errors.As(target, &errFileHasChanged)
	return ok
}

// PieceReader is a struct that represents a reader for pieces of data.
//
// Fields:
// ctx: The context in which the PieceReader operates. This can be used to cancel operations or set deadlines.
// fileSize: The size of the file being read.
// header: A byte slice representing the header of the file.
// sourceHandler: A Handler from the datasource package that is used to handle the source of the data.
// carBlocks: A slice of CarBlocks. These represent the blocks of data in the CAR (Content Addressable Archive) format.
// items: A map where the keys are item ID. This represents the items of data being read.
// reader: An io.ReadCloser that is used to read the data and close the reader when done.
// readerFor: A uint64 item ID that represents the current item being read.
// pos: An int64 that represents the current position in the data being read.
// blockIndex: An integer that represents the index of the current block being read.
type PieceReader struct {
	ctx           context.Context
	fileSize      int64
	header        []byte
	sourceHandler datasource.Handler
	carBlocks     []model.CarBlock
	items         map[uint64]model.Item
	reader        io.ReadCloser
	readerFor     uint64
	pos           int64
	blockIndex    int
}

// Seek is a method on the PieceReader struct that changes the position of the reader.
// It takes an offset and a 'whence' value as input, similar to the standard io.Seeker interface.
// The offset is added to the position determined by 'whence'.
// If 'whence' is io.SeekStart, the offset is from the start of the file.
// If 'whence' is io.SeekCurrent, the offset is from the current position.
// If 'whence' is io.SeekEnd, the offset is from the end of the file.
// If the resulting position is negative or beyond the end of the file, an error is returned.
// If a reader is currently open, it is closed before the position is changed.
//
// Parameters:
// offset: The offset to move the position by. Can be negative.
// whence: The position to move the offset from. Must be one of io.SeekStart, io.SeekCurrent, or io.SeekEnd.
//
// Returns:
// The new position after seeking, and an error if the seek operation failed.
func (pr *PieceReader) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		pr.pos = offset
	case io.SeekCurrent:
		pr.pos += offset
	case io.SeekEnd:
		pr.pos = pr.fileSize + offset
	default:
		return 0, ErrInvalidWhence
	}
	if pr.pos < 0 {
		return 0, ErrNegativeOffset
	}
	if pr.pos > pr.fileSize {
		return 0, ErrOffsetOutOfRange
	}
	if pr.reader != nil {
		pr.reader.Close()
		pr.reader = nil
		pr.readerFor = 0
	}

	if pr.pos < int64(len(pr.header)) {
		pr.blockIndex = -1
	} else {
		pr.blockIndex = sort.Search(len(pr.carBlocks), func(i int) bool {
			return pr.carBlocks[i].CarOffset > pr.pos
		}) - 1
	}

	return pr.pos, nil
}

// Clone is a method on the PieceReader struct that creates a new PieceReader with the same state as the original.
// It takes a context as input, which is used for the new PieceReader.
// The new PieceReader starts at the beginning of the data (position 0).
//
// Parameters:
// ctx: The context for the new PieceReader. This can be used to cancel operations or set deadlines.
//
// Returns:
// A new PieceReader that has the same state as the original, but with the provided context and starting at position 0.
func (pr *PieceReader) Clone(ctx context.Context) *PieceReader {
	reader := &PieceReader{
		ctx:           ctx,
		fileSize:      pr.fileSize,
		header:        pr.header,
		sourceHandler: pr.sourceHandler,
		carBlocks:     pr.carBlocks,
		items:         pr.items,
		reader:        pr.reader,
		readerFor:     pr.readerFor,
		pos:           pr.pos,
		blockIndex:    pr.blockIndex,
	}
	//nolint:errcheck
	reader.Seek(0, io.SeekStart)
	return reader
}

// NewPieceReader is a function that creates a new PieceReader.
// It takes a context, a Car model, a Source model, a slice of CarBlock models, a slice of Item models, and a HandlerResolver as input.
// It validates the input data and returns an error if any of it is invalid.
// The returned PieceReader starts at the beginning of the data (position 0).
//
// Parameters:
// ctx: The context for the new PieceReader. This can be used to cancel operations or set deadlines.
// car: A Car model that represents the CAR (Content Addressable Archive) file being read.
// source: A Source model that represents the source of the data.
// carBlocks: A slice of CarBlock models that represent the blocks of data in the CAR file.
// items: A slice of Item models that represent the items of data being read.
// resolver: A HandlerResolver that is used to resolve the handler for the source of the data.
//
// Returns:
// A new PieceReader that has been initialized with the provided data, and an error if the initialization failed.
func NewPieceReader(
	ctx context.Context,
	car model.Car,
	source model.Source,
	carBlocks []model.CarBlock,
	items []model.Item,
	resolver datasource.HandlerResolver,
) (
	*PieceReader,
	error,
) {
	itemsMap := make(map[uint64]model.Item)
	for _, item := range items {
		itemsMap[item.ID] = item
		if item.SourceID != source.ID {
			return nil, ErrSourceMismatch
		}
	}

	// Sanitize carBlocks
	if len(carBlocks) == 0 {
		return nil, ErrNoCarBlocks
	}

	if carBlocks[0].CarOffset != int64(len(car.Header)) {
		return nil, ErrInvalidStartOffset
	}

	lastBlock := carBlocks[len(carBlocks)-1]
	if lastBlock.CarOffset+int64(lastBlock.CarBlockLength) != car.FileSize {
		return nil, ErrInvalidEndOffset
	}

	for i := 0; i < len(carBlocks); i++ {
		if i != len(carBlocks)-1 {
			if carBlocks[i].CarOffset+int64(carBlocks[i].CarBlockLength) != carBlocks[i+1].CarOffset {
				return nil, ErrIncontiguousBlocks
			}
		}
		vint, read, err := varint.FromUvarint(carBlocks[i].Varint)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse varint")
		}
		if read != len(carBlocks[i].Varint) {
			return nil, ErrInvalidVarintLength
		}
		if uint64(carBlocks[i].BlockLength()) != vint-uint64(cid.Cid(carBlocks[i].CID).ByteLen()) {
			return nil, ErrVarintDoesNotMatchBlockLength
		}
		if carBlocks[i].RawBlock == nil {
			_, ok := itemsMap[*carBlocks[i].ItemID]
			if !ok {
				return nil, ErrItemNotProvided
			}
		}
	}

	sourceHandler, err := resolver.Resolve(ctx, source)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve source handler")
	}

	return &PieceReader{
		ctx:           ctx,
		header:        car.Header,
		fileSize:      car.FileSize,
		sourceHandler: sourceHandler,
		carBlocks:     carBlocks,
		items:         itemsMap,
		blockIndex:    -1,
	}, nil
}

// Read is a method on the PieceReader struct that reads data into the provided byte slice.
// It reads data from the current position of the PieceReader and advances the position accordingly.
// If the context of the PieceReader has been cancelled, it returns an error immediately.
// If the end of the file has been reached, it returns io.EOF.
// If the PieceReader is currently at a block boundary, it advances to the next block before reading data.
// If the PieceReader is currently at a varint or CID boundary within a block, it reads the varint or CID data.
// If the PieceReader is currently at a raw block boundary within a block, it reads the raw block data.
// If the PieceReader is currently at an item boundary within a block, it reads the item data.
// If the PieceReader encounters an error while reading data, it returns the error.
//
// Parameters:
// p: The byte slice to read data into.
//
// Returns:
// The number of bytes read, and an error if the read operation failed.
func (pr *PieceReader) Read(p []byte) (n int, err error) {
	if pr.ctx.Err() != nil {
		return 0, pr.ctx.Err()
	}

	// Read car header
	if pr.blockIndex == -1 {
		n = copy(p, pr.header[pr.pos:])
		pr.pos += int64(n)
		if pr.pos == int64(len(pr.header)) {
			pr.blockIndex = 0
		}
		return
	}

	if pr.pos >= pr.fileSize {
		return 0, io.EOF
	}

	carBlock := pr.carBlocks[pr.blockIndex]
	if pr.pos >= carBlock.CarOffset+int64(carBlock.CarBlockLength) {
		pr.blockIndex++
		carBlock = pr.carBlocks[pr.blockIndex]
	}

	if pr.pos < carBlock.CarOffset+int64(len(carBlock.Varint)) {
		n = copy(p, carBlock.Varint[pr.pos-carBlock.CarOffset:])
		pr.pos += int64(n)
		return
	}

	if pr.pos < carBlock.CarOffset+int64(len(carBlock.Varint))+int64(cid.Cid(carBlock.CID).ByteLen()) {
		cidBytes := cid.Cid(carBlock.CID).Bytes()
		n = copy(p, cidBytes[pr.pos-carBlock.CarOffset-int64(len(carBlock.Varint)):])
		pr.pos += int64(n)
		return
	}

	if carBlock.RawBlock != nil {
		n = copy(p, carBlock.RawBlock[pr.pos-carBlock.CarOffset-int64(len(carBlock.Varint))-int64(cid.Cid(carBlock.CID).ByteLen()):])
		pr.pos += int64(n)
		return
	}

	if pr.reader != nil && pr.readerFor != *carBlock.ItemID {
		pr.reader.Close()
		pr.reader = nil
	}

	if pr.reader == nil {
		item := pr.items[*carBlock.ItemID]
		itemOffset := pr.pos - carBlock.CarOffset - int64(len(carBlock.Varint)) - int64(cid.Cid(carBlock.CID).ByteLen())
		itemOffset += carBlock.ItemOffset
		logger.Infow("reading item", "sourceID", item.SourceID, "path", item.Path, "offset", itemOffset)
		var obj fs.Object
		pr.reader, obj, err = pr.sourceHandler.Read(pr.ctx, item.Path, itemOffset, item.Size-itemOffset)
		if err != nil {
			return 0, errors.Wrap(err, "failed to read file")
		}
		isSameEntry, explanation := pack.IsSameEntry(pr.ctx, item, obj)
		if !isSameEntry {
			return 0, &FileHasChangedError{Message: "file has changed: " + explanation}
		}

		pr.readerFor = item.ID
	}

	maxToRead := carBlock.CarOffset + int64(carBlock.CarBlockLength) - pr.pos
	if maxToRead > int64(len(p)) {
		maxToRead = int64(len(p))
	}
	limitReader := io.LimitReader(pr.reader, maxToRead)
	n, err = limitReader.Read(p)
	pr.pos += int64(n)
	if errors.Is(err, io.EOF) {
		err = nil
		pr.reader.Close()
		pr.reader = nil
		if pr.pos != carBlock.CarOffset+int64(carBlock.CarBlockLength) {
			err = ErrTruncated
		}
		return
	}
	if err != nil {
		return
	}
	return
}

func (pr *PieceReader) Close() error {
	if pr.reader == nil {
		return nil
	}
	return pr.reader.Close()
}
