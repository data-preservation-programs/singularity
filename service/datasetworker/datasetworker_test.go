package datasetworker

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/analytics"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func init() {
	analytics.Enabled = false
}

func TestDatasetWorker_ExitOnComplete(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		worker := NewWorker(db, Config{
			Concurrency:    2,
			ExitOnComplete: true,
			EnableScan:     true,
			EnablePack:     true,
			EnableDag:      true,
			ExitOnError:    true,
		})

		// Create preparation
		prep := model.Preparation{
			Name: "test-prep",
		}
		err := db.Create(&prep).Error
		require.NoError(t, err)

		// Create storage
		storage := model.Storage{
			Name: "test-storage",
			Type: "local",
			Path: t.TempDir(),
		}
		err = db.Create(&storage).Error
		require.NoError(t, err)

		// Create source attachment
		attachment := model.SourceAttachment{
			PreparationID: prep.ID,
			StorageID:     storage.ID,
		}
		err = db.Create(&attachment).Error
		require.NoError(t, err)

		// Create job referencing the attachment
		job := model.Job{
			Type:         model.Scan,
			State:        model.Ready,
			AttachmentID: attachment.ID,
		}
		err = db.Create(&job).Error
		require.NoError(t, err)

		// Create root directory for the attachment
		dir := model.Directory{
			AttachmentID: attachment.ID,
			Name:         "root",
			ParentID:     nil, // This makes it a root directory
		}
		err = db.Create(&dir).Error
		require.NoError(t, err)

		err = worker.Run(ctx)
		require.NoError(t, err)
	})
}

func TestDatasetWorker_ExitOnError(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		worker := NewWorker(db, Config{
			Concurrency:    1, // Use single worker to avoid race conditions
			ExitOnComplete: true,
			EnableScan:     false, // Disable scan to focus on DagGen
			EnablePack:     false, // Disable pack to focus on DagGen
			EnableDag:      true,
			ExitOnError:    true,
		})

		// Create preparation with NoDag=false (default) to allow DAG generation
		prep := model.Preparation{
			Name:  "test-prep-error",
			NoDag: false,
		}
		err := db.Create(&prep).Error
		require.NoError(t, err)

		// Create storage
		tmp := t.TempDir()
		storage := model.Storage{
			Name: "test-storage-error",
			Type: "local",
			Path: tmp,
		}
		err = db.Create(&storage).Error
		require.NoError(t, err)

		// Create source attachment
		attachment := model.SourceAttachment{
			PreparationID: prep.ID,
			StorageID:     storage.ID,
		}
		err = db.Create(&attachment).Error
		require.NoError(t, err)

		// Create job referencing the attachment (DagGen job)
		job := model.Job{
			Type:         model.DagGen,
			State:        model.Ready,
			AttachmentID: attachment.ID,
		}
		err = db.Create(&job).Error
		require.NoError(t, err)

		// Note: We intentionally do NOT create a root directory here
		// This should cause the RootDirectoryCID call to fail with record not found
		// which is what this test expects

		err = worker.Run(ctx)
		require.Error(t, err)
		// Check if the error contains gorm.ErrRecordNotFound in the error chain
		require.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})
}
