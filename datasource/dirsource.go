package datasource

import (
	"context"
	"io"
)

type DirSource struct{}

type FileSource struct{}

func (f FileSource) Open(ctx context.Context, path string, offset uint64, length uint64) (io.ReadCloser, error) {
	//TODO implement me
	panic("implement me")
}

func (f FileSource) GetDir(path string) string {
	//TODO implement me
	panic("implement me")
}

func (d DirSource) Scan(ctx context.Context, path string, last string) (<-chan Entry, error) {
	//TODO implement me
	panic("implement me")
}
