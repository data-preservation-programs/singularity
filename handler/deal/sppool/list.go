package sppool

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func (DefaultHandler) ListHandler(
	ctx context.Context,
	db *gorm.DB,
) ([]model.SPPool, error) {
	db = db.WithContext(ctx)
	var pools []model.SPPool
	err := db.Preload("Providers").Preload("Preparations").Find(&pools).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return pools, nil
}

// @ID ListSPPools
// @Summary List all SP Pools
// @Tags SP Pool
// @Produce json
// @Success 200 {array} model.SPPool
// @Failure 500 {object} api.HTTPError
// @Router /sp-pool [get]
func _() {}
