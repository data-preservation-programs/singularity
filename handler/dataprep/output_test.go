package dataprep

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestAddOutputStorageHandler_StorageNotFound(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{}).Error
	require.NoError(t, err)

	_, err = AddOutputStorageHandler(ctx, db, 1, "not found")
	require.ErrorIs(t, err, handlererror.ErrNotFound)
	require.ErrorContains(t, err, "output storage")
}

func TestAddOutputStorageHandler_PreparationNotFound(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		OutputStorages: []model.Storage{{
			Name: "output",
		}},
	}).Error
	require.NoError(t, err)

	_, err = AddOutputStorageHandler(ctx, db, 100, "output")
	require.ErrorIs(t, err, handlererror.ErrNotFound)
	require.ErrorContains(t, err, "preparation")
}

func TestAddOutputStorageHandler_AlreadyAttached(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		OutputStorages: []model.Storage{{
			Name: "output",
		}},
	}).Error
	require.NoError(t, err)

	_, err = AddOutputStorageHandler(ctx, db, 1, "output")
	require.ErrorIs(t, err, handlererror.ErrDuplicateRecord)
	require.ErrorContains(t, err, "already")
}

func TestAddOutputStorageHandler_Success(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		OutputStorages: []model.Storage{{
			Name: "output",
		}},
	}).Error
	require.NoError(t, err)

	err = db.Create(&model.Storage{
		Name: "output2",
	}).Error

	preparation, err := AddOutputStorageHandler(ctx, db, 1, "output2")
	require.NoError(t, err)
	require.Len(t, preparation.OutputStorages, 2)
}

func TestRemoveOutputStorageHandler_StorageNotFound(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{}).Error
	require.NoError(t, err)

	_, err = RemoveOutputStorageHandler(ctx, db, 1, "not found")
	require.ErrorIs(t, err, handlererror.ErrNotFound)
	require.ErrorContains(t, err, "output storage")
}

func TestRemoveOutputStorageHandler_PreparationNotFound(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		OutputStorages: []model.Storage{{
			Name: "output",
		}},
	}).Error
	require.NoError(t, err)

	_, err = RemoveOutputStorageHandler(ctx, db, 100, "output")
	require.ErrorIs(t, err, handlererror.ErrNotFound)
	require.ErrorContains(t, err, "preparation")
}

func TestRemoveOutputStorageHandler_InlinePreparationWithEncryption(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		EncryptionRecipients: []string{"aaa"},
		OutputStorages: []model.Storage{{
			Name: "output",
		}},
	}).Error
	require.NoError(t, err)

	_, err = RemoveOutputStorageHandler(ctx, db, 1, "output")
	require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	require.ErrorContains(t, err, "inline")
}

func TestRemoveOutputStorageHandler_NotAttached(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{}).Error
	require.NoError(t, err)

	err = db.Create(&model.Storage{
		Name: "output",
	}).Error

	_, err = RemoveOutputStorageHandler(ctx, db, 1, "output")
	require.ErrorIs(t, err, handlererror.ErrNotFound)
	require.ErrorContains(t, err, "not attached")
}

func TestRemoveOutputStorageHandler_Success(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		OutputStorages: []model.Storage{{
			Name: "output",
		}},
	}).Error
	require.NoError(t, err)

	preparation, err := RemoveOutputStorageHandler(ctx, db, 1, "output")
	require.NoError(t, err)
	require.Len(t, preparation.OutputStorages, 0)
}
