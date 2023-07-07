package encryption

import (
	"bytes"
	"testing"

	"filippo.io/age"
	"github.com/stretchr/testify/require"
)

func TestEncryptionFullPiece(t *testing.T) {
	state, err := NewAgeEncryptor([]string{"age1th55qj77d32vhumd72de2m3y0nzsxyeahuddz770s8qadz3h6v8quedwf3"})
	require.NoError(t, err)

	input := bytes.NewReader([]byte("This is a test"))
	reader, err := state.Encrypt(input)
	require.NoError(t, err)
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(reader)
	require.NoError(t, err)
	require.Equal(t, 214, len(buf.Bytes()))

	identity, err := age.ParseX25519Identity("AGE-SECRET-KEY-1HZG3ESWDVPE3S4AM8WWCZG3H66A6RVJPXPZZEAC04FWZVT6RJ7XQAUV49J")
	require.NoError(t, err)
	decrypted, err := age.Decrypt(buf, identity)
	require.NoError(t, err)
	decryptedBuf := new(bytes.Buffer)
	_, err = decryptedBuf.ReadFrom(decrypted)
	require.NoError(t, err)
	require.Equal(t, "This is a test", decryptedBuf.String())
}
