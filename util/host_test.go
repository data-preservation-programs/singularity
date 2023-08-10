package util

import (
	"testing"

	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"
)

func TestInitHost(t *testing.T) {
	listen, err := multiaddr.NewMultiaddr("/ip4/0.0.0.0/tcp/0")
	require.NoError(t, err)
	h, err := InitHost(nil, listen)
	require.NoError(t, err)
	require.NotNil(t, h)
	require.NoError(t, h.Close())
}
