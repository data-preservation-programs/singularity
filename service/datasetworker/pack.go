package datasetworker

import (
	"context"

	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/model"
)

func (w *DatasetWorkerThread) pack(
	ctx context.Context, packJob model.PackJob,
) error {
	_, err := datasource.Pack(ctx, w.db, packJob, w.datasourceHandlerResolver)
	if err != nil {
		return err
	}

	return nil
}
