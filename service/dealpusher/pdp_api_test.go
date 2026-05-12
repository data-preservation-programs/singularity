package dealpusher

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPDPSchedulingConfigValidate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		cfg := PDPSchedulingConfig{
			BatchSize:            100,
			MaxPiecesPerProofSet: 1024,
			PullTimeout:          5 * time.Minute,
		}
		require.NoError(t, cfg.Validate())
	})

	t.Run("invalid (zero values)", func(t *testing.T) {
		cfg := PDPSchedulingConfig{}
		require.Error(t, cfg.Validate())
	})

	t.Run("invalid max pieces per proof set", func(t *testing.T) {
		cfg := PDPSchedulingConfig{
			BatchSize:   100,
			PullTimeout: 5 * time.Minute,
		}
		err := cfg.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "pdp max pieces per proof set must be greater than 0")
	})

	t.Run("invalid pull timeout", func(t *testing.T) {
		cfg := PDPSchedulingConfig{
			BatchSize:            100,
			MaxPiecesPerProofSet: 1024,
		}
		err := cfg.Validate()
		require.Error(t, err)
		require.Contains(t, err.Error(), "pdp pull timeout must be greater than 0")
	})
}
