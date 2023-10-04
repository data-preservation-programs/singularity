package storagesystem

import (
	"bytes"
	"context"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type faultyReader struct {
	willFail bool
}

func (f *faultyReader) Read(p []byte) (n int, err error) {
	if f.willFail {
		return 0, errors.New("test")
	}
	p[0] = 'a'
	return 1, io.EOF
}

func (f *faultyReader) Close() error {
	return nil
}

func TestScanWithConcurrency(t *testing.T) {
	tmp := t.TempDir()
	for i := 0; i < 10; i++ {
		err := os.MkdirAll(filepath.Join(tmp, strconv.Itoa(i)), 0755)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(tmp, strconv.Itoa(i), "test.txt"), []byte("test"), 0644)
		require.NoError(t, err)
		for j := 0; j < 10; j++ {
			err = os.MkdirAll(filepath.Join(tmp, strconv.Itoa(i), strconv.Itoa(j)), 0755)
			require.NoError(t, err)
			err = os.WriteFile(filepath.Join(tmp, strconv.Itoa(i), strconv.Itoa(j), "test.txt"), []byte("test"), 0644)
			require.NoError(t, err)
			for k := 0; k < 10; k++ {
				err = os.MkdirAll(filepath.Join(tmp, strconv.Itoa(i), strconv.Itoa(j), strconv.Itoa(k)), 0755)
				require.NoError(t, err)
				err = os.WriteFile(filepath.Join(tmp, strconv.Itoa(i), strconv.Itoa(j), strconv.Itoa(k), "test.txt"), []byte("test"), 0644)
				require.NoError(t, err)
			}
		}
	}

	handler, err := NewRCloneHandler(context.Background(), model.Storage{Type: "local", Path: tmp, ClientConfig: model.ClientConfig{ScanConcurrency: ptr.Of(10)}})
	require.NoError(t, err)
	ch := handler.Scan(context.Background(), "")
	var entries []Entry
	for entry := range ch {
		entries = append(entries, entry)
	}
	require.Len(t, entries, 2220)
}

func TestReaderWithRetry(t *testing.T) {
	ctx := context.Background()
	mockObject := new(MockObject)
	mockObject.On("Open", ctx, mock.Anything).Return(&faultyReader{willFail: false}, nil)
	reader := &readerWithRetry{
		ctx:                     ctx,
		object:                  mockObject,
		reader:                  &faultyReader{willFail: true},
		offset:                  0,
		retryDelay:              time.Second,
		retryBackoff:            time.Second,
		retryCountMax:           10,
		retryBackoffExponential: 1.0,
	}
	out, err := io.ReadAll(reader)
	require.NoError(t, err)
	require.EqualValues(t, "a", out)
}

func TestRCloneHandler_OverrideConfig(t *testing.T) {
	tmp := t.TempDir()

	ctx := context.Background()
	handler, err := NewRCloneHandler(ctx, model.Storage{Type: "local", Path: tmp, ClientConfig: model.ClientConfig{
		ConnectTimeout:          ptr.Of(time.Hour),
		Timeout:                 ptr.Of(time.Hour),
		ExpectContinueTimeout:   ptr.Of(time.Hour),
		InsecureSkipVerify:      ptr.Of(true),
		NoGzip:                  ptr.Of(true),
		UserAgent:               ptr.Of("test"),
		CaCert:                  []string{"a"},
		ClientCert:              ptr.Of("test"),
		ClientKey:               ptr.Of("test"),
		Headers:                 map[string]string{"a": "b"},
		DisableHTTP2:            ptr.Of(true),
		DisableHTTPKeepAlives:   ptr.Of(true),
		RetryMaxCount:           ptr.Of(10),
		RetryDelay:              ptr.Of(time.Second),
		RetryBackoff:            ptr.Of(time.Second),
		RetryBackoffExponential: ptr.Of(1.0),
		SkipInaccessibleFile:    ptr.Of(true),
		UseServerModTime:        ptr.Of(true),
		LowLevelRetries:         ptr.Of(10),
		ScanConcurrency:         ptr.Of(10),
	}})
	require.NoError(t, err)
	entries, err := handler.List(ctx, "")
	require.NoError(t, err)
	require.Len(t, entries, 0)
}

func TestRCloneHandler_EmptyS3File(t *testing.T) {
	ctx := context.Background()
	handler, err := NewRCloneHandler(ctx, model.Storage{
		Type:   "s3",
		Path:   "public-dataset-test",
		Config: map[string]string{"provider": "AWS", "region": "us-west-2", "chunk_size": "5Mi"},
	})
	require.NoError(t, err)
	stream, obj, err := handler.Read(ctx, "subfolder/empty.bin", 0, 0)
	require.NoError(t, err)
	defer stream.Close()
	require.NotNil(t, stream)
	require.NotNil(t, obj)
	require.EqualValues(t, 0, obj.Size())
	content, err := io.ReadAll(stream)
	require.NoError(t, err)
	require.Len(t, content, 0)
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

	entryChan := handler.Scan(ctx, "")
	require.NotNil(t, entryChan)
	var scannedEntries []Entry
	for entry := range entryChan {
		scannedEntries = append(scannedEntries, entry)
	}
	require.Len(t, scannedEntries, 1)
}
