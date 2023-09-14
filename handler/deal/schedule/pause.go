package schedule

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

var pausableStates = []model.ScheduleState{
	model.ScheduleActive,
}

// PauseHandler attempts to pause an active schedule based on the provided scheduleID.
// If the schedule is already completed, an error will be returned.
//
// Parameters:
//   - ctx: The context for the operation, which can include cancellation signals, timeout details, etc.
//   - db: The database connection used for CRUD operations.
//   - scheduleID: The ID of the schedule to be paused.
//
// Returns:
//   - A pointer to the updated Schedule if successful.
//   - An error if there are issues during the operation, e.g., if the schedule is not found or already completed.
func (DefaultHandler) PauseHandler(
	ctx context.Context,
	db *gorm.DB,
	scheduleID uint32,
) (*model.Schedule, error) {
	db = db.WithContext(ctx)
	var schedule model.Schedule
	err := db.First(&schedule, "id = ?", scheduleID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "schedule %d not found", scheduleID)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if !slices.Contains(pausableStates, schedule.State) {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "schedule %d is not pausable, current state: %s", scheduleID, schedule.State)
	}

	schedule.State = model.SchedulePaused
	err = database.DoRetry(ctx, func() error {
		return db.Model(&model.Schedule{}).Where("id = ?", scheduleID).Update("state", model.SchedulePaused).Error
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &schedule, nil
}

// @ID PauseSchedule
// @Summary Pause a specific schedule
// @Tags Deal Schedule
// @Produce json
// @Param id path int true "Schedule ID"
// @Success 200 {object} model.Schedule
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /schedule/{id}/pause [post]
func _() {}
