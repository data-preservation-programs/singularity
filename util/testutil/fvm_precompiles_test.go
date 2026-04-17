package testutil

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestIntegration_FVMPrecompilesResolveRegisteredAddress(t *testing.T) {
	anvil := StartAnvil(t, CalibnetRPC)

	SetupFVMPrecompiles(t, anvil.RPCURL)

	account0 := crypto.PubkeyToAddress(AnvilFunderKey(t).PublicKey)
	RegisterActorID(t, anvil.RPCURL, account0, 5103)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, anvil.RPCURL)
	require.NoError(t, err)
	defer client.Close()

	result, err := client.CallContract(ctx, ethereum.CallMsg{
		To:   &FVMResolveAddress,
		Data: encodeEAMDelegatedAddress(account0),
	}, nil)
	require.NoError(t, err)
	require.Len(t, result, 32, "resolve precompile should return a single ABI word")
	require.EqualValues(t, 5103, decodeUint64Word(result))

	for name, addr := range map[string]common.Address{
		"RESOLVE_ADDRESS":       FVMResolveAddress,
		"CALL_ACTOR_BY_ID":      FVMCallActorByID,
		"CALL_ACTOR_BY_ADDRESS": FVMCallActorByAddress,
	} {
		code, err := client.CodeAt(ctx, addr, nil)
		require.NoError(t, err)
		require.NotEmpty(t, code, "%s should have code deployed", name)
	}
}

func TestIntegration_FVMActorCallPrecompilesAcceptMinimalSupportedRequests(t *testing.T) {
	anvil := StartAnvil(t, CalibnetRPC)

	SetupFVMPrecompiles(t, anvil.RPCURL)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, anvil.RPCURL)
	require.NoError(t, err)
	defer client.Close()

	result, err := client.CallContract(ctx, ethereum.CallMsg{
		To:   &FVMCallActorByID,
		Data: packCallActorByIDInput(t, 0, 99),
	}, nil)
	require.NoError(t, err)
	exitCode, returnCodec, returnValue := decodeActorCallResponse(t, result)
	require.Zero(t, exitCode)
	require.Zero(t, returnCodec)
	require.Empty(t, returnValue)

	account0 := crypto.PubkeyToAddress(AnvilFunderKey(t).PublicKey)
	result, err = client.CallContract(ctx, ethereum.CallMsg{
		To:   &FVMCallActorByAddress,
		Data: packCallActorByAddressInput(t, 0, encodeEAMDelegatedAddress(account0)),
	}, nil)
	require.NoError(t, err)
	exitCode, returnCodec, returnValue = decodeActorCallResponse(t, result)
	require.Zero(t, exitCode)
	require.Zero(t, returnCodec)
	require.Empty(t, returnValue)
}

func TestIntegration_FVMResolveUnregisteredAddressReturnsEmpty(t *testing.T) {
	anvil := StartAnvil(t, CalibnetRPC)

	SetupFVMPrecompiles(t, anvil.RPCURL)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, anvil.RPCURL)
	require.NoError(t, err)
	defer client.Close()

	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	unregistered := crypto.PubkeyToAddress(key.PublicKey)

	result, err := client.CallContract(ctx, ethereum.CallMsg{
		To:   &FVMResolveAddress,
		Data: encodeEAMDelegatedAddress(unregistered),
	}, nil)
	require.NoError(t, err)
	require.Empty(t, result, "vendored resolve-address mock returns no data for unknown addresses")
}

func decodeUint64Word(result []byte) uint64 {
	return new(big.Int).SetBytes(result).Uint64()
}

func packCallActorByIDInput(t *testing.T, methodNum uint64, target uint64) []byte {
	t.Helper()
	args := abi.Arguments{
		{Type: mustABIType(t, "uint64")},
		{Type: mustABIType(t, "uint256")},
		{Type: mustABIType(t, "uint64")},
		{Type: mustABIType(t, "uint64")},
		{Type: mustABIType(t, "bytes")},
		{Type: mustABIType(t, "uint64")},
	}
	data, err := args.Pack(methodNum, big.NewInt(0), uint64(0), uint64(0), []byte{}, target)
	require.NoError(t, err)
	return data
}

func packCallActorByAddressInput(t *testing.T, methodNum uint64, actorAddress []byte) []byte {
	t.Helper()
	args := abi.Arguments{
		{Type: mustABIType(t, "uint64")},
		{Type: mustABIType(t, "uint256")},
		{Type: mustABIType(t, "uint64")},
		{Type: mustABIType(t, "uint64")},
		{Type: mustABIType(t, "bytes")},
		{Type: mustABIType(t, "bytes")},
	}
	data, err := args.Pack(methodNum, big.NewInt(0), uint64(0), uint64(0), []byte{}, actorAddress)
	require.NoError(t, err)
	return data
}

func decodeActorCallResponse(t *testing.T, result []byte) (int64, uint64, []byte) {
	t.Helper()
	args := abi.Arguments{
		{Type: mustABIType(t, "int256")},
		{Type: mustABIType(t, "uint64")},
		{Type: mustABIType(t, "bytes")},
	}
	values, err := args.Unpack(result)
	require.NoError(t, err)
	require.Len(t, values, 3)
	return values[0].(*big.Int).Int64(), values[1].(uint64), values[2].([]byte)
}

func mustABIType(t *testing.T, typ string) abi.Type {
	t.Helper()
	parsed, err := abi.NewType(typ, "", nil)
	require.NoError(t, err)
	return parsed
}
