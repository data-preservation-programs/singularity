package testutil

import (
	"crypto/rand"
	"os"
	"strings"
	"testing"

	rand2 "math/rand"

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

func RandomLetterString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"

	b := make([]byte, length)
	for i := range b {
		//nolint:gosec
		b[i] = charset[rand2.Intn(len(charset))]
	}
	return string(b)
}

func GetFileTimestamp(t *testing.T, path string) int64 {
	t.Helper()
	info, err := os.Stat(path)
	require.NoError(t, err)
	return info.ModTime().UnixNano()
}

var TestCid = cid.NewCidV1(cid.Raw, util.Hash([]byte("test")))

// https://github.com/jsign/go-filsigner/blob/main/wallet/wallet_test.go

const TestWalletAddr = "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa"

const TestPrivateKeyHex = "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a226b35507976337148327349586343595a58594f5775453149326e32554539436861556b6c4e36695a5763453d227d"

func EscapePath(p string) string {
	return "'" + strings.ReplaceAll(p, `\`, `\\`) + "'"
}
