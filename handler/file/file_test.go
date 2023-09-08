package file

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestGetFile_NotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, err := Default.GetFileHandler(ctx, db, 1)
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestGetFile(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tmpdir := t.TempDir()
		file := model.File{
			Path: "test.txt",
			Attachment: &model.SourceAttachment{
				Preparation: &model.Preparation{
					Name: "prep",
				},
				Storage: &model.Storage{
					Name: "source",
					Type: "local",
					Path: tmpdir,
				},
			},
			FileRanges: []model.FileRange{{
				Job: &model.Job{
					AttachmentID: 1,
				},
			}},
		}
		err := db.Create(&file).Error
		require.NoError(t, err)

		found, err := Default.GetFileHandler(ctx, db, 1)
		require.NoError(t, err)
		require.NotNil(t, found)
		require.NotZero(t, found.ID)
	})
}
