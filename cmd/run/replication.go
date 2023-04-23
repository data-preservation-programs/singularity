package run

import (
	"github.com/urfave/cli/v2"
)

var ReplicationCmd = &cli.Command{
	Name:  "replication",
	Usage: "Start a replication worker to process deal making",
	Action: func(c *cli.Context) error {
		return nil
	},
}
