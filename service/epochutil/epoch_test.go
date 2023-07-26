package epochutil

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultValue(t *testing.T) {
	err := Initialize(context.Background(), "https://api.node.glif.io/rpc/v0", "")
	require.NoError(t, err)
	require.EqualValues(t, 1598306400, GENESIS_TIMESTAMP)
}

func TestCalibNet(t *testing.T) {
	err := Initialize(context.Background(), "https://api.calibration.node.glif.io/rpc/v0", "")
	require.NoError(t, err)
	require.EqualValues(t, 1667326380, GENESIS_TIMESTAMP)
}
