package dealpusher

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestF05PaidSchedulingConfigValidate(t *testing.T) {
	require.NoError(t, defaultF05PaidSchedulingConfig().Validate())

	err := (F05PaidSchedulingConfig{}).Validate()
	require.ErrorContains(t, err, "must be set")

	err = (F05PaidSchedulingConfig{MinWalletBalanceAttoFIL: big.NewInt(-1)}).Validate()
	require.ErrorContains(t, err, "cannot be negative")
}

func TestParseFILAmount(t *testing.T) {
	value, err := ParseFILAmount("1.5")
	require.NoError(t, err)
	require.Equal(t, "1500000000000000000", value.String())

	value, err = ParseFILAmount("")
	require.NoError(t, err)
	require.Zero(t, value.Sign())

	_, err = ParseFILAmount("0.1234567890123456789")
	require.ErrorContains(t, err, "at most 18 decimal places")

	_, err = ParseFILAmount("-1")
	require.ErrorContains(t, err, "cannot be negative")
}
