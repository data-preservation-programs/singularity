package datasource

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/urfave/cli/v2"
)

var RepackCmd = &cli.Command{
	Name:      "repack",
	Usage:     "Retry packing a chunk or all errored chunks of a data source",
	ArgsUsage: "<source_id> or --chunk-id <chunk_id>",
	Flags: []cli.Flag{
		&cli.Uint64Flag{
			Name:  "chunk-id",
			Usage: "Chunk ID to retry packing",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		var chunkID *uint64
		if c.IsSet("chunk-id") {
			c2 := c.Uint64("chunk-id")
			chunkID = &c2
		}
		chunks, err := datasource.RepackHandler(
			c.Context,
			db,
			c.Args().Get(0),
			datasource.RepackRequest{
				ChunkID: chunkID,
			},
		)
		if err != nil {
			return err
		}

		cliutil.PrintToConsole(chunks, c.Bool("json"), nil)
		return nil
	},
}
