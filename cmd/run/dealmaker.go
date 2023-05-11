package run

import (
	"github.com/urfave/cli/v2"
)

var DealMakerCmd = &cli.Command{
	Name:  "dealmaker",
	Usage: "Start a deal making/tracking worker to process deal making",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "lotus-api",
			Category: "Lotus",
			Usage:    "Lotus RPC API endpoint, only used to get miner info",
			Value:    "https://api.node.glif.io/rpc/v1",
		},
		&cli.StringFlag{
			Name:     "lotus-token",
			Category: "Lotus",
			Usage:    "Lotus RPC API token, only used to get miner info",
			Value:    "",
		},
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
