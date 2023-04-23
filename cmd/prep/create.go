package prep

import (
	"github.com/data-preservation-programs/go-singularity/prep/handler"
	"github.com/urfave/cli/v2"
	"time"
)

var CreateCmd = &cli.Command{
	Name:      "create",
	Usage:     "Create a new dataset preparation request",
	ArgsUsage: "DATASET_NAME DATASET_SOURCE [DATASET_SOURCE...]",
	Description: "DATASET_NAME must be a unique identifier for a dataset\n" +
		"DATASET_SOURCE is a list of valid dataset sources. You can also leave it empty and add it later.\n" +
		"Example: local directory  - /mnt/dataset\n" +
		"         website explorer - https://my.website.com/dataset/directory\n" +
		"         s3 bucket        - s3://my-bucket/dataset/directory\n",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "min-size",
			Aliases:  []string{"m"},
			Usage:    "Minimum size of the CAR files to be created",
			Value:    "20GiB",
			Category: "Preparation Parameters",
		},
		&cli.StringFlag{
			Name:     "max-size",
			Aliases:  []string{"M"},
			Usage:    "Maximum size of the CAR files to be created",
			Value:    "30GiB",
			Category: "Preparation Parameters",
		},
		&cli.StringFlag{
			Name:        "piece-size",
			Aliases:     []string{"s"},
			Usage:       "Target piece size of the CAR files used for piece commitment calculation",
			DefaultText: "inferred",
			Category:    "Preparation Parameters",
		},
		&cli.StringSliceFlag{
			Name:        "output-dir",
			Aliases:     []string{"o"},
			Usage:       "Output directory for CAR files",
			DefaultText: "not needed",
			Category:    "Inline Preparation",
		},
		&cli.DurationFlag{
			Name:        "rescan-interval",
			Aliases:     []string{"interval"},
			Usage:       "Interval to rescan the dataset source for changes",
			DefaultText: "disabled",
			Category:    "Live Monitoring",
			Value:       0,
		},
		&cli.DurationFlag{
			Name:     "max-wait",
			Aliases:  []string{"wait"},
			Usage:    "Maximum time to wait before queueing up the CAR generation",
			Category: "Live Monitoring",
			Value:    time.Hour,
		},
		&cli.StringSliceFlag{
			Name:     "encryption-recipient",
			Usage:    "Public key of the encryption recipient",
			Category: "Encryption",
		},
		&cli.StringFlag{
			Name:     "encryption-script",
			Usage:    "Script command to run for custom encryption",
			Category: "Encryption",
		},
	},
	Action: handler.CreateHandler,
}
