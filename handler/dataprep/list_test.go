package dataprep

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestListHandler(t *testing.T) {
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

		preparations, err := Default.ListHandler(ctx, db)
		require.NoError(t, err)
		require.Len(t, preparations, 1)
		require.Len(t, preparations[0].SourceStorages, 1)
		require.Len(t, preparations[0].OutputStorages, 1)
	})
}
