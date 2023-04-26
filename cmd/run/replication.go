package run

import (
	"github.com/urfave/cli/v2"
)

var ReplicationCmd = &cli.Command{
	Name:  "replication",
	Usage: "Start a replication worker to process deal making",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "lotus-url",
			Category: "Lotus API",
			Usage:    "Lotus API URL",
			Value:    "https://api.node.glif.io/rpc/v0",
		},
		&cli.StringFlag{
			Name:     "lotus-token",
			Category: "Lotus API",
			Usage:    "Lotus API Token",
		},
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
