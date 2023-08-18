package run

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/contentprovider"
	"github.com/urfave/cli/v2"
)

var ContentProviderCmd = &cli.Command{
	Name:  "content-provider",
	Usage: "Start a content provider that serves retrieval requests",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Category: "HTTP Retrieval",
			Name:     "http-bind",
			Usage:    "Address to bind the HTTP server to",
			Value:    "127.0.0.1:7777",
		},
		&cli.BoolFlag{
			Category: "HTTP Retrieval",
			Name:     "enable-http",
			Usage:    "Enable HTTP retrieval",
			Value:    true,
		},
		&cli.BoolFlag{
			Category: "Bitswap Retrieval",
			Name:     "enable-bitswap",
			Usage:    "Enable bitswap retrieval",
			Value:    false,
		},
		&cli.StringFlag{
			Category:    "Bitswap Retrieval",
			Name:        "libp2p-identity-key",
			Usage:       "The base64 encoded private key for libp2p peer",
			Value:       "",
			DefaultText: "AutoGenerated",
		},
		&cli.StringSliceFlag{
			Category: "Bitswap Retrieval",
			Name:     "libp2p-listen",
			Usage:    "Addresses to listen on for libp2p connections",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		if err := model.AutoMigrate(db); err != nil {
			return errors.WithStack(err)
		}

		config := contentprovider.Config{
			HTTP: contentprovider.HTTPConfig{
				Enable: c.Bool("enable-http"),
				Bind:   c.String("http-bind"),
			},
			Bitswap: contentprovider.BitswapConfig{
				Enable:           c.Bool("enable-bitswap"),
				IdentityKey:      c.String("libp2p-identity-key"),
				ListenMultiAddrs: c.StringSlice("libp2p-listen"),
			},
		}

		s, err := contentprovider.NewService(db, config)
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}
		return s.Start(c.Context)
	},
}
