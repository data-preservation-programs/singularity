package run

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/storage"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/service/downloadserver"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
)

var DownloadServerCmd = &cli.Command{
	Name:        "download-server",
	Usage:       "An HTTP server connecting to remote metadata API to offer CAR file downloads",
	Description: "Example Usage:\n  singularity run download-server --metadata-api \"http://remote-metadata-api:7777\" --bind \"127.0.0.1:8888\"",
	Flags: func() []cli.Flag {
		flags := []cli.Flag{
			&cli.StringFlag{
				Name:     "metadata-api",
				Usage:    "URL of the metadata API",
				Value:    "http://127.0.0.1:7777",
				Category: "General Config",
			},
			&cli.StringFlag{
				Name:     "bind",
				Usage:    "Address to bind the HTTP server to",
				Value:    "127.0.0.1:8888",
				Category: "General Config",
			},
		}

		flags = append(flags, storage.HTTPClientConfigFlagsForUpdate...)
		flags = append(flags, storage.CommonConfigFlags...)

		keys := make(map[string]struct{})
		for _, backend := range storagesystem.Backends {
			var providers []string
			for _, providerOptions := range backend.ProviderOptions {
				providers = append(providers, providerOptions.Provider)
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
		api := c.String("metadata-api")
		bind := c.String("bind")
		config := map[string]string{}
		for _, key := range c.LocalFlagNames() {
			if c.IsSet(key) {
				if slices.Contains([]string{"api", "bind"}, key) {
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

		server := downloadserver.NewDownloadServer(bind, api, config, *clientConfig)
		return service.StartServers(c.Context, downloadserver.Logger, server)
	},
}
