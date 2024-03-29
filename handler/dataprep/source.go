package dataprep

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"gorm.io/gorm"
)

// AddSourceStorageHandler associates a given source storage to a Preparation based on the provided ID.
// It first checks if the source storage exists. If it does, it then creates an association
// between the source storage and the specified Preparation.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - id: The ID or name of the Preparation to which the source storage should be attached.
//   - source: The ID or name of the source storage to be attached.
//
// Returns:
//   - A pointer to the updated Preparation model with the new source storage associated.
//   - An error, if any occurred during the verification or attachment process.
//
// Note:
// This function ensures that the given source storage exists and that the given Preparation exists
// before creating an association. It also ensures there are no duplicate associations and handles
// potential errors accordingly.
func (DefaultHandler) AddSourceStorageHandler(ctx context.Context, db *gorm.DB, id string, source string) (*model.Preparation, error) {
	db = db.WithContext(ctx)
	var storage model.Storage
	err := storage.FindByIDOrName(db, source)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "source storage '%s' does not exist", source)
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
		attachment := model.SourceAttachment{
			StorageID:     storage.ID,
			PreparationID: preparation.ID,
		}
		err := db.Create(&attachment).Error
		if err != nil {
			return errors.WithStack(err)
		}

		err = db.Create(&model.Directory{
			AttachmentID: attachment.ID,
		}).Error
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	if util.IsDuplicateKeyError(err) {
		return nil, errors.Wrapf(handlererror.ErrDuplicateRecord, "source storage %s is already attached to preparation %d", source, id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = db.Preload("SourceStorages").Preload("OutputStorages").First(&preparation, preparation.ID).Error

	return &preparation, errors.WithStack(err)
}

// @ID AddSourceStorage
// @Summary Attach a source storage with a preparation
// @Tags Preparation
// @Accept json
// @Produce json
// @Param id path string true "Preparation ID or name"
// @Param name path string true "Source storage ID or name"
// @Success 200 {object} model.Preparation
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/source/{name} [post]
func _() {}
