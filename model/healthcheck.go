package model

import (
	"context"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func StartHealthCheck(ctx context.Context, db *gorm.DB, workerID string) {
	logger := log.Logger("healthcheck").With("worker", workerID)
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		// Send heartbeat
		worker := Worker{
			ID:            workerID,
			LastHeartbeat: time.Now().UTC(),
		}
		err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"last_heartbeat"}),
		}).Create(&worker).Error

		if err != nil {
			logger.Errorw("failed to send heartbeat", "error", err)
		}

		// Remove all workers that haven't sent heartbeat for 5 minutes
		err = db.Where("last_heartbeat < ?", time.Now().UTC().Add(-5*time.Minute)).Delete(&Worker{}).Error
		if err != nil {
			logger.Errorw("failed to remove dead workers", "error", err)
		}

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			continue
		}
	}
}
