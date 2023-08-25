package storage

import (
	"strings"

	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/rjNemo/underscore"
	"github.com/urfave/cli/v2"
)

var CreateCmd = &cli.Command{
	Name:  "create",
	Usage: "Create a new storage which can be used as source or output",
	Subcommands: underscore.Map(storagesystem.Backends, func(backend storagesystem.Backend) *cli.Command {
		if len(backend.ProviderOptions) > 0 {
			return &cli.Command{
				Name:  backend.Prefix,
				Usage: backend.Description,
				Subcommands: underscore.Map(backend.ProviderOptions, func(providerOption storagesystem.ProviderOptions) *cli.Command {
					return providerOption.ToCLICommand(strings.ToLower(providerOption.Provider), providerOption.Provider, providerOption.ProviderDescription)
				}),
			}
		}
		return backend.ProviderOptions[0].ToCLICommand(backend.Prefix, backend.Name, backend.Description)
	}),
}
