package dataset

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddSourceHandler_NoDatasetName(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := AddSourceHandler(db, AddSourceRequest{})
	assert.ErrorContains(err, "dataset name is required")
}

func TestAddSourceHandler_DatasetNotFound(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := AddSourceHandler(db, AddSourceRequest{
		DatasetName: "test",
		SourcePath: "/",
	})
	assert.ErrorContains(err, "failed to find dataset")
}

func TestAddSourceHandler_SourceInvalid(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MinSizeStr: "1GB", MaxSizeStr: "2GB"})
	assert.Nil(err)
	_, err = AddSourceHandler(db, AddSourceRequest{
		DatasetName: "test",
		SourcePath: "not exist",
	})
	assert.ErrorContains(err, "no such file or directory")
}

func TestAddSourceHandler_SourceResolved(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	_, err := CreateHandler(db, CreateRequest{Name: "test", MinSizeStr: "1GB", MaxSizeStr: "2GB"})
	assert.Nil(err)
	source, err := AddSourceHandler(db, AddSourceRequest{
		DatasetName: "test",
		SourcePath: ".",
	})
	assert.Nil(err)
	assert.NotEqual(".", source.Path)
}

func TestAddSourceHandler_WithHTTPHeader(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	dataset, err := CreateHandler(db, CreateRequest{Name: "test", MinSizeStr: "1GB", MaxSizeStr: "2GB"})
	assert.Nil(err)
	_, err = AddSourceHandler(db, AddSourceRequest{
		DatasetName: "test",
		SourcePath: "https://example.com",
		HTTPHeaders: []string{
			"key1=value1",
			"key2=value2",
		},
	})
	var source model.Source
	err = db.Where("dataset_id = ?", dataset.ID).First(&source).Error
	assert.Nil(err)
	metadata, err := source.Metadata.GetHTTPMetadata()
	assert.Nil(err)
	assert.Equal("value1", metadata.Headers["key1"])
	assert.Equal("value2", metadata.Headers["key2"])
}

func TestAddSourceHandler_WithS3Metadata(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	dataset, err := CreateHandler(db, CreateRequest{Name: "test", MinSizeStr: "1GB", MaxSizeStr: "2GB"})
	assert.Nil(err)
	_, err = AddSourceHandler(db, AddSourceRequest{
		DatasetName:       "test",
		SourcePath:        "s3://example.com",
		S3Endpoint:        "endpoint",
		S3AccessKeyID:     "access_key_id",
		S3SecretAccessKey: "secret_access_key",
	})
	var source model.Source
	err = db.Where("dataset_id = ?", dataset.ID).First(&source).Error
	assert.Nil(err)
	metadata, err := source.Metadata.GetS3Metadata()
	assert.Nil(err)
	assert.Equal("access_key_id", metadata.AccessKeyID)
	assert.Equal("endpoint", metadata.Endpoint)
	assert.Equal("secret_access_key", metadata.SecretAccessKey)
}
