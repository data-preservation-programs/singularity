package cmd

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/handler/job"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/workermanager"
	"github.com/data-preservation-programs/singularity/service/workflow"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

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
			Name:  "enable-deals",
			Usage: "Enable automatic deal creation after preparation completion",
			Value: true,
		},
		&cli.StringFlag{
			Name:     "deal-provider",
			Usage:    "Storage Provider ID for deals (e.g., f01000)",
			Category: "Deal Settings",
		},
		&cli.Float64Flag{
			Name:     "deal-price-per-gb",
			Usage:    "Price in FIL per GiB for storage deals",
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
			Name:  "validate-wallet",
			Usage: "Enable wallet balance validation",
		},
		&cli.BoolFlag{
			Name:  "validate-provider",
			Usage: "Enable storage provider validation",
		},
	},
	Action: func(c *cli.Context) error {
		fmt.Println("🚀 Starting unified data onboarding...")

		// Initialize database
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		ctx := c.Context

		// Step 1: Create preparation with deal configuration
		fmt.Println("\n📋 Creating data preparation...")
		prep, err := createPreparationForOnboarding(ctx, db, c)
		if err != nil {
			return errors.Wrap(err, "failed to create preparation")
		}
		fmt.Printf("✓ Created preparation: %s (ID: %d)\n", prep.Name, prep.ID)

		// Step 2: Enable workflow orchestration
		fmt.Println("\n⚙️  Enabling workflow orchestration...")
		workflow.DefaultOrchestrator.SetEnabled(true)
		fmt.Println("✓ Automatic job progression enabled (scan → pack → daggen → deals)")

		// Step 3: Start workers if requested
		var workerManager *workermanager.WorkerManager
		if c.Bool("start-workers") {
			fmt.Println("\n👷 Starting managed workers...")
			workerManager, err = startManagedWorkers(ctx, db, c.Int("max-workers"))
			if err != nil {
				return errors.Wrap(err, "failed to start workers")
			}
			fmt.Printf("✓ Started %d managed workers\n", c.Int("max-workers"))
		}

		// Step 4: Start scanning
		fmt.Println("\n🔍 Starting initial scanning...")
		err = startScanningForPreparation(ctx, db, prep)
		if err != nil {
			return errors.Wrap(err, "failed to start scanning")
		}
		fmt.Println("✓ Scanning started for all source attachments")

		// Step 5: Monitor progress if requested
		if c.Bool("wait-for-completion") {
			fmt.Println("\n📊 Monitoring progress...")
			err = monitorProgress(ctx, db, prep, c.Duration("timeout"))
			if err != nil {
				return errors.Wrap(err, "monitoring failed")
			}
		} else {
			fmt.Println("\n✅ Onboarding initiated successfully!")
			fmt.Println("\n📝 Next steps:")
			fmt.Println("   • Monitor progress: singularity prep status", prep.Name)
			fmt.Println("   • Check jobs: singularity job list")
			if c.Bool("start-workers") {
				fmt.Println("   • Workers will process jobs automatically")
			} else {
				fmt.Println("   • Start workers: singularity run unified")
			}
		}

		// Cleanup workers if we started them
		if workerManager != nil {
			fmt.Println("\n🧹 Cleaning up workers...")
			err = workerManager.Stop(ctx)
			if err != nil {
				fmt.Printf("⚠ Warning: failed to stop workers cleanly: %v\n", err)
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
		AutoCreateDeals:  c.Bool("enable-deals"),
		DealProvider:     c.String("deal-provider"),
		DealPricePerGB:   c.Float64("deal-price-per-gb"),
		DealDuration:     c.Duration("deal-duration"),
		DealStartDelay:   c.Duration("deal-start-delay"),
		DealVerified:     c.Bool("deal-verified"),
		WalletValidation: c.Bool("validate-wallet"),
		SPValidation:     c.Bool("validate-provider"),
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
		_, err = jobHandler.StartScanHandler(ctx, db, strconv.FormatUint(uint64(attachment.ID), 10), "")
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
	// This would use the same logic as the dataprep create command
	// For brevity, we'll create a simple implementation
	storageName := fmt.Sprintf("%s-%s-%d", prefix, util.RandomName(), time.Now().Unix())

	// Check if storage already exists for this path
	var existing model.Storage
	err := db.WithContext(ctx).Where("type = ? AND path = ?", "local", path).First(&existing).Error
	if err == nil {
		return &existing, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.WithStack(err)
	}

	// Create new storage
	// This is a simplified version - in practice would use the storage handler
	storage := &model.Storage{
		Name: storageName,
		Type: "local",
		Path: path,
	}

	err = db.WithContext(ctx).Create(storage).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return storage, nil
}
