package encryption

import (
	"bytes"
	"testing"

	"filippo.io/age"
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
)

func TestEncryptionFullPiece(t *testing.T) {
	state, err := NewAgeEncryptor([]string{testutil.TestRecipient})
	require.NoError(t, err)

	input := bytes.NewReader([]byte("This is a test"))
	reader, err := state.Encrypt(input)
	require.NoError(t, err)
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(reader)
	require.NoError(t, err)
	require.Equal(t, 214, len(buf.Bytes()))

	identity, err := age.ParseX25519Identity(testutil.TestSecretKey)
	require.NoError(t, err)
	decrypted, err := age.Decrypt(buf, identity)
	require.NoError(t, err)
	decryptedBuf := new(bytes.Buffer)
	_, err = decryptedBuf.ReadFrom(decrypted)
	require.NoError(t, err)
	require.Equal(t, "This is a test", decryptedBuf.String())
}

func TestEncryptionWithReadError(t *testing.T) {
	state, err := NewAgeEncryptor([]string{testutil.TestRecipient})
	require.NoError(t, err)

	readErr := errors.New("read error")
	input := NewTestReader("This is a test", readErr)
	reader, err := state.Encrypt(input)
	require.NoError(t, err)
	buf := new(bytes.Buffer)
	n, err := buf.ReadFrom(reader)
	require.ErrorIs(t, err, readErr)
	require.EqualValues(t, 184, n)
}

func TestEncryptionWithAgeError(t *testing.T) {
	state, err := NewAgeEncryptor([]string{})
	require.NoError(t, err)

	input := bytes.NewReader([]byte("This is a test"))
	reader, err := state.Encrypt(input)
	require.NoError(t, err)
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(reader)
	require.ErrorContains(t, err, "no recipients")
}

type TestReader struct {
	data       []byte
	readOffset int
	err        error
}

func NewTestReader(data string, err error) *TestReader {
	return &TestReader{
		data:       []byte(data),
		readOffset: 0,
		err:        err,
	}
}

func (r *TestReader) Read(p []byte) (n int, err error) {
	if r.readOffset >= len(r.data) {
		return 0, r.err // Return the read error when all data has been read
	}

	n = copy(p, r.data[r.readOffset:])
	r.readOffset += n
	return n, nil
}
