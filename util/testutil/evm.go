package testutil

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

// AnvilImpersonate enables Anvil's impersonation for the given address,
// allowing transactions to be sent from it without a private key.
func AnvilImpersonate(t *testing.T, rpcURL string, addr common.Address) {
	t.Helper()
	anvilRPC(t, rpcURL, "anvil_impersonateAccount", []any{addr.Hex()})
	t.Cleanup(func() {
		anvilRPC(t, rpcURL, "anvil_stopImpersonatingAccount", []any{addr.Hex()})
	})
}

// ReadContractOwner reads the Diamond contract owner by calling the owner() function.
func ReadContractOwner(t *testing.T, rpcURL string, contractAddr common.Address) common.Address {
	t.Helper()
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, rpcURL)
	require.NoError(t, err)
	defer client.Close()

	selector := crypto.Keccak256([]byte("owner()"))[:4]
	result, err := client.CallContract(ctx, ethereum.CallMsg{
		To:   &contractAddr,
		Data: selector,
	}, nil)
	require.NoError(t, err)
	require.True(t, len(result) >= 32, "owner() returned %d bytes", len(result))
	return common.BytesToAddress(result[12:32])
}

// TransferERC20 sends ERC20 tokens from Anvil's pre-funded account 0 to
// a recipient. This works because account 0 already has a real private key
// (the well-known Anvil test key).
func TransferERC20(t *testing.T, rpcURL string, tokenAddr, recipient common.Address, amount *big.Int) {
	t.Helper()
	ctx := context.Background()

	client, err := ethclient.DialContext(ctx, rpcURL)
	require.NoError(t, err)
	defer client.Close()

	funderKey := AnvilFunderKey(t)
	funderAddr := crypto.PubkeyToAddress(funderKey.PublicKey)

	// ABI-encode transfer(address,uint256)
	transferSelector := crypto.Keccak256([]byte("transfer(address,uint256)"))[:4]
	data := make([]byte, 4+64)
	copy(data[0:4], transferSelector)
	copy(data[4+12:4+32], recipient.Bytes())
	amount.FillBytes(data[4+32 : 4+64])

	nonce, err := client.PendingNonceAt(ctx, funderAddr)
	require.NoError(t, err)
	chainID, err := client.ChainID(ctx)
	require.NoError(t, err)
	gasPrice, err := client.SuggestGasPrice(ctx)
	require.NoError(t, err)

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &tokenAddr,
		Value:    big.NewInt(0),
		Gas:      100000,
		GasPrice: gasPrice,
		Data:     data,
	})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), funderKey)
	require.NoError(t, err)

	err = client.SendTransaction(ctx, signedTx)
	require.NoError(t, err)

	for range 30 {
		receipt, err := client.TransactionReceipt(ctx, signedTx.Hash())
		if err == nil {
			require.EqualValues(t, 1, receipt.Status, "ERC20 transfer failed")
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	t.Fatal("ERC20 transfer not mined after 3s")
}

// SendImpersonatedTx sends a transaction from an impersonated account via
// Anvil's eth_sendTransaction and waits for the receipt. The caller must
// have already called AnvilImpersonate for the from address.
func SendImpersonatedTx(t *testing.T, rpcURL string, from, to common.Address, data []byte) common.Hash {
	t.Helper()
	ctx := context.Background()

	client, err := ethclient.DialContext(ctx, rpcURL)
	require.NoError(t, err)
	defer client.Close()

	nonce, err := client.PendingNonceAt(ctx, from)
	require.NoError(t, err)
	gasPrice, err := client.SuggestGasPrice(ctx)
	require.NoError(t, err)

	type txParams struct {
		From     string `json:"from"`
		To       string `json:"to"`
		Data     string `json:"data"`
		Gas      string `json:"gas"`
		GasPrice string `json:"gasPrice"`
		Nonce    string `json:"nonce"`
	}

	params := txParams{
		From:     from.Hex(),
		To:       to.Hex(),
		Data:     fmt.Sprintf("0x%x", data),
		Gas:      fmt.Sprintf("0x%x", 500000),
		GasPrice: fmt.Sprintf("0x%x", gasPrice),
		Nonce:    fmt.Sprintf("0x%x", nonce),
	}

	var txHash string
	result := anvilRPCResult(t, rpcURL, "eth_sendTransaction", []any{params})
	err = json.Unmarshal(result, &txHash)
	require.NoError(t, err, "failed to parse tx hash: %s", string(result))

	hash := common.HexToHash(txHash)
	for range 30 {
		receipt, err := client.TransactionReceipt(ctx, hash)
		if err == nil {
			require.EqualValues(t, 1, receipt.Status, "impersonated tx %s failed", txHash)
			return hash
		}
		time.Sleep(100 * time.Millisecond)
	}
	t.Fatalf("impersonated tx %s not mined after 3s", txHash)
	return hash
}

func anvilRPC(t *testing.T, rpcURL, method string, params []any) {
	t.Helper()
	anvilRPCResult(t, rpcURL, method, params)
}

func anvilRPCResult(t *testing.T, rpcURL, method string, params []any) json.RawMessage {
	t.Helper()

	body, _ := json.Marshal(map[string]any{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      1,
	})
	resp, err := http.Post(rpcURL, "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, 200, resp.StatusCode)

	var rpcResp struct {
		Result json.RawMessage `json:"result"`
		Error  *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&rpcResp))
	if rpcResp.Error != nil {
		t.Fatalf("rpc %s failed: %s", method, rpcResp.Error.Message)
	}
	return rpcResp.Result
}
