package dealpusher

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidatePDPProofSetPieceSize(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		pieceSize int64
		wantError string
	}{
		{
			name:      "valid",
			pieceSize: 1 << 20,
		},
		{
			name:      "zero",
			pieceSize: 0,
			wantError: "must be greater than 0",
		},
		{
			name:      "negative",
			pieceSize: -1024,
			wantError: "must be greater than 0",
		},
		{
			name:      "non_power_of_two",
			pieceSize: 1536,
			wantError: "must be a power of two",
		},
		{
			name:      "too_large",
			pieceSize: 1 << 31,
			wantError: "exceeds max allowed",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := validatePDPProofSetPieceSize(tc.pieceSize)
			if tc.wantError == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			require.ErrorContains(t, err, tc.wantError)
		})
	}
}
