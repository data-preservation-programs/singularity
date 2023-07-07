package inspect

import (
	"fmt"

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
	Action: func(c *cli.Context) error {
		db, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		result, err := inspect.GetSourceChunksHandler(
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
		fmt.Println("Chunks:")
		cliutil.PrintToConsole(result, false, []string{"PackingWorkerID"})
		fmt.Println("Pieces:")
		var cars []model.Car
		for _, chunk := range result {
			cars = append(cars, chunk.Cars...)
		}
		cliutil.PrintToConsole(cars, false, nil)

		return nil
	},
}
