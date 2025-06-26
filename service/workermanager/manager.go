package workermanager

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/datasetworker"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

var logger = log.Logger("worker-manager")

// WorkerManager manages the lifecycle of dataset workers
type WorkerManager struct {
	db                *gorm.DB
	config            ManagerConfig
	activeWorkers     map[string]*ManagedWorker
	mutex             sync.RWMutex
	enabled           bool
	stopChan          chan struct{}
	monitoringStopped chan struct{}
}

// ManagerConfig configures the worker manager
type ManagerConfig struct {
	CheckInterval      time.Duration `json:"checkInterval"`      // How often to check for work availability
	MinWorkers         int           `json:"minWorkers"`         // Minimum number of workers to keep running
	MaxWorkers         int           `json:"maxWorkers"`         // Maximum number of workers to run
	ScaleUpThreshold   int           `json:"scaleUpThreshold"`   // Number of ready jobs to trigger scale-up
	ScaleDownThreshold int           `json:"scaleDownThreshold"` // Number of ready jobs below which to scale down
	WorkerIdleTimeout  time.Duration `json:"workerIdleTimeout"`  // How long a worker can be idle before shutdown
	AutoScaling        bool          `json:"autoScaling"`        // Enable automatic scaling
	ScanWorkerRatio    float64       `json:"scanWorkerRatio"`    // Proportion of workers for scan jobs
	PackWorkerRatio    float64       `json:"packWorkerRatio"`    // Proportion of workers for pack jobs
	DagGenWorkerRatio  float64       `json:"dagGenWorkerRatio"`  // Proportion of workers for daggen jobs
}

// DefaultManagerConfig returns sensible defaults
func DefaultManagerConfig() ManagerConfig {
	return ManagerConfig{
		CheckInterval:      30 * time.Second,
		MinWorkers:         1,
		MaxWorkers:         10,
		ScaleUpThreshold:   5,
		ScaleDownThreshold: 2,
		WorkerIdleTimeout:  5 * time.Minute,
		AutoScaling:        true,
		ScanWorkerRatio:    0.3, // 30% scan workers
		PackWorkerRatio:    0.5, // 50% pack workers
		DagGenWorkerRatio:  0.2, // 20% daggen workers
	}
}

// ManagedWorker represents a worker managed by the WorkerManager
type ManagedWorker struct {
	ID           string
	Worker       *datasetworker.Worker
	Config       datasetworker.Config
	StartTime    time.Time
	LastActivity time.Time
	Context      context.Context
	Cancel       context.CancelFunc
	ExitErr      chan error
	Done         chan struct{}
	JobTypes     []model.JobType
}

// NewWorkerManager creates a new worker manager
func NewWorkerManager(db *gorm.DB, config ManagerConfig) *WorkerManager {
	return &WorkerManager{
		db:                db,
		config:            config,
		activeWorkers:     make(map[string]*ManagedWorker),
		enabled:           true,
		stopChan:          make(chan struct{}),
		monitoringStopped: make(chan struct{}),
	}
}

// Start begins the worker management service
func (m *WorkerManager) Start(ctx context.Context) error {
	logger.Info("Starting worker manager")

	// Start minimum workers
	err := m.ensureMinimumWorkers(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	// Start monitoring goroutine
	go m.monitorLoop(ctx)

	return nil
}

// Stop shuts down the worker manager and all managed workers
func (m *WorkerManager) Stop(ctx context.Context) error {
	logger.Info("Stopping worker manager")

	m.mutex.Lock()
	m.enabled = false
	m.mutex.Unlock()

	// Signal monitoring to stop
	close(m.stopChan)

	// Wait for monitoring to stop
	select {
	case <-m.monitoringStopped:
	case <-ctx.Done():
		return ctx.Err()
	}

	// Stop all workers
	return m.stopAllWorkers(ctx)
}

// monitorLoop continuously monitors job availability and manages workers
func (m *WorkerManager) monitorLoop(ctx context.Context) {
	defer close(m.monitoringStopped)

	ticker := time.NewTicker(m.config.CheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-m.stopChan:
			return
		case <-ticker.C:
			if m.isEnabled() && m.config.AutoScaling {
				err := m.evaluateScaling(ctx)
				if err != nil {
					logger.Errorf("Failed to evaluate scaling: %v", err)
				}
			}

			// Clean up idle workers
			err := m.cleanupIdleWorkers(ctx)
			if err != nil {
				logger.Errorf("Failed to cleanup idle workers: %v", err)
			}
		}
	}
}

