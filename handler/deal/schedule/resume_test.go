package schedule

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestResumeHandler(t *testing.T) {
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
		State:         model.SchedulePaused,
	}).Error
	require.NoError(t, err)
	schedule, err := ResumeHandler(context.Background(), db, 1)
	require.NoError(t, err)
	require.Equal(t, model.ScheduleActive, schedule.State)

	_, err = ResumeHandler(context.Background(), db, 1)
	require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	_, err = ResumeHandler(context.Background(), db, 2)
	require.ErrorIs(t, err, handlererror.ErrNotFound)
}
