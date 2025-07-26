package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"slices"
	"sort"
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
	"github.com/data-preservation-programs/singularity/service/errorlog"
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

SUPPORTED STORAGE BACKENDS:
The onboard command supports all 40+ storage backends available in the storage create command, including:
  • Cloud providers: S3, GCS, Azure Blob, Dropbox, OneDrive, Box, etc.
  • Protocol-based: FTP, SFTP, WebDAV, HTTP, SMB, etc.
  • Specialized: Storj, Sia, HDFS, Internet Archive, etc.

COMMON USAGE PATTERNS:
  • Basic local data onboarding:
    singularity onboard --name "my-dataset" --source "/path/to/data" --deal-template-id "1"

  • S3 to local with custom output:
    singularity onboard --name "s3-data" \
      --source "s3://bucket/data" --source-type "s3" \
      --source-s3-region us-east-1 --source-s3-access-key-id "key" \
      --output "/mnt/storage/cars" \
      --deal-template-id "template1"

  • Multiple sources with monitoring:
    singularity onboard --name "multi-source" \
      --source "/data1" --source "/data2" \
      --wait-for-completion --max-workers 5 \
      --deal-template-id "prod-template"

  • Cloud-to-cloud transfer:
    singularity onboard --name "gcs-to-s3" \
      --source-type "gcs" --source "gs://source-bucket/data" \
      --output-type "s3" --output "s3://dest-bucket/cars" \
      --deal-template-id "cloud-template"

GETTING HELP:
  • Use --help-examples to see more detailed examples
  • Use --help-backends to list all available storage backends
  • Use --help-backend=<type> to see only flags for specific backends (e.g., s3, gcs)
  • Use --help-all to see all available flags including backend-specific options

BACKEND-SPECIFIC OPTIONS:
Each storage backend has its own configuration options. For example:
  • S3: --source-s3-region, --source-s3-access-key-id, --source-s3-secret-access-key
  • GCS: --source-gcs-project-number, --source-gcs-service-account-file
  • Azure: --source-azureblob-account, --source-azureblob-key

Use --help-backend=<type> to see all available options for a specific backend.

