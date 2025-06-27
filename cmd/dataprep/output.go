package dataprep

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/urfave/cli/v2"
)

var AttachOutputCmd = &cli.Command{
	Name:      "attach-output",
	Usage:     "Attach a output storage to a preparation",
	ArgsUsage: "<preparation id|name> <storage id|name>",
	Category:  "Preparation Management",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		prep, err := dataprep.Default.AddOutputStorageHandler(c.Context, db, c.Args().Get(0), c.Args().Get(1))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, prep)
		return nil
	},
}

var DetachOutputCmd = &cli.Command{
	Name:      "detach-output",
	Usage:     "Detach a output storage to a preparation",
	ArgsUsage: "<preparation id|name> <storage id|name>",
	Category:  "Preparation Management",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		prep, err := dataprep.Default.RemoveOutputStorageHandler(c.Context, db, c.Args().Get(0), c.Args().Get(1))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, prep)
		return nil
	},
}
