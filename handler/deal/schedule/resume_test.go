package schedule

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestResumeHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			Actors: []model.Actor{{
				ID: "f01",
			}},
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.Schedule{
			PreparationID: 1,
			State:         model.SchedulePaused,
		}).Error
		require.NoError(t, err)
		schedule, err := Default.ResumeHandler(ctx, db, 1)
		require.NoError(t, err)
		require.Equal(t, model.ScheduleActive, schedule.State)

		_, err = Default.ResumeHandler(ctx, db, 1)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		_, err = Default.ResumeHandler(ctx, db, 2)
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}
