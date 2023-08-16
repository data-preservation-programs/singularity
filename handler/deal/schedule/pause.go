package schedule

import (
	"context"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func PauseHandler(
	ctx context.Context,
	db *gorm.DB,
	scheduleID string,
) (*model.Schedule, error) {
	return pauseHandler(ctx, db.WithContext(ctx), scheduleID)
}

// @Summary Pause a specific schedule
// @Tags Deal Schedule
// @Produce json
// @Param id path string true "Schedule ID"
// @Success 200 {object} model.Schedule
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /schedule/{id}/pause [post]
func pauseHandler(
	ctx context.Context,
	db *gorm.DB,
	scheduleID string,
) (*model.Schedule, error) {
	var schedule model.Schedule
	err := db.First(&schedule, "id = ?", scheduleID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr("schedule not found")
	}
	if err != nil {
		return nil, err
	}

	if schedule.State == model.ScheduleCompleted {
		return nil, handler.NewInvalidParameterErr("schedule is already completed")
	}

	schedule.State = model.SchedulePaused
	err = database.DoRetry(ctx, func() error {
		return db.Model(&model.Schedule{}).Where("id = ?", scheduleID).Update("state", model.SchedulePaused).Error
	})
	if err != nil {
		return nil, err
	}

	return &schedule, nil
}
