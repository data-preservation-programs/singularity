//go:build !windows

package device

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathWithMostSpace(t *testing.T) {
	home, err := os.UserHomeDir()
	require.NoError(t, err)
	paths := []string{os.TempDir(), home}
	path, err := GetPathWithMostSpace(paths)
	require.NoError(t, err)
	require.NotEmpty(t, path)
}
