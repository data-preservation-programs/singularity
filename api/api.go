package api

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	fs2 "github.com/rclone/rclone/fs"
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
	"unicode"

	_ "github.com/data-preservation-programs/singularity/api/docs"
	"github.com/data-preservation-programs/singularity/cmd/embed"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/handler"
	datasource2 "github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
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
// @Tags Data Source
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
		ScannedAt: now,
		SourceID:  source.ID,
		// TODO Type:         model.File,
		Path:                      dstPath,
		Size:                      written,
		LastModifiedTimestampNano: lastModified.UnixNano(),
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
				// TODO RootDirectoryID:      root.ID,
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
	// TODO Type     model.ItemType `json:"type"`
	Path     string `json:"path"`
	SourceID uint32 `json:"sourceId"`
}

// PushItem godoc
// @Summary Push an item to the staging area
// @Description Push an item to the staging area
// @Tags Data Source
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
	reader, _, err := s.contentProvider.FindPiece(c.Request().Context(), pieceCID)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}
	return c.JSON(http.StatusOK, reader)
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

	handler, err := s.datasourceHandlerResolver.Resolve(c.Request().Context(), source)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}

	// TODO: source and item does not match

	// item is not a subpath of source path
	if !strings.HasPrefix(itemInfo.Path, source.Path) {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: item path %s is not a subpath of source path %s", itemInfo.Path, source.Path))
	}

	entry, err := handler.Check(c.Request().Context(), itemInfo.Path)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
	}

	err = s.db.WithContext(c.Request().Context()).Create(&model.Item{
		ScannedAt:                 time.Now().UTC(),
		SourceID:                  source.ID,
		Path:                      itemInfo.Path,
		Size:                      entry.Size(),
		LastModifiedTimestampNano: entry.ModTime(c.Request().Context()).UnixNano(),
	}).Error
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}

	return c.String(http.StatusOK, "OK")
}

func Run(c *cli.Context) error {
	db := database.MustOpenFromCLI(c)
	err := model.AutoMigrate(db)
	if err != nil {
		return handler.NewHandlerError(err)
	}
	bind := c.String("bind")
	stagingDir := c.String("staging-dir")
	return Server{db: db, bind: bind, stagingDir: stagingDir,
		datasourceHandlerResolver: &datasource.DefaultHandlerResolver{},
		datasets:                  make(map[string]model.Source),
		contentProvider:           &service.ContentProviderService{DB: db, Resolver: &datasource.DefaultHandlerResolver{}},
	}.Run(c)
}

func (s Server) toEchoHandler(handlerFunc interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		handlerFuncValue := reflect.ValueOf(handlerFunc)
		handlerFuncType := handlerFuncValue.Type()

		// Check the number of input parameters
		if handlerFuncType.NumIn() == 0 || handlerFuncType.In(0) != reflect.TypeOf(s.db) {
			logger.Error("Invalid handler function signature.")
			return echo.NewHTTPError(http.StatusInternalServerError, "Invalid handler function signature.")
		}

		// Prepare input parameters
		inputParams := []reflect.Value{reflect.ValueOf(s.db.WithContext(c.Request().Context()))}

		// Get path parameters
		for i := 1; i < handlerFuncType.NumIn(); i++ {
			paramType := handlerFuncType.In(i)
			if paramType == reflect.TypeOf(c.Request().Context()) {
				inputParams = append(inputParams, reflect.ValueOf(c.Request().Context()))
				continue
			}
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
				return err.HTTPResponse(c)
			}
			return c.NoContent(http.StatusNoContent)
		}

		// Handle the returned error
		if err, ok := results[1].Interface().(*handler.Error); ok && err != nil {
			return err.HTTPResponse(c)
		}

		// Handle the returned data
		data := results[0].Interface()
		if data == nil {
			return c.NoContent(http.StatusNoContent)
		}
		return c.JSON(http.StatusOK, data)
	}
}

func lowerCamelToSnake(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}

