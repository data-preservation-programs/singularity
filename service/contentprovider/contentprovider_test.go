package contentprovider

import (
	"context"
	"encoding/base64"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestContentProviderStart(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		service, err := NewService(db, Config{
			HTTP: HTTPConfig{
				EnablePiece:         true,
				EnablePieceMetadata: true,
				Bind:                ":0",
			},
			Bitswap: BitswapConfig{
				Enable:           true,
				IdentityKey:      "",
				ListenMultiAddrs: nil,
			},
		})
		require.NoError(t, err)
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		err = service.Start(ctx)
		require.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

func TestContentProviderStart_WithIdentityKey(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		private, _, _, err := util.GenerateNewPeer()
		require.NoError(t, err)
		encoded := base64.StdEncoding.EncodeToString(private)
		service, err := NewService(db, Config{
			HTTP: HTTPConfig{
				EnablePiece:         true,
				EnablePieceMetadata: true,
				Bind:                ":0",
			},
			Bitswap: BitswapConfig{
				Enable:           true,
				IdentityKey:      encoded,
				ListenMultiAddrs: nil,
			},
		})
		require.NoError(t, err)
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		err = service.Start(ctx)
		require.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

func TestContentProviderStart_NoneEnabled(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		s, err := NewService(db, Config{
			HTTP: HTTPConfig{
				EnablePiece:         true,
				EnablePieceMetadata: true,
			},
			Bitswap: BitswapConfig{
				Enable: false,
			},
		})
		require.NoError(t, err)
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		err = s.Start(ctx)
		require.ErrorIs(t, err, service.ErrNoService)
	})
}
