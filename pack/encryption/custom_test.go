package encryption

import (
	"io"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustomEncryptor_Encrypt(t *testing.T) {
	// Setup the command for the CustomEncryptor
	cmd := exec.Command("echo", "Hello, World!")
	encryptor := NewCustomEncryptor(cmd)

	// Call Encrypt with an empty reader
	rc, err := encryptor.Encrypt(strings.NewReader(""))

	// Assert that the call to Encrypt does not return an error
	require.NoError(t, err)

	// Create a buffer to store the output
	buffer := new(strings.Builder)

	// Read the output from the reader
	_, err = io.Copy(buffer, rc)
	require.NoError(t, err)

	// Assert that the output is as expected
	require.Equal(t, "Hello, World!\n", buffer.String())
}

func TestCustomEncryptor_FailedCommand(t *testing.T) {
	// Setup the command for the CustomEncryptor
	cmd := exec.Command("false")
	encryptor := NewCustomEncryptor(cmd)

	// Call Encrypt with an empty reader
	rc, err := encryptor.Encrypt(strings.NewReader(""))

	require.NoError(t, err)
	_, err = io.ReadFull(rc, make([]byte, 10))
	require.Equal(t, "EOF", err.Error())
	err = rc.Close()
	require.Equal(t, "exit status 1", err.Error())
}
