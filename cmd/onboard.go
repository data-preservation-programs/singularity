package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/handler/job"
	storageHandlers "github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/workermanager"
	"github.com/data-preservation-programs/singularity/service/workflow"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/dustin/go-humanize"
	"github.com/gotidy/ptr"
	"github.com/rclone/rclone/fs"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

// OnboardResult represents the JSON output for the onboard command
type OnboardResult struct {
	Success       bool     `json:"success"`
	PreparationID uint32   `json:"preparationId"`
	Name          string   `json:"name"`
	SourcePaths   []string `json:"sourcePaths"`
	OutputPaths   []string `json:"outputPaths"`
	AutoDeals     bool     `json:"autoDeals"`
	WorkersCount  int      `json:"workersCount"`
	NextSteps     []string `json:"nextSteps"`
	Error         string   `json:"error,omitempty"`
}

// OnboardCmd provides a single command for complete data onboarding
var OnboardCmd = &cli.Command{
	Name:  "onboard",
	Usage: "Complete data onboarding workflow (storage → preparation → scanning → deal creation)",
	Description: `The onboard command provides a unified workflow for complete data onboarding.

It performs the following steps automatically:
1. Creates storage connections (if paths provided)
2. Creates data preparation with deal template configuration
3. Starts scanning immediately
4. Enables automatic job progression (scan → pack → daggen → deals)
5. Optionally starts managed workers to process jobs

This is the simplest way to onboard data from source to storage deals.
Use deal templates to configure deal parameters - individual deal flags are not supported.

EXAMPLES:
  # Basic local data onboarding
  singularity onboard --name "my-dataset" --source "/path/to/data" --deal-template-id "1"

  # S3 to local with custom output
  singularity onboard --name "s3-data" \
    --source "s3://bucket/data" --source-type "s3" \
    --output "/mnt/storage/cars" \
    --deal-template-id "template1"

  # Multiple sources with monitoring
  singularity onboard --name "multi-source" \
    --source "/data1" --source "/data2" \
    --wait-for-completion --max-workers 5 \
    --deal-template-id "prod-template"

For backend-specific options, use --help to see all available flags.`,
	Flags: []cli.Flag{
		// Required Options
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "Name for the preparation",
			Required: true,
			Category: "Required Options",
		},
		&cli.StringSliceFlag{
			Name:     "source",
			Aliases:  []string{"s"},
			Usage:    "Source path(s) to onboard (local paths or remote URLs like s3://bucket/path)",
			Required: true,
			Category: "Required Options",
		},

		// Data Source Configuration
		&cli.StringFlag{
			Name:     "source-type",
			Usage:    "Source storage type (local, s3, gcs, azure, etc.)",
			Value:    "local",
			Category: "Data Source",
		},
		&cli.StringFlag{
			Name:     "source-provider",
			Usage:    "Source storage provider (for s3: aws, minio, wasabi, etc.)",
			Category: "Data Source",
		},
		&cli.StringFlag{
			Name:     "source-name",
			Usage:    "Custom name for source storage (auto-generated if not provided)",
			Category: "Data Source",
		},
		&cli.StringFlag{
			Name:     "source-config",
			Usage:    "Source storage configuration in JSON format (key-value pairs)",
			Category: "Data Source",
		},

		// Data Output Configuration
		&cli.StringSliceFlag{
			Name:     "output",
			Aliases:  []string{"o"},
			Usage:    "Output path(s) for CAR files (local paths or remote URLs like s3://bucket/path)",
			Category: "Data Output",
		},
		&cli.StringFlag{
			Name:     "output-type",
			Usage:    "Output storage type (local, s3, gcs, azure, etc.)",
			Value:    "local",
			Category: "Data Output",
		},
		&cli.StringFlag{
			Name:     "output-provider",
			Usage:    "Output storage provider",
			Category: "Data Output",
		},
		&cli.StringFlag{
			Name:     "output-name",
			Usage:    "Custom name for output storage (auto-generated if not provided)",
			Category: "Data Output",
		},
		&cli.StringFlag{
			Name:     "output-config",
			Usage:    "Output storage configuration in JSON format (key-value pairs)",
			Category: "Data Output",
		},

		// Data Preparation
		&cli.StringFlag{
			Name:     "max-size",
			Usage:    "Maximum size of a single CAR file",
			Value:    "31.5GiB",
			Category: "Data Preparation",
		},
		&cli.BoolFlag{
			Name:     "no-dag",
			Usage:    "Disable maintaining folder DAG structure",
			Category: "Data Preparation",
		},

		// Deal Configuration
		&cli.BoolFlag{
			Name:     "auto-create-deals",
			Usage:    "Enable automatic deal creation after preparation completion",
			Value:    true,
			Category: "Deal Configuration",
		},
		&cli.StringFlag{
			Name:     "deal-template-id",
			Aliases:  []string{"t"},
			Usage:    "Deal template ID to use for deal configuration (required when auto-create-deals is enabled)",
			Category: "Deal Configuration",
		},

		// Worker Management
		&cli.BoolFlag{
			Name:     "start-workers",
			Usage:    "Start managed workers to process jobs automatically",
			Value:    true,
			Category: "Worker Management",
		},
		&cli.IntFlag{
			Name:     "max-workers",
			Aliases:  []string{"w"},
			Usage:    "Maximum number of workers to run",
			Value:    3,
			Category: "Worker Management",
		},

		// Monitoring & Progress
		&cli.BoolFlag{
			Name:     "wait-for-completion",
			Usage:    "Wait and monitor until all jobs complete",
			Category: "Monitoring",
		},
		&cli.DurationFlag{
			Name:     "timeout",
			Usage:    "Timeout for waiting for completion (0 = no timeout)",
			Value:    0,
			Category: "Monitoring",
		},
		&cli.BoolFlag{
			Name:     "json",
			Usage:    "Output result in JSON format for automation",
			Category: "Monitoring",
		},

		// Validation Options
		&cli.BoolFlag{
			Name:     "wallet-validation",
			Usage:    "Enable wallet balance validation",
			Category: "Validation",
		},
		&cli.BoolFlag{
			Name:     "sp-validation",
			Usage:    "Enable storage provider validation",
			Category: "Validation",
		},
	},
}

