package source

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/model"
	"time"
)

type Entry struct {
	ScannedAt    time.Time
	Type         model.ItemType
	Path         string
	Size         uint64
	LastModified *time.Time
}

type DataSource interface {
	Scan(ctx context.Context, path string, last string) (<-chan Entry, error)
}
