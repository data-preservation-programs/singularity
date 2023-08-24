package dataprep

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestListHandler(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		SourceStorages: []model.Storage{{
			Name: "source",
		}},
		OutputStorages: []model.Storage{{
			Name: "output",
		}},
	}).Error
	require.NoError(t, err)

	preparations, err := ListHandler(ctx, db)
	require.NoError(t, err)
	require.Len(t, preparations, 1)
	require.Len(t, preparations[0].SourceStorages, 1)
	require.Len(t, preparations[0].OutputStorages, 1)
}
