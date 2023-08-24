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

const TestRecipient = "age1th55qj77d32vhumd72de2m3y0nzsxyeahuddz770s8qadz3h6v8quedwf3"
const TestSecretKey = "AGE-SECRET-KEY-1HZG3ESWDVPE3S4AM8WWCZG3H66A6RVJPXPZZEAC04FWZVT6RJ7XQAUV49J"

// https://github.com/jsign/go-filsigner/blob/main/wallet/wallet_test.go

const TestWalletAddr = "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa"

const TestPrivateKeyHex = "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a226b35507976337148327349586343595a58594f5775453149326e32554539436861556b6c4e36695a5763453d227d"
