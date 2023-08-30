package storage

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/urfave/cli/v2"
)

var ExploreCmd = &cli.Command{
	Name:      "explore",
	Usage:     "Explore a storage by listing all entries under a path",
	ArgsUsage: "<name> [path]",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		entries, err := storage.Default.ExploreHandler(c.Context, db, c.Args().Get(0), c.Args().Get(1))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, entries)
		return nil
	},
}
