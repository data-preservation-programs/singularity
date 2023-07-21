//go:build !windows

package healthcheck

import (
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestHealthCheck(t *testing.T) {
	req := require.New(t)
	db, err := database.OpenInMemory()
	require.NoError(t, err)
	defer database.DropAll(db)

	id := uuid.New()
	HealthCheck(db, id, func() State {
		return State{
			WorkType:  model.Packing,
			WorkingOn: "something",
		}
	})
	var worker model.Worker
	err = db.Where("id = ?", id.String()).First(&worker).Error
	req.Nil(err)
	req.Equal(model.Packing, worker.WorkType)
	req.Equal("something", worker.WorkingOn)
	req.NotEmpty(worker.Hostname)
	lastHeatbeat := worker.LastHeartbeat

	HealthCheck(db, id, func() State {
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
