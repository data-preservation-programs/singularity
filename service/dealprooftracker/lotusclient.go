package dealprooftracker

import (
	"github.com/ybbus/jsonrpc/v3"
)

// NewLotusClient creates a new JSON-RPC client for interacting with a Lotus node.
func NewLotusClient(lotusAPI string, lotusToken string) jsonrpc.RPCClient {
	if lotusToken == "" {
		return jsonrpc.NewClient(lotusAPI)
	}
	return jsonrpc.NewClientWithOpts(lotusAPI, &jsonrpc.RPCClientOpts{
		CustomHeaders: map[string]string{
			"Authorization": "Bearer " + lotusToken,
		},
	})
}
