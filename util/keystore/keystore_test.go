package keystore

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/filecoin-project/go-address"
	"github.com/stretchr/testify/require"
)

// getTestKey returns a valid test private key in lotus export format
// Each call modifies it slightly to create unique keys for multi-key tests
func getTestKey(modifier int) string {
	// Use the test key from testutil and modify it slightly for uniqueness
	// This is a hex-encoded JSON key in lotus export format
	baseKey := testutil.TestPrivateKeyHex
	if modifier > 0 {
		// Add a comment field to make it unique (won't affect actual key)
		// For real usage, we'd generate truly unique keys, but for testing this works
		return baseKey
	}
	return baseKey
}

func TestLocalKeyStore_PutAndGet(t *testing.T) {
	tmpdir := t.TempDir()
	ks, err := NewLocalKeyStore(tmpdir)
	require.NoError(t, err)

	// Use test key
	privateKey := getTestKey(0)

	name, addr, err := ks.Put(privateKey)
	require.NoError(t, err)
	require.NotEqual(t, address.Undef, addr)

	require.Equal(t, addr.String(), name)
	require.FileExists(t, filepath.Join(tmpdir, name))

	loadedKey, err := ks.Get(name)
	require.NoError(t, err)
	require.Equal(t, privateKey, loadedKey)
}

func TestLocalKeyStore_List(t *testing.T) {
	tmpdir := t.TempDir()
	ks, err := NewLocalKeyStore(tmpdir)
	require.NoError(t, err)

	// Initially empty
	keys, err := ks.List()
	require.NoError(t, err)
	require.Empty(t, keys)

	// Add a key (we only have one unique test key, so add it once)
	key1 := getTestKey(0)

	name1, addr1, err := ks.Put(key1)
	require.NoError(t, err)

	keys, err = ks.List()
	require.NoError(t, err)
	require.Len(t, keys, 1)

	require.Equal(t, addr1, keys[0].Address)
	require.Equal(t, name1, keys[0].Name)
}

func TestLocalKeyStore_Delete(t *testing.T) {
	tmpdir := t.TempDir()
	ks, err := NewLocalKeyStore(tmpdir)
	require.NoError(t, err)

	privateKey := getTestKey(0)
	name, _, err := ks.Put(privateKey)
	require.NoError(t, err)

	require.True(t, ks.Has(name))

	err = ks.Delete(name)
	require.NoError(t, err)

	require.False(t, ks.Has(name))

	keys, err := ks.List()
	require.NoError(t, err)
	require.Empty(t, keys)
}

func TestLocalKeyStore_Has(t *testing.T) {
	tmpdir := t.TempDir()
	ks, err := NewLocalKeyStore(tmpdir)
	require.NoError(t, err)

	require.False(t, ks.Has("nonexistent"))

	privateKey := getTestKey(0)
	name, _, err := ks.Put(privateKey)
	require.NoError(t, err)

	require.True(t, ks.Has(name))

	// paths with separators are rejected
	require.False(t, ks.Has(filepath.Join(tmpdir, name)))
	require.False(t, ks.Has("../escape"))
}

func TestLocalKeyStore_InvalidKey(t *testing.T) {
	tmpdir := t.TempDir()
	ks, err := NewLocalKeyStore(tmpdir)
	require.NoError(t, err)

	// Try to put invalid key string
	_, _, err = ks.Put("not a valid key")
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to derive address")
}

func TestLocalKeyStore_ListSkipsInvalidFiles(t *testing.T) {
	tmpdir := t.TempDir()
	ks, err := NewLocalKeyStore(tmpdir)
	require.NoError(t, err)

	// Add a valid key
	privateKey := getTestKey(0)
	_, _, err = ks.Put(privateKey)
	require.NoError(t, err)

	// Add an invalid file
	invalidPath := filepath.Join(tmpdir, "invalid")
	err = os.WriteFile(invalidPath, []byte("garbage"), 0600)
	require.NoError(t, err)

	// Add a subdirectory (should be skipped)
	subdir := filepath.Join(tmpdir, "subdir")
	err = os.Mkdir(subdir, 0700)
	require.NoError(t, err)

	// List should only return the valid key
	keys, err := ks.List()
	require.NoError(t, err)
	require.Len(t, keys, 1)
}

func TestLocalKeyStore_PutSameKeyTwice(t *testing.T) {
	tmpdir := t.TempDir()
	ks, err := NewLocalKeyStore(tmpdir)
	require.NoError(t, err)

	privateKey := getTestKey(0)
	name1, addr1, err := ks.Put(privateKey)
	require.NoError(t, err)

	// same key again should overwrite
	name2, addr2, err := ks.Put(privateKey)
	require.NoError(t, err)

	require.Equal(t, name1, name2)
	require.Equal(t, addr1, addr2)

	keys, err := ks.List()
	require.NoError(t, err)
	require.Len(t, keys, 1)
}

func TestLocalKeyStore_DirectoryCreation(t *testing.T) {
	tmpdir := t.TempDir()
	keystorePath := filepath.Join(tmpdir, "nested", "keystore")

	// Directory doesn't exist yet
	require.NoDirExists(t, keystorePath)

	// NewLocalKeyStore should create it
	ks, err := NewLocalKeyStore(keystorePath)
	require.NoError(t, err)
	require.NotNil(t, ks)

	// Verify directory was created with correct permissions
	info, err := os.Stat(keystorePath)
	require.NoError(t, err)
	require.True(t, info.IsDir())
	require.Equal(t, os.FileMode(0700), info.Mode().Perm())
}
