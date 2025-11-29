package job

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestGetStatusHandler(t *testing.T) {
	for _, name := range []string{"1", "name"} {
		t.Run(name, func(t *testing.T) {
			testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
				err := db.Create(&model.Preparation{
					Name: "name",
					SourceStorages: []model.Storage{{
						Name: "source",
					}},
					OutputStorages: []model.Storage{{
						Name: "output",
					}},
				}).Error
				require.NoError(t, err)

				err = db.Create(&model.Job{
					AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
					State:        model.Ready,
					Type:         model.Pack,
				}).Error
				require.NoError(t, err)

				status, err := Default.GetStatusHandler(ctx, db, name)
				require.NoError(t, err)
				require.Len(t, status, 1)
				require.Len(t, status[0].Jobs, 1)
				require.Len(t, status[0].OutputStorages, 1)
				require.Equal(t, model.Ready, status[0].Jobs[0].State)
			})
		})
	}
}

func TestGetStatusHandler_NotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, err := Default.GetStatusHandler(ctx, db, "1")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}
