package dataprep

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"path/filepath"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/handler/job"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/workflow"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var CreateCmd = &cli.Command{
	Name:     "create",
	Usage:    "Create a new preparation",
	Category: "Preparation Management",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Usage:       "The name for the preparation",
			DefaultText: "Auto generated",
		},
		&cli.StringSliceFlag{
			Name:  "source",
			Usage: "The id or name of the source storage to be used for the preparation",
		},
		&cli.StringSliceFlag{
			Name:  "output",
			Usage: "The id or name of the output storage to be used for the preparation",
		},
		&cli.StringSliceFlag{
			Name:     "local-source",
			Category: "Quick creation with local source paths",
			Usage:    "The local source path to be used for the preparation. This is a convenient flag that will create a source storage with the provided path",
		},
		&cli.StringSliceFlag{
			Name:     "local-output",
			Category: "Quick creation with local output paths",
			Usage:    "The local output path to be used for the preparation. This is a convenient flag that will create a output storage with the provided path",
		},
		&cli.StringFlag{
			Name:  "max-size",
			Usage: "The maximum size of a single CAR file",
			Value: "31.5GiB",
		},
		&cli.StringFlag{
			Name:        "piece-size",
			Usage:       "The target piece size of the CAR files used for piece commitment calculation",
			Value:       "",
			DefaultText: "Determined by --max-size",
		},
		&cli.StringFlag{
			Name:        "min-piece-size",
			Usage:       "The minimum size of a piece. Pieces smaller than this will be padded up to this size. It's recommended to leave this as the default",
			Value:       "1MiB",
			DefaultText: "1MiB",
		},
		&cli.BoolFlag{
			Name:  "delete-after-export",
			Usage: "Whether to delete the source files after export to CAR files",
		},
		&cli.BoolFlag{
			Name:  "no-inline",
			Usage: "Whether to disable inline storage for the preparation. Can save database space but requires at least one output storage.",
		},
		&cli.BoolFlag{
			Name:  "no-dag",
			Usage: "Whether to disable maintaining folder dag structure for the sources. If disabled, DagGen will not be possible and folders will not have an associated CID.",
		},
		&cli.BoolFlag{
			Name:     "auto-create-deals",
			Usage:    "Enable automatic deal schedule creation after preparation completion",
			Category: "Auto Deal Creation",
		},
		&cli.StringFlag{
			Name:     "deal-template",
			Usage:    "Name or ID of deal template to use (optional - can specify deal parameters directly instead)",
			Category: "Auto Deal Creation",
		},
		&cli.Float64Flag{
			Name:     "deal-price-per-gb",
			Usage:    "Price in FIL per GiB for storage deals",
			Value:    0.0,
			Category: "Auto Deal Creation",
		},
		&cli.Float64Flag{
			Name:     "deal-price-per-gb-epoch",
			Usage:    "Price in FIL per GiB per epoch for storage deals",
			Value:    0.0,
			Category: "Auto Deal Creation",
		},
		&cli.Float64Flag{
			Name:     "deal-price-per-deal",
			Usage:    "Price in FIL per deal for storage deals",
			Value:    0.0,
			Category: "Auto Deal Creation",
		},
		&cli.DurationFlag{
			Name:     "deal-duration",
			Usage:    "Duration for storage deals (e.g., 535 days)",
			Value:    0,
			Category: "Auto Deal Creation",
		},
		&cli.DurationFlag{
			Name:     "deal-start-delay",
			Usage:    "Start delay for storage deals (e.g., 72h)",
			Value:    0,
			Category: "Auto Deal Creation",
		},
		&cli.BoolFlag{
			Name:     "deal-verified",
			Usage:    "Whether deals should be verified",
			Category: "Auto Deal Creation",
		},
		&cli.BoolFlag{
			Name:     "deal-keep-unsealed",
			Usage:    "Whether to keep unsealed copy of deals",
			Category: "Auto Deal Creation",
		},
		&cli.BoolFlag{
			Name:     "deal-announce-to-ipni",
			Usage:    "Whether to announce deals to IPNI",
			Category: "Auto Deal Creation",
		},
		&cli.StringFlag{
			Name:     "deal-provider",
			Usage:    "Storage Provider ID for deals (e.g., f01000)",
			Category: "Auto Deal Creation",
		},
		&cli.StringFlag{
			Name:     "deal-url-template",
			Usage:    "URL template for deals",
			Category: "Auto Deal Creation",
		},
		&cli.StringFlag{
			Name:     "deal-http-headers",
			Usage:    "HTTP headers for deals in JSON format",
			Category: "Auto Deal Creation",
		},
		&cli.BoolFlag{
			Name:     "wallet-validation",
			Usage:    "Enable wallet balance validation before deal creation",
			Category: "Validation",
			Value:    true,
		},
		&cli.BoolFlag{
			Name:     "sp-validation",
			Usage:    "Enable storage provider validation before deal creation",
			Category: "Validation",
			Value:    true,
		},
		&cli.BoolFlag{
			Name:     "auto-start",
			Usage:    "Automatically start scanning after preparation creation",
			Category: "Workflow Automation",
		},
		&cli.BoolFlag{
			Name:     "auto-progress",
			Usage:    "Enable automatic job progression (scan → pack → daggen → deals)",
			Category: "Workflow Automation",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		db = db.WithContext(c.Context)
		name := c.String("name")
		if name == "" {
			name = util.RandomName()
		}
		sourceStorages := c.StringSlice("source")
		outputStorages := c.StringSlice("output")
		maxSizeStr := c.String("max-size")
		pieceSizeStr := c.String("piece-size")
		minPieceSizeStr := c.String("min-piece-size")
		for _, sourcePath := range c.StringSlice("local-source") {
			source, err := createStorageIfNotExist(c.Context, db, sourcePath)
			if err != nil {
				return errors.WithStack(err)
			}
			sourceStorages = append(sourceStorages, source.Name)
		}
		for _, outputPath := range c.StringSlice("local-output") {
			output, err := createStorageIfNotExist(c.Context, db, outputPath)
			if err != nil {
				return errors.WithStack(err)
			}
			outputStorages = append(outputStorages, output.Name)
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

		prep, err := dataprep.Default.CreatePreparationHandler(c.Context, db, dataprep.CreateRequest{
			SourceStorages:      sourceStorages,
			OutputStorages:      outputStorages,
			MaxSizeStr:          maxSizeStr,
			PieceSizeStr:        pieceSizeStr,
			MinPieceSizeStr:     minPieceSizeStr,
			DeleteAfterExport:   c.Bool("delete-after-export"),
			Name:                name,
			NoInline:            c.Bool("no-inline"),
			NoDag:               c.Bool("no-dag"),
			AutoCreateDeals:     c.Bool("auto-create-deals"),
			DealTemplate:        c.String("deal-template"),
			DealPricePerGB:      c.Float64("deal-price-per-gb"),
			DealPricePerGBEpoch: c.Float64("deal-price-per-gb-epoch"),
			DealPricePerDeal:    c.Float64("deal-price-per-deal"),
			DealDuration:        c.Duration("deal-duration"),
			DealStartDelay:      c.Duration("deal-start-delay"),
			DealVerified:        c.Bool("deal-verified"),
			DealKeepUnsealed:    c.Bool("deal-keep-unsealed"),
			DealAnnounceToIPNI:  c.Bool("deal-announce-to-ipni"),
			DealProvider:        c.String("deal-provider"),
			DealURLTemplate:     c.String("deal-url-template"),
			DealHTTPHeaders:     dealHTTPHeaders,
			WalletValidation:    c.Bool("wallet-validation"),
			SPValidation:        c.Bool("sp-validation"),
		})
		if err != nil {
			return errors.WithStack(err)
		}

		// Enable workflow orchestration if auto-progress is requested
		if c.Bool("auto-progress") {
			enableWorkflowOrchestration(c.Context)
		}

		// Auto-start scanning if requested
		if c.Bool("auto-start") {
			err = autoStartScanning(c.Context, db, prep)
			if err != nil {
				return errors.Wrap(err, "failed to auto-start scanning")
			}
		}

		cliutil.Print(c, *prep)
		return nil
	},
}

