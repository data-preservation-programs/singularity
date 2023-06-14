package inspect

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/urfave/cli/v2"
)

var DagsCmd = &cli.Command{
	Name:      "dags",
	Usage:     "Get all piece details for generated dags",
	ArgsUsage: "<source_id>",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		cars, err := inspect.GetDagsHandler(
			db,
			c.Args().Get(0),
		)
		if err != nil {
			return err.CliError()
		}

		cliutil.PrintToConsole(cars, c.Bool("json"))

		return nil
	},
}
