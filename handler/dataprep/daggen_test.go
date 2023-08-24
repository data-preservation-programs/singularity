package dataprep

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestStartDagGenHandler_StorageNotFound(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{}).Error
	require.NoError(t, err)
	_, err = StartDagGenHandler(ctx, db, 1, "not found")
	require.ErrorIs(t, err, handlererror.ErrNotFound)
}

func TestPauseDagGenHandler_NoJob(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		SourceStorages: []model.Storage{{
			Name: "source",
		}},
	}).Error
	require.NoError(t, err)
	_, err = PauseDagGenHandler(ctx, db, 1, "source")
	require.ErrorIs(t, err, handlererror.ErrNotFound)
}
