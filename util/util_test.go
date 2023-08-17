package util

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/rjNemo/underscore"
	"github.com/stretchr/testify/require"
)

func TestGetLotusHeadTime(t *testing.T) {
	head, err := GetLotusHeadTime(context.Background(), "https://api.node.glif.io/", "")
	require.NoError(t, err)
	require.Greater(t, time.Since(head).Seconds(), float64(0))
	require.Less(t, time.Since(head).Seconds(), float64(3600))
}

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
	for _, token := range []string{"token", ""} {
		t.Run(token, func(t *testing.T) {
			client := NewLotusClient("https://api.node.glif.io/", token)
			resp, err := client.Call(context.Background(), "Filecoin.Version")
			require.NoError(t, err)
			require.NotNil(t, resp.Result)
		})
	}
}

func TestPackJobSlice(t *testing.T) {
	tests := []struct {
		name        string
		slice       []int
		packJobSize int
		expected    [][]int
	}{
		{
			name:        "basic case",
			slice:       []int{1, 2, 3, 4, 5, 6, 7},
			packJobSize: 3,
			expected:    [][]int{{1, 2, 3}, {4, 5, 6}, {7}},
		},
		{
			name:        "packJobSize greater than slice length",
			slice:       []int{1, 2, 3},
			packJobSize: 5,
			expected:    [][]int{{1, 2, 3}},
		},
		{
			name:        "empty slice",
			slice:       []int{},
			packJobSize: 2,
			expected:    [][]int(nil),
		},
		{
			name:        "packJobSize zero",
			slice:       []int{1, 2, 3, 4, 5},
			packJobSize: 0,
			expected:    [][]int(nil),
		},
		{
			name:        "packJobSize equals slice length",
			slice:       []int{1, 2, 3, 4, 5},
			packJobSize: 5,
			expected:    [][]int{{1, 2, 3, 4, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PackJobSlice(tt.slice, tt.packJobSize)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPackJobMapKeys(t *testing.T) {
	tests := []struct {
		name        string
		m           map[string]int
		packJobSize int
	}{
		{
			name:        "basic case",
			m:           map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7},
			packJobSize: 3,
		},
		{
			name:        "packJobSize greater than map size",
			m:           map[string]int{"a": 1, "b": 2, "c": 3},
			packJobSize: 5,
		},
		{
			name:        "empty map",
			m:           map[string]int{},
			packJobSize: 2,
		},
		{
			name:        "packJobSize zero",
			m:           map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			packJobSize: 0,
		},
		{
			name:        "packJobSize equals map size",
			m:           map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			packJobSize: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PackJobMapKeys(tt.m, tt.packJobSize)
			if tt.packJobSize <= 0 {
				require.Equal(t, [][]string(nil), result)
				return
			}
			total := underscore.SumMap(result, func(keys []string) int {
				return len(keys)
			})
			require.Equal(t, len(tt.m), total)
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
