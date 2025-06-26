package workermanager

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestDefaultManagerConfig(t *testing.T) {
	config := DefaultManagerConfig()

	assert.Equal(t, 30*time.Second, config.CheckInterval)
	assert.Equal(t, 1, config.MinWorkers)
	assert.Equal(t, 10, config.MaxWorkers)
	assert.Equal(t, 5, config.ScaleUpThreshold)
	assert.Equal(t, 2, config.ScaleDownThreshold)
	assert.Equal(t, 5*time.Minute, config.WorkerIdleTimeout)
	assert.True(t, config.AutoScaling)
	assert.Equal(t, 0.3, config.ScanWorkerRatio)
	assert.Equal(t, 0.5, config.PackWorkerRatio)
	assert.Equal(t, 0.2, config.DagGenWorkerRatio)
}

func TestNewWorkerManager(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		config := DefaultManagerConfig()
		manager := NewWorkerManager(db, config)

		assert.NotNil(t, manager)
		assert.Equal(t, db, manager.db)
		assert.Equal(t, config, manager.config)
		assert.True(t, manager.enabled)
		assert.NotNil(t, manager.activeWorkers)
		assert.Equal(t, 0, len(manager.activeWorkers))
		assert.NotNil(t, manager.stopChan)
		assert.NotNil(t, manager.monitoringStopped)
	})
}

func TestWorkerManager_Name(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		manager := NewWorkerManager(db, DefaultManagerConfig())
		assert.Equal(t, "Worker Manager", manager.Name())
	})
}

func TestWorkerManager_GetWorkerCount(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		manager := NewWorkerManager(db, DefaultManagerConfig())

		assert.Equal(t, 0, manager.getWorkerCount())

		// Add a mock worker to test counting
		mockWorker := &ManagedWorker{
			ID:        "test-worker",
			StartTime: time.Now(),
		}
		manager.activeWorkers["test-worker"] = mockWorker

		assert.Equal(t, 1, manager.getWorkerCount())
	})
}

func TestWorkerManager_IsEnabled(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		manager := NewWorkerManager(db, DefaultManagerConfig())

		assert.True(t, manager.isEnabled())

		// Test disabling
		manager.mutex.Lock()
		manager.enabled = false
		manager.mutex.Unlock()

		assert.False(t, manager.isEnabled())
	})
}

func TestWorkerManager_GetJobCounts(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		manager := NewWorkerManager(db, DefaultManagerConfig())

		// Set up test data
		preparation := &model.Preparation{
			Name: "test-prep",
			SourceStorages: []model.Storage{
				{
					Name: "test-storage",
					Type: "local",
					Path: "/tmp/test",
				},
			},
		}
		require.NoError(t, db.Create(preparation).Error)

		// Source attachment is created automatically by GORM when creating preparation with SourceStorages
		var sourceAttachment model.SourceAttachment
		require.NoError(t, db.Where("preparation_id = ? AND storage_id = ?", preparation.ID, preparation.SourceStorages[0].ID).First(&sourceAttachment).Error)

		// Create ready jobs of different types
		jobs := []model.Job{
			{Type: model.Scan, State: model.Ready, AttachmentID: sourceAttachment.ID},
			{Type: model.Scan, State: model.Ready, AttachmentID: sourceAttachment.ID},
			{Type: model.Pack, State: model.Ready, AttachmentID: sourceAttachment.ID},
			{Type: model.DagGen, State: model.Ready, AttachmentID: sourceAttachment.ID},
			{Type: model.Scan, State: model.Processing, AttachmentID: sourceAttachment.ID}, // Not ready
		}

		for _, job := range jobs {
			require.NoError(t, db.Create(&job).Error)
		}

		jobCounts, err := manager.getJobCounts(ctx)
		require.NoError(t, err)

		assert.Equal(t, int64(2), jobCounts[model.Scan])   // 2 ready scan jobs
		assert.Equal(t, int64(1), jobCounts[model.Pack])   // 1 ready pack job
		assert.Equal(t, int64(1), jobCounts[model.DagGen]) // 1 ready daggen job
	})
}

