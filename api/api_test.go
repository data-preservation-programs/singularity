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
	"github.com/data-preservation-programs/singularity/handler/item"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestHandlePostSource(t *testing.T) {
	tmp := os.TempDir()
	tmp = strings.TrimSuffix(tmp, "/")
	db, err := database.OpenInMemory()
	require.NoError(t, err)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/dataset/test/source/local", bytes.NewBuffer([]byte(
		`{"deleteAfterExport":true,"sourcePath":"`+tmp+`","rescanInterval":"1h","caseInsensitive":"false"}`,
	)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/dataset/:datasetName/source/:type")
	c.SetParamNames("datasetName", "type")
	c.SetParamValues("test", "local")
	server := Server{
		db:                        db,
		datasourceHandlerResolver: &datasource.DefaultHandlerResolver{},
	}
	err = db.Create(&model.Dataset{
		Name:      "test",
		MaxSize:   3 * 1024 * 1024,
		PieceSize: 4 * 1024 * 1024,
	}).Error
	require.NoError(t, err)
	err = server.HandlePostSource(c)
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

func TestPushItem_InvalidID(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
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
	err = server.toEchoHandler(item.PushItemHandler)(c)
	require.Error(t, err)
	var httpError *echo.HTTPError
	require.ErrorAs(t, err, &httpError)
	require.Equal(t, http.StatusBadRequest, httpError.Code)
	require.Contains(t, "Failed to parse path parameter as number.", httpError.Message)
}

func TestPushItem_InvalidPayload(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
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
	err = server.toEchoHandler(item.PushItemHandler)(c)
	require.Error(t, err)
	var httpError *echo.HTTPError
	require.ErrorAs(t, err, &httpError)
	require.Equal(t, http.StatusBadRequest, httpError.Code)
	require.Contains(t, "Failed to bind request body.", httpError.Message)
}

func TestPushItem_SourceNotFound(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
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
	err = server.toEchoHandler(item.PushItemHandler)(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Contains(t, rec.Body.String(), "source 1 not found")
}

func TestPushItem_EntryNotFound(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
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
	err = db.Create(&model.Dataset{
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
	err = server.toEchoHandler(item.PushItemHandler)(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Contains(t, rec.Body.String(), "object not found")
}

func TestPushItem_DuplicateItem(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
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
	err = db.Create(&model.Dataset{
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
	err = server.toEchoHandler(item.PushItemHandler)(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
	var newItem model.Item
	err = json.Unmarshal(rec.Body.Bytes(), &newItem)
	require.NoError(t, err)
	require.Equal(t, "test.txt", newItem.Path)
	require.Len(t, newItem.ItemParts, 1)

	req = httptest.NewRequest(http.MethodPost, "/api/source/1/push", bytes.NewBuffer([]byte(`{"path":"test.txt"}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/api/source/:id/push")
	c.SetParamNames("id")
	c.SetParamValues("1")
	err = server.toEchoHandler(item.PushItemHandler)(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusConflict, rec.Code)
	require.Contains(t, rec.Body.String(), "already exists")
}

func TestPushItem(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
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
	err = db.Create(&model.Dataset{
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
	err = server.toEchoHandler(item.PushItemHandler)(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
	var newItem model.Item
	err = json.Unmarshal(rec.Body.Bytes(), &newItem)
	require.NoError(t, err)
	require.Equal(t, "test.txt", newItem.Path)
	require.Len(t, newItem.ItemParts, 1)
}
