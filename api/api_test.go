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
