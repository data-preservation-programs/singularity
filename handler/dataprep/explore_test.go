package dataprep

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

func TestExploreHandler_StorageNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, err := Default.ExploreHandler(ctx, db, "1", "source", "path")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestExploreHandler_PrepNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Storage{}).Error
		require.NoError(t, err)
		_, err = Default.ExploreHandler(ctx, db, "1", "source", "path")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestExploreHandler(t *testing.T) {
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
				err = db.Create([]model.Directory{
					{
						AttachmentID: 1,
					},
					{
						AttachmentID: 1,
						ParentID:     ptr.Of(model.DirectoryID(1)),
						Name:         "sub1",
					},
					{
						AttachmentID: 1,
						ParentID:     ptr.Of(model.DirectoryID(2)),
						Name:         "sub2",
					},
				}).Error
				require.NoError(t, err)
				err = db.Create([]model.File{
					{
						Path:             "sub1/test1.tst",
						Hash:             "hash",
						Size:             100,
						LastModifiedNano: 100,
						AttachmentID:     1,
						DirectoryID:      ptr.Of(model.DirectoryID(2)),
					},
					{
						Path:             "sub1/test2.tst",
						Hash:             "hash",
						Size:             100,
						LastModifiedNano: 100,
						AttachmentID:     1,
						DirectoryID:      ptr.Of(model.DirectoryID(2)),
					},
					{
						Path:             "sub1/test1.tst",
						Hash:             "hash2",
						Size:             200,
						LastModifiedNano: 200,
						AttachmentID:     1,
						DirectoryID:      ptr.Of(model.DirectoryID(2)),
					},
				}).Error
				require.NoError(t, err)
				result, err := Default.ExploreHandler(ctx, db, name, "source", "sub1")
				require.NoError(t, err)
				require.Len(t, result.SubEntries, 3)
			})
		})
	}
}
