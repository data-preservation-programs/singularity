package deal

import (
	"context"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

type ListDealRequest struct {
	Preparations []string          `json:"preparations"` // preparation ID or name filter
	Sources      []string          `json:"sources"`      // source ID or name filter
	Schedules    []uint32          `json:"schedules"`    // schedule id filter
	Providers    []string          `json:"providers"`    // provider filter
	States       []model.DealState `json:"states"`       // state filter
}

// ListHandler retrieves a list of deals from the database based on the specified filtering criteria in ListDealRequest.
//
// The function takes advantage of the conditional nature of the ListDealRequest to construct the final query. It
// filters deals based on various conditions such as preparations, storages, schedules, providers, and states
// as specified in the request.
//
// The function begins by associating the provided context with the database connection. It then successively builds
// upon a GORM statement by appending where clauses based on the parameters in the request.
//
// It's important to note that there aren't indexes for all the query fields in the current database setup.
// This might be sufficient for smaller datasets but could affect performance on larger datasets or under heavy query loads.
//
// Parameters:
//   - ctx:      The context for the operation which provides facilities for timeouts and cancellations.
//   - db:       The database connection for performing CRUD operations related to deals.
//   - request:  The request object which contains the filtering criteria for the deals retrieval.
//
// Returns:
//   - A slice of model.Deal objects matching the filtering criteria.
//   - An error indicating any issues that occurred during the database operation.
func (DefaultHandler) ListHandler(ctx context.Context, db *gorm.DB, request ListDealRequest) ([]model.Deal, error) {
	db = db.WithContext(ctx)
	var deals []model.Deal
	statement := db
	if len(request.Preparations) > 0 {
		var ids []uint64
		var names []string
		for _, preparation := range request.Preparations {
			if id, err := strconv.ParseUint(preparation, 10, 32); err == nil {
				ids = append(ids, id)
			} else {
				names = append(names, preparation)
			}
		}
		statement = statement.Where("schedule_id IN (?)", db.Model(&model.Schedule{}).Select("id").
			Where("preparation_id in (?)", db.Model(&model.Preparation{}).Select("id").
				Where("id in ? OR name in ?", ids, names)))
	}

	if len(request.Sources) > 0 {
		var ids []uint64
		var names []string
		for _, source := range request.Sources {
			if id, err := strconv.ParseUint(source, 10, 32); err == nil {
				ids = append(ids, id)
			} else {
				names = append(names, source)
			}
		}
		statement = statement.Where("schedule_id IN (?)", db.Model(&model.Schedule{}).Select("id").
			Where("preparation_id in (?)", db.Model(&model.SourceAttachment{}).Select("preparation_id").
				Where("storage_id in (?)", db.Model(&model.Storage{}).Select("id").
					Where("id in ? OR name in ?", ids, names))))
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

// @ID ListDeals
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
