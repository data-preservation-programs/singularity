package storagesystem

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBackends(t *testing.T) {
	require.EqualValues(t, 41, len(Backends)) // Was 42 before amazonclouddrive removal
	local := BackendMap["local"]
	require.Equal(t, "local", local.Name)
}
