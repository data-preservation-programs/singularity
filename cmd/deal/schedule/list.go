package schedule

import (
	"github.com/data-preservation-programs/go-singularity/cmd/cliutil"
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler/deal/schedule"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:      "list",
	Usage:     "List all deal making schedules",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		schedules, err := schedule.ListHandler(db)
		if err != nil {
			return err.CliError()
		}

		cliutil.PrintToConsole(schedules, c.Bool("json"))
		return nil
	},
}
