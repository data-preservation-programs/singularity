package file

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestPushFileHandler_NotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, err := Default.PushFileHandler(ctx, db, "prep", "source", Info{})
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestPushFileHandler_FileNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tmpdir := t.TempDir()
		err := db.Create(&model.Preparation{
			Name:      "prep",
			MaxSize:   1 << 34,
			PieceSize: 1 << 35,
			SourceStorages: []model.Storage{{
				Name: "source",
				Type: "local",
				Path: tmpdir,
			}},
		}).Error
		require.NoError(t, err)
		_, err = Default.PushFileHandler(ctx, db, "prep", "source", Info{Path: "notexist"})
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "object not found")
	})
}

func TestPushFileHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tmpdir := t.TempDir()
		err := os.WriteFile(filepath.Join(tmpdir, "test.txt"), []byte("test"), 0o644)
		require.NoError(t, err)
		err = db.Create(&model.Preparation{
			Name:      "prep",
			MaxSize:   1 << 34,
			PieceSize: 1 << 35,
			SourceStorages: []model.Storage{{
				Name: "source",
				Type: "local",
				Path: tmpdir,
			}},
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.Directory{
			AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
		}).Error
		require.NoError(t, err)
		file, err := Default.PushFileHandler(ctx, db, "prep", "source", Info{Path: "test.txt"})
		require.NoError(t, err)
		require.NotNil(t, file)
		require.Len(t, file.FileRanges, 1)
	})
}
