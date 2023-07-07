package dataset

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/urfave/cli/v2"
)

var ListDatasetCmd = &cli.Command{
	Name:  "list",
	Usage: "List all datasets",
	Action: func(c *cli.Context) error {
		db, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		datasets, err := dataset.ListHandler(
			db,
		)
		if err != nil {
			return err
		}
		cliutil.PrintToConsole(datasets, c.Bool("json"), nil)
		return nil
	},
}