// evaluateScaling checks job availability and scales workers accordingly
func (m *WorkerManager) evaluateScaling(ctx context.Context) error {
	// Get job counts by type
	jobCounts, err := m.getJobCounts(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	totalReadyJobs := jobCounts[model.Scan] + jobCounts[model.Pack] + jobCounts[model.DagGen]
	currentWorkerCount := m.getWorkerCount()

	logger.Debugf("Job counts: scan=%d, pack=%d, daggen=%d, workers=%d",
		jobCounts[model.Scan], jobCounts[model.Pack], jobCounts[model.DagGen], currentWorkerCount)

	// Scale up if needed
	if totalReadyJobs >= int64(m.config.ScaleUpThreshold) && currentWorkerCount < m.config.MaxWorkers {
		workersToAdd := workerMin(m.config.MaxWorkers-currentWorkerCount, int(totalReadyJobs/int64(m.config.ScaleUpThreshold)))
		logger.Infof("Scaling up: adding %d workers (ready jobs: %d)", workersToAdd, totalReadyJobs)

		for i := 0; i < workersToAdd; i++ {
			err = m.startOptimalWorker(ctx, jobCounts)
			if err != nil {
				logger.Errorf("Failed to start worker: %v", err)
				break
			}
		}
	}

	// Scale down if needed (but keep minimum)
	if totalReadyJobs <= int64(m.config.ScaleDownThreshold) && currentWorkerCount > m.config.MinWorkers {
		workersToRemove := workerMin(currentWorkerCount-m.config.MinWorkers, 1) // Remove one at a time
		logger.Infof("Scaling down: removing %d workers (ready jobs: %d)", workersToRemove, totalReadyJobs)

		for i := 0; i < workersToRemove; i++ {
			err = m.stopOldestWorker(ctx)
			if err != nil {
				logger.Errorf("Failed to stop worker: %v", err)
				break
			}
		}
	}

	return nil
}

// startOptimalWorker starts a worker optimized for current job distribution
func (m *WorkerManager) startOptimalWorker(ctx context.Context, jobCounts map[model.JobType]int64) error {
	// Determine optimal job types for this worker based on current distribution
	var jobTypes []model.JobType
	if jobCounts[model.DagGen] > 0 {
		jobTypes = append(jobTypes, model.DagGen) // Prioritize DagGen (final stage)
	}
	if jobCounts[model.Scan] > 0 {
		jobTypes = append(jobTypes, model.Scan)
	}
	if jobCounts[model.Pack] > 0 {
		jobTypes = append(jobTypes, model.Pack)
	}

	// If no specific jobs, create a general-purpose worker
	if len(jobTypes) == 0 {
		jobTypes = []model.JobType{model.Scan, model.Pack, model.DagGen}
	}

	return m.startWorker(ctx, jobTypes, 1)
}

// startWorker starts a new worker with specified configuration
func (m *WorkerManager) startWorker(ctx context.Context, jobTypes []model.JobType, concurrency int) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	workerID := fmt.Sprintf("managed-worker-%d", time.Now().UnixNano())

	config := datasetworker.Config{
		Concurrency:    concurrency,
		ExitOnComplete: false, // Managed workers should not exit automatically
		EnableScan:     contains(jobTypes, model.Scan),
		EnablePack:     contains(jobTypes, model.Pack),
		EnableDag:      contains(jobTypes, model.DagGen),
		ExitOnError:    false, // Managed workers should be resilient
		MinInterval:    5 * time.Second,
		MaxInterval:    30 * time.Second,
	}

	worker := datasetworker.NewWorker(m.db, config)
	workerCtx, cancel := context.WithCancel(ctx)
	exitErr := make(chan error, 1)
	done := make(chan struct{})

	managedWorker := &ManagedWorker{
		ID:           workerID,
		Worker:       worker,
		Config:       config,
		StartTime:    time.Now(),
		LastActivity: time.Now(),
		Context:      workerCtx,
		Cancel:       cancel,
		ExitErr:      exitErr,
		Done:         done,
		JobTypes:     jobTypes,
	}

	// Start worker in goroutine
	go func() {
		defer close(done)
		defer cancel()

		logger.Infof("Starting managed worker %s with job types: %v", workerID, jobTypes)
		err := worker.Run(workerCtx)
		if err != nil && !errors.Is(err, context.Canceled) {
			logger.Errorf("Managed worker %s exited with error: %v", workerID, err)
			select {
			case exitErr <- err:
			default:
			}
		} else {
			logger.Infof("Managed worker %s exited normally", workerID)
		}

		// Remove from active workers
		m.mutex.Lock()
		delete(m.activeWorkers, workerID)
		m.mutex.Unlock()
	}()

	m.activeWorkers[workerID] = managedWorker
	logger.Infof("Started managed worker %s (total workers: %d)", workerID, len(m.activeWorkers))

	return nil
}

