package store

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/datasource"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/pkg/errors"
	"io"
)

type pieceBlock interface {
	PieceOffset() uint64
}

type itemBlockMetadata struct {
	pieceOffset uint64
	varint      []byte
	cid         []byte
	itemOffset  uint64
	itemLength  int
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
	return len(i.varint) + len(i.cid) + i.itemLength
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
	sourceType    model.SourceType
	itemPath      string
	itemOffset    uint64
	itemLength    uint64
	meta          []itemBlockMetadata
}

func (i itemBlock) PieceOffset() uint64 {
	return i.pieceOffset
}

type PieceReader struct {
	blocks       []pieceBlock
	reader       io.ReadCloser
	pos          uint64
	blockID      int
	innerBlockID int
}

func (i *PieceReader) Read(p []byte) (n int, err error) {
	if i.blockID >= len(i.blocks) {
		return 0, io.EOF
	}
	currentBlock := i.blocks[i.blockID]
	if rawBlock, ok := currentBlock.(rawBlock); ok {
		if i.pos < rawBlock.CidOffset() {
			n += copy(p[n:], rawBlock.varint[i.pos-rawBlock.pieceOffset:])
			i.pos += uint64(n)
			if n == len(p) {
				return n, nil
			}
		}
		if i.pos < rawBlock.BlockOffset() {
			n += copy(p[n:], rawBlock.cid[i.pos-rawBlock.CidOffset():])
			i.pos += uint64(n)
			if n == len(p) {
				return n, nil
			}
		}
		if i.pos < rawBlock.EndOffset() {
			n += copy(p[n:], rawBlock.blockData[i.pos-rawBlock.BlockOffset():])
			i.pos += uint64(n)
			if n == len(p) {
				return n, nil
			}
		}
		i.blockID++
		i.innerBlockID = 0
		return n, nil
	}

	itemBlock, _ := currentBlock.(itemBlock)
	if i.innerBlockID >= len(itemBlock.meta) {
		i.blockID++
		i.innerBlockID = 0
		return n, nil
	}
	if i.reader == nil {
		i.reader, err = itemBlock.sourceHandler.Read(context.Background(), itemBlock.itemPath, itemBlock.itemOffset, itemBlock.itemLength)
		if err != nil {
			return 0, errors.Wrap(err, "failed to read item")
		}
	}
	innerBlock := itemBlock.meta[i.innerBlockID]
	if i.pos < innerBlock.CidOffset() {
		n += copy(p[n:], innerBlock.varint[i.pos-innerBlock.pieceOffset:])
		i.pos += uint64(n)
		if n == len(p) {
			return n, nil
		}
	}
	if i.pos < innerBlock.BlockOffset() {
		n += copy(p[n:], innerBlock.cid[i.pos-innerBlock.CidOffset():])
		i.pos += uint64(n)
		if n == len(p) {
			return n, nil
		}
	}
	if i.pos < innerBlock.EndOffset() {
		read, err := i.reader.Read(p[n : n+int(innerBlock.EndOffset()-i.pos)])
		n += read
		if err != nil {
			return n, errors.Wrap(err, "failed to read item")
		}
		i.pos += uint64(n)
		if i.pos == innerBlock.EndOffset() {
			i.innerBlockID++
			if i.innerBlockID >= len(itemBlock.meta) {
				i.blockID++
				i.innerBlockID = 0
				i.reader.Close()
				i.reader = nil
			}
		}
		if n == len(p) {
			return n, nil
		}
	}
	return n, nil
}

func (i *PieceReader) Close() error {
	if i.reader == nil {
		return nil
	}
	return i.reader.Close()
}
