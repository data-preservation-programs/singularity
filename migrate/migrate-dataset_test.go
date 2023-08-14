package migrate

import (
	"context"
	"flag"
	"path/filepath"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-cid"
	util "github.com/ipfs/go-ipfs-util"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Using 27018 intentionally to avoid deleting default singularity V1 database
var localMongoDB = "mongodb://localhost:27018"

func TestMigrateDataset(t *testing.T) {
	err := setupMongoDBDataset()
	if err != nil {
		t.Log(err)
		t.Skip("Skipping test because MongoDB is not available")
	}

	// Make sure we have connection to sqlite inmemory to prevent it being garbage collected
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	flagSet := flag.NewFlagSet("", 0)
	flagSet.String("mongo-connection-string", localMongoDB, "")
	flagSet.String("database-connection-string", database.TestConnectionString, "")
	cctx := cli.NewContext(nil, flagSet, nil)
	err = MigrateDataset(cctx)
	require.NoError(t, err)
	// Migrate again does nothing
	err = MigrateDataset(cctx)
	require.NoError(t, err)

	var datasets []model.Dataset
	err = db.Find(&datasets).Error
	require.NoError(t, err)
	require.Len(t, datasets, 2)
	require.Equal(t, "test", datasets[0].Name)
	require.EqualValues(t, 18*1024*1024*1024, datasets[0].MaxSize)
	require.EqualValues(t, 32*1024*1024*1024, datasets[0].PieceSize)
	require.Equal(t, "test2", datasets[1].Name)
	require.Equal(t, []string{filepath.Join("out", "dir")}, []string(datasets[0].OutputDirs))

	var sources []model.Source
	err = db.Find(&sources).Error
	require.NoError(t, err)
	require.Len(t, sources, 2)
	require.Equal(t, "s3path", sources[1].Path)
	require.Equal(t, "/path", sources[0].Path)
	require.Equal(t, "s3", sources[1].Type)
	require.Equal(t, "local", sources[0].Type)
	require.EqualValues(t, 1, sources[0].DatasetID)
	require.Equal(t, model.Complete, sources[0].ScanningState)

	var dirs []model.Directory
	err = db.Find(&dirs).Error
	require.NoError(t, err)
	require.Len(t, dirs, 3)
	require.Equal(t, "/path", dirs[0].Name)
	require.Equal(t, "dir", dirs[1].Name)
	require.Equal(t, "s3path", dirs[2].Name)
	require.EqualValues(t, 1, dirs[0].SourceID)
	require.EqualValues(t, 1, dirs[1].SourceID)
	require.EqualValues(t, 2, dirs[2].SourceID)
	require.EqualValues(t, 1, *dirs[1].ParentID)

	var items []model.Item
	err = db.Find(&items).Error
	require.NoError(t, err)
	require.Len(t, items, 3)
	require.Equal(t, "1.txt", items[0].Path)
	require.Equal(t, "2.txt", items[1].Path)
	require.Equal(t, "dir/3.txt", items[2].Path)
	require.EqualValues(t, 1, items[0].SourceID)
	require.EqualValues(t, 1, items[1].SourceID)
	require.EqualValues(t, 1, items[2].SourceID)
	require.EqualValues(t, 1, *items[0].DirectoryID)
	require.EqualValues(t, 1, *items[1].DirectoryID)
	require.EqualValues(t, 2, *items[2].DirectoryID)

	var itemParts []model.ItemPart
	err = db.Find(&itemParts).Error
	require.NoError(t, err)
	require.Len(t, itemParts, 5)
	require.EqualValues(t, 0, itemParts[0].Offset)
	require.EqualValues(t, 100, itemParts[0].Length)
	require.EqualValues(t, 0, itemParts[1].Offset)
	require.EqualValues(t, 20, itemParts[1].Length)
	require.EqualValues(t, 20, itemParts[2].Offset)
	require.EqualValues(t, 60, itemParts[2].Length)
	require.EqualValues(t, 80, itemParts[3].Offset)
	require.EqualValues(t, 20, itemParts[3].Length)

	var chunks []model.Chunk
	err = db.Find(&chunks).Error
	require.NoError(t, err)
	require.Len(t, chunks, 2)
	require.EqualValues(t, 1, chunks[0].SourceID)
	require.Equal(t, model.Complete, chunks[0].PackingState)
	require.Equal(t, "error message", chunks[0].ErrorMessage)

	var cars []model.Car
	err = db.Find(&cars).Error
	require.NoError(t, err)
	require.Len(t, cars, 2)
	require.EqualValues(t, 32*1024*1024*1024, cars[0].PieceSize)
	require.EqualValues(t, 20*1024*1024*1024, cars[0].FileSize)
	require.EqualValues(t, filepath.Join("out", "dir", "test.car"), cars[0].FilePath)
	require.NotEmpty(t, cars[0].PieceCID.String())
	require.NotEmpty(t, cars[0].RootCID.String())
	require.EqualValues(t, 1, *cars[0].ChunkID)
	require.EqualValues(t, 1, *cars[0].SourceID)
}

func setupMongoDBDataset() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(localMongoDB))
	if err != nil {
		return err
	}
	defer db.Disconnect(context.Background())
	err = db.Database("singularity").Drop(ctx)
	if err != nil {
		return err
	}
	insertScanningResult, err := db.Database("singularity").Collection("scanningrequests").InsertMany(ctx, []any{ScanningRequest{
		Name:                  "test",
		Path:                  "/path",
		OutDir:                filepath.Join("out", "dir"),
		MinSize:               16 * 1024 * 1024 * 1024,
		MaxSize:               18 * 1024 * 1024 * 1024,
		Status:                ScanningStatusCompleted,
		ErrorMessage:          "error message",
		TmpDir:                "/tmp/dir",
		SkipInaccessibleFiles: false,
	}, ScanningRequest{
		Name:                  "test2",
		Path:                  "s3://s3path",
		OutDir:                filepath.Join("out", "dir"),
		MinSize:               16 * 1024 * 1024 * 1024,
		MaxSize:               18 * 1024 * 1024 * 1024,
		Status:                ScanningStatusCompleted,
		ErrorMessage:          "error message",
		TmpDir:                "/tmp/dir",
		SkipInaccessibleFiles: false,
	}})
	if err != nil {
		return err
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
		CarSize:               20 * 1024 * 1024 * 1024,
		PieceCID:              pieceCID.String(),
		PieceSize:             32 * 1024 * 1024 * 1024,
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
		CarSize:               20 * 1024 * 1024 * 1024,
		PieceCID:              pieceCID.String(),
		PieceSize:             32 * 1024 * 1024 * 1024,
		FilenameOverride:      "test2.car",
		TmpDir:                "/tmp/dir",
		SkipInaccessibleFiles: false,
		CreatedAt:             time.Now(),
	}})
	if err != nil {
		return err
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
	return err
}
