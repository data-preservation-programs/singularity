package prep

import "github.com/urfave/cli/v2"

var RemoveCmd = &cli.Command{
	Name:      "remove",
	Usage:     "Remove all relevant data of a dataset",
	ArgsUsage: "DATASET_NAME",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "--purge",
			Aliases: []string{"p"},
			Usage:   "Also delete all exported CAR files",
		},
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
