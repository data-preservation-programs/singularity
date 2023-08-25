package storage

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/urfave/cli/v2"
)

var RemoveCmd = &cli.Command{
	Name:      "remove",
	Usage:     "Remove a storage connection if it's not used by any preparation",
	ArgsUsage: "<name>",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		err = storage.RemoveHandler(c.Context, db, c.Args().Get(0))
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	},
}
