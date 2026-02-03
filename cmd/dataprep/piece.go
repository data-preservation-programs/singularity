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
	Name:  "add-piece",
	Usage: "Add a piece to a preparation. If the piece exists in the database, metadata is copied. Otherwise, --piece-size is required.",
	Description: `Add a piece to a preparation by piece CID.

If the piece CID already exists in the database (from a previous preparation),
the metadata (size, root CID, etc.) is automatically copied. This is useful for
reorganizing pieces between preparations (e.g., consolidating small pieces for
batch deal scheduling).

For external pieces not in the database, --piece-size must be provided.

NOTE: This is an advanced feature. When overriding file-path for an existing
piece, ensure the new file has matching content. File paths must be accessible
to any workers or content providers that will serve this piece.`,
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
			Name:  "piece-size",
			Usage: "Size of the piece (e.g. 32GiB). Required only for external pieces not in database.",
		},
		&cli.StringFlag{
			Name:  "file-path",
			Usage: "Path to the CAR file, used to determine the size of the file and root CID",
		},
		&cli.StringFlag{
			Name:  "root-cid",
			Usage: "Root CID of the CAR file",
		},
		&cli.Int64Flag{
			Name:  "file-size",
			Usage: "Size of the CAR file, this is required for boost online deal. If not set, it will be determined from the file path if provided.",
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
			FileSize:  c.Int64("file-size"),
		})
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, pieces)
		return nil
	},
}

var DeletePieceCmd = &cli.Command{
	Name:      "delete-piece",
	Usage:     "Delete a piece from a preparation",
	Category:  "Piece Management",
	ArgsUsage: "<preparation id|name> <piece-cid>",
	Before:    cliutil.CheckNArgs,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "delete-car",
			Usage: "Delete the physical CAR file from storage",
			Value: true,
		},
		&cli.BoolFlag{
			Name:  "force",
			Usage: "Delete even if deals reference this piece",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		return dataprep.Default.DeletePieceHandler(
			c.Context, db,
			c.Args().Get(0),
			c.Args().Get(1),
			dataprep.DeletePieceRequest{
				DeleteCar: c.Bool("delete-car"),
				Force:     c.Bool("force"),
			})
	},
}
