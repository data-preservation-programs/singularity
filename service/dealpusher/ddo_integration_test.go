package dealpusher

import (
	"context"
	"math/big"
	"testing"
	"time"

	"strings"

	ddocontract "github.com/Eastore-project/ddo-client/pkg/contract/ddo"
	ddotypes "github.com/Eastore-project/ddo-client/pkg/types"
	"github.com/data-preservation-programs/go-synapse/signer"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

func startCalibnetFork(t *testing.T) *testutil.AnvilInstance {
	t.Helper()
	return testutil.StartAnvil(t, testutil.CalibnetRPC)
}

// TestIntegration_DDOClientConnectivity verifies that OnChainDDO can connect
// to a calibnet fork and detect the correct chain ID.
func TestIntegration_DDOClientConnectivity(t *testing.T) {
	anvil := startCalibnetFork(t)
	rpcURL := anvil.RPCURL

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Verify raw RPC connectivity and chain ID
	ethClient, err := ethclient.DialContext(ctx, rpcURL)
	require.NoError(t, err)
	defer ethClient.Close()

	chainID, err := ethClient.ChainID(ctx)
	require.NoError(t, err)
	require.EqualValues(t, testutil.CalibnetChainID, chainID.Int64())

	// Initialize OnChainDDO client with calibnet contract addresses
	ddo, err := NewOnChainDDO(ctx, rpcURL,
		testutil.CalibnetDDOContract,
		testutil.CalibnetPaymentsContract,
		testutil.CalibnetUSDFC,
	)
	require.NoError(t, err)
	defer ddo.Close()

	require.EqualValues(t, testutil.CalibnetChainID, ddo.chainID.Int64())
	t.Logf("DDO client connected: chainID=%d, ddo=%s, payments=%s",
		ddo.chainID, ddo.ddoContractAddr.Hex(), ddo.paymentsContractAddr.Hex())
}

// TestIntegration_DDOWalletFunding verifies the testutil wallet funding helper
// works against an Anvil fork.
func TestIntegration_DDOWalletFunding(t *testing.T) {
	anvil := startCalibnetFork(t)
	rpcURL := anvil.RPCURL

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Generate a fresh test wallet
	_, addr := testutil.GenerateTestKey(t)

	// Fund it with 10 ETH (FIL on calibnet)
	testutil.FundEVMWallet(t, rpcURL, addr, testutil.OneEther)

	// Verify the balance
	client, err := ethclient.DialContext(ctx, rpcURL)
	require.NoError(t, err)
	defer client.Close()

	balance, err := client.BalanceAt(ctx, addr, nil)
	require.NoError(t, err)
	require.Equal(t, testutil.OneEther, balance)
	t.Logf("funded wallet %s with %s wei", addr.Hex(), balance.String())
}

// TestIntegration_DDOValidateSP exercises the contract read path for SP
// validation, first with a bogus provider (should be inactive) and then
// with the real calibnet FF miner t0178773.
func TestIntegration_DDOValidateSP(t *testing.T) {
	anvil := startCalibnetFork(t)
	rpcURL := anvil.RPCURL

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ddo, err := NewOnChainDDO(ctx, rpcURL,
		testutil.CalibnetDDOContract,
		testutil.CalibnetPaymentsContract,
		testutil.CalibnetUSDFC,
	)
	require.NoError(t, err)
	defer ddo.Close()

	// Bogus provider — should return inactive, not error.
	bogus, err := ddo.ValidateSP(ctx, 99999)
	require.NoError(t, err)
	require.NotNil(t, bogus)
	t.Logf("ValidateSP(99999): active=%v, minPiece=%d, maxPiece=%d",
		bogus.IsActive, bogus.MinPieceSize, bogus.MaxPieceSize)

	// Real calibnet FF miner t0178773.
	real, err := ddo.ValidateSP(ctx, testutil.CalibnetDDOProviderActorID)
	require.NoError(t, err)
	require.NotNil(t, real)
	t.Logf("ValidateSP(%d): active=%v, minPiece=%d, maxPiece=%d, minTerm=%d, maxTerm=%d",
		testutil.CalibnetDDOProviderActorID,
		real.IsActive, real.MinPieceSize, real.MaxPieceSize,
		real.MinTermLen, real.MaxTermLen)
	if real.IsActive {
		require.Greater(t, real.MaxPieceSize, uint64(0))
	}
}

// TestIntegration_DDORegisterSPOnFork verifies that we can register an SP
// in the DDO contract on an Anvil fork by impersonating the contract owner.
func TestIntegration_DDORegisterSPOnFork(t *testing.T) {
	anvil := startCalibnetFork(t)
	rpcURL := anvil.RPCURL

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	ddoAddr := common.HexToAddress(testutil.CalibnetDDOContract)
	usdfcAddr := common.HexToAddress(testutil.CalibnetUSDFC)

	// Read the DDO contract owner
	owner := testutil.ReadContractOwner(t, rpcURL, ddoAddr)
	t.Logf("DDO contract owner: %s", owner.Hex())

	// Impersonate the owner and fund it with gas
	testutil.AnvilImpersonate(t, rpcURL, owner)
	testutil.FundEVMWallet(t, rpcURL, owner, testutil.OneEther)

	// Register the SP via impersonated eth_sendTransaction
	registerSPViaImpersonation(t, rpcURL, ddoAddr, owner, testutil.CalibnetDDOProviderActorID, usdfcAddr)

	// Verify the SP is now registered and active
	ddo, err := NewOnChainDDO(ctx, rpcURL,
		testutil.CalibnetDDOContract,
		testutil.CalibnetPaymentsContract,
		testutil.CalibnetUSDFC,
	)
	require.NoError(t, err)
	defer ddo.Close()

	cfg, err := ddo.ValidateSP(ctx, testutil.CalibnetDDOProviderActorID)
	require.NoError(t, err)
	require.True(t, cfg.IsActive, "SP should be active after registration")
	require.Greater(t, cfg.MaxPieceSize, uint64(0))
	t.Logf("Registered SP %d: active=%v, minPiece=%d, maxPiece=%d",
		testutil.CalibnetDDOProviderActorID, cfg.IsActive, cfg.MinPieceSize, cfg.MaxPieceSize)
}

// TestIntegration_DDOFullDealFlow exercises the complete DDO deal lifecycle:
// register SP → fund wallet → ensure payments → create allocations → parse IDs.
func TestIntegration_DDOFullDealFlow(t *testing.T) {
	anvil := startCalibnetFork(t)
	rpcURL := anvil.RPCURL

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	ddoAddr := common.HexToAddress(testutil.CalibnetDDOContract)
	usdfcAddr := common.HexToAddress(testutil.CalibnetUSDFC)

	// --- Step 1: Register the SP on the fork ---
	owner := testutil.ReadContractOwner(t, rpcURL, ddoAddr)
	testutil.AnvilImpersonate(t, rpcURL, owner)
	testutil.FundEVMWallet(t, rpcURL, owner, testutil.OneEther)
	registerSPViaImpersonation(t, rpcURL, ddoAddr, owner, testutil.CalibnetDDOProviderActorID, usdfcAddr)
	t.Log("SP registered in DDO contract")

	// --- Step 2: Initialize OnChainDDO client ---
	ddo, err := NewOnChainDDO(ctx, rpcURL,
		testutil.CalibnetDDOContract,
		testutil.CalibnetPaymentsContract,
		testutil.CalibnetUSDFC,
	)
	require.NoError(t, err)
	defer ddo.Close()

	// Verify SP is active
	spCfg, err := ddo.ValidateSP(ctx, testutil.CalibnetDDOProviderActorID)
	require.NoError(t, err)
	require.True(t, spCfg.IsActive)
	t.Logf("SP config: minPiece=%d, maxPiece=%d, minTerm=%d, maxTerm=%d",
		spCfg.MinPieceSize, spCfg.MaxPieceSize, spCfg.MinTermLen, spCfg.MaxTermLen)

	// --- Step 3: Use Anvil's pre-funded account 0 as client ---
	// Account 0 already has USDFC tokens and deposited funds on calibnet.
	clientKey := testutil.AnvilFunderKey(t)
	clientSigner, err := signer.NewSecp256k1SignerFromECDSA(clientKey)
	require.NoError(t, err)
	evmSigner, ok := signer.AsEVM(clientSigner)
	require.True(t, ok)
	t.Logf("Client wallet (Anvil account 0): %s", evmSigner.EVMAddress().Hex())

	// --- Step 4: Build test pieces ---
	// Use a valid piece CID (commp)
	pieceCID := calculateCommp(t, generateRandomBytes(1000), 2048)
	pieces := []DDOPieceSubmission{{
		PieceCID:    pieceCID,
		PieceSize:   2048,
		ProviderID:  testutil.CalibnetDDOProviderActorID,
		DownloadURL: "https://example.test/piece/" + pieceCID.String(),
	}}

	cfg := DDOSchedulingConfig{
		BatchSize:         10,
		ConfirmationDepth: 1,
		PollingInterval:   100 * time.Millisecond,
		TermMin:           518400,
		TermMax:           5256000,
		ExpirationOffset:  172800,
	}

	// --- Step 5: EnsurePayments ---
	err = ddo.EnsurePayments(ctx, evmSigner, pieces, cfg)
	require.NoError(t, err)
	t.Log("EnsurePayments succeeded — deposit and operator approval completed on-chain")

	// --- Step 6: Deploy MockAllocationFacet ---
	// The real createAllocationRequests calls DataCapAPI.transfer() which
	// requires Filecoin built-in actors (CALL_ACTOR_ID precompile). Anvil
	// doesn't have those, so we deploy the MockAllocationFacet which skips
	// the DataCap path and generates mock allocation IDs.
	mockFacet := testutil.DeployMockAllocationFacet(t, rpcURL, ddoAddr, owner)

	// --- Step 7: Call mockCreateAllocationRequests ---
	// ABI-encode the pieceInfos for the mock function call.
	pieceInfos, err := ddo.buildPieceInfos(pieces, cfg)
	require.NoError(t, err)

	parsedABI, err := abi.JSON(strings.NewReader(ddocontract.DDOClientABI))
	require.NoError(t, err)

	// Use mockCreateRawAllocationRequests which skips DataCap AND payment
	// rails — it only emits AllocationCreated events with mock IDs.
	// It takes the same parameter shape as createAllocationRequests, so we
	// pack the real method and derive the mock selector by replacing the
	// function name in the canonical signature. This stays in sync if the
	// pieceInfos struct ever changes.
	realCallData, err := parsedABI.Pack("createAllocationRequests", pieceInfos)
	require.NoError(t, err)
	realMethod := parsedABI.Methods["createAllocationRequests"]
	mockSig := strings.Replace(realMethod.Sig, "createAllocationRequests", "mockCreateRawAllocationRequests", 1)
	mockSelector := crypto.Keccak256([]byte(mockSig))[:4]
	copy(realCallData[0:4], mockSelector)

	clientAddr := evmSigner.EVMAddress()
	testutil.AnvilImpersonate(t, rpcURL, clientAddr)
	mockTxHash := testutil.SendImpersonatedTx(t, rpcURL, clientAddr, ddoAddr, realCallData)
	t.Logf("mockCreateAllocationRequests tx: %s", mockTxHash.Hex())

	// --- Step 8: Parse AllocationCreated events ---
	allocationIDs, err := ddo.ParseAllocationIDs(ctx, mockTxHash.Hex())
	require.NoError(t, err)
	require.Len(t, allocationIDs, 1, "should get exactly 1 allocation ID for 1 piece")
	t.Logf("Mock allocation ID: %d", allocationIDs[0])

	_ = mockFacet
}

// registerSPViaImpersonation registers an SP in the DDO contract using
// Anvil's account impersonation. The caller must have already called
// AnvilImpersonate and FundEVMWallet for the owner address.
func registerSPViaImpersonation(
	t *testing.T,
	rpcURL string,
	ddoAddr, owner common.Address,
	actorID uint64,
	paymentToken common.Address,
) {
	t.Helper()

	// ABI-encode the registerSP call
	abiJSON := ddocontract.DDOClientABI
	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	require.NoError(t, err)

	callData, err := parsedABI.Pack("registerSP",
		actorID,
		owner,          // payment address
		uint64(512),    // min piece size
		uint64(1<<35),  // max piece size (32 GiB)
		int64(172800),  // min term (~60 days)
		int64(5256000), // max term (~5 years)
		[]ddotypes.TokenConfig{{
			Token:                paymentToken,
			PricePerBytePerEpoch: big.NewInt(1),
			IsActive:             true,
		}},
	)
	require.NoError(t, err)

	testutil.SendImpersonatedTx(t, rpcURL, owner, ddoAddr, callData)
}

// calculateCommp and generateRandomBytes are defined in dealpusher_test.go
// and available within the same test package.

// Ensure cid is used (it's needed by calculateCommp)
var _ = cid.Undef
