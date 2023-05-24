package datasource

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/urfave/cli/v2"
)

var RescanCmd = &cli.Command{
	Name:      "rescan",
	Usage:     "Rescan a data source",
	ArgsUsage: "<source_id>",
	Description: "This command will clear any error of a data source and rescan it",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		entries, err := datasource.RescanSourceHandler(
			db,
			c.Args().Get(0),
		)
		if err != nil {
			return err.CliError()
		}

		cliutil.PrintToConsole(entries, c.Bool("json"))
		return nil
	},
}
