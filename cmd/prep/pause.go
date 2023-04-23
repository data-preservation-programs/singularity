package prep

import "github.com/urfave/cli/v2"

var PauseCmd = &cli.Command{
	Name:      "pause",
	Usage:     "Pause the preparation of a dataset",
	ArgsUsage: "DATASET_NAME",
	Action: func(c *cli.Context) error {
		return nil
	},
}
