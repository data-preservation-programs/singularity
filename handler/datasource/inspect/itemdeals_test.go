package inspect_test

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"testing"

	libclient "github.com/data-preservation-programs/singularity/client/lib"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	datasourcehandler "github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/stretchr/testify/require"
)

func TestItemDeals(t *testing.T) {
	ctx := context.Background()

	// Create test file to add
	testSourcePath := path.Join(os.TempDir(), "singularity-test-source")
	require.NoError(t, os.RemoveAll(testSourcePath))
	require.NoError(t, os.Mkdir(testSourcePath, 0744))
	require.NoError(t, os.WriteFile(path.Join(testSourcePath, "a"), []byte("test file a"), fs.FileMode(os.O_WRONLY)))
	require.NoError(t, os.WriteFile(path.Join(testSourcePath, "b"), []byte("test file b"), fs.FileMode(os.O_WRONLY)))

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
	datasource, err := client.CreateLocalSource(ctx, datasetName, datasourcehandler.LocalRequest{
		SourcePath:     testSourcePath,
		RescanInterval: "10s",
	})
	require.NoError(t, err)

	// Push items
	_, err = client.PushItem(ctx, datasource.ID, datasourcehandler.ItemInfo{Path: "a"})
	require.NoError(t, err)

	_, err = client.PushItem(ctx, datasource.ID, datasourcehandler.ItemInfo{Path: "b"})
	require.NoError(t, err)

	items, err := client.GetSourceItems(ctx, datasource.ID)
	require.NoError(t, err)

	fmt.Printf("item count: %v\n", len(items))

	for _, item := range items {
		deals, err := client.GetItemDeals(ctx, item.ID)
		require.NoError(t, err)

		fmt.Printf("%#v", deals)
	}
}
