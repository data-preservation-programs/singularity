package fssource

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/source"
)

type DirSource struct{}

func (d DirSource) Scan(ctx context.Context, path string, last string) (<-chan source.Entry, error) {
	//TODO implement me
	panic("implement me")
}
