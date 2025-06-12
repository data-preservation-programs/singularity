package run

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/service/workermanager"
	"github.com/data-preservation-programs/singularity/service/workflow"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var logger = log.Logger("unified-service")

// UnifiedServiceCmd provides a single command to run both workflow orchestration and worker management
var UnifiedServiceCmd = &cli.Command{
	Name:    "unified",
	Aliases: []string{"auto"},
	Usage:   "Run unified auto-preparation service (workflow orchestration + worker management)",
	Description: `The unified service combines workflow orchestration and worker lifecycle management.
	
It automatically:
- Manages dataset worker lifecycle (start/stop workers based on job availability)
- Orchestrates job progression (scan → pack → daggen → deals)
- Scales workers up/down based on job queue
- Handles automatic deal creation when preparations complete

This is the recommended way to run fully automated data preparation.`,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "min-workers",
			Usage: "Minimum number of workers to keep running",
			Value: 1,
		},
		&cli.IntFlag{
			Name:  "max-workers",
			Usage: "Maximum number of workers to run",
			Value: 5,
		},
		&cli.IntFlag{
			Name:  "scale-up-threshold",
			Usage: "Number of ready jobs to trigger worker scale-up",
			Value: 5,
		},
		&cli.IntFlag{
			Name:  "scale-down-threshold",
			Usage: "Number of ready jobs below which to scale down workers",
			Value: 2,
		},
		&cli.DurationFlag{
			Name:  "check-interval",
			Usage: "How often to check for scaling and workflow progression",
			Value: 30 * time.Second,
		},
		&cli.DurationFlag{
			Name:  "worker-idle-timeout",
			Usage: "How long a worker can be idle before shutdown (0 = never)",
			Value: 5 * time.Minute,
		},
		&cli.BoolFlag{
			Name:  "disable-auto-scaling",
			Usage: "Disable automatic worker scaling",
		},
		&cli.BoolFlag{
			Name:  "disable-workflow-orchestration",
			Usage: "Disable automatic job progression",
		},
		&cli.BoolFlag{
			Name:  "disable-auto-deals",
			Usage: "Disable automatic deal creation",
		},
		&cli.BoolFlag{
			Name:  "disable-scan-to-pack",
			Usage: "Disable automatic scan → pack transitions",
		},
		&cli.BoolFlag{
			Name:  "disable-pack-to-daggen",
			Usage: "Disable automatic pack → daggen transitions",
		},
		&cli.BoolFlag{
			Name:  "disable-daggen-to-deals",
			Usage: "Disable automatic daggen → deals transitions",
		},
	},
	Action: func(c *cli.Context) error {
		// Initialize database
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		// Create worker manager
		workerConfig := workermanager.ManagerConfig{
			CheckInterval:      c.Duration("check-interval"),
			MinWorkers:         c.Int("min-workers"),
			MaxWorkers:         c.Int("max-workers"),
			ScaleUpThreshold:   c.Int("scale-up-threshold"),
			ScaleDownThreshold: c.Int("scale-down-threshold"),
			WorkerIdleTimeout:  c.Duration("worker-idle-timeout"),
			AutoScaling:        !c.Bool("disable-auto-scaling"),
			ScanWorkerRatio:    0.3,
			PackWorkerRatio:    0.5,
			DagGenWorkerRatio:  0.2,
		}

		workerManager := workermanager.NewWorkerManager(db, workerConfig)

		// Configure workflow orchestrator
		orchestratorConfig := workflow.OrchestratorConfig{
			EnableJobProgression: !c.Bool("disable-workflow-orchestration"),
			EnableAutoDeal:       !c.Bool("disable-auto-deals"),
			CheckInterval:        c.Duration("check-interval"),
			ScanToPack:           !c.Bool("disable-scan-to-pack"),
			PackToDagGen:         !c.Bool("disable-pack-to-daggen"),
			DagGenToDeals:        !c.Bool("disable-daggen-to-deals"),
		}

		orchestrator := workflow.NewWorkflowOrchestrator(orchestratorConfig)

		// Start unified service
		return runUnifiedService(c.Context, db, workerManager, orchestrator)
	},
}

