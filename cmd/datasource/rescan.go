package datasource

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/urfave/cli/v2"
)

var RescanCmd = &cli.Command{
	Name:        "rescan",
	Usage:       "Rescan a data source",
	ArgsUsage:   "<source_id>",
	Description: "This command will clear any error of a data source scanning work and rescan it",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		source, err := datasource.RescanSourceHandler(
			c.Context,
			db,
			c.Args().Get(0),
		)
		if err != nil {
			return err
		}

		cliutil.PrintToConsole(source, c.Bool("json"), exclude)
		return nil
	},
}
