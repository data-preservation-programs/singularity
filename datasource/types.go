package datasource

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/rclone/rclone/fs"
	"io"
	"time"
)

// Entry is a struct that represents a single item during a data source scan
type Entry struct {
	Error     error
	IsDir     bool
	Info      fs.DirEntry
	ScannedAt time.Time
}

// Handler is an interface for scanning, reading, opening, and checking items in a data source.
type Handler interface {
	// Scan scans the data source starting at the given path and returns a channel of entries.
	// The `last` parameter is used to resume scanning from the last entry returned by a previous scan. It is exclusive.
	// The returned entries must be sorted by path in ascending order.
	Scan(ctx context.Context, path string, last string) <-chan Entry

	// Read reads data from the data source starting at the given path and offset, and returns a ReadCloser.
	// The `length` parameter specifies the number of bytes to read.
	// This method is most likely used for retrieving a single block of data.
	Read(ctx context.Context, path string, offset int64, length int64) (io.ReadCloser, fs.DirEntry, error)

	// Check checks the size and last modified time of the item at the given path.
	Check(ctx context.Context, path string) (fs.DirEntry, error)
}

type HandlerResolver interface {
	Resolve(ctx context.Context, source model.Source) (Handler, error)
}

type DefaultHandlerResolver struct {
}

func (DefaultHandlerResolver) Resolve(ctx context.Context, source model.Source) (Handler, error) {
	return NewRCloneHandler(ctx, source)
}

// EmptyReadCloser is a ReadCloser that always returns EOF.
type EmptyReadCloser struct{}

func (e *EmptyReadCloser) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

func (e *EmptyReadCloser) Close() error {
	return nil
}
