package datasetworker

import (
	"context"
	"sync"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

const jobCheckInterval = 5 * time.Second

func NewStateMonitor(db *gorm.DB) *StateMonitor {
	return &StateMonitor{
		db:   db,
		jobs: make(map[model.JobID]context.CancelFunc),
		done: make(chan struct{}),
	}
}

type StateMonitor struct {
	db   *gorm.DB
	jobs map[model.JobID]context.CancelFunc
	mu   sync.Mutex
	done chan struct{}
}

func (s *StateMonitor) AddJob(jobID model.JobID, cancel context.CancelFunc) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jobs[jobID] = cancel
}

func (s *StateMonitor) RemoveJob(jobID model.JobID) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.jobs, jobID)
}

func (s *StateMonitor) Start(ctx context.Context) {
	db := s.db.WithContext(ctx)
	go func() {
		defer close(s.done)
		var timer *time.Timer
		for {
			var i int
			s.mu.Lock()
			jobIDs := make([]model.JobID, len(s.jobs))
			for jobID := range s.jobs {
				jobIDs[i] = jobID
				i++
			}
			s.mu.Unlock()

			var jobs []model.Job
			if len(jobIDs) > 0 {
				err := db.Where("state = ?", model.Paused).Find(&jobs, jobIDs).Error
				if err != nil {
					logger.Errorf("failed to fetch paused jobs: %v", err)
				}
			}

			s.mu.Lock()
			for _, job := range jobs {
				jobID := job.ID
				cancel, ok := s.jobs[jobID]
				if ok {
					cancel()
					delete(s.jobs, jobID)
				}
			}
			s.mu.Unlock()

			if timer == nil {
				timer = time.NewTimer(jobCheckInterval)
				defer timer.Stop()
			} else {
				timer.Reset(jobCheckInterval)
			}
			select {
			case <-ctx.Done():
				return
			case <-timer.C:
			}
		}
	}()
}

func (s *StateMonitor) Done() <-chan struct{} {
	return s.done
}
