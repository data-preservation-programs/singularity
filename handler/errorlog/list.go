package errorlog

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/errorlog"
	"gorm.io/gorm"
)

// ListErrorLogsHandler fetches and returns a list of error logs with filtering and pagination.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - filters: QueryFilters containing filtering and pagination options.
//
// Returns:
//   - A slice containing error log records that match the filter criteria.
//   - Total count of error logs matching the filters (before pagination).
//   - An error, if any occurred during the database query operation.
func (DefaultHandler) ListErrorLogsHandler(
	ctx context.Context,
	db *gorm.DB,
	filters errorlog.QueryFilters,
) ([]model.ErrorLog, int64, error) {
	logs, total, err := errorlog.QueryErrorLogs(ctx, db, filters)
	return logs, total, errors.WithStack(err)
}

// @ID ListErrorLogs
// @Summary List error logs with filtering and pagination
// @Tags Error Logs
// @Accept json
// @Produce json
// @Param entity_type query string false "Filter by entity type (e.g., deal, preparation, schedule)"
// @Param entity_id query string false "Filter by entity ID"
// @Param component query string false "Filter by component (e.g., onboard, deal_schedule)"
// @Param level query string false "Filter by error level (info, warning, error, critical)"
// @Param event_type query string false "Filter by event type"
// @Param start_time query string false "Filter logs after this time (RFC3339 format)"
// @Param end_time query string false "Filter logs before this time (RFC3339 format)"
// @Param limit query int false "Maximum number of logs to return (default: 50, max: 1000)"
// @Param offset query int false "Number of logs to skip (default: 0)"
// @Success 200 {object} ErrorLogsResponse
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /api/errors [get]
func _() {}

// ErrorLogsResponse represents the API response for listing error logs
type ErrorLogsResponse struct {
	ErrorLogs []model.ErrorLog `json:"errorLogs"`
	Total     int64            `json:"total"`
	Limit     int              `json:"limit"`
	Offset    int              `json:"offset"`
	HasMore   bool             `json:"hasMore"`
}
