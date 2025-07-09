package cliutil

import (
	"time"

	"github.com/urfave/cli/v2"
)

// CommonDealFlags contains reusable deal configuration flags for CLI commands
var CommonDealFlags = []cli.Flag{
	&cli.Float64Flag{
		Name:     "deal-price-per-gb",
		Usage:    "Price in FIL per GiB for storage deals",
		Value:    0.0,
		Category: "Deal Settings",
	},
	&cli.Float64Flag{
		Name:     "deal-price-per-gb-epoch",
		Usage:    "Price in FIL per GiB per epoch for storage deals",
		Value:    0.0,
		Category: "Deal Settings",
	},
	&cli.Float64Flag{
		Name:     "deal-price-per-deal",
		Usage:    "Price in FIL per deal for storage deals",
		Value:    0.0,
		Category: "Deal Settings",
	},
	&cli.DurationFlag{
		Name:     "deal-duration",
		Usage:    "Duration for storage deals (e.g., 535 days)",
		Value:    12840 * time.Hour, // ~535 days
		Category: "Deal Settings",
	},
	&cli.DurationFlag{
		Name:     "deal-start-delay",
		Usage:    "Start delay for storage deals (e.g., 72h)",
		Value:    72 * time.Hour,
		Category: "Deal Settings",
	},
	&cli.BoolFlag{
		Name:     "deal-verified",
		Usage:    "Whether deals should be verified",
		Category: "Deal Settings",
	},
	&cli.BoolFlag{
		Name:     "deal-keep-unsealed",
		Usage:    "Whether to keep unsealed copy of deals",
		Category: "Deal Settings",
	},
	&cli.BoolFlag{
		Name:     "deal-announce-to-ipni",
		Usage:    "Whether to announce deals to IPNI",
		Value:    true,
		Category: "Deal Settings",
	},
	&cli.StringFlag{
		Name:     "deal-provider",
		Usage:    "Storage Provider ID for deals (e.g., f01000)",
		Category: "Deal Settings",
	},
	&cli.StringFlag{
		Name:     "deal-url-template",
		Usage:    "URL template for deals",
		Category: "Deal Settings",
	},
	&cli.StringFlag{
		Name:     "deal-http-headers",
		Usage:    "HTTP headers for deals in JSON format",
		Category: "Deal Settings",
	},
	&cli.StringFlag{
		Name:     "deal-template",
		Usage:    "Name or ID of deal template to use for defaults",
		Category: "Deal Settings",
	},
}

// CommonStorageClientFlags contains reusable storage client configuration flags
var CommonStorageClientFlags = []cli.Flag{
	&cli.IntFlag{
		Name:        "client-retry-max",
		Usage:       "Max number of retries for IO read errors",
		Value:       10,
		Category:    "Storage Client Config",
	},
	&cli.DurationFlag{
		Name:        "client-retry-delay",
		Usage:       "The initial delay before retrying IO read errors",
		Value:       time.Second,
		Category:    "Storage Client Config",
	},
	&cli.DurationFlag{
		Name:        "client-retry-backoff",
		Usage:       "The constant delay backoff for retrying IO read errors",
		Value:       time.Second,
		Category:    "Storage Client Config",
	},
	&cli.Float64Flag{
		Name:        "client-retry-backoff-exp",
		Usage:       "The exponential delay backoff for retrying IO read errors",
		Value:       1.0,
		Category:    "Storage Client Config",
	},
	&cli.BoolFlag{
		Name:     "client-skip-inaccessible",
		Usage:    "Skip inaccessible files when opening",
		Category: "Storage Client Config",
	},
	&cli.IntFlag{
		Name:        "client-low-level-retries",
		Usage:       "Maximum number of retries for low-level client errors",
		Value:       10,
		Category:    "Storage Client Config",
	},
	&cli.IntFlag{
		Name:        "client-scan-concurrency",
		Usage:       "Max number of concurrent listing requests when scanning data source",
		Value:       1,
		Category:    "Storage Client Config",
	},
	&cli.DurationFlag{
		Name:        "client-connect-timeout",
		Usage:       "HTTP Client Connect timeout",
		Category:    "Storage Client Config",
	},
	&cli.DurationFlag{
		Name:        "client-timeout",
		Usage:       "IO idle timeout",
		Category:    "Storage Client Config",
	},
	&cli.DurationFlag{
		Name:        "client-expect-continue-timeout",
		Usage:       "Timeout when using expect / 100-continue in HTTP",
		Category:    "Storage Client Config",
	},
	&cli.BoolFlag{
		Name:        "client-insecure-skip-verify",
		Usage:       "Do not verify the server SSL certificate (insecure)",
		Category:    "Storage Client Config",
	},
	&cli.BoolFlag{
		Name:        "client-no-gzip",
		Usage:       "Don't set Accept-Encoding: gzip",
		Category:    "Storage Client Config",
	},
	&cli.StringFlag{
		Name:        "client-user-agent",
		Usage:       "Set the user-agent to a specified string",
		Category:    "Storage Client Config",
	},
	&cli.PathFlag{
		Name:     "client-ca-cert",
		Usage:    "Path to CA certificate used to verify servers",
		Category: "Storage Client Config",
	},
	&cli.PathFlag{
		Name:     "client-cert",
		Usage:    "Path to Client SSL certificate (PEM) for mutual TLS auth",
		Category: "Storage Client Config",
	},
	&cli.PathFlag{
		Name:     "client-key",
		Usage:    "Path to Client SSL private key (PEM) for mutual TLS auth",
		Category: "Storage Client Config",
	},
	&cli.StringSliceFlag{
		Name:     "client-header",
		Usage:    "Set HTTP header for all transactions (i.e. key=value)",
		Category: "Storage Client Config",
	},
	&cli.BoolFlag{
		Name:     "client-use-server-mod-time",
		Usage:    "Use server modified time if possible",
		Category: "Storage Client Config",
	},
}