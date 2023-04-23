package prep

import "github.com/urfave/cli/v2"

var StatusCmd = &cli.Command{
	Name:      "status",
	Usage:     "Check the preparation status of a dataset",
	ArgsUsage: "DATASET_NAME",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "watch",
			Usage:   "Watch the progress of the preparation",
			Aliases: []string{"w"},
			Value:   false,
		},
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
