package api

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/data-preservation-programs/singularity/analytics"
	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/singularity/handler/file"
	"github.com/data-preservation-programs/singularity/handler/job"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/retriever"
	"github.com/data-preservation-programs/singularity/retriever/endpointfinder"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/service/contentprovider"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/filecoin-project/lassie/pkg/lassie"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	_ "github.com/data-preservation-programs/singularity/docs/swagger"
	logging "github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

type Server struct {
	db              *gorm.DB
	listener        net.Listener
	lotusClient     jsonrpc.RPCClient
	dealMaker       replication.DealMaker
	closer          io.Closer
	host            host.Host
	retriever       *retriever.Retriever
	adminHandler    admin.Handler
	storageHandler  storage.Handler
	dataprepHandler dataprep.Handler
	dealHandler     deal.Handler
	walletHandler   wallet.Handler
	fileHandler     file.Handler
	jobHandler      job.Handler
	scheduleHandler schedule.Handler
}

func (s Server) Name() string {
	return "api"
}

// @Summary Get metadata for a piece
// @Description Get metadata for a piece for how it may be reassembled from the data source
// @Tags Piece
// @Produce json
// @Param id path string true "Piece CID"
// @Success 200 {object} store.PieceReader
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /piece/{id}/metadata [get]
func (s Server) getMetadataHandler(c echo.Context) error {
	return contentprovider.GetMetadataHandler(c, s.db)
}

