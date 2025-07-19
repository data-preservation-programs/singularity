package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Integration test for starting and stopping the worker
func TestStartStopWorkerIntegration(t *testing.T) {
	// Protect global state
	workerMutex.Lock()
	// Reset any existing worker process
	workerProcess = nil
	workerMutex.Unlock()
	
	// Inject mock command runner before creating any processes
	oldRunner := commandRunner
	commandRunner = &MockCommandRunner{} 
	defer func() { 
		commandRunner = oldRunner
		// Clean up after test
		workerMutex.Lock()
		workerProcess = nil 
		workerMutex.Unlock()
	}()

	e := echo.New()
	jsonPayload := `{"concurrency":1,"enableScan":true,"enablePack":true,"enableDag":true,"exitOnComplete":false,"exitOnError":false,"minInterval":"5s","maxInterval":"160s"}`
	req := httptest.NewRequest(http.MethodPost, "/api/dataset-worker/start", bytes.NewBufferString(jsonPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	c := e.NewContext(req, w)

	// Start worker
	err := StartDatasetWorkerHandler(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code, "Expected OK status code when starting worker")

	// Verify worker process exists and has expected PID
	if assert.NotNil(t, workerProcess, "Worker process should not be nil") {
		if assert.NotNil(t, workerProcess.Cmd, "Worker command should not be nil") {
			pid := workerProcess.Cmd.Pid()
			assert.Equal(t, 12345, pid, "Worker should have the mock PID")
		}
	}

	// Stop worker
	reqStop := httptest.NewRequest(http.MethodPost, "/api/dataset-worker/stop", nil)
	wStop := httptest.NewRecorder()
	cStop := e.NewContext(reqStop, wStop)
	err = StopDatasetWorkerHandler(cStop)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, wStop.Code)

	// Verify the worker is stopped
	assert.Nil(t, workerProcess)
}

// Integration test for status endpoint
func TestStatusWorkerIntegration(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/dataset-worker/status", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(req, w)
	_ = StatusDatasetWorkerHandler(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

// Integration test for logs endpoint
func TestLogsWorkerIntegration(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/dataset-worker/logs", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(req, w)
	_ = LogsDatasetWorkerHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

// Integration test for jobs endpoint with mock DB
func TestJobsWorkerIntegration(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/dataset-worker/jobs", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(req, w)
	
	// Return empty response with OK status since DB initialization isn't critical
	err := JobsDatasetWorkerHandler(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}
