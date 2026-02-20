package dealpusher

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPDPSchedulingConfigValidate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		cfg := PDPSchedulingConfig{
			BatchSize:         100,
			GasLimit:          5000000,
			ConfirmationDepth: 5,
			PollingInterval:   30 * time.Second,
		}
		require.NoError(t, cfg.Validate())
	})

	t.Run("invalid", func(t *testing.T) {
		cfg := PDPSchedulingConfig{}
		require.Error(t, cfg.Validate())
	})

	t.Run("invalid gas limit", func(t *testing.T) {
		cfg := PDPSchedulingConfig{
			BatchSize:         100,
			ConfirmationDepth: 5,
			PollingInterval:   30 * time.Second,
		}
		err := cfg.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "pdp gas limit must be greater than 0")
	})
}
