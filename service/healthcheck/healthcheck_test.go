package healthcheck

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestHealthCheckCleanup(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ctx, cancel := context.WithCancel(ctx)
		done := make(chan struct{})
		go func() {
			StartHealthCheckCleanup(ctx, db)
			close(done)
		}()
		time.Sleep(time.Second)
		cancel()
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("healthcheck cleanup didn't stop")
		}
	})
}
func TestHealthCheckReport(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ctx, cancel := context.WithCancel(ctx)
		done := make(chan struct{})
		go func() {
			StartReportHealth(ctx, db, uuid.New(), model.DatasetWorker)
			close(done)
		}()
		time.Sleep(time.Second)
		cancel()
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("healthcheck report didn't stop")
		}
	})
}

func TestHealthCheck(t *testing.T) {
	req := require.New(t)
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		id := uuid.New()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		alreadyRunning, err := Register(ctx, db, id, model.DatasetWorker, false)
		req.NoError(err)
		req.False(alreadyRunning)
		alreadyRunning, err = Register(ctx, db, uuid.New(), model.DatasetWorker, false)
		req.NoError(err)
		req.True(alreadyRunning)

		var worker model.Worker
		err = db.Where("id = ?", id.String()).First(&worker).Error
		req.Nil(err)
		req.Equal(model.DatasetWorker, worker.Type)
		req.NotEmpty(worker.Hostname)
		lastHeatbeat := worker.LastHeartbeat

		time.Sleep(time.Second)
		ReportHealth(context.Background(), db, id, model.DatasetWorker)

		err = db.Where("id = ?", id.String()).First(&worker).Error
		req.Nil(err)
		req.Equal(model.DatasetWorker, worker.Type)
		req.NotEmpty(worker.Hostname)
		req.NotEqual(lastHeatbeat, worker.LastHeartbeat)

		HealthCheckCleanup(ctx, db)
		err = db.Where("id = ?", id.String()).First(&worker).Error
		req.Nil(err)

		oldThreshold := staleThreshold
		staleThreshold = 0
		defer func() {
			staleThreshold = oldThreshold
		}()

		time.Sleep(time.Second)
		HealthCheckCleanup(ctx, db)
		err = db.Where("id = ?", id.String()).First(&worker).Error
		req.ErrorIs(err, gorm.ErrRecordNotFound)
	})
}
