package dataset

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/urfave/cli/v2"
)

var CreateCmd = &cli.Command{
	Name:      "create",
	Usage:     "Create a new dataset",
	ArgsUsage: "<dataset_name>",
	Description: "<dataset_name> must be a unique identifier for a dataset\n" +
		"The dataset is a top level object to distinguish different dataset.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "max-size",
			Aliases:  []string{"M"},
			Usage:    "Maximum size of the CAR files to be created",
			Value:    "31.5GiB",
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
		&cli.StringSliceFlag{
			Name:     "encryption-recipient",
			Usage:    "[Alpha] Public key of the encryption recipient",
			Category: "Encryption",
		},
		&cli.StringFlag{
			Name:     "encryption-script",
			Usage:    "[WIP] EncryptionScript command to run for custom encryption",
			Category: "Encryption",
		},
	},
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		dataset, err := dataset.CreateHandler(
			db,
			dataset.CreateRequest{
				Name:                 c.Args().Get(0),
				MaxSizeStr:           c.String("max-size"),
				PieceSizeStr:         c.String("piece-size"),
				OutputDirs:           c.StringSlice("output-dir"),
				EncryptionRecipients: c.StringSlice("encryption-recipient"),
				EncryptionScript:     c.String("encryption-script")},
		)
		if err != nil {
			return err
		}
		cliutil.PrintToConsole(dataset, c.Bool("json"))
		return nil
	},
}
