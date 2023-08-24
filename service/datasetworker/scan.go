package datasetworker

import (
	"context"

	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/model"
)

// scan scans the data source and inserts the packJobing strategy back to database
// scanSource is true if the source will be actually scanned in addition to just picking up remaining ones
// resume is true if the scan will be resumed from the last scanned file, which is useful for resuming a failed scan
func (w *Thread) scan(ctx context.Context, source model.Source, scanSource bool) error {
	return datasource.PrepareSource(ctx, w.dbNoContext.WithContext(ctx), w.datasourceHandlerResolver, source, scanSource)
}
