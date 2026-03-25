package sppool

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func (DefaultHandler) GetHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
) (*model.SPPool, error) {
	db = db.WithContext(ctx)
	var pool model.SPPool
	err := db.Preload("Providers").Preload("Preparations").First(&pool, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "sp pool %d not found", id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &pool, nil
}

// @ID GetSPPool
// @Summary Get an SP Pool by ID
// @Tags SP Pool
// @Produce json
// @Param id path int true "SP Pool ID"
// @Success 200 {object} model.SPPool
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /sp-pool/{id} [get]
func _() {}
