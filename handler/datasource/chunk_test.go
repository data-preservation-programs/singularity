package datasource_test

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path"
	"testing"

	"github.com/data-preservation-programs/singularity/client"
	"github.com/data-preservation-programs/singularity/client/testutil"

	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestCreatePackingManifest(t *testing.T) {
	ctx := context.Background()

	// Create test file to add
	testSourcePath := path.Join(os.TempDir(), "singularity-test-source")
	require.NoError(t, os.RemoveAll(testSourcePath))
	require.NoError(t, os.Mkdir(testSourcePath, 0744))
	require.NoError(t, os.WriteFile(path.Join(testSourcePath, "a"), []byte("test file a"), fs.FileMode(os.O_WRONLY)))
	require.NoError(t, os.WriteFile(path.Join(testSourcePath, "b"), []byte("test file b"), fs.FileMode(os.O_WRONLY)))

	testutil.TestWithAllClients(ctx, t, func(t *testing.T, client client.Client) {

		datasetName := "test"

		_, err := client.CreateDataset(ctx, dataset.CreateRequest{
			Name:       datasetName,
			MaxSizeStr: "31GiB",
		})
		require.NoError(t, err)

		// Create source
		source, err := client.CreateLocalSource(ctx, datasetName, datasource.LocalRequest{
			SourcePath:     testSourcePath,
			RescanInterval: "10s",
			ScanningState:  model.Created,
		})
		require.NoError(t, err)

		// Push files
		fileA, err := client.PushFile(ctx, source.ID, datasource.FileInfo{Path: "a"})
		require.NoError(t, err)

		fileB, err := client.PushFile(ctx, source.ID, datasource.FileInfo{Path: "b"})
		require.NoError(t, err)

		// Create packing manifest
		packingManifest, err := client.CreatePackingManifest(ctx, source.ID, datasource.PackingManifestRequest{FileIDs: []uint64{fileA.ID, fileB.ID}})
		require.NoError(t, err)
		fmt.Printf("%#v\n", packingManifest)

		// Check that packing manifest exists
		packingManifests, err := client.GetSourcePackingManifests(ctx, source.ID, inspect.GetSourcePackingManifestsRequest{})
		require.NoError(t, err)

		require.Len(t, packingManifests, 1)
	})
}
