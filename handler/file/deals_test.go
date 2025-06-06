package file

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestGetFileDealsHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		testCid1 := cid.NewCidV1(cid.Raw, util.Hash([]byte("test1")))
		testCid2 := cid.NewCidV1(cid.Raw, util.Hash([]byte("test2")))
		tmpdir := t.TempDir()
		attachment := model.SourceAttachment{
			Preparation: &model.Preparation{
				Name: "prep",
			},
			Storage: &model.Storage{
				Name: "source",
				Type: "local",
				Path: tmpdir,
			},
		}
		err := db.Create(&attachment).Error
		require.NoError(t, err)
		jobs := []model.Job{
			{
				AttachmentID: attachment.ID,
			},
			{
				AttachmentID: attachment.ID,
			},
			{
				AttachmentID: attachment.ID,
			},
		}
		err = db.Create(&jobs).Error
		require.NoError(t, err)
		file := model.File{
			Path:         "test.txt",
			AttachmentID: attachment.ID,
			FileRanges: []model.FileRange{
				{
					JobID: ptr.Of(model.JobID(1)),
				},
				{
					JobID: ptr.Of(model.JobID(1)),
				},
				{
					JobID: ptr.Of(model.JobID(2)),
				},
				{
					JobID: ptr.Of(model.JobID(3)),
				},
			},
		}
		err = db.Create(&file).Error
		require.NoError(t, err)
		cars := []model.Car{{
			JobID:         ptr.Of(model.JobID(1)),
			PieceCID:      model.CID(testCid1),
			PreparationID: 1,
		}, {
			JobID:         ptr.Of(model.JobID(2)),
			PieceCID:      model.CID(testCid2),
			PreparationID: 1,
		}}
		err = db.Create(cars).Error
		require.NoError(t, err)

		wallet := &model.Wallet{ActorID: "f01", Address: "f11"}
		err = db.Create(wallet).Error
		require.NoError(t, err)
		deals := []model.Deal{{
			PieceCID: model.CID(testCid1),
			Wallet:   wallet,
		}, {
			PieceCID: model.CID(testCid2),
			Wallet:   wallet,
		}, {
			PieceCID: model.CID(testCid2),
			Wallet:   wallet,
		}}
		err = db.Create(deals).Error
		require.NoError(t, err)
		result, err := Default.GetFileDealsHandler(ctx, db, 1)
		require.NoError(t, err)
		require.Len(t, result, 4)
		require.Len(t, result[0].Deals, 1)
		require.Len(t, result[1].Deals, 1)
		require.Len(t, result[2].Deals, 2)
		require.Len(t, result[3].Deals, 0)
	})
}
