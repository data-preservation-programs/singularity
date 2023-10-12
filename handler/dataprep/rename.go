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

type RenameRequest struct {
	Name string `binding:"required" json:"name"`
}

// RenamePreparationHandler updates the name of a preparation entry in the database.
// This handler finds a preparation entry by its ID or name and then updates its name with the new name
// provided in the request payload. The new name cannot be entirely numeric or empty.
//
// Parameters:
//   - ctx: The context for managing timeouts and cancellation.
//   - db: The gorm.DB instance for making database queries.
//   - name: The current name or ID of the preparation entry to be renamed.
//   - request: A RenameRequest object containing the new name for the preparation entry.
//
// Returns:
//   - A pointer to the updated model.Preparation entry, reflecting the new name.
//   - An error if any issues occur during the operation, especially if the provided new name is invalid,
//     or if there are database-related errors.
func (DefaultHandler) RenamePreparationHandler(
	ctx context.Context,
	db *gorm.DB,
	name string,
	request RenameRequest,
) (*model.Preparation, error) {
	db = db.WithContext(ctx)
	if util.IsAllDigits(request.Name) || request.Name == "" {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "preparation name %s cannot be all digits or empty", name)
	}

	var preparation model.Preparation
	err := preparation.FindByIDOrName(db, name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %s does not exist", name)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	preparation.Name = request.Name
	err = database.DoRetry(ctx, func() error {
		return db.Model(&model.Preparation{}).Where("id = ?", preparation.ID).Update("name", preparation.Name).Error
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &preparation, nil
}

// @ID RenamePreparation
// @Summary Rename a preparation
// @Tags Preparation
// @Param name path string true "Preparation ID or name"
// @Param request body RenameRequest true "Preparation Request"
// @Accept json
// @Produce json
// @Success 200 {object} model.Preparation
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{name}/rename [patch]
func _() {}
