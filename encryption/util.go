package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/pkg/errors"
	"golang.org/x/crypto/pbkdf2"
	"gorm.io/gorm"
	"io"
	"sync"
)

var encryptionKey []byte

var lock = sync.Mutex{}

func getSalt(db *gorm.DB) ([]byte, error) {
	row := model.Global{}
	err := db.Where("key = ?", "salt").First(&row).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to get salt")
	}
	decoded, err := base64.StdEncoding.DecodeString(row.Value)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode salt")
	}
	return decoded, nil
}

func Init(keyStr string, db *gorm.DB) error {
	lock.Lock()
	defer lock.Unlock()
	if encryptionKey != nil {
		salt, err := getSalt(db)
		if err != nil {
			return errors.Wrap(err, "failed to get salt")
		}

		encryptionKey = pbkdf2.Key([]byte(keyStr), salt, 4096, 32, sha256.New)
	}
	return nil
}

func EncryptToBytes(payload []byte) ([]byte, error) {
	if encryptionKey == nil {
		return nil, errors.New("encryption key not initialized")
	}
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, 12)
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, payload, nil)

	// nolint:makezero
	return append(nonce, ciphertext...), nil
}

func EncryptToBase64String(payload []byte) (string, error) {
	encrypted, err := EncryptToBytes(payload)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func DecryptFromBytes(payload []byte) ([]byte, error) {
	if encryptionKey == nil {
		return nil, errors.New("encryption key not initialized")
	}
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := payload[:12]
	ciphertext := payload[12:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func DecryptFromBase64String(payload string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}
	return DecryptFromBytes(decoded)
}
