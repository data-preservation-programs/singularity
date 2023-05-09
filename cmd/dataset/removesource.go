package dataset

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler/dataset"
	"github.com/urfave/cli/v2"
)

var RemoveSourceCmd = &cli.Command{
	Name:      "remove-source",
	Usage:     "Remove a data source from a dataset.",
	ArgsUsage: "DATASET_NAME SOURCE_PATH",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		err := dataset.RemoveSourceHandler(
			db,
			c.Args().Get(0),
			c.Args().Get(1),
		)
		return err.CliError()
	},
}