// stopWorker stops a specific worker
func (m *WorkerManager) stopWorker(ctx context.Context, workerID string) error {
	m.mutex.Lock()
	worker, exists := m.activeWorkers[workerID]
	if !exists || worker == nil {
		m.mutex.Unlock()
		return errors.Errorf("worker %s not found or is nil", workerID)
	}
	delete(m.activeWorkers, workerID)
	m.mutex.Unlock()

	logger.Infof("Stopping managed worker %s", workerID)
	if worker.Cancel == nil {
		return errors.Errorf("worker %s has nil Cancel function", workerID)
	}
	worker.Cancel()

	// Wait for worker to stop with timeout
	stopCtx, stopCancel := context.WithTimeout(ctx, 30*time.Second)
	defer stopCancel()

	if worker.Done != nil {
		select {
		case <-worker.Done:
			logger.Infof("Managed worker %s stopped successfully", workerID)
		case <-stopCtx.Done():
			logger.Warnf("Timeout waiting for worker %s to stop", workerID)
		}
	} else {
		logger.Warnf("Worker %s has nil Done channel, cannot wait for stop confirmation", workerID)
	}

	return nil
}

// stopOldestWorker stops the worker that has been running the longest
func (m *WorkerManager) stopOldestWorker(ctx context.Context) error {
	m.mutex.RLock()
	var oldestWorkerID string
	var oldestTime time.Time

	for id, worker := range m.activeWorkers {
		if oldestWorkerID == "" || worker.StartTime.Before(oldestTime) {
			oldestWorkerID = id
			oldestTime = worker.StartTime
		}
	}
	m.mutex.RUnlock()

	if oldestWorkerID == "" {
		return errors.New("no workers to stop")
	}

	return m.stopWorker(ctx, oldestWorkerID)
}

// stopAllWorkers stops all managed workers
func (m *WorkerManager) stopAllWorkers(ctx context.Context) error {
	m.mutex.RLock()
	var workerIDs []string
	for id := range m.activeWorkers {
		workerIDs = append(workerIDs, id)
	}
	m.mutex.RUnlock()

	for _, id := range workerIDs {
		err := m.stopWorker(ctx, id)
		if err != nil {
			logger.Errorf("Failed to stop worker %s: %v", id, err)
		}
	}

	return nil
}

