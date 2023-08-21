package datasource

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func PackHandler(
	db *gorm.DB,
	ctx context.Context,
	resolver storagesystem.HandlerResolver,
	packJobID uint64,
) ([]model.Car, error) {
	return packHandler(db, ctx, resolver, packJobID)
}

// @Summary Pack a pack job into car files
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Pack job ID"
// @Success 201 {object} []model.Car
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /packjob/{id}/pack [post]
func packHandler(
	db *gorm.DB,
	ctx context.Context,
	resolver storagesystem.HandlerResolver,
	packJobID uint64,
) ([]model.Car, error) {
	var packJob model.PackJob
	err := db.Where("id = ?", packJobID).Find(&packJob).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return Pack(ctx, db, packJob, resolver)
}
