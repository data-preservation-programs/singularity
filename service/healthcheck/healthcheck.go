package healthcheck

import (
	"context"
	"os"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/google/uuid"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var staleThreshold = time.Minute * 5

type State struct {
	WorkType  model.WorkType
	WorkingOn string
}

func StartHealthCheckCleanup(ctx context.Context, db *gorm.DB) {
	for {
		HealthCheckCleanup(db)
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Minute):
			continue
		}
	}
}

func HealthCheckCleanup(db *gorm.DB) {
	// Remove all workers that haven't sent heartbeat for 5 minutes
	err := database.DoRetry(func() error {
		return db.Where("last_heartbeat < ?", time.Now().UTC().Add(-staleThreshold)).Delete(&model.Worker{}).Error
	})
	if err != nil {
		log.Logger("healthcheck").Errorw("failed to remove dead workers", "error", err)
	}

	// In case there are some works that have stale foreign key referenced to dead workers, we need to remove them
	err = database.DoRetry(func() error {
		return db.Model(&model.Source{}).Where("(dag_gen_worker_id NOT IN (?) OR dag_gen_worker_id IS NULL) AND dag_gen_state = ?",
			db.Table("workers").Select("id"), model.Processing).
			Updates(map[string]any{
				"dag_gen_worker_id": nil,
				"dag_gen_state":     model.Ready,
			}).Error
	})
	if err != nil {
		log.Logger("healthcheck").Errorw("failed to remove stale daggen worker", "error", err)
	}

	err = database.DoRetry(func() error {
		return db.Model(&model.Source{}).Where("(scanning_worker_id NOT IN (?) OR scanning_worker_id IS NULL) AND scanning_state = ?",
			db.Table("workers").Select("id"), model.Processing).
			Updates(map[string]any{
				"scanning_worker_id": nil,
				"scanning_state":     model.Ready,
			}).Error
	})
	if err != nil {
		log.Logger("healthcheck").Errorw("failed to remove stale scanning worker", "error", err)
	}

	err = database.DoRetry(func() error {
		return db.Model(&model.Chunk{}).Where("(packing_worker_id NOT IN (?) OR packing_worker_id IS NULL) AND packing_state = ?",
			db.Table("workers").Select("id"), model.Processing).
			Updates(map[string]any{
				"packing_worker_id": nil,
				"packing_state":     model.Ready,
			}).Error
	})
	if err != nil {
		log.Logger("healthcheck").Errorw("failed to remove stale packing worker", "error", err)
	}
}

func HealthCheck(db *gorm.DB, workerID uuid.UUID, getState func() State) {
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
	err = database.DoRetry(func() error {
		return db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"last_heartbeat", "work_type", "working_on"}),
		}).Create(&worker).Error
	})

	if err != nil {
		log.Logger("healthcheck").Errorw("failed to send heartbeat", "error", err)
	}
}

func StartHealthCheck(ctx context.Context, db *gorm.DB, workerID uuid.UUID, getState func() State) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Minute):
			HealthCheck(db.WithContext(ctx), workerID, getState)
			continue
		}
	}
}
