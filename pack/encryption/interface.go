package encryption

import (
	"github.com/data-preservation-programs/singularity/model"
	"github.com/google/shlex"
	"github.com/pkg/errors"
	"io"
	"os/exec"
)

// Encryptor is an interface that defines the methods required to encrypt data in a resumable way.
type Encryptor interface {
	// Encrypt encrypts the data from the given reader and returns an io.ReadCloser
	// that can be used to read the encrypted data. The last parameter indicates whether
	// this is the last piece of data to be encrypted. This function is expected to be called only once.
	// To resume previous encryption, use LoadState before calling Encrypt. To save the encryption state
	// for later resumption, use GetState after calling Encrypt.
	Encrypt(in io.Reader, last bool) (io.ReadCloser, error)

	// GetState returns the current state of the Encryptor as a byte slice.
	// This state can be used to resume encryption at a later time.
	GetState() ([]byte, error)

	// LoadState loads the state of the Encryptor from a byte slice.
	LoadState([]byte) error
}

func GetEncryptor(dataset model.Dataset) (Encryptor, error){
	if len(dataset.EncryptionRecipients) > 0 {
		return NewAgeEncryptor(dataset.EncryptionRecipients)
	}

	if dataset.EncryptionScript != "" {
		parts, err := shlex.Split(dataset.EncryptionScript)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse encryption script")
		}
		cmd := exec.Command(parts[0], parts[1:]...)
		return NewCustomEncryptor(cmd), nil
	}

	return nil, nil
}
