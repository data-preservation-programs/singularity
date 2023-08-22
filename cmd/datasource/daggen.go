package datasource

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/urfave/cli/v2"
)

var DagGenCmd = &cli.Command{
	Name:      "daggen",
	Usage:     "Generate and export the DAG which represents the full folder structure of the data source",
	ArgsUsage: "<source_id>",
	Description: "This step is required for:\n" +
		"  1. Lookup and download files or folders using unixfs path\n" +
		"  2. Retrieve file that are splited across multiple pieces/deals",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		source, err := datasource.DagGenHandler(c.Context, db, c.Args().Get(0))
		if err != nil {
			return err
		}
		cliutil.PrintToConsole(source, c.Bool("json"), exclude)
		return nil
	},
}
