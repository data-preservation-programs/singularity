package tool

import (
	"github.com/data-preservation-programs/singularity/handler/tool"
	"github.com/ipfs/go-cid"
	"github.com/pkg/errors"
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
			Name:     "output",
			Usage:    "Output directory or file to extract to. It will be created if it does not exist",
			Required: true,
			Aliases:  []string{"o"},
			Value:    ".",
		},
		&cli.StringFlag{
			Name:     "cid",
			Usage:    "CID of the folder or file to extract",
			Required: true,
			Aliases:  []string{"c"},
		},
	},
	Action: func(context *cli.Context) error {
		inputDir := context.String("input-dir")
		output := context.String("output")
		id := context.String("cid")
		c, err := cid.Decode(id)
		if err != nil {
			return errors.Wrap(err, "failed to decode CID")
		}

		return tool.ExtractCarHandler(context.Context, inputDir, output, c)
	},
}
