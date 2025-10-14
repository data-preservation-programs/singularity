package storage

import (
	"net/url"
	"path/filepath"
	"slices"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/gotidy/ptr"
	"github.com/rclone/rclone/fs"
	"github.com/rjNemo/underscore"
	"github.com/urfave/cli/v2"
)

var defaultClientConfig = fs.NewConfig()

var CommonConfigFlags = []cli.Flag{
	&cli.IntFlag{
		Name:        "client-retry-max",
		Usage:       "Max number of retries for IO read errors",
		DefaultText: "10",
		Category:    "Retry Strategy",
	},
	&cli.DurationFlag{
		Name:        "client-retry-delay",
		Usage:       "The initial delay before retrying IO read errors",
		DefaultText: "1s",
		Category:    "Retry Strategy",
	},
	&cli.DurationFlag{
		Name:        "client-retry-backoff",
		Usage:       "The constant delay backoff for retrying IO read errors",
		DefaultText: "1s",
		Category:    "Retry Strategy",
	},
	&cli.Float64Flag{
		Name:        "client-retry-backoff-exp",
		Usage:       "The exponential delay backoff for retrying IO read errors",
		DefaultText: "1.0",
		Category:    "Retry Strategy",
	},
	&cli.BoolFlag{
		Name:     "client-skip-inaccessible",
		Usage:    "Skip inaccessible files when opening",
		Category: "Retry Strategy",
	},
	&cli.IntFlag{
		Name:        "client-low-level-retries",
		Usage:       "Maximum number of retries for low-level client errors",
		DefaultText: "10",
		Category:    "Retry Strategy",
	},
	&cli.IntFlag{
		Name:        "client-scan-concurrency",
		Usage:       "Max number of concurrent listing requests when scanning data source",
		DefaultText: "1",
		Category:    "Client Config",
	},
}

var httpClientConfigFlags = []cli.Flag{
	&cli.DurationFlag{
		Name:        "client-connect-timeout",
		Usage:       "HTTP Client Connect timeout",
		DefaultText: defaultClientConfig.ConnectTimeout.String(),
		Category:    "Client Config",
	},
	&cli.DurationFlag{
		Name:        "client-timeout",
		Usage:       "IO idle timeout",
		DefaultText: defaultClientConfig.Timeout.String(),
		Category:    "Client Config",
	},
	&cli.DurationFlag{
		Name:        "client-expect-continue-timeout",
		Usage:       "Timeout when using expect / 100-continue in HTTP",
		DefaultText: defaultClientConfig.ExpectContinueTimeout.String(),
		Category:    "Client Config",
	},
	&cli.BoolFlag{
		Name:        "client-insecure-skip-verify",
		Usage:       "Do not verify the server SSL certificate (insecure)",
		DefaultText: "false",
		Category:    "Client Config",
	},
	&cli.BoolFlag{
		Name:        "client-no-gzip",
		Usage:       "Don't set Accept-Encoding: gzip",
		DefaultText: "false",
		Category:    "Client Config",
	},
	&cli.StringFlag{
		Name:        "client-user-agent",
		Usage:       "Set the user-agent to a specified string",
		DefaultText: defaultClientConfig.UserAgent,
		Category:    "Client Config",
	},
	&cli.PathFlag{
		Name:     "client-ca-cert",
		Usage:    "Path to CA certificate used to verify servers",
		Category: "Client Config",
	},
	&cli.PathFlag{
		Name:     "client-cert",
		Usage:    "Path to Client SSL certificate (PEM) for mutual TLS auth",
		Category: "Client Config",
	},
	&cli.PathFlag{
		Name:     "client-key",
		Usage:    "Path to Client SSL private key (PEM) for mutual TLS auth",
		Category: "Client Config",
	},
	&cli.StringSliceFlag{
		Name:     "client-header",
		Usage:    "Set HTTP header for all transactions (i.e. key=value)",
		Category: "Client Config",
	},
	&cli.BoolFlag{
		Name:     "client-use-server-mod-time",
		Usage:    "Use server modified time if possible",
		Category: "Client Config",
	},
}

const localStorageType = "local"

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
					command.Flags = append(command.Flags, httpClientConfigFlags...)
					command.Flags = append(command.Flags, CommonConfigFlags...)
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
		if backend.Prefix != localStorageType {
			command.Flags = append(command.Flags, httpClientConfigFlags...)
		}
		command.Flags = append(command.Flags, CommonConfigFlags...)
		return command
	}),
}

func createAction(c *cli.Context, storageType string, provider string) error {
	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() { _ = closer.Close() }()
	name := c.String("name")
	if name == "" {
		name = util.RandomName()
	}
	path := c.String("path")
	if storageType == localStorageType {
		path, err = filepath.Abs(path)
		if err != nil {
			return errors.Wrapf(err, "failed to get absolute path of %s", path)
		}
	}
	config := make(map[string]string)
	extraFlagNames := []string{"name", "path"}
	for _, flag := range httpClientConfigFlags {
		extraFlagNames = append(extraFlagNames, flag.Names()...)
	}
	for _, flag := range CommonConfigFlags {
		extraFlagNames = append(extraFlagNames, flag.Names()...)
	}
	for _, flagName := range c.LocalFlagNames() {
		if slices.Contains(extraFlagNames, flagName) {
			continue
		}
		if c.IsSet(flagName) {
			config[strings.ReplaceAll(flagName, "-", "_")] = c.String(flagName)
		}
	}
	clientConfig, err := getClientConfig(c)
	if err != nil {
		return errors.WithStack(err)
	}
	s, err := storage.Default.CreateStorageHandler(c.Context, db, storageType, storage.CreateRequest{
		Provider:     provider,
		Name:         name,
		Path:         path,
		Config:       config,
		ClientConfig: *clientConfig,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	cliutil.Print(c, s)
	return nil
}

func getClientConfig(c *cli.Context) (*model.ClientConfig, error) {
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
	if c.IsSet("client-scan-concurrency") {
		config.ScanConcurrency = ptr.Of(c.Int("client-scan-concurrency"))
	}
	return &config, nil
}
