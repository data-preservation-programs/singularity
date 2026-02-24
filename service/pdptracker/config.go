package pdptracker

import (
	"time"

	"github.com/cockroachdb/errors"
)

type PDPConfig struct {
	PollingInterval time.Duration
}

func (c PDPConfig) Validate() error {
	if c.PollingInterval <= 0 {
		return errors.New("pdp polling interval must be greater than 0")
	}
	return nil
}
