package datasetworker

import (
	"context"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func (w *Thread) findDagWork(ctx context.Context) (*model.Source, error) {
	db := w.dbNoContext.WithContext(ctx)
	if !w.config.EnableDag {
		return nil, nil
	}
	w.logger.Debugw("finding ExportDag work")
	var sources []model.Source

	err := database.DoRetry(ctx, func() error {
		return db.Transaction(func(db *gorm.DB) error {
			// First, find the id of the record to update
			err := db.
				Where("dag_gen_state = ? OR (dag_gen_state = ? AND dag_gen_worker_id is null)",
					model.Ready, model.Processing).
				Order("id asc").
				Limit(1).
				Find(&sources).Error
			if err != nil {
				return err
			}

			if len(sources) == 0 {
				return nil
			}

			// Then, perform the update using the found id
			return db.Model(&sources[0]).
				Updates(map[string]any{
					"dag_gen_state":         model.Processing,
					"dag_gen_worker_id":     w.id,
					"dag_gen_error_message": "",
				}).Error
		})
	})

	if err != nil {
		return nil, err
	}
	if len(sources) == 0 {
		//nolint: nilnil
		return nil, nil
	}

	err = db.Model(&sources[0]).Association("Dataset").Find(&sources[0].Dataset)
	if err != nil {
		return nil, err
	}

	return &sources[0], nil
}

func (w *Thread) findPackWork(ctx context.Context) (*model.PackJob, error) {
	db := w.dbNoContext.WithContext(ctx)
	if !w.config.EnablePack {
		return nil, nil
	}
	w.logger.Debugw("finding pack work")
	var packJobs []model.PackJob

	err := database.DoRetry(ctx, func() error {
		return db.Transaction(func(db *gorm.DB) error {
			// First, find the id of the record to update
			err := db.
				Where("packing_state = ? OR (packing_state = ? AND packing_worker_id is null)", model.Ready, model.Processing).
				Order("id asc").
				Limit(1).
				Find(&packJobs).Error
			if err != nil {
				return err
			}

			if len(packJobs) == 0 {
				return nil
			}

			// Then, perform the update using the found id
			return db.Model(&packJobs[0]).
				Updates(map[string]any{
					"packing_state":     model.Processing,
					"packing_worker_id": w.id,
					"error_message":     "",
				}).Error
		})
	})

	if err != nil {
		return nil, err
	}
	if len(packJobs) == 0 {
		//nolint: nilnil
		return nil, nil
	}

	var src model.Source
	err = db.Joins("Dataset").Where("sources.id = ?", packJobs[0].SourceID).First(&src).Error
	if err != nil {
		return nil, err
	}
	packJobs[0].Source = &src

	var fileRanges []model.FileRange
	err = db.Joins("File").Where("file_ranges.pack_job_id = ?", packJobs[0].ID).Order("file_ranges.id asc").Find(&fileRanges).Error
	if err != nil {
		return nil, err
	}
	packJobs[0].FileRanges = fileRanges

	return &packJobs[0], nil
}

func (w *Thread) findScanWork(ctx context.Context) (*model.Source, error) {
	db := w.dbNoContext.WithContext(ctx)
	if !w.config.EnableScan {
		return nil, nil
	}
	w.logger.Debugw("finding scan work")
	var sources []model.Source
	// Find all ready sources or sources that is being processed but does not have a worker id,
	// or all source that is complete but needs rescanning
	err := database.DoRetry(ctx, func() error {
		return db.Transaction(func(db *gorm.DB) error {
			err := db.
				Where(
					"(scanning_state = ? OR (scanning_state = ? AND scanning_worker_id is null)) OR "+
						"(scanning_state = ? AND scan_interval_seconds > 0 AND last_scanned_timestamp + scan_interval_seconds < ?)",
					model.Ready,
					model.Processing,
					model.Complete,
					time.Now().UTC().Unix()).
				Order("id asc").
				Limit(1).Find(&sources).Error
			if err != nil {
				return err
			}
			if len(sources) == 0 {
				return nil
			}
			err = db.Model(&sources[0]).
				Updates(map[string]any{
					"scanning_state":     model.Processing,
					"scanning_worker_id": w.id,
					"error_message":      "",
				}).Error
			if err != nil {
				return err
			}
			return nil
		})
	})

	if err != nil {
		return nil, err
	}
	if len(sources) == 0 {
		//nolint: nilnil
		return nil, nil
	}

	err = db.Model(&sources[0]).Association("Dataset").Find(&sources[0].Dataset)
	if err != nil {
		return nil, err
	}

	return &sources[0], nil
}
