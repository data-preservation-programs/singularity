package repl

import "github.com/urfave/cli/v2"

var ResumeCmd = &cli.Command{
	Name:      "resume",
	Usage:     "Resume the replication request",
	ArgsUsage: "REQUEST_ID",
	Action: func(c *cli.Context) error {
		return nil
	},
}
