package file

import (
	"bytes"
	"io"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type testReadCloser struct {
	*bytes.Reader
	closed bool
}

func (rc *testReadCloser) Close() error {
	rc.closed = true
	return nil
}

func TestRangeReader(t *testing.T) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	testData := make([]byte, 257)
	for i := range testData {
		testData[i] = charset[seededRand.Intn(len(charset))]
	}

	rc := &testReadCloser{
		Reader: bytes.NewReader(testData),
	}
	rr := rangeReader{
		reader:    rc,
		remaining: rc.Size(),
	}

	outBuf := new(bytes.Buffer)

	var totalRead int64
	const readLen = 23
	for {
		n, err := rr.writeToN(outBuf, readLen)
		totalRead += n
		require.Equal(t, rc.Size()-totalRead, rr.remaining)
		if n < readLen {
			break
		}
		require.NoError(t, err)
	}
	require.Zero(t, rr.remaining)
	require.Equal(t, rc.Size(), totalRead)
	require.Equal(t, testData, outBuf.Bytes())

	n, err := rr.writeToN(outBuf, readLen)
	require.ErrorIs(t, err, io.EOF)
	require.Zero(t, n)

	require.NoError(t, rr.close())
	require.True(t, rc.closed)
	require.Nil(t, rr.reader)
}
