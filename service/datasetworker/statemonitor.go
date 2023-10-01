package datasetworker

import (
	"context"
	"sync"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service"
	"gorm.io/gorm"
)

func NewStateMonitor(db *gorm.DB) *StateMonitor {
	return &StateMonitor{
		db:   db,
		jobs: make(map[model.JobID]context.CancelFunc),
	}
}

type StateMonitor struct {
	db   *gorm.DB
	jobs map[model.JobID]context.CancelFunc
	mu   sync.RWMutex
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

func (s *StateMonitor) Start(ctx context.Context) service.Done {
	db := s.db.WithContext(ctx)
	monitorDone := make(chan struct{})
	go func() {
		defer close(monitorDone)
		for {
			var jobIDs []model.JobID
			s.mu.RLock()
			for jobID := range s.jobs {
				jobIDs = append(jobIDs, jobID)
			}
			s.mu.RUnlock()
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
			select {
			case <-ctx.Done():
				return
			case <-time.After(5 * time.Second):
			}
		}
	}()
	return monitorDone
}
