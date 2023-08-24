package api

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/stretchr/testify/require"
)

func TestServerStart(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	server, err := InitServer(ctx, APIParams{
		Listener:   listener,
		LotusAPI:   "https://api.node.glif.io",
		LotusToken: "",
		ConnString: database.TestConnectionString,
	})
	require.NoError(t, err)
	require.Equal(t, "api", server.Name())
	_, _, err = server.Start(ctx)
	require.NoError(t, err)
}
