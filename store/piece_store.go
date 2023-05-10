package store

import (
	"context"
	"io"

	"github.com/data-preservation-programs/go-singularity/datasource"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-varint"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

type pieceBlock interface {
	PieceOffset() uint64
}

type itemBlockMetadata struct {
	pieceOffset uint64
	varint      []byte
	cid         []byte
	itemOffset  uint64
	itemLength  uint64
}

func (i itemBlockMetadata) PieceOffset() uint64 {
	return i.pieceOffset
}
func (i itemBlockMetadata) CidOffset() uint64 {
	return i.pieceOffset + uint64(len(i.varint))
}
func (i itemBlockMetadata) BlockOffset() uint64 {
	return i.pieceOffset + uint64(len(i.varint)) + uint64(len(i.cid))
}
func (i itemBlockMetadata) EndOffset() uint64 {
	return i.pieceOffset + uint64(len(i.varint)) + uint64(len(i.cid)) + uint64(i.itemLength)
}
func (i itemBlockMetadata) Length() int {
	return len(i.varint) + len(i.cid) + int(i.itemLength)
}

type rawBlock struct {
	pieceOffset uint64
	varint      []byte
	cid         []byte
	blockData   []byte
}

func (r rawBlock) PieceOffset() uint64 {
	return r.pieceOffset
}
func (r rawBlock) CidOffset() uint64 {
	return r.pieceOffset + uint64(len(r.varint))
}
func (r rawBlock) BlockOffset() uint64 {
	return r.pieceOffset + uint64(len(r.varint)) + uint64(len(r.cid))
}
func (r rawBlock) EndOffset() uint64 {
	return r.pieceOffset + uint64(len(r.varint)) + uint64(len(r.cid)) + uint64(len(r.blockData))
}

func (r rawBlock) Length() int {
	return len(r.varint) + len(r.cid) + len(r.blockData)
}

type itemBlock struct {
	pieceOffset   uint64
	sourceHandler datasource.Handler
	item          *model.Item
	meta          []itemBlockMetadata
}

func (i itemBlock) PieceOffset() uint64 {
	return i.pieceOffset
}

type PieceReader struct {
	ctx          context.Context
	blocks       []pieceBlock
	reader       io.ReadCloser
	pos          uint64
	blockID      int
	innerBlockID int
	blockOffset  uint64
	header       []byte
}

func (pr *PieceReader) MakeCopy(ctx context.Context, offset uint64) (*PieceReader, error) {
	newReader := &PieceReader{
		ctx:    ctx,
		blocks: pr.blocks,
		reader: nil,
		pos:    offset,
		header: pr.header,
	}

	if offset < uint64(len(pr.header)) {
		return newReader, nil
	}

	index, _ := slices.BinarySearchFunc(
		pr.blocks, offset, func(b pieceBlock, o uint64) int {
			return int(b.PieceOffset() - o)
		},
	)
	newReader.blockID = index
	switch block := pr.blocks[index].(type) {
	case rawBlock:
		newReader.blockOffset = offset - block.PieceOffset()
	case itemBlock:
		innerIndex, _ := slices.BinarySearchFunc(
			block.meta, offset, func(b itemBlockMetadata, o uint64) int {
				return int(b.PieceOffset() - o)
			},
		)
		newReader.innerBlockID = innerIndex
		newReader.blockOffset = offset - block.meta[innerIndex].PieceOffset()
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
		return nil, errors.New("no blocks provided")
	}

	if carBlocks[0].CarOffset != uint64(len(car.Header)) {
		return nil, errors.New("first block must start at car header")
	}

	lastBlock := carBlocks[len(carBlocks)-1]
	if lastBlock.CarOffset+lastBlock.CarBlockLength != car.FileSize {
		return nil, errors.New("last block must end at car footer")
	}

	for i := 0; i < len(carBlocks)-1; i++ {
		if carBlocks[i].CarOffset+carBlocks[i].CarBlockLength != carBlocks[i+1].CarOffset {
			return nil, errors.New("blocks must be contiguous")
		}
		if carBlocks[i].RawBlock == nil && (carBlocks[i].Item == nil || carBlocks[i].Source == nil) {
			return nil, errors.New("block must be either raw or item, and the item/source needs to be preloaded")
		}
	}

	// Combine nearby clocks with same item
	blocks := make([]pieceBlock, 0)
	var lastItemBlock *itemBlock
	for _, carBlock := range carBlocks {
		if lastItemBlock != nil && (carBlock.RawBlock != nil || lastItemBlock.item.ID != carBlock.Item.ID) {
			blocks = append(blocks, *lastItemBlock)
			lastItemBlock = nil
		}
		if carBlock.RawBlock != nil {
			blocks = append(
				blocks, rawBlock{
					pieceOffset: carBlock.CarOffset,
					varint:      varint.ToUvarint(carBlock.Varint),
					cid:         cid.MustParse(carBlock.CID).Bytes(),
					blockData:   carBlock.RawBlock,
				},
			)
			continue
		}
		if lastItemBlock == nil {
			handler, err := resolver.GetHandler(*carBlock.Source)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get handler")
			}
			lastItemBlock = &itemBlock{
				pieceOffset:   carBlock.CarOffset,
				sourceHandler: handler,
				item:          carBlock.Item,
				meta: []itemBlockMetadata{
					{
						pieceOffset: carBlock.CarOffset,
						varint:      varint.ToUvarint(carBlock.Varint),
						cid:         cid.MustParse(carBlock.CID).Bytes(),
						itemOffset:  carBlock.ItemOffset,
						itemLength:  carBlock.BlockLength,
					},
				},
			}
			continue
		}
		// merge last item with the new item
		lastItemBlock.meta = append(
			lastItemBlock.meta, itemBlockMetadata{
				pieceOffset: carBlock.CarOffset,
				varint:      varint.ToUvarint(carBlock.Varint),
				cid:         cid.MustParse(carBlock.CID).Bytes(),
				itemOffset:  carBlock.ItemOffset,
				itemLength:  carBlock.BlockLength,
			},
		)
	}
	if lastItemBlock != nil {
		blocks = append(blocks, *lastItemBlock)
	}

	return &PieceReader{
		ctx:          ctx,
		blocks:       blocks,
		reader:       nil,
		pos:          0,
		blockID:      0,
		innerBlockID: 0,
		header:       car.Header,
	}, nil
}