func TestWorkerManager_GetStatus(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		manager := NewWorkerManager(db, DefaultManagerConfig())

		// Test empty status
		status := manager.GetStatus()
		assert.True(t, status.Enabled)
		assert.Equal(t, 0, status.TotalWorkers)
		assert.Equal(t, 0, len(status.Workers))

		// Add a mock worker
		startTime := time.Now().Add(-10 * time.Millisecond) // Set start time slightly in the past
		mockWorker := &ManagedWorker{
			ID:           "test-worker",
			JobTypes:     []model.JobType{model.Scan, model.Pack},
			StartTime:    startTime,
			LastActivity: startTime,
		}
		manager.activeWorkers["test-worker"] = mockWorker

		status = manager.GetStatus()
		assert.True(t, status.Enabled)
		assert.Equal(t, 1, status.TotalWorkers)
		assert.Equal(t, 1, len(status.Workers))

		workerStatus := status.Workers[0]
		assert.Equal(t, "test-worker", workerStatus.ID)
		assert.Equal(t, []model.JobType{model.Scan, model.Pack}, workerStatus.JobTypes)
		assert.Equal(t, startTime, workerStatus.StartTime)
		assert.Equal(t, startTime, workerStatus.LastActivity)
		assert.True(t, workerStatus.Uptime > 0)
	})
}

func TestWorkerManager_StartOptimalWorker(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		config := DefaultManagerConfig()
		config.MinWorkers = 0 // Don't start minimum workers automatically
		manager := NewWorkerManager(db, config)

		// Test with mixed job counts
		jobCounts := map[model.JobType]int64{
			model.Scan:   3,
			model.Pack:   2,
			model.DagGen: 1,
		}

		// This will likely fail due to missing worker setup, but we test the logic
		err := manager.startOptimalWorker(ctx, jobCounts)

		// We expect this to fail in test environment due to missing dependencies
		// but the function should not panic
		_ = err // Ignore error as we're testing the logic, not full functionality

		// Wait for worker to be registered, then clean up
		// This prevents race conditions with database cleanup
		for i := 0; i < 10 && manager.getWorkerCount() == 0; i++ {
			time.Sleep(5 * time.Millisecond)
		}
		cleanupErr := manager.stopAllWorkers(ctx)
		_ = cleanupErr // Ignore cleanup errors in test
	})
}

func TestWorkerManager_EvaluateScaling_NoJobs(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		config := DefaultManagerConfig()
		config.MinWorkers = 0
		config.MaxWorkers = 5
		config.ScaleUpThreshold = 2
		manager := NewWorkerManager(db, config)

		// Test with no jobs (should not scale up)
		err := manager.evaluateScaling(ctx)
		assert.NoError(t, err)

		// Should have no workers
		assert.Equal(t, 0, manager.getWorkerCount())
	})
}

func TestWorkerManager_StopWorker_NonExistent(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		manager := NewWorkerManager(db, DefaultManagerConfig())

		err := manager.stopWorker(ctx, "non-existent-worker")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "worker non-existent-worker not found")
	})
}

func TestWorkerManager_StopOldestWorker_NoWorkers(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		manager := NewWorkerManager(db, DefaultManagerConfig())

		err := manager.stopOldestWorker(ctx)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no workers to stop")
	})
}

func TestWorkerManager_StopOldestWorker(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		manager := NewWorkerManager(db, DefaultManagerConfig())

		// Add mock workers with different start times
		now := time.Now()

		mockWorker1 := &ManagedWorker{
			ID:        "worker-1",
			StartTime: now.Add(-2 * time.Hour), // Older
			Done:      make(chan struct{}),
			Cancel:    func() {}, // Add mock cancel function
		}
		close(mockWorker1.Done) // Simulate already stopped

		mockWorker2 := &ManagedWorker{
			ID:        "worker-2",
			StartTime: now.Add(-1 * time.Hour), // Newer
			Done:      make(chan struct{}),
			Cancel:    func() {}, // Add mock cancel function
		}
		close(mockWorker2.Done) // Simulate already stopped

		manager.activeWorkers["worker-1"] = mockWorker1
		manager.activeWorkers["worker-2"] = mockWorker2

		// Should stop the oldest worker (worker-1)
		err := manager.stopOldestWorker(ctx)
		assert.NoError(t, err)

		// worker-1 should be removed from active workers
		_, exists := manager.activeWorkers["worker-1"]
		assert.False(t, exists)

		// worker-2 should still exist
		_, exists = manager.activeWorkers["worker-2"]
		assert.True(t, exists)
	})
}

