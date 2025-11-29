package datasetworker

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/localstack"
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
			{AttachmentID: ptr.Of(model.SourceAttachmentID(1)), Data: dir1Data, CID: model.CID(packutil.EmptyFileCid)},
			{AttachmentID: ptr.Of(model.SourceAttachmentID(1)), Data: dir2Data, CID: model.CID(packutil.EmptyFileCid), ParentID: ptr.Of(model.DirectoryID(1)), Name: "sub"},
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
			{AttachmentID: ptr.Of(model.SourceAttachmentID(1)), Data: dir1Data, CID: model.CID(packutil.EmptyFileCid)},
			{AttachmentID: ptr.Of(model.SourceAttachmentID(1)), Data: dir2Data, CID: model.CID(packutil.EmptyFileCid), ParentID: ptr.Of(model.DirectoryID(1)), Name: "sub"},
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

func TestExportDag_WithMinPieceSize_LocalStorage(t *testing.T) {
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
					NoInline:    true,
					MinPieceSize: 512 * 1024, // 512 KiB - will force padding for small DAG
				},
				StorageID: 1,
			},
		}
		err := thread.dbNoContext.Create(&job).Error
		require.NoError(t, err)

		// Create a small DAG that will need padding
		dir1 := daggen.NewDirectoryData()
		dir1.AddBlocks(ctx, []blocks.Block{
			blocks.NewBlock([]byte("small data")),
		})
		dir1Data, err := dir1.MarshalBinary(ctx)
		require.NoError(t, err)

		attachmentID := model.SourceAttachmentID(1)
		dirs := []model.Directory{
			{AttachmentID: &attachmentID, Data: dir1Data, CID: model.CID(packutil.EmptyFileCid)},
		}
		err = thread.dbNoContext.Create(&dirs).Error
		require.NoError(t, err)

		err = thread.ExportDag(ctx, job)
		require.NoError(t, err)

		var cars []model.Car
		err = db.Find(&cars).Error
		require.NoError(t, err)
		require.Len(t, cars, 1)

		// Verify the piece size is at least the minPieceSize
		require.GreaterOrEqual(t, cars[0].PieceSize, int64(512*1024))

		// Verify the file size matches the expected padded size (127/128 of piece size)
		expectedFileSize := (cars[0].PieceSize * 127) / 128
		require.Equal(t, expectedFileSize, cars[0].FileSize)

		// Verify MinPieceSizePadding is 0 (padding is physical, not virtual)
		require.Equal(t, int64(0), cars[0].MinPieceSizePadding)
	})
}

func TestExportDag_WithMinPieceSize_RemoteStorage(t *testing.T) {

	// Set up localstack S3
	bucketName := "testbucket"
	tempDir := t.TempDir()

	// Create bucket directory for localstack
	err := os.MkdirAll(filepath.Join(tempDir, bucketName), 0777)
	require.NoError(t, err)

	p := localstack.Preset(
		localstack.WithServices(localstack.S3),
		localstack.WithS3Files(tempDir),
	)
	localS3, err := gnomock.Start(p)
	if err != nil && strings.HasPrefix(err.Error(), "can't start container") {
		t.Skip("Docker required for S3 tests")
	}
	require.NoError(t, err)
	defer func() { _ = gnomock.Stop(localS3) }()

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
							Type: "s3",
							Path: bucketName,
							Config: map[string]string{
								"region":            "us-east-1",
								"provider":          "Other",
								"force_path_style":  "true",
								"chunk_size":        "5Mi",
								"copy_cutoff":       "5Mi",
								"upload_cutoff":     "5Mi",
								"list_chunk":        "1000",
								"endpoint":          fmt.Sprint("http://", localS3.Address(localstack.APIPort)),
								"env_auth":          "false",
								"access_key_id":     "test",
								"secret_access_key": "test",
							},
						},
					},
					NoInline:     true,
					MinPieceSize: 512 * 1024, // 512 KiB - will force padding
				},
				StorageID: 1,
			},
		}
		err := thread.dbNoContext.Create(&job).Error
		require.NoError(t, err)

		// Create a small DAG that will need padding
		dir1 := daggen.NewDirectoryData()
		dir1.AddBlocks(ctx, []blocks.Block{
			blocks.NewBlock([]byte("small data for remote")),
		})
		dir1Data, err := dir1.MarshalBinary(ctx)
		require.NoError(t, err)

		attachmentID := model.SourceAttachmentID(1)
		dirs := []model.Directory{
			{AttachmentID: &attachmentID, Data: dir1Data, CID: model.CID(packutil.EmptyFileCid)},
		}
		err = thread.dbNoContext.Create(&dirs).Error
		require.NoError(t, err)

		err = thread.ExportDag(ctx, job)
		require.NoError(t, err)

		var cars []model.Car
		err = db.Find(&cars).Error
		require.NoError(t, err)
		require.Len(t, cars, 1)

		// Verify the piece size is at least the minPieceSize
		require.GreaterOrEqual(t, cars[0].PieceSize, int64(512*1024))

		// Verify the file size matches the expected padded size (127/128 of piece size)
		expectedFileSize := (cars[0].PieceSize * 127) / 128
		require.Equal(t, expectedFileSize, cars[0].FileSize)

		// Verify MinPieceSizePadding is 0 (padding is physical, not virtual)
		require.Equal(t, int64(0), cars[0].MinPieceSizePadding)
	})
}
