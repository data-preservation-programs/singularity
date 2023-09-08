package dataprep

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/urfave/cli/v2"
)

var ListPiecesCmd = &cli.Command{
	Name:      "list-pieces",
	Usage:     "List all generated pieces for a preparation",
	Category:  "Piece Management",
	ArgsUsage: "<preparation id|name>",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		pieces, err := dataprep.Default.ListPiecesHandler(c.Context, db, c.Args().Get(0))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, pieces)
		return nil
	},
}

var AddPieceCmd = &cli.Command{
	Name:      "add-piece",
	Usage:     "Manually add piece info to a preparation. This is useful for pieces prepared by external tools.",
	Category:  "Piece Management",
	ArgsUsage: "<preparation id|name>",
	Before:    cliutil.CheckNArgs,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "piece-cid",
			Usage:    "CID of the piece",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "piece-size",
			Usage:    "Size of the piece",
			Required: true,
			Value:    "32GiB",
		},
		&cli.StringFlag{
			Name:  "file-path",
			Usage: "Path to the CAR file, used to determine the size of the file and root CID",
		},
		&cli.StringFlag{
			Name:  "root-cid",
			Usage: "Root CID of the CAR file",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		pieces, err := dataprep.Default.AddPieceHandler(c.Context, db, c.Args().Get(0), dataprep.AddPieceRequest{
			PieceCID:  c.String("piece-cid"),
			PieceSize: c.String("piece-size"),
			FilePath:  c.String("file-path"),
			RootCID:   c.String("root-cid"),
		})
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, pieces)
		return nil
	},
}