func (pr *PieceReader) Read(p []byte) (n int, err error) {
	if pr.blockID >= len(pr.blocks) {
		return 0, io.EOF
	}
	if pr.pos < uint64(len(pr.header)) {
		copied := copy(p[n:], pr.header[pr.pos:])
		pr.pos += uint64(copied)
		n += copied
		if n == len(p) {
			return n, nil
		}
	}
	currentBlock := pr.blocks[pr.blockID]
	if rawBlock, ok := currentBlock.(rawBlock); ok {
		if pr.pos < rawBlock.CidOffset() {
			copied := copy(p[n:], rawBlock.varint[pr.pos-rawBlock.pieceOffset:])
			pr.pos += uint64(copied)
			n += copied
			if n == len(p) {
				return n, nil
			}
		}
		if pr.pos < rawBlock.BlockOffset() {
			copied := copy(p[n:], rawBlock.cid[pr.pos-rawBlock.CidOffset():])
			pr.pos += uint64(copied)
			n += copied
			if n == len(p) {
				return n, nil
			}
		}
		if pr.pos < rawBlock.EndOffset() {
			copied := copy(p[n:], rawBlock.blockData[pr.pos-rawBlock.BlockOffset():])
			pr.pos += uint64(copied)
			n += copied
			if n == len(p) {
				return n, nil
			}
		}
		pr.blockID++
		pr.innerBlockID = 0
		return n, nil
	}

	itemBlock, _ := currentBlock.(itemBlock)
	innerBlock := itemBlock.meta[pr.innerBlockID]
	if pr.reader == nil {
		pr.reader, err = itemBlock.sourceHandler.Read(
			pr.ctx,
			itemBlock.item.Path,
			innerBlock.itemOffset+pr.blockOffset,
			itemBlock.item.Size-(innerBlock.itemOffset+pr.blockOffset),
		)
		if err != nil {
			return 0, errors.Wrap(err, "failed to read item")
		}
	}
	if pr.pos < innerBlock.CidOffset() {
		copied := copy(p[n:], innerBlock.varint[pr.pos-innerBlock.pieceOffset:])
		pr.pos += uint64(copied)
		n += copied
		if n == len(p) {
			return n, nil
		}
	}
	if pr.pos < innerBlock.BlockOffset() {
		copied := copy(p[n:], innerBlock.cid[pr.pos-innerBlock.CidOffset():])
		pr.pos += uint64(copied)
		n += copied
		if n == len(p) {
			return n, nil
		}
	}
	if pr.pos < innerBlock.EndOffset() {
		readTill := min(len(p), n+int(innerBlock.EndOffset()-pr.pos))
		read, err := pr.reader.Read(p[n:readTill])
		n += read
		pr.pos += uint64(read)
		if err != nil {
			return n, errors.Wrap(err, "failed to read item")
		}
		if pr.pos == innerBlock.EndOffset() {
			pr.innerBlockID++
			if pr.innerBlockID >= len(itemBlock.meta) {
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
