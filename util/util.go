package util

import "github.com/ybbus/jsonrpc/v3"

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

func ChunkMapKeys[T1 comparable, T2 any](m map[T1]T2, chunkSize int) [][]T1 {
	keys := make([]T1, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return ChunkSlice(keys, chunkSize)
}

func ChunkSlice[T any](slice []T, chunkSize int) [][]T {
	var chunks [][]T

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
