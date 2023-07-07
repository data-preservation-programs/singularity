package device

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathWithMostSpace(t *testing.T) {
	paths := []string{"/", "/tmp", "/var"}
	path, err := GetPathWithMostSpace(paths)
	require.NoError(t, err)
	require.NotEmpty(t, path)
}
