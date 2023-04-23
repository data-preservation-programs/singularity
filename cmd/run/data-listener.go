package run

import (
	"github.com/urfave/cli/v2"
)

var DataListenerCmd = &cli.Command{
	Name:  "data-listener",
	Usage: "Start an API that listens to new data source events",
	Action: func(c *cli.Context) error {
		return nil
	},
}
