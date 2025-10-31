package epochutil

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
)

func TestDefaultValue(t *testing.T) {
	err := Initialize(context.Background(), testutil.TestLotusAPI, "")
	require.NoError(t, err)
	require.EqualValues(t, 1598306400, GenesisTimestamp)
}

func TestCalibNet(t *testing.T) {
	// This test may fail when calibnet resets
	err := Initialize(context.Background(), "https://api.calibration.node.glif.io/rpc/v1", "")
	require.NoError(t, err)
	require.EqualValues(t, 1667326380, GenesisTimestamp)
}

func TestEpochToTime(t *testing.T) {
	err := Initialize(context.Background(), testutil.TestLotusAPI, "")
	require.NoError(t, err)
	require.EqualValues(t, 1598306400, GenesisTimestamp)
	require.EqualValues(t, 1598306400, EpochToTime(0).Unix())
	require.EqualValues(t, 1598306430, EpochToTime(1).Unix())
}

func TestUnixToEpoch(t *testing.T) {
	err := Initialize(context.Background(), testutil.TestLotusAPI, "")
	require.NoError(t, err)
	require.EqualValues(t, 1598306400, GenesisTimestamp)
	require.EqualValues(t, 0, UnixToEpoch(1598306400))
	require.EqualValues(t, 1, UnixToEpoch(1598306430))
}

func TestTimeToEpoch(t *testing.T) {
	err := Initialize(context.Background(), testutil.TestLotusAPI, "")
	require.NoError(t, err)
	require.EqualValues(t, 1598306400, GenesisTimestamp)
	require.EqualValues(t, 0, TimeToEpoch(EpochToTime(0)))
	require.EqualValues(t, 1, TimeToEpoch(EpochToTime(1)))
}
