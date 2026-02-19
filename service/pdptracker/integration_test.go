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

func startCalibnetFork(t *testing.T) string {
	t.Helper()
	anvil := testutil.StartAnvil(t, calibnetRPC)
	return anvil.RPCURL
}

func TestIntegration_NetworkDetection(t *testing.T) {
	rpcURL := startCalibnetFork(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ethClient, err := ethclient.DialContext(ctx, rpcURL)
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

func TestIntegration_ShovelConfig(t *testing.T) {
	contractAddr := constants.GetPDPVerifierAddress(constants.NetworkCalibration)
	conf := buildShovelConfig(
		"postgres://localhost/test",
		calibnetRPC,
		uint64(constants.ChainIDCalibration),
		contractAddr,
	)

	require.Len(t, conf.Integrations, 7)
	require.Len(t, conf.Sources, 1)
	require.Equal(t, uint64(314159), conf.Sources[0].ChainID)
}

func TestIntegration_ShovelIndexer(t *testing.T) {
	rpcURL := startCalibnetFork(t)

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		if db.Dialector.Name() != "postgres" {
			t.Skip("Shovel requires Postgres")
			return
		}

		connStr := os.Getenv("DATABASE_CONNECTION_STRING")
		require.NotEmpty(t, connStr)

		contractAddr := constants.GetPDPVerifierAddress(constants.NetworkCalibration)

		indexer, err := NewPDPIndexer(ctx, connStr, rpcURL, uint64(constants.ChainIDCalibration), contractAddr)
		require.NoError(t, err)

		indexCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()

		exitErr := make(chan error, 1)
		err = indexer.Start(indexCtx, exitErr)
		require.NoError(t, err)

		time.Sleep(10 * time.Second)

		var schemaExists bool
		err = db.Raw("SELECT EXISTS(SELECT 1 FROM information_schema.schemata WHERE schema_name = 'shovel')").Scan(&schemaExists).Error
		require.NoError(t, err)
		require.True(t, schemaExists, "shovel schema should exist")

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

		cancel()
		select {
		case err := <-exitErr:
			require.NoError(t, err)
		case <-time.After(5 * time.Second):
		}
	})
}

func TestIntegration_RPCClient(t *testing.T) {
	rpcURL := startCalibnetFork(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	contractAddr := constants.GetPDPVerifierAddress(constants.NetworkCalibration)
	client, err := NewPDPClient(ctx, rpcURL, contractAddr)
	require.NoError(t, err)
	defer client.Close()

	_, err = client.GetDataSetListener(ctx, 0)
	t.Logf("GetDataSetListener(0): err=%v", err)

	_, err = client.GetActivePieces(ctx, 0)
	t.Logf("GetActivePieces(0): err=%v", err)
}
