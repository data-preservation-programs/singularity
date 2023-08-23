package datasetworker

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestDatasetWorker_ExitOnComplete(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	worker := NewWorker(db, Config{
		Concurrency:    2,
		ExitOnComplete: true,
		EnableScan:     true,
		EnablePack:     true,
		EnableDag:      true,
		ExitOnError:    true,
	})

	job := model.Job{
		Type:  model.Scan,
		State: model.Ready,
		Attachment: &model.SourceAttachment{
			Preparation: &model.Preparation{},
			Storage: &model.Storage{
				Type: "local",
				Path: t.TempDir(),
			},
		},
	}
	err = db.Create(&job).Error
	require.NoError(t, err)
	dir := model.Directory{
		AttachmentID: 1,
	}
	err = db.Create(&dir).Error
	require.NoError(t, err)

	err = worker.Run(context.Background())
	require.NoError(t, err)
}

func TestDatasetWorker_ExitOnError(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	worker := NewWorker(db, Config{
		Concurrency:    2,
		ExitOnComplete: true,
		EnableScan:     true,
		EnablePack:     true,
		EnableDag:      true,
		ExitOnError:    true,
	})

	tmp := t.TempDir()
	job := model.Job{
		Type:  model.DagGen,
		State: model.Ready,
		Attachment: &model.SourceAttachment{
			Preparation: &model.Preparation{},
			Storage: &model.Storage{
				Type: "local",
				Path: tmp,
			},
		},
	}
	err = db.Create(&job).Error
	require.NoError(t, err)
	err = worker.Run(context.Background())
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
}
