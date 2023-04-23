package run

import (
	"github.com/urfave/cli/v2"
)

var ContentProviderCmd = &cli.Command{
	Name:  "content-provider",
	Usage: "Start a content provider that serves retrieval requests",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "http-piece",
			Usage:   "Enable HTTP piece retrieval (for downloading CAR files)",
			EnvVars: []string{"CONTENT_PROVIDER_HTTP_PIECE"},
			Value:   true,
		},
		&cli.BoolFlag{
			Name:    "http-payload",
			Usage:   "Enable HTTP payload retrieval",
			EnvVars: []string{"CONTENT_PROVIDER_HTTP_PAYLOAD"},
			Value:   false,
		},
		&cli.BoolFlag{
			Name:    "bitswap",
			Usage:   "Enable Bitswap retrieval",
			EnvVars: []string{"CONTENT_PROVIDER_BITSWAP"},
			Value:   false,
		},
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
