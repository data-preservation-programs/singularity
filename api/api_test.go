package api

import (
	"context"
	"net"
	"os"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestServerStart(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()

		listener, err := net.Listen("tcp", ":0")
		require.NoError(t, err)
		server, err := InitServer(ctx, APIParams{
			Listener:   listener,
			LotusAPI:   "https://api.node.glif.io",
			LotusToken: "",
			ConnString: os.Getenv("DATABASE_CONNECTION_STRING"),
		})
		require.NoError(t, err)
		require.Equal(t, "api", server.Name())
		_, _, err = server.Start(ctx)
		require.NoError(t, err)
	})
}
