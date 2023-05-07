package datasource

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCheckItem(t *testing.T) {
	fs := Filesystem{}

	// Create a temporary file for the test
	file, err := ioutil.TempFile("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(file.Name())
	defer file.Close()
	_, err = file.Write([]byte("Hello, world!"))
	assert.NoError(t, err)

	// Test with temporary file
	size, modTime, err := fs.CheckItem(context.Background(), file.Name())
	assert.NoError(t, err)
	assert.Equal(t, uint64(13), size)
	assert.NotNil(t, modTime)
	assert.True(t, modTime.After(time.Now().Add(-time.Minute))) // Check that modTime is recent enough
}

func TestDirSource_Scan_Last(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "testdir")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create nested directories and files
	nestedDir1 := filepath.Join(tmpDir, "nested1")
	nestedDir2 := filepath.Join(tmpDir, "nested2")
	require.NoError(t, os.Mkdir(nestedDir1, 0755))
	require.NoError(t, os.Mkdir(nestedDir2, 0755))

	fileNames := []string{"file1.txt", "file2.txt", "file3.txt"}
	for _, fileName := range fileNames {
		filePath1 := filepath.Join(nestedDir1, fileName)
		err = ioutil.WriteFile(filePath1, []byte("content"), 0644)
		require.NoError(t, err)
		filePath2 := filepath.Join(nestedDir2, fileName)
		err = ioutil.WriteFile(filePath2, []byte("content"), 0644)
		require.NoError(t, err)
	}

	dirSource := Filesystem{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	lastFile := filepath.Join(nestedDir2, fileNames[1])
	entries := dirSource.Scan(ctx, tmpDir, lastFile)

	foundFiles := make(map[string]bool)
	for entry := range entries {
		if entry.Error != nil {
			t.Errorf("unexpected error: %v", entry.Error)
			continue
		}

		foundFiles[entry.Path] = true
	}

	// Check that only the files after the "last" file are found
	assert.False(t, foundFiles[filepath.Join(nestedDir1, fileNames[0])], "file should not be found: %s", fileNames[0])
	assert.False(t, foundFiles[filepath.Join(nestedDir1, fileNames[1])], "file should not be found: %s", fileNames[1])
	assert.False(t, foundFiles[filepath.Join(nestedDir1, fileNames[2])], "file should not be found: %s", fileNames[2])
	assert.False(t, foundFiles[filepath.Join(nestedDir2, fileNames[0])], "file should not be found: %s", fileNames[0])
	assert.False(t, foundFiles[filepath.Join(nestedDir2, fileNames[1])], "file should not be found: %s", fileNames[1])
	assert.True(t, foundFiles[filepath.Join(nestedDir2, fileNames[2])], "file not found: %s", fileNames[2])
}

func TestDirSource_Scan_FileAccessError(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "testdir")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	fileName := "inaccessible.txt"
	filePath := filepath.Join(tmpDir, fileName)
	err = ioutil.WriteFile(filePath, []byte("content"), 0644)
	require.NoError(t, err)

	// Change the parent directory permission to make it inaccessible
	require.NoError(t, os.Chmod(tmpDir, 0000))
	defer os.Chmod(tmpDir, 0755) // Restore the permission after the test

	dirSource := Filesystem{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	entries := dirSource.Scan(ctx, tmpDir, "")

	foundFileAccessError := false
	for entry := range entries {
		if entry.Error != nil {
			// Expect a permission error
			if os.IsPermission(entry.Error) {
				foundFileAccessError = true
			} else {
				t.Errorf("unexpected error: %v", entry.Error)
			}
			continue
		}
	}

	assert.True(t, foundFileAccessError, "file access error not found")
}

func TestDirSource_Scan_BrokenSymlink(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "testdir")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create a broken symlink
	symlinkName := "brokensymlink"
	symlinkTarget := filepath.Join(tmpDir, "nonexistent")
	symlinkPath := filepath.Join(tmpDir, symlinkName)
	err = os.Symlink(symlinkTarget, symlinkPath)
	require.NoError(t, err)

	dirSource := Filesystem{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	entries := dirSource.Scan(ctx, tmpDir, "")

	foundBrokenSymlink := false
	for entry := range entries {
		if entry.Error != nil {
			// Expect a broken symlink error
			if os.IsNotExist(entry.Error) {
				foundBrokenSymlink = true
			} else {
				t.Errorf("unexpected error: %v", entry.Error)
			}
			continue
		}
	}

	assert.True(t, foundBrokenSymlink, "broken symlink not found")
}

func TestDirSource_Scan_NestedDirectories(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "testdir")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create nested directories and files
	nestedDir := filepath.Join(tmpDir, "nested")
	require.NoError(t, os.Mkdir(nestedDir, 0755))

	fileNames := []string{"file1.txt", "file2.txt"}
	for _, fileName := range fileNames {
		filePath := filepath.Join(nestedDir, fileName)
		err = ioutil.WriteFile(filePath, []byte("content"), 0644)
		require.NoError(t, err)
	}

	dirSource := Filesystem{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	entries := dirSource.Scan(ctx, tmpDir, "")

	foundFiles := make(map[string]bool)
	for entry := range entries {
		if entry.Error != nil {
			t.Errorf("unexpected error: %v", entry.Error)
			continue
		}

		foundFiles[entry.Path] = true
	}

	for _, fileName := range fileNames {
		assert.True(t, foundFiles[filepath.Join(nestedDir, fileName)], "file not found: %s", fileName)
	}
}
func TestFileSource_Open(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := ioutil.TempFile("", "testfile")
	assert.NoError(t, err)

	// Write some content to the temporary file
	content := "The quick brown fox jumps over the lazy dog"
	_, err = tmpfile.WriteString(content)
	assert.NoError(t, err)

	tmpfile.Close()

	// Defer removal of the temporary file
	defer os.Remove(tmpfile.Name())

	// Test cases
	tests := []struct {
		name     string
		offset   uint64
		length   uint64
		expected string
	}{
		{
			name:     "Read from the beginning",
			offset:   0,
			length:   9,
			expected: "The quick",
		},
		{
			name:     "Read from the middle",
			offset:   10,
			length:   5,
			expected: "brown",
		},
		{
			name:     "Read till the end",
			offset:   35,
			length:   10,
			expected: "lazy dog",
		},
	}

	fileSource := Filesystem{}
	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readCloser, err := fileSource.Read(ctx, tmpfile.Name(), tt.offset, tt.length)
			assert.NoError(t, err)
			defer readCloser.Close()

			buf := make([]byte, tt.length)
			n, err := readCloser.Read(buf)
			assert.NoError(t, err)

			result := string(buf[:n])
			assert.Equal(t, tt.expected, result)
		})
	}
}
