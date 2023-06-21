package store

import (
	"context"
	"io"

	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-cid"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

type PieceBlock interface {
	GetPieceOffset() int64
}

type ItemBlockMetadata struct {
	PieceOffset int64   `json:"pieceOffset"`
	Varint      []byte  `json:"varint"`
	Cid         cid.Cid `json:"cid"`
	ItemOffset  int64   `json:"itemOffset"`
	ItemLength  int32   `json:"itemLength"`
}

func (i ItemBlockMetadata) GetPieceOffset() int64 {
	return i.PieceOffset
}
func (i ItemBlockMetadata) CidOffset() int64 {
	return i.PieceOffset + int64(len(i.Varint))
}
func (i ItemBlockMetadata) BlockOffset() int64 {
	return i.PieceOffset + int64(len(i.Varint)) + int64(i.Cid.ByteLen())
}
func (i ItemBlockMetadata) EndOffset() int64 {
	return i.PieceOffset + int64(len(i.Varint)) + int64(i.Cid.ByteLen()) + int64(i.ItemLength)
}
func (i ItemBlockMetadata) Length() int {
	return len(i.Varint) + i.Cid.ByteLen() + int(i.ItemLength)
}

type RawBlock struct {
	PieceOffset int64   `json:"pieceOffset"`
	Varint      []byte  `json:"varint"`
	Cid         cid.Cid `json:"cid"`
	BlockData   []byte  `json:"blockData"`
}

func (r RawBlock) GetPieceOffset() int64 {
	return r.PieceOffset
}
func (r RawBlock) CidOffset() int64 {
	return r.PieceOffset + int64(len(r.Varint))
}
func (r RawBlock) BlockOffset() int64 {
	return r.PieceOffset + int64(len(r.Varint)) + int64(r.Cid.ByteLen())
}
func (r RawBlock) EndOffset() int64 {
	return r.PieceOffset + int64(len(r.Varint)) + int64(r.Cid.ByteLen()) + int64(len(r.BlockData))
}

func (r RawBlock) Length() int {
	return len(r.Varint) + r.Cid.ByteLen() + len(r.BlockData)
}

type ItemBlock struct {
	PieceOffset   int64               `json:"pieceOffset"`
	SourceHandler datasource.Handler  `json:"-"`
	Item          *model.Item         `json:"item"`
	Meta          []ItemBlockMetadata `json:"meta"`
}

func (i ItemBlock) GetPieceOffset() int64 {
	return i.PieceOffset
}

type PieceReader struct {
	ctx          context.Context
	Blocks       []PieceBlock `json:"blocks"`
	reader       io.ReadCloser
	pos          int64
	blockID      int
	innerBlockID int
	blockOffset  int64
	Header       []byte `json:"header"`
}

func (pr *PieceReader) MakeCopy(ctx context.Context, offset int64) (*PieceReader, error) {
	newReader := &PieceReader{
		ctx:    ctx,
		Blocks: pr.Blocks,
		reader: nil,
		pos:    offset,
		Header: pr.Header,
	}

	if offset < int64(len(pr.Header)) {
		return newReader, nil
	}

	index, _ := slices.BinarySearchFunc(
		pr.Blocks, offset, func(b PieceBlock, o int64) int {
			return int(b.GetPieceOffset() - o)
		},
	)
	newReader.blockID = index
	switch block := pr.Blocks[index].(type) {
	case RawBlock:
		newReader.blockOffset = offset - block.GetPieceOffset()
	case ItemBlock:
		innerIndex, _ := slices.BinarySearchFunc(
			block.Meta, offset, func(b ItemBlockMetadata, o int64) int {
				return int(b.GetPieceOffset() - o)
			},
		)
		newReader.innerBlockID = innerIndex
		newReader.blockOffset = offset - block.Meta[innerIndex].GetPieceOffset()
	}

	return newReader, nil
}

func NewPieceReader(
	ctx context.Context,
	car model.Car,
	carBlocks []model.CarBlock,
	resolver datasource.HandlerResolver,
) (
	*PieceReader,
	error,
) {
	// Sanitize carBlocks
	if len(carBlocks) == 0 {
		return nil, errors.New("no Blocks provided")
	}

	if carBlocks[0].CarOffset != int64(len(car.Header)) {
		return nil, errors.New("first block must start at car Header")
	}

	lastBlock := carBlocks[len(carBlocks)-1]
	if lastBlock.CarOffset+int64(lastBlock.CarBlockLength) != car.FileSize {
		return nil, errors.New("last block must end at car footer")
	}

	for i := 0; i < len(carBlocks)-1; i++ {
		if carBlocks[i].CarOffset+int64(carBlocks[i].CarBlockLength) != carBlocks[i+1].CarOffset {
			return nil, errors.New("Blocks must be contiguous")
		}
		if carBlocks[i].RawBlock == nil && (carBlocks[i].Item == nil || carBlocks[i].Item.Source == nil) {
			return nil, errors.New("block must be either raw or Item, and the Item/source needs to be preloaded")
		}
	}

	// Combine nearby clocks with same Item
	blocks := make([]PieceBlock, 0)
	var lastItemBlock *ItemBlock
	for _, carBlock := range carBlocks {
		if lastItemBlock != nil && (carBlock.RawBlock != nil || lastItemBlock.Item.ID != carBlock.Item.ID) {
			blocks = append(blocks, *lastItemBlock)
			lastItemBlock = nil
		}
		if carBlock.RawBlock != nil {
			blocks = append(
				blocks, RawBlock{
					PieceOffset: carBlock.CarOffset,
					Varint:      carBlock.Varint,
					Cid:         cid.Cid(carBlock.CID),
					BlockData:   carBlock.RawBlock,
				},
			)
			continue
		}
		if lastItemBlock == nil {
			handler, err := resolver.Resolve(ctx, *carBlock.Item.Source)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get handler")
			}
			lastItemBlock = &ItemBlock{
				PieceOffset:   carBlock.CarOffset,
				SourceHandler: handler,
				Item:          carBlock.Item,
				Meta: []ItemBlockMetadata{
					{
						PieceOffset: carBlock.CarOffset,
						Varint:      carBlock.Varint,
						Cid:         cid.Cid(carBlock.CID),
						ItemOffset:  carBlock.ItemOffset,
						ItemLength:  carBlock.BlockLength(),
					},
				},
			}
			continue
		}
		// merge last Item with the new Item
		lastItemBlock.Meta = append(
			lastItemBlock.Meta, ItemBlockMetadata{
				PieceOffset: carBlock.CarOffset,
				Varint:      carBlock.Varint,
				Cid:         cid.Cid(carBlock.CID),
				ItemOffset:  carBlock.ItemOffset,
				ItemLength:  carBlock.BlockLength(),
			},
		)
	}
	if lastItemBlock != nil {
		blocks = append(blocks, *lastItemBlock)
	}

	return &PieceReader{
		ctx:          ctx,
		Blocks:       blocks,
		reader:       nil,
		pos:          0,
		blockID:      0,
		innerBlockID: 0,
		Header:       car.Header,
	}, nil
}

