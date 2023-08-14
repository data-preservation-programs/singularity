package inspect

import (
	"fmt"

	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/urfave/cli/v2"
)

var ItemDetailCmd = &cli.Command{
	Name:      "itemdetail",
	Usage:     "Get details about a specific item",
	ArgsUsage: "<item_id>",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		result, err := inspect.GetSourceItemDetailHandler(
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

		fmt.Println("Item:")
		cliutil.PrintToConsole(result, false, nil)
		fmt.Println("Item Parts:")
		cliutil.PrintToConsole(result.FileRanges, false, nil)
		return nil
	},
}
