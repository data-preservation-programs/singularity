package dealtemplate

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dealtemplate"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/urfave/cli/v2"
)

var UpdateCmd = &cli.Command{
	Name:      "update",
	Usage:     "Update an existing deal template",
	Category:  "Deal Template Management",
	ArgsUsage: "<template_id_or_name>",
	Description: `Update an existing deal template with new values. Only specified flags will be updated.

Key flags:
  --name               New name for the template
  --provider           Storage Provider ID (e.g., f01234)
  --duration           Deal duration (e.g., 12840h)
  --start-delay        Deal start delay (e.g., 72h)
  --verified           Propose deals as verified
  --keep-unsealed      Keep unsealed copy
  --ipni               Announce deals to IPNI
  --http-header        HTTP headers (key=value)
  --allowed-piece-cid  List of allowed piece CIDs
  --allowed-piece-cid-file File with allowed piece CIDs

Piece CID Handling:
  By default, piece CIDs are merged with existing ones. 
  Use --replace-piece-cids to completely replace the existing list.

See --help for all options.`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "name",
			Usage: "New name for the deal template",
		},
		&cli.StringFlag{
			Name:  "description",
			Usage: "Description of the deal template",
		},
		&cli.StringFlag{
			Name:  "provider",
			Usage: "Storage Provider ID (e.g., f01000)",
		},
		&cli.Float64Flag{
			Name:  "price-per-gb",
			Usage: "Price in FIL per GiB for storage deals",
		},
		&cli.Float64Flag{
			Name:  "price-per-gb-epoch",
			Usage: "Price in FIL per GiB per epoch for storage deals",
		},
		&cli.Float64Flag{
			Name:  "price-per-deal",
			Usage: "Price in FIL per deal for storage deals",
		},
		&cli.DurationFlag{
			Name:  "duration",
			Usage: "Duration for storage deals (e.g., 12840h for 535 days)",
		},
		&cli.DurationFlag{
			Name:  "start-delay",
			Usage: "Start delay for storage deals",
		},
		&cli.BoolFlag{
			Name:  "verified",
			Usage: "Whether deals should be verified",
		},
		&cli.BoolFlag{
			Name:  "keep-unsealed",
			Usage: "Whether to keep unsealed copy of deals",
		},
		&cli.BoolFlag{
			Name:  "ipni",
			Usage: "Whether to announce deals to IPNI",
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
			Usage: "Force deals regardless of replication restrictions",
		},
		&cli.StringSliceFlag{
			Name:  "allowed-piece-cid",
			Usage: "List of allowed piece CIDs for this template",
		},
		&cli.StringFlag{
			Name:  "allowed-piece-cid-file",
			Usage: "File containing list of allowed piece CIDs",
		},
		&cli.BoolFlag{
			Name:  "replace-piece-cids",
			Usage: "Replace existing piece CIDs instead of merging (use with --allowed-piece-cid or --allowed-piece-cid-file)",
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
		},
	},
	Action: func(c *cli.Context) error {
		if c.NArg() != 1 {
			return errors.New("template ID or name is required")
		}

		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		db = db.WithContext(c.Context)

		templateIdentifier := c.Args().First()

		// Validate inputs
		if err := validateUpdateTemplateInputs(c); err != nil {
			return errors.Wrap(err, "validation failed")
		}

		// Build update request with only set fields
		request := dealtemplate.UpdateRequest{}

		// Update basic fields only if set
		if c.IsSet("name") {
			name := c.String("name")
			request.Name = &name
		}
		if c.IsSet("description") {
			description := c.String("description")
			request.Description = &description
		}
		if c.IsSet("provider") {
			provider := c.String("provider")
			request.DealProvider = &provider
		}
		if c.IsSet("price-per-gb") {
			price := c.Float64("price-per-gb")
			request.DealPricePerGB = &price
		}
		if c.IsSet("price-per-gb-epoch") {
			price := c.Float64("price-per-gb-epoch")
			request.DealPricePerGBEpoch = &price
		}
		if c.IsSet("price-per-deal") {
			price := c.Float64("price-per-deal")
			request.DealPricePerDeal = &price
		}
		if c.IsSet("duration") {
			duration := c.Duration("duration")
			request.DealDuration = &duration
		}
		if c.IsSet("start-delay") {
			startDelay := c.Duration("start-delay")
			request.DealStartDelay = &startDelay
		}
		if c.IsSet("verified") {
			verified := c.Bool("verified")
			request.DealVerified = &verified
		}
		if c.IsSet("keep-unsealed") {
			keepUnsealed := c.Bool("keep-unsealed")
			request.DealKeepUnsealed = &keepUnsealed
		}
		if c.IsSet("ipni") {
			ipni := c.Bool("ipni")
			request.DealAnnounceToIPNI = &ipni
		}
		if c.IsSet("url-template") {
			urlTemplate := c.String("url-template")
			request.DealURLTemplate = &urlTemplate
		}
		if c.IsSet("notes") {
			notes := c.String("notes")
			request.DealNotes = &notes
		}
		if c.IsSet("force") {
			force := c.Bool("force")
			request.DealForce = &force
		}

		// Handle HTTP headers
		if c.IsSet("http-header") {
			httpHeaders := c.StringSlice("http-header")
			var dealHTTPHeaders model.ConfigMap
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
			request.DealHTTPHeaders = &dealHTTPHeaders
			// Also set HTTPHeaders as string slice
			httpHeaderSlice := model.StringSlice(httpHeaders)
			request.HTTPHeaders = &httpHeaderSlice
		}

		// Handle piece CIDs with merge/replace logic
		if c.IsSet("allowed-piece-cid") || c.IsSet("allowed-piece-cid-file") {
			var newPieceCIDs model.StringSlice
			cidSet := make(map[string]bool) // Use map to track unique CIDs

			// If not replacing, get existing piece CIDs first
			if !c.Bool("replace-piece-cids") {
				existingTemplate, err := dealtemplate.Default.GetHandler(c.Context, db, templateIdentifier)
				if err != nil {
					return errors.Wrap(err, "failed to get existing template for piece CID merging")
				}

				// Add existing CIDs to the set
				for _, cid := range existingTemplate.DealConfig.DealAllowedPieceCIDs {
					cid = strings.TrimSpace(cid)
					if cid != "" && !cidSet[cid] {
						cidSet[cid] = true
						newPieceCIDs = append(newPieceCIDs, cid)
					}
				}
			}

			// Add piece CIDs from flag
			if flagCIDs := c.StringSlice("allowed-piece-cid"); len(flagCIDs) > 0 {
				for _, cid := range flagCIDs {
					cid = strings.TrimSpace(cid)
					if cid != "" && !cidSet[cid] {
						cidSet[cid] = true
						newPieceCIDs = append(newPieceCIDs, cid)
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
						newPieceCIDs = append(newPieceCIDs, line)
					}
				}
			}

			request.DealAllowedPieceCIDs = &newPieceCIDs
		}

		// Handle scheduling fields
		if c.IsSet("schedule-cron") {
			scheduleCron := c.String("schedule-cron")
			request.ScheduleCron = &scheduleCron
		}
		if c.IsSet("schedule-deal-number") {
			scheduleDealNumber := c.Int("schedule-deal-number")
			request.ScheduleDealNumber = &scheduleDealNumber
		}
		if c.IsSet("schedule-deal-size") {
			scheduleDealSize := c.String("schedule-deal-size")
			request.ScheduleDealSize = &scheduleDealSize
		}

		// Handle restriction fields
		if c.IsSet("total-deal-number") {
			totalDealNumber := c.Int("total-deal-number")
			request.TotalDealNumber = &totalDealNumber
		}
		if c.IsSet("total-deal-size") {
			totalDealSize := c.String("total-deal-size")
			request.TotalDealSize = &totalDealSize
		}
		if c.IsSet("max-pending-deal-number") {
			maxPendingDealNumber := c.Int("max-pending-deal-number")
			request.MaxPendingDealNumber = &maxPendingDealNumber
		}
		if c.IsSet("max-pending-deal-size") {
			maxPendingDealSize := c.String("max-pending-deal-size")
			request.MaxPendingDealSize = &maxPendingDealSize
		}

		template, err := dealtemplate.Default.UpdateHandler(c.Context, db, templateIdentifier, request)
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

// validateUpdateTemplateInputs validates the inputs for updating a deal template
func validateUpdateTemplateInputs(c *cli.Context) error {
	// Validate pricing fields are non-negative (if set)
	if c.IsSet("price-per-gb") && c.Float64("price-per-gb") < 0 {
		return errors.New("deal price per GB must be non-negative")
	}
	if c.IsSet("price-per-gb-epoch") && c.Float64("price-per-gb-epoch") < 0 {
		return errors.New("deal price per GB epoch must be non-negative")
	}
	if c.IsSet("price-per-deal") && c.Float64("price-per-deal") < 0 {
		return errors.New("deal price per deal must be non-negative")
	}

	// Validate durations are non-negative (if set)
	if c.IsSet("duration") && c.Duration("duration") < 0 {
		return errors.New("deal duration cannot be negative")
	}
	if c.IsSet("start-delay") && c.Duration("start-delay") < 0 {
		return errors.New("deal start delay cannot be negative")
	}

	// Validate deal provider format if provided
	if c.IsSet("provider") {
		if provider := c.String("provider"); provider != "" {
			if len(provider) < 3 || (provider[:2] != "f0" && provider[:2] != "t0") {
				return errors.New("deal provider must be a valid storage provider ID (e.g., f01234 or t01234)")
			}
		}
	}

	// Validate HTTP headers if provided
	if c.IsSet("http-header") {
		httpHeaders := c.StringSlice("http-header")
		for _, header := range httpHeaders {
			parts := strings.SplitN(header, "=", 2)
			if len(parts) != 2 {
				return errors.Errorf("invalid HTTP header format: %s (expected key=value)", header)
			}
			if parts[0] == "" {
				return errors.New("HTTP header keys cannot be empty")
			}
			// Allow empty values for deletion
		}
	}

	// Validate allowed piece CID file if provided
	if c.IsSet("allowed-piece-cid-file") {
		if filePath := c.String("allowed-piece-cid-file"); filePath != "" {
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				return errors.Wrapf(err, "allowed-piece-cid-file does not exist: %s", filePath)
			}
		}
	}

	// Validate scheduling fields
	if c.IsSet("schedule-cron") {
		if scheduleCron := c.String("schedule-cron"); scheduleCron != "" {
			// Basic validation for cron format - could be enhanced with actual cron parsing
			if !strings.HasPrefix(scheduleCron, "@") && len(strings.Fields(scheduleCron)) < 5 {
				return errors.New("invalid cron format - use descriptors like @daily or standard cron format")
			}
		}
	}

	// Validate deal numbers are non-negative (if set)
	if c.IsSet("schedule-deal-number") && c.Int("schedule-deal-number") < 0 {
		return errors.New("schedule deal number cannot be negative")
	}
	if c.IsSet("total-deal-number") && c.Int("total-deal-number") < 0 {
		return errors.New("total deal number cannot be negative")
	}
	if c.IsSet("max-pending-deal-number") && c.Int("max-pending-deal-number") < 0 {
		return errors.New("max pending deal number cannot be negative")
	}

	// Validate that replace-piece-cids is only used with piece CID flags
	if c.Bool("replace-piece-cids") && !c.IsSet("allowed-piece-cid") && !c.IsSet("allowed-piece-cid-file") {
		return errors.New("--replace-piece-cids can only be used with --allowed-piece-cid or --allowed-piece-cid-file")
	}

	return nil
}