NOTE: All backends supported by 'storage create' are also supported by 'onboard'.
      Use SINGULARITY_LIMIT_BACKENDS=true to show only common backends in help.`,
	Flags: []cli.Flag{
		// Required Options
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "Name for the preparation",
			Required: false, // Will be validated in action function
			Category: "Required Options",
		},
		&cli.StringSliceFlag{
			Name:     "source",
			Aliases:  []string{"s"},
			Usage:    "Source path(s) to onboard (local paths or remote URLs like s3://bucket/path)",
			Required: false, // Will be validated in action function
			Category: "Required Options",
		},
		&cli.StringFlag{
			Name:     "deal-template-id",
			Aliases:  []string{"t"},
			Usage:    "Deal template ID to use for deal configuration (required when auto-create-deals is enabled)",
			Category: "Required Options",
		},

		// Data Source Configuration
		&cli.StringFlag{
			Name:     "source-type",
			Usage:    "Source storage type (local, s3, gcs, azure, etc.)",
			Value:    "local",
			Category: "Data Source Configuration",
		},
		&cli.StringFlag{
			Name:     "source-provider",
			Usage:    "Source storage provider (for s3: aws, minio, wasabi, etc.)",
			Category: "Data Source Configuration",
		},
		&cli.StringFlag{
			Name:     "source-name",
			Usage:    "Custom name for source storage (auto-generated if not provided)",
			Category: "Data Source Configuration",
		},
		&cli.StringFlag{
			Name:     "source-config",
			Usage:    "Source storage configuration in JSON format (key-value pairs)",
			Category: "Data Source Configuration",
		},

		// Data Output Configuration
		&cli.StringSliceFlag{
			Name:     "output",
			Aliases:  []string{"o"},
			Usage:    "Output path(s) for CAR files (local paths or remote URLs like s3://bucket/path)",
			Category: "Data Output Configuration",
		},
		&cli.StringFlag{
			Name:     "output-type",
			Usage:    "Output storage type (local, s3, gcs, azure, etc.)",
			Value:    "local",
			Category: "Data Output Configuration",
		},
		&cli.StringFlag{
			Name:     "output-provider",
			Usage:    "Output storage provider",
			Category: "Data Output Configuration",
		},
		&cli.StringFlag{
			Name:     "output-name",
			Usage:    "Custom name for output storage (auto-generated if not provided)",
			Category: "Data Output Configuration",
		},
		&cli.StringFlag{
			Name:     "output-config",
			Usage:    "Output storage configuration in JSON format (key-value pairs)",
			Category: "Data Output Configuration",
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
			Category: "Monitoring & Progress",
		},
		&cli.DurationFlag{
			Name:     "timeout",
			Usage:    "Timeout for waiting for completion (0 = no timeout)",
			Value:    0,
			Category: "Monitoring & Progress",
		},
		&cli.BoolFlag{
			Name:     "json",
			Usage:    "Output result in JSON format for automation",
			Category: "Output & Formatting",
		},

		// Validation Options
		&cli.BoolFlag{
			Name:     "wallet-validation",
			Usage:    "Enable wallet balance validation",
			Category: "Validation & Security",
		},
		&cli.BoolFlag{
			Name:     "sp-validation",
			Usage:    "Enable storage provider validation",
			Category: "Validation & Security",
		},

		// Help Options
		&cli.BoolFlag{
			Name:     "help-all",
			Usage:    "Show all available options including all backend-specific flags",
			Category: "Help Options",
		},
		&cli.BoolFlag{
			Name:     "help-backends",
			Usage:    "List all available storage backends",
			Category: "Help Options",
		},
		&cli.StringFlag{
			Name:     "help-backend",
			Usage:    "Show options for specific backend (e.g., s3, gcs, local)",
			Category: "Help Options",
		},
		&cli.BoolFlag{
			Name:     "help-examples",
			Usage:    "Show common usage examples",
			Category: "Help Options",
		},
		&cli.BoolFlag{
			Name:     "help-json",
			Usage:    "Output help in JSON format for machine processing",
			Category: "Help Options",
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

	// Set the custom help template
	OnboardCmd.CustomHelpTemplate = getCustomHelpTemplate()

	// Set custom help printer to handle special help flags
	OnboardCmd.HelpName = "singularity onboard"
	OnboardCmd.SkipFlagParsing = false
}

// Action function for the onboard command
func onboardAction(c *cli.Context) error {
	// Check for special help flags first
	if c.Bool("help-backends") {
		return showBackendsList(c)
	}
	if c.String("help-backend") != "" {
		return showBackendHelp(c, c.String("help-backend"))
	}
	if c.Bool("help-examples") {
		return showExamples(c)
	}
	if c.Bool("help-json") {
		return showHelpJSON(c)
	}

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

	// Validate CLI inputs before proceeding (skip for help flags)
	if err := validateOnboardInputs(c); err != nil {
		return outputJSONError("input validation failed", err)
	}

	// Initialize database
	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return outputJSONError("failed to initialize database", err)
	}
	defer func() { _ = closer.Close() }()

	// Initialize error logging
	errorlog.Init(db)

	ctx := c.Context

	// Log onboarding start
	errorlog.LogOnboardEvent(model.ErrorLevelInfo, "", "onboard_started",
		"Onboarding started for preparation: "+c.String("name"), nil,
		map[string]interface{}{
			"sources":       c.StringSlice("source"),
			"outputs":       c.StringSlice("output"),
			"auto_deals":    c.Bool("auto-create-deals"),
			"deal_template": c.String("deal-template-id"),
		})

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
		errorlog.LogOnboardEvent(model.ErrorLevelError, "", "preparation_creation_failed",
			"Failed to create preparation during onboarding", err,
			map[string]interface{}{
				"preparation_name": c.String("name"),
			})
		return outputJSONError("failed to create preparation", err)
	}

	// Log successful preparation creation
	prepIDStr := strconv.FormatUint(uint64(prep.ID), 10)
	errorlog.LogOnboardEvent(model.ErrorLevelInfo, prepIDStr, "preparation_created",
		"Preparation created successfully: "+prep.Name, nil,
		map[string]interface{}{
			"preparation_id":   prep.ID,
			"preparation_name": prep.Name,
		})

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

// generateDynamicStorageFlags creates CLI flags for all or specific storage backends
func generateDynamicStorageFlags() []cli.Flag {
	// Check for environment variable to limit backends (for backwards compatibility)
	limitBackends := os.Getenv("SINGULARITY_LIMIT_BACKENDS") == "true"

	// Default: show all backends (like storage create command)
	if !limitBackends {
		return generateAllStorageFlags()
	}

	// Fallback: limited backends (only when explicitly requested)
	defaultBackends := map[string]bool{
		"local":     true,
		"s3":        true,
		"gcs":       true,
		"azureblob": true,
	}

	// Generate flags for limited backends only
	return generateSelectiveStorageFlags(defaultBackends)
}

// generateAllStorageFlags creates CLI flags for all supported storage backends
func generateAllStorageFlags() []cli.Flag {
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

// generateSelectiveStorageFlags creates CLI flags for specific storage backends
func generateSelectiveStorageFlags(allowedBackends map[string]bool) []cli.Flag {
	var flags []cli.Flag
	seenFlags := make(map[string]bool)

	// Add flags for allowed backends only
	for _, backend := range storagesystem.Backends {
		// Skip if this backend is not in the allowed list
		if !allowedBackends[backend.Prefix] {
			continue
		}

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
		val := c.Duration("client-connect-timeout")
		config.ConnectTimeout = ptr.Of(int64(val))
	}
	if c.IsSet("client-timeout") {
		val := c.Duration("client-timeout")
		config.Timeout = ptr.Of(int64(val))
	}
	if c.IsSet("client-expect-continue-timeout") {
		val := c.Duration("client-expect-continue-timeout")
		config.ExpectContinueTimeout = ptr.Of(int64(val))
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
		val := c.Duration("client-retry-delay")
		config.RetryDelay = ptr.Of(int64(val))
	}
	if c.IsSet("client-retry-backoff") {
		val := c.Duration("client-retry-backoff")
		config.RetryBackoff = ptr.Of(int64(val))
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

// showBackendsList displays all available storage backends
func showBackendsList(c *cli.Context) error {
	isJSON := c.Bool("json")

	if isJSON {
		type BackendInfo struct {
			Name        string   `json:"name"`
			Prefix      string   `json:"prefix"`
			Description string   `json:"description"`
			Providers   []string `json:"providers"`
		}

		var backends []BackendInfo
		for _, backend := range storagesystem.Backends {
			var providers []string
			for _, provider := range backend.ProviderOptions {
				providers = append(providers, provider.Provider)
			}
			backends = append(backends, BackendInfo{
				Name:        backend.Name,
				Prefix:      backend.Prefix,
				Description: backend.Description,
				Providers:   providers,
			})
		}

		data, err := json.MarshalIndent(backends, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(data))
		return nil
	}

	fmt.Println("Available Storage Backends:")
	fmt.Println("")
	for _, backend := range storagesystem.Backends {
		fmt.Printf("  %-20s %s\n", backend.Prefix, backend.Name)
		if backend.Description != "" {
			fmt.Printf("  %-20s %s\n", "", backend.Description)
		}
		if len(backend.ProviderOptions) > 0 {
			var providers []string
			for _, provider := range backend.ProviderOptions {
				providers = append(providers, provider.Provider)
			}
			fmt.Printf("  %-20s Providers: %s\n", "", strings.Join(providers, ", "))
		}
		fmt.Println("")
	}
	fmt.Println("Use --help-backend=<prefix> to see specific backend options.")
	return nil
}

// showBackendHelp displays help for a specific backend with filtered flags
func showBackendHelp(c *cli.Context, backendPrefix string) error {
	backend, exists := storagesystem.BackendMap[backendPrefix]
	if !exists {
		return errors.Wrapf(storagesystem.ErrBackendNotFound, "backend '%s' not found. Use --help-backends to see available backends", backendPrefix)
	}

	isJSON := c.Bool("json")

	if isJSON {
		return showBackendHelpJSON(backend)
	}

	// Show command usage and description
	fmt.Printf("Usage: %s [flags]\n", OnboardCmd.Name)
	fmt.Printf("Complete data onboarding workflow for %s storage\n\n", backend.Name)

	// Show backend-specific information
	fmt.Printf("Backend: %s (%s)\n", backend.Name, backend.Prefix)
	if backend.Description != "" {
		fmt.Printf("Description: %s\n", backend.Description)
	}
	fmt.Println("")

	// Show available flags organized by category
	flags := OnboardCmd.Flags
	filteredFlags := filterFlagsForBackend(flags, backendPrefix)

	// Group flags by category
	categories := make(map[string][]cli.Flag)
	for _, flag := range filteredFlags {
		category := getCategoryForFlag(flag)
		categories[category] = append(categories[category], flag)
	}

	// Display flags by category
	categoryOrder := []string{
		"Required Options",
		"Data Source Configuration",
		"Data Output Configuration",
		"Data Preparation",
		"Deal Configuration",
		"Worker Management",
		"Monitoring & Progress",
		"Validation & Security",
		"Output & Formatting",
		fmt.Sprintf("Source Storage (%s)", strings.ToUpper(backend.Prefix)),
		fmt.Sprintf("Output Storage (%s)", strings.ToUpper(backend.Prefix)),
		"Client Config",
		"Retry Strategy",
	}

	for _, category := range categoryOrder {
		if flags, exists := categories[category]; exists && len(flags) > 0 {
			fmt.Printf("%s:\n", category)
			for _, flag := range flags {
				fmt.Printf("  %s\n", formatFlagUsage(flag))
			}
			fmt.Println("")
		}
	}

	// Show provider-specific information
	for _, provider := range backend.ProviderOptions {
		if len(backend.ProviderOptions) > 1 {
			fmt.Printf("Provider: %s\n", provider.Provider)
			if provider.ProviderDescription != "" {
				fmt.Printf("  %s\n", provider.ProviderDescription)
			}
			fmt.Println("")
		}
	}

	return nil
}

// showBackendHelpJSON displays backend help in JSON format
func showBackendHelpJSON(backend storagesystem.Backend) error {
	type OptionInfo struct {
		Name         string `json:"name"`
		Type         string `json:"type"`
		Help         string `json:"help"`
		Required     bool   `json:"required"`
		DefaultValue string `json:"default,omitempty"`
	}

	type ProviderInfo struct {
		Provider    string       `json:"provider"`
		Description string       `json:"description"`
		Options     []OptionInfo `json:"options"`
	}

	type BackendHelpInfo struct {
		Name        string         `json:"name"`
		Prefix      string         `json:"prefix"`
		Description string         `json:"description"`
		Providers   []ProviderInfo `json:"providers"`
	}

	var providers []ProviderInfo
	for _, provider := range backend.ProviderOptions {
		var options []OptionInfo
		for _, option := range provider.Options {
			options = append(options, OptionInfo{
				Name:         option.Name,
				Type:         (*fs.Option)(&option).Type(),
				Help:         option.Help,
				Required:     option.Required,
				DefaultValue: fmt.Sprintf("%v", option.Default),
			})
		}
		providers = append(providers, ProviderInfo{
			Provider:    provider.Provider,
			Description: provider.ProviderDescription,
			Options:     options,
		})
	}

	backendHelp := BackendHelpInfo{
		Name:        backend.Name,
		Prefix:      backend.Prefix,
		Description: backend.Description,
		Providers:   providers,
	}

	data, err := json.MarshalIndent(backendHelp, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

// filterFlagsForBackend filters flags to show only those relevant to the specified backend
func filterFlagsForBackend(flags []cli.Flag, backendPrefix string) []cli.Flag {
	var filteredFlags []cli.Flag

	for _, flag := range flags {
		flagName := flag.Names()[0]

		// Always include common flags
		if isCommonFlag(flagName) {
			filteredFlags = append(filteredFlags, flag)
			continue
		}

		// Include backend-specific flags
		if isBackendSpecificFlag(flagName, backendPrefix) {
			filteredFlags = append(filteredFlags, flag)
			continue
		}
	}

	return filteredFlags
}

// isCommonFlag checks if a flag is a common flag (not backend-specific)
func isCommonFlag(flagName string) bool {
	commonFlags := []string{
		"name", "source", "deal-template-id", "source-type", "source-provider", "source-name", "source-config",
		"output", "output-type", "output-provider", "output-name", "output-config", "max-size", "no-dag",
		"auto-create-deals", "start-workers", "max-workers", "wait-for-completion", "timeout", "json",
		"wallet-validation", "sp-validation", "help-all", "help-backends", "help-backend", "help-examples", "help-json",
		"client-connect-timeout", "client-timeout", "client-expect-continue-timeout", "client-insecure-skip-verify",
		"client-no-gzip", "client-user-agent", "client-ca-cert", "client-cert", "client-key", "client-header",
		"client-retry-max", "client-retry-delay", "client-retry-backoff", "client-retry-backoff-exp",
		"client-skip-inaccessible", "client-low-level-retries", "client-use-server-mod-time", "client-scan-concurrency",
	}

	return slices.Contains(commonFlags, flagName)
}

// isBackendSpecificFlag checks if a flag is specific to the given backend
func isBackendSpecificFlag(flagName, backendPrefix string) bool {
	sourcePrefix := "source-" + backendPrefix + "-"
	outputPrefix := "output-" + backendPrefix + "-"

	return strings.HasPrefix(flagName, sourcePrefix) || strings.HasPrefix(flagName, outputPrefix)
}

// getCategoryForFlag returns the category for a flag
func getCategoryForFlag(flag cli.Flag) string {
	switch f := flag.(type) {
	case *cli.StringFlag:
		return f.Category
	case *cli.BoolFlag:
		return f.Category
	case *cli.IntFlag:
		return f.Category
	case *cli.DurationFlag:
		return f.Category
	case *cli.StringSliceFlag:
		return f.Category
	case *cli.Float64Flag:
		return f.Category
	case *cli.PathFlag:
		return f.Category
	default:
		return "Other"
	}
}

// formatFlagUsage returns formatted usage string for a flag
func formatFlagUsage(flag cli.Flag) string {
	names := flag.Names()
	usage := getUsageForFlag(flag)

	var nameStr string
	if len(names) > 1 {
		nameStr = fmt.Sprintf("--%s, -%s", names[0], names[1])
	} else {
		nameStr = "--" + names[0]
	}

	return fmt.Sprintf("%-30s %s", nameStr, usage)
}

// getUsageForFlag returns the usage string for a flag
func getUsageForFlag(flag cli.Flag) string {
	switch f := flag.(type) {
	case *cli.StringFlag:
		return f.Usage
	case *cli.BoolFlag:
		return f.Usage
	case *cli.IntFlag:
		return f.Usage
	case *cli.DurationFlag:
		return f.Usage
	case *cli.StringSliceFlag:
		return f.Usage
	case *cli.Float64Flag:
		return f.Usage
	case *cli.PathFlag:
		return f.Usage
	default:
		return ""
	}
}

// showExamples displays common usage examples
func showExamples(c *cli.Context) error {
	isJSON := c.Bool("json")

	examples := []map[string]string{
		{
			"title":       "Basic local data onboarding",
			"description": "Onboard local data to local storage with automatic deal creation",
			"command":     "singularity onboard --name my-dataset --source /path/to/data --deal-template-id 1",
		},
		{
			"title":       "S3 to local storage",
			"description": "Onboard data from S3 bucket to local storage",
			"command":     "singularity onboard --name s3-data --source s3://bucket/data --source-type s3 --source-s3-region us-east-1 --deal-template-id template1",
		},
		{
			"title":       "Multiple sources with monitoring",
			"description": "Onboard from multiple local sources with progress monitoring",
			"command":     "singularity onboard --name multi-source --source /data1 --source /data2 --wait-for-completion --max-workers 5 --deal-template-id prod-template",
		},
		{
			"title":       "Google Cloud Storage to S3",
			"description": "Onboard from GCS bucket to S3 output storage",
			"command":     "singularity onboard --name gcs-to-s3 --source-type gcs --source gs://source-bucket/data --output-type s3 --output s3://dest-bucket/cars --deal-template-id gcs-template",
		},
		{
			"title":       "Custom configuration with JSON",
			"description": "Use JSON configuration for complex storage setup",
			"command":     "singularity onboard --name custom-config --source-type s3 --source s3://bucket/data --source-config '{\"access_key_id\":\"AKIAIOSFODNN7EXAMPLE\",\"secret_access_key\":\"secret\"}' --deal-template-id custom-template",
		},
	}

	if isJSON {
		data, err := json.MarshalIndent(examples, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(data))
		return nil
	}

	fmt.Println("Common Usage Examples:")
	fmt.Println("")
	for i, example := range examples {
		fmt.Printf("%d. %s\n", i+1, example["title"])
		fmt.Printf("   %s\n", example["description"])
		fmt.Printf("   %s\n", example["command"])
		fmt.Println("")
	}

	return nil
}

// showHelpJSON outputs help information in JSON format
func showHelpJSON(_ *cli.Context) error {
	type FlagInfo struct {
		Name         string   `json:"name"`
		Aliases      []string `json:"aliases,omitempty"`
		Usage        string   `json:"usage"`
		Category     string   `json:"category"`
		Required     bool     `json:"required"`
		DefaultValue string   `json:"default,omitempty"`
		Type         string   `json:"type"`
	}

	type HelpInfo struct {
		Name        string     `json:"name"`
		Usage       string     `json:"usage"`
		Description string     `json:"description"`
		Flags       []FlagInfo `json:"flags"`
		Backends    []string   `json:"backends"`
	}

	var flags []FlagInfo
	for _, flag := range OnboardCmd.Flags {
		var aliases []string
		var defaultValue string
		var usage string
		var category string
		flagType := "string"
		required := false

		switch f := flag.(type) {
		case *cli.StringFlag:
			aliases = f.Aliases
			defaultValue = f.Value
			usage = f.Usage
			category = f.Category
			required = f.Required
			flagType = "string"
		case *cli.BoolFlag:
			aliases = f.Aliases
			defaultValue = strconv.FormatBool(f.Value)
			usage = f.Usage
			category = f.Category
			flagType = "bool"
		case *cli.IntFlag:
			aliases = f.Aliases
			defaultValue = strconv.Itoa(f.Value)
			usage = f.Usage
			category = f.Category
			flagType = "int"
		case *cli.DurationFlag:
			aliases = f.Aliases
			defaultValue = f.Value.String()
			usage = f.Usage
			category = f.Category
			flagType = "duration"
		case *cli.StringSliceFlag:
			aliases = f.Aliases
			if f.Value != nil {
				defaultValue = strings.Join(f.Value.Value(), ",")
			}
			usage = f.Usage
			category = f.Category
			required = f.Required
			flagType = "string[]"
		case *cli.Float64Flag:
			aliases = f.Aliases
			defaultValue = fmt.Sprintf("%f", f.Value)
			usage = f.Usage
			category = f.Category
			flagType = "float64"
		case *cli.PathFlag:
			aliases = f.Aliases
			defaultValue = f.Value
			usage = f.Usage
			category = f.Category
			flagType = "path"
		}

		flags = append(flags, FlagInfo{
			Name:         flag.Names()[0],
			Aliases:      aliases,
			Usage:        usage,
			Category:     category,
			Required:     required,
			DefaultValue: defaultValue,
			Type:         flagType,
		})
	}

	var backends []string
	for _, backend := range storagesystem.Backends {
		backends = append(backends, backend.Prefix)
	}
	sort.Strings(backends)

	helpInfo := HelpInfo{
		Name:        OnboardCmd.Name,
		Usage:       OnboardCmd.Usage,
		Description: OnboardCmd.Description,
		Flags:       flags,
		Backends:    backends,
	}

	data, err := json.MarshalIndent(helpInfo, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

// getCustomHelpTemplate returns a custom help template that can conditionally show flags
func getCustomHelpTemplate() string {
	return `NAME:
   {{.HelpName}} - {{.Usage}}

USAGE:
   {{.HelpName}} [command options] [arguments...]

DESCRIPTION:
   {{.Description}}

OPTIONS:
{{range .VisibleFlags}}   {{.}}
{{end}}

SPECIALIZED HELP OPTIONS:
   --help-all             Show all available options including all backend-specific flags
   --help-backends        List all available storage backends (40+ supported)
   --help-backend=<type>  Show filtered options for specific backend (e.g., s3, gcs, local)
   --help-examples        Show common usage examples with backend configurations
   --help-json            Output help in JSON format for machine processing

BACKEND SUPPORT:
   This command supports all 40+ storage backends available in 'storage create'.
   Each backend has its own configuration options (e.g., --source-s3-region, --source-gcs-project-number).
   Use --help-backend=<type> to see only the flags relevant to your specific backend.

NOTE: By default, all backend flags are shown. Use --help-backend=<type> for filtered help.
      Use SINGULARITY_LIMIT_BACKENDS=true to show only common backends in help.
`
}
