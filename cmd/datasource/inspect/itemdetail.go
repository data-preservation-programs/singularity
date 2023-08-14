package inspect

import (
	"fmt"

	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/urfave/cli/v2"
)

var FileDetailCmd = &cli.Command{
	Name:      "filedetail",
	Usage:     "Get details about a specific file",
	ArgsUsage: "<file_id>",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		result, err := inspect.GetSourceFileDetailHandler(
			db,
			c.Args().Get(0),
		)
		if err != nil {
			return err
		}

		if c.Bool("json") {
			cliutil.PrintToConsole(result, true, nil)
			return nil
		}

		fmt.Println("File:")
		cliutil.PrintToConsole(result, false, nil)
		fmt.Println("File Parts:")
		cliutil.PrintToConsole(result.FileRanges, false, nil)
		return nil
	},
}
