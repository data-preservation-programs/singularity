package spadepolicy

import (
	"github.com/urfave/cli/v2"
)

var RemoveCmd = &cli.Command{
	Name:      "remove",
	Usage:     "Remove a SPADE policy for self deal proposal",
	ArgsUsage: "DATASET_NAME POLICY_ID",
	Action: func(c *cli.Context) error {
		return nil
	},
}
