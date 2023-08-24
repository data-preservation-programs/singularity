package deal

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

type ListDealRequest struct {
	Preparations []uint32          `json:"preparations"` // preparation ID filter
	Storages     []string          `json:"storages"`     // storage filter
	Schedules    []uint32          `json:"schedules"`    // schedule id filter
	Providers    []string          `json:"providers"`    // provider filter
	States       []model.DealState `json:"states"`       // state filter
}

func ListHandler(ctx context.Context, db *gorm.DB, request ListDealRequest) ([]model.Deal, error) {
	db = db.WithContext(ctx)
	var deals []model.Deal
	statement := db
	if len(request.Preparations) > 0 {
		statement = statement.Where("schedule_id IN (?)", db.Model(&model.Schedule{}).Select("id").
			Where("preparation_id in ?", request.Preparations))
	}

	if len(request.Storages) > 0 {
		statement = statement.Where("schedule_id IN (?)", db.Model(&model.Schedule{}).Select("id").
			Where("preparation_id in (?)", db.Model(&model.SourceAttachment{}).Select("preparation_id").
				Where("storage_id in (?)", db.Model(&model.Storage{}).Select("id").
					Where("name in ?", request.Storages))))
	}

	if len(request.Schedules) > 0 {
		statement = statement.Where("schedule_id IN ?", request.Schedules)
	}

	if len(request.Providers) > 0 {
		statement = statement.Where("provider IN ?", request.Providers)
	}

	if len(request.States) > 0 {
		statement = statement.Where("state IN ?", request.States)
	}

	// We did not create indexes for all above query and it should be fine for now
	err := db.Where(statement).Find(&deals).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return deals, nil
}

// @Summary List all deals
// @Description List all deals
// @Tags Deal
// @Accept json
// @Produce json
// @Param request body ListDealRequest true "ListDealRequest"
// @Success 200 {array} model.Deal
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /deal [post]
func _() {}
