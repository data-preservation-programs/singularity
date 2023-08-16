package datasource

import (
	"context"
	"strconv"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
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
		return handler.NewInvalidParameterErr("invalid source id")
	}
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return handler.NewInvalidParameterErr("source not found")
	}
	if err != nil {
		return err
	}
	err = database.DoRetry(ctx, func() error { return db.Delete(&source).Error })
	if err != nil {
		return err
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
