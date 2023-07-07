package encryption

import (
	"io"

	"filippo.io/age"
	"github.com/pkg/errors"
)

// AgeEncryptor implements the Encryptor interface using age
type AgeEncryptor struct {
	rs     []age.Recipient
	target io.WriteCloser
}

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
			return
		}
		// Write final encryption block
		e.target.Close()
		// Close the pipe writer to signal the end of the stream
		writer.Close()
	}()

	return reader, nil
}
