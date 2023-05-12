package spadepolicy

import (
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:      "list",
	Usage:     "List all SPADE policies for self deal proposal",
	ArgsUsage: "DATASET_NAME",
	Action: func(c *cli.Context) error {
		return nil
	},
}
