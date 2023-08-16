package datasource

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/urfave/cli/v2"
)

var RepackCmd = &cli.Command{
	Name:      "repack",
	Usage:     "Retry packing a packingmanifest or all errored packingmanifests of a data source",
	ArgsUsage: "<source_id> or --packing-manifest-id <packing_manifest_id>",
	Flags: []cli.Flag{
		&cli.Uint64Flag{
			Name:  "packing-manifest-id",
			Usage: "Packing manifest ID to retry packing",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		var packingManifestID *uint64
		if c.IsSet("packing-manifest-id") {
			c2 := c.Uint64("packing-manifest-id")
			packingManifestID = &c2
		}
		packingManifests, err := datasource.RepackHandler(
			db,
			c.Args().Get(0),
			datasource.RepackRequest{
				PackingManifestID: packingManifestID,
			},
		)
		if err != nil {
			return err
		}

		cliutil.PrintToConsole(packingManifests, c.Bool("json"), nil)
		return nil
	},
}
