package testutil

import (
	"context"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

//go:embed mock_allocation_facet_bytecode.txt
var mockAllocationFacetBytecodeHex string

// MockDDOFacet holds the deployed mock AllocationFacet address and
// provides helpers to call mock-specific functions.
type MockDDOFacet struct {
	Address common.Address
	RPCURL  string
}

// minimal ABI for diamondCut, sufficient to construct the calldata for
// adding facets to a Diamond proxy. matches EIP-2535.
const diamondCutABI = `[{
	"inputs": [
		{
			"components": [
				{"internalType": "address", "name": "facetAddress", "type": "address"},
				{"internalType": "uint8", "name": "action", "type": "uint8"},
				{"internalType": "bytes4[]", "name": "functionSelectors", "type": "bytes4[]"}
			],
			"internalType": "struct IDiamondCut.FacetCut[]",
			"name": "_diamondCut",
			"type": "tuple[]"
		},
		{"internalType": "address", "name": "_init", "type": "address"},
		{"internalType": "bytes", "name": "_calldata", "type": "bytes"}
	],
	"name": "diamondCut",
	"outputs": [],
	"stateMutability": "nonpayable",
	"type": "function"
}]`

// facetCut mirrors the IDiamondCut.FacetCut Solidity struct for ABI encoding.
type facetCut struct {
	FacetAddress      common.Address `abi:"facetAddress"`
	Action            uint8          `abi:"action"`
	FunctionSelectors [][4]byte      `abi:"functionSelectors"`
}

// DeployMockAllocationFacet deploys the MockAllocationFacet contract on
// an Anvil fork and performs a Diamond cut to add mock selectors to the
// DDO Diamond proxy. The caller must have already impersonated the
// Diamond owner.
func DeployMockAllocationFacet(t *testing.T, rpcURL string, ddoAddr, owner common.Address) *MockDDOFacet {
	t.Helper()
	ctx := context.Background()

	client, err := ethclient.DialContext(ctx, rpcURL)
	require.NoError(t, err)
	defer client.Close()

	// Deploy the mock facet by sending the creation bytecode from the
	// impersonated owner.
	bytecodeHex := mockAllocationFacetBytecodeHex
	if len(bytecodeHex) > 2 && bytecodeHex[:2] == "0x" {
		bytecodeHex = bytecodeHex[2:]
	}

	nonce, err := client.PendingNonceAt(ctx, owner)
	require.NoError(t, err)
	gasPrice, err := client.SuggestGasPrice(ctx)
	require.NoError(t, err)

	// Send deployment tx via impersonated account
	type txParams struct {
		From     string `json:"from"`
		Data     string `json:"data"`
		Gas      string `json:"gas"`
		GasPrice string `json:"gasPrice"`
		Nonce    string `json:"nonce"`
	}
	params := txParams{
		From:     owner.Hex(),
		Data:     "0x" + bytecodeHex,
		Gas:      fmt.Sprintf("0x%x", 10_000_000),
		GasPrice: fmt.Sprintf("0x%x", gasPrice),
		Nonce:    fmt.Sprintf("0x%x", nonce),
	}

	var txHash string
	result := anvilRPCResult(t, rpcURL, "eth_sendTransaction", []any{params})
	err = json.Unmarshal(result, &txHash)
	require.NoError(t, err, "deploy mock facet tx: %s", string(result))

	receipt := waitForReceipt(t, client, common.HexToHash(txHash))
	require.EqualValues(t, 1, receipt.Status, "mock facet deployment failed")
	mockAddr := receipt.ContractAddress
	require.NotEqual(t, common.Address{}, mockAddr, "mock facet deployment: no contract address")
	t.Logf("MockAllocationFacet deployed at %s", mockAddr.Hex())

	// Diamond cut: add mock selectors to the Diamond proxy.
	// Selectors must match MockAllocationFacet's exposed functions.
	mockSelectors := []string{
		"bdc87581", // mockCreateAllocationRequests
		"76e92deb", // mockCreateRawAllocationRequests
		"2b8e3274", // mockActivateAllocation
		"b44039a0", // mockActivateAllocationWithSector
		"3e7b404c", // setMockMiner
		"617864ed", // mockMinerActorIds
	}
	diamondCutViaRPC(t, rpcURL, ddoAddr, owner, mockAddr, mockSelectors)

	return &MockDDOFacet{Address: mockAddr, RPCURL: rpcURL}
}

// diamondCutViaRPC encodes and sends a diamondCut transaction via impersonation.
// Uses the standard EIP-2535 ABI rather than hand-rolled byte packing.
func diamondCutViaRPC(t *testing.T, rpcURL string, ddoAddr, owner, facetAddr common.Address, selectors []string) {
	t.Helper()

	parsedABI, err := abi.JSON(strings.NewReader(diamondCutABI))
	require.NoError(t, err)

	selectorBytes := make([][4]byte, len(selectors))
	for i, s := range selectors {
		b, err := hex.DecodeString(s)
		require.NoError(t, err, "invalid selector %q", s)
		require.Len(t, b, 4, "selector %q must be 4 bytes", s)
		copy(selectorBytes[i][:], b)
	}

	// FacetCutAction.Add = 0
	cuts := []facetCut{{
		FacetAddress:      facetAddr,
		Action:            0,
		FunctionSelectors: selectorBytes,
	}}

	data, err := parsedABI.Pack("diamondCut", cuts, common.Address{}, []byte{})
	require.NoError(t, err)

	SendImpersonatedTx(t, rpcURL, owner, ddoAddr, data)
	t.Log("Diamond cut completed — mock selectors added")
}

// CallMockCreateAllocations calls mockCreateAllocationRequests on the
// DDO Diamond proxy via the impersonated client wallet.
func (m *MockDDOFacet) CallMockCreateAllocations(t *testing.T, from common.Address, ddoAddr common.Address, pieceInfosABI []byte) common.Hash {
	t.Helper()
	// mockCreateAllocationRequests selector = 0xbdc87581
	selector, _ := hex.DecodeString("bdc87581")
	data := append(selector, pieceInfosABI...)
	return SendImpersonatedTx(t, m.RPCURL, from, ddoAddr, data)
}
