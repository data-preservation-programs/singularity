package datasetworker

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/scan"
)

func (w *Thread) scan(ctx context.Context, attachment model.SourceAttachment) error {
	return scan.Scan(ctx, w.dbNoContext, attachment)
}
