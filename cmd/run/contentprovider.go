package run

import (
	"github.com/urfave/cli/v2"
)

var ContentProviderCmd = &cli.Command{
	Name:  "content-provider",
	Usage: "Start a content provider that serves retrieval requests",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "http-piece",
			Usage: "Enable HTTP piece retrieval (for downloading CAR files)",
			Value: true,
		},
		&cli.BoolFlag{
			Name:  "http-payload",
			Usage: "Enable HTTP payload retrieval",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "bitswap",
			Usage: "Enable Bitswap retrieval",
			Value: false,
		},
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
