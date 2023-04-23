package prep

import "github.com/urfave/cli/v2"

var ResumeCmd = &cli.Command{
	Name:      "resume",
	Usage:     "Resume the preparation of a dataset",
	ArgsUsage: "DATASET_NAME",
	Action: func(c *cli.Context) error {
		return nil
	},
}
