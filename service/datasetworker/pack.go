package datasetworker

import (
	"context"

	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/model"
)

func (w *Thread) pack(
	ctx context.Context, chunk model.Chunk,
) error {
	_, err := datasource.Pack(ctx, w.dbNoContext, chunk, w.datasourceHandlerResolver)
	if err != nil {
		return err
	}

	return nil
}
