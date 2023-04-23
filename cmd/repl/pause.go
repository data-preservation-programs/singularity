package repl

import "github.com/urfave/cli/v2"

var PauseCmd = &cli.Command{
	Name:      "pause",
	Usage:     "Pause the replication request",
	ArgsUsage: "REQUEST_ID",
	Action: func(c *cli.Context) error {
		return nil
	},
}
