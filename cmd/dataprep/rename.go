package dataprep

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/urfave/cli/v2"
)

var RenameCmd = &cli.Command{
	Name:      "rename",
	Usage:     "Rename a preparation",
	ArgsUsage: "<name|id> <new_name>",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		preparation, err := dataprep.Default.RenamePreparationHandler(c.Context, db, c.Args().Get(0), dataprep.RenameRequest{Name: c.Args().Get(1)})
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, preparation)
		return nil
	},
}
