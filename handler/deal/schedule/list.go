package schedule

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// ListRequest holds optional filters for listing schedules.
type ListRequest struct {
	Group string `json:"group"` // Filter by group label (empty = all)
}

// @ID ListSchedules
// @Summary List all deal making schedules
// @Tags Deal Schedule
// @Produce json
// @Param group query string false "Filter by group label"
// @Success 200 {array} model.Schedule
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /schedule [get]
func _() {}

// ListHandler retrieves schedules from the database, optionally filtered by group.
func (DefaultHandler) ListHandler(
	ctx context.Context,
	db *gorm.DB,
	request ListRequest,
) ([]model.Schedule, error) {
	db = db.WithContext(ctx)
	if request.Group != "" {
		db = db.Where("`group` = ?", request.Group)
	}
	var schedules []model.Schedule
	err := db.Find(&schedules).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return schedules, nil
}
