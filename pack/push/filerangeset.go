package push

import (
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/packutil"
	"github.com/rjNemo/underscore"
)

type FileRangeSet struct {
	fileRanges []model.FileRange
	carSize    int64
}

var carHeaderSize = len(packutil.EmptyCarHeader)

func NewFileRangeSet() *FileRangeSet {
	return &FileRangeSet{
		fileRanges: make([]model.FileRange, 0),
		// Some buffer for header
		carSize: int64(carHeaderSize),
	}
}

func (r *FileRangeSet) CarSize() int64 {
	return r.carSize
}

func (r *FileRangeSet) FileRanges() []model.FileRange {
	return r.fileRanges
}

func (r *FileRangeSet) Add(fileRanges ...model.FileRange) {
	r.fileRanges = append(r.fileRanges, fileRanges...)
	for _, fileRange := range fileRanges {
		r.carSize += toCarSize(fileRange.Length)
	}
}

func (r *FileRangeSet) AddIfFits(fileRange model.FileRange, maxSize int64) bool {
	nextSize := toCarSize(fileRange.Length)
	if r.carSize+nextSize > maxSize {
		return false
	}
	r.fileRanges = append(r.fileRanges, fileRange)
	r.carSize += nextSize
	return true
}

func (r *FileRangeSet) Reset() {
	r.fileRanges = make([]model.FileRange, 0)
	r.carSize = int64(carHeaderSize)
}

func (r *FileRangeSet) FileRangeIDs() []model.FileRangeID {
	return underscore.Map(r.fileRanges, func(fileRange model.FileRange) model.FileRangeID {
		return fileRange.ID
	})
}

func toCarSize(size int64) int64 {
	if size == 0 {
		return 37
	}
	out := size
	nBlocks := (size-1)/1024/1024 + 1

	// For each block, we need to Add the bytes for the CID as well as varint
	out += nBlocks * (36 + 3)

	// Estimate the parent block for those blocks. The parent block stores the CID and the size
	if nBlocks > 1 {
		out += nBlocks*52 + 44
	}

	return out
}
