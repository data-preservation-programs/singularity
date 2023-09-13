package file

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

func TestPrepareToPackFileHandler_NotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, err := Default.PrepareToPackFileHandler(ctx, db, 1)
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestPrepareToPackFileHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		file1 := model.File{
			Size: 900000,
			Attachment: &model.SourceAttachment{
				Preparation: &model.Preparation{
					Name:      "prep",
					MaxSize:   1 << 20,
					PieceSize: 1 << 20,
				},
				Storage: &model.Storage{
					Name: "source",
					Type: "local",
					Path: t.TempDir(),
				},
			},
			Directory: &model.Directory{
				AttachmentID: 1,
			},
			FileRanges: []model.FileRange{{
				Offset: 0,
				Length: 900000,
			}},
		}
		err := db.Create(&file1).Error
		require.NoError(t, err)
		_, err = Default.PrepareToPackFileHandler(ctx, db, 1)
		require.NoError(t, err)
		var jobs []model.Job
		err = db.Find(&jobs).Error
		require.NoError(t, err)
		require.Len(t, jobs, 1)
		require.Equal(t, jobs[0].State, model.Created)

		file2 := model.File{
			Size:         900000,
			AttachmentID: 1,
			DirectoryID:  ptr.Of(model.DirectoryID(1)),
			FileRanges: []model.FileRange{{
				Offset: 0,
				Length: 900000,
			}},
		}
		err = db.Create(&file2).Error
		require.NoError(t, err)
		_, err = Default.PrepareToPackFileHandler(ctx, db, 2)
		require.NoError(t, err)
		err = db.Find(&jobs).Error
		require.NoError(t, err)
		require.Len(t, jobs, 2)
		require.Equal(t, model.Ready, jobs[0].State)
		require.Equal(t, model.Created, jobs[1].State)
	})
}
