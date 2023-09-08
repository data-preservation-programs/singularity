package schedule

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:  "list",
	Usage: "List all deal making schedules",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		schedules, err := schedule.Default.ListHandler(c.Context, db)
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.Print(c, schedules)
		return nil
	},
}
