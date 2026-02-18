package pdptracker

import (
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPDPConfigValidate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		cfg := PDPConfig{
			BatchSize:         1,
			GasCap:            big.NewInt(1),
			ConfirmationDepth: 1,
			PollingInterval:   time.Second,
		}
		require.NoError(t, cfg.Validate())
	})

	t.Run("invalid batch size", func(t *testing.T) {
		cfg := PDPConfig{BatchSize: 0, ConfirmationDepth: 1, PollingInterval: time.Second}
		require.Error(t, cfg.Validate())
	})

	t.Run("invalid gas cap", func(t *testing.T) {
		cfg := PDPConfig{BatchSize: 1, GasCap: big.NewInt(0), ConfirmationDepth: 1, PollingInterval: time.Second}
		require.Error(t, cfg.Validate())
	})

	t.Run("invalid confirmation depth", func(t *testing.T) {
		cfg := PDPConfig{BatchSize: 1, ConfirmationDepth: 0, PollingInterval: time.Second}
		require.Error(t, cfg.Validate())
	})

	t.Run("invalid polling interval", func(t *testing.T) {
		cfg := PDPConfig{BatchSize: 1, ConfirmationDepth: 1, PollingInterval: 0}
		require.Error(t, cfg.Validate())
	})
}
