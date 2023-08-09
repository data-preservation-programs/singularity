package inspect

import (
	"fmt"
	"strings"

	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/urfave/cli/v2"
)

var ChunksCmd = &cli.Command{
	Name:      "chunks",
	Usage:     "Get all chunk details of a data source",
	ArgsUsage: "<source_id>",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "state",
			Usage:       "Filter chunks by state. Valid values are: " + strings.Join(model.WorkStateStrings, ", "),
			Aliases:     []string{"s"},
			DefaultText: "All states",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		result, err := inspect.GetSourceChunksHandler(
			db,
			c.Args().Get(0),
			inspect.GetSourceChunksRequest{State: c.String("state")},
		)
		if err != nil {
			return err
		}

		if c.Bool("json") {
			cliutil.PrintToConsole(result, true, nil)
			return nil
		}
		fmt.Println("Chunks:")
		cliutil.PrintToConsole(result, false, []string{"PackingWorkerID"})

		return nil
	},
}
