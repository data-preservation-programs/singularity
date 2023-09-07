package dataprep

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestListSchedulesHandler_NotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, err := Default.ListSchedulesHandler(ctx, db, "1")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestListSchedulesHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
			OutputStorages: []model.Storage{{
				Name: "output",
			}},
		}).Error
		require.NoError(t, err)

		err = db.Create(&model.Schedule{
			PreparationID: 1,
		}).Error
		require.NoError(t, err)

		schedules, err := Default.ListSchedulesHandler(ctx, db, "1")
		require.NoError(t, err)
		require.Len(t, schedules, 1)
	})
}