func Run(c *cli.Context) error {
	connString := c.String("database-connection-string")
	bind := c.String("bind")
	lotusAPI := c.String("lotus-api")
	lotusToken := c.String("lotus-token")

	listener, err := net.Listen("tcp", bind)
	if err != nil {
		return errors.WithStack(err)
	}
	server, err := InitServer(c.Context, APIParams{
		ConnString: connString,
		Listener:   listener,
		LotusAPI:   lotusAPI,
		LotusToken: lotusToken,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	logger.Info("Starting Singularity API server...")
	return service.StartServers(c.Context, logger, server)
}

type APIParams struct {
	Listener   net.Listener
	LotusAPI   string
	LotusToken string
	ConnString string
}

func InitServer(ctx context.Context, params APIParams) (Server, error) {
	db, closer, err := database.OpenWithLogger(params.ConnString)
	if err != nil {
		return Server{}, errors.WithStack(err)
	}
	h, err := util.InitHost(nil)
	if err != nil {
		return Server{}, errors.Wrap(err, "failed to init host")
	}
	lassie, err := lassie.NewLassie(ctx, lassie.WithHost(h))
	if err != nil {
		return Server{}, errors.Wrap(err, "failed to init lassie")
	}
	infoFetcher := replication.MinerInfoFetcher{
		Client: util.NewLotusClient(params.LotusAPI, params.LotusToken),
	}
	endpointFinder := endpointfinder.NewEndpointFinder(
		infoFetcher,
		h,
		endpointfinder.WithLruSize(128),
		endpointfinder.WithLruTimeout(time.Hour*2),
		endpointfinder.WithErrorLruSize(128),
		endpointfinder.WithErrorLruTimeout(time.Minute*5),
	)
	return Server{
		db:          db,
		host:        h,
		listener:    params.Listener,
		lotusClient: util.NewLotusClient(params.LotusAPI, params.LotusToken),
		dealMaker: replication.NewDealMaker(
			util.NewLotusClient(params.LotusAPI, params.LotusToken),
			h,
			time.Hour,
			time.Minute*5,
		),
		retriever:       retriever.NewRetriever(lassie, endpointFinder),
		closer:          closer,
		adminHandler:    &admin.DefaultHandler{},
		storageHandler:  &storage.DefaultHandler{},
		dataprepHandler: &dataprep.DefaultHandler{},
		dealHandler:     &deal.DefaultHandler{},
		walletHandler:   &wallet.DefaultHandler{},
		fileHandler:     &file.DefaultHandler{},
		jobHandler:      &job.DefaultHandler{},
		scheduleHandler: &schedule.DefaultHandler{},
	}, nil
}

func (s Server) toEchoHandler(handlerFunc any) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("💥 Panic in handler:", r)
				_ = httpResponseFromError(c, fmt.Errorf("panic: %v", r))
			}
		}()

		handlerFuncValue := reflect.ValueOf(handlerFunc)
		handlerFuncType := handlerFuncValue.Type()

		if handlerFuncType.NumIn() < 2 ||
			handlerFuncType.In(1).String() != "*gorm.DB" ||
			handlerFuncType.In(0).String() != "context.Context" {
			logger.Error("Invalid handler function signature.")
			return echo.NewHTTPError(http.StatusInternalServerError, "Invalid handler function signature.")
		}

		var inputParams []reflect.Value
		var j int

		for i := 0; i < handlerFuncType.NumIn(); i++ {
			paramType := handlerFuncType.In(i)

			switch paramType.String() {
			case "context.Context":
				inputParams = append(inputParams, reflect.ValueOf(c.Request().Context()))
			case "*gorm.DB":
				inputParams = append(inputParams, reflect.ValueOf(s.db.WithContext(c.Request().Context())))
			case "jsonrpc.RPCClient":
				inputParams = append(inputParams, reflect.ValueOf(s.lotusClient))
			case "replication.DealMaker":
				inputParams = append(inputParams, reflect.ValueOf(s.dealMaker))
			default:
				if paramType.Kind() == reflect.String || isIntKind(paramType.Kind()) || isUIntKind(paramType.Kind()) {
					if j >= len(c.ParamValues()) {
						return c.JSON(http.StatusInternalServerError, HTTPError{Err: "invalid handler function signature"})
					}
					paramValue := c.ParamValues()[j]
					j++

					switch paramType.Kind() {
					case reflect.String:
						decoded, err := url.QueryUnescape(paramValue)
						if err != nil {
							return c.JSON(http.StatusInternalServerError, HTTPError{Err: "failed to decode path parameter"})
						}
						inputParams = append(inputParams, reflect.ValueOf(decoded))
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						val, err := strconv.ParseInt(paramValue, 10, paramType.Bits())
						if err != nil {
							return c.JSON(http.StatusBadRequest, HTTPError{Err: "failed to parse int"})
						}
						v := reflect.New(paramType).Elem()
						v.SetInt(val)
						inputParams = append(inputParams, v)
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						val, err := strconv.ParseUint(paramValue, 10, paramType.Bits())
						if err != nil {
							return c.JSON(http.StatusBadRequest, HTTPError{Err: "failed to parse uint"})
						}
						v := reflect.New(paramType).Elem()
						v.SetUint(val)
						inputParams = append(inputParams, v)
					}
				} else {
					bodyParam := reflect.New(paramType).Elem()
					if bodyParam.Kind() == reflect.Map {
						bodyParam.Set(reflect.MakeMap(bodyParam.Type()))
					}
					if err := c.Bind(bodyParam.Addr().Interface()); err != nil {
						return c.JSON(http.StatusBadRequest, HTTPError{Err: fmt.Sprintf("failed to bind request body: %s", err)})
					}
					inputParams = append(inputParams, bodyParam)
				}
			}
		}

		results := handlerFuncValue.Call(inputParams)

		switch len(results) {
		case 1:
			if errVal := results[0]; !errVal.IsNil() {
				err, _ := errVal.Interface().(error)
				logger.Warnf("🪝 toEchoHandler caught error (1-result): %+v", err)
				return httpResponseFromError(c, err)
			}
			return c.NoContent(http.StatusNoContent)

		case 2:
			if errVal := results[1]; !errVal.IsNil() {
				err, _ := errVal.Interface().(error)
				logger.Warnf("🪝 toEchoHandler caught error (2-result): %+v", err)
				return httpResponseFromError(c, err)
			}
			return c.JSON(http.StatusOK, results[0].Interface())

		default:
			return c.JSON(http.StatusInternalServerError, HTTPError{Err: "invalid handler return signature"})
		}
	}
}

