package cmd

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/handler/job"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/google/uuid"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func swapJobHandler(mockHandler job.Handler) func() {
	actual := job.Default
	job.Default = mockHandler
	return func() {
		job.Default = actual
	}
}

var testDagGenJob = model.Job{
	ID:           1,
	Type:         model.DagGen,
	State:        model.Ready,
	WorkerID:     nil,
	AttachmentID: 1,
}

func TestDataPrepStartDagGenHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(job.MockJob)
		defer swapJobHandler(mockHandler)()

		mockHandler.On("StartDagGenHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testDagGenJob, nil)
		_, _, err := runner.Run(ctx, "singularity prep start-daggen 1 name")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep start-daggen 1 name")
		require.NoError(t, err)
	})
}

func TestDataPrepPauseDagGenHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(job.MockJob)
		defer swapJobHandler(mockHandler)()

		mockHandler.On("PauseDagGenHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testDagGenJob, nil)
		_, _, err := runner.Run(ctx, "singularity prep pause-daggen 1 name")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep pause-daggen 1 name")
		require.NoError(t, err)
	})
}

var testScanJob = model.Job{
	ID:           1,
	Type:         model.Scan,
	State:        model.Ready,
	WorkerID:     nil,
	AttachmentID: 1,
}

func TestDataPrepStartScanHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(job.MockJob)
		defer swapJobHandler(mockHandler)()

		mockHandler.On("StartScanHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testScanJob, nil)
		_, _, err := runner.Run(ctx, "singularity prep start-scan 1 name")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep start-scan 1 name")
		require.NoError(t, err)
	})
}

func TestDataPrepPauseScanHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(job.MockJob)
		defer swapJobHandler(mockHandler)()

		mockHandler.On("PauseScanHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testScanJob, nil)
		_, _, err := runner.Run(ctx, "singularity prep pause-scan 1 name")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep pause-scan 1 name")
		require.NoError(t, err)
	})
}

var testPackJob = model.Job{
	ID:           1,
	Type:         model.Pack,
	State:        model.Ready,
	WorkerID:     nil,
	AttachmentID: 1,
}

func TestDataPrepStartPackHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(job.MockJob)
		defer swapJobHandler(mockHandler)()

		mockHandler.On("StartPackHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]model.Job{testPackJob}, nil)
		_, _, err := runner.Run(ctx, "singularity prep start-pack 1 name")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep start-pack 1 name")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity prep start-pack 1 name 1")
		require.NoError(t, err)
		_, _, err = runner.Run(ctx, "singularity --verbose prep start-pack 1 name 1")
		require.NoError(t, err)
	})
}

func TestDataPrepPausePackHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(job.MockJob)
		defer swapJobHandler(mockHandler)()

		mockHandler.On("PausePackHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]model.Job{testPackJob}, nil)
		_, _, err := runner.Run(ctx, "singularity prep pause-pack 1 name")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep pause-pack 1 name")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity prep pause-pack 1 name 1")
		require.NoError(t, err)
		_, _, err = runner.Run(ctx, "singularity --verbose prep pause-pack 1 name 1")
		require.NoError(t, err)
	})
}

func TestDataPreparationGetStatusHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(job.MockJob)
		defer swapJobHandler(mockHandler)()

		mockHandler.On("GetStatusHandler", mock.Anything, mock.Anything, mock.Anything).Return([]job.SourceStatus{
			{
				AttachmentID:    ptr.Of(uint32(1)),
				SourceStorageID: ptr.Of(uint32(1)),
				SourceStorage: &model.Storage{
					ID:        1,
					Name:      "source",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					Type:      "local",
					Path:      "/tmp",
				},
				Jobs: []model.Job{
					{
						ID:           1,
						Type:         model.Pack,
						State:        model.Processing,
						WorkerID:     ptr.Of(uuid.NewString()),
						AttachmentID: 1,
					},
				},
			},
		}, nil)
		_, _, err := runner.Run(ctx, "singularity prep status 1")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep status 1")
		require.NoError(t, err)
	})
}
