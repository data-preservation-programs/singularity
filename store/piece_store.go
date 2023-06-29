package store

import (
	"context"
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-varint"
	"github.com/rclone/rclone/fs"
	"io"
	"sort"

	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-cid"
	"github.com/pkg/errors"
)

var logger = log.Logger("piece_store")

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

func (pr *PieceReader) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		pr.pos = offset
	case io.SeekCurrent:
		pr.pos += offset
	case io.SeekEnd:
		pr.pos = pr.fileSize + offset
	default:
		return 0, errors.New("invalid whence")
	}
	if pr.pos < 0 {
		return 0, errors.New("negative position")
	}
	if pr.pos > pr.fileSize {
		return 0, errors.New("position past end of file")
	}
	if pr.reader != nil {
		pr.reader.Close()
		pr.reader = nil
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

func (pr *PieceReader) Clone(ctx context.Context) *PieceReader {
	return &PieceReader{
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
}

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
			return nil, errors.New("item source does not match source")
		}
	}

	// Sanitize carBlocks
	if len(carBlocks) == 0 {
		return nil, errors.New("no Blocks provided")
	}

	if carBlocks[0].CarOffset != int64(len(car.Header)) {
		return nil, errors.New("first block must start at car Header")
	}

	lastBlock := carBlocks[len(carBlocks)-1]
	if lastBlock.CarOffset+int64(lastBlock.CarBlockLength) != car.FileSize {
		return nil, errors.New("last block must end at car end")
	}

	for i := 0; i < len(carBlocks); i++ {
		if i != len(carBlocks)-1 {
			if carBlocks[i].CarOffset+int64(carBlocks[i].CarBlockLength) != carBlocks[i+1].CarOffset {
				return nil, errors.New("Blocks must be contiguous")
			}
		}
		vint, read, err := varint.FromUvarint(carBlocks[i].Varint)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse varint")
		}
		if read != len(carBlocks[i].Varint) {
			return nil, errors.New("varint does not match byte array length")
		}
		if uint64(carBlocks[i].BlockLength()) != vint-uint64(cid.Cid(carBlocks[i].CID).ByteLen()) {
			return nil, errors.New("varint does not match block length")
		}
		if carBlocks[i].RawBlock == nil {
			_, ok := itemsMap[*carBlocks[i].ItemID]
			if !ok {
				return nil, errors.New("item not found")
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

func (pr *PieceReader) Read(p []byte) (n int, err error) {
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
		if pr.reader != nil {
			pr.reader.Close()
			pr.reader = nil
		}
	}

	if pr.reader == nil {
		item := pr.items[*carBlock.ItemID]
		itemOffset := pr.pos - carBlock.CarOffset - int64(len(carBlock.Varint)) - int64(cid.Cid(carBlock.CID).ByteLen())
		itemOffset += carBlock.ItemOffset
		logger.Infow("reading item", "sourceID", item.SourceID, "path", item.Path, "offset", itemOffset)
		var obj fs.Object
		pr.reader, obj, err = pr.sourceHandler.Read(pr.ctx, item.Path, itemOffset, item.Size-itemOffset)
		if err != nil {
			return 0, errors.Wrap(err, "failed to get item")
		}
		isSameEntry, explanation := pack.IsSameEntry(pr.ctx, item, obj)
		if !isSameEntry {
			return 0, errors.New("item has changed: " + explanation)
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
			// This can be caused by original data source truncation
			err = errors.New("failed to read full block")
		}
	}
	return
}

func (pr *PieceReader) Close() error {
	if pr.reader == nil {
		return nil
	}
	return pr.reader.Close()
}
