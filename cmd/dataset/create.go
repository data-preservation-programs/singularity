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
			Usage:    "Public key of the encryption recipient",
			Category: "Encryption",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		dataset, err := dataset.CreateHandler(
			c.Context,
			db,
			dataset.CreateRequest{
				Name:                 c.Args().Get(0),
				MaxSizeStr:           c.String("max-size"),
				PieceSizeStr:         c.String("piece-size"),
				OutputDirs:           c.StringSlice("output-dir"),
				EncryptionRecipients: c.StringSlice("encryption-recipient"),
			},
		)
		if err != nil {
			return err
		}
		cliutil.PrintToConsole(dataset, c.Bool("json"), nil)
		return nil
	},
}
