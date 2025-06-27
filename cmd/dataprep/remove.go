package dataprep

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/urfave/cli/v2"
)

var RemoveCmd = &cli.Command{
	Name:  "remove",
	Usage: "Remove a preparation",
	Description: `This will remove all relevant information, including:
  * All related jobs
  * All related piece info
  * Mapping used for Inline Preparation
  * All File and Directory data and CIDs
  * All Schedules
This will not remove
  * All deals ever made`,
	ArgsUsage: "<name|id>",
	Before:    cliutil.CheckNArgs,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "cars",
			Usage: "Also remove prepared CAR files",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()

		removeCars := c.Bool("cars")

		err = dataprep.Default.RemovePreparationHandler(
			c.Context, db, c.Args().Get(0),
			dataprep.RemoveRequest{
				RemoveCars: removeCars,
			})

		return errors.WithStack(err)
	},
}
