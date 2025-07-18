package statechange

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/statetracker"
	"gorm.io/gorm"
)

// @ID ListStateChanges
// @Summary List all deal state changes with filtering and pagination
// @Description Retrieve a list of deal state changes with optional filtering by deal ID, state, provider, client, and time range. Supports pagination and sorting.
// @Tags State Changes
// @Accept json
// @Produce json
// @Param dealId query string false "Filter by deal ID"
// @Param state query string false "Filter by new state" Enums(proposed,published,active,expired,proposal_expired,rejected,slashed,error)
// @Param providerId query string false "Filter by storage provider ID"
// @Param clientAddress query string false "Filter by client wallet address"
// @Param startTime query string false "Filter changes after this time (RFC3339 format)"
// @Param endTime query string false "Filter changes before this time (RFC3339 format)"
// @Param offset query int false "Number of records to skip for pagination" default(0)
// @Param limit query int false "Maximum number of records to return" default(100)
// @Param orderBy query string false "Field to sort by" default(timestamp) Enums(timestamp,dealId,newState,providerId,clientAddress)
// @Param order query string false "Sort order" default(desc) Enums(asc,desc)
// @Success 200 {object} statechange.StateChangeResponse
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /api/state-changes [get]
func (DefaultHandler) ListStateChangesHandler(ctx context.Context, db *gorm.DB, query model.DealStateChangeQuery) (StateChangeResponse, error) {
	tracker := statetracker.NewStateChangeTracker(db)

	stateChanges, total, err := tracker.GetStateChanges(ctx, query)
	if err != nil {
		return StateChangeResponse{}, errors.Wrap(err, "failed to retrieve state changes")
	}

	response := StateChangeResponse{
		StateChanges: stateChanges,
		Total:        total,
		Offset:       query.Offset,
		Limit:        query.Limit,
	}

	return response, nil
}

// @ID GetDealStateChanges
// @Summary Get state changes for a specific deal
// @Description Retrieve all state changes for a specific deal ordered by timestamp
// @Tags State Changes
// @Accept json
// @Produce json
// @Param id path int true "Deal ID"
// @Success 200 {array} model.DealStateChange
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /api/deals/{id}/state-changes [get]
func (DefaultHandler) GetDealStateChangesHandler(ctx context.Context, db *gorm.DB, dealID model.DealID) ([]model.DealStateChange, error) {
	// First verify the deal exists
	var deal model.Deal
	err := db.First(&deal, dealID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, handlererror.ErrNotFound
		}
		return nil, errors.Wrap(err, "failed to verify deal exists")
	}

	tracker := statetracker.NewStateChangeTracker(db)

	stateChanges, err := tracker.GetStateChangesForDeal(ctx, dealID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve state changes for deal")
	}

	return stateChanges, nil
}

// @ID GetStateChangeStats
// @Summary Get statistics about deal state changes
// @Description Retrieve various statistics about deal state changes including total counts, distribution by state, and recent activity
// @Tags State Changes
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} api.HTTPError
// @Router /api/state-changes/stats [get]
func (DefaultHandler) GetStateChangeStatsHandler(ctx context.Context, db *gorm.DB) (map[string]interface{}, error) {
	tracker := statetracker.NewStateChangeTracker(db)

	stats, err := tracker.GetStateChangeStats(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve state change statistics")
	}

	return stats, nil
}
