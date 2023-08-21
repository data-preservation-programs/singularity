package datasetworker

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/model"
)

func (w *Thread) pack(
	ctx context.Context, job model.Job,
) error {
	_, err := datasource.Pack(ctx, w.dbNoContext, job)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
