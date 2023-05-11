package schedule

import (
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// ResumeHandler godoc
// @Summary Resume a specific schedule
// @Tags Deal Schedule
// @Produce json
// @Param scheduleID path string true "Schedule ID"
// @Success 200 {object} model.Schedule
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /deal/schedule/{scheduleID}/resume [post]
func ResumeHandler(
	db *gorm.DB,
	scheduleID string,
) (*model.Schedule, *handler.Error) {
	var schedule model.Schedule
	err := db.Transaction(func(db *gorm.DB) error {
		err := db.First(&schedule, "id = ?", scheduleID).Error
		if err != nil {
			return err
		}
		schedule.State = model.ScheduleActive
		err = db.Where("id = ?", scheduleID).Update("state", model.ScheduleActive).Error
		if err != nil {
			return err
		}
		return nil
	})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewBadRequestString("schedule not found")
	}
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}
	return &schedule, nil
}
