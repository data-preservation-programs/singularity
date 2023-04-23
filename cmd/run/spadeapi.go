package run

import (
	"github.com/urfave/cli/v2"
)

var SpadeAPICmd = &cli.Command{
	Name:  "spade-api",
	Usage: "Start a Spade compatible API for storage provider deal proposal self service",
	Action: func(c *cli.Context) error {
		return nil
	},
}
