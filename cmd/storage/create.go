package storage

import (
	"path/filepath"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/rjNemo/underscore"
	"github.com/urfave/cli/v2"
)

var CreateCmd = &cli.Command{
	Name:  "create",
	Usage: "Create a new storage which can be used as source or output",
	Subcommands: underscore.Map(storagesystem.Backends, func(backend storagesystem.Backend) *cli.Command {
		if len(backend.ProviderOptions) > 1 {
			return &cli.Command{
				Name:  backend.Prefix,
				Usage: backend.Description,
				Subcommands: underscore.Map(backend.ProviderOptions, func(providerOption storagesystem.ProviderOptions) *cli.Command {
					command := providerOption.ToCLICommand(strings.ToLower(providerOption.Provider), providerOption.Provider, providerOption.ProviderDescription)
					command.Action = func(c *cli.Context) error {
						return createAction(c, backend.Prefix, providerOption.Provider)
					}
					command.Flags = append(command.Flags, &cli.StringFlag{
						Name:        "name",
						Usage:       "Name of the storage",
						DefaultText: "Auto generated",
						Category:    "General",
					}, &cli.StringFlag{
						Name:     "path",
						Usage:    "Path of the storage",
						Category: "General",
						Required: true,
					})
					return command
				}),
			}
		}
		command := backend.ProviderOptions[0].ToCLICommand(backend.Prefix, backend.Name, backend.Description)
		command.Action = func(c *cli.Context) error {
			return createAction(c, backend.Prefix, "")
		}
		command.Flags = append(command.Flags, &cli.StringFlag{
			Name:        "name",
			Usage:       "Name of the storage",
			DefaultText: "Auto generated",
			Category:    "General",
		}, &cli.StringFlag{
			Name:     "path",
			Usage:    "Path of the storage",
			Category: "General",
			Required: true,
		})
		return command
	}),
}

func createAction(c *cli.Context, storageType string, provider string) error {
	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return errors.WithStack(err)
	}
	defer closer.Close()
	name := c.String("name")
	if name == "" {
		name = gofakeit.Noun()
	}
	path := c.String("path")
	if storageType == "local" {
		path, err = filepath.Abs(path)
		if err != nil {
			return errors.Wrapf(err, "failed to get absolute path of %s", path)
		}
	}
	config := make(map[string]string)
	for _, flagName := range c.LocalFlagNames() {
		if flagName == "name" || flagName == "path" {
			continue
		}
		if c.IsSet(flagName) {
			config[strings.ReplaceAll(flagName, "-", "_")] = c.String(flagName)
		}
	}
	s, err := storage.Default.CreateStorageHandler(c.Context, db, storageType, storage.CreateRequest{
		Provider: provider,
		Name:     name,
		Path:     path,
		Config:   config,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	cliutil.Print(c, s)
	return nil
}
