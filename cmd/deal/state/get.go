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

// getDealStateChangesAction handles the get command action
func getDealStateChangesAction(c *cli.Context) error {
	if c.NArg() != 1 {
		return errors.New("deal ID is required")
	}

	dealIDStr := c.Args().Get(0)
	dealID, err := strconv.ParseUint(dealIDStr, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid deal ID format")
	}

	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() { _ = closer.Close() }()

	// Get state changes for the specific deal
	stateChanges, err := statechange.Default.GetDealStateChangesHandler(c.Context, db, model.DealID(dealID))
	if err != nil {
		return errors.WithStack(err)
	}

	// Handle export if requested
	exportFormat := c.String("export")
	if exportFormat != "" {
		return handleDealStateExport(c, stateChanges, dealIDStr, exportFormat)
	}

	// Print results to console
	if len(stateChanges) == 0 {
		cliutil.Print(c, map[string]interface{}{
			"message": "No state changes found for deal " + dealIDStr,
		})
		return nil
	}

	cliutil.Print(c, stateChanges)
	return nil
}

// handleDealStateExport handles exporting deal state changes to file
func handleDealStateExport(c *cli.Context, stateChanges []model.DealStateChange, dealIDStr string, exportFormat string) error {
	outputPath := c.String("output")
	if outputPath == "" {
		timestamp := time.Now().Format("20060102-150405")
		switch exportFormat {
		case "csv":
			outputPath = "deal-" + dealIDStr + "-states-" + timestamp + ".csv"
		case "json":
			outputPath = "deal-" + dealIDStr + "-states-" + timestamp + ".json"
		default:
			return errors.Errorf("unsupported export format: %s (supported: csv, json)", exportFormat)
		}
	}

	err := exportStateChanges(stateChanges, exportFormat, outputPath)
	if err != nil {
		return errors.WithStack(err)
	}

	cliutil.Print(c, map[string]interface{}{
		"message":    "Deal state changes exported successfully",
		"dealId":     dealIDStr,
		"format":     exportFormat,
		"outputPath": outputPath,
		"count":      len(stateChanges),
	})
	return nil
}

var GetCmd = &cli.Command{
	Name:      "get",
	Usage:     "Get state changes for a specific deal",
	ArgsUsage: "<deal-id>",
	Description: `Get all state changes for a specific deal ordered by timestamp.
This command shows the complete state transition history for a given deal.
Note: deal-id refers to the internal database ID, not a content CID.`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "export",
			Usage: "Export format (csv, json). If specified, results will be exported to a file instead of displayed",
		},
		&cli.StringFlag{
			Name:  "output",
			Usage: "Output file path for export (optional, defaults to deal-<deal-id>-states-<timestamp>.csv/json)",
		},
	},
	Action: getDealStateChangesAction,
}
