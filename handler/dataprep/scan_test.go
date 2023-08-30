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

func TestStartScanHandler_StorageNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{}).Error
		require.NoError(t, err)
		_, err = Default.StartScanHandler(ctx, db, 1, "not found")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestStartScanHandler_PreparationNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)
		_, err = Default.StartScanHandler(ctx, db, 2, "source")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestStartScanHandler_NewScanJob(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)
		job, err := Default.StartScanHandler(ctx, db, 1, "source")
		require.NoError(t, err)
		require.EqualValues(t, 1, job.ID)
		require.Equal(t, model.Ready, job.State)
		require.Equal(t, model.Scan, job.Type)
	})
}

func TestStartScanHandler_StartExisting(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.Job{
			AttachmentID: 1,
			State:        model.Error,
			Type:         model.Scan,
		}).Error
		require.NoError(t, err)
		job, err := Default.StartScanHandler(ctx, db, 1, "source")
		require.NoError(t, err)
		require.Equal(t, model.Ready, job.State)
		require.Equal(t, model.Scan, job.Type)
		require.EqualValues(t, 1, job.ID)
	})
}

func TestStartScanHandler_AlreadyProcessing(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.Job{
			AttachmentID: 1,
			State:        model.Processing,
			Type:         model.Scan,
		}).Error
		require.NoError(t, err)
		_, err = Default.StartScanHandler(ctx, db, 1, "source")
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})
}

func TestPauseScanHandler_NoJob(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)
		_, err = Default.PauseScanHandler(ctx, db, 1, "source")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestPauseScanHandler_JobNotRunning(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)
		require.NoError(t, err)
		err = db.Create(&model.Job{
			AttachmentID: 1,
			State:        model.Error,
			Type:         model.Scan,
		}).Error
		require.NoError(t, err)

		_, err = Default.PauseScanHandler(ctx, db, 1, "source")
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})
}

func TestPauseScanHandler_JobPaused(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)
		require.NoError(t, err)
		err = db.Create(&model.Job{
			AttachmentID: 1,
			State:        model.Processing,
			Type:         model.Scan,
		}).Error
		require.NoError(t, err)

		job, err := Default.PauseScanHandler(ctx, db, 1, "source")
		require.NoError(t, err)
		require.Equal(t, model.Paused, job.State)
	})
}
