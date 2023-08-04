package datasource_test

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
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/stretchr/testify/require"
)

func TestChunk(t *testing.T) {
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

	// Create source
	source, err := client.CreateLocalSource(ctx, datasetName, datasource.LocalRequest{
		SourcePath:     testSourcePath,
		RescanInterval: "10s",
	})
	require.NoError(t, err)

	// Push items
	itemA, err := client.PushItem(ctx, source.ID, datasource.ItemInfo{Path: "a"})
	require.NoError(t, err)

	itemB, err := client.PushItem(ctx, source.ID, datasource.ItemInfo{Path: "b"})
	require.NoError(t, err)

	// Chunk
	chunk, err := client.Chunk(ctx, source.ID, datasource.ChunkRequest{ItemIDs: []uint64{itemA.ID, itemB.ID}})
	require.NoError(t, err)
	fmt.Printf("%#v\n", chunk)

	// Check that chunk exists
	chunks, err := client.GetSourceChunks(ctx, source.ID)
	require.NoError(t, err)

	require.Len(t, chunks, 1)
}
