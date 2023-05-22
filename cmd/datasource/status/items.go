package status

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource/status"
	"github.com/urfave/cli/v2"
)

var ItemsCmd = &cli.Command{
	Name:        "items",
	Usage:       "Get all item details of a data source",
	ArgsUsage:   "<source_id>",
	Description: "This command will list all items in a data source. This may be very large if not filtered by chunk id.",
	Flags: []cli.Flag{
		&cli.UintFlag{
			Name:  "chunk-id",
			Usage: "Filter by chunk ID",
		},
	},
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		result, err := status.GetSourceItemsHandler(
			db,
			c.Args().Get(0),
			c.String("chunk-id"),
		)
		if err != nil {
			return err.CliError()
		}

		cliutil.PrintToConsole(result, c.Bool("json"))
		return nil
	},
}
