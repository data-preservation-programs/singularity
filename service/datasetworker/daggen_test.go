package datasetworker

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/daggen"
	"github.com/data-preservation-programs/singularity/pack/packutil"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/google/uuid"
	"github.com/gotidy/ptr"
	"github.com/ipfs/boxo/util"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestExportDag_DagNotEnabled(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		thread := &Thread{
			id:          uuid.New(),
			dbNoContext: db,
			logger:      log.Logger("test").With("test", true),
		}
		job := model.Job{
			Type:  model.DagGen,
			State: model.Ready,
			Attachment: &model.SourceAttachment{
				Preparation: &model.Preparation{
					NoDag: true,
				},
				Storage: &model.Storage{
					Type: "local",
				},
			},
		}
		err := thread.dbNoContext.Create(&job).Error
		require.NoError(t, err)

		err = thread.ExportDag(ctx, job)
		require.ErrorIs(t, err, ErrDagDisabled)
	})
}

func TestExportDag(t *testing.T) {
	tmp := t.TempDir()
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		thread := &Thread{
			id:          uuid.New(),
			dbNoContext: db,
			logger:      log.Logger("test").With("test", true),
		}
		job := model.Job{
			Type:  model.DagGen,
			State: model.Ready,
			Attachment: &model.SourceAttachment{
				Preparation: &model.Preparation{},
				Storage: &model.Storage{
					Type: "local",
					Path: tmp,
				},
			},
		}
		err := thread.dbNoContext.Create(&job).Error
		require.NoError(t, err)

		dir1 := daggen.NewDirectoryData()
		dir1.AddBlocks(ctx, []blocks.Block{
			blocks.NewBlock([]byte("hello")),
			daggen.NewDummyNode(5, cid.NewCidV1(cid.Raw, util.Hash([]byte("world")))),
		})
		dir1Data, err := dir1.MarshalBinary(ctx)
		require.NoError(t, err)
		dir2 := daggen.NewDirectoryData()
		dir2Data, err := dir2.MarshalBinary(ctx)
		require.NoError(t, err)

		dirs := []model.Directory{
			{AttachmentID: 1, Data: dir1Data, CID: model.CID(packutil.EmptyFileCid)},
			{AttachmentID: 1, Data: dir2Data, CID: model.CID(packutil.EmptyFileCid), ParentID: ptr.Of(model.DirectoryID(1)), Name: "sub"},
		}
		err = thread.dbNoContext.Create(&dirs).Error
		require.NoError(t, err)

		err = thread.ExportDag(ctx, job)
		require.NoError(t, err)

		var carBlocks []model.CarBlock
		var cars []model.Car
		err = db.Find(&carBlocks).Error
		require.NoError(t, err)
		err = db.Find(&cars).Error
		require.NoError(t, err)
		dirs = []model.Directory{}
		err = db.Find(&dirs).Error
		require.NoError(t, err)
		require.Len(t, carBlocks, 3)
		require.Len(t, cars, 1)
		for _, dir := range dirs {
			require.True(t, dir.Exported)
		}
	})
}

func TestExportDag_WithOutputStorage_NoInline(t *testing.T) {
	tmp := t.TempDir()
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		thread := &Thread{
			id:          uuid.New(),
			dbNoContext: db,
			logger:      log.Logger("test").With("test", true),
		}
		job := model.Job{
			Type:  model.DagGen,
			State: model.Ready,
			Attachment: &model.SourceAttachment{
				Preparation: &model.Preparation{
					OutputStorages: []model.Storage{
						{
							Type: "local",
							Path: tmp,
						},
					},
					NoInline: true,
				},
				StorageID: 1,
			},
		}
		err := thread.dbNoContext.Create(&job).Error
		require.NoError(t, err)

		dir1 := daggen.NewDirectoryData()
		dir1.AddBlocks(ctx, []blocks.Block{
			blocks.NewBlock([]byte("hello")),
			daggen.NewDummyNode(5, cid.NewCidV1(cid.Raw, util.Hash([]byte("world")))),
		})
		dir1Data, err := dir1.MarshalBinary(ctx)
		require.NoError(t, err)
		dir2 := daggen.NewDirectoryData()
		dir2Data, err := dir2.MarshalBinary(ctx)
		require.NoError(t, err)

		dirs := []model.Directory{
			{AttachmentID: 1, Data: dir1Data, CID: model.CID(packutil.EmptyFileCid)},
			{AttachmentID: 1, Data: dir2Data, CID: model.CID(packutil.EmptyFileCid), ParentID: ptr.Of(model.DirectoryID(1)), Name: "sub"},
		}
		err = thread.dbNoContext.Create(&dirs).Error
		require.NoError(t, err)

		err = thread.ExportDag(ctx, job)
		require.NoError(t, err)

		var carBlocks []model.CarBlock
		var cars []model.Car
		err = db.Find(&carBlocks).Error
		require.NoError(t, err)
		err = db.Find(&cars).Error
		require.NoError(t, err)
		dirs = []model.Directory{}
		err = db.Find(&dirs).Error
		require.NoError(t, err)
		// no inline
		require.Len(t, carBlocks, 0)
		require.Len(t, cars, 1)
		for _, dir := range dirs {
			require.True(t, dir.Exported)
		}
	})
}
