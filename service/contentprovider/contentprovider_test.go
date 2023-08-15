package contentprovider

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/stretchr/testify/require"
)

func TestContentProviderStart(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	service, err := NewService(db, Config{
		HTTP: HTTPConfig{
			Enable: true,
			Bind:   ":0",
		},
		Bitswap: BitswapConfig{
			Enable:           true,
			IdentityKey:      "",
			ListenMultiAddrs: nil,
		},
	})
	require.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = service.Start(ctx)
	require.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestContentProviderStart_NoneEnabled(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	s, err := NewService(db, Config{
		HTTP: HTTPConfig{
			Enable: false,
		},
		Bitswap: BitswapConfig{
			Enable: false,
		},
	})
	require.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = s.Start(ctx)
	require.ErrorIs(t, err, service.ErrNoService)
}
