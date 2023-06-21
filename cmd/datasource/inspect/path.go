package inspect

import (
	"fmt"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/urfave/cli/v2"
)

var PathCmd = &cli.Command{
	Name:      "path",
	Usage:     "Get details about a path within a data source",
	ArgsUsage: "<source_id> <path>",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		result, err := inspect.InspectPathHandler(
			db,
			c.Args().Get(0),
			c.Args().Get(1),
		)
		if err != nil {
			return err.CliError()
		}

		if c.Bool("json") {
			cliutil.PrintToConsole(result, true, nil)
			return nil
		}

		fmt.Println("Current Directory:")
		cliutil.PrintToConsole(result.Current, false, nil)
		if len(result.Dirs) > 0 {
			fmt.Println("SubDirs:")
			cliutil.PrintToConsole(result.Dirs, false, nil)
		}
		if len(result.Items) > 0 {
			fmt.Println("SubItems:")
			cliutil.PrintToConsole(result.Items, false, nil)
		}
		return nil
	},
}
