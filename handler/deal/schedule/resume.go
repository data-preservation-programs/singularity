package schedule

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func ResumeHandler(
	db *gorm.DB,
	scheduleID string,
) (*model.Schedule, error) {
	return resumeHandler(db, scheduleID)
}

// @Summary Resume a specific schedule
// @Tags Deal Schedule
// @Produce json
// @Param id path string true "Schedule ID"
// @Success 200 {object} model.Schedule
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /schedule/{id}/resume [post]
func resumeHandler(
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
	if schedule.State != model.SchedulePaused {
		return nil, handler.NewInvalidParameterErr("schedule is not paused")
	}
	schedule.State = model.ScheduleActive
	err = db.Model(&model.Schedule{}).Where("id = ?", scheduleID).Update("state", model.ScheduleActive).Error
	if err != nil {
		return nil, err
	}

	return &schedule, nil
}