// runUnifiedService runs the unified auto-preparation service
func runUnifiedService(ctx context.Context, db *gorm.DB, workerManager *workermanager.WorkerManager, orchestrator *workflow.WorkflowOrchestrator) error {
	logger.Info("Starting unified auto-preparation service")

	// Start worker manager
	err := workerManager.Start(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to start worker manager")
	}

	// Start workflow monitor (for batch processing of pending workflows)
	workflowDone := make(chan struct{})
	go func() {
		defer close(workflowDone)
		runWorkflowMonitor(ctx, db, orchestrator)
	}()

	// Print status periodically
	statusTicker := time.NewTicker(2 * time.Minute)
	defer statusTicker.Stop()

	statusDone := make(chan struct{})
	go func() {
		defer close(statusDone)
		for {
			select {
			case <-ctx.Done():
				return
			case <-statusTicker.C:
				printServiceStatus(db, workerManager, orchestrator)
			}
		}
	}()

	// Wait for context cancellation
	<-ctx.Done()
	logger.Info("Shutting down unified auto-preparation service")

	// Stop worker manager
	err = workerManager.Stop(ctx)
	if err != nil {
		logger.Errorf("Failed to stop worker manager: %v", err)
	}

	// Wait for background tasks to complete
	<-workflowDone
	<-statusDone

	logger.Info("Unified auto-preparation service stopped")
	return nil
}

// runWorkflowMonitor runs periodic workflow progression checks
func runWorkflowMonitor(ctx context.Context, db *gorm.DB, orchestrator *workflow.WorkflowOrchestrator) {
	logger.Info("Starting workflow monitor")

	// Create a lotus client for workflow operations
	lotusClient := util.NewLotusClient("", "")

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logger.Info("Workflow monitor stopped")
			return
		case <-ticker.C:
			err := orchestrator.ProcessPendingWorkflows(ctx, db, lotusClient)
			if err != nil {
				logger.Errorf("Failed to process pending workflows: %v", err)
			}
		}
	}
}

// printServiceStatus logs the current status of the unified service
func printServiceStatus(db *gorm.DB, workerManager *workermanager.WorkerManager, orchestrator *workflow.WorkflowOrchestrator) {
	// Get worker manager status
	workerStatus := workerManager.GetStatus()

	// Get job counts
	var jobCounts []struct {
		Type  string `json:"type"`
		State string `json:"state"`
		Count int64  `json:"count"`
	}

	db.Model(&struct {
		Type  string `gorm:"column:type"`
		State string `gorm:"column:state"`
		Count int64  `gorm:"column:count"`
	}{}).
		Table("jobs").
		Select("type, state, count(*) as count").
		Group("type, state").
		Find(&jobCounts)

	// Log comprehensive status
	logger.Infof("=== UNIFIED SERVICE STATUS ===")
	logger.Infof("Workers: %d active (enabled: %t)", workerStatus.TotalWorkers, workerStatus.Enabled)
	logger.Infof("Orchestrator enabled: %t", orchestrator.IsEnabled())

	// Log job counts
	readyJobs := map[string]int64{"scan": 0, "pack": 0, "daggen": 0}
	totalJobs := map[string]int64{"scan": 0, "pack": 0, "daggen": 0}

	for _, jc := range jobCounts {
		if _, exists := totalJobs[jc.Type]; exists {
			totalJobs[jc.Type] += jc.Count
			if jc.State == "ready" {
				readyJobs[jc.Type] = jc.Count
			}
		}
	}

	logger.Infof("Jobs - Scan: %d ready/%d total, Pack: %d ready/%d total, DagGen: %d ready/%d total",
		readyJobs["scan"], totalJobs["scan"],
		readyJobs["pack"], totalJobs["pack"],
		readyJobs["daggen"], totalJobs["daggen"])

	// Log worker details
	for _, worker := range workerStatus.Workers {
		logger.Infof("Worker %s: types=%v, uptime=%v",
			worker.ID[:8], worker.JobTypes, worker.Uptime.Truncate(time.Second))
	}
	logger.Infof("===============================")
}
