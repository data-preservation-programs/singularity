package dataset

import (
	"github.com/data-preservation-programs/go-singularity/cmd/cliutil"
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler/dataset"
	"github.com/urfave/cli/v2"
)

var ListPieceCmd = &cli.Command{
	Name:      "list-piece",
	Usage:     "List all pieces associated with a dataset",
	ArgsUsage: "DATASET_NAME",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)

		cars, err := dataset.ListPieceHandler(
			db, c.Args().Get(0),
		)
		if err != nil {
			return err.CliError()
		}

		cliutil.PrintToConsole(cars, c.Bool("json"))
		return nil
	},
}
