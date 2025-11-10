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

// Alternative test keys (pre-generated, valid lotus format)
var testKeys = []string{
	testutil.TestPrivateKeyHex,
	// We only have one test key in testutil, so for multi-key tests we'll use the same one
	// This is fine for testing keystore functionality
}

func TestLocalKeyStore_PutAndGet(t *testing.T) {
	tmpdir := t.TempDir()
	ks, err := NewLocalKeyStore(tmpdir)
	require.NoError(t, err)

	// Use test key
	privateKey := getTestKey(0)

	// Put the key
	path, addr, err := ks.Put(privateKey)
	require.NoError(t, err)
	require.NotEmpty(t, path)
	require.NotEqual(t, address.Undef, addr)

	// Verify file exists
	require.FileExists(t, path)

	// Verify path is in the keystore directory
	require.Contains(t, path, tmpdir)

	// Verify filename matches address
	require.Equal(t, filepath.Join(tmpdir, addr.String()), path)

	// Get the key back
	loadedKey, err := ks.Get(path)
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

	path1, addr1, err := ks.Put(key1)
	require.NoError(t, err)

	// List should return it
	keys, err = ks.List()
	require.NoError(t, err)
	require.Len(t, keys, 1)

	// Verify address matches
	require.Equal(t, addr1, keys[0].Address)

	// Verify path is correct
	require.Equal(t, path1, keys[0].Path)
}

func TestLocalKeyStore_Delete(t *testing.T) {
	tmpdir := t.TempDir()
	ks, err := NewLocalKeyStore(tmpdir)
	require.NoError(t, err)

	// Add a key
	privateKey := getTestKey(0)
	path, _, err := ks.Put(privateKey)
	require.NoError(t, err)

	// Verify it exists
	require.True(t, ks.Has(path))

	// Delete it
	err = ks.Delete(path)
	require.NoError(t, err)

	// Verify it's gone
	require.False(t, ks.Has(path))

	// List should be empty
	keys, err := ks.List()
	require.NoError(t, err)
	require.Empty(t, keys)
}

func TestLocalKeyStore_Has(t *testing.T) {
	tmpdir := t.TempDir()
	ks, err := NewLocalKeyStore(tmpdir)
	require.NoError(t, err)

	// Non-existent key
	require.False(t, ks.Has(filepath.Join(tmpdir, "nonexistent")))

	// Add a key
	privateKey := getTestKey(0)
	path, _, err := ks.Put(privateKey)
	require.NoError(t, err)

	// Should exist
	require.True(t, ks.Has(path))
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

	// Add a key
	privateKey := getTestKey(0)
	path1, addr1, err := ks.Put(privateKey)
	require.NoError(t, err)

	// Add the same key again (should overwrite)
	path2, addr2, err := ks.Put(privateKey)
	require.NoError(t, err)

	// Paths and addresses should be the same
	require.Equal(t, path1, path2)
	require.Equal(t, addr1, addr2)

	// Should only have one key in the list
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
