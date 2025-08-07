package state

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/statechange"
	"github.com/urfave/cli/v2"
)

// getStateChangeStatsAction handles the stats command action
func getStateChangeStatsAction(c *cli.Context) error {
	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() { _ = closer.Close() }()

	// Get state change statistics
	stats, err := statechange.Default.GetStateChangeStatsHandler(c.Context, db)
	if err != nil {
		return errors.WithStack(err)
	}

	// Print statistics
	cliutil.Print(c, stats)
	return nil
}

var StatsCmd = &cli.Command{
	Name:  "stats",
	Usage: "Get statistics about deal state changes",
	Description: `Get comprehensive statistics about deal state changes including:
- Total number of state changes
- Distribution by state
- Recent activity
- Provider statistics
- Client statistics`,
	Action: getStateChangeStatsAction,
}
