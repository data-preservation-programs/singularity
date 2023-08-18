package tool

import (
	"github.com/data-preservation-programs/singularity/handler/tool"
	"github.com/ipfs/go-cid"
	"github.com/cockroachdb/errors"
	"github.com/urfave/cli/v2"
)

var ExtractCarCmd = &cli.Command{
	Name:  "extract-car",
	Usage: "Extract folders or files from a folder of CAR files to a local directory",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "input-dir",
			Usage:    "Input directory containing CAR files. This directory will be scanned recursively",
			Required: true,
			Aliases:  []string{"i"},
		},
		&cli.StringFlag{
			Name:    "output",
			Usage:   "Output directory or file to extract to. It will be created if it does not exist",
			Aliases: []string{"o"},
			Value:   ".",
		},
		&cli.StringFlag{
			Name:     "cid",
			Usage:    "CID of the folder or file to extract",
			Required: true,
			Aliases:  []string{"c"},
		},
	},
	Action: func(c *cli.Context) error {
		inputDir := c.String("input-dir")
		output := c.String("output")
		id := c.String("cid")
		cidValue, err := cid.Decode(id)
		if err != nil {
			return errors.Wrap(err, "failed to decode CID")
		}

		return tool.ExtractCarHandler(c.Context, inputDir, output, cidValue)
	},
}
