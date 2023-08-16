package inspect

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/urfave/cli/v2"
)

var PackingManifestsCmd = &cli.Command{
	Name:      "packingmanifests",
	Usage:     "Get all packing manifest details of a data source",
	ArgsUsage: "<source_id>",
	Flags: []cli.Flag{
		&cli.GenericFlag{
			Name:        "state",
			Value:       new(model.WorkState),
			Usage:       "Filter packing manifests by state. Valid values are: " + strings.Join(model.WorkStateStrings, ", "),
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
		sourceID, err := strconv.ParseUint(c.Args().Get(0), 10, 32)
		if err != nil {
			return err
		}
		result, err := inspect.GetSourcePackingManifestsHandler(
			db,
			uint32(sourceID),
			inspect.GetSourcePackingManifestsRequest{State: *(c.Generic("state").(*model.WorkState))},
		)
		if err != nil {
			return err
		}

		if c.Bool("json") {
			cliutil.PrintToConsole(result, true, nil)
			return nil
		}
		fmt.Println("Packing manifests:")
		cliutil.PrintToConsole(result, false, []string{"PackingWorkerID"})

		return nil
	},
}
