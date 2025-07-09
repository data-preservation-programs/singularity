package dealtemplate

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dealtemplate"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/urfave/cli/v2"
)

var CreateCmd = &cli.Command{
	Name:     "create",
	Usage:    "Create a new deal template",
	Category: "Deal Template Management",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Usage:    "Name of the deal template",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "description",
			Usage: "Description of the deal template",
		},
		&cli.Float64Flag{
			Name:  "deal-price-per-gb",
			Usage: "Price in FIL per GiB for storage deals",
			Value: 0.0,
		},
		&cli.Float64Flag{
			Name:  "deal-price-per-gb-epoch",
			Usage: "Price in FIL per GiB per epoch for storage deals",
			Value: 0.0,
		},
		&cli.Float64Flag{
			Name:  "deal-price-per-deal",
			Usage: "Price in FIL per deal for storage deals",
			Value: 0.0,
		},
		&cli.DurationFlag{
			Name:  "deal-duration",
			Usage: "Duration for storage deals (e.g., 535 days)",
			Value: 0,
		},
		&cli.DurationFlag{
			Name:  "deal-start-delay",
			Usage: "Start delay for storage deals (e.g., 72h)",
			Value: 0,
		},
		&cli.BoolFlag{
			Name:  "deal-verified",
			Usage: "Whether deals should be verified",
		},
		&cli.BoolFlag{
			Name:  "deal-keep-unsealed",
			Usage: "Whether to keep unsealed copy of deals",
		},
		&cli.BoolFlag{
			Name:  "deal-announce-to-ipni",
			Usage: "Whether to announce deals to IPNI",
		},
		&cli.StringFlag{
			Name:  "deal-provider",
			Usage: "Storage Provider ID for deals (e.g., f01000)",
		},
		&cli.StringFlag{
			Name:  "deal-url-template",
			Usage: "URL template for deals",
		},
		&cli.StringFlag{
			Name:  "deal-http-headers",
			Usage: "HTTP headers for deals in JSON format",
		},
		&cli.StringFlag{
			Name:  "notes",
			Usage: "Notes or tags for tracking purposes",
		},
		&cli.BoolFlag{
			Name:  "force",
			Usage: "Force deals regardless of replication restrictions (overrides max pending/total deal limits and piece CID restrictions)",
		},
		&cli.StringSliceFlag{
			Name:  "allowed-piece-cid",
			Usage: "List of allowed piece CIDs for this template",
		},
		&cli.StringFlag{
			Name:  "allowed-piece-cid-file",
			Usage: "File containing list of allowed piece CIDs",
		},

		// Scheduling flags (matching deal schedule create command)
		&cli.StringFlag{
			Name:     "schedule-cron",
			Usage:    "Cron schedule to send out batch deals (e.g., @daily, @hourly, '0 0 * * *')",
			Category: "Scheduling",
		},
		&cli.IntFlag{
			Name:     "schedule-deal-number",
			Usage:    "Max deal number per triggered schedule (0 = unlimited)",
			Category: "Scheduling",
		},
		&cli.StringFlag{
			Name:     "schedule-deal-size",
			Usage:    "Max deal sizes per triggered schedule (e.g., 500GiB, 0 = unlimited)",
			Category: "Scheduling",
			Value:    "0",
		},

		// Restriction flags (matching deal schedule create command)
		&cli.IntFlag{
			Name:     "total-deal-number",
			Usage:    "Max total deal number for this template (0 = unlimited)",
			Category: "Restrictions",
		},
		&cli.StringFlag{
			Name:     "total-deal-size",
			Usage:    "Max total deal sizes for this template (e.g., 100TiB, 0 = unlimited)",
			Category: "Restrictions",
			Value:    "0",
		},
		&cli.IntFlag{
			Name:     "max-pending-deal-number",
			Usage:    "Max pending deal number overall (0 = unlimited)",
			Category: "Restrictions",
		},
		&cli.StringFlag{
			Name:     "max-pending-deal-size",
			Usage:    "Max pending deal sizes overall (e.g., 1000GiB, 0 = unlimited)",
			Category: "Restrictions",
			Value:    "0",
		},

		// HTTP headers as string slice (matching deal schedule create command)
		&cli.StringSliceFlag{
			Name:     "http-header",
			Usage:    "HTTP headers to be passed with the request (key=value format)",
			Category: "Boost Only",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		db = db.WithContext(c.Context)

		// Validate inputs
		if err := validateCreateTemplateInputs(c); err != nil {
			return errors.Wrap(err, "validation failed")
		}

		// Parse deal HTTP headers if provided
		var dealHTTPHeaders model.ConfigMap
		if headersStr := c.String("deal-http-headers"); headersStr != "" {
			var tempMap map[string]string
			if err := json.Unmarshal([]byte(headersStr), &tempMap); err != nil {
				return errors.Wrapf(err, "invalid JSON format for deal-http-headers: %s", headersStr)
			}
			dealHTTPHeaders = model.ConfigMap(tempMap)
		}

		// Parse allowed piece CIDs from flags and file
		var allowedPieceCIDs model.StringSlice

		// Add piece CIDs from flag
		if flagCIDs := c.StringSlice("allowed-piece-cid"); len(flagCIDs) > 0 {
			allowedPieceCIDs = append(allowedPieceCIDs, flagCIDs...)
		}

		// Add piece CIDs from file
		if filePath := c.String("allowed-piece-cid-file"); filePath != "" {
			cleanPath := filepath.Clean(filePath)
			fileContent, err := os.ReadFile(cleanPath)
			if err != nil {
				return errors.Wrapf(err, "failed to read allowed-piece-cid-file: %s", filePath)
			}

			// Split by lines and filter out empty ones
			lines := strings.Split(string(fileContent), "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line != "" && !strings.HasPrefix(line, "#") { // Skip empty lines and comments
					allowedPieceCIDs = append(allowedPieceCIDs, line)
				}
			}
		}

		// Parse HTTP headers from string slice to model.StringSlice
		var httpHeaders model.StringSlice
		if flagHeaders := c.StringSlice("http-header"); len(flagHeaders) > 0 {
			httpHeaders = model.StringSlice(flagHeaders)
		}

		template, err := dealtemplate.Default.CreateHandler(c.Context, db, dealtemplate.CreateRequest{
			Name:                 c.String("name"),
			Description:          c.String("description"),
			DealPricePerGB:       c.Float64("deal-price-per-gb"),
			DealPricePerGBEpoch:  c.Float64("deal-price-per-gb-epoch"),
			DealPricePerDeal:     c.Float64("deal-price-per-deal"),
			DealDuration:         c.Duration("deal-duration"),
			DealStartDelay:       c.Duration("deal-start-delay"),
			DealVerified:         c.Bool("deal-verified"),
			DealKeepUnsealed:     c.Bool("deal-keep-unsealed"),
			DealAnnounceToIPNI:   c.Bool("deal-announce-to-ipni"),
			DealProvider:         c.String("deal-provider"),
			DealURLTemplate:      c.String("deal-url-template"),
			DealHTTPHeaders:      dealHTTPHeaders,
			DealNotes:            c.String("notes"),
			DealForce:            c.Bool("force"),
			DealAllowedPieceCIDs: allowedPieceCIDs,

			// New scheduling fields
			ScheduleCron:       c.String("schedule-cron"),
			ScheduleDealNumber: c.Int("schedule-deal-number"),
			ScheduleDealSize:   c.String("schedule-deal-size"),

			// New restriction fields
			TotalDealNumber:      c.Int("total-deal-number"),
			TotalDealSize:        c.String("total-deal-size"),
			MaxPendingDealNumber: c.Int("max-pending-deal-number"),
			MaxPendingDealSize:   c.String("max-pending-deal-size"),

			// HTTP headers as string slice
			HTTPHeaders: httpHeaders,
		})
		if err != nil {
			return errors.WithStack(err)
		}

		// Print success confirmation
		if !c.Bool("json") {
			println("✓ Deal template \"" + template.Name + "\" created successfully")
		}

		cliutil.Print(c, *template)
		return nil
	},
}

