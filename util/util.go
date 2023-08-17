package util

import (
	"context"
	"crypto/rand"
	"time"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pkg/errors"
	"github.com/ybbus/jsonrpc/v3"
)

// NextPowerOfTwo calculates the smallest power of two that is greater than or equal to x.
// If x is already a power of two, it returns x. For x equal to 0, the result is 1.
//
// Parameters:
// - x: The input value for which the next power of two needs to be calculated.
//
// Returns:
// The smallest power of two that is greater than or equal to x.
func NextPowerOfTwo(x uint64) uint64 {
	if x == 0 {
		return 1
	}

	// Find the position of the highest bit set to 1
	pos := uint(0)
	for shifted := x; shifted > 0; shifted >>= 1 {
		pos++
	}

	// If x is already a power of two, return x
	if x == 1<<(pos-1) {
		return x
	}

	// Otherwise, return the next power of two
	return 1 << pos
}

// NewLotusClient is a function that creates a new JSON-RPC client for interacting with a Lotus node.
// It takes the Lotus API endpoint and an optional Lotus token as input.
// If the Lotus token is provided, it is included in the 'Authorization' header of the JSON-RPC requests.
//
// Parameters:
// lotusAPI: The Lotus API endpoint. Must be a valid URL.
// lotusToken: An optional Lotus token. If provided, it is included in the 'Authorization' header of the JSON-RPC requests.
//
// Returns:
// A JSON-RPC client configured to interact with the specified Lotus node. If a Lotus token is provided, the client includes it in the 'Authorization' header of its requests.
func NewLotusClient(lotusAPI string, lotusToken string) jsonrpc.RPCClient {
	if lotusToken == "" {
		return jsonrpc.NewClient(lotusAPI)
	} else {
		return jsonrpc.NewClientWithOpts(lotusAPI, &jsonrpc.RPCClientOpts{
			CustomHeaders: map[string]string{
				"Authorization": "Bearer " + lotusToken,
			},
		})
	}
}

// GetLotusHeadTime retrieves the timestamp of the latest block in the Lotus API and returns it as a time.Time value.
//
//nolint:tagliatelle
func GetLotusHeadTime(ctx context.Context, lotusAPI string, lotusToken string) (time.Time, error) {
	client := NewLotusClient(lotusAPI, lotusToken)
	var resp struct {
		Blocks []struct {
			Timestamp int64 `json:"Timestamp"`
		} `json:"Blocks"`
	}
	err := client.CallFor(ctx, &resp, "Filecoin.ChainHead")
	if err != nil {
		return time.Time{}, errors.Wrap(err, "failed to get chain head")
	}

	if len(resp.Blocks) == 0 {
		return time.Time{}, errors.New("chain head is empty")
	}

	return time.Unix(resp.Blocks[0].Timestamp, 0), nil
}

// ChunkMapKeys is a generic function that takes a map with keys of any comparable type and values of any type, and an integer as input.
// It divides the keys of the input map into packJobs of size 'packJobSize' and returns a 2D slice of keys.
// It uses the ChunkSlice function to divide the keys into packJobs.
//
// Parameters:
// m: A map with keys of any comparable type and values of any type. The keys of this map will be packJobed.
// packJobSize: The size of each packJob. Must be a positive integer.
//
// Returns:
// A 2D slice where each inner slice is of length 'packJobSize'. The last inner slice may be shorter if the number of keys in 'm' is not a multiple of 'packJobSize'.
func ChunkMapKeys[T1 comparable, T2 any](m map[T1]T2, packJobSize int) [][]T1 {
	keys := make([]T1, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return ChunkSlice(keys, packJobSize)
}

// ChunkSlice is a generic function that takes a slice of any type and an integer as input.
// It divides the input slice into packJobs of size 'packJobSize' and returns a 2D slice.
// If 'packJobSize' is less than or equal to zero, it returns an empty 2D slice.
//
// Parameters:
// slice: A slice of any type that needs to be packJobed.
// packJobSize: The size of each packJob. Must be a positive integer.
//
// Returns:
// A 2D slice where each inner slice is of length 'packJobSize'.
// The last inner slice may be shorter if the length of 'slice' is not a multiple of 'packJobSize'.
func ChunkSlice[T any](slice []T, packJobSize int) [][]T {
	var packJobs [][]T
	if packJobSize <= 0 {
		return packJobs
	}

	for i := 0; i < len(slice); i += packJobSize {
		end := i + packJobSize
		if end > len(slice) {
			end = len(slice)
		}
		packJobs = append(packJobs, slice[i:end])
	}

	return packJobs
}

const BatchSize = 100

// GenerateNewPeer is a function that generates a new peer for the IPFS network.
// It generates a new Ed25519 key pair using a secure random source and derives a peer ID from the public key.
// It then marshals the private and public keys into byte slices.
//
// Returns:
// The private key as a byte slice, the public key as a byte slice, the peer ID, and an error if the operation failed.
// If any of the operations (key generation, peer ID derivation, key marshalling) fail, it returns an error wrapped with a descriptive message.
func GenerateNewPeer() ([]byte, []byte, peer.ID, error) {
	private, public, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, nil, "", errors.Wrap(err, "cannot generate new peer")
	}

	peerID, err := peer.IDFromPublicKey(public)
	if err != nil {
		return nil, nil, "", errors.Wrap(err, "cannot generate peer id")
	}

	privateBytes, err := crypto.MarshalPrivateKey(private)
	if err != nil {
		return nil, nil, "", errors.Wrap(err, "cannot marshal private key")
	}

	publicBytes, err := crypto.MarshalPublicKey(public)
	if err != nil {
		return nil, nil, "", errors.Wrap(err, "cannot marshal public key")
	}
	return privateBytes, publicBytes, peerID, nil
}
