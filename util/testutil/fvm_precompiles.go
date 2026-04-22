package testutil

import (
	_ "embed"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

//go:embed fvm_resolve_address_bytecode.txt
var fvmResolveAddressBytecodeHex string

//go:embed fvm_call_actor_by_id_bytecode.txt
var fvmCallActorByIDBytecodeHex string

//go:embed fvm_call_actor_by_address_bytecode.txt
var fvmCallActorByAddressBytecodeHex string

var (
	FVMResolveAddress     = common.HexToAddress("0xfe00000000000000000000000000000000000001")
	FVMCallActorByAddress = common.HexToAddress("0xfe00000000000000000000000000000000000003")
	FVMCallActorByID      = common.HexToAddress("0xfe00000000000000000000000000000000000005")
)

// decoded once at init so SetupFVMPrecompiles doesn't re-parse ~4KB hex per test
var (
	fvmResolveAddressBytecode     = mustDecodeBytecode(fvmResolveAddressBytecodeHex)
	fvmCallActorByIDBytecode      = mustDecodeBytecode(fvmCallActorByIDBytecodeHex)
	fvmCallActorByAddressBytecode = mustDecodeBytecode(fvmCallActorByAddressBytecodeHex)

	mockResolveAddressSelector = crypto.Keccak256([]byte("mockResolveAddress(bytes,uint64)"))[:4]
)

// SetupFVMPrecompiles deploys narrow Filecoin precompile mocks on an Anvil fork.
//
// These mocks are useful for tests that would otherwise revert on empty Filecoin
// precompile addresses. They cover address resolution plus minimal actor-call
// happy paths, but they do not emulate DataCap / VerifReg flows end-to-end.
func SetupFVMPrecompiles(t *testing.T, rpcURL string) {
	t.Helper()
	anvilSetCode(t, rpcURL, FVMResolveAddress, fvmResolveAddressBytecode)
	anvilSetCode(t, rpcURL, FVMCallActorByID, fvmCallActorByIDBytecode)
	anvilSetCode(t, rpcURL, FVMCallActorByAddress, fvmCallActorByAddressBytecode)
	t.Log("FVM mock precompiles deployed")
}

// RegisterActorID adds an EVM address -> actor ID mapping to the mocked
// RESOLVE_ADDRESS precompile using the Filecoin delegated-address encoding.
func RegisterActorID(t *testing.T, rpcURL string, evmAddr common.Address, actorID uint64) {
	t.Helper()

	filAddr := encodeEAMDelegatedAddress(evmAddr)
	// single-byte length field below; addresses must fit
	require.LessOrEqual(t, len(filAddr), 255, "filAddr length exceeds single-byte ABI length field")

	// abi encode mockResolveAddress(bytes,uint64)
	data := make([]byte, 4+32+32+32+32)
	copy(data[0:4], mockResolveAddressSelector)
	data[4+31] = 64
	binary.BigEndian.PutUint64(data[4+32+24:4+32+32], actorID)
	data[4+64+31] = byte(len(filAddr))
	copy(data[4+96:4+96+len(filAddr)], filAddr)

	funderAddr := crypto.PubkeyToAddress(AnvilFunderKey(t).PublicKey)
	AnvilImpersonate(t, rpcURL, funderAddr)
	SendImpersonatedTx(t, rpcURL, funderAddr, FVMResolveAddress, data)
	t.Logf("registered actor ID %d for %s", actorID, evmAddr.Hex())
}

func encodeEAMDelegatedAddress(evmAddr common.Address) []byte {
	filAddr := make([]byte, 22)
	filAddr[0] = 0x04
	filAddr[1] = 0x0a
	copy(filAddr[2:], evmAddr.Bytes())
	return filAddr
}

func anvilSetCode(t *testing.T, rpcURL string, addr common.Address, code []byte) {
	t.Helper()
	anvilRPC(t, rpcURL, "anvil_setCode", []any{addr.Hex(), "0x" + hex.EncodeToString(code)})
}

func mustDecodeBytecode(s string) []byte {
	s = strings.TrimPrefix(strings.TrimSpace(s), "0x")
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(fmt.Sprintf("failed to decode embedded bytecode: %v", err))
	}
	return b
}
