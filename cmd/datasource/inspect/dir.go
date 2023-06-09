package inspect

import (
	"fmt"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/urfave/cli/v2"
)

var DirCmd = &cli.Command{
	Name:        "dir",
	Usage:       "Get all item details within a directory of a data source",
	ArgsUsage:   "<source_id> <path>",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		result, err := inspect.GetDirectoryHandler(
			db,
			c.Args().Get(0),
			c.Args().Get(1),
		)
		if err != nil {
			return err.CliError()
		}

		if c.Bool("json") {
			cliutil.PrintToConsole(result, true)
			return nil
		}

		fmt.Println("Dirs:")
		cliutil.PrintToConsole(result.Dirs, false)
		fmt.Println("Items:")
		cliutil.PrintToConsole(result.Items, false)
		return nil
	},
}
