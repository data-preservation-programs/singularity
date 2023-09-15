package storage

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestListStoragesHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tmp := t.TempDir()
		_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "name", tmp, nil, model.ClientConfig{}})
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
		storages, err := Default.ListStoragesHandler(ctx, db)
		require.NoError(t, err)
		require.Len(t, storages, 1)
		require.Len(t, storages[0].PreparationsAsSource, 1)
		require.Len(t, storages[0].PreparationsAsOutput, 1)
	})
}
