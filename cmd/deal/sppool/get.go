package sppool

import (
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal/sppool"
	"github.com/urfave/cli/v2"
)

var GetCmd = &cli.Command{
	Name:      "get",
	Usage:     "Get an SP Pool by ID",
	Before:    cliutil.CheckNArgs,
	ArgsUsage: "<pool_id>",
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

		pool, err := sppool.Default.GetHandler(c.Context, db, uint32(poolID))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, pool)
		return nil
	},
}
