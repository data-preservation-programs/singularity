package api

import (
	"context"
	"fmt"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"strconv"
	"syscall"
	"time"
	"unicode"

	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/service/contentprovider"
	"github.com/data-preservation-programs/singularity/service/datasetworker"
	"github.com/data-preservation-programs/singularity/util"
	fs2 "github.com/rclone/rclone/fs"
	"github.com/ybbus/jsonrpc/v3"
	"golang.org/x/net/webdav"

	_ "github.com/data-preservation-programs/singularity/api/docs"
	"github.com/data-preservation-programs/singularity/cmd/embed"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/handler"
	datasource2 "github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/model"
	logging "github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

type Server struct {
	db                        *gorm.DB
	adminBind                 string
	datasourceHandlerResolver datasource.HandlerResolver
	lotusClient               jsonrpc.RPCClient
	dealMaker                 replication.DealMaker
	adminEnabled              bool
	webDavEnabled             bool
	webDavBind                string
	webDavSourceID            uint32
}

type ItemInfo struct {
	Path string `json:"path"` // Path to the new item, relative to the source
}

// @Summary Push an item to be queued
// @Description Tells Singularity that something is ready to be grabbed for data preparation
// @Tags Data Source
// @Accept json
// @Produce json
// @Param item body ItemInfo true "Item"
// @Success 201 {object} model.Item
// @Failure 400 {string} string "Bad Request"
// @Failure 409 {string} string "Item already exists"
// @Failure 500 {string} string "Internal Server Error"
// @Router /source/{id}/push [post]
func (s Server) handlePushItem(c echo.Context) error {
	id := c.Param("id")
	sourceID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid source ID")
	}

	var itemInfo ItemInfo
	err = c.Bind(&itemInfo)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
	}
	return s.pushItem(c, uint32(sourceID), itemInfo)
}

// @Summary Get metadata for a piece
// @Description Get metadata for a piece for how it may be reassembled from the data source
// @Tags Metadata
// @Produce json
// @Param piece path string true "Piece CID"
// @Success 200 {object} store.PieceReader
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /piece/{id}/metadata [get]
func (s Server) getMetadataHandler(c echo.Context) error {
	return contentprovider.GetMetadataHandler(c, s.db)
}

func (s Server) SendManualHandler(c echo.Context) error {
	db := s.db.WithContext(c.Request().Context())
	ctx := c.Request().Context()
	var request deal.Proposal
	err := c.Bind(&request)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
	}
	dealModel, err := deal.SendManualHandler(db, ctx, request, s.dealMaker)
	if err != nil {
		var err2 *handler.Error
		if errors.As(err, &err2) {
			return err2.HTTPResponse(c)
		}
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}
	return c.JSON(http.StatusOK, dealModel)
}

func (s Server) pushItem(c echo.Context, sourceID uint32, itemInfo ItemInfo) error {
	var source model.Source
	err := s.db.WithContext(c.Request().Context()).Preload("Dataset").Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: source %d not found.", sourceID))
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}

	handler, err := s.datasourceHandlerResolver.Resolve(c.Request().Context(), source)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}

	entry, err := handler.Check(c.Request().Context(), itemInfo.Path)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
	}

	obj, ok := entry.(fs2.ObjectInfo)
	if !ok {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s is not an object", itemInfo.Path))
	}

	item, itemParts, err := datasetworker.PushItem(c.Request().Context(), s.db, obj, source, *source.Dataset, map[string]uint64{})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}

	if item == nil {
		return c.String(http.StatusConflict, fmt.Sprintf("Error: %s already exists", obj.Remote()))
	}

	item.ItemParts = itemParts

	return c.JSON(http.StatusCreated, item)
}

