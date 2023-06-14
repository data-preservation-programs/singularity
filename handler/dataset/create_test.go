package dataset

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateHandler_NoDatasetName(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{})
	assert.ErrorContains(err, "name is required")
}

func TestCreateHandler_MaxSizeNotValid(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "not valid"})
	assert.ErrorContains(err, "invalid value for max-size")
}

func TestCreateHandler_PieceSizeNotValid(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "2GB", PieceSizeStr: "not valid"})
	assert.ErrorContains(err, "invalid value for piece-size")
}

func TestCreateHandler_PieceSizeNotPowerOfTwo(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "2GB", PieceSizeStr: "3GB"})
	assert.ErrorContains(err, "piece size must be a power of two")
}

func TestCreateHandler_PieceSizeTooLarge(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "2GB", PieceSizeStr: "128GiB"})
	assert.ErrorContains(err, "piece size cannot be larger than 64 GiB")
}

func TestCreateHandler_MaxSizeTooLarge(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "63.9GiB"})
	assert.ErrorContains(err, "max size needs to be reduced to leave space for padding")
}

func TestCreateHandler_OutDirDoesNotExist(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "2GB", OutputDirs: []string{"not exist"}})
	assert.ErrorContains(err, "output directory does not exist")
}

func TestCreateHandler_RecipientsScriptCannotBeUsedTogether(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "2GB", EncryptionRecipients: []string{"test"}, EncryptionScript: "test"})
	assert.ErrorContains(err, "encryption recipients and script cannot be used together")
}

func TestCreateHandler_EncryptionNeedsOutputDir(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "2GB", EncryptionRecipients: []string{"test"}})
	assert.ErrorContains(err, "encryption is not compatible with inline preparation")
}

func TestCreateHandler_Success(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "2GB"})
	assert.Nil(err)
	dataset := model.Dataset{}
	db.Where("name = ?", "test").First(&dataset)
	assert.Equal("test", dataset.Name)
}
