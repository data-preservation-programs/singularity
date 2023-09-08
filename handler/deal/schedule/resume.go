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

var resumableStates = []model.ScheduleState{
	model.SchedulePaused,
	model.ScheduleError,
}

// ResumeHandler attempts to resume a previously paused schedule based on the provided scheduleID.
//
// Parameters:
// - ctx: The context for the operation, allowing for cancellation and timeouts.
// - db: The database connection used for operations.
// - scheduleID: The ID of the schedule to be resumed.
//
// Returns:
// - A pointer to the updated Schedule if successful.
// - An error if any issues occur, e.g., if the schedule is not found or not in a paused state.
func (DefaultHandler) ResumeHandler(
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
	if !slices.Contains(resumableStates, schedule.State) {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "schedule %d is not resumable, current state: %s", scheduleID, schedule.State)
	}
	schedule.State = model.ScheduleActive
	err = database.DoRetry(ctx, func() error {
		return db.Model(&model.Schedule{}).Where("id = ?", scheduleID).Update("state", model.ScheduleActive).Error
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &schedule, nil
}

// @Summary Resume a specific schedule
// @Tags Deal Schedule
// @Produce json
// @Param id path string true "Schedule ID"
// @Success 200 {object} model.Schedule
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /schedule/{id}/resume [post]
func _() {}
