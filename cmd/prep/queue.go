package prep

import "github.com/urfave/cli/v2"

var QueueCmd = &cli.Command{
	Name:      "queue",
	Usage:     "Queue a single item for preparation",
	ArgsUsage: "DATASET_NAME ITEM_URI",
	Action: func(c *cli.Context) error {
		return nil
	},
}
