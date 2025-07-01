package migrate

import (
	"context"
	"flag"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

// Using 27018 intentionally to avoid deleting default singularity V1 database
var localMongoDB = "mongodb://localhost:27018"

func TestMigrateDataset(t *testing.T) {
	err := setupMongoDBDataset()
	if err != nil {
		t.Log(err)
		t.Skip("Skipping test because MongoDB is not available")
	}
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		flagSet := flag.NewFlagSet("", 0)
		flagSet.String("mongo-connection-string", localMongoDB, "")
		flagSet.String("database-connection-string", os.Getenv("DATABASE_CONNECTION_STRING"), "")
		cctx := cli.NewContext(&cli.App{
			Writer: os.Stdout,
		}, flagSet, nil)
		err = MigrateDataset(cctx)
		require.NoError(t, err)
		// Migrate again does nothing
		err = MigrateDataset(cctx)
		require.NoError(t, err)

		var preparations []model.Preparation
		err = db.Preload("SourceStorages").Preload("OutputStorages").Find(&preparations).Error
		require.NoError(t, err)
		require.Len(t, preparations, 2)
		require.Equal(t, "test-source", preparations[0].SourceStorages[0].Name)
		require.Equal(t, "/path", preparations[0].SourceStorages[0].Path)
		require.Equal(t, "local", preparations[0].SourceStorages[0].Type)
		require.EqualValues(t, int64(18*1024*1024*1024), preparations[0].MaxSize)
		require.EqualValues(t, int64(32*1024*1024*1024), preparations[0].PieceSize)
		require.Equal(t, "test2-source", preparations[1].SourceStorages[0].Name)
		require.Equal(t, "s3path", preparations[1].SourceStorages[0].Path)
		require.Equal(t, "s3", preparations[1].SourceStorages[0].Type)
		require.Equal(t, filepath.Join("out", "dir"), preparations[0].OutputStorages[0].Path)

		var dirs []model.Directory
		err = db.Find(&dirs).Error
		require.NoError(t, err)
		require.Len(t, dirs, 3)
		require.Equal(t, "/path", dirs[0].Name)
		require.Equal(t, "dir", dirs[1].Name)
		require.Equal(t, "s3path", dirs[2].Name)

		var files []model.File
		err = db.Find(&files).Error
		require.NoError(t, err)
		require.Len(t, files, 3)
		require.Equal(t, "1.txt", files[0].Path)
		require.Equal(t, "2.txt", files[1].Path)
		require.Equal(t, "dir/3.txt", files[2].Path)
		require.EqualValues(t, 1, *files[0].DirectoryID)
		require.EqualValues(t, 1, *files[1].DirectoryID)
		require.EqualValues(t, 2, *files[2].DirectoryID)

		var fileRanges []model.FileRange
		err = db.Find(&fileRanges).Error
		require.NoError(t, err)
		require.Len(t, fileRanges, 5)
		require.EqualValues(t, 0, fileRanges[0].Offset)
		require.EqualValues(t, 100, fileRanges[0].Length)
		require.EqualValues(t, 0, fileRanges[1].Offset)
		require.EqualValues(t, 20, fileRanges[1].Length)
		require.EqualValues(t, 20, fileRanges[2].Offset)
		require.EqualValues(t, 60, fileRanges[2].Length)
		require.EqualValues(t, 80, fileRanges[3].Offset)
		require.EqualValues(t, 20, fileRanges[3].Length)

		var packJobs []model.Job
		err = db.Find(&packJobs).Error
		require.NoError(t, err)
		require.Len(t, packJobs, 2)
		require.EqualValues(t, 1, packJobs[0].AttachmentID)
		require.Equal(t, model.Complete, packJobs[0].State)
		require.Equal(t, "error message", packJobs[0].ErrorMessage)

		var cars []model.Car
		err = db.Find(&cars).Error
		require.NoError(t, err)
		require.Len(t, cars, 2)
		require.EqualValues(t, int64(32*1024*1024*1024), cars[0].PieceSize)
		require.EqualValues(t, int64(20*1024*1024*1024), cars[0].FileSize)
		require.EqualValues(t, filepath.Join("out", "dir", "test.car"), cars[0].StoragePath)
		require.NotEmpty(t, cars[0].PieceCID.String())
		require.NotEmpty(t, cars[0].RootCID.String())
		require.EqualValues(t, 1, *cars[0].AttachmentID)
	})
}

