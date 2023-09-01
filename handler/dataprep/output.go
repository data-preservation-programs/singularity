package dataprep

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

// AddOutputStorageHandler associates a given output storage to a Preparation based on the provided ID.
// The function verifies the existence of the output storage and the specified Preparation.
// If both are valid, it creates an association between the output storage and the Preparation.
//
// Parameters:
// - ctx: The context for database transactions and other operations.
// - db: A pointer to the gorm.DB instance representing the database connection.
// - id: The ID of the Preparation to which the output storage should be attached.
// - output: The name of the output storage to be attached.
//
// Returns:
// - A pointer to the updated Preparation model with the new output storage associated.
// - An error, if any occurred during the verification or attachment process.
//
// Note:
// This function performs several checks to ensure the output storage and the Preparation exist.
// It also checks for potential duplicate associations and handles potential errors accordingly.
func (DefaultHandler) AddOutputStorageHandler(ctx context.Context, db *gorm.DB, id uint32, output string) (*model.Preparation, error) {
	db = db.WithContext(ctx)
	var storage model.Storage
	err := db.Where("name = ?", output).First(&storage).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "output storage '%s' does not exist", output)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = database.DoRetry(ctx, func() error {
		return db.Create(&model.OutputAttachment{
			StorageID:     storage.ID,
			PreparationID: id,
		}).Error
	})
	if util.IsForeignKeyConstraintError(err) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %d does not exist", id)
	}
	if util.IsDuplicateKeyError(err) {
		return nil, errors.Wrapf(handlererror.ErrDuplicateRecord, "output storage %s is already attached to preparation %d", output, id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var preparation model.Preparation
	err = db.Preload("SourceStorages").Preload("OutputStorages").First(&preparation, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %d does not exist", id)
	}

	return &preparation, errors.WithStack(err)
}

// @Summary Attach an output storage with a preparation
// @Tags Preparation
// @Accept json
// @Produce json
// @Param id path int true "Preparation ID"
// @Param name path string true "Output storage name"
// @Success 200 {object} model.Preparation
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/output/{name} [post]
func _() {}

// @Summary Detach an output storage from a preparation
// @Tags Preparation
// @Accept json
// @Produce json
// @Param id path int true "Preparation ID"
// @Param name path string true "Output storage name"
// @Success 200 {object} model.Preparation
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/output/{name} [delete]
func _() {}

// RemoveOutputStorageHandler disassociates a specified output storage from a Preparation using the provided ID.
// It ensures that the output storage and Preparation both exist before attempting the removal.
// Special checks are in place to ensure:
//  1. The output storage is currently attached to the Preparation.
//  2. Removing the only output storage while using encryption is disallowed.
//
// Parameters:
// - ctx: The context for database transactions and other operations.
// - db: A pointer to the gorm.DB instance representing the database connection.
// - id: The ID of the Preparation from which the output storage should be detached.
// - output: The name of the output storage to be detached.
//
// Returns:
// - A pointer to the updated Preparation model with the output storage removed.
// - An error, if any occurred during the verification or detachment process.
//
// Note:
// This function performs several validation steps to ensure integrity while removing the association.
// It also preloads associated storages to return an updated version of the Preparation.
func (DefaultHandler) RemoveOutputStorageHandler(ctx context.Context, db *gorm.DB, id uint32, output string) (*model.Preparation, error) {
	db = db.WithContext(ctx)
	var storage model.Storage
	err := db.Where("name = ?", output).First(&storage).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "output storage '%s' does not exist", output)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var preparation model.Preparation
	err = db.Preload("OutputStorages").First(&preparation, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %d does not exist", id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if !underscore.Any(preparation.OutputStorages, func(s model.Storage) bool { return s.ID == storage.ID }) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "output storage %s is not attached to preparation %d", output, id)
	}

	if preparation.DeleteAfterExport && len(preparation.OutputStorages) == 1 {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "cannot remove the only output storage from a preparation with deleteAfterExport enabled")
	}

	err = database.DoRetry(ctx, func() error {
		return db.Where("storage_id = ? AND preparation_id = ?", storage.ID, id).Delete(&model.OutputAttachment{}).Error
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = db.Preload("SourceStorages").Preload("OutputStorages").First(&preparation, id).Error

	return &preparation, errors.WithStack(err)
}
