package inspect

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/urfave/cli/v2"
)

var FilesCmd = &cli.Command{
	Name:        "files",
	Usage:       "Get all file details of a data source",
	ArgsUsage:   "<source_id>",
	Description: "This command will list all files in a data source. This may be very large list.",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		result, err := inspect.GetSourceFilesHandler(
			c.Context,
			db,
			c.Args().Get(0),
		)
		if err != nil {
			return err
		}

		cliutil.PrintToConsole(result, c.Bool("json"), nil)
		return nil
	},
}
