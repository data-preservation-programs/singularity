package encryption

import (
	"io"

	"filippo.io/age"
	"github.com/cockroachdb/errors"
)

// AgeEncryptor implements the Encryptor interface using age
type AgeEncryptor struct {
	rs     []age.Recipient
	target io.WriteCloser
}

// NewAgeEncryptor creates a new Encryptor instance that uses the age encryption tool
// to encrypt data for a specified list of recipients. Each recipient is identified
// by a public key string that is parsed and used to encrypt the data.
//
// Parameters:
// - recipients: A list of strings, where each string is a recipient's public key.
//
// Returns:
//   - An Encryptor instance that is configured to encrypt data for the specified recipients.
//     The Encryptor instance is assumed to conform to an interface that provides
//     methods for encrypting data.
//   - An error if any recipient's public key string fails to parse or if any other error occurs
//     while creating the Encryptor.
func NewAgeEncryptor(recipients []string) (Encryptor, error) {
	var rs []age.Recipient
	for _, recipient := range recipients {
		r, err := age.ParseX25519Recipient(recipient)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse recipient")
		}
		rs = append(rs, r)
	}
	return &AgeEncryptor{rs: rs}, nil
}

// Encrypt takes an io.Reader as input, which is the source of plaintext data that needs to be encrypted.
// It encrypts this data using the AGE encryption protocol, streaming the encrypted data to the caller through
// an io.ReadCloser.
//
// Parameters:
// - in: An io.Reader from which the plaintext data will be read. This could be a file, a buffer, a network connection, etc.
//
// Returns:
// - An io.ReadCloser through which the caller can read the encrypted data.
// - An error if any error occurs while setting up the encryption.
//
// Note:
//   - This function runs concurrently; it spawns a goroutine to handle the encryption process. The caller can
//     immediately start reading from the returned io.ReadCloser, and the data will be encrypted on-the-fly as it is read.
func (e *AgeEncryptor) Encrypt(in io.Reader) (io.ReadCloser, error) {
	var err error
	reader, writer := io.Pipe()
	go func() {
		e.target, err = age.Encrypt(writer, e.rs...)
		if err != nil {
			writer.CloseWithError(err)
			return
		}
		_, err = io.Copy(e.target, in)
		if err != nil {
			writer.CloseWithError(err)
			e.target.Close()
			return
		}
		// Write final encryption block
		e.target.Close()
		// Close the pipe writer to signal the end of the stream
		writer.Close()
	}()

	return reader, nil
}
