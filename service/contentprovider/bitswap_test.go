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
		s := BitswapServer{
			dbNoContext: db,
			host:        h,
		}
		require.Equal(t, "Bitswap", s.Name())
		ctx, cancel := context.WithCancel(ctx)
		done, _, err := s.Start(ctx)
		require.NoError(t, err)
		time.Sleep(200 * time.Millisecond)
		cancel()
		select {
		case <-time.After(1 * time.Second):
			t.Fatal("bitswap server did not stop")
		case <-done[0]:
		}
	})
}
