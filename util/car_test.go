package util

import (
	"testing"

	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
)

func TestGenerateCarHeader(t *testing.T) {
	header, err := GenerateCarHeader(testutil.TestCid)
	require.NoError(t, err)
	require.Len(t, header, 59)
}
