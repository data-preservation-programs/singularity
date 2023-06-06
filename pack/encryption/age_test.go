package encryption

import (
	"bytes"
	"crypto/rand"
	"filippo.io/age"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptionFullPiece(t *testing.T) {
	state, err := NewAgeEncryptor([]string{"age1th55qj77d32vhumd72de2m3y0nzsxyeahuddz770s8qadz3h6v8quedwf3"})
	assert.NoError(t, err)

	input := bytes.NewReader([]byte("This is a test"))
	reader, err := state.Encrypt(input, true)
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


func TestEncryptionPartialPiece(t *testing.T) {
	state, err := NewAgeEncryptor([]string{"age1th55qj77d32vhumd72de2m3y0nzsxyeahuddz770s8qadz3h6v8quedwf3"})
	assert.NoError(t, err)

	input := bytes.NewReader([]byte("This is a test"))
	reader, err := state.Encrypt(input, false)
	assert.NoError(t, err)
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(reader)
	assert.NoError(t, err)
	assert.Equal(t, 184, len(buf.Bytes()))

	s, err := state.GetState()
	assert.Equal(t, 65631, len(s))

	newState, err := NewAgeEncryptor([]string{"age1th55qj77d32vhumd72de2m3y0nzsxyeahuddz770s8qadz3h6v8quedwf3"})
	_ = newState.LoadState(s)
	input2 := bytes.NewReader([]byte("This is another test"))
	newReader, err := newState.Encrypt(input2, true)
	assert.NoError(t, err)
	_, err = buf.ReadFrom(newReader)
	assert.NoError(t, err)
	assert.Equal(t, 234, len(buf.Bytes()))

	identity, err := age.ParseX25519Identity("AGE-SECRET-KEY-1HZG3ESWDVPE3S4AM8WWCZG3H66A6RVJPXPZZEAC04FWZVT6RJ7XQAUV49J")
	assert.NoError(t, err)
	decrypted, err := age.Decrypt(buf, identity)
	assert.NoError(t, err)
	decryptedBuf := new(bytes.Buffer)
	_, err = decryptedBuf.ReadFrom(decrypted)
	assert.NoError(t, err)
	assert.Equal(t, "This is a testThis is another test", decryptedBuf.String())
}

func TestEncryptionLargePiece(t *testing.T) {
	state, err := NewAgeEncryptor([]string{"age1th55qj77d32vhumd72de2m3y0nzsxyeahuddz770s8qadz3h6v8quedwf3"})
	assert.NoError(t, err)

	randomBytes1 := make([]byte, 1000000)
	_, _ = rand.Read(randomBytes1)
	input := bytes.NewReader(randomBytes1)
	reader, err := state.Encrypt(input, false)
	assert.NoError(t, err)
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(reader)
	assert.NoError(t, err)

	s, err := state.GetState()

	newState, err := NewAgeEncryptor([]string{"age1th55qj77d32vhumd72de2m3y0nzsxyeahuddz770s8qadz3h6v8quedwf3"})
	_ = newState.LoadState(s)
	randomBytes2 := make([]byte, 1000000)
	_, _ = rand.Read(randomBytes2)
	input2 := bytes.NewReader(randomBytes2)
	newReader, err := newState.Encrypt(input2, true)
	assert.NoError(t, err)
	_, err = buf.ReadFrom(newReader)
	assert.NoError(t, err)

	identity, err := age.ParseX25519Identity("AGE-SECRET-KEY-1HZG3ESWDVPE3S4AM8WWCZG3H66A6RVJPXPZZEAC04FWZVT6RJ7XQAUV49J")
	assert.NoError(t, err)
	decrypted, err := age.Decrypt(buf, identity)
	assert.NoError(t, err)
	decryptedBuf := new(bytes.Buffer)
	_, err = decryptedBuf.ReadFrom(decrypted)
	assert.NoError(t, err)

	expected := append(randomBytes1, randomBytes2...)
	assert.Equal(t, expected, decryptedBuf.Bytes())
}

