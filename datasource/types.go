package datasource

import (
	"context"
	"io"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/rclone/rclone/fs"
)

// Entry is a struct that represents a single file during a data source scan
type Entry struct {
	Error error
	Info  fs.Object
}

// Handler is an interface for scanning, reading, opening, and checking files in a data source.
type Handler interface {
	// List lists the files at the given path.
	List(ctx context.Context, path string) ([]fs.DirEntry, error)

	// Scan scans the data source starting at the given path and returns a channel of entries.
	// The `last` parameter is used to resume scanning from the last entry returned by a previous scan. It is exclusive.
	// The returned entries must be sorted by path in ascending order.
	Scan(ctx context.Context, path string, last string) <-chan Entry

	// Check checks the size and last modified time of the file at the given path.
	Check(ctx context.Context, path string) (fs.DirEntry, error)

	ReadHandler
}

type ReadHandler interface {
	// Read reads data from the data source starting at the given path and offset, and returns a ReadCloser.
	// The `length` parameter specifies the number of bytes to read.
	// This method is most likely used for retrieving a single block of data.
	Read(ctx context.Context, path string, offset int64, length int64) (io.ReadCloser, fs.Object, error)
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
