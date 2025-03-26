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
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - id: The ID or name of the Preparation to which the output storage should be attached.
//   - output: The ID or name of the output storage to be attached.
//
// Returns:
//   - A pointer to the updated Preparation model with the new output storage associated.
//   - An error, if any occurred during the verification or attachment process.
//
// Note:
// This function performs several checks to ensure the output storage and the Preparation exist.
// It also checks for potential duplicate associations and handles potential errors accordingly.
func (DefaultHandler) AddOutputStorageHandler(ctx context.Context, db *gorm.DB, id string, output string) (*model.Preparation, error) {
	db = db.WithContext(ctx)
	var storage model.Storage
	err := storage.FindByIDOrName(db, output)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "output storage '%s' does not exist", output)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var preparation model.Preparation
	err = preparation.FindByIDOrName(db, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation '%s' does not exist", id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = database.DoRetry(ctx, func() error {
		return db.Create(&model.OutputAttachment{
			StorageID:     storage.ID,
			PreparationID: preparation.ID,
		}).Error
	})
	if util.IsDuplicateKeyError(err) {
		return nil, errors.Wrapf(handlererror.ErrDuplicateRecord, "output storage %s is already attached to preparation %d", output, id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = db.Preload("SourceStorages").Preload("OutputStorages").First(&preparation, preparation.ID).Error

	return &preparation, errors.WithStack(err)
}

// @ID AddOutputStorage
// @Summary Attach an output storage with a preparation
// @Tags Preparation
// @Accept json
// @Produce json
// @Param id path string true "Preparation ID or name"
// @Param name path string true "Output storage ID or name"
// @Success 200 {object} model.Preparation
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/output/{name} [post]
func _() {}

// @ID RemoveOutputStorage
// @Summary Detach an output storage from a preparation
// @Tags Preparation
// @Accept json
// @Produce json
// @Param id path string true "Preparation ID or name"
// @Param name path string true "Output storage ID or name"
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
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - id: The ID or name of the Preparation from which the output storage should be detached.
//   - output: The ID or name of the output storage to be detached.
//
// Returns:
//   - A pointer to the updated Preparation model with the output storage removed.
//   - An error, if any occurred during the verification or detachment process.
//
// Note:
// This function performs several validation steps to ensure integrity while removing the association.
// It also preloads associated storages to return an updated version of the Preparation.
func (DefaultHandler) RemoveOutputStorageHandler(ctx context.Context, db *gorm.DB, id string, output string) (*model.Preparation, error) {
	db = db.WithContext(ctx)
	var storage model.Storage
	err := storage.FindByIDOrName(db, output)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "output storage '%s' does not exist", output)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var preparation model.Preparation
	err = preparation.FindByIDOrName(db, id, "OutputStorages")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation '%s' does not exist", id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if !underscore.Any(preparation.OutputStorages, func(s model.Storage) bool { return s.ID == storage.ID }) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "output storage %s is not attached to preparation %s", output, id)
	}

	if preparation.DeleteAfterExport && len(preparation.OutputStorages) == 1 {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "cannot remove the only output storage from a preparation with deleteAfterExport enabled")
	}

	if preparation.NoInline && len(preparation.OutputStorages) == 1 {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "cannot remove the only output storage from a preparation in non-inline mode")
	}

	if !preparation.NoInline && !preparation.NoDag && len(preparation.OutputStorages) == 1 {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "cannot remove the only output storage from a preparation in inline mode with DAG generation enabled")
	}

	err = database.DoRetry(ctx, func() error {
		return db.Where("storage_id = ? AND preparation_id = ?", storage.ID, preparation.ID).Delete(&model.OutputAttachment{}).Error
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = db.Preload("SourceStorages").Preload("OutputStorages").First(&preparation, preparation.ID).Error

	return &preparation, errors.WithStack(err)
}
