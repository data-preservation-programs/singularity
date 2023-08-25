package dataprep

import (
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/urfave/cli/v2"
)

var ExploreCmd = &cli.Command{
	Name:      "explore",
	Usage:     "Explore prepared source by path",
	ArgsUsage: "<preparation_id> <storage_name> [path]",
	Category:  "Preparation Management",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		id, err := strconv.ParseUint(c.Args().Get(0), 10, 32)
		if err != nil {
			return errors.Wrapf(err, "invalid preparation ID '%s'", c.Args().Get(0))
		}

		entries, err := dataprep.ExploreHandler(c.Context, db, uint32(id), c.Args().Get(1), c.Args().Get(2))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.PrintToConsole(c, entries)
		return nil
	},
}
