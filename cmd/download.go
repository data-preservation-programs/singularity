package cmd

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/ipfs/go-log"
	"github.com/urfave/cli/v2"
)

var DownloadCmd = &cli.Command{
	Name:      "download",
	Usage:     "Download a CAR file from the metadata API",
	Category:  "Utility",
	ArgsUsage: "PIECE_CID",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "api",
			Usage: "URL of the metadata API",
			Value: "http://127.0.0.1:7777",
		},
		&cli.StringFlag{
			Name:    "out-dir",
			Usage:   "Directory to write CAR files to",
			Value:   ".",
			Aliases: []string{"o"},
		},
		&cli.IntFlag{
			Name:    "concurrency",
			Usage:   "Number of concurrent downloads",
			Value:   10,
			Aliases: []string{"j"},
		},
	},
	Action: func(c *cli.Context) error {
		api := c.String("api")
		outDir := c.String("out-dir")
		concurrency := c.Int("concurrency")
		piece := c.Args().First()
		err := handler.DownloadHandler(c.Context, piece, api, nil, outDir, concurrency)
		if err == nil {
			log.Logger("download").Info("Download complete")
			return nil
		}
		return cli.Exit(err.Error(), 1)
	},
}
