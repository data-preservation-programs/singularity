package dataprep

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/stretchr/testify/require"
)

func TestCreatePreparationHandler_MaxSizeNotValid(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	_, err = CreatePreparationHandler(context.Background(), db, CreateRequest{MaxSizeStr: "not valid"})
	require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	require.ErrorContains(t, err, "invalid value for maxSize")
}

func TestCreatePreparationHandler_PieceSizeNotValid(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	_, err = CreatePreparationHandler(context.Background(), db, CreateRequest{MaxSizeStr: "2GB", PieceSizeStr: "not valid"})
	require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	require.ErrorContains(t, err, "invalid value for pieceSize")
}

func TestCreatePreparationHandler_PieceSizeNotPowerOfTwo(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	_, err = CreatePreparationHandler(context.Background(), db, CreateRequest{MaxSizeStr: "2GB", PieceSizeStr: "3GB"})
	require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	require.ErrorContains(t, err, "pieceSize must be a power of two")
}

func TestCreatePreparationHandler_PieceSizeTooLarge(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	_, err = CreatePreparationHandler(context.Background(), db, CreateRequest{MaxSizeStr: "2GB", PieceSizeStr: "128GiB"})
	require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	require.ErrorContains(t, err, "pieceSize cannot be larger than 64 GiB")
}

func TestCreatePreparationHandler_MaxSizeTooLarge(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	_, err = CreatePreparationHandler(context.Background(), db, CreateRequest{MaxSizeStr: "63.9GiB"})
	require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	require.ErrorContains(t, err, "maxSize needs to be reduced to leave space for padding")
}

func TestCreatePreparationHandler_EncryptionNeedsOutputDir(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	_, err = CreatePreparationHandler(context.Background(), db, CreateRequest{MaxSizeStr: "2GB", EncryptionRecipients: []string{"test"}})
	require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	require.ErrorContains(t, err, "encryption is not compatible with inline preparation")
}

func TestCreatePreparationHandler_NoSource(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	_, err = CreatePreparationHandler(context.Background(), db, CreateRequest{MaxSizeStr: "2GB"})
	require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	require.ErrorContains(t, err, "at least one source storage must be specified")
}

func TestCreatePreparationHandler_Success(t *testing.T) {
	tmp1 := t.TempDir()
	tmp2 := t.TempDir()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	_, err = storage.CreateStorageHandler(context.Background(), db, "local", "", "source", tmp1, nil)
	require.NoError(t, err)
	_, err = storage.CreateStorageHandler(context.Background(), db, "local", "", "output", tmp2, nil)
	require.NoError(t, err)
	preparation, err := CreatePreparationHandler(context.Background(), db, CreateRequest{
		MaxSizeStr:     "2GB",
		SourceStorages: []string{"source"},
		OutputStorages: []string{"output"},
	})
	require.NoError(t, err)
	require.Greater(t, preparation.ID, uint32(0))
}
