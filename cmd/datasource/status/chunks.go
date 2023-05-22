package status

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource/status"
	"github.com/urfave/cli/v2"
)

var ChunksCmd = &cli.Command{
	Name:      "chunks",
	Usage:     "Get all chunk details of a data source",
	ArgsUsage: "<source_id>",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		result, err := status.GetSourceChunksHandler(
			db,
			c.Args().Get(0),
		)
		if err != nil {
			return err.CliError()
		}

		cliutil.PrintToConsole(result, c.Bool("json"))
		return nil
	},
}
