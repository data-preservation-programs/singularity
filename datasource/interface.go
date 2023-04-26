package datasource

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/model"
	"io"
	"time"
)

type Entry struct {
	ScannedAt    time.Time
	Type         model.ItemType
	Path         string
	Size         uint64
	LastModified *time.Time
}

type Scanner interface {
	Scan(ctx context.Context, path string, last string) (<-chan Entry, error)
}

type Streamer interface {
	Open(ctx context.Context, path string, offset uint64, length uint64) (io.ReadCloser, error)
	GetDir(path string) string
}

var Scanners = map[model.SourceType]Scanner{
	model.Dir: DirSource{},
}

var Streamers = map[model.ItemType]Streamer{
	model.File: FileSource{},
}
