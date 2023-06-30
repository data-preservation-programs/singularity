package encryption

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os/exec"
	"strings"
	"testing"
)

func TestCustomEncryptor_Encrypt(t *testing.T) {
	// Setup the command for the CustomEncryptor
	cmd := exec.Command("echo", "Hello, World!")
	encryptor := NewCustomEncryptor(cmd)

	// Call Encrypt with an empty reader
	rc, err := encryptor.Encrypt(strings.NewReader(""))

	// Assert that the call to Encrypt does not return an error
	assert.NoError(t, err)

	// Create a buffer to store the output
	buffer := new(strings.Builder)

	// Read the output from the reader
	_, err = io.Copy(buffer, rc)
	assert.NoError(t, err)

	// Assert that the output is as expected
	assert.Equal(t, "Hello, World!\n", buffer.String())
}

func TestCustomEncryptor_FailedCommand(t *testing.T) {
	// Setup the command for the CustomEncryptor
	cmd := exec.Command("false")
	encryptor := NewCustomEncryptor(cmd)

	// Call Encrypt with an empty reader
	rc, err := encryptor.Encrypt(strings.NewReader(""))

	assert.NoError(t, err)
	_, err = io.ReadFull(rc, make([]byte, 10))
	assert.Equal(t, "EOF", err.Error())
	err = rc.Close()
	assert.Equal(t, "exit status 1", err.Error())
}
