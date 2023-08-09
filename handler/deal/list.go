package deal

import (
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

type ListDealRequest struct {
	Datasets  []string `json:"datasets"`  // dataset name filter
	Schedules []uint   `json:"schedules"` // schedule id filter
	Providers []string `json:"providers"` // provider filter
	States    []string `json:"states"`    // state filter
}

func ListHandler(db *gorm.DB, request ListDealRequest) ([]model.Deal, error) {
	return listHandler(db, request)
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
func listHandler(db *gorm.DB, request ListDealRequest) ([]model.Deal, error) {
	var deals []model.Deal
	statement := db
	if len(request.Datasets) > 0 {
		statement = statement.Where("dataset_id IN (?)", statement.Model(&model.Dataset{}).Select("id").Where("name in ?", request.Datasets))
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
		return nil, err
	}
	return deals, nil
}
