package inspect_test

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"testing"

	libclient "github.com/data-preservation-programs/singularity/client/lib"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/datasetworker"
	"github.com/stretchr/testify/require"
)

func TestItemDeals(t *testing.T) {
	ctx := context.Background()

	// Create test file to add
	testSourcePath := path.Join(t.TempDir(), "singularity-test-source")
	require.NoError(t, os.Mkdir(testSourcePath, 0744))
	require.NoError(t, os.WriteFile(path.Join(testSourcePath, "a"), []byte("test file a"), 0744))
	require.NoError(t, os.WriteFile(path.Join(testSourcePath, "b"), []byte("test file b"), 0744))

	db, closer, err := database.OpenWithLogger("sqlite:" + filepath.Join(t.TempDir(), "singularity.db"))
	require.NoError(t, err)
	defer closer.Close()

	client, err := libclient.NewClient(db)
	require.NoError(t, err)

	datasetName := "test"

	_, err = client.CreateDataset(ctx, dataset.CreateRequest{
		Name:       datasetName,
		MaxSizeStr: "31GiB",
	})
	require.NoError(t, err)

	// Create datasource
	source, err := client.CreateLocalSource(ctx, datasetName, datasource.LocalRequest{
		SourcePath:     testSourcePath,
		RescanInterval: "10s",
		ScanningState:  model.Ready,
	})
	require.NoError(t, err)

	// Push items
	itemA, err := client.PushItem(ctx, source.ID, datasource.ItemInfo{Path: "a"})
	require.NoError(t, err)

	_, err = client.PushItem(ctx, source.ID, datasource.ItemInfo{Path: "b"})
	require.NoError(t, err)

	dsworker := datasetworker.NewDatasetWorker(db, datasetworker.DatasetWorkerConfig{
		Concurrency:    1,
		EnableScan:     true,
		EnablePack:     true,
		EnableDag:      true,
		ExitOnComplete: true,
		ExitOnError:    true,
	})

	require.NoError(t, dsworker.Run(ctx))

	chunks, err := inspect.GetSourceChunksHandler(db, strconv.FormatUint(uint64(source.ID), 10))
	require.NoError(t, err)

	// Manually add fake deals
	for _, car := range chunks[0].Cars {
		err := db.Create(&model.Deal{
			PieceCID: car.PieceCID,
		}).Error
		require.NoError(t, err)
	}

	// Test item deals
	deals, err := client.GetItemDeals(ctx, itemA.ID)
	require.NoError(t, err)
	require.Len(t, deals, 1)

	fmt.Printf("%#v\n", deals)
}
