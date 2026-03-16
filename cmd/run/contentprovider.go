package run

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/service/contentprovider"
	"github.com/urfave/cli/v2"
)

var ContentProviderCmd = &cli.Command{
	Name:  "content-provider",
	Usage: "Start a content provider that serves retrieval requests",
	Flags: []cli.Flag{
		NoAutoMigrateFlag,
		&cli.StringFlag{
			Category: "HTTP Retrieval",
			Name:     "http-bind",
			Usage:    "Address to bind the HTTP server to",
			Value:    "127.0.0.1:7777",
		},
		&cli.BoolFlag{
			Category: "HTTP Piece Retrieval",
			Name:     "enable-http-piece",
			Usage:    "Enable HTTP Piece retrieval",
			Aliases:  []string{"enable-http"},
			Value:    true,
		},
		&cli.BoolFlag{
			Category: "HTTP Piece Metadata Retrieval",
			Name:     "enable-http-piece-metadata",
			Usage:    "Enable HTTP Piece Metadata, this is to be used with the download server",
			Value:    true,
		},
		&cli.BoolFlag{
			Category: "HTTP IPFS Gateway",
			Name:     "enable-http-ipfs",
			Usage:    "Enable trustless IPFS gateway on /ipfs/",
			Value:    true,
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := openAndMigrate(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		config := contentprovider.Config{
			HTTP: contentprovider.HTTPConfig{
				EnablePiece:         c.Bool("enable-http-piece"),
				EnablePieceMetadata: c.Bool("enable-http-piece-metadata"),
				EnableIPFS:          c.Bool("enable-http-ipfs"),
				Bind:                c.String("http-bind"),
			},
		}

		s, err := contentprovider.NewService(db, config)
		if err != nil {
			return errors.WithStack(err)
		}
		return s.Start(c.Context)
	},
}
