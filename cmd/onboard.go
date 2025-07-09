package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/handler/job"
	storageHandlers "github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/workermanager"
	"github.com/data-preservation-programs/singularity/service/workflow"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/dustin/go-humanize"
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
2. Creates data preparation with deal parameters
3. Starts scanning immediately
4. Enables automatic job progression (scan → pack → daggen → deals)
5. Optionally starts managed workers to process jobs

This is the simplest way to onboard data from source to storage deals.`,
	Flags: []cli.Flag{
		// Data source flags
		&cli.StringFlag{
			Name:     "name",
			Usage:    "Name for the preparation",
			Required: true,
		},
		&cli.StringSliceFlag{
			Name:     "source",
			Usage:    "Local source path(s) to onboard",
			Required: true,
		},
		&cli.StringSliceFlag{
			Name:  "output",
			Usage: "Local output path(s) for CAR files (optional)",
		},

		// Preparation settings
		&cli.StringFlag{
			Name:  "max-size",
			Usage: "Maximum size of a single CAR file",
			Value: "31.5GiB",
		},
		&cli.BoolFlag{
			Name:  "no-dag",
			Usage: "Disable maintaining folder DAG structure",
		},

		// Deal configuration
		&cli.BoolFlag{
			Name:  "auto-create-deals",
			Usage: "Enable automatic deal creation after preparation completion",
			Value: true,
		},
		&cli.StringFlag{
			Name:     "deal-template-id",
			Usage:    "Deal template ID to use for deal configuration (required when auto-create-deals is enabled)",
			Category: "Deal Settings",
		},

		// Worker management
		&cli.BoolFlag{
			Name:  "start-workers",
			Usage: "Start managed workers to process jobs automatically",
			Value: true,
		},
		&cli.IntFlag{
			Name:  "max-workers",
			Usage: "Maximum number of workers to run",
			Value: 3,
		},

		// Progress monitoring
		&cli.BoolFlag{
			Name:  "wait-for-completion",
			Usage: "Wait and monitor until all jobs complete",
		},
		&cli.DurationFlag{
			Name:  "timeout",
			Usage: "Timeout for waiting for completion (0 = no timeout)",
			Value: 0,
		},

		// Validation
		&cli.BoolFlag{
			Name:  "wallet-validation",
			Usage: "Enable wallet balance validation",
			Value: true,
		},
		&cli.BoolFlag{
			Name:  "sp-validation",
			Usage: "Enable storage provider validation",
			Value: true,
		},

		// Output format
		&cli.BoolFlag{
			Name:  "json",
			Usage: "Output result in JSON format for automation",
		},
	},
	Action: func(c *cli.Context) error {
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

		if !isJSON {
			fmt.Println("🚀 Starting unified data onboarding...")
		}

		// Initialize database
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return outputJSONError("failed to initialize database", err)
		}
		defer func() { _ = closer.Close() }()

		ctx := c.Context

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
		} else {
			if !c.Bool("wait-for-completion") {
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
		}

		return nil
	},
}

// createPreparationForOnboarding creates a preparation with all onboarding settings
func createPreparationForOnboarding(ctx context.Context, db *gorm.DB, c *cli.Context) (*model.Preparation, error) {
	// Convert source paths to storage names (create if needed)
	var sourceStorages []string
	for _, sourcePath := range c.StringSlice("source") {
		storage, err := createLocalStorageIfNotExist(ctx, db, sourcePath, "source")
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create source storage for %s", sourcePath)
		}
		sourceStorages = append(sourceStorages, storage.Name)
	}

	// Convert output paths to storage names (create if needed)
	var outputStorages []string
	for _, outputPath := range c.StringSlice("output") {
		storage, err := createLocalStorageIfNotExist(ctx, db, outputPath, "output")
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

// Helper function to create local storage if it doesn't exist
func createLocalStorageIfNotExist(ctx context.Context, db *gorm.DB, path, prefix string) (*model.Storage, error) {
	// Check if storage already exists for this path
	var existing model.Storage
	err := db.WithContext(ctx).Where("type = ? AND path = ?", "local", path).First(&existing).Error
	if err == nil {
		return &existing, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.WithStack(err)
	}

	// Generate a unique storage name
	storageName := fmt.Sprintf("%s-%s-%d", prefix, util.RandomName(), time.Now().Unix())

	// Use the storage handler to create new storage with proper validation
	storageHandler := storageHandlers.Default
	request := storageHandlers.CreateRequest{
		Name:         storageName,
		Path:         path,
		Provider:     "local",
		Config:       make(map[string]string),
		ClientConfig: model.ClientConfig{},
	}

	storage, err := storageHandler.CreateStorageHandler(ctx, db, "local", request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return storage, nil
}

// validateOnboardInputs validates CLI inputs for onboard command
func validateOnboardInputs(c *cli.Context) error {
	// Required fields validation
	if c.String("name") == "" {
		return errors.New("preparation name is required (--name)")
	}

	// Source validation
	sourcePaths := c.StringSlice("source")

	if len(sourcePaths) == 0 {
		return errors.New("at least one source path is required (--source)")
	}

	// Output paths are optional - no validation needed

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