func (s Server) setupRoutes(e *echo.Echo) {
	// Add debug middleware first
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger.Infow("Incoming request",
				"method", c.Request().Method,
				"path", c.Request().URL.Path,
			)
			return next(c)
		}
	})

	// Admin
	e.POST("/api/identity", s.toEchoHandler(s.adminHandler.SetIdentityHandler))

	// Storage - Modified DELETE route comes first
	e.DELETE("/api/storage/*", func(c echo.Context) error {
		path := strings.TrimPrefix(c.Request().URL.Path, "/api/storage/")
		logger.Infow("DELETE storage request", "path", path)
		return s.storageHandler.RemoveHandler(
			c.Request().Context(),
			s.db.WithContext(c.Request().Context()),
			path,
		)
	})

	// Other storage routes
	e.POST("/api/storage/:type", s.toEchoHandler(s.storageHandler.CreateStorageHandler))
	e.POST("/api/storage/:type/:provider", s.toEchoHandler(func(
		ctx context.Context,
		db *gorm.DB,
		storageType string,
		provider string,
		request storage.CreateRequest,
	) (*model.Storage, error) {
		request.Provider = provider
		return s.storageHandler.CreateStorageHandler(ctx, db, storageType, request)
	}))
	e.GET("/api/storage/:name/explore/:path", s.toEchoHandler(s.storageHandler.ExploreHandler))
	e.GET("/api/storage", s.toEchoHandler(s.storageHandler.ListStoragesHandler))
	e.PATCH("/api/storage/:name", s.toEchoHandler(s.storageHandler.UpdateStorageHandler))
	e.PATCH("/api/storage/:name/rename", func(c echo.Context) error {
		ctx := c.Request().Context()
		db := s.db.WithContext(ctx)
		name := c.Param("name")

		var req storage.RenameRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, HTTPError{
				Code: "INVALID_REQUEST",
				Err:  "Invalid request body",
			})
		}

		resp, err := s.storageHandler.RenameStorageHandler(ctx, db, name, req)
		if err != nil {
			// Ensure error is properly wrapped with stack trace
			return httpResponseFromError(c, errors.WithStack(err))
		}

		return c.JSON(http.StatusOK, resp)
	})
	// Preparation
	e.POST("/api/preparation", s.toEchoHandler(s.dataprepHandler.CreatePreparationHandler))
	e.DELETE("/api/preparation/:id", s.toEchoHandler(s.dataprepHandler.RemovePreparationHandler))
	e.GET("/api/preparation", s.toEchoHandler(s.dataprepHandler.ListHandler))
	e.GET("/api/preparation/:id", s.toEchoHandler(s.jobHandler.GetStatusHandler))
	e.GET("/api/preparation/:id/schedules", s.toEchoHandler(s.dataprepHandler.ListSchedulesHandler))
	e.PATCH("/api/preparation/:name/rename", s.toEchoHandler(s.dataprepHandler.RenamePreparationHandler))

	// Job management
	e.POST("/api/preparation/:id/source/:name/start-daggen", s.toEchoHandler(s.jobHandler.StartDagGenHandler))
	e.POST("/api/preparation/:id/source/:name/pause-daggen", s.toEchoHandler(s.jobHandler.PauseDagGenHandler))
	e.POST("/api/preparation/:id/source/:name/start-scan", s.toEchoHandler(s.jobHandler.StartScanHandler))
	e.POST("/api/preparation/:id/source/:name/pause-scan", s.toEchoHandler(s.jobHandler.PauseScanHandler))
	e.POST("/api/preparation/:id/source/:name/start-pack/:job_id", s.toEchoHandler(s.jobHandler.StartPackHandler))
	e.POST("/api/preparation/:id/source/:name/pause-pack/:job_id", s.toEchoHandler(s.jobHandler.PausePackHandler))
	e.POST("/api/preparation/:id/source/:name/finalize", s.toEchoHandler(s.jobHandler.PrepareToPackSourceHandler))
	e.POST("/api/job/:id/pack", s.toEchoHandler(s.jobHandler.PackHandler))

	// storage attachment
	e.POST("/api/preparation/:id/output/:name", s.toEchoHandler(s.dataprepHandler.AddOutputStorageHandler))
	e.POST("/api/preparation/:id/source/:name", s.toEchoHandler(s.dataprepHandler.AddSourceStorageHandler))
	e.DELETE("/api/preparation/:id/output/:name", s.toEchoHandler(s.dataprepHandler.RemoveOutputStorageHandler))

	// Explore
	e.GET("/api/preparation/:id/source/:name/explore/:path", s.toEchoHandler(s.dataprepHandler.ExploreHandler))

	// Piece
	e.GET("/api/preparation/:id/piece", s.toEchoHandler(s.dataprepHandler.ListPiecesHandler))
	e.POST("/api/preparation/:id/piece", s.toEchoHandler(s.dataprepHandler.AddPieceHandler))

	// Wallet
	e.POST("/api/wallet", s.toEchoHandler(s.walletHandler.ImportHandler))
	e.GET("/api/wallet", s.toEchoHandler(s.walletHandler.ListHandler))
	e.DELETE("/api/wallet/:address", s.toEchoHandler(s.walletHandler.RemoveHandler))

	// Wallet Association
	e.POST("/api/preparation/:id/wallet/:wallet", s.toEchoHandler(s.walletHandler.AttachHandler))
	e.GET("/api/preparation/:id/wallet", s.toEchoHandler(s.walletHandler.ListAttachedHandler))
	e.DELETE("/api/preparation/:id/wallet/:wallet", s.toEchoHandler(s.walletHandler.DetachHandler))

	// Piece metadata
	e.GET("/api/piece/:id/metadata", s.getMetadataHandler)

	// Deal Schedule
	e.POST("/api/send_deal", s.toEchoHandler(s.dealHandler.SendManualHandler))
	e.POST("/api/schedule", s.toEchoHandler(s.scheduleHandler.CreateHandler))
	e.GET("/api/schedule", s.toEchoHandler(s.scheduleHandler.ListHandler))
	e.POST("/api/schedule/:id/pause", s.toEchoHandler(s.scheduleHandler.PauseHandler))
	e.POST("/api/schedule/:id/resume", s.toEchoHandler(s.scheduleHandler.ResumeHandler))
	e.PATCH("/api/schedule/:id", s.toEchoHandler(s.scheduleHandler.UpdateHandler))
	e.DELETE("/api/schedule/:id", s.toEchoHandler(s.scheduleHandler.RemoveHandler))

	// Deal
	e.POST("/api/deal", s.toEchoHandler(s.dealHandler.ListHandler))

	// File
	e.GET("/api/file/:id/deals", s.toEchoHandler(s.fileHandler.GetFileDealsHandler))
	e.GET("/api/file/:id", s.toEchoHandler(s.fileHandler.GetFileHandler))
	e.POST("/api/file/:id/prepare_to_pack", s.toEchoHandler(s.fileHandler.PrepareToPackFileHandler))
	e.GET("/api/file/:id/retrieve", s.retrieveFile)
	e.POST("/api/preparation/:id/source/:name/file", s.toEchoHandler(s.fileHandler.PushFileHandler))
}

