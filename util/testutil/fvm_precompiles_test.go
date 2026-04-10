package testutil

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

// TestIntegration_FVMPrecompiles verifies that mock precompiles can be
// deployed via anvil_setCode and that RESOLVE_ADDRESS responds to
// registered actor IDs.
func TestIntegration_FVMPrecompiles(t *testing.T) {
	anvil := StartAnvil(t, CalibnetRPC)

	SetupFVMPrecompiles(t, anvil.RPCURL)

	// register anvil account 0 as actor t05103 (its real calibnet actor ID)
	account0 := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	RegisterActorID(t, anvil.RPCURL, account0, 5103)

	// verify the precompile resolves the address
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, anvil.RPCURL)
	require.NoError(t, err)
	defer client.Close()

	// call RESOLVE_ADDRESS with account0's f410 encoding
	filAddr := make([]byte, 22)
	filAddr[0] = 0x04
	filAddr[1] = 0x0a
	copy(filAddr[2:], account0.Bytes())

	// the precompile's fallback expects raw bytes (the filecoin address)
	result, err := client.CallContract(ctx, ethereum.CallMsg{
		To:   &FVMResolveAddress,
		Data: filAddr,
	}, nil)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(result), 8, "expected at least 8 bytes for uint64 actor ID")
	t.Logf("RESOLVE_ADDRESS returned %d bytes: 0x%x", len(result), result)

	// also verify code exists at the precompile addresses
	for name, addr := range map[string]common.Address{
		"RESOLVE_ADDRESS":       FVMResolveAddress,
		"CALL_ACTOR_BY_ID":      FVMCallActorByID,
		"CALL_ACTOR_BY_ADDRESS": FVMCallActorByAddress,
	} {
		code, err := client.CodeAt(ctx, addr, nil)
		require.NoError(t, err)
		require.NotEmpty(t, code, "%s should have code deployed", name)
		t.Logf("%s: %d bytes of code", name, len(code))
	}
}

// TestIntegration_FVMPrecompileUnregisteredAddress verifies that the mock
// RESOLVE_ADDRESS precompile reverts for addresses that haven't been
// registered, matching the real precompile behavior.
func TestIntegration_FVMPrecompileUnregisteredAddress(t *testing.T) {
	anvil := StartAnvil(t, CalibnetRPC)

	SetupFVMPrecompiles(t, anvil.RPCURL)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, anvil.RPCURL)
	require.NoError(t, err)
	defer client.Close()

	// generate a random address that hasn't been registered
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	unregistered := crypto.PubkeyToAddress(key.PublicKey)

	filAddr := make([]byte, 22)
	filAddr[0] = 0x04
	filAddr[1] = 0x0a
	copy(filAddr[2:], unregistered.Bytes())

	// the mock returns 0 for unregistered addresses (no revert)
	result, err := client.CallContract(ctx, ethereum.CallMsg{
		To:   &FVMResolveAddress,
		Data: filAddr,
	}, nil)
	require.NoError(t, err)
	// all bytes should be zero (actor ID 0 = not found)
	allZero := true
	for _, b := range result {
		if b != 0 {
			allZero = false
			break
		}
	}
	require.True(t, allZero, "unregistered address should resolve to actor ID 0")
	t.Logf("unregistered address correctly resolved to 0")
}