// Initialize common flags for the onboard command
func init() {
	// Add common storage client flags to the onboard command
	OnboardCmd.Flags = append(OnboardCmd.Flags, cliutil.CommonStorageClientFlags...)

	// Add dynamic storage flags for all supported backends
	OnboardCmd.Flags = append(OnboardCmd.Flags, generateDynamicStorageFlags()...)

	// Set the action function
	OnboardCmd.Action = onboardAction
}

// Action function for the onboard command
func onboardAction(c *cli.Context) error {
	isJSON := c.Bool("json")

	// Helper function to output JSON error and exit
	outputJSONError := func(msg string, err error) error {
		if isJSON {
			result := OnboardResult{
				Success: false,
				Error:   fmt.Sprintf("%s: %v", msg, err),
			}
			if data, err := json.Marshal(result); err == nil {
				fmt.Println(string(data))
			}
		}
		return errors.Wrap(err, msg)
	}

	// Validate CLI inputs before proceeding
	if err := validateOnboardInputs(c); err != nil {
		return outputJSONError("input validation failed", err)
	}

	// Initialize database
	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return outputJSONError("failed to initialize database", err)
	}
	defer func() { _ = closer.Close() }()

	ctx := c.Context

	// Validate deal template exists if specified
	if c.Bool("auto-create-deals") {
		if err := validateDealTemplateExists(ctx, db, c.String("deal-template-id")); err != nil {
			return outputJSONError("deal template validation failed", err)
		}
	}

	if !isJSON {
		fmt.Println("🚀 Starting unified data onboarding...")
	}

	// Step 1: Create preparation with deal configuration
	if !isJSON {
		fmt.Println("\n📋 Creating data preparation...")
	}
	prep, err := createPreparationForOnboarding(ctx, db, c)
	if err != nil {
		return outputJSONError("failed to create preparation", err)
	}
	if !isJSON {
		fmt.Printf("✓ Created preparation: %s (ID: %d)\n", prep.Name, prep.ID)
	}

	// Step 2: Enable workflow orchestration
	if !isJSON {
		fmt.Println("\n⚙️  Enabling workflow orchestration...")
	}
	workflow.DefaultOrchestrator.SetEnabled(true)
	if !isJSON {
		fmt.Println("✓ Automatic job progression enabled (scan → pack → daggen → deals)")
	} else {
		// Include orchestration state in JSON output
		result := OnboardResult{
			Success: true,
			// WorkflowOrchestrationEnabled will be set to true in final output
		}
		_ = result // Use later in final output
	}

	// Step 3: Start workers if requested
	var workerManager *workermanager.WorkerManager
	workersCount := 0
	if c.Bool("start-workers") {
		if !isJSON {
			fmt.Println("\n👷 Starting managed workers...")
		}
		workerManager, err = startManagedWorkers(ctx, db, c.Int("max-workers"))
		if err != nil {
			return outputJSONError("failed to start workers", err)
		}
		workersCount = c.Int("max-workers")
		if !isJSON {
			fmt.Printf("✓ Started %d managed workers\n", workersCount)
		}
	}

	// Step 4: Start scanning
	if !isJSON {
		fmt.Println("\n🔍 Starting initial scanning...")
	}
	err = startScanningForPreparation(ctx, db, prep)
	if err != nil {
		return outputJSONError("failed to start scanning", err)
	}
	if !isJSON {
		fmt.Println("✓ Scanning started for all source attachments")
	}

	// Step 5: Monitor progress if requested
	if c.Bool("wait-for-completion") {
		if !isJSON {
			fmt.Println("\n📊 Monitoring progress...")
		}
		err = monitorProgress(ctx, db, prep, c.Duration("timeout"))
		if err != nil {
			return outputJSONError("monitoring failed", err)
		}

		// Only cleanup workers after completion monitoring finishes successfully
		if workerManager != nil {
			if !isJSON {
				fmt.Println("\n🧹 Cleaning up workers...")
			}
			err = workerManager.Stop(ctx)
			if err != nil {
				if !isJSON {
					fmt.Printf("⚠ Warning: failed to stop workers cleanly: %v\n", err)
				}
			}
		}
	} else if workerManager != nil {
		// When not waiting for completion, leave workers running to process jobs
		if !isJSON {
			fmt.Println("\n✅ Workers will continue running to process jobs")
			fmt.Println("💡 Use --wait-for-completion to monitor progress and stop workers when done")
		}
	}

	// Output results
	if isJSON {
		// Prepare next steps
		nextSteps := []string{
			"Monitor progress: singularity prep status " + prep.Name,
			"Check jobs: singularity job list",
		}
		if c.Bool("start-workers") {
			if c.Bool("wait-for-completion") {
				nextSteps = append(nextSteps, "Workers have been stopped after completion")
			} else {
				nextSteps = append(nextSteps, "Workers are running and will process jobs automatically")
			}
		} else {
			nextSteps = append(nextSteps, "Start workers: singularity run unified")
		}

		result := OnboardResult{
			Success:       true,
			PreparationID: uint32(prep.ID),
			Name:          prep.Name,
			SourcePaths:   c.StringSlice("source"),
			OutputPaths:   c.StringSlice("output"),
			AutoDeals:     c.Bool("auto-create-deals"),
			WorkersCount:  workersCount,
			NextSteps:     nextSteps,
		}
		data, err := json.Marshal(result)
		if err != nil {
			return errors.Wrap(err, "failed to marshal JSON result")
		}
		fmt.Println(string(data))
	} else if !c.Bool("wait-for-completion") {
		fmt.Println("\n✅ Onboarding initiated successfully!")
		fmt.Println("\n📝 Next steps:")
		fmt.Println("   • Monitor progress: singularity prep status", prep.Name)
		fmt.Println("   • Check jobs: singularity job list")
		if c.Bool("start-workers") {
			fmt.Println("   • Workers are running and will process jobs automatically")
		} else {
			fmt.Println("   • Start workers: singularity run unified")
		}
	}

	return nil
}

