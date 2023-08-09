package testutil

import (
	"context"
	"net/http"
	"path/filepath"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/api"
	"github.com/data-preservation-programs/singularity/client"
	httpclient "github.com/data-preservation-programs/singularity/client/http"
	libclient "github.com/data-preservation-programs/singularity/client/lib"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/stretchr/testify/require"
)

func TestWithAllClients(ctx context.Context, t *testing.T, test func(*testing.T, client.Client)) {
	t.Run("http", func(t *testing.T) {
		ctx, cancel := context.WithCancel(ctx)
		httpErr := make(chan error, 1)
		defer func() {
			cancel()
			err := <-httpErr
			require.ErrorIs(t, err, http.ErrServerClosed)
		}()
		server, err := api.InitServer(api.APIParams{
			ConnString: "sqlite:" + filepath.Join(t.TempDir(), "singularity.db"),
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
		db, closer, err := database.OpenWithLogger("sqlite:" + filepath.Join(t.TempDir(), "singularity.db"))
		require.NoError(t, err)
		defer closer.Close()
		client, err := libclient.NewClient(db)
		require.NoError(t, err)
		test(t, client)
	})
}
