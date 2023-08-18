package dataset

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/urfave/cli/v2"
)

var RemoveDatasetCmd = &cli.Command{
	Name:        "remove",
	Usage:       "Remove a specific dataset. This will not remove the CAR files.",
	Description: "Important! If the dataset is large, this command will take some time to remove all relevant data.",
	ArgsUsage:   "<dataset_name>",
	Flags: []cli.Flag{
		cliutil.ReallyDotItFlag,
	},
	Action: func(c *cli.Context) error {
		if err := cliutil.HandleReallyDoIt(c); err != nil {
			return errors.WithStack(err)
		}
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		return dataset.RemoveHandler(
			c.Context,
			db,
			c.Args().Get(0),
		)
	},
}