func (s Server) HandlePostSource(c echo.Context) error {
	t := c.Param("type")
	datasetName := c.Param("datasetName")
	dataset, err := database.FindDatasetByName(s.db.WithContext(c.Request().Context()), datasetName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Dataset not found.")
	}
	if err != nil {
		return errors.Wrap(err, "Failed to find dataset.")
	}

	r, err := fs2.Find(t)
	if err != nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("Error: %s", err.Error()))
	}
	body := map[string]string{}
	err = c.Bind(&body)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
	}
	path := body["sourcePath"]
	if path == "" {
		return c.String(http.StatusBadRequest, "Error: sourcePath is required")
	}
	if r.Prefix == "local" {
		path, err = filepath.Abs(path)
		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("failed to get absolute path: %s", err.Error()))
		}
	}
	deleteAfterExportStr := body["deleteAfterExport"]
	deleteAfterExport := false
	if deleteAfterExportStr != "" {
		deleteAfterExport, err = strconv.ParseBool(deleteAfterExportStr)
		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("failed to parse deleteAfterExport: %s", err.Error()))
		}
	}
	delete(body, "sourcePath")
	delete(body, "deleteAfterExport")
	config := map[string]string{}
	for k, v := range body {
		config[lowerCamelToSnake(k)] = v
	}

	source := model.Source{
		DatasetID:           dataset.ID,
		Type:                r.Prefix,
		Path:                path,
		Metadata:            model.Metadata(config),
		ScanIntervalSeconds: 0,
		ScanningState:       model.Ready,
		DeleteAfterExport:   deleteAfterExport,
		DagGenState:         model.Created,
	}

	handler, err := datasource.DefaultHandlerResolver{}.Resolve(c.Request().Context(), source)
	if err != nil {
		return errors.Wrap(err, "failed to resolve handler")
	}

	_, err = handler.List(c.Request().Context(), "")
	if err != nil {
		return errors.Wrap(err, "failed to check source")
	}

	dir := model.Directory{
		Name: path,
	}
	err = s.db.WithContext(c.Request().Context()).Create(&dir).Error
	if err != nil {
		return errors.Wrap(err, "failed to create directory")
	}

	// TODO source.RootDirectoryID = dir.ID
	err = s.db.WithContext(c.Request().Context()).Create(&source).Error
	if err != nil {
		return errors.Wrap(err, "failed to create source")
	}

	return c.JSON(http.StatusOK, source)
}

func (s Server) setupRoutes(e *echo.Echo) {
	// Admin
	e.POST("/api/admin/reset", s.toEchoHandler(admin.ResetHandler))
	e.POST("/api/admin/init", s.toEchoHandler(admin.InitHandler))

	// Dataset
	e.POST("/api/dataset", s.toEchoHandler(dataset.CreateHandler))
	e.PATCH("/api/dataset/:datasetName", s.toEchoHandler(dataset.UpdateHandler))
	e.DELETE("/api/dataset/:datasetName", s.toEchoHandler(dataset.RemoveHandler))
	e.POST("/api/dataset/:datasetName/piece", s.toEchoHandler(dataset.AddPieceHandler))
	e.GET("/api/datasets", s.toEchoHandler(dataset.ListHandler))
	e.GET("/api/dataset/:datasetName/pieces", s.toEchoHandler(dataset.ListPiecesHandler))

	// Wallet
	e.POST("/api/wallet", s.toEchoHandler(wallet.ImportHandler))
	e.GET("/api/wallets", s.toEchoHandler(wallet.ListHandler))
	e.POST("/api/wallet/remote", s.toEchoHandler(wallet.AddRemoteHandler))
	e.DELETE("/api/wallet/:address", s.toEchoHandler(wallet.RemoveHandler))

	// Data source
	e.POST("/api/dataset/:datasetName/source/:type", s.HandlePostSource)
	e.GET("/api/sources", func(c echo.Context) error {
		datasetName := c.QueryParam("dataset")
		sources, err := datasource2.ListSourceHandler(s.db.WithContext(c.Request().Context()), datasetName)
		if err != nil {
			return err.HTTPResponse(c)
		}
		return c.JSON(http.StatusOK, sources)
	})
	e.PATCH("/api/source/:id", s.toEchoHandler(datasource2.UpdateSourceHandler))
	e.POST("/api/source/:id/rescan", s.toEchoHandler(datasource2.RescanSourceHandler))
	// Data source status
	e.DELETE("/api/source/:id", s.toEchoHandler(datasource2.RemoveSourceHandler))
	e.POST("/api/source/:id/check", s.toEchoHandler(datasource2.CheckSourceHandler))
	e.GET("/api/source/:id/summary", s.toEchoHandler(datasource2.GetSourceStatusHandler))
	e.GET("/api/source/:id/chunks", s.toEchoHandler(inspect.GetSourceChunksHandler))
	e.GET("/api/source/:id/items", s.toEchoHandler(inspect.GetSourceItemsHandler))

	e.POST("/api/deal/send_manual", s.toEchoHandler(deal.SendManualHandler))

	e.POST("/api/deal/schedule", s.toEchoHandler(schedule.CreateHandler))

	e.GET("/api/deal/schedules", s.toEchoHandler(schedule.ListHandler))

	e.POST("/api/deal/schedule/:id/pause", s.toEchoHandler(schedule.PauseHandler))

	e.POST("/api/deal/schedule/:id/resume", s.toEchoHandler(schedule.ResumeHandler))

	e.POST("/api/dataset/:name/wallet/:wallet", s.toEchoHandler(wallet.AddWalletHandler))

	e.GET("/api/dataset/:name/wallets", s.toEchoHandler(wallet.ListWalletHandler))

	e.DELETE("/api/dataset/:name/wallet/:wallet", s.toEchoHandler(wallet.RemoveWalletHandler))

	e.POST("/api/deal/list", s.toEchoHandler(deal.ListHandler))

	e.GET("/api/piece/metadata/:piece", s.GetMetadataHandler)

	e.POST("/api/dataset/upload", s.UploadFile)

	e.POST("/api/dataset/push", s.PushItem)
}