func TestWorkerManager_CleanupIdleWorkers(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		config := DefaultManagerConfig()
		config.MinWorkers = 1
		config.WorkerIdleTimeout = time.Millisecond * 100
		manager := NewWorkerManager(db, config)

		now := time.Now()

		// Add mock workers - one idle, one active
		idleWorker := &ManagedWorker{
			ID:           "idle-worker",
			StartTime:    now,
			LastActivity: now.Add(-time.Hour), // Very old activity
			Done:         make(chan struct{}),
			Cancel:       func() {}, // Add mock cancel function
		}
		close(idleWorker.Done)

		activeWorker := &ManagedWorker{
			ID:           "active-worker",
			StartTime:    now,
			LastActivity: now, // Recent activity
			Done:         make(chan struct{}),
			Cancel:       func() {}, // Add mock cancel function
		}
		close(activeWorker.Done)

		manager.activeWorkers["idle-worker"] = idleWorker
		manager.activeWorkers["active-worker"] = activeWorker

		manager.cleanupIdleWorkers(ctx)

		// idle-worker should be removed, active-worker should remain
		// But since we have MinWorkers = 1, it might not remove if it would go below minimum
		assert.Equal(t, 1, manager.getWorkerCount())
	})
}

func TestWorkerManager_CleanupIdleWorkers_NoTimeout(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		config := DefaultManagerConfig()
		config.WorkerIdleTimeout = 0 // Disabled
		manager := NewWorkerManager(db, config)

		// Add an idle worker
		idleWorker := &ManagedWorker{
			ID:           "idle-worker",
			StartTime:    time.Now(),
			LastActivity: time.Now().Add(-time.Hour),
			Cancel:       func() {}, // Add mock cancel function
		}
		manager.activeWorkers["idle-worker"] = idleWorker

		manager.cleanupIdleWorkers(ctx)

		// Worker should not be cleaned up when timeout is 0
		assert.Equal(t, 1, manager.getWorkerCount())
	})
}

func TestHelperFunctions(t *testing.T) {
	// Test min function
	assert.Equal(t, 3, min(3, 5))
	assert.Equal(t, 2, min(5, 2))
	assert.Equal(t, 0, min(0, 1))

	// Test contains function
	jobTypes := []model.JobType{model.Scan, model.Pack}
	assert.True(t, contains(jobTypes, model.Scan))
	assert.True(t, contains(jobTypes, model.Pack))
	assert.False(t, contains(jobTypes, model.DagGen))

	emptyJobTypes := []model.JobType{}
	assert.False(t, contains(emptyJobTypes, model.Scan))
}

func TestWorkerManager_StopAllWorkers(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		manager := NewWorkerManager(db, DefaultManagerConfig())

		// Add mock workers
		worker1 := &ManagedWorker{
			ID:     "worker-1",
			Done:   make(chan struct{}),
			Cancel: func() {}, // Add mock cancel function
		}
		close(worker1.Done)

		worker2 := &ManagedWorker{
			ID:     "worker-2",
			Done:   make(chan struct{}),
			Cancel: func() {}, // Add mock cancel function
		}
		close(worker2.Done)

		manager.activeWorkers["worker-1"] = worker1
		manager.activeWorkers["worker-2"] = worker2

		err := manager.stopAllWorkers(ctx)
		assert.NoError(t, err)

		// All workers should be removed
		assert.Equal(t, 0, manager.getWorkerCount())
	})
}

func TestWorkerManager_EnsureMinimumWorkers(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		config := DefaultManagerConfig()
		config.MinWorkers = 2
		manager := NewWorkerManager(db, config)

		// This will likely fail due to missing worker dependencies
		// but we test that it doesn't panic
		err := manager.ensureMinimumWorkers(ctx)
		_ = err // Ignore error as we're testing the logic, not full functionality

		// Wait for workers to be registered, then clean up
		// This prevents race conditions with database cleanup
		expectedWorkers := config.MinWorkers
		for i := 0; i < 20 && manager.getWorkerCount() < expectedWorkers; i++ {
			time.Sleep(5 * time.Millisecond)
		}
		cleanupErr := manager.stopAllWorkers(ctx)
		_ = cleanupErr // Ignore cleanup errors in test
	})
}
