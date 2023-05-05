package service

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/google/uuid"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
	"time"
)

var staleThreshold = time.Minute * 5

type State struct {
	WorkType  model.WorkType
	WorkingOn string
}

func StartHealthCheckCleanup(ctx context.Context, db *gorm.DB) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		healthCheckCleanup(db)
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			continue
		}
	}
}

func healthCheckCleanup(db *gorm.DB) {
	// Remove all workers that haven't sent heartbeat for 5 minutes
	err := db.Where("last_heartbeat < ?", time.Now().UTC().Add(-staleThreshold)).Delete(&model.Worker{}).Error
	if err != nil {
		log.Logger("healthcheck").Errorw("failed to remove dead workers", "error", err)
	}
}

func healthCheck(db *gorm.DB, workerID uuid.UUID, getState func() State) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Logger("healthcheck").Warnw("failed to get hostname", "error", err)
	}
	state := getState()
	worker := model.Worker{
		ID:            workerID.String(),
		LastHeartbeat: time.Now().UTC(),
		Hostname:      hostname,
		WorkType:      state.WorkType,
		WorkingOn:     state.WorkingOn,
	}
	err = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"last_heartbeat", "work_type", "working_on"}),
	}).Create(&worker).Error

	if err != nil {
		log.Logger("healthcheck").Errorw("failed to send heartbeat", "error", err)
	}
}

func StartHealthCheck(ctx context.Context, db *gorm.DB, workerID uuid.UUID, getState func() State) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		healthCheck(db, workerID, getState)
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			continue
		}
	}
}
