package testutil

import (
	"context"
	"net"
	"net/http"
	"testing"

	"github.com/data-preservation-programs/singularity/api"
	"github.com/data-preservation-programs/singularity/client"
	httpclient "github.com/data-preservation-programs/singularity/client/http"
	libclient "github.com/data-preservation-programs/singularity/client/lib"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/stretchr/testify/require"
)

func TestWithAllClients(ctx context.Context, t *testing.T, test func(*testing.T, client.Client)) {
	t.Run("http", func(t *testing.T) {
		//nolint:contextcheck
		_, closer, err := database.OpenInMemory()
		require.NoError(t, err)
		closer.Close()
		ctx, cancel := context.WithCancel(ctx)
		httpErr := make(chan error, 1)
		defer func() {
			cancel()
			err := <-httpErr
			require.ErrorIs(t, err, http.ErrServerClosed)
		}()
		listener, err := net.Listen("tcp", "127.0.0.1:0")
		require.NoError(t, err)
		server, err := api.InitServer(api.APIParams{
			ConnString: database.TestConnectionString,
			Listener:   listener,
			LotusAPI:   "https://api.node.glif.io/rpc/v1",
		})
		require.NoError(t, err)
		go func() {
			err := server.Run(ctx)
			httpErr <- err
		}()
		client := httpclient.NewHTTPClient(http.DefaultClient, "http://"+listener.Addr().String())
		test(t, client)
	})
	t.Run("lib", func(t *testing.T) {
		db, closer, err := database.OpenInMemory()
		require.NoError(t, err)
		defer closer.Close()
		client, err := libclient.NewClient(db)
		require.NoError(t, err)
		test(t, client)
	})
}
