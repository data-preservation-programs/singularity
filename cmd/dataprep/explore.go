package dataprep

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/urfave/cli/v2"
)

var ExploreCmd = &cli.Command{
	Name:      "explore",
	Usage:     "Explore prepared source by path",
	ArgsUsage: "<preparation id|name> <storage id|name> [path]",
	Category:  "Preparation Management",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()

		entries, err := dataprep.Default.ExploreHandler(c.Context, db, c.Args().Get(0), c.Args().Get(1), c.Args().Get(2))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, entries)
		return nil
	},
}
