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

func TestCleanupOrphanedRecords(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create orphaned records (NULL foreign keys)
		db.Exec("INSERT INTO jobs (type, state) VALUES (?, ?)", model.Pack, model.Complete)
		db.Exec("INSERT INTO cars (piece_size) VALUES (1024)")
		db.Exec("INSERT INTO directories (name) VALUES ('orphan')")
		db.Exec("INSERT INTO files (path, size) VALUES ('orphan.txt', 100)")
		db.Exec("INSERT INTO car_blocks (car_offset) VALUES (0)")

		// Verify records exist
		var counts [5]int64
		db.Model(&model.Job{}).Count(&counts[0])
		db.Model(&model.Car{}).Count(&counts[1])
		db.Model(&model.Directory{}).Count(&counts[2])
		db.Model(&model.File{}).Count(&counts[3])
		db.Model(&model.CarBlock{}).Count(&counts[4])
		for i, c := range counts {
			require.Equal(t, int64(1), c, "table %d should have 1 record", i)
		}

		cleanupOrphanedRecords(ctx, db)

		// Verify all orphaned records deleted
		db.Model(&model.Job{}).Count(&counts[0])
		db.Model(&model.Car{}).Count(&counts[1])
		db.Model(&model.Directory{}).Count(&counts[2])
		db.Model(&model.File{}).Count(&counts[3])
		db.Model(&model.CarBlock{}).Count(&counts[4])
		for i, c := range counts {
			require.Equal(t, int64(0), c, "table %d should be empty", i)
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
