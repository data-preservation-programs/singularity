package epochutil

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultValue(t *testing.T) {
	// Skip if network is not available
	err := Initialize(context.Background(), "https://api.node.glif.io/rpc/v0", "")
	if err != nil {
		t.Skipf("Skipping test because network connection failed: %v", err)
	}
	require.EqualValues(t, 1598306400, GenesisTimestamp)
}

func TestCalibNet(t *testing.T) {
	// This test may fail when calibnet resets
	err := Initialize(context.Background(), "https://api.calibration.node.glif.io/rpc/v0", "")
	if err != nil {
		t.Skipf("Skipping test because network connection failed: %v", err)
	}
	require.EqualValues(t, 1667326380, GenesisTimestamp)
}

func TestEpochToTime(t *testing.T) {
	// Test with mock data
	GenesisTimestamp = int32(1598306400)
	require.EqualValues(t, 1598306400, EpochToTime(0).Unix())
	require.EqualValues(t, 1598306430, EpochToTime(1).Unix())
}

func TestUnixToEpoch(t *testing.T) {
	// Test with mock data
	GenesisTimestamp = int32(1598306400)
	require.EqualValues(t, 0, UnixToEpoch(1598306400))
	require.EqualValues(t, 1, UnixToEpoch(1598306430))
}

func TestTimeToEpoch(t *testing.T) {
	// Test with mock data
	GenesisTimestamp = int32(1598306400)
	require.EqualValues(t, 0, TimeToEpoch(EpochToTime(0)))
	require.EqualValues(t, 1, TimeToEpoch(EpochToTime(1)))
}
