package api

import (
	"bytes"
	"encoding/json"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandlePostSource(t *testing.T) {
	db := database.OpenInMemory()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/dataset/test/source/local", bytes.NewBuffer([]byte(
		`{"deleteAfterExport":true,"sourcePath":"/tmp","rescanInterval":"1h","caseInsensitive":"false"}`,
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
	err := db.Create(&model.Dataset{
		Name:      "test",
		MaxSize:   3 * 1024 * 1024,
		PieceSize: 4 * 1024 * 1024,
	}).Error
	assert.NoError(t, err)
	err = server.HandlePostSource(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var sources []model.Source
	err = db.Find(&sources).Error
	assert.NoError(t, err)
	assert.Len(t, sources, 1)
	assert.EqualValues(t, 1, sources[0].DatasetID)
	assert.Equal(t, "/tmp", sources[0].Path)
	assert.EqualValues(t, 3600, sources[0].ScanIntervalSeconds)
	assert.True(t, sources[0].DeleteAfterExport)
}

func TestPushItem_InvalidID(t *testing.T) {
	db := database.OpenInMemory()
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
	err := server.PushItem(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid source ID")
}

func TestPushItem_InvalidPayload(t *testing.T) {
	db := database.OpenInMemory()
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
	err := server.PushItem(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Syntax error")
}

func TestPushItem_SourceNotFound(t *testing.T) {
	db := database.OpenInMemory()
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
	err := server.PushItem(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "source 1 not found")
}

func TestPushItem_EntryNotFound(t *testing.T) {
	db := database.OpenInMemory()
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
	err := db.Create(&model.Dataset{
		Name:      "test",
		MaxSize:   3 * 1024 * 1024,
		PieceSize: 4 * 1024 * 1024,
	}).Error
	assert.NoError(t, err)
	temp := t.TempDir()
	err = db.Create(&model.Source{
		DatasetID: 1,
		Type:      "local",
		Path:      temp,
	}).Error
	assert.NoError(t, err)
	err = db.Create(&model.Directory{SourceID: 1}).Error
	assert.NoError(t, err)
	err = server.PushItem(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "object not found")
}

func TestPushItem_DuplicateItem(t *testing.T) {
	db := database.OpenInMemory()
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
	err := db.Create(&model.Dataset{
		Name:      "test",
		MaxSize:   3 * 1024 * 1024,
		PieceSize: 4 * 1024 * 1024,
	}).Error
	assert.NoError(t, err)
	temp := t.TempDir()
	err = db.Create(&model.Source{
		DatasetID: 1,
		Type:      "local",
		Path:      temp,
	}).Error
	assert.NoError(t, err)
	err = db.Create(&model.Directory{SourceID: 1}).Error
	assert.NoError(t, err)
	err = os.WriteFile(temp+"/test.txt", []byte("test"), 0644)
	assert.NoError(t, err)
	err = server.PushItem(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	var newItem model.Item
	err = json.Unmarshal(rec.Body.Bytes(), &newItem)
	assert.NoError(t, err)
	assert.Equal(t, "test.txt", newItem.Path)
	assert.Len(t, newItem.ItemParts, 1)

	req = httptest.NewRequest(http.MethodPost, "/api/source/1/push", bytes.NewBuffer([]byte(`{"path":"test.txt"}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/api/source/:id/push")
	c.SetParamNames("id")
	c.SetParamValues("1")
	err = server.PushItem(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusConflict, rec.Code)
	assert.Contains(t, rec.Body.String(), "already exists")
}

func TestPushItem(t *testing.T) {
	db := database.OpenInMemory()
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
	err := db.Create(&model.Dataset{
		Name:      "test",
		MaxSize:   3 * 1024 * 1024,
		PieceSize: 4 * 1024 * 1024,
	}).Error
	assert.NoError(t, err)
	temp := t.TempDir()
	err = db.Create(&model.Source{
		DatasetID: 1,
		Type:      "local",
		Path:      temp,
	}).Error
	assert.NoError(t, err)
	err = db.Create(&model.Directory{SourceID: 1}).Error
	assert.NoError(t, err)
	err = os.WriteFile(temp+"/test.txt", []byte("test"), 0644)
	assert.NoError(t, err)
	err = server.PushItem(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	var newItem model.Item
	err = json.Unmarshal(rec.Body.Bytes(), &newItem)
	assert.NoError(t, err)
	assert.Equal(t, "test.txt", newItem.Path)
	assert.Len(t, newItem.ItemParts, 1)
}
