package testutil

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

// anvilAccount0Key is the private key for Anvil's first pre-funded account
// (0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266, 10000 ETH).
// This is a well-known deterministic test key — never use in production.
const anvilAccount0Key = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

// AnvilFunderKey returns the parsed private key for Anvil's first pre-funded account.
func AnvilFunderKey(t *testing.T) *ecdsa.PrivateKey {
	t.Helper()
	key, err := crypto.HexToECDSA(anvilAccount0Key)
	require.NoError(t, err)
	return key
}

// FundEVMWallet sends ETH from Anvil's pre-funded account 0 to the given address.
// amount is in wei. Returns the transaction hash.
func FundEVMWallet(t *testing.T, rpcURL string, to common.Address, amount *big.Int) common.Hash {
	t.Helper()

	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, rpcURL)
	require.NoError(t, err)
	defer client.Close()

	funderKey := AnvilFunderKey(t)
	funderAddr := crypto.PubkeyToAddress(funderKey.PublicKey)

	nonce, err := client.PendingNonceAt(ctx, funderAddr)
	require.NoError(t, err)

	chainID, err := client.ChainID(ctx)
	require.NoError(t, err)

	gasPrice, err := client.SuggestGasPrice(ctx)
	require.NoError(t, err)

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Value:    amount,
		Gas:      21000,
		GasPrice: gasPrice,
	})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), funderKey)
	require.NoError(t, err)

	err = client.SendTransaction(ctx, signedTx)
	require.NoError(t, err)

	// Wait for the transaction to be mined (Anvil uses --block-time 1).
	deadline := time.Now().Add(10 * time.Second)
	for time.Now().Before(deadline) {
		receipt, err := client.TransactionReceipt(ctx, signedTx.Hash())
		if err == nil && receipt != nil {
			require.EqualValues(t, 1, receipt.Status, "funding transaction failed")
			return signedTx.Hash()
		}
		time.Sleep(500 * time.Millisecond)
	}
	t.Fatal("funding transaction not mined within 10s")
	return common.Hash{}
}

// GenerateTestKey creates a fresh ECDSA key pair for testing and returns
// the private key and its corresponding EVM address.
func GenerateTestKey(t *testing.T) (*ecdsa.PrivateKey, common.Address) {
	t.Helper()
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	addr := crypto.PubkeyToAddress(key.PublicKey)
	return key, addr
}

// OneEther is 1e18 wei, useful as a unit for test funding amounts.
var OneEther = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