var logger = logging.Logger("api")

func (s Server) Start(ctx context.Context, exitErr chan<- error) error {
	err := analytics.Init(ctx, s.db)
	if err != nil {
		return errors.WithStack(err)
	}
	e := echo.New()
	e.Debug = true
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		logger.Warnf("🔥 Global error handler triggered: %+v", err)
		if herr := httpResponseFromError(c, err); herr != nil {
			logger.Error("❌ Failed to write error response:", herr)
			_ = c.JSON(http.StatusInternalServerError, HTTPError{
				Code: "INTERNAL_ERROR",
				Err:  herr.Error(),
			})
		}
	}

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		Skipper:           middleware.DefaultSkipper,
		StackSize:         4 << 10, // 4 KiB
		DisableStackAll:   false,
		DisablePrintStack: false,
		LogLevel:          0,
		//LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
		//	logger.Errorw("panic", "err", err, "stack", string(stack))
		//	return nil
		//},
	}))
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

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.Listener = s.listener

	done := make(chan struct{})
	eventsFlushed := make(chan struct{})

	go func() {
		err := e.Start("")
		<-eventsFlushed
		<-done

		if exitErr != nil {
			exitErr <- err
		}
	}()

	go func() {
		defer close(done)
		<-ctx.Done()
		ctx2, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		err := e.Shutdown(ctx2)
		if err != nil {
			logger.Errorw("failed to shutdown api server", "err", err)
		}
		err = s.closer.Close()
		if err != nil {
			logger.Errorw("failed to close database connection", "err", err)
		}
		s.host.Close()
	}()

	go func() {
		defer close(eventsFlushed)
		analytics.Default.Start(ctx)
		analytics.Default.Flush()
	}()

	return nil
}

func isIntKind(kind reflect.Kind) bool {
	return kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64
}

func isUIntKind(kind reflect.Kind) bool {
	return kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64
}

type HTTPError struct {
	Code string `json:"code,omitempty"`
	Err  string `json:"err"`
}

func httpResponseFromError(c echo.Context, e error) error {
	// Force unwrap all error layers
	unwrapped := errors.UnwrapAll(e)

	// Check the raw error message
	if strings.Contains(unwrapped.Error(), "duplicated key not allowed") {
		return c.JSON(http.StatusConflict, HTTPError{
			Code: "DUPLICATE_NAME",
			Err:  "Cannot rename: a storage with that name already exists. Please choose a different name.",
		})
	}

	// Default error response
	return c.JSON(http.StatusInternalServerError, HTTPError{
		Code: "INTERNAL_ERROR",
		Err:  e.Error(),
	})
}
