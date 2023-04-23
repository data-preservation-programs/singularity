package repl

import "github.com/urfave/cli/v2"

var ListCmd = &cli.Command{
	Name:  "list",
	Usage: "List all dataset replication requests",
	Action: func(c *cli.Context) error {
		return nil
	},
}