var logger = logging.Logger("api")

func (s Server) Run(c *cli.Context) error {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			uri := v.URI
			status := v.Status
			latency := time.Since(v.StartTime)
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
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	s.setupRoutes(e)
	efs, err := fs.Sub(embed.DashboardStaticFiles, "build")
	if err != nil {
		return err
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/api/datasets", s.GetDatasets)
	e.GET("/api/dataset/:id/sources", s.GetSources)
	e.GET("/api/source/:id/cars", s.GetCars)
	e.GET("/api/car/:id/items", s.GetItems)
	e.GET("/api/car/:id/deals", s.GetDealsForCar)
	e.GET("/api/item/:id/deals", s.GetDealsForItem)
	e.GET("/api/directory/:id/entries", s.GetDirectoryEntries)
	e.GET("/api/dataset/:id/deal_stats", s.GetDealStats)
	e.GET("/api/deal_stats", s.GetOverallDealStats)
	e.GET("/*", echo.WrapHandler(http.FileServer(http.FS(efs))))
	return e.Start(s.bind)
}

func (s Server) GetOverallDealStats(c echo.Context) error {
	var stats []DealStats
	err := s.db.Table("deals").
		Select("provider, state, DATE(sector_start) as day, SUM(piece_size) as deal_size").
		Group("provider, state, day").
		Find(&stats).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, stats)
}

func (s Server) GetDealStats(c echo.Context) error {
	datasetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	var stats []DealStats

	err = s.db.Table("deals").
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

func (s Server) GetDatasets(c echo.Context) error {
	var datasets []model.Dataset
	err := s.db.Find(&datasets).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, datasets)
}

func (s Server) GetSources(c echo.Context) error {
	datasetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	var sources []model.Source
	err = s.db.Where("dataset_id = ?", datasetID).Find(&sources).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, sources)
}

func (s Server) GetCars(c echo.Context) error {
	sourceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	var cars []model.Car
	err = s.db.Where("chunk_id in (?)",
		s.db.Table("chunks").Where("source_id", sourceID).Select("id"),
	).Find(&cars).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, cars)
}

func (s Server) GetItems(c echo.Context) error {
	carID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	var items []model.Item
	err = s.db.Where("chunk_id in (?)",
		s.db.Table("cars").Where("id = ?", carID).Select("chunk_id")).Find(&items).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, items)
}

func (s Server) GetDealsForCar(c echo.Context) error {
	carID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	var deals []model.Deal
	err = s.db.Where("piece_cid in (?)",
		s.db.Table("cars").Where("id = ?", carID).Select("piece_cid")).Find(&deals).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, deals)
}

func (s Server) GetDealsForItem(c echo.Context) error {
	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	var deals []model.Deal
	err = s.db.Where("piece_cid in (?)",
		s.db.Table("cars").Where("chunk_id in (?)",
			s.db.Table("items").Where("id = ?", itemID).Select("chunk_id")).
			Select("piece_cid")).Find(&deals).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, deals)
}

func (s Server) GetDirectoryEntries(c echo.Context) error {
	directoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	var dirs []model.Directory
	err = s.db.Where("parent_id = ?", directoryID).Find(&dirs).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	var items []model.Item
	err = s.db.Where("directory_id = ?", directoryID).Find(&items).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, map[string]interface{}{
		"Directories": dirs,
		"Items":       items,
	})
}
