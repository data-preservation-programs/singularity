package dealpusher

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

// Calibnet DDO contract addresses from
// https://github.com/Eastore-project/ddo-client#filecoin-calibration-testnet
const (
	calibnetRPC              = "https://api.calibration.node.glif.io/rpc/v1"
	calibnetDDOContract      = "0x889fD50196BE300D06dc4b8F0F17fdB0af587095"
	calibnetPaymentsContract = "0x09a0fDc2723fAd1A7b8e3e00eE5DF73841df55a0"
	calibnetUSDFC            = "0xb3042734b608a1B16e9e86B374A3f3e389B4cDf0"
	calibnetChainID          = 314159
)

func startCalibnetFork(t *testing.T) *testutil.AnvilInstance {
	t.Helper()
	return testutil.StartAnvil(t, calibnetRPC)
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
	require.EqualValues(t, calibnetChainID, chainID.Int64())

	// Initialize OnChainDDO client with calibnet contract addresses
	ddo, err := NewOnChainDDO(ctx, rpcURL,
		calibnetDDOContract,
		calibnetPaymentsContract,
		calibnetUSDFC,
	)
	require.NoError(t, err)
	defer ddo.Close()

	require.EqualValues(t, calibnetChainID, ddo.chainID.Int64())
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

// TestIntegration_DDOValidateSP attempts to validate an SP on calibnet.
// This test exercises the contract read path. Without a registered SP,
// ValidateSP returns an empty (inactive) config — we verify the call
// succeeds and returns a well-formed response.
func TestIntegration_DDOValidateSP(t *testing.T) {
	anvil := startCalibnetFork(t)
	rpcURL := anvil.RPCURL

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ddo, err := NewOnChainDDO(ctx, rpcURL,
		calibnetDDOContract,
		calibnetPaymentsContract,
		calibnetUSDFC,
	)
	require.NoError(t, err)
	defer ddo.Close()

	// Use a known-invalid provider ID — should return inactive, not error.
	// When a real calibnet SP is registered, this test should be updated
	// to use that provider ID and assert IsActive=true.
	cfg, err := ddo.ValidateSP(ctx, 99999)
	require.NoError(t, err)
	require.NotNil(t, cfg)
	t.Logf("ValidateSP(99999): active=%v, minPiece=%d, maxPiece=%d",
		cfg.IsActive, cfg.MinPieceSize, cfg.MaxPieceSize)

	// TODO: Once a calibnet SP is registered in the DDO contract, add a
	// test here with the real provider actor ID and assert:
	//   require.True(t, cfg.IsActive)
	//   require.Greater(t, cfg.MaxPieceSize, uint64(0))
}

// TODO: TestIntegration_DDOFullDealFlow
// This test requires a registered, active SP on calibnet. Once FF provides
// the SP and it's registered in the DDO contract:
//
// 1. Fork calibnet via Anvil
// 2. Fund a test wallet with FIL
// 3. Create a test preparation with a piece in the database
// 4. Create a DDO schedule pointing to the funded wallet and the SP
// 5. Run the deal pusher schedule
// 6. Verify allocation was created on-chain
// 7. Initialize DDOTrackingClient
// 8. Verify allocation tracking returns the correct status
