package spadepolicy

import (
	"github.com/urfave/cli/v2"
)

var CreateCmd = &cli.Command{
	Name:      "create",
	Usage:     "Create a SPADE policy for self deal proposal",
	ArgsUsage: "DATASET_NAME [...PROVIDER_ID]",
	Flags: []cli.Flag{
		&cli.Float64Flag{
			Name:  "min-delay",
			Usage: "Minimum delay in days for the deal start epoch",
			Value: 3.0,
		},
		&cli.Float64Flag{
			Name:  "max-delay",
			Usage: "Maximum delay in days for the deal start epoch",
			Value: 3.0,
		},
		&cli.Float64Flag{
			Name:  "min-duration",
			Usage: "Minimum duration in days for the deal start epoch",
			Value: 535.0,
		},
		&cli.Float64Flag{
			Name:  "max-duration",
			Usage: "Maximum duration in days for the deal start epoch",
			Value: 535.0,
		},
		&cli.BoolFlag{
			Name:  "verified",
			Usage: "Whether to propose dea as verified",
			Value: true,
		},
		&cli.Float64Flag{
			Name:  "price",
			Usage: "The price of the deal measured by per 32GiB over the whole duration",
			Value: 0.0,
		},
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