// createPreparationForOnboarding creates a preparation with all onboarding settings
func createPreparationForOnboarding(ctx context.Context, db *gorm.DB, c *cli.Context) (*model.Preparation, error) {
	// Log warning for insecure client configurations
	logInsecureClientConfigWarning(c)

	// Convert source paths to storage names (create if needed)
	var sourceStorages []string
	for i, sourcePath := range c.StringSlice("source") {
		storageName := c.String("source-name")
		if storageName == "" {
			storageName = fmt.Sprintf("source-%s-%d", util.RandomName(), time.Now().Unix())
		} else if len(c.StringSlice("source")) > 1 {
			storageName = fmt.Sprintf("%s-%d", storageName, i)
		}
		storage, err := createStorageIfNotExistWithConfig(ctx, db, sourcePath, c.String("source-type"), c.String("source-provider"), c, storageName, "source")
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create source storage for %s", sourcePath)
		}
		sourceStorages = append(sourceStorages, storage.Name)
	}

	// Convert output paths to storage names (create if needed)
	var outputStorages []string
	for i, outputPath := range c.StringSlice("output") {
		storageName := c.String("output-name")
		if storageName == "" {
			storageName = fmt.Sprintf("output-%s-%d", util.RandomName(), time.Now().Unix())
		} else if len(c.StringSlice("output")) > 1 {
			storageName = fmt.Sprintf("%s-%d", storageName, i)
		}
		storage, err := createStorageIfNotExistWithConfig(ctx, db, outputPath, c.String("output-type"), c.String("output-provider"), c, storageName, "output")
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create output storage for %s", outputPath)
		}
		outputStorages = append(outputStorages, storage.Name)
	}

	// Create preparation
	prep, err := dataprep.Default.CreatePreparationHandler(ctx, db, dataprep.CreateRequest{
		Name:             c.String("name"),
		SourceStorages:   sourceStorages,
		OutputStorages:   outputStorages,
		MaxSizeStr:       c.String("max-size"),
		NoDag:            c.Bool("no-dag"),
		AutoCreateDeals:  c.Bool("auto-create-deals"),
		DealTemplate:     c.String("deal-template-id"),
		WalletValidation: c.Bool("wallet-validation"),
		SPValidation:     c.Bool("sp-validation"),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return prep, nil
}

// startManagedWorkers starts the worker manager for automatic job processing
func startManagedWorkers(ctx context.Context, db *gorm.DB, maxWorkers int) (*workermanager.WorkerManager, error) {
	config := workermanager.ManagerConfig{
		CheckInterval:      10 * time.Second,
		MinWorkers:         1,
		MaxWorkers:         maxWorkers,
		ScaleUpThreshold:   3,
		ScaleDownThreshold: 1,
		WorkerIdleTimeout:  2 * time.Minute,
		AutoScaling:        true,
		ScanWorkerRatio:    0.3,
		PackWorkerRatio:    0.5,
		DagGenWorkerRatio:  0.2,
	}

	manager := workermanager.NewWorkerManager(db, config)
	err := manager.Start(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return manager, nil
}

// startScanningForPreparation starts scanning for all source attachments
func startScanningForPreparation(ctx context.Context, db *gorm.DB, prep *model.Preparation) error {
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

// monitorProgress monitors the progress of the onboarding workflow
func monitorProgress(ctx context.Context, db *gorm.DB, prep *model.Preparation, timeout time.Duration) error {
	fmt.Println("Monitoring job progress (Ctrl+C to stop monitoring)...")

	var monitorCtx context.Context
	var cancel context.CancelFunc

	if timeout > 0 {
		monitorCtx, cancel = context.WithTimeout(ctx, timeout)
		fmt.Printf("⏰ Timeout set to %v\n", timeout)
	} else {
		monitorCtx, cancel = context.WithCancel(ctx)
	}
	defer cancel()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	lastStatus := ""

	for {
		select {
		case <-monitorCtx.Done():
			if errors.Is(monitorCtx.Err(), context.DeadlineExceeded) {
				fmt.Printf("⏰ Monitoring timeout reached\n")
				return nil
			}
			fmt.Printf("\n🛑 Monitoring stopped\n")
			return nil

		case <-ticker.C:
			status, complete, err := getPreparationStatus(ctx, db, prep)
			if err != nil {
				fmt.Printf("⚠ Error checking status: %v\n", err)
				continue
			}

			if status != lastStatus {
				fmt.Printf("📊 %s\n", status)
				lastStatus = status
			}

			if complete {
				fmt.Printf("🎉 Onboarding completed successfully!\n")
				return nil
			}
		}
	}
}

// getPreparationStatus returns the current status of the preparation
func getPreparationStatus(ctx context.Context, db *gorm.DB, prep *model.Preparation) (string, bool, error) {
	// Get job counts by type and state
	type JobCount struct {
		Type  string `json:"type"`
		State string `json:"state"`
		Count int64  `json:"count"`
	}

	var jobCounts []JobCount
	err := db.WithContext(ctx).Model(&model.Job{}).
		Select("type, state, count(*) as count").
		Joins("JOIN source_attachments ON jobs.attachment_id = source_attachments.id").
		Where("source_attachments.preparation_id = ?", prep.ID).
		Group("type, state").
		Find(&jobCounts).Error
	if err != nil {
		return "", false, errors.WithStack(err)
	}

	// Analyze status
	jobStats := make(map[string]map[string]int64)
	totalJobs := int64(0)
	completeJobs := int64(0)

	for _, jc := range jobCounts {
		if jobStats[jc.Type] == nil {
			jobStats[jc.Type] = make(map[string]int64)
		}
		jobStats[jc.Type][jc.State] = jc.Count
		totalJobs += jc.Count
		if jc.State == "complete" {
			completeJobs += jc.Count
		}
	}

	if totalJobs == 0 {
		return "No jobs created yet", false, nil
	}

	// Check for deal schedules
	var scheduleCount int64
	err = db.WithContext(ctx).Model(&model.Schedule{}).
		Where("preparation_id = ?", prep.ID).Count(&scheduleCount).Error
	if err != nil {
		return "", false, errors.WithStack(err)
	}

	// Build status message
	status := fmt.Sprintf("Progress: %d/%d jobs complete", completeJobs, totalJobs)

	if scan := jobStats["scan"]; len(scan) > 0 {
		status += fmt.Sprintf(" | Scan: %d ready, %d processing, %d complete",
			scan["ready"], scan["processing"], scan["complete"])
	}

	if pack := jobStats["pack"]; len(pack) > 0 {
		status += fmt.Sprintf(" | Pack: %d ready, %d processing, %d complete",
			pack["ready"], pack["processing"], pack["complete"])
	}

	if daggen := jobStats["daggen"]; len(daggen) > 0 {
		status += fmt.Sprintf(" | DagGen: %d ready, %d processing, %d complete",
			daggen["ready"], daggen["processing"], daggen["complete"])
	}

	if scheduleCount > 0 {
		status += fmt.Sprintf(" | Deals: %d schedule(s) created", scheduleCount)
		return status, true, nil // Complete when deals are created
	}

	return status, false, nil
}

// capitalizeFirst capitalizes the first letter of a string
func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// generateDynamicStorageFlags creates CLI flags for all supported storage backends
func generateDynamicStorageFlags() []cli.Flag {
	var flags []cli.Flag
	seenFlags := make(map[string]bool)

	// Add flags for each supported backend
	for _, backend := range storagesystem.Backends {
		for _, providerOptions := range backend.ProviderOptions {
			for _, option := range providerOptions.Options {
				// Convert option name to CLI flag format
				flagName := strings.ReplaceAll(option.Name, "_", "-")

				// Create flag for each storage type context (source/output)
				for _, context := range []string{"source", "output"} {
					fullFlagName := context + "-" + backend.Prefix + "-" + flagName

					// Skip if we've already seen this flag name
					if seenFlags[fullFlagName] {
						continue
					}
					seenFlags[fullFlagName] = true

					// Create better category names for backend-specific flags
					var category string
					if context == "source" {
						category = fmt.Sprintf("Source Storage (%s)", strings.ToUpper(backend.Prefix))
					} else {
						category = fmt.Sprintf("Output Storage (%s)", strings.ToUpper(backend.Prefix))
					}

					// Clean up the usage text to be more concise
					usage := strings.Split(option.Help, "\n")[0]
					if usage == "" {
						usage = fmt.Sprintf("%s configuration for %s", option.Name, backend.Name)
					}

					// Create flag based on option type
					switch (*fs.Option)(&option).Type() {
					case "bool":
						flags = append(flags, &cli.BoolFlag{
							Name:     fullFlagName,
							Usage:    usage,
							Category: category,
						})
					case "string":
						flags = append(flags, &cli.StringFlag{
							Name:     fullFlagName,
							Usage:    usage,
							Category: category,
						})
					case "int":
						flags = append(flags, &cli.IntFlag{
							Name:     fullFlagName,
							Usage:    usage,
							Category: category,
						})
					case "SizeSuffix":
						flags = append(flags, &cli.StringFlag{
							Name:     fullFlagName,
							Usage:    usage,
							Category: category,
						})
					case "Duration":
						flags = append(flags, &cli.DurationFlag{
							Name:     fullFlagName,
							Usage:    usage,
							Category: category,
						})
					default:
						flags = append(flags, &cli.StringFlag{
							Name:     fullFlagName,
							Usage:    usage,
							Category: category,
						})
					}
				}
			}
		}
	}

	return flags
}

// getDefaultProvider returns the default provider for a storage type
func getDefaultProvider(storageType string) string {
	switch storageType {
	case "s3":
		return "aws"
	case "gcs":
		return "google"
	case "local":
		return "local"
	default:
		return ""
	}
}

// validateOnboardInputs validates CLI inputs for onboard command
func validateOnboardInputs(c *cli.Context) error {
	// Required fields validation
	if c.String("name") == "" {
		return errors.New("preparation name is required (--name)")
	}

	// Source and output validation
	sourcePaths := c.StringSlice("source")

	if len(sourcePaths) == 0 {
		return errors.New("at least one source path is required (--source)")
	}

	// Basic storage type validation - detailed validation is handled by the storage handler
	sourceType := c.String("source-type")
	outputType := c.String("output-type")

	// Validate storage types are supported using BackendMap
	if sourceType != "" && sourceType != "local" {
		if _, exists := storagesystem.BackendMap[sourceType]; !exists {
			return errors.Errorf("source storage type '%s' is not supported", sourceType)
		}
	}
	if outputType != "" && outputType != "local" {
		if _, exists := storagesystem.BackendMap[outputType]; !exists {
			return errors.Errorf("output storage type '%s' is not supported", outputType)
		}
	}

	// Validate custom storage configurations if provided
	if sourceConfig := c.String("source-config"); sourceConfig != "" {
		if err := validateStorageConfig(sourceConfig, "source"); err != nil {
			return err
		}
	}
	if outputConfig := c.String("output-config"); outputConfig != "" {
		if err := validateStorageConfig(outputConfig, "output"); err != nil {
			return err
		}
	}

	// Auto-deal validation
	if c.Bool("auto-create-deals") {
		// Deal template is required when auto-create-deals is enabled
		if c.String("deal-template-id") == "" {
			return errors.New("deal template ID is required when auto-create-deals is enabled (--deal-template-id)")
		}
	}

	// Validate max-size format if provided
	if maxSize := c.String("max-size"); maxSize != "" {
		if _, err := humanize.ParseBytes(maxSize); err != nil {
			return errors.Wrapf(err, "invalid max-size format")
		}
	}

	// Validate worker count
	if maxWorkers := c.Int("max-workers"); maxWorkers < 1 {
		return errors.New("max workers must be at least 1")
	}

	return nil
}

// getOnboardClientConfig parses client configuration from CLI flags
func getOnboardClientConfig(c *cli.Context) (*model.ClientConfig, error) {
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

// validateStorageConfig validates that a storage configuration is valid JSON
func validateStorageConfig(configStr string, prefix string) error {
	var config map[string]string
	if err := json.Unmarshal([]byte(configStr), &config); err != nil {
		return errors.Wrapf(err, "invalid JSON format for %s-config", prefix)
	}

	// Validate that keys and values are non-empty strings
	for key, value := range config {
		if key == "" {
			return errors.Errorf("%s-config cannot have empty keys", prefix)
		}
		if value == "" {
			return errors.Errorf("%s-config cannot have empty values", prefix)
		}
	}

	return nil
}

// logInsecureClientConfigWarning logs warnings for insecure client configurations
func logInsecureClientConfigWarning(c *cli.Context) {
	if c.Bool("client-insecure-skip-verify") {
		fmt.Printf("⚠️  WARNING: Insecure SSL/TLS configuration detected. Certificate verification is disabled.\n")
		fmt.Printf("   This may expose your data to man-in-the-middle attacks.\n")
		fmt.Printf("   Consider using proper certificates in production environments.\n\n")
	}
}

// getCustomStorageConfig returns custom storage configuration for the given storage type
func getCustomStorageConfig(c *cli.Context, storageContext string) map[string]string {
	var configFlag string

	// Determine which config flag to use based on storage context
	switch storageContext {
	case "source":
		configFlag = c.String("source-config")
	case "output":
		configFlag = c.String("output-config")
	}

	if configFlag == "" {
		return nil
	}

	var config map[string]string
	if err := json.Unmarshal([]byte(configFlag), &config); err != nil {
		// This should have been caught in validation, but handle gracefully
		return nil
	}

	return config
}

// createStorageIfNotExistWithConfig creates storage with context-aware custom configuration
func createStorageIfNotExistWithConfig(ctx context.Context, db *gorm.DB, path, storageType, provider string, c *cli.Context, storageName, storageContext string) (*model.Storage, error) {
	// Check if storage already exists for this path and type
	var existing model.Storage
	err := db.WithContext(ctx).Where("type = ? AND path = ?", storageType, path).First(&existing).Error
	if err == nil {
		return &existing, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "failed to check existing storage")
	}

	// Build storage configuration from CLI flags using the same pattern as storage create
	config := make(map[string]string)

	// Get backend configuration to understand what flags to look for
	backend, exists := storagesystem.BackendMap[storageType]
	if !exists {
		return nil, errors.Errorf("storage type %s is not supported", storageType)
	}

	// Parse configuration for all provider options with the storageContext prefix
	for _, providerOptions := range backend.ProviderOptions {
		for _, option := range providerOptions.Options {
			// Convert option name to CLI flag format
			flagName := strings.ReplaceAll(option.Name, "_", "-")
			fullFlagName := storageContext + "-" + backend.Prefix + "-" + flagName

			// Check if the flag is set and extract value based on type
			if c.IsSet(fullFlagName) {
				switch (*fs.Option)(&option).Type() {
				case "bool":
					if c.Bool(fullFlagName) {
						config[option.Name] = "true"
					}
				case "string":
					if value := c.String(fullFlagName); value != "" {
						config[option.Name] = value
					}
				case "int":
					config[option.Name] = strconv.Itoa(c.Int(fullFlagName))
				case "SizeSuffix":
					if value := c.String(fullFlagName); value != "" {
						config[option.Name] = value
					}
				case "Duration":
					config[option.Name] = c.Duration(fullFlagName).String()
				default:
					if value := c.String(fullFlagName); value != "" {
						config[option.Name] = value
					}
				}
			}
		}
	}

	// Merge custom config if provided via JSON config flag
	if customConfig := getCustomStorageConfig(c, storageContext); customConfig != nil {
		for key, value := range customConfig {
			config[key] = value
		}
	}

	// Set default provider if not specified
	if provider == "" {
		provider = getDefaultProvider(storageType)
	}

	// Get client configuration from flags
	clientConfig, err := getOnboardClientConfig(c)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse client configuration")
	}

	// Use the storage handler to create new storage - it will handle validation, defaults, and connectivity testing
	storageHandler := storageHandlers.Default
	request := storageHandlers.CreateRequest{
		Name:         storageName,
		Path:         path,
		Provider:     provider,
		Config:       config,
		ClientConfig: *clientConfig,
	}

	storage, err := storageHandler.CreateStorageHandler(ctx, db, storageType, request)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create storage '%s'", storageName)
	}

	return storage, nil
}

// validateDealTemplateExists validates that the deal template ID exists in the database
// and checks basic fields like provider format, pricing, and duration
func validateDealTemplateExists(ctx context.Context, db *gorm.DB, templateID string) error {
	if templateID == "" {
		return nil // No template specified, validation not needed
	}

	// Check if template exists by ID or name
	var template model.DealTemplate
	err := db.WithContext(ctx).Where("id = ? OR name = ?", templateID, templateID).First(&template).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Errorf("deal template '%s' not found. Please create the template first using 'singularity deal-schedule-template create'", templateID)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	// Validate basic template fields
	if template.DealConfig.DealProvider == "" {
		return errors.Errorf("deal template '%s' has no provider specified", templateID)
	}

	// Validate provider format (should be like f01000)
	if len(template.DealConfig.DealProvider) < 4 || template.DealConfig.DealProvider[:1] != "f" {
		return errors.Errorf("deal template '%s' has invalid provider format '%s' (expected format: f01000)", templateID, template.DealConfig.DealProvider)
	}

	// Validate duration is set
	if template.DealConfig.DealDuration <= 0 {
		return errors.Errorf("deal template '%s' has invalid duration %v (must be > 0)", templateID, template.DealConfig.DealDuration)
	}

	// Validate pricing is non-negative
	if template.DealConfig.DealPricePerGb < 0 || template.DealConfig.DealPricePerGbEpoch < 0 || template.DealConfig.DealPricePerDeal < 0 {
		return errors.Errorf("deal template '%s' has negative pricing values", templateID)
	}

	return nil
}
