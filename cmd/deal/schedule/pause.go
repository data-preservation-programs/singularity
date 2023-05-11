package schedule

import (
	"github.com/data-preservation-programs/go-singularity/cmd/cliutil"
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler/deal/schedule"
	"github.com/urfave/cli/v2"
)

var PauseCmd = &cli.Command{
	Name:      "pause",
	Usage:     "Pause a specific schedule",
	ArgsUsage: "SCHEDULE_ID",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		schedule, err := schedule.PauseHandler(c.Args().Get(0), db)
		if err != nil {
			return err.CliError()
		}
		cliutil.PrintToConsole(schedule, c.Bool("json"))
		return nil
	},
}
