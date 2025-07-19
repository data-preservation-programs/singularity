package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestStartDatasetWorkerHandler_Success(t *testing.T) {
	// Setup
	e := echo.New()
	jsonPayload := `{
		"concurrency": 1,
		"enableScan": true,
		"enablePack": true,
		"enableDag": true,
		"exitOnComplete": false,
		"exitOnError": false,
		"minInterval": "5s",
		"maxInterval": "160s"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/dataset-worker/start", strings.NewReader(jsonPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Inject mock command runner
	oldRunner := commandRunner
	commandRunner = &MockCommandRunner{}
	defer func() {
		commandRunner = oldRunner
		workerProcess = nil // Cleanup
	}()

	// Test
	err := StartDatasetWorkerHandler(c)
	
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotNil(t, workerProcess)
	assert.NotNil(t, workerProcess.Cmd)
	assert.Equal(t, 12345, workerProcess.Cmd.Pid())
}

func TestStartDatasetWorkerHandler_InvalidConfig(t *testing.T) {
	// Setup
	e := echo.New()
	invalidJson := `{"concurrency": "invalid"}`
	
	req := httptest.NewRequest(http.MethodPost, "/api/dataset-worker/start", strings.NewReader(invalidJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	err := StartDatasetWorkerHandler(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestStartDatasetWorkerHandler_AlreadyRunning(t *testing.T) {
	// Setup
	e := echo.New()
	jsonPayload := `{"concurrency": 1}`

	// Inject mock command runner and start a worker
	oldRunner := commandRunner
	commandRunner = &MockCommandRunner{}
	defer func() {
		commandRunner = oldRunner
		workerProcess = nil
	}()

	// Start first worker
	req1 := httptest.NewRequest(http.MethodPost, "/api/dataset-worker/start", strings.NewReader(jsonPayload))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	_ = StartDatasetWorkerHandler(c1)

	// Try to start second worker
	req2 := httptest.NewRequest(http.MethodPost, "/api/dataset-worker/start", strings.NewReader(jsonPayload))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	
	// Test
	err := StartDatasetWorkerHandler(c2)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusConflict, rec2.Code)
}

func TestStopDatasetWorkerHandler_Success(t *testing.T) {
	// Setup
	e := echo.New()
	
	// Start a worker first
	oldRunner := commandRunner
	commandRunner = &MockCommandRunner{}
	defer func() {
		commandRunner = oldRunner
		workerProcess = nil
	}()

	jsonPayload := `{"concurrency": 1}`
	reqStart := httptest.NewRequest(http.MethodPost, "/api/dataset-worker/start", strings.NewReader(jsonPayload))
	reqStart.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recStart := httptest.NewRecorder()
	cStart := e.NewContext(reqStart, recStart)
	_ = StartDatasetWorkerHandler(cStart)

	// Now try to stop it
	req := httptest.NewRequest(http.MethodPost, "/api/dataset-worker/stop", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(req, w)
	
	err := StopDatasetWorkerHandler(c)
	
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, workerProcess)
}

func TestStopDatasetWorkerHandler_NotRunning(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/dataset-worker/stop", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(req, w)

	// Ensure worker process is nil
	workerProcess = nil
	
	err := StopDatasetWorkerHandler(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestStatusDatasetWorkerHandler_Running(t *testing.T) {
	// Setup
	e := echo.New()
	
	// Start a worker first
	oldRunner := commandRunner
	commandRunner = &MockCommandRunner{}
	defer func() {
		commandRunner = oldRunner
		workerProcess = nil
	}()

	jsonPayload := `{"concurrency": 1}`
	reqStart := httptest.NewRequest(http.MethodPost, "/api/dataset-worker/start", strings.NewReader(jsonPayload))
	reqStart.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recStart := httptest.NewRecorder()
	cStart := e.NewContext(reqStart, recStart)
	_ = StartDatasetWorkerHandler(cStart)

	// Now check status
	req := httptest.NewRequest(http.MethodGet, "/api/dataset-worker/status", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(req, w)

	err := StatusDatasetWorkerHandler(c)
	
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"running":true`)
	assert.Contains(t, w.Body.String(), `"pid":12345`)
}

func TestStatusDatasetWorkerHandler_NotRunning(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/dataset-worker/status", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(req, w)

	workerProcess = nil
	err := StatusDatasetWorkerHandler(c)
	
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"running":false`)
}

func TestLogsDatasetWorkerHandler_Success(t *testing.T) {
	// Setup
	e := echo.New()
	
	// Start a worker first
	oldRunner := commandRunner
	commandRunner = &MockCommandRunner{}
	defer func() {
		commandRunner = oldRunner
		workerProcess = nil
	}()

	jsonPayload := `{"concurrency": 1}`
	reqStart := httptest.NewRequest(http.MethodPost, "/api/dataset-worker/start", strings.NewReader(jsonPayload))
	reqStart.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recStart := httptest.NewRecorder()
	cStart := e.NewContext(reqStart, recStart)
	_ = StartDatasetWorkerHandler(cStart)

	// Now get logs
	req := httptest.NewRequest(http.MethodGet, "/api/dataset-worker/logs", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(req, w)

	err := LogsDatasetWorkerHandler(c)
	
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"stdout":""`)
	assert.Contains(t, w.Body.String(), `"stderr":""`)
}

func TestLogsDatasetWorkerHandler_NotRunning(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/dataset-worker/logs", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(req, w)

	_ = LogsDatasetWorkerHandler(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestJobsDatasetWorkerHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/dataset-worker/jobs", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(req, w)

	err := JobsDatasetWorkerHandler(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}
