package dataset

import "github.com/urfave/cli/v2"

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
		&cli.StringFlag{
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
	Action: func(c *cli.Context) error {},
}
