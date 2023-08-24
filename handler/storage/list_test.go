package storage

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestListStoragesHandler(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	tmp := t.TempDir()
	_, err = CreateStorageHandler(ctx, db, "local", "", "name", tmp, nil)
	require.NoError(t, err)
	err = db.Create(&model.Preparation{
		SourceStorages: []model.Storage{
			{
				ID: 1,
			},
		},
		OutputStorages: []model.Storage{
			{
				ID: 1,
			},
		},
	}).Error
	require.NoError(t, err)
	storages, err := ListStoragesHandler(ctx, db)
	require.NoError(t, err)
	require.Len(t, storages, 1)
	require.Len(t, storages[0].PreparationsAsSource, 1)
	require.Len(t, storages[0].PreparationsAsOutput, 1)
}
