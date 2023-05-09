package dataset

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/datasource"
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
	"strings"
	"time"
)

type AddSourceRequest struct {
	DatasetName       string
	SourcePath        string
	ScanInterval      time.Duration
	HTTPHeaders       []string
	S3Region          string
	S3Endpoint        string
	S3AccessKeyID     string
	S3SecretAccessKey string
	PushOnly          bool
}

func AddSourceHandler(
	db *gorm.DB,
	request AddSourceRequest,
) (*model.Source, *handler.Error) {
	logger := log.Logger("cli")
	log.SetAllLoggers(log.LevelInfo)
	if request.DatasetName == "" {
		return nil, handler.NewBadRequestString("dataset name is required")
	}

	if request.SourcePath == "" {
		return nil, handler.NewBadRequestString("source path is required")
	}

	headers := map[string]string{}
	for _, header := range request.HTTPHeaders {
		parts := strings.SplitN(header, "=", 2)
		if len(parts) != 2 {
			return nil, handler.NewBadRequestString("invalid header: " + header)
		}

		headers[parts[0]] = parts[1]
	}

	dataset, err := database.FindDatasetByName(db, request.DatasetName)
	if err != nil {
		return nil, handler.NewBadRequestString("failed to find dataset: " + err.Error())
	}

	sourceType, sourcePath, err := datasource.ResolveSourceType(request.SourcePath)
	if err != nil {
		return nil, handler.NewBadRequestString("failed to resolve source type: " + err.Error())
	}

	var metadata model.Metadata
	if sourceType == model.S3Path {
		m := model.S3Metadata{
			Region:          request.S3Region,
			Endpoint:        request.S3Endpoint,
			AccessKeyID:     request.S3AccessKeyID,
			SecretAccessKey: request.S3SecretAccessKey,
		}
		metadata, err = m.Encode()
		if err != nil {
			return nil, handler.NewBadRequestString("failed to encode metadata: " + err.Error())
		}
	} else if sourceType == model.Website {
		m := model.HTTPMetadata{
			Headers: headers,
		}
		metadata, err = m.Encode()
		if err != nil {
			return nil, handler.NewBadRequestString("failed to encode metadata: " + err.Error())
		}
	}

	if request.ScanInterval == 0 && request.PushOnly {
		return nil, handler.NewBadRequestString("scan interval is required to handle pushed data")
	}

	dir := model.Directory{
		Name: sourcePath,
	}
	err = db.Create(&dir).Error
	if err != nil {
		return nil, handler.NewBadRequestString("failed to create root directory: " + err.Error())
	}

	source := model.Source{
		DatasetID:           dataset.ID,
		Type:                sourceType,
		Path:                sourcePath,
		Metadata:            metadata,
		ScanIntervalSeconds: uint64(request.ScanInterval.Seconds()),
		ScanningState:       model.Ready,
		ScanningWorkerID:    nil,
		ErrorMessage:        "",
		RootDirectoryID:     dir.ID,
		PushOnly:            request.PushOnly,
	}
	err = db.Create(&source).Error
	if err != nil {
		return nil, handler.NewBadRequestString("failed to create source: " + err.Error())
	}
	logger.Infof("created source %d", source.ID)
	return &source, nil
}
