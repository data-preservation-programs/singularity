package testutil

import (
	"crypto/rand"
	"os"
	"testing"

	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)
	//nolint:errcheck
	rand.Read(b)
	return b
}

func GetFileTimestamp(t *testing.T, path string) int64 {
	t.Helper()
	info, err := os.Stat(path)
	require.NoError(t, err)
	return info.ModTime().UnixNano()
}

var TestCid = cid.NewCidV1(cid.Raw, util.Hash([]byte("test")))

var TestRecipient = "age1th55qj77d32vhumd72de2m3y0nzsxyeahuddz770s8qadz3h6v8quedwf3"
var TestSecretKey = "AGE-SECRET-KEY-1HZG3ESWDVPE3S4AM8WWCZG3H66A6RVJPXPZZEAC04FWZVT6RJ7XQAUV49J"
