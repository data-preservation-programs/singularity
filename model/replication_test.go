package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStoragePricePerEpochToPricePerDeal(t *testing.T) {
	tests := []struct {
		price          string
		dealSize       int64
		durationEpoch  int32
		expectedResult float64
	}{
		{
			price:          "1000000000000000000", // 1 in terms of 1e18
			dealSize:       1 << 35,
			durationEpoch:  1,
			expectedResult: 1.0,
		},
		// Add more test cases as required
	}

	for _, test := range tests {
		result := StoragePricePerEpochToPricePerDeal(test.price, test.dealSize, test.durationEpoch)
		require.Equal(t, test.expectedResult, result)
	}
}
