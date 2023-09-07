package inspect_test

import (
	"context"
	"os"
	"path"
	"strconv"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	dshandler "github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestFileDeals(t *testing.T) {
	ctx := context.Background()

	// Create test file to add
	testSourcePath := path.Join(t.TempDir(), "singularity-test-source")
	require.NoError(t, os.Mkdir(testSourcePath, 0744))
	require.NoError(t, os.WriteFile(path.Join(testSourcePath, "a"), []byte("test file a"), 0744))
	require.NoError(t, os.WriteFile(path.Join(testSourcePath, "b"), []byte("test file b"), 0744))

	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = model.AutoMigrate(db)
	require.NoError(t, err)

	datasourceHandlerResolver := &datasource.DefaultHandlerResolver{}

	datasetName := "test"

	_, err = dataset.CreateHandler(ctx, db.WithContext(ctx), dataset.CreateRequest{
		Name:       datasetName,
		MaxSizeStr: "31GiB",
	})
	require.NoError(t, err)

	// Create datasource
	source, err := dshandler.CreateDatasourceHandler(ctx, db.WithContext(ctx), "local", datasetName, map[string]any{
		"sourcePath":        testSourcePath,
		"rescanInterval":    "10s",
		"scanningState":     string(model.Ready),
		"deleteAfterExport": false,
	})
	require.NoError(t, err)

	// Push files
	fileA, err := dshandler.PushFileHandler(ctx, db.WithContext(ctx), datasourceHandlerResolver, source.ID, dshandler.FileInfo{Path: "a"})
	require.NoError(t, err)

	fileB, err := dshandler.PushFileHandler(ctx, db.WithContext(ctx), datasourceHandlerResolver, source.ID, dshandler.FileInfo{Path: "b"})
	require.NoError(t, err)

	dshandler.PrepareToPackSourceHandler(ctx, db.WithContext(ctx), datasourceHandlerResolver, source.ID)
	require.NoError(t, err)

	packJobs, err := inspect.GetSourcePackJobsHandler(ctx, db.WithContext(ctx), source.ID, inspect.GetSourcePackJobsRequest{})
	require.NoError(t, err)
	require.Len(t, packJobs, 1)
	packJob := packJobs[0]

	cars, err := dshandler.PackHandler(db.WithContext(ctx), ctx, datasourceHandlerResolver, packJob.ID)
	require.NoError(t, err)

	// Manually add fake deals
	for _, car := range cars {
		wallet := model.Wallet{
			ID: "Apples",
		}
		err := db.Create(&wallet).Error
		require.NoError(t, err)
		err = db.Create(&model.Deal{
			PieceCID: car.PieceCID,
			ClientID: wallet.ID,
		}).Error
		require.NoError(t, err)
	}

	// Test file deals
	deals, err := inspect.GetFileDealsHandler(db.WithContext(ctx), strconv.FormatUint(fileA.ID, 10))
	require.NoError(t, err)
	require.Len(t, deals, 1)

	deals, err = inspect.GetFileDealsHandler(db.WithContext(ctx), strconv.FormatUint(fileB.ID, 10))
	require.NoError(t, err)
	require.Len(t, deals, 1)
}
