package storage

import (
	"net/url"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/gotidy/ptr"
	"github.com/rjNemo/underscore"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var ClientConfigFlagsForUpdate = []cli.Flag{
	&cli.DurationFlag{
		Name:        "client-connect-timeout",
		Usage:       "HTTP Client Connect timeout",
		DefaultText: defaultClientConfig.ConnectTimeout.String(),
		Category:    "HTTP Client Config",
	},
	&cli.DurationFlag{
		Name:        "client-timeout",
		Usage:       "IO idle timeout",
		DefaultText: defaultClientConfig.Timeout.String(),
		Category:    "HTTP Client Config",
	},
	&cli.DurationFlag{
		Name:        "client-expect-continue-timeout",
		Usage:       "Timeout when using expect / 100-continue in HTTP",
		DefaultText: defaultClientConfig.ExpectContinueTimeout.String(),
		Category:    "HTTP Client Config",
	},
	&cli.BoolFlag{
		Name:        "client-insecure-skip-verify",
		Usage:       "Do not verify the server SSL certificate (insecure)",
		DefaultText: "false",
		Category:    "HTTP Client Config",
	},
	&cli.BoolFlag{
		Name:        "client-no-gzip",
		Usage:       "Don't set Accept-Encoding: gzip",
		DefaultText: "false",
		Category:    "HTTP Client Config",
	},
	&cli.StringFlag{
		Name:        "client-user-agent",
		Usage:       "Set the user-agent to a specified string. To remove, use empty string.",
		DefaultText: defaultClientConfig.UserAgent,
		Category:    "HTTP Client Config",
	},
	&cli.PathFlag{
		Name:     "client-ca-cert",
		Usage:    "Path to CA certificate used to verify servers. To remove, use empty string.",
		Category: "HTTP Client Config",
	},
	&cli.PathFlag{
		Name:     "client-cert",
		Usage:    "Path to Client SSL certificate (PEM) for mutual TLS auth. To remove, use empty string.",
		Category: "HTTP Client Config",
	},
	&cli.PathFlag{
		Name:     "client-key",
		Usage:    "Path to Client SSL private key (PEM) for mutual TLS auth. To remove, use empty string.",
		Category: "HTTP Client Config",
	},
	&cli.StringSliceFlag{
		Name:     "client-header",
		Usage:    "Set HTTP header for all transactions (i.e. key=value). This will replace the existing header values. To remove a header, use --http-header \"key=\"\". To remove all headers, use --http-header \"\"",
		Category: "HTTP Client Config",
	},
	&cli.BoolFlag{
		Name:     "client-use-server-mod-time",
		Usage:    "Use server modified time if possible",
		Category: "HTTP Client Config",
	},
}

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
					command.ArgsUsage = "<name|id>"
					command.Before = cliutil.CheckNArgs
					command.Flags = append(command.Flags, ClientConfigFlagsForUpdate...)
					command.Flags = append(command.Flags, RetryConfigFlags...)
					return command
				}),
			}
		}
		command := backend.ProviderOptions[0].ToCLICommand(backend.Prefix, backend.Name, backend.Description)
		command.Action = func(c *cli.Context) error {
			return updateAction(c, backend.Prefix, "")
		}
		command.ArgsUsage = "<name|id>"
		command.Before = cliutil.CheckNArgs
		if backend.Prefix != "local" {
			command.Flags = append(command.Flags, ClientConfigFlagsForUpdate...)
		}
		command.Flags = append(command.Flags, RetryConfigFlags...)
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
	err = s.FindByIDOrName(db, name)
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
			config[strings.ReplaceAll(flagName, "-", "_")] = c.String(flagName)
		}
	}

	clientConfig, err := GetClientConfigForUpdate(c)
	if err != nil {
		return errors.WithStack(err)
	}
	s2, err := storage.Default.UpdateStorageHandler(c.Context, db, name, storage.UpdateRequest{
		Config:       config,
		ClientConfig: *clientConfig,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	cliutil.Print(c, s2)
	return nil
}

func GetClientConfigForUpdate(c *cli.Context) (*model.ClientConfig, error) {
	var config model.ClientConfig
	if c.IsSet("client-connect-timeout") {
		config.ConnectTimeout = ptr.Of(c.Duration("client-connect-timeout"))
	}
	if c.IsSet("client-timeout") {
		config.Timeout = ptr.Of(c.Duration("client-timeout"))
	}
	if c.IsSet("client-expect-continue-timeout") {
		config.ExpectContinueTimeout = ptr.Of(c.Duration("client-expect-continue-timeout"))
	}
	if c.IsSet("client-insecure-skip-verify") {
		config.InsecureSkipVerify = ptr.Of(c.Bool("client-insecure-skip-verify"))
	}
	if c.IsSet("client-no-gzip") {
		config.NoGzip = ptr.Of(c.Bool("client-no-gzip"))
	}
	if c.IsSet("client-user-agent") {
		config.UserAgent = ptr.Of(c.String("client-user-agent"))
	}
	if c.IsSet("client-ca-cert") {
		config.CaCert = []string{c.Path("client-ca-cert")}
	}
	if c.IsSet("client-cert") {
		config.ClientCert = ptr.Of(c.Path("client-cert"))
	}
	if c.IsSet("client-key") {
		config.ClientKey = ptr.Of(c.Path("client-key"))
	}
	if c.IsSet("client-header") {
		val := c.StringSlice("client-header")

		headers := make(map[string]string)
		for _, header := range val {
			if header == "" {
				headers = map[string]string{"": ""}
				break
			}
			kv := strings.SplitN(header, "=", 2)
			if len(kv) != 2 {
				return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid http header: %s", header)
			}
			var err error
			headers[kv[0]], err = url.QueryUnescape(kv[1])
			if err != nil {
				return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid http header: %s", header)
			}
		}
		config.Headers = headers
	}
	if c.IsSet("client-retry-max") {
		config.RetryMaxCount = ptr.Of(c.Int("client-retry-max"))
	}
	if c.IsSet("client-retry-delay") {
		config.RetryDelay = ptr.Of(c.Duration("client-retry-delay"))
	}
	if c.IsSet("client-retry-backoff") {
		config.RetryBackoff = ptr.Of(c.Duration("client-retry-backoff"))
	}
	if c.IsSet("client-retry-backoff-exp") {
		config.RetryBackoffExponential = ptr.Of(c.Float64("client-retry-backoff-exp"))
	}
	if c.IsSet("client-skip-inaccessible") {
		config.SkipInaccessibleFile = ptr.Of(c.Bool("client-skip-inaccessible"))
	}
	if c.IsSet("client-low-level-retries") {
		config.LowLevelRetries = ptr.Of(c.Int("client-low-level-retries"))
	}
	if c.IsSet("client-use-server-mod-time") {
		config.UseServerModTime = ptr.Of(c.Bool("client-use-server-mod-time"))
	}
	return &config, nil
}
