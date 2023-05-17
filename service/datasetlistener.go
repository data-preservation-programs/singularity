package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/data-preservation-programs/go-singularity/datasource"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type DatasetListenerService struct {
	db                        *gorm.DB
	stagingDir                string
	bind                      string
	datasets                  map[string]model.Source
	logger                    *log.ZapEventLogger
	datasourceHandlerResolver datasource.HandlerResolver
}

func (s DatasetListenerService) getSource(ctx context.Context, datasetName string) (model.Source, error) {
	if source, ok := s.datasets[datasetName]; ok {
		return source, nil
	}
	var dataset model.Dataset
	err := s.db.WithContext(ctx).Where("name = ?", datasetName).First(&dataset).Error
	if err != nil {
		return model.Source{}, err
	}

	var source model.Source
	err = s.db.WithContext(ctx).Transaction(func(db *gorm.DB) error {
		err := db.Where("dataset_id = ? AND type = ?", dataset.ID, model.Upload).First(&source).Error
		if err == nil {
			return nil
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			root := model.Directory{}
			err = db.Create(&root).Error
			if err != nil {
				return err
			}
			source = model.Source{
				DatasetID:            dataset.ID,
				Type:                 model.Upload,
				Path:                 filepath.Join(s.stagingDir, encodeDatasetName(datasetName)),
				ScanIntervalSeconds:  0,
				ScanningState:        "",
				LastScannedTimestamp: time.Now().Unix(),
				RootDirectoryID:      root.ID,
			}
			err = db.Create(&source).Error
			if err != nil {
				return err
			}
			return nil
		}

		return err
	})
	if err != nil {
		return model.Source{}, err
	}
	s.datasets[datasetName] = source
	return source, nil
}

func NewDatasetListenerService(db *gorm.DB, stagingDir string, bind string) DatasetListenerService {
	return DatasetListenerService{
		db:                        db,
		stagingDir:                stagingDir,
		bind:                      bind,
		datasets:                  map[string]model.Source{},
		logger:                    log.Logger("dataset-listener"),
		datasourceHandlerResolver: datasource.NewDefaultHandlerResolver(),
	}
}

func (s DatasetListenerService) Start() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/upload", s.uploadFile)
	e.GET("/push", func(c echo.Context) error {
		t := c.QueryParam("type")
		if !slices.Contains(model.ItemTypes, model.ItemType(t)) {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Error: invalid type %s", t))
		}
		path := c.QueryParam("path")
		sourceIDStr := c.QueryParam("source_id")
		sourceID, err := strconv.ParseUint(sourceIDStr, 10, 32)
		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Error: invalid source_id %s", sourceIDStr))
		}
		itemInfo := ItemInfo{
			Type:     model.ItemType(t),
			Path:     path,
			SourceID: uint32(sourceID),
		}
		return s.pushItem(c, itemInfo)
	})
	e.POST("/push", func(c echo.Context) error {
		var itemInfo ItemInfo
		err := c.Bind(&itemInfo)
		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		}
		return s.pushItem(c, itemInfo)
	})
	err := e.Start(s.bind)
	if err != nil {
		panic(err)
	}
}

type ItemInfo struct {
	Type     model.ItemType `json:"type"`
	Path     string         `json:"path"`
	SourceID uint32         `json:"sourceId"`
}

func (s DatasetListenerService) pushItem(c echo.Context, itemInfo ItemInfo) error {
	var source model.Source
	err := s.db.WithContext(c.Request().Context()).Where("id = ?", itemInfo.SourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: source %d not found.", itemInfo.SourceID))
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}

	handler, err := s.datasourceHandlerResolver.GetHandler(source)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}

	// source and item does not match
	if source.Type.GetSupportedItemType() != itemInfo.Type {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: source %d does not support item type %s", itemInfo.SourceID, itemInfo.Type))
	}

	// item is not a subpath of source path
	if !strings.HasPrefix(itemInfo.Path, source.Path) {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: item path %s is not a subpath of source path %s", itemInfo.Path, source.Path))
	}

	size, lastModified, err := handler.CheckItem(c.Request().Context(), itemInfo.Path)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
	}

	err = s.db.WithContext(c.Request().Context()).Create(&model.Item{
		ScannedAt:    time.Now().UTC(),
		SourceID:     source.ID,
		Type:         itemInfo.Type,
		Path:         itemInfo.Path,
		Size:         size,
		Length:       size,
		LastModified: lastModified,
	}).Error
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}

	return c.String(http.StatusOK, "OK")
}

func encodeDatasetName(str string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9_-]+`)
	encoded := reg.ReplaceAllString(str, "_")

	hashBytes := sha256.Sum256([]byte(str))
	hashStr := hex.EncodeToString(hashBytes[:])

	return encoded + "-" + hashStr[:8]
}

func (s DatasetListenerService) uploadFile(c echo.Context) error {
	datasetName := c.QueryParams().Get("dataset")
	if datasetName == "" {
		return c.String(http.StatusBadRequest, "Error: dataset name is required. Use &dataset=<dataset_name> in the query string.")
	}

	source, err := s.getSource(c.Request().Context(), datasetName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: dataset %s not found.", datasetName))
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
	}

	filename := file.Filename
	itemPath := filepath.Clean("/" + filename)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to get the best staging directory: %s", err.Error()))
	}
	dstPath := filepath.Join(s.stagingDir, encodeDatasetName(datasetName), itemPath)
	// ensure the directory exists
	if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}
	defer dst.Close()

	var written int64
	if written, err = io.Copy(dst, src); err != nil {
		os.Remove(dstPath)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}

	now := time.Now().UTC()
	lastModified := now
	fileInfo, err := os.Stat(dstPath)
	if err != nil {
		s.logger.Errorw("Failed to get file info", "error", err.Error())
	} else {
		lastModified = fileInfo.ModTime().UTC()
	}

	s.db.WithContext(c.Request().Context()).Create(&model.Item{
		ScannedAt:    now,
		SourceID:     source.ID,
		Type:         model.File,
		Path:         dstPath,
		Size:         uint64(written),
		Offset:       0,
		Length:       uint64(written),
		LastModified: &lastModified,
		Version:      0,
	})

	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully to %s.", file.Filename, dstPath))
}
