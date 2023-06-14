package encryption

import (
	"bytes"
	"filippo.io/age"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptionFullPiece(t *testing.T) {
	state, err := NewAgeEncryptor([]string{"age1th55qj77d32vhumd72de2m3y0nzsxyeahuddz770s8qadz3h6v8quedwf3"})
	assert.NoError(t, err)

	input := bytes.NewReader([]byte("This is a test"))
	reader, err := state.Encrypt(input)
	assert.NoError(t, err)
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(reader)
	assert.NoError(t, err)
	assert.Equal(t, 214, len(buf.Bytes()))

	identity, err := age.ParseX25519Identity("AGE-SECRET-KEY-1HZG3ESWDVPE3S4AM8WWCZG3H66A6RVJPXPZZEAC04FWZVT6RJ7XQAUV49J")
	assert.NoError(t, err)
	decrypted, err := age.Decrypt(buf, identity)
	assert.NoError(t, err)
	decryptedBuf := new(bytes.Buffer)
	_, err = decryptedBuf.ReadFrom(decrypted)
	assert.NoError(t, err)
	assert.Equal(t, "This is a test", decryptedBuf.String())
}
