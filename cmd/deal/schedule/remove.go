package schedule

import (
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/urfave/cli/v2"
)

var RemoveCmd = &cli.Command{
	Name:        "remove",
	Usage:       "Remove a paused or completed schedule",
	Before:      cliutil.CheckNArgs,
	ArgsUsage:   "<schedule_id>",
	Description: "Note: all deals made by this schedule will remain for tracking purpose.",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		scheduleID, err := strconv.ParseUint(c.Args().Get(0), 10, 32)
		if err != nil {
			return errors.Wrapf(err, "failed to parse schedule ID %s", c.Args().Get(0))
		}

		err = schedule.Default.RemoveHandler(c.Context, db, uint32(scheduleID))
		return errors.WithStack(err)
	},
}