func Run(c *cli.Context) error {
	db, err := database.OpenFromCLI(c)
	if err != nil {
		return err
	}
	if err := model.AutoMigrate(db); err != nil {
		return handler.NewHandlerError(err)
	}
	adminBind := c.String("admin-bind")
	adminEnabled := c.Bool("enable-admin")
	webDavBind := c.String("webdav-bind")
	webDavEnabled := c.Bool("enable-webdav")
	webDavSourceID := c.Uint("webdav-source-id")

	lotusAPI := c.String("lotus-api")
	lotusToken := c.String("lotus-token")

	h, err := util.InitHost(nil)
	if err != nil {
		return errors.Wrap(err, "failed to init host")
	}

	return Server{db: db,
		datasourceHandlerResolver: &datasource.DefaultHandlerResolver{},
		lotusClient:               util.NewLotusClient(lotusAPI, lotusToken),
		dealMaker: replication.NewDealMaker(
			util.NewLotusClient(lotusAPI, lotusToken),
			h,
			time.Hour,
			time.Minute*5,
		),
		adminBind:      adminBind,
		adminEnabled:   adminEnabled,
		webDavBind:     webDavBind,
		webDavEnabled:  webDavEnabled,
		webDavSourceID: uint32(webDavSourceID),
	}.Run(c.Context)
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

		var j int
		// Get path parameters
		for i := 1; i < handlerFuncType.NumIn(); i++ {
			paramType := handlerFuncType.In(i)
			if paramType.String() == "context.Context" {
				inputParams = append(inputParams, reflect.ValueOf(c.Request().Context()))
				continue
			}
			if paramType.String() == "jsonrpc.RPCClient" {
				inputParams = append(inputParams, reflect.ValueOf(s.lotusClient))
				continue
			}
			if paramType.Kind() == reflect.String {
				if j >= len(c.ParamValues()) {
					logger.Error("Invalid handler function signature.")
					return echo.NewHTTPError(http.StatusInternalServerError, "Invalid handler function signature.")
				}
				paramValue := c.ParamValues()[j]
				decoded, err := url.QueryUnescape(paramValue)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, "Failed to decode path parameter.")
				}
				inputParams = append(inputParams, reflect.ValueOf(decoded))
				j += 1
				continue
			}

			bodyParam := reflect.New(paramType).Elem()
			if bodyParam.Kind() == reflect.Map {
				bodyParam.Set(reflect.MakeMap(bodyParam.Type()))
			}
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
			err, ok := results[0].Interface().(*handler.Error)
			if results[0].Interface() != nil && !ok {
				err, ok := results[1].Interface().(error)
				if !ok {
					return echo.NewHTTPError(http.StatusInternalServerError, "Invalid handler function signature.")
				}
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			if err != nil {
				return err.HTTPResponse(c)
			}
			return c.NoContent(http.StatusNoContent)
		}

		// Handle the returned error
		err, ok := results[1].Interface().(*handler.Error)
		if results[1].Interface() != nil && !ok {
			err, ok := results[1].Interface().(error)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, "Invalid handler function signature.")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if err != nil {
			return err.HTTPResponse(c)
		}

		// Handle the returned data
		data := results[0].Interface()
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
	body := map[string]interface{}{}
	err = c.Bind(&body)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
	}
	sourcePath := body["sourcePath"]
	path, ok := sourcePath.(string)
	if !ok {
		return c.String(http.StatusBadRequest, "Error: sourcePath needs to be a string")
	}
	if path == "" {
		return c.String(http.StatusBadRequest, "Error: sourcePath is required")
	}
	if r.Prefix == "local" {
		path, err = filepath.Abs(path)
		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("failed to get absolute path: %s", err.Error()))
		}
	}
	deleteAfterExportValue := body["deleteAfterExport"]
	deleteAfterExport, ok := deleteAfterExportValue.(bool)
	if !ok {
		return c.String(http.StatusBadRequest, "Error: deleteAfterExport needs to be a boolean")
	}
	rescanIntervalValue := body["rescanInterval"]
	rescanInterval, ok := rescanIntervalValue.(string)
	if !ok {
		return c.String(http.StatusBadRequest, "Error: rescanInterval needs to be a string")
	}
	rescan, err := time.ParseDuration(rescanInterval)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error: failed to parse rescanInterval: %s", err.Error()))
	}
	delete(body, "sourcePath")
	delete(body, "deleteAfterExport")
	delete(body, "rescanInterval")
	delete(body, "type")
	delete(body, "datasetName")
	config := map[string]string{}
	for k, v := range body {
		str, ok := v.(string)
		if !ok {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s needs to be a string", k))
		}
		config[lowerCamelToSnake(k)] = str
	}

	source := model.Source{
		DatasetID:           dataset.ID,
		Type:                r.Prefix,
		Path:                path,
		Metadata:            model.Metadata(config),
		ScanIntervalSeconds: uint64(rescan.Seconds()),
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

	err = s.db.WithContext(c.Request().Context()).Create(&source).Error
	if err != nil {
		return errors.Wrap(err, "failed to create source")
	}

	dir := model.Directory{
		Name:     path,
		SourceID: source.ID,
	}
	err = s.db.WithContext(c.Request().Context()).Create(&dir).Error
	if err != nil {
		return errors.Wrap(err, "failed to create directory")
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
	e.GET("/api/dataset", s.toEchoHandler(dataset.ListHandler))
	e.GET("/api/dataset/:datasetName/piece", s.toEchoHandler(dataset.ListPiecesHandler))

	// Wallet
	e.POST("/api/wallet", s.toEchoHandler(wallet.ImportHandler))
	e.GET("/api/wallet", s.toEchoHandler(wallet.ListHandler))
	e.POST("/api/wallet/remote", s.toEchoHandler(wallet.AddRemoteHandler))
	e.DELETE("/api/wallet/:address", s.toEchoHandler(wallet.RemoveHandler))

	// Wallet Association
	e.POST("/api/dataset/:datasetName/wallet/:wallet", s.toEchoHandler(wallet.AddWalletHandler))
	e.GET("/api/dataset/:datasetName/wallet", s.toEchoHandler(wallet.ListWalletHandler))
	e.DELETE("/api/dataset/:datasetName/wallet/:wallet", s.toEchoHandler(wallet.RemoveWalletHandler))

	// Data source
	e.POST("/api/source/:type/dataset/:datasetName", s.HandlePostSource)
	e.GET("/api/source", s.toEchoHandler(datasource2.ListSourceHandler))
	e.PATCH("/api/source/:id", s.toEchoHandler(datasource2.UpdateSourceHandler))
	e.DELETE("/api/source/:id", s.toEchoHandler(datasource2.RemoveSourceHandler))
	e.POST("/api/source/:id/rescan", s.toEchoHandler(datasource2.RescanSourceHandler))
	e.POST("/api/source/:id/daggen", s.toEchoHandler(datasource2.DagGenHandler))
	e.POST("/api/source/:id/push", s.handlePushItem)

	// Piece metadata
	e.GET("/api/piece/:id/metadata", s.getMetadataHandler)

	// Data source status
	e.POST("/api/source/:id/check", s.toEchoHandler(datasource2.CheckSourceHandler))
	e.GET("/api/source/:id/summary", s.toEchoHandler(datasource2.GetSourceStatusHandler))
	e.GET("/api/source/:id/chunks", s.toEchoHandler(inspect.GetSourceChunksHandler))
	e.GET("/api/source/:id/items", s.toEchoHandler(inspect.GetSourceItemsHandler))
	e.GET("/api/source/:id/dags", s.toEchoHandler(inspect.GetDagsHandler))
	e.GET("/api/source/:id/path", s.toEchoHandler(inspect.GetPathHandler))
	e.GET("/api/chunk/:id", s.toEchoHandler(inspect.GetSourceChunkDetailHandler))
	e.GET("/api/item/:id", s.toEchoHandler(inspect.GetSourceItemDetailHandler))

	// Deal Schedule
	e.POST("/api/send_deal", s.SendManualHandler)
	e.POST("/api/schedule", s.toEchoHandler(schedule.CreateHandler))
	e.GET("/api/schedule", s.toEchoHandler(schedule.ListHandler))
	e.POST("/api/schedule/:id/pause", s.toEchoHandler(schedule.PauseHandler))
	e.POST("/api/schedule/:id/resume", s.toEchoHandler(schedule.ResumeHandler))

	// Deal
	e.POST("/api/deal", s.toEchoHandler(deal.ListHandler))
}

var logger = logging.Logger("api")

type WebDavHandler struct {
	webdav.Handler
}

func (h *WebDavHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handler.ServeHTTP(w, r)
}

type cleanup func()

func (s Server) runWebdav(sourceHandler *datasource.RCloneHandler) (<-chan error, cleanup) {
	webdavHandler := &webdav.Handler{
		Prefix:     "/",
		FileSystem: sourceHandler,
		// TODO: use a real lock system with a shared backend
		LockSystem: webdav.NewMemLS(),
	}

	server := &http.Server{
		Addr:              s.webDavBind,
		Handler:           webdavHandler,
		ReadHeaderTimeout: time.Second * 10,
	}

	errCh := make(chan error)
	go func() {
		defer close(errCh)
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	return errCh, func() {
		logger.Info("shutting down webdav server")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			logger.Errorw("failed to shutdown webdav server", "err", err)
		}
	}
}

func (s Server) runAPI() (<-chan error, cleanup) {
	errChan := make(chan error)
	e := echo.New()
	e.Debug = true
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
		errChan <- err
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	/* deprecated API
	e.GET("/api/datasets", s.GetDatasets)
	e.GET("/api/dataset/:id/sources", s.GetSources)
	e.GET("/api/source/:id/cars", s.GetCars)
	e.GET("/api/car/:id/items", s.GetItems)
	e.GET("/api/car/:id/deals", s.GetDealsForCar)
	e.GET("/api/item/:id/deals", s.GetDealsForItem)
	e.GET("/api/directory/:id/entries", s.GetDirectoryEntries)
	*/
	e.GET("/*", echo.WrapHandler(http.FileServer(http.FS(efs))))

	go func() {
		defer close(errChan)
		err := e.Start(s.adminBind)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()

	return errChan, func() {
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()
		//nolint:contextcheck
		if err := e.Shutdown(shutdownCtx); err != nil {
			fmt.Printf("Error shutting down the server: %v\n", err)
		}
	}
}

func (s Server) Run(ctx context.Context) error {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGTRAP)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var (
		webdavErrChan <-chan error
		webdavCleanup func()
		apiErrChan    <-chan error
		apiCleanup    func()
	)
	if s.adminEnabled {
		//nolint:contextcheck
		apiErrChan, apiCleanup = s.runAPI()
		defer apiCleanup()
	}
	if s.webDavEnabled {
		var source model.Source
		err := s.db.WithContext(ctx).Where("id = ?", s.webDavSourceID).First(&source).Error
		if err != nil {
			return errors.Wrap(err, "failed to get source")
		}
		rCloneHandler, err := datasource.NewRCloneHandler(ctx, source)
		if err != nil {
			return errors.Wrap(err, "failed to create rclone handler")
		}
		if !rCloneHandler.Features().CanHaveEmptyDirectories {
			return errors.New("source does not support empty directories")
		}
		_, ok := rCloneHandler.Fs.(fs2.Mover)
		if !ok {
			return errors.New("source does not support moving files")
		}
		_, ok = rCloneHandler.Fs.(fs2.DirMover)
		if !ok {
			return errors.New("source does not support moving directories")
		}
		_, ok = rCloneHandler.Fs.(fs2.Purger)
		if !ok {
			return errors.New("source does not support purging")
		}
		_, ok = rCloneHandler.Fs.(fs2.PutStreamer)
		if !ok {
			return errors.New("source does not support streaming uploads")
		}
		//nolint:contextcheck
		webdavErrChan, webdavCleanup = s.runWebdav(rCloneHandler)
		defer webdavCleanup()
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-webdavErrChan:
		return errors.Wrap(err, "webdav server error")
	case err := <-apiErrChan:
		return errors.Wrap(err, "api server error")
	case sig := <-signalChan:
		logger.Infow("received signal", "signal", sig)
		cancel()
		return cli.Exit("shutting down", 130)
	}
}

/* deprecated API
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
*/
