package push

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/config/configmap"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestMaxSizeToSplitSize(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected int64
	}{
		{
			name:     "Input is 1G",
			input:    1 << 30,
			expected: 1 << 28,
		},
		{
			name:     "Input is 4G",
			input:    1 << 32,
			expected: 1 << 30,
		},
		{
			name:     "Input is 32G",
			input:    1 << 35,
			expected: 1 << 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, MaxSizeToSplitSize(tt.input))
		})
	}
}

func TestPushFile(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		cache := map[string]model.DirectoryID{"": 1, ".": 1}
		attachment := model.SourceAttachment{
			Preparation: &model.Preparation{
				MaxSize: 16,
			},
			Storage: &model.Storage{},
		}
		err := db.Create(&attachment).Error
		require.NoError(t, err)
		root := model.Directory{
			AttachmentID: &attachment.ID,
		}
		err = db.Create(&root).Error
		require.NoError(t, err)
		tmp := t.TempDir()
		err = os.WriteFile(filepath.Join(tmp, "test.txt"), []byte("hello world"), 0644)
		require.NoError(t, err)
		_ = storagesystem.Backends
		backend, err := fs.Find("local")
		require.NoError(t, err)
		f, err := backend.NewFs(ctx, "local", tmp, make(configmap.Simple))
		require.NoError(t, err)
		obj, err := f.NewObject(ctx, "test.txt")
		require.NoError(t, err)

		// First push
		file, fileRanges, err := PushFile(ctx, db, obj, attachment, cache)
		require.NoError(t, err)
		require.Equal(t, "test.txt", file.Path)
		require.EqualValues(t, 11, file.Size)
		require.Len(t, fileRanges, 3)

		// Second push with same file
		file, fileRanges, err = PushFile(ctx, db, obj, attachment, cache)
		require.NoError(t, err)
		require.Nil(t, file)
		require.Nil(t, fileRanges)
	})
}

func TestEnsureParentDirectories(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		cache := map[string]model.DirectoryID{}
		attachment := model.SourceAttachment{
			Preparation: &model.Preparation{},
			Storage:     &model.Storage{},
		}
		err := db.Create(&attachment).Error
		require.NoError(t, err)
		root := model.Directory{
			AttachmentID: &attachment.ID,
		}
		err = db.Create(&root).Error
		require.NoError(t, err)
		err = EnsureParentDirectories(ctx, db, &model.File{
			Path:         "sub1/sub2/sub3/sub4/test.txt",
			AttachmentID: &attachment.ID,
		}, 1, cache)
		require.NoError(t, err)
		var dirs []model.Directory
		err = db.Find(&dirs).Error
		require.NoError(t, err)
		require.Len(t, dirs, 5)

		err = EnsureParentDirectories(ctx, db, &model.File{
			Path:         "sub1/sub2/c/d/test.txt",
			AttachmentID: &attachment.ID,
		}, 1, cache)
		require.NoError(t, err)

		err = EnsureParentDirectories(ctx, db, &model.File{
			Path:         "x/y/z/test.txt",
			AttachmentID: &attachment.ID,
		}, 1, cache)
		require.NoError(t, err)
		err = db.Find(&dirs).Error
		require.NoError(t, err)
		require.Len(t, dirs, 10)
	})
}

func TestCreatePackJob(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		attachment := model.SourceAttachment{
			Preparation: &model.Preparation{},
			Storage:     &model.Storage{},
		}
		err := db.Create(&attachment).Error
		require.NoError(t, err)
		fileRanges := []model.FileRange{
			{
				File: &model.File{
					AttachmentID: &attachment.ID,
				},
			},
		}
		err = db.Create(&fileRanges).Error
		require.NoError(t, err)
		job, err := CreatePackJob(ctx, db, attachment.ID, []model.FileRangeID{1})
		require.NoError(t, err)
		require.Equal(t, attachment.ID, *job.AttachmentID)
	})
}
