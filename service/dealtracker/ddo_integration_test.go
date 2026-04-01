package dealtracker

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
)

func startCalibnetFork(t *testing.T) *testutil.AnvilInstance {
	t.Helper()
	return testutil.StartAnvil(t, testutil.CalibnetRPC)
}

// TestIntegration_DDOTrackingClientConnectivity verifies that the DDO tracking
// client can connect to a calibnet fork and query allocation info.
func TestIntegration_DDOTrackingClientConnectivity(t *testing.T) {
	anvil := startCalibnetFork(t)
	rpcURL := anvil.RPCURL

	client, err := NewDDOTrackingClient(rpcURL, testutil.CalibnetDDOContract)
	require.NoError(t, err)
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Query a non-existent allocation — should return non-nil status, not error.
	status, err := client.GetAllocationInfo(ctx, 0)
	require.NoError(t, err)
	require.NotNil(t, status)
	require.False(t, status.Activated)
	t.Logf("GetAllocationInfo(0): activated=%v, sectorNumber=%d",
		status.Activated, status.SectorNumber)
}
