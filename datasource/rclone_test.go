package datasource

import (
	"context"
	"io"
	"io/fs"
	"os"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestRCloneHandler(t *testing.T) {
	temp := t.TempDir()
	err := os.WriteFile(temp+"/test.txt", []byte("test"), 0644)
	require.NoError(t, err)
	ctx := context.Background()
	h, err := NewRCloneHandler(ctx, model.Source{
		Type: "local",
		Path: temp,
	})
	require.NoError(t, err)

	entry, err := h.Check(ctx, "test.txt")
	require.NoError(t, err)
	require.Equal(t, "test.txt", entry.Remote())

	readCloser, obj, err := h.Read(ctx, "test.txt", 1, 2)
	require.NoError(t, err)
	require.Equal(t, "test.txt", obj.Remote())
	defer readCloser.Close()
	content, err := io.ReadAll(readCloser)
	require.NoError(t, err)
	require.Equal(t, "es", string(content))

	ch := h.Scan(ctx, "", "")
	var entries []Entry
	for entry := range ch {
		entries = append(entries, entry)
	}
	require.Len(t, entries, 1)
	require.NoError(t, entries[0].Error)
	require.Equal(t, "test.txt", entries[0].Info.Remote())

	list, err := h.List(ctx, "")
	require.NoError(t, err)
	require.Len(t, list, 1)
	require.Equal(t, "test.txt", list[0].Remote())

	stat, err := h.Stat(ctx, "")
	require.NoError(t, err)
	require.True(t, stat.IsDir())
	require.EqualValues(t, fs.ModeDir, stat.Mode())
	require.Nil(t, stat.Sys())

	stat, err = h.Stat(ctx, "test.txt")
	require.NoError(t, err)
	require.False(t, stat.IsDir())
	require.Equal(t, "test.txt", stat.Name())
	require.Equal(t, int64(4), stat.Size())
	require.Greater(t, stat.ModTime().Unix(), int64(0))
	require.EqualValues(t, 0, stat.Mode())

	stat, err = h.Stat(ctx, "nonexisting.txt")
	require.ErrorIs(t, err, os.ErrNotExist)

	err = h.Mkdir(ctx, "test", 0755)
	require.NoError(t, err)

	err = h.Rename(ctx, "test.txt", "test/test2.txt")
	require.NoError(t, err)

	err = h.Rename(ctx, "test", "test2")
	require.NoError(t, err)

	stat, err = h.Stat(ctx, "test2/test2.txt")
	require.NoError(t, err)

	ch = h.Scan(ctx, "", "")
	entries = []Entry{}
	for entry := range ch {
		entries = append(entries, entry)
	}
	require.Len(t, entries, 1)

	defer func() {
		err = h.RemoveAll(ctx, "test2")
		require.NoError(t, err)
	}()

	read, err := h.OpenFile(ctx, "test2/test2.txt", os.O_RDONLY, 0644)
	require.NoError(t, err)
	content, err = io.ReadAll(read)
	require.NoError(t, err)
	require.Equal(t, "test", string(content))

	_, err = read.Seek(1, io.SeekStart)
	require.NoError(t, err)
	content, err = io.ReadAll(read)
	require.NoError(t, err)
	require.Equal(t, "est", string(content))

	_, err = read.Seek(-2, io.SeekEnd)
	require.NoError(t, err)
	content, err = io.ReadAll(read)
	require.NoError(t, err)
	require.Equal(t, "st", string(content))

	_, err = read.Seek(-1, io.SeekCurrent)
	require.NoError(t, err)
	content, err = io.ReadAll(read)
	require.NoError(t, err)
	require.Equal(t, "t", string(content))

	require.NoError(t, read.Close())

	write, err := h.OpenFile(ctx, "test2/test2.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	require.NoError(t, err)
	_, err = write.Write([]byte("hello"))
	require.NoError(t, err)
	_, err = write.Write([]byte("world"))
	require.NoError(t, err)
	require.NoError(t, write.Close())

	content, err = os.ReadFile(temp + "/test2/test2.txt")
	require.NoError(t, err)
	require.Equal(t, "helloworld", string(content))

	write, err = h.OpenFile(ctx, "test2/test3.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	require.NoError(t, err)
	_, err = write.Write([]byte("hello2"))
	require.NoError(t, err)
	_, err = write.Write([]byte("world2"))
	require.NoError(t, err)
	require.NoError(t, write.Close())

	content, err = os.ReadFile(temp + "/test2/test3.txt")
	require.NoError(t, err)
	require.Equal(t, "hello2world2", string(content))
}