func createStorageIfNotExist(ctx context.Context, db *gorm.DB, sourcePath string) (*model.Storage, error) {
	db = db.WithContext(ctx)
	path, err := filepath.Abs(sourcePath)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid path: %s", sourcePath)
	}
	existing := &model.Storage{}
	err = db.Where("type = ? AND path = ?", "local", path).
		First(existing).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.WithStack(err)
	}
	if err != nil {
		name := filepath.Base(path)
		if name == "." {
			name = ""
		}
		name += "-" + randomReadableString(4)
		existing, err = storage.Default.CreateStorageHandler(
			ctx,
			db,
			"local", storage.CreateRequest{
				Name: name,
				Path: path,
			})
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return existing, nil
}

func randomReadableString(length int) string {
	const charset = "0123456789abcdef"

	b := make([]byte, length)
	for i := range b {
		//nolint:gosec
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// enableWorkflowOrchestration enables the workflow orchestrator for automatic job progression
func enableWorkflowOrchestration(_ context.Context) {
	workflow.DefaultOrchestrator.SetEnabled(true)
	fmt.Printf("✓ Workflow orchestration enabled (automatic scan → pack → daggen → deals)\n")
}

// autoStartScanning automatically starts scanning for all source attachments in the preparation
func autoStartScanning(ctx context.Context, db *gorm.DB, prep *model.Preparation) error {
	// Get all source attachments for this preparation
	var attachments []model.SourceAttachment
	err := db.WithContext(ctx).Where("preparation_id = ?", prep.ID).Find(&attachments).Error
	if err != nil {
		return errors.WithStack(err)
	}

	if len(attachments) == 0 {
		fmt.Printf("⚠ No source attachments found for preparation %s\n", prep.Name)
		return nil
	}

	jobHandler := &job.DefaultHandler{}
	successCount := 0

	// Start scan jobs for each source attachment
	for _, attachment := range attachments {
		_, err = jobHandler.StartScanHandler(ctx, db, strconv.FormatUint(uint64(attachment.PreparationID), 10), strconv.FormatUint(uint64(attachment.StorageID), 10))
		if err != nil {
			fmt.Printf("⚠ Failed to start scan for attachment %d: %v\n", attachment.ID, err)
			continue
		}
		successCount++
	}

	if successCount > 0 {
		fmt.Printf("✓ Started scanning for %d source attachment(s) in preparation %s\n", successCount, prep.Name)
		if successCount < len(attachments) {
			fmt.Printf("⚠ %d attachment(s) failed to start scanning\n", len(attachments)-successCount)
		}
	} else {
		return errors.New("failed to start scanning for any attachments")
	}

	return nil
}

// StartScanningForPreparation starts scanning for all source attachments in a preparation
func StartScanningForPreparation(ctx context.Context, db *gorm.DB, prep *model.Preparation) error {
	return autoStartScanning(ctx, db, prep)
}
