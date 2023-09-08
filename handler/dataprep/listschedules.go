package dataprep

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// ListSchedulesHandler retrieves and returns the list of schedules associated with the specified preparation.
//
// This function searches for a preparation in the database based on the given ID (which could be either a primary key ID or name).
// Once the preparation is successfully fetched, it retrieves all schedules linked with this preparation.
//
// Parameters:
// - ctx: The context for managing timeouts and cancellation.
// - db: The gorm.DB instance for database operations.
// - id: The ID or name of the preparation to find associated schedules for.
//
// Returns:
// - A slice of model.Schedule if the operation is successful.
// - An error if any issues occur during the operation, including database retrieval errors.
func (DefaultHandler) ListSchedulesHandler(
	ctx context.Context,
	db *gorm.DB,
	id string) ([]model.Schedule, error) {
	db = db.WithContext(ctx)

	var preparation model.Preparation
	err := preparation.FindByIDOrName(db, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation '%s' does not exist", id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var schedules []model.Schedule
	err = db.Where("preparation_id = ?", preparation.ID).Find(&schedules).Error
	return schedules, errors.WithStack(err)
}

// @Summary List all schedules for a preparation
// @Tags Deal Schedule
// @Accept json
// @Produce json
// @Param id path string true "Preparation ID or name"
// @Success 200 {array} model.Schedule
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/schedules [get]
func _() {}
