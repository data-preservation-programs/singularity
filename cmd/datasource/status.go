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
			return errors.WithStack(err)
		}
		defer closer.Close()
		result, err := datasource.GetSourceStatusHandler(
			c.Context,
			db,
			c.Args().Get(0),
		)
		if err != nil {
			return errors.WithStack(err)
		}

		if c.Bool("json") {
			objJSON, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				return errors.WithStack(err)
			}
			fmt.Println(string(objJSON))
			return nil
		}

		fmt.Println("Pack jobs by state:")
		cliutil.PrintToConsole(result.PackJobSummary, c.Bool("json"), nil)
		fmt.Println("Files by state:")
		cliutil.PrintToConsole(result.FileSummary, c.Bool("json"), nil)
		if len(result.FailedPackJobs) > 0 {
			fmt.Println("Failed pack jobs:")
			cliutil.PrintToConsole(result.FailedPackJobs, c.Bool("json"), nil)
		}
		return nil
	},
}
