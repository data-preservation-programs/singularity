package util

import (
	"context"
	"reflect"
	"testing"

	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/libp2p/go-libp2p/core/crypto"
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

func TestNewLotusClient(t *testing.T) {
	testutil.SkipIfNotExternalAPI(t)
	client := NewLotusClient(testutil.TestLotusAPI, testutil.TestLotusToken)
	resp, err := client.Call(context.Background(), "Filecoin.Version")
	require.NoError(t, err)
	require.NotNil(t, resp.Result)
}

func TestGetLotusHeadTime(t *testing.T) {
	testutil.SkipIfNotExternalAPI(t)
	headTime, err := GetLotusHeadTime(context.Background(), testutil.TestLotusAPI, testutil.TestLotusToken)
	require.NoError(t, err)
	require.NotZero(t, headTime)
}

func TestPackJobSlice(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		chunkSize int
		expected  [][]int
	}{
		{
			name:      "basic case",
			slice:     []int{1, 2, 3, 4, 5, 6, 7},
			chunkSize: 3,
			expected:  [][]int{{1, 2, 3}, {4, 5, 6}, {7}},
		},
		{
			name:      "chunkSize greater than slice length",
			slice:     []int{1, 2, 3},
			chunkSize: 5,
			expected:  [][]int{{1, 2, 3}},
		},
		{
			name:      "empty slice",
			slice:     []int{},
			chunkSize: 2,
			expected:  [][]int(nil),
		},
		{
			name:      "chunkSize zero",
			slice:     []int{1, 2, 3, 4, 5},
			chunkSize: 0,
			expected:  [][]int(nil),
		},
		{
			name:      "chunkSize equals slice length",
			slice:     []int{1, 2, 3, 4, 5},
			chunkSize: 5,
			expected:  [][]int{{1, 2, 3, 4, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ChunkSlice(tt.slice, tt.chunkSize)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGenerateNewPeer(t *testing.T) {
	privateBytes, publicBytes, peerID, err := GenerateNewPeer()

	require.NoError(t, err, "GenerateNewPeer should not return an error")

	privateKey, err := crypto.UnmarshalPrivateKey(privateBytes)
	require.NoError(t, err, "UnmarshalPrivateKey should not return an error")
	require.NotNil(t, privateKey, "privateKey should not be nil")

	publicKey, err := crypto.UnmarshalPublicKey(publicBytes)
	require.NoError(t, err, "UnmarshalPublicKey should not return an error")
	require.NotNil(t, publicKey, "publicKey should not be nil")

	require.NotEmpty(t, peerID, "peerID should not be empty")

	err = peerID.Validate()
	require.NoError(t, err)
}
