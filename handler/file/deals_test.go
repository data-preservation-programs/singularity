package file

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestGetFileDealsHandler(t *testing.T) {
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
		car := model.Car{
			JobID:         ptr.Of(uint64(1)),
			PieceCID:      model.CID(testutil.TestCid),
			PreparationID: 1,
		}
		err = db.Create(&car).Error
		require.NoError(t, err)

		deal := model.Deal{
			PieceCID: model.CID(testutil.TestCid),
			Wallet:   &model.Wallet{},
		}
		err = db.Create(&deal).Error
		require.NoError(t, err)
		deals, err := Default.GetFileDealsHandler(ctx, db, "1")
		require.NoError(t, err)
		require.Len(t, deals, 1)
	})
}
