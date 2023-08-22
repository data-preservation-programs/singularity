package datasetworker

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
)

func (w *Thread) pack(
	ctx context.Context, job model.Job,
) error {
	_, err := pack.Pack(ctx, w.dbNoContext, job)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
