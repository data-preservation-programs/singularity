package datasource

import (
	"encoding/json"
	"fmt"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/urfave/cli/v2"
)

var StatusCmd = &cli.Command{
	Name:      "status",
	Usage:     "Get the data preparation summary of a data source",
	ArgsUsage: "<source_id>",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		result, err := datasource.GetSourceSummaryHandler(
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

		fmt.Println("Chunks by state:")
		cliutil.PrintToConsole(result.ChunkSummary, c.Bool("json"))
		fmt.Println("Items by state:")
		cliutil.PrintToConsole(result.ItemSummary, c.Bool("json"))
		return nil
	},
}
