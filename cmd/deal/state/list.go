package state

import (
	"strconv"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/statechange"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/urfave/cli/v2"
)

// listStateChangesAction handles the list command action
func listStateChangesAction(c *cli.Context) error {
	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() { _ = closer.Close() }()

	query, err := buildStateChangeQuery(c)
	if err != nil {
		return err
	}

	// Get state changes
	response, err := statechange.Default.ListStateChangesHandler(c.Context, db, *query)
	if err != nil {
		return errors.WithStack(err)
	}

	// Handle export if requested
	exportFormat := c.String("export")
	if exportFormat != "" {
		return handleStateChangeExport(c, response.StateChanges, exportFormat, response.Total)
	}

	// Print results to console
	cliutil.Print(c, response)
	return nil
}

// buildStateChangeQuery builds the query from CLI flags
func buildStateChangeQuery(c *cli.Context) (*model.DealStateChangeQuery, error) {
	query := &model.DealStateChangeQuery{}

	// Parse deal ID if provided
	if dealIDStr := c.String("deal-id"); dealIDStr != "" {
		dealID, err := strconv.ParseUint(dealIDStr, 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, "invalid deal ID format")
		}
		dealIDValue := model.DealID(dealID)
		query.DealID = &dealIDValue
	}

	// Parse state if provided
	if stateStr := c.String("state"); stateStr != "" {
		state := model.DealState(stateStr)
		query.State = &state
	}

	// Parse provider ID if provided
	if providerStr := c.String("provider"); providerStr != "" {
		query.ProviderID = &providerStr
	}

	// Parse client address if provided
	if clientStr := c.String("client"); clientStr != "" {
		query.ClientAddress = &clientStr
	}

	// Parse start time if provided
	if startTimeStr := c.String("start-time"); startTimeStr != "" {
		startTime, err := time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			return nil, errors.Wrap(err, "invalid start-time format, expected RFC3339 (e.g., 2023-01-01T00:00:00Z)")
		}
		query.StartTime = &startTime
	}

	// Parse end time if provided
	if endTimeStr := c.String("end-time"); endTimeStr != "" {
		endTime, err := time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			return nil, errors.Wrap(err, "invalid end-time format, expected RFC3339 (e.g., 2023-12-31T23:59:59Z)")
		}
		query.EndTime = &endTime
	}

	// Set pagination
	if c.IsSet("offset") {
		offset := c.Int("offset")
		query.Offset = &offset
	}

	if c.IsSet("limit") {
		limit := c.Int("limit")
		query.Limit = &limit
	}

	// Set sorting
	if orderBy := c.String("order-by"); orderBy != "" {
		query.OrderBy = &orderBy
	}

	if order := c.String("order"); order != "" {
		query.Order = &order
	}

	return query, nil
}

// handleStateChangeExport handles exporting state changes to file
func handleStateChangeExport(c *cli.Context, stateChanges []model.DealStateChange, exportFormat string, totalInDB int64) error {
	outputPath := c.String("output")
	if outputPath == "" {
		timestamp := time.Now().Format("20060102-150405")
		switch exportFormat {
		case "csv":
			outputPath = "statechanges-" + timestamp + ".csv"
		case "json":
			outputPath = "statechanges-" + timestamp + ".json"
		default:
			return errors.Errorf("unsupported export format: %s (supported: csv, json)", exportFormat)
		}
	}

	err := exportStateChanges(stateChanges, exportFormat, outputPath)
	if err != nil {
		return errors.WithStack(err)
	}

	cliutil.Print(c, map[string]interface{}{
		"message":    "State changes exported successfully",
		"format":     exportFormat,
		"outputPath": outputPath,
		"totalCount": len(stateChanges),
		"totalInDB":  totalInDB,
	})
	return nil
}

var ListCmd = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "List deal state changes with optional filtering and pagination",
	Description: `List deal state changes with comprehensive filtering options:
- Filter by deal ID (internal database ID), state, provider, client address, and time range
- Support for pagination and sorting
- Export results to CSV or JSON formats
- Note: Deal IDs and state change IDs are database IDs, not content CIDs`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "deal-id",
			Usage: "Filter by specific deal ID (internal database ID, not CID)",
		},
		&cli.StringFlag{
			Name:  "state",
			Usage: "Filter by deal state (proposed, published, active, expired, proposal_expired, rejected, slashed, error)",
		},
		&cli.StringFlag{
			Name:  "provider",
			Usage: "Filter by storage provider ID (e.g., f01234)",
		},
		&cli.StringFlag{
			Name:  "client",
			Usage: "Filter by client wallet address",
		},
		&cli.StringFlag{
			Name:  "start-time",
			Usage: "Filter changes after this time (RFC3339 format, e.g., 2023-01-01T00:00:00Z)",
		},
		&cli.StringFlag{
			Name:  "end-time",
			Usage: "Filter changes before this time (RFC3339 format, e.g., 2023-12-31T23:59:59Z)",
		},
		&cli.IntFlag{
			Name:        "offset",
			Usage:       "Number of records to skip for pagination",
			DefaultText: "0",
		},
		&cli.IntFlag{
			Name:        "limit",
			Usage:       "Maximum number of records to return",
			DefaultText: "100",
		},
		&cli.StringFlag{
			Name:        "order-by",
			Usage:       "Field to sort by (timestamp, dealId, newState, providerId, clientAddress)",
			DefaultText: "timestamp",
		},
		&cli.StringFlag{
			Name:        "order",
			Usage:       "Sort order (asc, desc)",
			DefaultText: "desc",
		},
		&cli.StringFlag{
			Name:  "export",
			Usage: "Export format (csv, json). If specified, results will be exported to a file instead of displayed",
		},
		&cli.StringFlag{
			Name:  "output",
			Usage: "Output file path for export (optional, defaults to statechanges-<timestamp>.csv/json)",
		},
	},
	Action: listStateChangesAction,
}
