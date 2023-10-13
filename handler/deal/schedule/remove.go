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

var removableStates = []model.ScheduleState{
	model.ScheduleError, model.ScheduleCompleted, model.SchedulePaused,
}

func (DefaultHandler) RemoveHandler(
	ctx context.Context,
	db *gorm.DB,
	scheduleID uint32,
) error {
	db = db.WithContext(ctx)
	var schedule model.Schedule
	err := db.First(&schedule, "id = ?", scheduleID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrapf(handlererror.ErrNotFound, "schedule %d not found", scheduleID)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if !slices.Contains(removableStates, schedule.State) {
		return errors.Wrapf(handlererror.ErrInvalidParameter, "schedule %d is not removable, current state: %s", scheduleID, schedule.State)
	}

	err = database.DoRetry(ctx, func() error {
		return db.Model(&model.Schedule{}).Delete(&schedule).Error
	})

	return errors.WithStack(err)
}

// @ID RemoveSchedule
// @Summary Delete a specific schedule
// @Tags Deal Schedule
// @Produce json
// @Param id path int true "Schedule ID"
// @Success 204
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /schedule/{id} [delete]
func _() {}
