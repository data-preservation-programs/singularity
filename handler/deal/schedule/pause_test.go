package schedule

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestPauseHandler(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	err = db.Create(&model.Preparation{
		Wallets: []model.Wallet{{
			ID: "f01",
		}},
	}).Error
	require.NoError(t, err)
	err = db.Create(&model.Schedule{
		PreparationID: 1,
		State:         model.ScheduleActive,
	}).Error
	require.NoError(t, err)
	schedule, err := PauseHandler(context.Background(), db, 1)
	require.NoError(t, err)
	require.Equal(t, model.SchedulePaused, schedule.State)

	_, err = PauseHandler(context.Background(), db, 1)
	require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	_, err = PauseHandler(context.Background(), db, 2)
	require.ErrorIs(t, err, handlererror.ErrNotFound)
}
