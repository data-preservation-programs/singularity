package scan

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/rjNemo/underscore"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestScan(t *testing.T) {
	tmp := t.TempDir()
	files := map[string]int{
		"empty.bin":    0,
		"1.bin":        1,
		"1k.bin":       1 << 10,
		"1m.bin":       1 << 20,
		"16m.bin":      16 << 20,
		"1/2/3/10.bin": 10,
		"1/2/3/11.bin": 11,
		"1/2/3 1.bin":  31,
		"1/2/32.bin":   32,
	}
	for path, size := range files {
		err := os.MkdirAll(filepath.Join(tmp, filepath.Dir(path)), 0755)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(tmp, path), testutil.GenerateRandomBytes(size), 0644)
		require.NoError(t, err)
	}

	// Create empty folder
	err := os.MkdirAll(filepath.Join(tmp, "emptyfolder"), 0755)
	require.NoError(t, err)

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		job := model.Job{
			Type:  model.Scan,
			State: model.Ready,
			Attachment: &model.SourceAttachment{
				Preparation: &model.Preparation{
					MaxSize: 2_000_000,
				},
				Storage: &model.Storage{
					Type: "local",
					Path: tmp,
				},
			},
		}
		err := db.Create(&job).Error
		require.NoError(t, err)
		dir := model.Directory{
			AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
		}
		err = db.Create(&dir).Error
		require.NoError(t, err)
		err = Scan(ctx, db, *job.Attachment)
		require.NoError(t, err)

		var dirs []model.Directory
		err = db.Find(&dirs).Error
		require.NoError(t, err)
		require.Len(t, dirs, 5)
		require.True(t, underscore.Any(dirs, func(dir model.Directory) bool {
			return dir.Name == "emptyfolder"
		}))
		var jobs []model.Job
		err = db.Preload("FileRanges").Find(&jobs).Error
		require.NoError(t, err)
		require.Len(t, jobs, 13)
	})
}
