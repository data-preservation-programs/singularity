package errorlog

import (
	"fmt"
	"os"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/errorlog"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:     "list",
	Usage:    "List error logs with filtering and pagination",
	Category: "Error Log Management",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "entity-type",
			Usage: "Filter by entity type (e.g., deal, preparation, schedule)",
		},
		&cli.StringFlag{
			Name:  "entity-id",
			Usage: "Filter by entity ID",
		},
		&cli.StringFlag{
			Name:  "component",
			Usage: "Filter by component (e.g., onboard, deal_schedule)",
		},
		&cli.StringFlag{
			Name:  "level",
			Usage: "Filter by error level (info, warning, error, critical)",
		},
		&cli.StringFlag{
			Name:  "event-type",
			Usage: "Filter by event type",
		},
		&cli.StringFlag{
			Name:  "start-time",
			Usage: "Filter logs after this time (RFC3339 format, e.g., 2023-01-01T00:00:00Z)",
		},
		&cli.StringFlag{
			Name:  "end-time",
			Usage: "Filter logs before this time (RFC3339 format, e.g., 2023-12-31T23:59:59Z)",
		},
		&cli.IntFlag{
			Name:  "limit",
			Usage: "Maximum number of logs to return",
			Value: 50,
		},
		&cli.IntFlag{
			Name:  "offset",
			Usage: "Number of logs to skip",
			Value: 0,
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()

		// Parse time filters
		var startTime, endTime *time.Time
		if startTimeStr := c.String("start-time"); startTimeStr != "" {
			t, err := time.Parse(time.RFC3339, startTimeStr)
			if err != nil {
				return errors.Wrap(err, "invalid start-time format, expected RFC3339")
			}
			startTime = &t
		}
		if endTimeStr := c.String("end-time"); endTimeStr != "" {
			t, err := time.Parse(time.RFC3339, endTimeStr)
			if err != nil {
				return errors.Wrap(err, "invalid end-time format, expected RFC3339")
			}
			endTime = &t
		}

		// Validate error level if provided
		level := c.String("level")
		if level != "" {
			switch model.ErrorLevel(level) {
			case model.ErrorLevelInfo, model.ErrorLevelWarning, model.ErrorLevelError, model.ErrorLevelCritical:
				// Valid level
			default:
				return errors.Errorf("invalid error level: %s. Valid levels are: info, warning, error, critical", level)
			}
		}

		filters := errorlog.QueryFilters{
			EntityType: c.String("entity-type"),
			EntityID:   c.String("entity-id"),
			Component:  c.String("component"),
			Level:      model.ErrorLevel(level),
			EventType:  c.String("event-type"),
			StartTime:  startTime,
			EndTime:    endTime,
			Limit:      c.Int("limit"),
			Offset:     c.Int("offset"),
		}

		logs, total, err := errorlog.QueryErrorLogs(c.Context, db, filters)
		if err != nil {
			return errors.WithStack(err)
		}

		// Create response with pagination info
		response := struct {
			ErrorLogs     []model.ErrorLog       `json:"errorLogs"`
			Total         int64                  `json:"total"`
			Limit         int                    `json:"limit"`
			Offset        int                    `json:"offset"`
			HasMore       bool                   `json:"hasMore"`
			FilterSummary map[string]interface{} `json:"filterSummary,omitempty"`
		}{
			ErrorLogs: logs,
			Total:     total,
			Limit:     filters.Limit,
			Offset:    filters.Offset,
			HasMore:   int64(filters.Offset+len(logs)) < total,
		}

		// Add filter summary if verbose mode is enabled
		if c.Bool("verbose") || c.Bool("json") {
			filterSummary := make(map[string]interface{})
			if filters.EntityType != "" {
				filterSummary["entityType"] = filters.EntityType
			}
			if filters.EntityID != "" {
				filterSummary["entityId"] = filters.EntityID
			}
			if filters.Component != "" {
				filterSummary["component"] = filters.Component
			}
			if filters.Level != "" {
				filterSummary["level"] = filters.Level
			}
			if filters.EventType != "" {
				filterSummary["eventType"] = filters.EventType
			}
			if filters.StartTime != nil {
				filterSummary["startTime"] = filters.StartTime.Format(time.RFC3339)
			}
			if filters.EndTime != nil {
				filterSummary["endTime"] = filters.EndTime.Format(time.RFC3339)
			}
			if len(filterSummary) > 0 {
				response.FilterSummary = filterSummary
			}
		}

		// Print pagination info to stderr if not in JSON mode
		if !c.Bool("json") {
			fmt.Fprintf(os.Stderr, "Total error logs: %d\n", total)
			if response.HasMore {
				fmt.Fprintf(os.Stderr, "Showing %d of %d results\n", len(logs), total)
				nextOffset := filters.Offset + len(logs)
				fmt.Fprintf(os.Stderr, "To see more results, use: --offset %d\n", nextOffset)
			}
		}

		cliutil.Print(c, response)
		return nil
	},
}
