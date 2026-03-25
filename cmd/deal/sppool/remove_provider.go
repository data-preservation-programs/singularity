package sppool

import (
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal/sppool"
	"github.com/urfave/cli/v2"
)

var RemoveProviderCmd = &cli.Command{
	Name:        "remove-provider",
	Usage:       "Remove a provider from an SP Pool",
	Before:      cliutil.CheckNArgs,
	ArgsUsage:   "<pool_id> <provider_id>",
	Description: "Associated schedules are paused and unlinked, not deleted.",
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

		providerID, err := strconv.ParseUint(c.Args().Get(1), 10, 32)
		if err != nil {
			return errors.Wrapf(err, "failed to parse provider ID %s", c.Args().Get(1))
		}

		return errors.WithStack(sppool.Default.RemoveProviderHandler(c.Context, db, uint32(poolID), uint32(providerID)))
	},
}
