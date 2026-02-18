package pdptracker

import (
	"time"

	"github.com/cockroachdb/errors"
)

// PDPConfig configures the PDP tracker operations layer.
type PDPConfig struct {
	PollingInterval time.Duration
}

// Validate ensures the PDPConfig values are sane.
func (c PDPConfig) Validate() error {
	if c.PollingInterval <= 0 {
		return errors.New("pdp polling interval must be greater than 0")
	}
	return nil
}
