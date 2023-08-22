package datasetworker

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/data-preservation-programs/singularity/pack/daggen"
	"github.com/google/uuid"
	"github.com/gotidy/ptr"
	"github.com/ipfs/boxo/util"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	"github.com/stretchr/testify/require"
)

func TestExportDag(t *testing.T) {
	tmp := t.TempDir()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
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
	err = thread.dbNoContext.Create(&job).Error
	require.NoError(t, err)

	ctx := context.Background()
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
		{AttachmentID: 1, Data: dir1Data, CID: model.CID(pack.EmptyFileCid)},
		{AttachmentID: 1, Data: dir2Data, CID: model.CID(pack.EmptyFileCid), ParentID: ptr.Of(uint64(1)), Name: "sub"},
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
}

func TestExportDag_WithOutputStorage(t *testing.T) {
	tmp := t.TempDir()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
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
			},
			StorageID: 1,
		},
	}
	err = thread.dbNoContext.Create(&job).Error
	require.NoError(t, err)

	ctx := context.Background()
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
		{AttachmentID: 1, Data: dir1Data, CID: model.CID(pack.EmptyFileCid)},
		{AttachmentID: 1, Data: dir2Data, CID: model.CID(pack.EmptyFileCid), ParentID: ptr.Of(uint64(1)), Name: "sub"},
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
}
