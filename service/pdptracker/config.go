package pdptracker

import (
	"math/big"
	"time"

	"github.com/cockroachdb/errors"
)

// PDPConfig configures the PDP tracker operations layer.
type PDPConfig struct {
	BatchSize         int
	GasCap            *big.Int
	ConfirmationDepth uint64
	PollingInterval   time.Duration
}

// Validate ensures the PDPConfig values are sane.
func (c PDPConfig) Validate() error {
	if c.BatchSize <= 0 {
		return errors.New("pdp batch size must be greater than 0")
	}
	if c.GasCap != nil && c.GasCap.Sign() <= 0 {
		return errors.New("pdp gas cap must be greater than 0")
	}
	if c.ConfirmationDepth == 0 {
		return errors.New("pdp confirmation depth must be greater than 0")
	}
	if c.PollingInterval <= 0 {
		return errors.New("pdp polling interval must be greater than 0")
	}
	return nil
}
