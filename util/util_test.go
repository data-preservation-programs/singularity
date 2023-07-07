package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNextPowerOfTwo(t *testing.T) {
	require.Equal(t, uint64(1), NextPowerOfTwo(0))
	require.Equal(t, uint64(1), NextPowerOfTwo(1))
	require.Equal(t, uint64(2), NextPowerOfTwo(2))
	require.Equal(t, uint64(4), NextPowerOfTwo(3))
	require.Equal(t, uint64(4), NextPowerOfTwo(4))
	require.Equal(t, uint64(8), NextPowerOfTwo(5))
	require.Equal(t, uint64(16), NextPowerOfTwo(9))
	require.Equal(t, uint64(32), NextPowerOfTwo(17))
	require.Equal(t, uint64(64), NextPowerOfTwo(33))
}
