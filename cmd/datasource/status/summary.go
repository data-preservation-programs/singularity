package status

import (
	"encoding/json"
	"fmt"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource/status"
	"github.com/urfave/cli/v2"
)

var SummaryCmd = &cli.Command{
	Name:      "summary",
	Usage:     "Get the data preparation summary of a data source",
	ArgsUsage: "<source_id>",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		result, err := status.GetSourceSummaryHandler(
			db,
			c.Args().Get(0),
		)
		if err != nil {
			return err.CliError()
		}

		if c.Bool("json") {
			objJSON, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(objJSON))
			return nil
		}

		cliutil.PrintToConsole(result.Source, c.Bool("json"))
		cliutil.PrintToConsole(result.ChunkSummary, c.Bool("json"))
		cliutil.PrintToConsole(result.ItemSummary, c.Bool("json"))
		return nil
	},
}
