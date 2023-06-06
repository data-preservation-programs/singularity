package datasetworker

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm/clause"
	"time"
)

func (w *DatasetWorkerThread) findPackWork() (*model.Chunk, error) {
	if !w.config.EnablePack {
		return nil, nil
	}
	var chunks []model.Chunk
	err := database.DoRetry(func() error {
		return w.db.Model(&chunks).
			Where(
				"packing_state = ? OR (packing_state = ? AND packing_worker_id is null)",
				model.Ready,
				model.Processing,
			).
			Order("id asc").
			Limit(1).
			Clauses(clause.Returning{}).
			Updates(map[string]interface{}{
				"packing_state":     model.Processing,
				"packing_worker_id": w.id,
				"error_message":     "",
			}).Error
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
		return w.db.Model(&sources).
			Where("(scanning_state = ? OR (scanning_state = ? AND scanning_worker_id is null)) OR "+
				"(scanning_state = ? AND scan_interval_seconds > 0 AND last_scanned_timestamp + scan_interval_seconds < ?)",
				model.Ready,
				model.Processing,
				model.Complete,
				time.Now().UTC().Unix()).
			Order("id asc").
			Limit(1).
			Clauses(clause.Returning{}).
			Updates(map[string]interface{}{
				"scanning_state":     model.Processing,
				"scanning_worker_id": w.id,
				"error_message":      "",
			}).Error
	})
	if err != nil {
		return nil, err
	}
	if len(sources) == 0 {
		//nolint: nilnil
		return nil, nil
	}

	return &sources[0], nil
}
