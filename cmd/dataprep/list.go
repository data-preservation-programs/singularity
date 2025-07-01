package dataprep

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:     "list",
	Usage:    "List all preparations",
	Category: "Preparation Management",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		preps, err := dataprep.Default.ListHandler(c.Context, db)
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, preps)
		return nil
	},
}
