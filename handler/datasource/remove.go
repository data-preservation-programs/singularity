package datasource

import (
	"context"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// @Summary Remove a source
// @Tags Data Source
// @Param id path string true "Source ID"
// @Success 204
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /source/{id} [delete]
func removeSourceHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
) error {
	var source model.Source
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return handlererror.NewInvalidParameterErr("invalid source id")
	}
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return handlererror.NewInvalidParameterErr("source not found")
	}
	if err != nil {
		return errors.WithStack(err)
	}
	err = database.DoRetry(ctx, func() error { return db.Delete(&source).Error })
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func RemoveSourceHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
) error {
	return removeSourceHandler(ctx, db.WithContext(ctx), id)
}
