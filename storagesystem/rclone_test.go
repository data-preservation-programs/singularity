package storagesystem

import (
	"bytes"
	"context"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/require"
)

func TestRCloneHandler_OverrideConfig(t *testing.T) {
	tmp := t.TempDir()

	ctx := context.Background()
	handler, err := NewRCloneHandler(ctx, model.Storage{Type: "local", Path: tmp, ClientConfig: model.ClientConfig{
		ConnectTimeout:        ptr.Of(time.Hour),
		Timeout:               ptr.Of(time.Hour),
		ExpectContinueTimeout: ptr.Of(time.Hour),
		InsecureSkipVerify:    ptr.Of(true),
		NoGzip:                ptr.Of(true),
		UserAgent:             ptr.Of("test"),
		CaCert:                []string{"a"},
		ClientCert:            ptr.Of("test"),
		ClientKey:             ptr.Of("test"),
		Headers:               map[string]string{"a": "b"},
		DisableHTTP2:          ptr.Of(true),
		DisableHTTPKeepAlives: ptr.Of(true),
	}})
	require.NoError(t, err)
	entries, err := handler.List(ctx, "")
	require.NoError(t, err)
	require.Len(t, entries, 0)
}

func TestRCloneHandler(t *testing.T) {
	tmp := t.TempDir()

	ctx := context.Background()
	_, err := NewRCloneHandler(ctx, model.Storage{Type: "xxxxx", Path: tmp})
	require.ErrorIs(t, err, ErrBackendNotSupported)

	handler, err := NewRCloneHandler(ctx, model.Storage{Type: "local", Path: tmp})
	require.NoError(t, err)

	obj, err := handler.Write(ctx, "test.txt", bytes.NewReader([]byte("test")))
	require.NoError(t, err)
	require.EqualValues(t, 4, obj.Size())

	entries, err := handler.List(ctx, "")
	require.NoError(t, err)
	require.Len(t, entries, 1)

	readCloser, _, err := handler.Read(ctx, "test.txt", 0, 4)
	require.NoError(t, err)
	defer readCloser.Close()
	read, err := io.ReadAll(readCloser)
	require.NoError(t, err)
	require.EqualValues(t, "test", read)

	readCloser2, _, err := handler.Read(ctx, "test.txt", 0, 0)
	require.NoError(t, err)
	defer readCloser2.Close()
	read, err = io.ReadAll(readCloser2)
	require.NoError(t, err)
	require.EqualValues(t, "", read)

	usage, err := handler.About(ctx)
	require.NoError(t, err)
	require.NotNil(t, usage.Used)

	entry, err := handler.Check(ctx, "test.txt")
	require.NoError(t, err)
	require.EqualValues(t, 4, entry.Size())

	entryChan := handler.Scan(ctx, "", "")
	require.NotNil(t, entryChan)
	var scannedEntries []Entry
	for entry := range entryChan {
		scannedEntries = append(scannedEntries, entry)
	}
	require.Len(t, scannedEntries, 1)

	/*
	 | sub1
	   | sub2
	     | test.txt
	   | test.txt
	 | sub2.txt
	 | sub3
	   | sub3.txt
	   | sub4
	 | test.txt
	*/

	err = os.MkdirAll(filepath.Join(tmp, "sub1", "sub2"), 0755)
	require.NoError(t, err)

	err = os.MkdirAll(filepath.Join(tmp, "sub3", "sub4"), 0755)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(tmp, "sub1", "sub2", "test.txt"), []byte("test"), 0644)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(tmp, "sub1", "test.txt"), []byte("test"), 0644)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(tmp, "sub2.txt"), []byte("test"), 0644)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(tmp, "sub3", "sub3.txt"), []byte("test"), 0644)
	require.NoError(t, err)

	tests := map[string]int{
		"":                   5,
		"sub1/sub2/test.txt": 4,
		"sub1/test.txt":      3,
		"sub2.txt":           2,
		"sub3/sub3.txt":      1,
		"test.txt":           0,
	}

	for last, expect := range tests {
		t.Run("start from "+last, func(t *testing.T) {
			entryChan = handler.Scan(ctx, "", last)
			require.NotNil(t, entryChan)
			scannedEntries = []Entry{}
			for entry := range entryChan {
				if entry.Info == nil {
					continue
				}
				scannedEntries = append(scannedEntries, entry)
			}
			require.Len(t, scannedEntries, expect)
		})
	}
}
