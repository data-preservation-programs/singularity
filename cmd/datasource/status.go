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
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		result, err := datasource.GetSourceStatusHandler(
			c.Context,
			db,
			c.Args().Get(0),
		)
		if err != nil {
			return err
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
		cliutil.PrintToConsole(result.ChunkSummary, c.Bool("json"), nil)
		fmt.Println("Items by state:")
		cliutil.PrintToConsole(result.ItemSummary, c.Bool("json"), nil)
		if len(result.FailedChunks) > 0 {
			fmt.Println("Failed chunks:")
			cliutil.PrintToConsole(result.FailedChunks, c.Bool("json"), nil)
		}
		return nil
	},
}
