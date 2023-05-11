package dataset

import (
	"github.com/data-preservation-programs/go-singularity/cmd/cliutil"
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler/dataset"
	"github.com/urfave/cli/v2"
)

var AddPieceCmd = &cli.Command{
	Name:      "add-piece",
	Usage:     "Manually register a piece (CAR file) with the dataset for deal making purpose",
	ArgsUsage: "DATASET_ID PIECE_CID PIECE_SIZE",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "file-path",
			Usage:   "Path to the CAR file, used to determine the size of the file and root CID",
			Aliases: []string{"p"},
		},
		&cli.Uint64Flag{
			Name:    "file-size",
			Usage:   "Size of the CAR file, if not provided, will be determined by the CAR file",
			Aliases: []string{"s"},
		},
		&cli.StringFlag{
			Name:    "root-cid",
			Usage:   "Root CID of the CAR file, if not provided, will be determined by the CAR file header. Used to populate the label field of storage deal",
			Aliases: []string{"r"},
		},
	},
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)

		car, err := dataset.AddPieceHandler(
			db, c.Args().Get(0), dataset.AddPieceRequest{
				PieceCID:  c.Args().Get(1),
				PieceSize: c.Args().Get(2),
				FilePath:  c.String("file-path"),
				FileSize:  c.Uint64("file-size"),
				RootCID:   c.String("root-cid"),
			},
		)
		if err != nil {
			return err.CliError()
		}

		cliutil.PrintToConsole(car, c.Bool("json"))
		return nil
	},
}
