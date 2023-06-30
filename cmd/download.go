package cmd

import (
	datasource2 "github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/ipfs/go-log"
	"github.com/rclone/rclone/fs"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
)

var DownloadCmd = &cli.Command{
	Name:      "download",
	Usage:     "Download a CAR file from the metadata API",
	Category:  "Utility",
	ArgsUsage: "PIECE_CID",
	Flags: func() []cli.Flag {
		flags := []cli.Flag{
			&cli.StringFlag{
				Name:     "api",
				Usage:    "URL of the metadata API",
				Value:    "http://127.0.0.1:7777",
				Category: "General Options",
			},
			&cli.StringFlag{
				Name:     "out-dir",
				Usage:    "Directory to write CAR files to",
				Value:    ".",
				Aliases:  []string{"o"},
				Category: "General Options",
			},
			&cli.IntFlag{
				Name:     "concurrency",
				Usage:    "Number of concurrent downloads",
				Value:    10,
				Aliases:  []string{"j"},
				Category: "General Options",
			},
		}

		for _, r := range fs.Registry {
			if slices.Contains([]string{"crypt", "memory", "tardigrade"}, r.Prefix) {
				continue
			}
			cmd := datasource2.OptionsToCLIFlags(r)
			for _, flag := range cmd.Flags {
				stringFlag, ok := flag.(*cli.StringFlag)
				if !ok {
					flags = append(flags, flag)
					continue
				}
				stringFlag.Required = false
				stringFlag.Category = "Options for " + cmd.Name
				flags = append(flags, stringFlag)
			}
		}
		return flags
	}(),
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