func (pr *PieceReader) Read(p []byte) (n int, err error) {
	if pr.blockID >= len(pr.Blocks) {
		return 0, io.EOF
	}
	if pr.pos < int64(len(pr.Header)) {
		copied := copy(p[n:], pr.Header[pr.pos:])
		pr.pos += int64(copied)
		n += copied
		if n == len(p) {
			return n, nil
		}
	}
	currentBlock := pr.Blocks[pr.blockID]
	if rawBlock, ok := currentBlock.(RawBlock); ok {
		if pr.pos < rawBlock.CidOffset() {
			copied := copy(p[n:], rawBlock.Varint[pr.pos-rawBlock.PieceOffset:])
			pr.pos += int64(copied)
			n += copied
			if n == len(p) {
				return n, nil
			}
		}
		if pr.pos < rawBlock.BlockOffset() {
			copied := copy(p[n:], rawBlock.Cid.Bytes()[pr.pos-rawBlock.CidOffset():])
			pr.pos += int64(copied)
			n += copied
			if n == len(p) {
				return n, nil
			}
		}
		if pr.pos < rawBlock.EndOffset() {
			copied := copy(p[n:], rawBlock.BlockData[pr.pos-rawBlock.BlockOffset():])
			pr.pos += int64(copied)
			n += copied
			if n == len(p) {
				return n, nil
			}
		}
		pr.blockID++
		pr.innerBlockID = 0
		return n, nil
	}

	itemBlock, _ := currentBlock.(ItemBlock)
	innerBlock := itemBlock.Meta[pr.innerBlockID]
	if pr.reader == nil {
		pr.reader, _, err = itemBlock.SourceHandler.Read(
			pr.ctx,
			itemBlock.Item.Path,
			innerBlock.ItemOffset+pr.blockOffset,
			itemBlock.Item.Size-(innerBlock.ItemOffset+pr.blockOffset),
		)
		if err != nil {
			return 0, errors.Wrap(err, "failed to read Item")
		}
	}
	if pr.pos < innerBlock.CidOffset() {
		copied := copy(p[n:], innerBlock.Varint[pr.pos-innerBlock.PieceOffset:])
		pr.pos += int64(copied)
		n += copied
		if n == len(p) {
			return n, nil
		}
	}
	if pr.pos < innerBlock.BlockOffset() {
		copied := copy(p[n:], innerBlock.Cid.Bytes()[pr.pos-innerBlock.CidOffset():])
		pr.pos += int64(copied)
		n += copied
		if n == len(p) {
			return n, nil
		}
	}
	if pr.pos < innerBlock.EndOffset() {
		readTill := min(len(p), n+int(innerBlock.EndOffset()-pr.pos))
		read, err := pr.reader.Read(p[n:readTill])
		n += read
		pr.pos += int64(read)
		if err != nil && !errors.Is(err, io.EOF) {
			return n, errors.Wrap(err, "failed to read Item")
		}
		if pr.pos == innerBlock.EndOffset() {
			pr.innerBlockID++
			if pr.innerBlockID >= len(itemBlock.Meta) {
				pr.blockID++
				pr.innerBlockID = 0
				pr.reader.Close()
				pr.reader = nil
			}
		}
		if n == len(p) {
			return n, nil
		}
	}
	return n, nil
}

func min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}

func (pr *PieceReader) Close() error {
	if pr.reader == nil {
		return nil
	}
	return pr.reader.Close()
}
