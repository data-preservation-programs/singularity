package main

import "github.com/urfave/cli/v2"

var DownloadCmd = &cli.Command{
	Name:      "download",
	Usage:     "Download a CAR file from the metadata API",
	ArgsUsage: "PIECE_CID",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "metadata-api",
			Usage: "URL of the metadata API",
			Value: "http://localhost:8080",
		},
		&cli.StringFlag{
			Name:        "out-file",
			Aliases:     []string{"o"},
			Usage:       "The file to write the CAR file to",
			DefaultText: "./PIECE_CID.car",
		},
		&cli.IntFlag{
			Name:    "concurrency",
			Aliases: []string{"c"},
			Usage:   "The download concurrency",
			Value:   1,
		},
	},
}
