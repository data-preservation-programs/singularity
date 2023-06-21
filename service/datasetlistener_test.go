//go:build exclude

package service

import (
	"bytes"
	"fmt"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"path/filepath"
	"testing"
)

func TestDatasetListenerService_uploadFile(t *testing.T) {
	// Create a temporary directory to use as the staging directory
	tempDir := t.TempDir()

	// Create a new in-memory SQLite database for testing
	db := database.OpenInMemory()
	defer database.DropAll(db)

	// Create a new DatasetListenerService with the temporary directory and database
	ds := NewDatasetListenerService(db, tempDir, "")

	// Create a new Echo router and register the uploadFile handler
	e := echo.New()
	e.POST("/upload", ds.uploadFile)

	// Create a new test server using the Echo router
	ts := httptest.NewServer(e)
	defer ts.Close()

	// Create a test dataset
	dataset := model.Dataset{
		Name: "test",
	}
	assert.NoError(t, db.Create(&dataset).Error)

	// Create a new multipart form with a dummy file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileName := "test.txt"
	fileContents := "this is a test"

	partHeaders := textproto.MIMEHeader{}
	partHeaders.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", fileName))
	partHeaders.Set("Content-Type", "text/plain")

	part, err := writer.CreatePart(partHeaders)
	assert.NoError(t, err)

	_, err = io.WriteString(part, fileContents)
	assert.NoError(t, err)

	assert.NoError(t, writer.Close())

	// Create a new request to the test server with the multipart form
	req, err := http.NewRequest("POST", ts.URL+"/upload?dataset=test", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Perform the request and check the response
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Check that the file was created in the staging directory
	files, err := filepath.Glob(filepath.Join(tempDir, "*", fileName))
	assert.NoError(t, err)
	assert.Len(t, files, 1)

	// Check that the file was added to the database
	var item model.Item
	assert.NoError(t, db.Where("path = ?", files[0]).First(&item).Error)
	assert.Equal(t, uint64(len(fileContents)), item.Size)
}
