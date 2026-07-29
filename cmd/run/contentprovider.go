package run

import (
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/service/contentprovider"
	"github.com/data-preservation-programs/singularity/store"
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
		&cli.IntFlag{
			Category: "HTTP IPFS Gateway",
			Name:     "ipfs-span-blocks",
			Usage:    "Blocks (~1MiB each) fetched per backend read",
			Value:    8,
		},
		&cli.IntFlag{
			Category: "HTTP IPFS Gateway",
			Name:     "ipfs-prefetch-spans",
			Usage:    "Spans prefetched ahead on sequential access",
			Value:    2,
		},
		&cli.IntFlag{
			Category: "HTTP IPFS Gateway",
			Name:     "ipfs-max-backend-reads",
			Usage:    "Max concurrent source storage reads across all requests",
			Value:    64,
		},
		&cli.IntFlag{
			Category:    "HTTP IPFS Gateway",
			Name:        "ipfs-cache-blocks",
			Usage:       "Block cache capacity in blocks (~1MiB each); 0 sizes it to max-backend-reads * span-blocks * (1 + prefetch-spans)",
			Value:       0,
			DefaultText: "derived",
		},
		&cli.DurationFlag{
			Category: "HTTP IPFS Gateway",
			Name:     "ipfs-read-timeout",
			Usage:    "Max duration of a single backend span read once it holds a connection slot",
			Value:    2 * time.Minute,
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
				IPFSSpan: store.SpanConfig{
					SpanBlocks:      c.Int("ipfs-span-blocks"),
					PrefetchSpans:   c.Int("ipfs-prefetch-spans"),
					MaxBackendReads: c.Int("ipfs-max-backend-reads"),
					CacheBlocks:     c.Int("ipfs-cache-blocks"),
					ReadTimeout:     c.Duration("ipfs-read-timeout"),
				},
				Bind: c.String("http-bind"),
			},
		}

		s, err := contentprovider.NewService(db, config)
		if err != nil {
			return errors.WithStack(err)
		}
		return s.Start(c.Context)
	},
}
