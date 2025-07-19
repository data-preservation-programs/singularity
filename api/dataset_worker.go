package api

import (
	"bytes"
	"net/http"
	"sync"
	"time"
	"fmt"
	"github.com/labstack/echo/v4"
)

// WorkerConfig defines dataset worker configuration
type WorkerConfig struct {
	Concurrency    int    `json:"concurrency"`
	EnableScan     bool   `json:"enableScan"`
	EnablePack     bool   `json:"enablePack"`
	EnableDag      bool   `json:"enableDag"`
	ExitOnComplete bool   `json:"exitOnComplete"`
	ExitOnError    bool   `json:"exitOnError"`
	MinInterval    string `json:"minInterval"`
	MaxInterval    string `json:"maxInterval"`
}

// WorkerProcess holds the dataset-worker process and its config
type WorkerProcess struct {
	Cmd       WorkerCmd
	Config    WorkerConfig
	StartTime time.Time
	StdoutBuf *bytes.Buffer
	StderrBuf *bytes.Buffer
}

var (
	workerMutex   sync.Mutex
	workerProcess *WorkerProcess
	commandRunner CommandRunner = &DefaultCommandRunner{}
)
// DatasetWorkerAPI provides endpoints to manage the dataset-worker process
func RegisterDatasetWorkerAPI(e *echo.Echo) {
	e.POST("/api/dataset-worker/start", StartDatasetWorkerHandler)
	e.POST("/api/dataset-worker/stop", StopDatasetWorkerHandler)
	e.GET("/api/dataset-worker/status", StatusDatasetWorkerHandler)
	e.GET("/api/dataset-worker/logs", LogsDatasetWorkerHandler)
	e.GET("/api/dataset-worker/jobs", JobsDatasetWorkerHandler)
}

func StartDatasetWorkerHandler(c echo.Context) error {
	workerMutex.Lock()
	defer workerMutex.Unlock()

	if workerProcess != nil && workerProcess.Cmd != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Worker already running"})
	}

	var config WorkerConfig
	if err := c.Bind(&config); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid config"})
	}

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

	workerCmd, err := commandRunner.StartWorker(config)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create worker: " + err.Error(),
		})
	}

	stdoutBuf := &bytes.Buffer{}
	stderrBuf := &bytes.Buffer{}
	workerCmd.SetOutput(stdoutBuf, stderrBuf)
	
	workerProcess = &WorkerProcess{
		Cmd:       workerCmd,
		Config:    config,
		StartTime: time.Now(),
		StdoutBuf: stdoutBuf,
		StderrBuf: stderrBuf,
	}
	
	if err := workerCmd.Start(); err != nil {
		workerProcess = nil
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to start worker process: " + err.Error()})
	}

	pid := workerCmd.Pid()
	if pid == 0 {
		workerProcess = nil
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Worker process started but has no PID"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "started",
		"pid": pid,
	})
}

func StopDatasetWorkerHandler(c echo.Context) error {
	workerMutex.Lock()
	defer workerMutex.Unlock()

	// Check if worker exists and is running
	if workerProcess == nil || workerProcess.Cmd == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "No worker process is currently running",
		})
	}

	// Stop the worker process 
	err := workerProcess.Cmd.Kill()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to stop worker: " + err.Error(), 
		})
	}

	// Clean up the worker state
	workerProcess = nil
	return c.JSON(http.StatusOK, map[string]string{
		"status": "Worker process stopped successfully",
	})
}

func StatusDatasetWorkerHandler(c echo.Context) error {
	workerMutex.Lock()
	defer workerMutex.Unlock()

	status := map[string]interface{}{
		"running": false,
	}
	if workerProcess != nil && workerProcess.Cmd != nil {
		status["running"] = true
		status["pid"] = workerProcess.Cmd.Pid()
		status["config"] = workerProcess.Config
		status["startTime"] = workerProcess.StartTime
		if !workerProcess.StartTime.IsZero() {
			status["uptime"] = time.Since(workerProcess.StartTime).String()
		}
	}
	return c.JSON(http.StatusOK, status)
}

func LogsDatasetWorkerHandler(c echo.Context) error {
	workerMutex.Lock()
	defer workerMutex.Unlock()

	if workerProcess == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Worker not running"})
	}
	logs := map[string]string{
		"stdout": "",
		"stderr": "",
	}
	if workerProcess.StdoutBuf != nil {
		logs["stdout"] = workerProcess.StdoutBuf.String()
	}
	if workerProcess.StderrBuf != nil {
		logs["stderr"] = workerProcess.StderrBuf.String()
	}
	return c.JSON(http.StatusOK, logs)
}

func JobsDatasetWorkerHandler(c echo.Context) error {
	// Return empty jobs list - this is a valid response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"jobs": []interface{}{}, // Empty array indicates no running jobs
	})
}
