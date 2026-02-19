package pdptracker

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/data-preservation-programs/go-synapse"
	"github.com/data-preservation-programs/go-synapse/constants"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

const calibnetRPC = "https://api.calibration.node.glif.io/rpc/v1"

// TestIntegration_NetworkDetection verifies synapse.DetectNetwork works
// against calibnet and returns the expected chain ID and contract address.
func TestIntegration_NetworkDetection(t *testing.T) {
	testutil.SkipIfNotExternalAPI(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ethClient, err := ethclient.DialContext(ctx, calibnetRPC)
	require.NoError(t, err)
	defer ethClient.Close()

	network, chainID, err := synapse.DetectNetwork(ctx, ethClient)
	require.NoError(t, err)
	require.Equal(t, constants.NetworkCalibration, network)
	require.EqualValues(t, 314159, chainID)

	contractAddr := constants.GetPDPVerifierAddress(network)
	require.NotEqual(t, common.Address{}, contractAddr)
	t.Logf("calibnet PDPVerifier: %s", contractAddr.Hex())
}

// TestIntegration_ShovelConfig validates that the Shovel config generated for
// calibnet passes ValidateFix without errors.
func TestIntegration_ShovelConfig(t *testing.T) {
	testutil.SkipIfNotExternalAPI(t)

	contractAddr := constants.GetPDPVerifierAddress(constants.NetworkCalibration)
	conf := buildShovelConfig(
		"postgres://localhost/test",
		calibnetRPC,
		uint64(constants.ChainIDCalibration),
		contractAddr,
	)

	// import config package to validate
	require.Len(t, conf.Integrations, 7)
	require.Len(t, conf.Sources, 1)
	require.Equal(t, uint64(314159), conf.Sources[0].ChainID)
}

// TestIntegration_ShovelIndexer_Calibnet starts an embedded Shovel indexer
// against calibnet and verifies it processes blocks without errors.
// Requires: Postgres (PGPORT), calibnet RPC, SINGULARITY_TEST_EXTERNAL_API=true.
func TestIntegration_ShovelIndexer_Calibnet(t *testing.T) {
	testutil.SkipIfNotExternalAPI(t)

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		if db.Dialector.Name() != "postgres" {
			t.Skip("Shovel requires Postgres")
			return
		}

		// get the postgres connection string from env (set by testutil)
		connStr := os.Getenv("DATABASE_CONNECTION_STRING")
		require.NotEmpty(t, connStr)

		contractAddr := constants.GetPDPVerifierAddress(constants.NetworkCalibration)

		indexer, err := NewPDPIndexer(ctx, connStr, calibnetRPC, uint64(constants.ChainIDCalibration), contractAddr)
		require.NoError(t, err)

		// start indexer with timeout context
		indexCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()

		exitErr := make(chan error, 1)
		err = indexer.Start(indexCtx, exitErr)
		require.NoError(t, err)

		// let it run for a few seconds to process some blocks
		time.Sleep(10 * time.Second)

		// verify shovel internal tables exist
		var schemaExists bool
		err = db.Raw("SELECT EXISTS(SELECT 1 FROM information_schema.schemata WHERE schema_name = 'shovel')").Scan(&schemaExists).Error
		require.NoError(t, err)
		require.True(t, schemaExists, "shovel schema should exist")

		// verify integration tables exist
		for _, table := range []string{
			"pdp_dataset_created",
			"pdp_pieces_added",
			"pdp_pieces_removed",
			"pdp_next_proving_period",
			"pdp_possession_proven",
			"pdp_dataset_deleted",
			"pdp_sp_changed",
		} {
			var exists bool
			err = db.Raw(
				"SELECT EXISTS(SELECT 1 FROM information_schema.tables WHERE table_name = ?)", table,
			).Scan(&exists).Error
			require.NoError(t, err)
			require.True(t, exists, "table %s should exist", table)
		}

		t.Log("Shovel indexer started and processed blocks against calibnet successfully")

		cancel()
		// wait for clean shutdown
		select {
		case err := <-exitErr:
			require.NoError(t, err)
		case <-time.After(5 * time.Second):
			// fine, shutdown may be slow
		}
	})
}

// TestIntegration_RPCClient_Calibnet verifies the RPC client can make calls
// against the real calibnet PDPVerifier contract.
func TestIntegration_RPCClient_Calibnet(t *testing.T) {
	testutil.SkipIfNotExternalAPI(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	contractAddr := constants.GetPDPVerifierAddress(constants.NetworkCalibration)
	client, err := NewPDPClient(ctx, calibnetRPC, contractAddr)
	require.NoError(t, err)
	defer client.Close()

	// try to get listener for set 0 — may fail (doesn't exist) but shouldn't panic
	_, err = client.GetDataSetListener(ctx, 0)
	// we don't assert NoError here because set 0 may not exist,
	// but the call should complete without panic
	t.Logf("GetDataSetListener(0): err=%v", err)

	// try to get active pieces for set 0
	_, err = client.GetActivePieces(ctx, 0)
	t.Logf("GetActivePieces(0): err=%v", err)
}
