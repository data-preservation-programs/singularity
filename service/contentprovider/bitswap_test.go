package contentprovider

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestBitswapServer(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		h, err := util.InitHost(nil)
		require.NoError(t, err)
		defer func() { _ = h.Close() }()
		s := BitswapServer{
			dbNoContext: db,
			host:        h,
		}
		require.Equal(t, "Bitswap", s.Name())

		exitErr := make(chan error, 1)
		ctx, cancel := context.WithCancel(ctx)
		err = s.Start(ctx, exitErr)
		require.NoError(t, err)
		time.Sleep(200 * time.Millisecond)
		cancel()
		select {
		case <-time.After(1 * time.Second):
			t.Fatal("bitswap server did not stop")
		case err = <-exitErr:
			require.NoError(t, err)
		}
	})
}
