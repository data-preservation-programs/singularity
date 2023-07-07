package dataset

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/urfave/cli/v2"
)

var AddPieceCmd = &cli.Command{
	Name:  "add-piece",
	Usage: "Manually register a piece (CAR file) with the dataset for deal making purpose",
	Description: "If you already have the CAR file:\n" +
		"  singularity dataset add-piece -p <path_to_car_file> <dataset_name> <piece_cid> <piece_size>\n\n" +
		"If you don't have the CAR file but you know the RootCID:\n" +
		"  singularity dataset add-piece -r <root_cid> <dataset_name> <piece_cid> <piece_size>\n\n" +
		"If you don't have either:\n" +
		"  singularity dataset add-piece -r <root_cid> <dataset_name> <piece_cid> <piece_size>\n" +
		"However in this case, deals made will not have rootCID set correctly so it may not work well with retrieval testing.",
	ArgsUsage: "<dataset_name> <piece_cid> <piece_size>",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "file-path",
			Usage:   "Path to the CAR file, used to determine the size of the file and root CID",
			Aliases: []string{"p"},
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
				RootCID:   c.String("root-cid"),
			},
		)
		if err != nil {
			return err
		}

		cliutil.PrintToConsole(car, c.Bool("json"), nil)
		return nil
	},
}
