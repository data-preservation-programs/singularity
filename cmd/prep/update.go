package prep

import "github.com/urfave/cli/v2"

var UpdateCmd = &cli.Command{
	Name:      "update",
	Usage:     "Update a dataset preparation request",
	ArgsUsage: "DATASET_NAME",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "min-size",
			Aliases: []string{"m"},
			Usage:   "Minimum size of the CAR files to be created",
			Value:   "20GiB",
		},
		&cli.StringFlag{
			Name:    "max-size",
			Aliases: []string{"M"},
			Usage:   "Maximum size of the CAR files to be created",
			Value:   "30GiB",
		},
		&cli.StringFlag{
			Name:        "piece-size",
			Aliases:     []string{"s"},
			Usage:       "Target piece size of the CAR files used for piece commitment calculation",
			DefaultText: "Inferred",
		},
		&cli.BoolFlag{
			Name:    "skip-error",
			Aliases: []string{"force", "f"},
			Usage:   "Skip errors during dataset preparation",
			Value:   false,
		},
		&cli.StringSliceFlag{
			Name:        "output-dir",
			Aliases:     []string{"o"},
			Usage:       "Output directory for CAR files",
			DefaultText: "inline preparation",
		},
		&cli.DurationFlag{
			Name:        "rescan-interval",
			Aliases:     []string{"interval"},
			Usage:       "Interval to rescan the dataset source for changes",
			DefaultText: "disabled",
			Value:       0,
		},
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
