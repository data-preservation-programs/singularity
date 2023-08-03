package datasource

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	fs2 "github.com/rclone/rclone/fs"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

var ValidateSource = true

func CreateDatasourceHandler(
	db *gorm.DB,
	ctx context.Context,
	datasourceHandlerResolver datasource.HandlerResolver,
	sourceType string,
	datasetName string,
	sourceParameters map[string]any,
) (*model.Source, error) {
	return createDatasourceHandler(db, ctx, sourceType, datasetName, sourceParameters)
}

// @Summary Add acd source for a dataset
// @Tags Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {object} model.Source
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body AcdRequest true "Request body"
// @Router /source/acd/dataset/{datasetName} [post]
func createDatasourceHandler(
	db *gorm.DB,
	ctx context.Context,
	sourceType string,
	datasetName string,
	sourceParameters map[string]any,
) (*model.Source, error) {
	dataset, err := database.FindDatasetByName(db, datasetName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NotFoundError{Err: err}
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find dataset: %w", err)
	}

	r, err := fs2.Find(sourceType)
	if err != nil {
		return nil, handler.InvalidParameterError{Err: err}
	}
	sourcePath := sourceParameters["sourcePath"]
	path, ok := sourcePath.(string)
	if !ok {
		return nil, handler.NewInvalidParameterErr("sourcePath needs to be a string")
	}
	if path == "" {
		return nil, handler.NewInvalidParameterErr("sourcePath is required")
	}
	if r.Prefix == "local" {
		path, err = filepath.Abs(path)
		if err != nil {
			return nil, handler.InvalidParameterError{Err: fmt.Errorf("failed to get absolute path: %w", err)}
		}
	}
	deleteAfterExportValue := sourceParameters["deleteAfterExport"]
	deleteAfterExport, ok := deleteAfterExportValue.(bool)
	if !ok {
		return nil, handler.NewInvalidParameterErr("deleteAfterExport needs to be a boolean")
	}
	rescanIntervalValue := sourceParameters["rescanInterval"]
	rescanInterval, ok := rescanIntervalValue.(string)
	if !ok {
		return nil, handler.NewInvalidParameterErr("rescanInterval needs to be a string")
	}
	scanningStateValue := sourceParameters["scanningState"]
	scanningState, ok := scanningStateValue.(string)
	if !ok {
		return nil, handler.NewInvalidParameterErr("scanningState needs to be a string")
	}

	rescan, err := time.ParseDuration(rescanInterval)
	if err != nil {
		return nil, handler.InvalidParameterError{Err: fmt.Errorf("failed to parse rescanInterval: %w", err)}
	}

	var ws model.WorkState
	err = ws.Set(scanningState)
	if err != nil {
		return nil, handler.InvalidParameterError{Err: fmt.Errorf("failed to parse scanningState: %w", err)}
	}
	delete(sourceParameters, "sourcePath")
	delete(sourceParameters, "deleteAfterExport")
	delete(sourceParameters, "rescanInterval")
	delete(sourceParameters, "scanningState")
	delete(sourceParameters, "type")
	delete(sourceParameters, "datasetName")

	config := map[string]string{}
	for k, v := range sourceParameters {
		str, ok := v.(string)
		if !ok {
			return nil, handler.NewInvalidParameterErr(fmt.Sprintf("%s needs to be a string", k))
		}
		optionName := lowerCamelToSnake(k)
		if !underscore.Any(r.Options, func(o fs2.Option) bool {
			return o.Name == optionName
		}) {
			return nil, handler.NewInvalidParameterErr(fmt.Sprintf("%s is not a valid parameter for storage source type %s", k, r.Prefix))
		}
		config[optionName] = str
	}

	source := model.Source{
		DatasetID:           dataset.ID,
		Type:                r.Prefix,
		Path:                path,
		Metadata:            model.Metadata(config),
		ScanIntervalSeconds: uint64(rescan.Seconds()),
		ScanningState:       ws,
		DeleteAfterExport:   deleteAfterExport,
		DagGenState:         model.Created,
	}

	if ValidateSource {
		dsHandler, err := datasource.DefaultHandlerResolver{}.Resolve(ctx, source)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve handler: %w", err)
		}

		_, err = dsHandler.List(ctx, "")
		if err != nil {
			return nil, fmt.Errorf("failed to check source: %w", err)
		}
	}

	err = db.Create(&source).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return nil, handler.NewDuplicateRecordError(fmt.Sprintf("source at %s of type %s for dataset %s already exists", sourcePath, sourceType, datasetName))
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create source: %w", err)
	}

	dir := model.Directory{
		Name:     path,
		SourceID: source.ID,
	}
	err = db.Create(&dir).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	return &source, nil
}
