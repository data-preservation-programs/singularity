package dataprep

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestCreatePreparationHandler_MaxSizeNotValid(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, err := Default.CreatePreparationHandler(ctx, db, CreateRequest{Name: "name", MaxSizeStr: "not valid"})
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid value for maxSize")
	})
}

func TestCreatePreparationHandler_PieceSizeNotValid(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, err := Default.CreatePreparationHandler(ctx, db, CreateRequest{Name: "name", MaxSizeStr: "2GB", PieceSizeStr: "not valid"})
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid value for pieceSize")
	})
}

func TestCreatePreparationHandler_PieceSizeNotPowerOfTwo(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, err := Default.CreatePreparationHandler(ctx, db, CreateRequest{Name: "name", MaxSizeStr: "2GB", PieceSizeStr: "3GB"})
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "pieceSize must be a power of two")
	})
}

func TestCreatePreparationHandler_PieceSizeTooLarge(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, err := Default.CreatePreparationHandler(ctx, db, CreateRequest{Name: "name", MaxSizeStr: "2GB", PieceSizeStr: "128GiB"})
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "pieceSize cannot be larger than 64 GiB")
	})
}

func TestCreatePreparationHandler_MaxSizeTooLarge(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, err := Default.CreatePreparationHandler(ctx, db, CreateRequest{Name: "name", MaxSizeStr: "63.9GiB"})
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "maxSize needs to be reduced to leave space for padding")
	})
}

func TestCreatePreparationHandler_DeleteAfterExportWithoutOutput(t *testing.T) {
	tmp1 := t.TempDir()
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, err := storage.Default.CreateStorageHandler(ctx, db, "local", storage.CreateRequest{Name: "source", Path: tmp1})
		require.NoError(t, err)
		_, err = Default.CreatePreparationHandler(ctx, db, CreateRequest{
			Name:              "name",
			MaxSizeStr:        "2GB",
			SourceStorages:    []string{"source"},
			DeleteAfterExport: true,
		})
		require.ErrorContains(t, err, "deleteAfterExport cannot be set without output storages")
	})
}

func TestCreatePreparationHandler_NameAllDigits(t *testing.T) {
	tmp1 := t.TempDir()
	tmp2 := t.TempDir()
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, err := storage.Default.CreateStorageHandler(ctx, db, "local", storage.CreateRequest{Name: "source", Path: tmp1})
		require.NoError(t, err)
		_, err = storage.Default.CreateStorageHandler(ctx, db, "local", storage.CreateRequest{Name: "output", Path: tmp2})
		require.NoError(t, err)
		_, err = Default.CreatePreparationHandler(ctx, db, CreateRequest{
			Name:           "1234",
			MaxSizeStr:     "2GB",
			SourceStorages: []string{"source"},
			OutputStorages: []string{"output"},
		})
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "all digits")
	})
}

func TestCreatePreparationHandler_Success(t *testing.T) {
	tmp1 := t.TempDir()
	tmp2 := t.TempDir()
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, err := storage.Default.CreateStorageHandler(ctx, db, "local", storage.CreateRequest{Name: "source", Path: tmp1})
		require.NoError(t, err)
		_, err = storage.Default.CreateStorageHandler(ctx, db, "local", storage.CreateRequest{Name: "output", Path: tmp2})
		require.NoError(t, err)
		preparation, err := Default.CreatePreparationHandler(ctx, db, CreateRequest{
			Name:           "name",
			MaxSizeStr:     "2GB",
			SourceStorages: []string{"source"},
			OutputStorages: []string{"output"},
		})
		require.NoError(t, err)
		require.Greater(t, preparation.ID, uint32(0))
	})
}
