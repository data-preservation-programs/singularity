package model

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptWithoutInit(t *testing.T) {
	assert := assert.New(t)
	encryptionKey = nil
	encrypted, err := EncryptToBytes([]byte("test"))
	assert.ErrorContains(err, "encryption key not initialized")
	assert.Nil(encrypted)

	encryptedStr, err := EncryptToBase64String([]byte("test"))
	assert.ErrorContains(err, "encryption key not initialized")
	assert.Empty(encryptedStr)
}

func TestInit(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer DropAll(db)
	err := InitializeEncryption("password", db)
	assert.NoError(err)

	encrypted, err := EncryptToBytes([]byte("test"))
	assert.NoError(err)
	assert.NotEqual("test", string(encrypted))

	decrypted, err := DecryptFromBytes(encrypted)
	assert.NoError(err)
	assert.Equal("test", string(decrypted))

	encryptedStr, err := EncryptToBase64String([]byte("test2"))
	assert.NoError(err)
	assert.NotEqual("test2", encryptedStr)

	decryptedStr, err := DecryptFromBase64String(encryptedStr)
	assert.NoError(err)
	assert.Equal("test2", string(decryptedStr))
}
