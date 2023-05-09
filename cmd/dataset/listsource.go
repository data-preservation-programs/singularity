package dataset

import (
	"github.com/data-preservation-programs/go-singularity/cmd/cliutil"
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler/dataset"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/urfave/cli/v2"
)

var ListSourceCmd = &cli.Command{
	Name:      "list-source",
	Usage:     "List all sources of a dataset.",
	ArgsUsage: "DATASET_NAME",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		if err := model.Init(c.String("password"), db); err != nil {
			return cli.Exit("Cannot initialize encryption"+err.Error(), 1)
		}
		sources, err := dataset.ListSourceHandler(
			db,
			c.Args().Get(0),
		)
		if err != nil {
			return err.CliError()
		}
		cliutil.PrintToConsole(sources, c.Bool("json"))
		return nil
	},
}
