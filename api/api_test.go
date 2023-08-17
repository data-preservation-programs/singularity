package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	dshandler "github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestHandlePostSource(t *testing.T) {
	tmp := os.TempDir()
	tmp = strings.TrimSuffix(tmp, "/")
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	e := echo.New()
	payload := map[string]any{
		"deleteAfterExport": true,
		"sourcePath":        tmp,
		"rescanInterval":    "1h",
		"caseInsensitive":   "false",
		"scanningState":     "ready",
	}
	body, err := json.Marshal(payload)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/api/source/local/dataset/test", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/source/:type/dataset/:datasetName")
	c.SetParamNames("type", "datasetName")
	c.SetParamValues("local", "test")
	server := Server{
		db:                        db,
		datasourceHandlerResolver: &datasource.DefaultHandlerResolver{},
	}
	err = db.Create(&model.Preparation{
		Name:      "test",
		MaxSize:   3 * 1024 * 1024,
		PieceSize: 4 * 1024 * 1024,
	}).Error
	require.NoError(t, err)
	err = server.toEchoHandler(dshandler.CreateDatasourceHandler)(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
	var sources []model.Source
	err = db.Find(&sources).Error
	require.NoError(t, err)
	require.Len(t, sources, 1)
	require.EqualValues(t, 1, sources[0].DatasetID)
	require.Equal(t, tmp, sources[0].Path)
	require.EqualValues(t, 3600, sources[0].ScanIntervalSeconds)
	require.True(t, sources[0].DeleteAfterExport)
}

func TestPushFile_InvalidID(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/source/1/push", bytes.NewBuffer([]byte(`{"path":"test.txt"}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/source/:id/push")
	c.SetParamNames("id")
	c.SetParamValues("a")
	server := Server{
		db:                        db,
		datasourceHandlerResolver: &datasource.DefaultHandlerResolver{},
	}
	err = server.toEchoHandler(dshandler.PushFileHandler)(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Contains(t, rec.Body.String(), "failed to parse path parameter as number")
}

func TestPushFile_InvalidPayload(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/source/1/push", bytes.NewBuffer([]byte(`invalid payload`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/source/:id/push")
	c.SetParamNames("id")
	c.SetParamValues("1")
	server := Server{
		db:                        db,
		datasourceHandlerResolver: &datasource.DefaultHandlerResolver{},
	}
	err = server.toEchoHandler(dshandler.PushFileHandler)(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Contains(t, rec.Body.String(), "failed to bind request body")
}

func TestPushFile_SourceNotFound(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/source/1/push", bytes.NewBuffer([]byte(`{"path":"test.txt"}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/source/:id/push")
	c.SetParamNames("id")
	c.SetParamValues("1")
	server := Server{
		db:                        db,
		datasourceHandlerResolver: &datasource.DefaultHandlerResolver{},
	}
	err = server.toEchoHandler(dshandler.PushFileHandler)(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Contains(t, rec.Body.String(), "source 1 not found")
}

func TestPushFile_EntryNotFound(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/source/1/push", bytes.NewBuffer([]byte(`{"path":"test.txt"}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/source/:id/push")
	c.SetParamNames("id")
	c.SetParamValues("1")
	server := Server{
		db:                        db,
		datasourceHandlerResolver: &datasource.DefaultHandlerResolver{},
	}
	err = db.Create(&model.Preparation{
		Name:      "test",
		MaxSize:   3 * 1024 * 1024,
		PieceSize: 4 * 1024 * 1024,
	}).Error
	require.NoError(t, err)
	temp := t.TempDir()
	err = db.Create(&model.Source{
		DatasetID: 1,
		Type:      "local",
		Path:      temp,
	}).Error
	require.NoError(t, err)
	err = db.Create(&model.Directory{SourceID: 1}).Error
	require.NoError(t, err)
	err = server.toEchoHandler(dshandler.PushFileHandler)(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Contains(t, rec.Body.String(), "object not found")
}

func TestPushFile_Duplicate(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/source/1/push", bytes.NewBuffer([]byte(`{"path":"test.txt"}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/source/:id/push")
	c.SetParamNames("id")
	c.SetParamValues("1")
	server := Server{
		db:                        db,
		datasourceHandlerResolver: &datasource.DefaultHandlerResolver{},
	}
	err = db.Create(&model.Preparation{
		Name:      "test",
		MaxSize:   3 * 1024 * 1024,
		PieceSize: 4 * 1024 * 1024,
	}).Error
	require.NoError(t, err)
	temp := t.TempDir()
	err = db.Create(&model.Source{
		DatasetID: 1,
		Type:      "local",
		Path:      temp,
	}).Error
	require.NoError(t, err)
	err = db.Create(&model.Directory{SourceID: 1}).Error
	require.NoError(t, err)
	err = os.WriteFile(temp+"/test.txt", []byte("test"), 0644)
	require.NoError(t, err)
	err = server.toEchoHandler(dshandler.PushFileHandler)(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
	var newFile model.File
	err = json.Unmarshal(rec.Body.Bytes(), &newFile)
	require.NoError(t, err)
	require.Equal(t, "test.txt", newFile.Path)
	require.Len(t, newFile.FileRanges, 1)

	req = httptest.NewRequest(http.MethodPost, "/api/source/1/push", bytes.NewBuffer([]byte(`{"path":"test.txt"}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/api/source/:id/push")
	c.SetParamNames("id")
	c.SetParamValues("1")
	err = server.toEchoHandler(dshandler.PushFileHandler)(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusConflict, rec.Code)
	require.Contains(t, rec.Body.String(), "already exists")
}

func TestPushFile(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/source/1/push", bytes.NewBuffer([]byte(`{"path":"test.txt"}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/source/:id/push")
	c.SetParamNames("id")
	c.SetParamValues("1")
	server := Server{
		db:                        db,
		datasourceHandlerResolver: &datasource.DefaultHandlerResolver{},
	}
	err = db.Create(&model.Preparation{
		Name:      "test",
		MaxSize:   3 * 1024 * 1024,
		PieceSize: 4 * 1024 * 1024,
	}).Error
	require.NoError(t, err)
	temp := t.TempDir()
	err = db.Create(&model.Source{
		DatasetID: 1,
		Type:      "local",
		Path:      temp,
	}).Error
	require.NoError(t, err)
	err = db.Create(&model.Directory{SourceID: 1}).Error
	require.NoError(t, err)
	err = os.WriteFile(temp+"/test.txt", []byte("test"), 0644)
	require.NoError(t, err)
	err = server.toEchoHandler(dshandler.PushFileHandler)(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
	var newFile model.File
	err = json.Unmarshal(rec.Body.Bytes(), &newFile)
	require.NoError(t, err)
	require.Equal(t, "test.txt", newFile.Path)
	require.Len(t, newFile.FileRanges, 1)
}
