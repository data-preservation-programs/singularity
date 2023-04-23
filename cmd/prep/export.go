package prep

import "github.com/urfave/cli/v2"

var ExportCmd = &cli.Command{
	Name:      "export",
	Usage:     "Export the manifest or metadata for a dataset",
	ArgsUsage: "DATASET_NAME",
	Action: func(c *cli.Context) error {
		return nil
	},
}