// ensureMinimumWorkers ensures minimum number of workers are running
func (m *WorkerManager) ensureMinimumWorkers(ctx context.Context) error {
	currentCount := m.getWorkerCount()
	needed := m.config.MinWorkers - currentCount

	for i := 0; i < needed; i++ {
		// Start general-purpose workers for minimum baseline
		err := m.startWorker(ctx, []model.JobType{model.Scan, model.Pack, model.DagGen}, 1)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// cleanupIdleWorkers removes workers that have been idle too long
// Currently always returns nil, but error return is kept for future extensibility
func (m *WorkerManager) cleanupIdleWorkers(ctx context.Context) error {
	if m.config.WorkerIdleTimeout == 0 {
		return nil // No cleanup if timeout is 0
	}

	m.mutex.RLock()
	var idleWorkers []string
	now := time.Now()

	for id, worker := range m.activeWorkers {
		if now.Sub(worker.LastActivity) > m.config.WorkerIdleTimeout {
			idleWorkers = append(idleWorkers, id)
		}
	}
	m.mutex.RUnlock()

	// Don't cleanup if it would go below minimum
	if len(idleWorkers) > 0 && m.getWorkerCount()-len(idleWorkers) >= m.config.MinWorkers {
		for _, id := range idleWorkers {
			logger.Infof("Cleaning up idle worker %s", id)
			err := m.stopWorker(ctx, id)
			if err != nil {
				logger.Errorf("Failed to cleanup idle worker %s: %v", id, err)
			}
		}
	}

	return nil
}

// getJobCounts returns count of ready jobs by type
func (m *WorkerManager) getJobCounts(ctx context.Context) (map[model.JobType]int64, error) {
	type JobCount struct {
		Type  model.JobType `json:"type"`
		Count int64         `json:"count"`
	}

	var jobCounts []JobCount
	err := m.db.WithContext(ctx).Model(&model.Job{}).
		Select("type, count(*) as count").
		Where("state = ?", model.Ready).
		Group("type").
		Find(&jobCounts).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := map[model.JobType]int64{
		model.Scan:   0,
		model.Pack:   0,
		model.DagGen: 0,
	}

	for _, jc := range jobCounts {
		result[jc.Type] = jc.Count
	}

	return result, nil
}

// getWorkerCount returns the current number of active workers
func (m *WorkerManager) getWorkerCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.activeWorkers)
}

// isEnabled returns whether the manager is enabled
func (m *WorkerManager) isEnabled() bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.enabled
}

// GetStatus returns the current status of the worker manager
func (m *WorkerManager) GetStatus() ManagerStatus {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	status := ManagerStatus{
		Enabled:      m.enabled,
		TotalWorkers: len(m.activeWorkers),
		Workers:      make([]WorkerStatus, 0, len(m.activeWorkers)),
	}

	for _, worker := range m.activeWorkers {
		status.Workers = append(status.Workers, WorkerStatus{
			ID:           worker.ID,
			JobTypes:     worker.JobTypes,
			StartTime:    worker.StartTime,
			LastActivity: worker.LastActivity,
			Uptime:       time.Since(worker.StartTime),
		})
	}

	return status
}

// ManagerStatus represents the current status of the worker manager
type ManagerStatus struct {
	Enabled      bool           `json:"enabled"`
	TotalWorkers int            `json:"totalWorkers"`
	Workers      []WorkerStatus `json:"workers"`
}

// WorkerStatus represents the status of a single managed worker
type WorkerStatus struct {
	ID           string          `json:"id"`
	JobTypes     []model.JobType `json:"jobTypes"`
	StartTime    time.Time       `json:"startTime"`
	LastActivity time.Time       `json:"lastActivity"`
	Uptime       time.Duration   `json:"uptime"`
}

// Name returns the service name
func (m *WorkerManager) Name() string {
	return "Worker Manager"
}

// Helper functions
func workerMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func contains(slice []model.JobType, item model.JobType) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
