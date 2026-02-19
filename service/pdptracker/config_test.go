package pdptracker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPDPConfigValidate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		cfg := PDPConfig{
			PollingInterval: time.Second,
		}
		require.NoError(t, cfg.Validate())
	})

	t.Run("invalid polling interval", func(t *testing.T) {
		cfg := PDPConfig{PollingInterval: 0}
		require.Error(t, cfg.Validate())
	})
}
