package schedule

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestRemoveSchedule_Success(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			Wallets: []model.Wallet{{
				ID: "f01",
			}},
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.Schedule{
			PreparationID: 1,
			State:         model.ScheduleCompleted,
		}).Error
		require.NoError(t, err)

		err = db.Create(&model.Deal{
			ClientID:   "f01",
			ScheduleID: ptr.Of(model.ScheduleID(1)),
		}).Error
		require.NoError(t, err)

		err = Default.RemoveHandler(ctx, db, 1)
		require.NoError(t, err)

		var count int64
		err = db.Model(&model.Schedule{}).Count(&count).Error
		require.NoError(t, err)
		require.Zero(t, count)

		err = db.Model(&model.Deal{}).Count(&count).Error
		require.NoError(t, err)
		require.NotZero(t, count)
	})
}

func TestRemoveSchedule_NotExist(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := Default.RemoveHandler(ctx, db, 1)
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestRemoveSchedule_StillActive(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
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
		err = Default.RemoveHandler(ctx, db, 1)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})
}
