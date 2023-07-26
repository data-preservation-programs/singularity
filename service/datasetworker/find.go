package datasetworker

import (
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func (w *DatasetWorkerThread) findDagWork() (*model.Source, error) {
	if !w.config.EnableDag {
		return nil, nil
	}
	var sources []model.Source

	err := database.DoRetry(func() error {
		return w.db.Transaction(func(db *gorm.DB) error {
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

	err = w.db.Model(&sources[0]).Association("Dataset").Find(&sources[0].Dataset)
	if err != nil {
		return nil, err
	}

	return &sources[0], nil
}

func (w *DatasetWorkerThread) findPackWork() (*model.Chunk, error) {
	if !w.config.EnablePack {
		return nil, nil
	}
	var chunks []model.Chunk

	err := database.DoRetry(func() error {
		return w.db.Transaction(func(db *gorm.DB) error {
			// First, find the id of the record to update
			err := db.
				Where("packing_state = ? OR (packing_state = ? AND packing_worker_id is null)", model.Ready, model.Processing).
				Order("id asc").
				Limit(1).
				Find(&chunks).Error
			if err != nil {
				return err
			}

			if len(chunks) == 0 {
				return nil
			}

			// Then, perform the update using the found id
			return db.Model(&chunks[0]).
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
	if len(chunks) == 0 {
		//nolint: nilnil
		return nil, nil
	}

	err = w.db.Model(&chunks[0]).Preload("Dataset").Association("Source").Find(&chunks[0].Source)
	if err != nil {
		return nil, err
	}

	err = w.db.Model(&chunks[0]).Preload("Item").Association("ItemParts").Find(&chunks[0].ItemParts)
	if err != nil {
		return nil, err
	}

	return &chunks[0], nil
}

func (w *DatasetWorkerThread) findScanWork() (*model.Source, error) {
	if !w.config.EnableScan {
		return nil, nil
	}
	var sources []model.Source
	// Find all ready sources or sources that is being processed but does not have a worker id,
	// or all source that is complete but needs rescanning
	err := database.DoRetry(func() error {
		return w.db.Transaction(func(db *gorm.DB) error {
			err := w.db.
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

	err = w.db.Model(&sources[0]).Association("Dataset").Find(&sources[0].Dataset)
	if err != nil {
		return nil, err
	}

	return &sources[0], nil
}
