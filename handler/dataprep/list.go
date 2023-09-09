package dataprep

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
)

// ListHandler fetches a list of all the preparations from the database.
//
// This handler retrieves all preparations stored in the database along with their associated
// source and output storages. The result includes all fields and relationships of the preparation models.
//
// Parameters:
// - ctx: The context for managing timeouts and cancellation.
// - request: A handler request. This particular handler does not use the payload, so the request struct is empty.
// - dep: Contains the handler's dependencies, such as the gorm.DB instance.
//
// Returns:
//   - A slice of model.Preparation, each representing a preparation stored in the database
//     along with its associated source and output storages.
//   - An error if any issues occur during the operation, especially database-related errors.
func (DefaultHandler) ListHandler(ctx context.Context,
	request handler.Request[struct{}],
	dep handler.Dependency,
) ([]model.Preparation, error) {
	db := dep.DB.WithContext(ctx)
	var preparations []model.Preparation
	err := db.WithContext(ctx).Preload("SourceStorages").Preload("OutputStorages").Find(&preparations).Error
	return preparations, errors.WithStack(err)
}

// @ID ListPreparations
// @Summary List all preparations
// @Tags Preparation
// @Accept json
// @Produce json
// @Success 200 {array} model.Preparation
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation [get]
func _() {}
