package datasource

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:  "list",
	Usage: "List all sources",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "dataset",
			Usage: "Filter by dataset name",
		},
	},
	Action: func(c *cli.Context) error {
		db, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		datasetName := c.String("dataset")
		sources, err := datasource.ListSourcesByDatasetHandler(
			db,
			datasetName,
		)
		if err != nil {
			return err
		}
		cliutil.PrintToConsole(sources, c.Bool("json"), exclude)
		return nil
	},
}
