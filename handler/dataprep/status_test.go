package dataprep

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestGetStatusHandler(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	ctx := context.Background()

	err = db.Create(&model.Preparation{
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

	status, err := GetStatusHandler(ctx, db, 1)
	require.NoError(t, err)
	require.Len(t, status.Sources, 1)
	require.Len(t, status.Sources[0].Jobs, 1)
	require.Equal(t, model.Ready, status.Sources[0].Jobs[0].State)
}

func TestGetStatusHandler_NotFound(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	ctx := context.Background()

	_, err = GetStatusHandler(ctx, db, 1)
	require.ErrorIs(t, err, handlererror.ErrNotFound)
}
