package testutil

import (
	"context"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

	// Wait for receipt to get contract address
	hash := common.HexToHash(txHash)
	var mockAddr common.Address
	for range 30 {
		receipt, err := client.TransactionReceipt(ctx, hash)
		if err == nil {
			require.EqualValues(t, 1, receipt.Status, "mock facet deployment failed")
			mockAddr = receipt.ContractAddress
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	require.NotEqual(t, common.Address{}, mockAddr, "mock facet deployment: no contract address")
	t.Logf("MockAllocationFacet deployed at %s", mockAddr.Hex())

	// Diamond cut: add mock selectors to the Diamond proxy.
	// diamondCut(FacetCut[] calldata, address, bytes calldata)
	// FacetCut { address facetAddress; FacetCutAction action; bytes4[] functionSelectors }
	// FacetCutAction.Add = 0
	mockSelectors := []string{
		"bdc87581", // mockCreateAllocationRequests
		"76e92deb", // mockCreateRawAllocationRequests
		"2b8e3274", // mockActivateAllocation
		"b44039a0", // mockActivateAllocationWithSector
		"3e7b404c", // setMockMiner
		"617864ed", // mockMinerActorIds
	}

	diamondCutSig := crypto.Keccak256([]byte("diamondCut((address,uint8,bytes4[])[],address,bytes)"))[:4]

	// ABI encode: FacetCut[] with 1 element (Add action)
	// This is complex ABI encoding. Let's use a simpler approach via cast.
	selectorsHex := ""
	for _, s := range mockSelectors {
		selectorsHex += s
	}
	_ = selectorsHex
	_ = diamondCutSig

	// Use cast to do the ABI encoding and send the tx
	diamondCutViaRPC(t, rpcURL, ddoAddr, owner, mockAddr, mockSelectors)

	return &MockDDOFacet{Address: mockAddr, RPCURL: rpcURL}
}

// diamondCutViaRPC encodes and sends a diamondCut transaction via impersonation.
func diamondCutViaRPC(t *testing.T, rpcURL string, ddoAddr, owner, facetAddr common.Address, selectors []string) {
	t.Helper()

	// Manually ABI-encode diamondCut((address,uint8,bytes4[])[],address,bytes)
	// with one FacetCut: { facetAddress, Add(0), selectors }
	//
	// Layout:
	// 4 bytes: function selector
	// 32 bytes: offset to FacetCut[] (dynamic)
	// 32 bytes: _init address (zero)
	// 32 bytes: offset to _calldata (dynamic)
	// -- FacetCut[] --
	// 32 bytes: array length (1)
	// 32 bytes: offset to first element
	// -- FacetCut struct --
	// 32 bytes: facetAddress
	// 32 bytes: action (0 = Add)
	// 32 bytes: offset to bytes4[] (relative to struct start)
	// -- bytes4[] --
	// 32 bytes: array length
	// N * 32 bytes: each selector padded to 32 bytes

	funcSig := crypto.Keccak256([]byte("diamondCut((address,uint8,bytes4[])[],address,bytes)"))[:4]

	// Pre-compute sizes
	numSelectors := len(selectors)

	// Build calldata manually
	data := make([]byte, 0, 4+32*20) // generous initial capacity

	// Function selector
	data = append(data, funcSig...)

	// Param 1: offset to FacetCut[] array (3 params * 32 = 96 = 0x60)
	data = append(data, padLeft(big.NewInt(0x60).Bytes(), 32)...)
	// Param 2: _init address (zero)
	data = append(data, padLeft(nil, 32)...)
	// Param 3: offset to _calldata (we'll compute after)
	// Placeholder — we'll fill this in after computing the FacetCut[] size
	calldataOffsetPos := len(data)
	data = append(data, padLeft(nil, 32)...) // placeholder

	// FacetCut[] array
	// Array length = 1
	data = append(data, padLeft(big.NewInt(1).Bytes(), 32)...)
	// Offset to first element (32 bytes from array start)
	data = append(data, padLeft(big.NewInt(0x20).Bytes(), 32)...)

	// FacetCut struct (tuple)
	// facetAddress
	data = append(data, padLeft(facetAddr.Bytes(), 32)...)
	// action = 0 (Add)
	data = append(data, padLeft(nil, 32)...)
	// offset to bytes4[] (relative to struct start = 3 * 32 = 96 = 0x60)
	data = append(data, padLeft(big.NewInt(0x60).Bytes(), 32)...)

	// bytes4[] array
	data = append(data, padLeft(big.NewInt(int64(numSelectors)).Bytes(), 32)...)
	for _, sel := range selectors {
		selBytes, err := hex.DecodeString(sel)
		require.NoError(t, err)
		// bytes4 is left-aligned in 32 bytes
		padded := make([]byte, 32)
		copy(padded[0:4], selBytes)
		data = append(data, padded...)
	}

	// Now fill in the _calldata offset
	// _calldata is empty bytes, offset points to after all the FacetCut[] data
	calldataOffset := len(data) - 4 // subtract function selector
	copy(data[calldataOffsetPos:calldataOffsetPos+32], padLeft(big.NewInt(int64(calldataOffset)).Bytes(), 32))

	// Empty bytes _calldata
	data = append(data, padLeft(nil, 32)...) // length = 0

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

func padLeft(b []byte, size int) []byte {
	if len(b) >= size {
		return b[:size]
	}
	padded := make([]byte, size)
	copy(padded[size-len(b):], b)
	return padded
}
