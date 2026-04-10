package testutil

import (
	_ "embed"
	"encoding/binary"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

//go:embed fvm_resolve_address_bytecode.txt
var fvmResolveAddressBytecode string

//go:embed fvm_call_actor_by_id_bytecode.txt
var fvmCallActorByIDBytecode string

//go:embed fvm_call_actor_by_address_bytecode.txt
var fvmCallActorByAddressBytecode string

var (
	// canonical Filecoin precompile addresses
	FVMResolveAddress     = common.HexToAddress("0xFE00000000000000000000000000000000000001")
	FVMCallActorByAddress = common.HexToAddress("0xfe00000000000000000000000000000000000003")
	FVMCallActorByID      = common.HexToAddress("0xfe00000000000000000000000000000000000005")
)

// SetupFVMPrecompiles deploys mock Filecoin precompiles on an Anvil fork
// using anvil_setCode. Contracts that call RESOLVE_ADDRESS, CALL_ACTOR_BY_ID,
// etc. get working mock responses instead of reverting on empty addresses.
//
// Mock bytecodes are from github.com/filecoin-project/fvm-solidity/src/mocks/.
// After calling this, use RegisterActorID to map EVM addresses to f0 actor IDs.
func SetupFVMPrecompiles(t *testing.T, rpcURL string) {
	t.Helper()
	anvilSetCode(t, rpcURL, FVMResolveAddress, mustDecodeHex(t, fvmResolveAddressBytecode))
	anvilSetCode(t, rpcURL, FVMCallActorByID, mustDecodeHex(t, fvmCallActorByIDBytecode))
	anvilSetCode(t, rpcURL, FVMCallActorByAddress, mustDecodeHex(t, fvmCallActorByAddressBytecode))
	t.Log("FVM mock precompiles deployed")
}

// RegisterActorID maps an EVM address to a Filecoin actor ID in the mock
// RESOLVE_ADDRESS precompile. The EVM address is encoded as an f410
// delegated address (protocol 0x04, namespace 0x0a = EAM, 20-byte addr).
func RegisterActorID(t *testing.T, rpcURL string, evmAddr common.Address, actorID uint64) {
	t.Helper()

	// f410 encoding: 0x04 (protocol 4, delegated) + 0x0a (namespace 10, EAM) + 20 bytes
	filAddr := make([]byte, 22)
	filAddr[0] = 0x04
	filAddr[1] = 0x0a
	copy(filAddr[2:], evmAddr.Bytes())

	// ABI encode mockResolveAddress(bytes,uint64)
	selector := crypto.Keccak256([]byte("mockResolveAddress(bytes,uint64)"))[:4]
	data := make([]byte, 4+32+32+32+32) // selector + offset + actorID + len + data
	copy(data[0:4], selector)
	data[4+31] = 64                                            // offset to bytes param
	binary.BigEndian.PutUint64(data[4+32+24:4+32+32], actorID) // uint64 actorID
	data[4+64+31] = 22                                         // bytes length
	copy(data[4+96:4+96+22], filAddr)                          // bytes data

	funderAddr := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	AnvilImpersonate(t, rpcURL, funderAddr)
	SendImpersonatedTx(t, rpcURL, funderAddr, FVMResolveAddress, data)
	t.Logf("registered actor ID %d for %s", actorID, evmAddr.Hex())
}

func anvilSetCode(t *testing.T, rpcURL string, addr common.Address, code []byte) {
	t.Helper()
	anvilRPC(t, rpcURL, "anvil_setCode", []any{addr.Hex(), "0x" + hex.EncodeToString(code)})
}

func mustDecodeHex(t *testing.T, s string) []byte {
	t.Helper()
	s = strings.TrimPrefix(strings.TrimSpace(s), "0x")
	b, err := hex.DecodeString(s)
	require.NoError(t, err, "failed to decode hex bytecode")
	return b
}
