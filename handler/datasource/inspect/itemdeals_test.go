package inspect_test

import (
	"context"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	dshandler "github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/rjNemo/underscore"
	"github.com/stretchr/testify/require"
)

func TestItemDeals(t *testing.T) {
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

	_, err = dataset.CreateHandler(db.WithContext(ctx), dataset.CreateRequest{
		Name:       datasetName,
		MaxSizeStr: "31GiB",
	})
	require.NoError(t, err)

	// Create datasource
	source, err := dshandler.CreateDatasourceHandler(db.WithContext(ctx), ctx, datasourceHandlerResolver, "local", datasetName, map[string]any{
		"sourcePath":        testSourcePath,
		"rescanInterval":    "10s",
		"scanningState":     string(model.Ready),
		"deleteAfterExport": false,
	})
	require.NoError(t, err)

	// Push items
	itemA, err := dshandler.PushItemHandler(db.WithContext(ctx), ctx, datasourceHandlerResolver, source.ID, dshandler.ItemInfo{Path: "a"})
	require.NoError(t, err)

	itemB, err := dshandler.PushItemHandler(db.WithContext(ctx), ctx, datasourceHandlerResolver, source.ID, dshandler.ItemInfo{Path: "b"})
	require.NoError(t, err)

	chunk, err := dshandler.ChunkHandler(db.WithContext(ctx), source.ID, dshandler.ChunkRequest{
		ItemIDs: append(underscore.Map(itemA.ItemParts, func(itemPart model.ItemPart) uint64 { return itemPart.ID }),
			underscore.Map(itemB.ItemParts, func(itemPart model.ItemPart) uint64 { return itemPart.ID })...),
	})
	require.NoError(t, err)
	cars, err := dshandler.PackHandler(db.WithContext(ctx), ctx, datasourceHandlerResolver, chunk.ID)
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

	// Test item deals
	deals, err := inspect.GetItemDealsHandler(db.WithContext(ctx), itemA.ID)
	require.NoError(t, err)
	require.Len(t, deals, 1)
	fmt.Printf("%#v\n", deals)

	deals, err = inspect.GetItemDealsHandler(db.WithContext(ctx), itemB.ID)
	require.NoError(t, err)
	require.Len(t, deals, 1)
}
