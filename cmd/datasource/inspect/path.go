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
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		result, err := inspect.GetPathHandler(
			c.Context,
			db,
			c.Args().Get(0),
			inspect.GetPathRequest{
				Path: c.Args().Get(1),
			},
		)
		if err != nil {
			return err
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
		if len(result.Files) > 0 {
			fmt.Println("SubFiles:")
			cliutil.PrintToConsole(result.Files, false, nil)
		}
		return nil
	},
}
