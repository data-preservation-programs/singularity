package storage

import (
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/rjNemo/underscore"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var UpdateCmd = &cli.Command{
	Name:  "update",
	Usage: "Update the configuration of an existing storage connection",
	Subcommands: underscore.Map(storagesystem.Backends, func(backend storagesystem.Backend) *cli.Command {
		if len(backend.ProviderOptions) > 1 {
			return &cli.Command{
				Name:  backend.Prefix,
				Usage: backend.Description,
				Subcommands: underscore.Map(backend.ProviderOptions, func(providerOption storagesystem.ProviderOptions) *cli.Command {
					command := providerOption.ToCLICommand(strings.ToLower(providerOption.Provider), providerOption.Provider, providerOption.ProviderDescription)
					command.Action = func(c *cli.Context) error {
						return updateAction(c, backend.Prefix, providerOption.Provider)
					}
					command.ArgsUsage = "<name>"
					command.Before = cliutil.CheckNArgs
					return command
				}),
			}
		}
		command := backend.ProviderOptions[0].ToCLICommand(backend.Prefix, backend.Name, backend.Description)
		command.Action = func(c *cli.Context) error {
			return updateAction(c, backend.Prefix, "")
		}
		command.ArgsUsage = "<name>"
		command.Before = cliutil.CheckNArgs
		return command
	}),
}

func updateAction(c *cli.Context, storageType string, provider string) error {
	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return errors.WithStack(err)
	}
	defer closer.Close()
	name := c.Args().Get(0)

	var s model.Storage
	err = db.WithContext(c.Context).Where("name = ?", name).First(&s).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrapf(handlererror.ErrNotFound, "storage %s does not exist", name)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if s.Type != storageType {
		return errors.Wrapf(handlererror.ErrInvalidParameter, "storage %s is not of type %s", name, storageType)
	}

	if s.Config != nil && s.Config["provider"] != provider {
		return errors.Wrapf(handlererror.ErrInvalidParameter, "storage %s is not of provider %s", name, provider)
	}

	config := make(map[string]string)
	for _, flagName := range c.LocalFlagNames() {
		if c.IsSet(flagName) {
			config[flagName] = c.String(flagName)
		}
	}

	s2, err := storage.UpdateStorageHandler(c.Context, db, name, config)
	if err != nil {
		return errors.WithStack(err)
	}

	cliutil.PrintToConsole(c, s2)
	return nil
}
