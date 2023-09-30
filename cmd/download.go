package cmd

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/cmd/storage"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/ipfs/go-log"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
)

var DownloadCmd = &cli.Command{
	Name:      "download",
	Usage:     "Download a CAR file from the metadata API",
	Category:  "Utility",
	Before:    cliutil.CheckNArgs,
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
				Category: "General Config",
			},
			&cli.IntFlag{
				Name:     "concurrency",
				Usage:    "Number of concurrent downloads",
				Value:    10,
				Category: "General Config",
			},
			&cli.BoolFlag{
				Name:     "quiet",
				Usage:    "Suppress all output",
				Category: "General Config",
				Value:    false,
			},
		}

		flags = append(flags, storage.HTTPClientConfigFlagsForUpdate...)
		flags = append(flags, storage.CommonConfigFlags...)

		keys := make(map[string]struct{})
		for _, backend := range storagesystem.Backends {
			for _, providerOptions := range backend.ProviderOptions {
				for _, option := range providerOptions.Options {
					if !model.IsSecretConfigName(option.Name) {
						continue
					}
					flag := option.ToCLIFlag(backend.Prefix+"-", false, backend.Description)
					if _, ok := keys[flag.Names()[0]]; ok {
						continue
					}
					keys[flag.Names()[0]] = struct{}{}
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
				if slices.Contains([]string{"api", "out-dir", "concurrency", "quiet"}, key) {
					continue
				}
				value := c.String(key)
				config[key] = value
			}
		}
		clientConfig, err := storage.GetClientConfigForUpdate(c)
		if err != nil {
			return errors.WithStack(err)
		}
		err = handler.DownloadHandler(c, piece, api, config, *clientConfig, outDir, concurrency)
		if err == nil {
			log.Logger("Download").Info("Download complete")
			return nil
		}
		return errors.WithStack(err)
	},
}
