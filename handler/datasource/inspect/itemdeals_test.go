package inspect_test

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"testing"

	libclient "github.com/data-preservation-programs/singularity/client/lib"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource"
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

	itemB, err := client.PushItem(ctx, source.ID, datasource.ItemInfo{Path: "b"})
	require.NoError(t, err)

	chunk, err := client.Chunk(ctx, source.ID, datasource.ChunkRequest{
		ItemIDs: append(underscore.Map(itemA.ItemParts, func(itemPart model.ItemPart) uint64 { return itemPart.ID }),
			underscore.Map(itemB.ItemParts, func(itemPart model.ItemPart) uint64 { return itemPart.ID })...),
	})
	require.NoError(t, err)
	cars, err := client.Pack(ctx, chunk.ID)
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
	deals, err := client.GetItemDeals(ctx, itemA.ID)
	require.NoError(t, err)
	require.Len(t, deals, 1)
	fmt.Printf("%#v\n", deals)

	deals, err = client.GetItemDeals(ctx, itemB.ID)
	require.NoError(t, err)
	require.Len(t, deals, 1)
}
