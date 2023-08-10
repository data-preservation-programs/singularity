package util

import "github.com/ybbus/jsonrpc/v3"

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

// ChunkMapKeys is a generic function that takes a map with keys of any comparable type and values of any type, and an integer as input.
// It divides the keys of the input map into chunks of size 'chunkSize' and returns a 2D slice of keys.
// It uses the ChunkSlice function to divide the keys into chunks.
//
// Parameters:
// m: A map with keys of any comparable type and values of any type. The keys of this map will be chunked.
// chunkSize: The size of each chunk. Must be a positive integer.
//
// Returns:
// A 2D slice where each inner slice is of length 'chunkSize'. The last inner slice may be shorter if the number of keys in 'm' is not a multiple of 'chunkSize'.
func ChunkMapKeys[T1 comparable, T2 any](m map[T1]T2, chunkSize int) [][]T1 {
	keys := make([]T1, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return ChunkSlice(keys, chunkSize)
}

// ChunkSlice is a generic function that takes a slice of any type and an integer as input.
// It divides the input slice into chunks of size 'chunkSize' and returns a 2D slice.
// If 'chunkSize' is less than or equal to zero, it returns an empty 2D slice.
//
// Parameters:
// slice: A slice of any type that needs to be chunked.
// chunkSize: The size of each chunk. Must be a positive integer.
//
// Returns:
// A 2D slice where each inner slice is of length 'chunkSize'.
// The last inner slice may be shorter if the length of 'slice' is not a multiple of 'chunkSize'.
func ChunkSlice[T any](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	if chunkSize <= 0 {
		return chunks
	}

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

const BatchSize = 100