func setupMongoDBDataset() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(localMongoDB))
	if err != nil {
		return errors.WithStack(err)
	}
	defer db.Disconnect(context.Background())
	err = db.Database("singularity").Drop(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	insertScanningResult, err := db.Database("singularity").Collection("scanningrequests").InsertMany(ctx, []any{ScanningRequest{
		Name:                  "test",
		Path:                  "/path",
		OutDir:                filepath.Join("out", "dir"),
		MinSize:               uint64(16 * 1024 * 1024 * 1024),
		MaxSize:               uint64(18 * 1024 * 1024 * 1024),
		Status:                ScanningStatusCompleted,
		ErrorMessage:          "error message",
		TmpDir:                "/tmp/dir",
		SkipInaccessibleFiles: false,
	}, ScanningRequest{
		Name:                  "test2",
		Path:                  "s3://s3path",
		OutDir:                filepath.Join("out", "dir"),
		MinSize:               uint64(16 * 1024 * 1024 * 1024),
		MaxSize:               uint64(18 * 1024 * 1024 * 1024),
		Status:                ScanningStatusCompleted,
		ErrorMessage:          "error message",
		TmpDir:                "/tmp/dir",
		SkipInaccessibleFiles: false,
	}})
	if err != nil {
		return errors.WithStack(err)
	}

	dataCID := cid.NewCidV1(cid.Raw, util.Hash([]byte("test")))
	pieceCID := cid.NewCidV1(cid.FilCommitmentUnsealed, util.Hash([]byte("test")))
	insertGenerationResult, err := db.Database("singularity").Collection("generationrequests").InsertMany(ctx, []any{GenerationRequest{
		DatasetID:             insertScanningResult.InsertedIDs[0].(primitive.ObjectID).Hex(),
		DatasetName:           "test",
		Path:                  "/path",
		OutDir:                filepath.Join("out", "dir"),
		Index:                 0,
		Status:                GenerationStatusCompleted,
		ErrorMessage:          "error message",
		DataCID:               dataCID.String(),
		CarSize:               uint64(20 * 1024 * 1024 * 1024),
		PieceCID:              pieceCID.String(),
		PieceSize:             uint64(32 * 1024 * 1024 * 1024),
		FilenameOverride:      "test.car",
		TmpDir:                "/tmp/dir",
		SkipInaccessibleFiles: false,
		CreatedAt:             time.Now(),
	}, GenerationRequest{
		DatasetID:             insertScanningResult.InsertedIDs[0].(primitive.ObjectID).Hex(),
		DatasetName:           "test",
		Path:                  "/path",
		OutDir:                filepath.Join("out", "dir"),
		Index:                 1,
		Status:                GenerationStatusCompleted,
		ErrorMessage:          "error message",
		DataCID:               "unrecoverable",
		CarSize:               uint64(20 * 1024 * 1024 * 1024),
		PieceCID:              pieceCID.String(),
		PieceSize:             uint64(32 * 1024 * 1024 * 1024),
		FilenameOverride:      "test2.car",
		TmpDir:                "/tmp/dir",
		SkipInaccessibleFiles: false,
		CreatedAt:             time.Now(),
	}})
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = db.Database("singularity").Collection("outputfilelists").InsertMany(ctx, []any{OutputFileList{
		GenerationID: insertGenerationResult.InsertedIDs[0].(primitive.ObjectID).Hex(),
		Index:        0,
		GeneratedFileList: []GeneratedFile{{
			Path: "",
			Dir:  true,
		}, {
			Path:  "1.txt",
			Dir:   false,
			CID:   dataCID.String(),
			Size:  100,
			Start: 0,
			End:   100,
		}, {
			Path:  "2.txt",
			Dir:   false,
			CID:   dataCID.String(),
			Size:  100,
			Start: 0,
			End:   20,
		}, {
			Path:  "2.txt",
			Dir:   false,
			CID:   dataCID.String(),
			Size:  100,
			Start: 20,
			End:   80,
		}, {
			Path:  "2.txt",
			Dir:   false,
			CID:   dataCID.String(),
			Size:  100,
			Start: 80,
			End:   100,
		}, {
			Path: "dir",
			Dir:  true,
		}, {
			Path:  "dir/3.txt",
			Dir:   false,
			CID:   dataCID.String(),
			Size:  100,
			Start: 0,
			End:   0,
		}, {
			Path:  "dir/4.txt",
			Dir:   false,
			CID:   "unrecoverable",
			Size:  100,
			Start: 0,
			End:   0,
		}},
	}})
	return errors.WithStack(err)
}
