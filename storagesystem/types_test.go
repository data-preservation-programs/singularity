package storagesystem

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBackends(t *testing.T) {
	require.EqualValues(t, 40, len(Backends))
	local := BackendMap["local"]
	require.Equal(t, "local", local.Name)
}
