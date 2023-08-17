package encryption

import (
	"io"

	"github.com/data-preservation-programs/singularity/model"
)

// Encryptor is an interface that defines the methods required to encrypt data in a resumable way.
type Encryptor interface {
	// Encrypt encrypts the data from the given reader and returns an io.ReadCloser
	// that can be used to read the encrypted data. The last parameter indicates whether
	// this is the last piece of data to be encrypted. This function is expected to be called only once.
	// To resume previous encryption, use LoadState before calling Encrypt. To save the encryption state
	// for later resumption, use GetState after calling Encrypt.
	Encrypt(in io.Reader) (io.ReadCloser, error)
}

func GetEncryptor(dataset model.Preparation) (Encryptor, error) {
	if len(dataset.EncryptionRecipients) > 0 {
		return NewAgeEncryptor(dataset.EncryptionRecipients)
	}

	return nil, nil
}
