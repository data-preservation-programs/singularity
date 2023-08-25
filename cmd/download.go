package cmd

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/ipfs/go-log"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
)

var DownloadCmd = &cli.Command{
	Name:      "download",
	Usage:     "Download a CAR file from the metadata API",
	Category:  "Utility",
	ArgsUsage: "<piece_cid>",
	Flags: func() []cli.Flag {
		flags := []cli.Flag{
			&cli.StringFlag{
				Name:     "api",
				Usage:    "URL of the metadata API",
				Value:    "http://127.0.0.1:7777",
				Category: "General Config",
			},
			&cli.StringFlag{
				Name:     "out-dir",
				Usage:    "Directory to write CAR files to",
				Value:    ".",
				Aliases:  []string{"o"},
				Category: "General Config",
			},
			&cli.IntFlag{
				Name:     "concurrency",
				Usage:    "Number of concurrent downloads",
				Value:    10,
				Aliases:  []string{"j"},
				Category: "General Config",
			},
		}

		keys := make(map[string]struct{})
		for _, backend := range storagesystem.Backends {
			for _, providerOptions := range backend.ProviderOptions {
				for _, option := range providerOptions.Options {
					if _, ok := keys[option.Name]; ok {
						continue
					}
					keys[option.Name] = struct{}{}
					flag := option.ToCLIFlag(backend.Prefix + "-")
					flags = append(flags, flag)
				}
			}
		}
		return flags
	}(),
	Action: func(c *cli.Context) error {
		api := c.String("api")
		outDir := c.String("out-dir")
		concurrency := c.Int("concurrency")
		piece := c.Args().First()
		config := map[string]string{}
		for _, key := range c.LocalFlagNames() {
			if c.IsSet(key) {
				if slices.Contains([]string{"api", "out-dir", "concurrency", "o", "j"}, key) {
					continue
				}
				value := c.String(key)
				config[key] = value
			}
		}
		err := handler.DownloadHandler(c.Context, piece, api, config, outDir, concurrency)
		if err == nil {
			log.Logger("download").Info("Download complete")
			return nil
		}
		return cli.Exit(err.Error(), 1)
	},
}
