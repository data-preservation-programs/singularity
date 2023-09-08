package testutil

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomBytes(t *testing.T) {
	bytes := GenerateRandomBytes(0)
	require.Len(t, bytes, 0)
	bytes = GenerateRandomBytes(100)
	require.Len(t, bytes, 100)
}

func TestGetFileTimestamp(t *testing.T) {
	tmp := t.TempDir()
	err := os.WriteFile(filepath.Join(tmp, "test.txt"), []byte("test"), 0644)
	require.NoError(t, err)
	timestamp := GetFileTimestamp(t, filepath.Join(tmp, "test.txt"))
	require.NotZero(t, timestamp)
}
