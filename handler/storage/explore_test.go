package storage

import (
	"context"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestExploreHandler_NotFound(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	_, err = ExploreHandler(context.Background(), db, "not found", "")
	require.ErrorIs(t, err, handlererror.ErrNotFound)
}

func TestExploreHandler_InvalidPath(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	tmp := t.TempDir()
	err = db.Create(&model.Storage{
		Name: "test",
		Type: "local",
		Path: tmp,
	}).Error
	require.NoError(t, err)

	_, err = ExploreHandler(context.Background(), db, "test", "invalid")
	require.ErrorContains(t, err, "not found")
}

func TestExploreHandler(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	tmp := t.TempDir()
	err = os.WriteFile(filepath.Join(tmp, "test.txt"), []byte("test"), 0644)
	require.NoError(t, err)
	err = os.MkdirAll(filepath.Join(tmp, "test"), 0755)
	require.NoError(t, err)
	err = db.Create(&model.Storage{
		Name: "test",
		Type: "local",
		Path: tmp,
	}).Error
	require.NoError(t, err)

	entries, err := ExploreHandler(context.Background(), db, "test", "")
	require.NoError(t, err)
	require.Len(t, entries, 2)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Path < entries[j].Path
	})
	require.Equal(t, "test.txt", entries[1].Path)
	require.EqualValues(t, 4, entries[1].Size)
	require.False(t, entries[1].IsDir)
	require.Equal(t, "test", entries[0].Path)
	require.EqualValues(t, -1, entries[0].Size)
	require.True(t, entries[0].IsDir)
}
