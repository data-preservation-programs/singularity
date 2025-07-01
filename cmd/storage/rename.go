package storage

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/urfave/cli/v2"
)

var RenameCmd = &cli.Command{
	Name:      "rename",
	Usage:     "Rename a storage system connection",
	ArgsUsage: "<name|id> <new_name>",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		storage, err := storage.Default.RenameStorageHandler(c.Context, db, c.Args().Get(0), storage.RenameRequest{Name: c.Args().Get(1)})
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, storage)
		return nil
	},
}
