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

func TestGetStatusHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)

		err = db.Create(&model.Job{
			AttachmentID: 1,
			State:        model.Ready,
			Type:         model.Pack,
		}).Error
		require.NoError(t, err)

		status, err := Default.GetStatusHandler(ctx, db, 1)
		require.NoError(t, err)
		require.Len(t, status, 1)
		require.Len(t, status[0].Jobs, 1)
		require.Equal(t, model.Ready, status[0].Jobs[0].State)
	})
}

func TestGetStatusHandler_NotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, err := Default.GetStatusHandler(ctx, db, 1)
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}
