package schedule

import (
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/urfave/cli/v2"
)

var PauseCmd = &cli.Command{
	Name:      "pause",
	Usage:     "Pause a specific schedule",
	Before:    cliutil.CheckNArgs,
	ArgsUsage: "<schedule_id>",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()

		scheduleID, err := strconv.ParseUint(c.Args().Get(0), 10, 32)
		if err != nil {
			return errors.Wrapf(err, "failed to parse schedule ID %s", c.Args().Get(0))
		}

		schedule, err := schedule.Default.PauseHandler(c.Context, db, uint32(scheduleID))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, schedule)
		return nil
	},
}
