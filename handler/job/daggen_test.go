package job

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestStartDagGenHandler_StorageNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			Name: "name",
		}).Error
		require.NoError(t, err)
		_, err = Default.StartDagGenHandler(ctx, db, "1", "not found")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestPauseDagGenHandler_NoJob(t *testing.T) {
	for _, name := range []string{"1", "name"} {
		t.Run(name, func(t *testing.T) {
			testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
				err := db.Create(&model.Preparation{
					Name: "name",
					SourceStorages: []model.Storage{{
						Name: "source",
					}},
				}).Error
				require.NoError(t, err)
				_, err = Default.PauseDagGenHandler(ctx, db, name, "source")
				require.ErrorIs(t, err, handlererror.ErrNotFound)
			})
		})
	}
}

func TestStartDagGenHandler_DagDisabled(t *testing.T) {
	for _, name := range []string{"1", "name"} {
		t.Run(name, func(t *testing.T) {
			testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
				err := db.Create(&model.Preparation{
					Name: "name",
					SourceStorages: []model.Storage{{
						Name: "source",
					}},
					NoDag: true,
				}).Error
				require.NoError(t, err)
				_, err = Default.StartDagGenHandler(ctx, db, "1", "source")
				require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
				require.ErrorContains(t, err, "dag generation is disabled")
			})
		})
	}
}