// validateCreateTemplateInputs validates the inputs for creating a deal template
func validateCreateTemplateInputs(c *cli.Context) error {
	// Name is already required by CLI framework, but let's be explicit
	if c.String("name") == "" {
		return errors.New("template name is required")
	}

	// Validate pricing fields are non-negative
	if c.Float64("deal-price-per-gb") < 0 {
		return errors.New("deal price per GB must be non-negative")
	}
	if c.Float64("deal-price-per-gb-epoch") < 0 {
		return errors.New("deal price per GB epoch must be non-negative")
	}
	if c.Float64("deal-price-per-deal") < 0 {
		return errors.New("deal price per deal must be non-negative")
	}

	// Validate durations are non-negative
	if c.Duration("deal-duration") < 0 {
		return errors.New("deal duration cannot be negative")
	}
	if c.Duration("deal-start-delay") < 0 {
		return errors.New("deal start delay cannot be negative")
	}

	// Validate deal provider format if provided
	if provider := c.String("deal-provider"); provider != "" {
		if len(provider) < 3 || (provider[:2] != "f0" && provider[:2] != "t0") {
			return errors.New("deal provider must be a valid storage provider ID (e.g., f01234 or t01234)")
		}
	}

	// Validate HTTP headers if provided
	if headersStr := c.String("deal-http-headers"); headersStr != "" {
		var tempMap map[string]string
		if err := json.Unmarshal([]byte(headersStr), &tempMap); err != nil {
			return errors.Wrapf(err, "invalid JSON format for deal-http-headers")
		}

		// Validate header keys and values
		for key, value := range tempMap {
			if key == "" {
				return errors.New("HTTP header keys cannot be empty")
			}
			if value == "" {
				return errors.New("HTTP header values cannot be empty")
			}
		}
	}

	// Validate allowed piece CID file if provided
	if filePath := c.String("allowed-piece-cid-file"); filePath != "" {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return errors.Wrapf(err, "allowed-piece-cid-file does not exist: %s", filePath)
		}
	}

	// Validate scheduling fields
	if scheduleCron := c.String("schedule-cron"); scheduleCron != "" {
		// Basic validation for cron format - could be enhanced with actual cron parsing
		if !strings.HasPrefix(scheduleCron, "@") && len(strings.Fields(scheduleCron)) < 5 {
			return errors.New("invalid cron format - use descriptors like @daily or standard cron format")
		}
	}

	// Validate deal numbers are non-negative
	if c.Int("schedule-deal-number") < 0 {
		return errors.New("schedule deal number cannot be negative")
	}
	if c.Int("total-deal-number") < 0 {
		return errors.New("total deal number cannot be negative")
	}
	if c.Int("max-pending-deal-number") < 0 {
		return errors.New("max pending deal number cannot be negative")
	}

	// Validate HTTP headers format
	for _, header := range c.StringSlice("http-header") {
		if !strings.Contains(header, "=") {
			return errors.Errorf("invalid HTTP header format '%s' - use key=value format", header)
		}
		parts := strings.SplitN(header, "=", 2)
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			return errors.Errorf("invalid HTTP header format '%s' - both key and value must be non-empty", header)
		}
	}

	return nil
}
