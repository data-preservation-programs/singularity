package dataset

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler/dataset"
	"github.com/urfave/cli/v2"
)

var RemoveDatasetCmd = &cli.Command{
	Name:      "remove",
	Usage:     "Remove a specific dataset. This will not remove the CAR files.",
	Description: "Important! If the dataset is large, this command will take some time to remove all relevant data.",
	ArgsUsage: "DATASET_NAME",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		err := dataset.RemoveHandler(
			db,
			c.Args().Get(0),
		)
		if err != nil {
			return err.CliError()
		}
		return nil
	},
}
