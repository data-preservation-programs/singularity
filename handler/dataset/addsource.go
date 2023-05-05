package dataset

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/datasource"
	"github.com/data-preservation-programs/go-singularity/encryption"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
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
) (*model.Source, error) {
	logger := log.Logger("cli")
	log.SetAllLoggers(log.LevelInfo)
	if request.DatasetName == "" {
		return nil, cli.Exit("dataset name is required", 1)
	}

	if request.SourcePath == "" {
		return nil, cli.Exit("source path is required", 1)
	}

	headers := map[string]string{}
	for _, header := range request.HTTPHeaders {
		parts := strings.SplitN(header, "=", 2)
		if len(parts) != 2 {
			return nil, cli.Exit("invalid header: "+header, 1)
		}

		headers[parts[0]] = parts[1]
	}

	dataset, err := database.FindDatasetByName(db, request.DatasetName)
	if err != nil {
		return nil, cli.Exit("failed to find dataset: "+err.Error(), 1)
	}

	sourceType, sourcePath, err := datasource.ResolveSourceType(request.SourcePath)
	if err != nil {
		return nil, cli.Exit("failed to resolve source type: "+err.Error(), 1)
	}

	var metadata model.Metadata
	secret, err := encryption.EncryptToBase64String([]byte(request.S3SecretAccessKey))
	if err != nil {
		return nil, cli.Exit("failed to encrypt secret access key: "+err.Error(), 1)
	}
	if sourceType == model.S3Path {
		m := model.S3Metadata{
			Region:          request.S3Region,
			Endpoint:        request.S3Endpoint,
			AccessKeyID:     request.S3AccessKeyID,
			SecretAccessKey: secret,
		}
		metadata, err = m.Encode()
		if err != nil {
			return nil, cli.Exit("failed to encode metadata: "+err.Error(), 1)
		}
	} else if sourceType == model.Website {
		m := model.HTTPMetadata{
			Headers: headers,
		}
		metadata, err = m.Encode()
		if err != nil {
			return nil, cli.Exit("failed to encode metadata: "+err.Error(), 1)
		}
	}

	if request.ScanInterval == 0 && request.PushOnly {
		return nil, cli.Exit("scan interval is required to handle pushed data", 1)
	}

	dir := model.Directory{
		Name: sourcePath,
	}
	err = db.Create(&dir).Error
	if err != nil {
		return nil, cli.Exit("failed to create root directory: "+err.Error(), 1)
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
		return nil, cli.Exit("failed to create source: "+err.Error(), 1)
	}
	logger.Infof("created source %d", source.ID)
	return &source, nil
}
