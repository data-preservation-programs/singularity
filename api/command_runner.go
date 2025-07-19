package api

import (
	"fmt"
)

// CommandRunner abstracts worker process creation
type CommandRunner interface {
	StartWorker(config WorkerConfig) (WorkerCmd, error)
}

// DefaultCommandRunner implements real process creation
type DefaultCommandRunner struct{}

func (r *DefaultCommandRunner) StartWorker(config WorkerConfig) (WorkerCmd, error) {
	args := []string{"run", "dataset-worker"}
	args = append(args, "--concurrency", fmt.Sprintf("%d", config.Concurrency))
	if !config.EnableScan {
		args = append(args, "--enable-scan=false")
	}
	if !config.EnablePack {
		args = append(args, "--enable-pack=false")
	}
	if !config.EnableDag {
		args = append(args, "--enable-dag=false")
	}
	if config.ExitOnComplete {
		args = append(args, "--exit-on-complete=true")
	}
	if config.ExitOnError {
		args = append(args, "--exit-on-error=true")
	}
	if config.MinInterval != "" {
		args = append(args, "--min-interval", config.MinInterval)
	}
	if config.MaxInterval != "" {
		args = append(args, "--max-interval", config.MaxInterval)
	}

	return NewRealWorkerCmd("singularity", args...), nil
}

// MockCommandRunner for testing
type MockCommandRunner struct{}

func (r *MockCommandRunner) StartWorker(config WorkerConfig) (WorkerCmd, error) {
	return NewMockWorkerCmd(), nil
}
