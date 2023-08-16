package contentprovider

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/stretchr/testify/require"
)

func TestBitswapServer(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	h, err := util.InitHost(nil)
	require.NoError(t, err)
	s := BitswapServer{
		dbNoContext: db,
		host:        h,
	}
	require.Equal(t, "Bitswap", s.Name())
	ctx, cancel := context.WithCancel(context.Background())
	done, _, err := s.Start(ctx)
	require.NoError(t, err)
	time.Sleep(200 * time.Millisecond)
	cancel()
	select {
	case <-time.After(1 * time.Second):
		t.Fatal("bitswap server did not stop")
	case <-done[0]:
	}
}
