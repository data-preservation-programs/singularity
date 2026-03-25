package sppool

import (
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal/sppool"
	"github.com/urfave/cli/v2"
)

var RemoveCmd = &cli.Command{
	Name:        "remove",
	Usage:       "Remove an SP Pool",
	Before:      cliutil.CheckNArgs,
	ArgsUsage:   "<pool_id>",
	Description: "Removes the pool. Generated schedules are preserved but unlinked.",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		poolID, err := strconv.ParseUint(c.Args().Get(0), 10, 32)
		if err != nil {
			return errors.Wrapf(err, "failed to parse pool ID %s", c.Args().Get(0))
		}

		return errors.WithStack(sppool.Default.RemoveHandler(c.Context, db, uint32(poolID)))
	},
}
