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
	Name string `json:"name" binding:"required"`
}

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
// @Param request body RenameRequest true "New preparation name"
// @Accept json
// @Produce json
// @Success 200 {object} model.Preparation
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /Preparation/{name}/rename [patch]
func _() {}
