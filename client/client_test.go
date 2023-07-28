package client_test

import (
	"context"
	"crypto/rand"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/api"
	"github.com/data-preservation-programs/singularity/client"
	httpclient "github.com/data-preservation-programs/singularity/client/http"
	libclient "github.com/data-preservation-programs/singularity/client/lib"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/stretchr/testify/require"
)

func TestClients(t *testing.T) {
	ctx := context.Background()
	testWithAllClients(ctx, t, func(t *testing.T, client client.Client) {
		// createDataset
		ds, err := client.CreateDataset(ctx, dataset.CreateRequest{
			Name:       "test",
			MaxSizeStr: "31.5GiB",
		})
		require.NoError(t, err)
		require.Equal(t, "test", ds.Name)

		// cannot create dataset with same name

		dupDataset, err := client.CreateDataset(ctx, dataset.CreateRequest{
			Name:       "test",
			MaxSizeStr: "31.5GiB",
		})
		var asConflict handler.DuplicateRecordError
		require.ErrorAs(t, err, &asConflict)
		require.Nil(t, dupDataset)

		// cannot create dataset with invalid parameter
		invalidDataset, err := client.CreateDataset(ctx, dataset.CreateRequest{})
		var asInvalidParameter handler.InvalidParameterError
		require.ErrorAs(t, err, &asInvalidParameter)
		require.Nil(t, invalidDataset)

		path := t.TempDir()
		// create datasource
		source, err := client.CreateLocalSource(ctx, "test", datasource.LocalRequest{
			SourcePath:     path,
			RescanInterval: "0h",
			ScanningState:  model.Ready,
		})
		require.NoError(t, err)
		require.Equal(t, "local", source.Type)
		require.Equal(t, ds.ID, source.DatasetID)
		require.Equal(t, path, source.Path)
		require.Equal(t, model.Ready, source.ScanningState)

		// list sources for dataset
		sources, err := client.ListSourcesByDataset(ctx, "test")
		require.NoError(t, err)
		require.Len(t, sources, 1)
		require.Equal(t, ds.ID, sources[0].DatasetID)
		require.Equal(t, path, sources[0].Path)
		require.Equal(t, model.Ready, sources[0].ScanningState)

		// create datasource when dataset not found
		notFoundSource, err := client.CreateLocalSource(ctx, "apples", datasource.LocalRequest{
			SourcePath:     path,
			RescanInterval: "0h",
		})
		var asNotFoundError handler.NotFoundError
		require.ErrorAs(t, err, &asNotFoundError)
		require.Nil(t, notFoundSource)

		// push item
		file, err := os.CreateTemp(path, "push-*")
		require.NoError(t, err)
		buf := make([]byte, 1000)
		_, _ = rand.Read(buf)
		file.Write(buf)
		name := file.Name()
		err = file.Close()
		require.NoError(t, err)
		item, err := client.PushItem(ctx, source.ID, datasource.ItemInfo{Path: filepath.Base(name)})
		require.NoError(t, err)
		require.Equal(t, filepath.Base(name), item.Path)

		// get item
		returnedItem, err := client.GetItem(ctx, item.ID)
		require.NoError(t, err)
		require.Equal(t, item.Path, returnedItem.Path)
	})
}

func testWithAllClients(ctx context.Context, t *testing.T, test func(*testing.T, client.Client)) {
	t.Run("http", func(t *testing.T) {
		ctx, cancel := context.WithCancel(ctx)
		httpErr := make(chan error, 1)
		defer func() {
			cancel()
			err := <-httpErr
			require.ErrorIs(t, err, http.ErrServerClosed)
		}()
		server, err := api.InitServer(api.APIParams{
			ConnString: "sqlite:" + t.TempDir() + "/singularity.db",
			Bind:       "127.0.0.1:9090",
			LotusAPI:   "https://api.node.glif.io/rpc/v1",
		})
		require.NoError(t, err)
		go func() {
			err := server.Run(ctx)
			httpErr <- err
		}()
		time.Sleep(time.Second)
		client := httpclient.NewHTTPClient(http.DefaultClient, "http://127.0.0.1:9090")
		test(t, client)
	})
	t.Run("lib", func(t *testing.T) {
		db, err := database.OpenWithDefaults("sqlite:" + t.TempDir() + "/singularity.db")
		require.NoError(t, err)
		client, err := libclient.NewClient(db)
		require.NoError(t, err)
		test(t, client)
	})
}
