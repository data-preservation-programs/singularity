package dealtemplate

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dealtemplate"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/urfave/cli/v2"
)

var CreateCmd = &cli.Command{
	Name:  "create",
	Usage: "Create a new deal template with unified flags and defaults",
	Description: `Create a new deal template using the same flags and default values as deal schedule create.

Key flags:
  --provider           Storage Provider ID (e.g., f01234)
  --duration           Deal duration (default: 12840h)
  --start-delay        Deal start delay (default: 72h)
  --verified           Propose deals as verified (default: true)
  --keep-unsealed      Keep unsealed copy (default: true)
  --ipni               Announce deals to IPNI (default: true)
  --http-header        HTTP headers (key=value)
  --allowed-piece-cid  List of allowed piece CIDs
  --allowed-piece-cid-file File with allowed piece CIDs

See --help for all options.`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Usage:    "Name of the deal template",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "provider",
			Usage:    "Storage Provider ID (e.g., f01000)",
			Required: true,
		},
		&cli.Float64Flag{
			Name:  "price-per-gb",
			Usage: "Price in FIL per GiB for storage deals",
			Value: 0.0,
		},
		&cli.Float64Flag{
			Name:  "price-per-gb-epoch",
			Usage: "Price in FIL per GiB per epoch for storage deals",
			Value: 0.0,
		},
		&cli.Float64Flag{
			Name:  "price-per-deal",
			Usage: "Price in FIL per deal for storage deals",
			Value: 0.0,
		},
		&cli.DurationFlag{
			Name:  "duration",
			Usage: "Duration for storage deals (e.g., 12840h for 535 days)",
			Value: 12840 * time.Hour,
		},
		&cli.DurationFlag{
			Name:  "start-delay",
			Usage: "Start delay for storage deals",
			Value: 72 * time.Hour,
		},
		&cli.BoolFlag{
			Name:  "verified",
			Usage: "Whether deals should be verified",
			Value: true,
		},
		&cli.BoolFlag{
			Name:  "keep-unsealed",
			Usage: "Whether to keep unsealed copy of deals",
			Value: true,
		},
		&cli.BoolFlag{
			Name:  "ipni",
			Usage: "Whether to announce deals to IPNI",
			Value: true,
		},
		&cli.StringFlag{
			Name:  "url-template",
			Usage: "URL template for deals",
		},
		&cli.StringSliceFlag{
			Name:  "http-header",
			Usage: "HTTP headers to be passed with the request (key=value format)",
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
		// Scheduling flags
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
		// Restriction flags
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

		// Parse HTTP headers from string slice format to ConfigMap
		var dealHTTPHeaders model.ConfigMap
		httpHeaders := c.StringSlice("http-header")
		if len(httpHeaders) > 0 {
			tempMap := make(map[string]string)
			for _, header := range httpHeaders {
				parts := strings.SplitN(header, "=", 2)
				if len(parts) != 2 {
					return errors.Errorf("invalid HTTP header format: %s (expected key=value)", header)
				}
				tempMap[parts[0]] = parts[1]
			}
			dealHTTPHeaders = model.ConfigMap(tempMap)
		}

		// Parse allowed piece CIDs from flags and file
		var allowedPieceCIDs model.StringSlice
		cidSet := make(map[string]bool) // Use map to track unique CIDs

		// Add piece CIDs from flag
		if flagCIDs := c.StringSlice("allowed-piece-cid"); len(flagCIDs) > 0 {
			for _, cid := range flagCIDs {
				cid = strings.TrimSpace(cid)
				if cid != "" && !cidSet[cid] {
					cidSet[cid] = true
					allowedPieceCIDs = append(allowedPieceCIDs, cid)
				}
			}
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
				if line != "" && !strings.HasPrefix(line, "#") && !cidSet[line] { // Skip empty lines, comments, and duplicates
					cidSet[line] = true
					allowedPieceCIDs = append(allowedPieceCIDs, line)
				}
			}
		}

		template, err := dealtemplate.Default.CreateHandler(c.Context, db, dealtemplate.CreateRequest{
			Name:                 c.String("name"),
			Description:          c.String("description"),
			DealPricePerGB:       c.Float64("price-per-gb"),
			DealPricePerGBEpoch:  c.Float64("price-per-gb-epoch"),
			DealPricePerDeal:     c.Float64("price-per-deal"),
			DealDuration:         c.Duration("duration"),
			DealStartDelay:       c.Duration("start-delay"),
			DealVerified:         c.Bool("verified"),
			DealKeepUnsealed:     c.Bool("keep-unsealed"),
			DealAnnounceToIPNI:   c.Bool("ipni"),
			DealProvider:         c.String("provider"),
			DealURLTemplate:      c.String("url-template"),
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
			HTTPHeaders: model.StringSlice(httpHeaders),
		})
		if err != nil {
			return errors.WithStack(err)
		}

		// Always print as pretty JSON
		jsonBytes, err := json.MarshalIndent(template, "", "  ")
		if err != nil {
			return errors.Wrap(err, "failed to marshal template as JSON")
		}
		if _, err := os.Stdout.Write(jsonBytes); err != nil {
			return errors.Wrap(err, "failed to write JSON to stdout")
		}
		if _, err := os.Stdout.Write([]byte("\n")); err != nil {
			return errors.Wrap(err, "failed to write newline to stdout")
		}
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
	if c.Float64("price-per-gb") < 0 {
		return errors.New("deal price per GB must be non-negative")
	}
	if c.Float64("price-per-gb-epoch") < 0 {
		return errors.New("deal price per GB epoch must be non-negative")
	}
	if c.Float64("price-per-deal") < 0 {
		return errors.New("deal price per deal must be non-negative")
	}

	// Validate durations are non-negative
	if c.Duration("duration") < 0 {
		return errors.New("deal duration cannot be negative")
	}
	if c.Duration("start-delay") < 0 {
		return errors.New("deal start delay cannot be negative")
	}

	// Validate deal provider format if provided
	if provider := c.String("provider"); provider != "" {
		if len(provider) < 3 || (provider[:2] != "f0" && provider[:2] != "t0") {
			return errors.New("deal provider must be a valid storage provider ID (e.g., f01234 or t01234)")
		}
	}

	// Validate HTTP headers if provided
	httpHeaders := c.StringSlice("http-header")
	for _, header := range httpHeaders {
		parts := strings.SplitN(header, "=", 2)
		if len(parts) != 2 {
			return errors.Errorf("invalid HTTP header format: %s (expected key=value)", header)
		}
		if parts[0] == "" {
			return errors.New("HTTP header keys cannot be empty")
		}
		if parts[1] == "" {
			return errors.New("HTTP header values cannot be empty")
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
