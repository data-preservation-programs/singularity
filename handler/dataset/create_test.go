package dataset

import (
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestCreateHandler_NoDatasetName(t *testing.T) {
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{})
	require.ErrorContains(t, err, "name is required")
}

func TestCreateHandler_MaxSizeNotValid(t *testing.T) {
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "not valid"})
	require.ErrorContains(t, err, "invalid value for max-size")
}

func TestCreateHandler_PieceSizeNotValid(t *testing.T) {
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "2GB", PieceSizeStr: "not valid"})
	require.ErrorContains(t, err, "invalid value for piece-size")
}

func TestCreateHandler_PieceSizeNotPowerOfTwo(t *testing.T) {
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "2GB", PieceSizeStr: "3GB"})
	require.ErrorContains(t, err, "piece size must be a power of two")
}

func TestCreateHandler_PieceSizeTooLarge(t *testing.T) {
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "2GB", PieceSizeStr: "128GiB"})
	require.ErrorContains(t, err, "piece size cannot be larger than 64 GiB")
}

func TestCreateHandler_MaxSizeTooLarge(t *testing.T) {
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "63.9GiB"})
	require.ErrorContains(t, err, "max size needs to be reduced to leave space for padding")
}

func TestCreateHandler_OutDirDoesNotExist(t *testing.T) {
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "2GB", OutputDirs: []string{"not exist"}})
	require.ErrorContains(t, err, "output directory does not exist")
}

func TestCreateHandler_RecipientsScriptCannotBeUsedTogether(t *testing.T) {
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "2GB", EncryptionRecipients: []string{"test"}, EncryptionScript: "test"})
	require.ErrorContains(t, err, "encryption recipients and script cannot be used together")
}

func TestCreateHandler_EncryptionNeedsOutputDir(t *testing.T) {
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "2GB", EncryptionRecipients: []string{"test"}})
	require.ErrorContains(t, err, "encryption is not compatible with inline preparation")
}

func TestCreateHandler_Success(t *testing.T) {
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MaxSizeStr: "2GB"})
	require.NoError(t, err)
	dataset := model.Dataset{}
	db.Where("name = ?", "test").First(&dataset)
	require.Equal(t, "test", dataset.Name)
}
