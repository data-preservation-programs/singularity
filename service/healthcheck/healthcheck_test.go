package healthcheck

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestHealthCheckCleanup(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	ctx, cancel := context.WithCancel(context.Background())
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
}
func TestHealthCheckReport(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		StartReportHealth(ctx, db, uuid.New(), func() State {
			return State{
				WorkType:  model.Packing,
				WorkingOn: "something",
			}
		})
		close(done)
	}()
	time.Sleep(time.Second)
	cancel()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("healthcheck report didn't stop")
	}
}

func TestHealthCheck(t *testing.T) {
	req := require.New(t)
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	id := uuid.New()
	alreadyRunning, err := Register(context.Background(), db, id, func() State {
		return State{
			WorkType:  model.Packing,
			WorkingOn: "something",
		}
	}, false)
	req.NoError(err)
	req.False(alreadyRunning)
	alreadyRunning, err = Register(context.Background(), db, uuid.New(), func() State {
		return State{
			WorkType:  model.Packing,
			WorkingOn: "something",
		}
	}, false)
	req.NoError(err)
	req.True(alreadyRunning)

	var worker model.Worker
	err = db.Where("id = ?", id.String()).First(&worker).Error
	req.Nil(err)
	req.Equal(model.Packing, worker.WorkType)
	req.Equal("something", worker.WorkingOn)
	req.NotEmpty(worker.Hostname)
	lastHeatbeat := worker.LastHeartbeat

	time.Sleep(200 * time.Millisecond)
	ReportHealth(context.Background(), db, id, func() State {
		return State{
			WorkType:  model.Packing,
			WorkingOn: "something else",
		}
	})

	err = db.Where("id = ?", id.String()).First(&worker).Error
	req.Nil(err)
	req.Equal(model.Packing, worker.WorkType)
	req.Equal("something else", worker.WorkingOn)
	req.NotEmpty(worker.Hostname)
	req.NotEqual(lastHeatbeat, worker.LastHeartbeat)

	HealthCheckCleanup(db)
	err = db.Where("id = ?", id.String()).First(&worker).Error
	req.Nil(err)

	oldThreshold := staleThreshold
	staleThreshold = 0
	defer func() {
		staleThreshold = oldThreshold
	}()

	HealthCheckCleanup(db)
	err = db.Where("id = ?", id.String()).First(&worker).Error
	req.ErrorIs(err, gorm.ErrRecordNotFound)
}
