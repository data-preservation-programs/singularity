package pdptracker

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexsupply/shovel/shovel/config"
	"github.com/stretchr/testify/require"
)

func TestBuildShovelConfig(t *testing.T) {
	contract := common.HexToAddress("0xBADd0B92C1c71d02E7d520f64c0876538fa2557F")
	conf := buildShovelConfig("postgres://localhost/test", "https://rpc.example.com", 314, contract)
	require.NoError(t, config.ValidateFix(&conf))

	require.Len(t, conf.Sources, 1)
	require.Equal(t, "fevm", conf.Sources[0].Name)
	require.Equal(t, uint64(314), conf.Sources[0].ChainID)
	require.Len(t, conf.Sources[0].URLs, 1)

	// 7 event integrations
	require.Len(t, conf.Integrations, 7)

	names := make(map[string]bool)
	for _, ig := range conf.Integrations {
		names[ig.Name] = true
		require.True(t, ig.Enabled)
		require.Len(t, ig.Sources, 1)
		require.Equal(t, "fevm", ig.Sources[0].Name)

		// each integration must have a contract address filter
		require.NotEmpty(t, ig.Block)
		require.Equal(t, "log_addr", ig.Block[0].Name)
		require.Equal(t, "contains", ig.Block[0].Filter.Op)
		require.Contains(t, ig.Block[0].Filter.Arg[0], "0xbadd0b92c1c71d02e7d520f64c0876538fa2557f")
	}

	expectedNames := []string{
		"pdp_dataset_created",
		"pdp_pieces_added",
		"pdp_pieces_removed",
		"pdp_next_proving_period",
		"pdp_possession_proven",
		"pdp_dataset_deleted",
		"pdp_sp_changed",
	}
	for _, name := range expectedNames {
		require.True(t, names[name], "missing integration: %s", name)
	}
}

func TestBuildShovelConfig_EventInputs(t *testing.T) {
	contract := common.HexToAddress("0x85e366Cf9DD2c0aE37E963d9556F5f4718d6417C")
	conf := buildShovelConfig("postgres://localhost/test", "https://rpc.example.com", 314159, contract)

	// find DataSetCreated and verify inputs
	for _, ig := range conf.Integrations {
		if ig.Name == "pdp_dataset_created" {
			require.Equal(t, "DataSetCreated", ig.Event.Name)
			require.Len(t, ig.Event.Inputs, 2)
			require.True(t, ig.Event.Inputs[0].Indexed)
			require.Equal(t, "uint256", ig.Event.Inputs[0].Type)
			require.Equal(t, "set_id", ig.Event.Inputs[0].Column)
			require.True(t, ig.Event.Inputs[1].Indexed)
			require.Equal(t, "address", ig.Event.Inputs[1].Type)
			require.Equal(t, "storage_provider", ig.Event.Inputs[1].Column)
		}

		if ig.Name == "pdp_pieces_added" {
			require.Equal(t, "PiecesAdded", ig.Event.Name)
			require.Len(t, ig.Event.Inputs, 3)
			// only first input selected
			require.NotEmpty(t, ig.Event.Inputs[0].Column)
			require.Empty(t, ig.Event.Inputs[1].Column)
			require.Empty(t, ig.Event.Inputs[2].Column)
		}

		if ig.Name == "pdp_next_proving_period" {
			require.Equal(t, "NextProvingPeriod", ig.Event.Name)
			require.Len(t, ig.Event.Inputs, 3)
			require.Equal(t, "set_id", ig.Event.Inputs[0].Column)
			require.Equal(t, "challenge_epoch", ig.Event.Inputs[1].Column)
			require.Equal(t, "leaf_count", ig.Event.Inputs[2].Column)
		}
	}
}
