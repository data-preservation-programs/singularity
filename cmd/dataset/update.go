package dataset

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/urfave/cli/v2"
)

var UpdateCmd = &cli.Command{
	Name:      "update",
	Usage:     "Update an existing dataset",
	ArgsUsage: "<dataset_name>",
	Flags: []cli.Flag{
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
		&cli.StringSliceFlag{
			Name:     "encryption-recipient",
			Usage:    "Public key of the encryption recipient",
			Category: "Encryption",
		},
		&cli.StringFlag{
			Name:     "encryption-script",
			Usage:    "EncryptionScript command to run for custom encryption",
			Category: "Encryption",
		},
	},
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		var maxSizeStr *string
		if c.IsSet("max-size") {
			s := c.String("max-size")
			maxSizeStr = &s
		}
		var pieceSizeStr *string
		if c.IsSet("piece-size") {
			s := c.String("piece-size")
			pieceSizeStr = &s
		}
		var encryptionScript *string
		if c.IsSet("encryption-script") {
			s := c.String("encryption-script")
			encryptionScript = &s
		}
		dataset, err := dataset.UpdateHandler(
			db,
			c.Args().Get(0),
			dataset.UpdateRequest{
				MaxSizeStr:           maxSizeStr,
				PieceSizeStr:         pieceSizeStr,
				OutputDirs:           c.StringSlice("output-dir"),
				EncryptionRecipients: c.StringSlice("encryption-recipients"),
				EncryptionScript:     encryptionScript,
			},
		)
		if err != nil {
			return err
		}
		cliutil.PrintToConsole(dataset, c.Bool("json"))
		return nil
	},
}
