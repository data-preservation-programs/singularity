package dealpusher

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/data-preservation-programs/singularity/model"
)

// F05PaidDealManager owns the paid f05 schedule execution path.
// The first scaffold PR wires the type into Singularity; a later PR will
// provide the concrete implementation backed by the Singularity payments contract.
type F05PaidDealManager interface {
	RunSchedule(ctx context.Context, schedule *model.Schedule) (model.ScheduleState, error)
}

// F05PaidSchedulingConfig holds experimental paid-f05 scheduling knobs.
type F05PaidSchedulingConfig struct {
	MinWalletBalanceAttoFIL *big.Int
	SPRegistryAddress       string
	PaymentsAddress         string
}

func defaultF05PaidSchedulingConfig() F05PaidSchedulingConfig {
	return F05PaidSchedulingConfig{
		MinWalletBalanceAttoFIL: big.NewInt(0),
	}
}

// Validate validates paid-f05 scheduling configuration.
func (c F05PaidSchedulingConfig) Validate() error {
	if c.MinWalletBalanceAttoFIL == nil {
		return errors.New("f05 minimum wallet balance must be set")
	}
	if c.MinWalletBalanceAttoFIL.Sign() < 0 {
		return errors.New("f05 minimum wallet balance cannot be negative")
	}
	return nil
}

var attoFIL = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

// ParseFILAmount converts a FIL decimal string into attoFIL.
func ParseFILAmount(value string) (*big.Int, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return big.NewInt(0), nil
	}

	r, ok := new(big.Rat).SetString(value)
	if !ok {
		return nil, errors.New("invalid FIL amount")
	}
	if r.Sign() < 0 {
		return nil, errors.New("FIL amount cannot be negative")
	}

	r.Mul(r, new(big.Rat).SetInt(attoFIL))
	if !r.IsInt() {
		return nil, errors.New("FIL amount must have at most 18 decimal places")
	}

	return new(big.Int).Set(r.Num()), nil
}

func formatAttoFIL(value *big.Int) string {
	if value == nil {
		return "0"
	}
	return new(big.Rat).SetFrac(value, attoFIL).FloatString(18)
}
