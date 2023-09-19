//go:build !windows

package pack

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestAssembleCar_InaccessibleFile(t *testing.T) {
	tmp := t.TempDir()
	err := os.WriteFile(filepath.Join(tmp, "test.txt"), []byte("test"), 0000)
	require.NoError(t, err)
	stat, err := os.Stat(filepath.Join(tmp, "test.txt"))
	require.NoError(t, err)

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		job := model.Job{
			Type:  model.Pack,
			State: model.Processing,
			Attachment: &model.SourceAttachment{
				Preparation: &model.Preparation{
					MaxSize:   2000000,
					PieceSize: 1 << 21,
				},
				Storage: &model.Storage{
					Type: "local",
					Path: tmp,
				},
			},
			FileRanges: []model.FileRange{
				{
					Offset: 0,
					Length: 5,
					File: &model.File{
						Path:             "test.txt",
						Size:             stat.Size(),
						LastModifiedNano: stat.ModTime().UnixNano(),
						AttachmentID:     1,
						Directory: &model.Directory{
							AttachmentID: 1,
						},
					},
				},
			},
		}
		err := db.Create(&job).Error
		require.NoError(t, err)
		_, err = Pack(ctx, db, job)
		require.ErrorContains(t, err, "failed to open")
	})

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		job := model.Job{
			Type:  model.Pack,
			State: model.Processing,
			Attachment: &model.SourceAttachment{
				Preparation: &model.Preparation{
					MaxSize:   2000000,
					PieceSize: 1 << 21,
				},
				Storage: &model.Storage{
					Type: "local",
					Path: tmp,
					ClientConfig: model.ClientConfig{
						SkipInaccessibleFile: ptr.Of(true),
					},
				},
			},
			FileRanges: []model.FileRange{
				{
					Offset: 0,
					Length: 5,
					File: &model.File{
						Path:             "test.txt",
						Size:             stat.Size(),
						LastModifiedNano: stat.ModTime().UnixNano(),
						AttachmentID:     1,
						Directory: &model.Directory{
							AttachmentID: 1,
						},
					},
				},
			},
		}
		err := db.Create(&job).Error
		require.NoError(t, err)
		_, err = Pack(ctx, db, job)
		require.ErrorIs(t, err, ErrNoContent)
	})
}
