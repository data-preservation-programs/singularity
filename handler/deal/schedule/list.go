package schedule

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// @Summary List all deal making schedules
// @Tags Deal Schedule
// @Produce json
// @Success 200 {array} model.Schedule
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /schedules [get]
func _() {}

// ListHandler retrieves all the schedules from the database.
//
// Parameters:
// - ctx: The context for the operation, which can include cancellation signals, timeout details, etc.
// - db: The database connection used for CRUD operations.
//
// Returns:
// - A slice of Schedule models if successful.
// - An error if there are issues during the operation.
func (DefaultHandler) ListHandler(
	ctx context.Context,
	db *gorm.DB,
) ([]model.Schedule, error) {
	db = db.WithContext(ctx)
	var schedules []model.Schedule
	err := db.Find(&schedules).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return schedules, nil
}
