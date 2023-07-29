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

func TestHealthCheck(t *testing.T) {
	req := require.New(t)
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	defer database.DropAll(db)

	id := uuid.New()
	_, err = Register(context.Background(), db, id, func() State {
		return State{
			WorkType:  model.Packing,
			WorkingOn: "something",
		}
	}, true)
	req.NoError(err)

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
