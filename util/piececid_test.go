package util

import (
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

func TestIsAcceptedPieceCID(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want bool
	}{
		{
			name: "legacy commp",
			raw:  "baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq",
			want: true,
		},
		{
			name: "commpv2",
			raw:  "bafkzcibd6adqmxxmxtat74d6vqbczd5v3kvg5mdbhtkuawcyfeqvegqd4skzqqy3",
			want: true,
		},
		{
			name: "non-piece cid",
			raw:  "bafybeiejlvvmfokp5c6q2eqgbfjeaokz3nqho5c7yy3ov527vsatgsqfma",
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parsed, err := cid.Parse(tc.raw)
			require.NoError(t, err)
			require.Equal(t, tc.want, IsAcceptedPieceCID(parsed))
		})
	}
}
