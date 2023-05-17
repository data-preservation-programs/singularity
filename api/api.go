package api

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	_ "github.com/data-preservation-programs/go-singularity/api/docs"
	"github.com/data-preservation-programs/go-singularity/cmd/embed"
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/datasource"
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/handler/dataset"
	"github.com/data-preservation-programs/go-singularity/handler/deal"
	"github.com/data-preservation-programs/go-singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/go-singularity/handler/wallet"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/data-preservation-programs/go-singularity/service"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type DealStats struct {
	Provider string
	State    model.DealState
	Day      string
	DealSize int64
}

type Server struct {
	db                        *gorm.DB
	bind                      string
	stagingDir                string
	datasets                  map[string]model.Source
	datasourceHandlerResolver datasource.HandlerResolver
	contentProvider           *service.ContentProviderService
}

// UploadFile godoc
// @Summary Upload a file to a dataset
// @Description Upload a file to a dataset
// @Tags Dataset
// @Accept mpfd
// @Produce text/plain
// @Param dataset query string true "Dataset name"
// @Param file formData file true "File to upload"
// @Success 200 {string} string "File uploaded"
// @Failure 400 {string} string "Error: dataset name is required."
// @Failure 400 {string} string "Error: file is required."
// @Failure 400 {string} string "Error: dataset not found."
// @Failure 500 {string} string "Error: internal server error."
// @Router /dataset/upload [post]
func (s Server) UploadFile(c echo.Context) error {
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
		logger.Errorw("Failed to get file info", "error", err.Error())
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

func encodeDatasetName(str string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9_-]+`)
	encoded := reg.ReplaceAllString(str, "_")

	hashBytes := sha256.Sum256([]byte(str))
	hashStr := hex.EncodeToString(hashBytes[:])

	return encoded + "-" + hashStr[:8]
}

func (s Server) getSource(ctx context.Context, datasetName string) (model.Source, error) {
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

type ItemInfo struct {
	Type     model.ItemType `json:"type"`
	Path     string         `json:"path"`
	SourceID uint32         `json:"sourceId"`
}

// PushItem godoc
// @Summary Push an item to the staging area
// @Description Push an item to the staging area
// @Tags Dataset
// @Accept json
// @Produce json
// @Param item body ItemInfo true "Item"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /dataset/push [post]
func (s Server) PushItem(c echo.Context) error {
	var itemInfo ItemInfo
	err := c.Bind(&itemInfo)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
	}
	return s.pushItem(c, itemInfo)
}

// GetMetadataHandler godoc
// @Summary Get metadata for a piece
// @Description Get metadata for a piece
// @Tags Piece
// @Accept json
// @Produce json
// @Param piece path string true "Piece CID"
// @Success 200 {object} store.PieceReader
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /piece/metadata/{piece} [get]
func (s Server) GetMetadataHandler(c echo.Context) error {
	piece := c.Param("piece")
	if piece == "" {
		return c.String(http.StatusBadRequest, "Error: missing piece")
	}
	pieceCID, err := cid.Parse(piece)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
	}
	pieceReader, _, err := s.contentProvider.FindPieceAsPieceReader(c.Request().Context(), pieceCID)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}
	return c.JSON(http.StatusOK, pieceReader)
}

func (s Server) pushItem(c echo.Context, itemInfo ItemInfo) error {
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

func Run(c *cli.Context) error {
	db := database.MustOpenFromCLI(c)
	err := model.InitializeEncryption(c.String("password"), db)
	if err != nil {
		return err
	}
	bind := c.String("bind")
	stagingDir := c.String("staging-dir")
	return Server{db: db, bind: bind, stagingDir: stagingDir,
		datasourceHandlerResolver: datasource.NewDefaultHandlerResolver(),
		datasets:                  make(map[string]model.Source),
		contentProvider:           &service.ContentProviderService{DB: db, Resolver: datasource.NewDefaultHandlerResolver()},
	}.Run(c)
}

func (d Server) toEchoHandler(handlerFunc interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		handlerFuncValue := reflect.ValueOf(handlerFunc)
		handlerFuncType := handlerFuncValue.Type()

		// Check the number of input parameters
		if handlerFuncType.NumIn() == 0 || handlerFuncType.In(0) != reflect.TypeOf(d.db) {
			logger.Error("Invalid handler function signature.")
			return echo.NewHTTPError(http.StatusInternalServerError, "Invalid handler function signature.")
		}

		// Prepare input parameters
		inputParams := []reflect.Value{reflect.ValueOf(d.db.WithContext(c.Request().Context()))}

		// Get path parameters
		for i := 1; i < handlerFuncType.NumIn(); i++ {
			paramType := handlerFuncType.In(i)
			if paramType.Kind() == reflect.String {
				if len(c.ParamValues()) < i {
					logger.Error("Invalid handler function signature.")
					return echo.NewHTTPError(http.StatusInternalServerError, "Invalid handler function signature.")
				}
				paramValue := c.ParamValues()[i-1]
				decoded, err := url.QueryUnescape(paramValue)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, "Failed to decode path parameter.")
				}
				inputParams = append(inputParams, reflect.ValueOf(decoded))
				continue
			}

			bodyParam := reflect.New(paramType).Elem()
			if err := c.Bind(bodyParam.Addr().Interface()); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind request body.")
			}
			inputParams = append(inputParams, bodyParam)
			break
		}

		// Call the handler function
		results := handlerFuncValue.Call(inputParams)

		if len(results) == 1 {
			// Handle the returned error
			if err, ok := results[0].Interface().(*handler.Error); ok && err != nil {
				return err.HttpResponse(c)
			}
			return c.NoContent(http.StatusNoContent)
		}

		// Handle the returned error
		if err, ok := results[1].Interface().(*handler.Error); ok && err != nil {
			return err.HttpResponse(c)
		}

		// Handle the returned data
		data := results[0].Interface()
		if data == nil {
			return c.NoContent(http.StatusNoContent)
		}
		return c.JSON(http.StatusOK, data)
	}
}

func (d Server) setupRoutes(e *echo.Echo) {
	e.GET("/admin/api/piece/metadata/:piece", d.GetMetadataHandler)

	e.POST("/admin/api/dataset/upload", d.UploadFile)

	e.POST("/admin/api/dataset/push", d.PushItem)

	e.POST("/admin/api/init", d.toEchoHandler(handler.InitHandler))

	e.POST("/admin/api/dataset", d.toEchoHandler(dataset.CreateHandler))

	e.GET("/admin/api/datasets", d.toEchoHandler(dataset.ListHandler))

	e.DELETE("/admin/api/dataset/:name", d.toEchoHandler(dataset.RemoveHandler))

	e.POST("/admin/api/dataset/:name/source", d.toEchoHandler(dataset.AddSourceHandler))

	e.GET("/admin/api/dataset/:name/sources", d.toEchoHandler(dataset.ListSourceHandler))

	e.DELETE("/admin/api/dataset/:name/source/:sourcepath", d.toEchoHandler(dataset.RemoveSourceHandler))

	e.POST("/admin/api/dataset/:name/piece", d.toEchoHandler(dataset.AddPieceHandler))

	e.POST("/admin/api/deal/send_manual", d.toEchoHandler(deal.SendManualHandler))

	e.POST("/admin/api/wallet", d.toEchoHandler(wallet.ImportHandler))

	e.GET("/admin/api/wallets", d.toEchoHandler(wallet.ListHandler))

	e.POST("/admin/api/deal/schedule", d.toEchoHandler(schedule.CreateHandler))

	e.GET("/admin/api/deal/schedules", d.toEchoHandler(schedule.ListHandler))

	e.POST("/admin/api/deal/schedule/:id/pause", d.toEchoHandler(schedule.PauseHandler))

	e.POST("/admin/api/deal/schedule/:id/resume", d.toEchoHandler(schedule.ResumeHandler))

	e.POST("/admin/api/dataset/:name/wallet/:wallet", d.toEchoHandler(dataset.AddWalletHandler))

	e.GET("/admin/api/dataset/:name/wallets", d.toEchoHandler(dataset.ListWalletHandler))

	e.POST("/admin/api/wallet/remote", d.toEchoHandler(wallet.AddRemoteHandler))

	e.DELETE("/admin/api/dataset/:name/wallet/:wallet", d.toEchoHandler(dataset.RemoveWalletHandler))

	e.POST("/admin/api/deal/list", d.toEchoHandler(deal.ListHandler))
}

var logger = logging.Logger("api")

func (d Server) Run(c *cli.Context) error {
	e := echo.New()
	current := logging.GetConfig().Level
	if logging.LevelInfo < current {
		logging.SetAllLoggers(logging.LevelInfo)
	}
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			uri := v.URI
			status := v.Status
			latency := time.Now().Sub(v.StartTime)
			err := v.Error
			method := c.Request().Method
			if err != nil {
				logger.With("status", status, "latency_ms", latency.Milliseconds(), "err", err).Error(method + " " + uri)
			} else {
				logger.With("status", status, "latency_ms", latency.Milliseconds()).Info(method + " " + uri)
			}
			return nil
		},
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	d.setupRoutes(e)
	efs, err := fs.Sub(embed.DashboardStaticFiles, "build")
	if err != nil {
		return err
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/api/datasets", d.GetDatasets)
	e.GET("/api/dataset/:id/sources", d.GetSources)
	e.GET("/api/source/:id/cars", d.GetCars)
	e.GET("/api/car/:id/items", d.GetItems)
	e.GET("/api/car/:id/deals", d.GetDealsForCar)
	e.GET("/api/item/:id/deals", d.GetDealsForItem)
	e.GET("/api/directory/:id/entries", d.GetDirectoryEntries)
	e.GET("/api/dataset/:id/deal_stats", d.GetDealStats)
	e.GET("/api/deal_stats", d.GetOverallDealStats)
	e.GET("/*", echo.WrapHandler(http.FileServer(http.FS(efs))))
	return e.Start(d.bind)
}

func (d Server) GetOverallDealStats(c echo.Context) error {
	var stats []DealStats
	err := d.db.Table("deals").
		Select("provider, state, DATE(sector_start) as day, SUM(piece_size) as deal_size").
		Group("provider, state, day").
		Find(&stats).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, stats)
}

func (d Server) GetDealStats(c echo.Context) error {
	datasetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	var stats []DealStats

	err = d.db.Table("deals").
		Select("provider, state, DATE(sector_start) as day, SUM(deals.piece_size) as deal_size").
		Joins("JOIN cars ON deals.piece_cid = cars.piece_cid").
		Where("cars.dataset_id = ?", datasetID).
		Group("provider, state, day").
		Find(&stats).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, stats)
}

func (d Server) GetDatasets(c echo.Context) error {
	var datasets []model.Dataset
	err := d.db.Find(&datasets).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, datasets)
}

func (d Server) GetSources(c echo.Context) error {
	datasetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	var sources []model.Source
	err = d.db.Where("dataset_id = ?", datasetID).Find(&sources).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, sources)
}

func (d Server) GetCars(c echo.Context) error {
	sourceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	var cars []model.Car
	err = d.db.Where("chunk_id in (?)",
		d.db.Table("chunks").Where("source_id", sourceID).Select("id"),
	).Find(&cars).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, cars)
}

func (d Server) GetItems(c echo.Context) error {
	carID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	var items []model.Item
	err = d.db.Where("chunk_id in (?)",
		d.db.Table("cars").Where("id = ?", carID).Select("chunk_id")).Find(&items).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, items)
}

func (d Server) GetDealsForCar(c echo.Context) error {
	carID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	var deals []model.Deal
	err = d.db.Where("piece_cid in (?)",
		d.db.Table("cars").Where("id = ?", carID).Select("piece_cid")).Find(&deals).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, deals)
}

func (d Server) GetDealsForItem(c echo.Context) error {
	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	var deals []model.Deal
	err = d.db.Where("piece_cid in (?)",
		d.db.Table("cars").Where("chunk_id in (?)",
			d.db.Table("items").Where("id = ?", itemID).Select("chunk_id")).
			Select("piece_cid")).Find(&deals).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, deals)
}

func (d Server) GetDirectoryEntries(c echo.Context) error {
	directoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	var dirs []model.Directory
	err = d.db.Where("parent_id = ?", directoryID).Find(&dirs).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	var items []model.Item
	err = d.db.Where("directory_id = ?", directoryID).Find(&items).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, map[string]interface{}{
		"Directories": dirs,
		"Items":       items,
	})
}
