package api

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/service/contentprovider"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/data-preservation-programs/singularity/dashboard"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	_ "github.com/data-preservation-programs/singularity/docs/swagger"
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
	listener                  net.Listener
	datasourceHandlerResolver datasource.HandlerResolver
	lotusClient               jsonrpc.RPCClient
	dealMaker                 replication.DealMaker
	closer                    io.Closer
}

// @Summary Get metadata for a piece
// @Description Get metadata for a piece for how it may be reassembled from the data source
// @Tags Metadata
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
		return err
	}
	server, err := InitServer(APIParams{
		ConnString: connString,
		Listener:   listener,
		LotusAPI:   lotusAPI,
		LotusToken: lotusToken,
	})
	if err != nil {
		return err
	}

	logger.Info("Starting Singularity API server...")
	return server.Run(c.Context)
}

type APIParams struct {
	Ctx        context.Context
	Listener   net.Listener
	LotusAPI   string
	LotusToken string
	ConnString string
}

func InitServer(params APIParams) (Server, error) {
	db, closer, err := database.OpenWithLogger(params.ConnString)
	if err != nil {
		return Server{}, err
	}
	if err := model.AutoMigrate(db); err != nil {
		return Server{}, err
	}
	h, err := util.InitHost(nil)
	if err != nil {
		return Server{}, errors.Wrap(err, "failed to init host")
	}

	return Server{db: db, listener: params.Listener,
		datasourceHandlerResolver: &datasource.DefaultHandlerResolver{},
		lotusClient:               util.NewLotusClient(params.LotusAPI, params.LotusToken),
		dealMaker: replication.NewDealMaker(
			util.NewLotusClient(params.LotusAPI, params.LotusToken),
			h,
			time.Hour,
			time.Minute*5,
		),
		closer: closer,
	}, nil
}

func (s Server) toEchoHandler(handlerFunc any) echo.HandlerFunc {
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
			if paramType.String() == "datasource.HandlerResolver" {
				inputParams = append(inputParams, reflect.ValueOf(s.datasourceHandlerResolver))
				continue
			}
			if paramType.String() == "jsonrpc.RPCClient" {
				inputParams = append(inputParams, reflect.ValueOf(s.lotusClient))
				continue
			}
			if paramType.String() == "replication.DealMaker" {
				inputParams = append(inputParams, reflect.ValueOf(s.dealMaker))
				continue
			}
			if paramType.Kind() == reflect.String || isIntKind(paramType.Kind()) || isUIntKind(paramType.Kind()) {
				if j >= len(c.ParamValues()) {
					logger.Error("Invalid handler function signature.")
					return c.JSON(http.StatusInternalServerError, HTTPError{Err: "invalid handler function signature"})
				}
				paramValue := c.ParamValues()[j]
				switch {
				case paramType.Kind() == reflect.String:
					decoded, err := url.QueryUnescape(paramValue)
					if err != nil {
						return c.JSON(http.StatusInternalServerError, HTTPError{Err: "failed to decode path parameter"})
					}
					inputParams = append(inputParams, reflect.ValueOf(decoded))
				case isIntKind(paramType.Kind()):
					decoded, err := strconv.ParseInt(paramValue, 10, paramType.Bits())
					if err != nil {
						return c.JSON(http.StatusBadRequest, HTTPError{Err: "failed to parse path parameter as number"})
					}
					val := reflect.New(paramType).Elem()
					val.SetInt(decoded)
					inputParams = append(inputParams, val)
				case isUIntKind(paramType.Kind()):
					decoded, err := strconv.ParseUint(paramValue, 10, paramType.Bits())
					if err != nil {
						return c.JSON(http.StatusBadRequest, HTTPError{Err: "failed to parse path parameter as number"})
					}
					val := reflect.New(paramType).Elem()
					val.SetUint(decoded)
					inputParams = append(inputParams, val)
				default:
				}
				j += 1
				continue
			}
			bodyParam := reflect.New(paramType).Elem()
			if bodyParam.Kind() == reflect.Map {
				bodyParam.Set(reflect.MakeMap(bodyParam.Type()))
			}
			if err := c.Bind(bodyParam.Addr().Interface()); err != nil {
				return c.JSON(http.StatusBadRequest, HTTPError{Err: fmt.Sprintf("failed to bind request body: %s", err)})
			}
			inputParams = append(inputParams, bodyParam)
			break
		}

		// Call the handler function
		results := handlerFuncValue.Call(inputParams)

		if len(results) == 1 {
			// Handle the returned error
			if results[0].Interface() != nil {
				err, ok := results[1].Interface().(error)
				if !ok {
					return c.JSON(http.StatusInternalServerError, HTTPError{Err: "invalid handler function signature"})
				}
				return httpResponseFromError(c, err)
			}
			return c.NoContent(http.StatusNoContent)
		}

		// Handle the returned error
		if results[1].Interface() != nil {
			err, ok := results[1].Interface().(error)
			if !ok {
				return c.JSON(http.StatusInternalServerError, HTTPError{Err: "invalid handler function signature"})
			}
			return httpResponseFromError(c, err)
		}

		// Handle the returned data
		data := results[0].Interface()
		return c.JSON(http.StatusOK, data)
	}
}

func (s Server) setupRoutes(e *echo.Echo) {
	// Admin
	e.POST("/api/admin/reset", s.toEchoHandler(admin.ResetHandler))
	e.POST("/api/admin/init", s.toEchoHandler(admin.InitHandler))

	// Dataset
	e.POST("/api/dataset", s.toEchoHandler(dataset.CreateHandler))
	e.GET("/api/dataset/:datasetName/sources", s.toEchoHandler(datasource2.ListSourcesByDatasetHandler))
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
	e.POST("/api/source/:type/dataset/:datasetName", s.toEchoHandler(datasource2.CreateDatasourceHandler))
	e.GET("/api/source", s.toEchoHandler(datasource2.ListSourceHandler))
	e.PATCH("/api/source/:id", s.toEchoHandler(datasource2.UpdateSourceHandler))
	e.DELETE("/api/source/:id", s.toEchoHandler(datasource2.RemoveSourceHandler))
	e.POST("/api/source/:id/rescan", s.toEchoHandler(datasource2.RescanSourceHandler))
	e.POST("/api/source/:id/daggen", s.toEchoHandler(datasource2.DagGenHandler))
	e.POST("/api/source/:id/push", s.toEchoHandler(datasource2.PushItemHandler))
	e.POST("/api/source/:id/repack", s.toEchoHandler(datasource2.RepackHandler))
	e.POST("/api/source/:id/chunk", s.toEchoHandler(datasource2.ChunkHandler))
	e.POST("/api/chunk/:id/pack", s.toEchoHandler(datasource2.PackHandler))

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
	e.POST("/api/send_deal", s.toEchoHandler(deal.SendManualHandler))
	e.POST("/api/schedule", s.toEchoHandler(schedule.CreateHandler))
	e.GET("/api/schedule", s.toEchoHandler(schedule.ListHandler))
	e.POST("/api/schedule/:id/pause", s.toEchoHandler(schedule.PauseHandler))
	e.POST("/api/schedule/:id/resume", s.toEchoHandler(schedule.ResumeHandler))

	// Deal
	e.POST("/api/deal", s.toEchoHandler(deal.ListHandler))
}

var logger = logging.Logger("api")

func (s Server) Run(ctx context.Context) error {
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

	s.setupRoutes(e) //nolint: contextcheck
	efs, err := fs.Sub(dashboard.DashboardStaticFiles, "build")
	if err != nil {
		return err
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/*", echo.WrapHandler(http.FileServer(http.FS(efs))))
	e.Listener = s.listener
	shutdownDone := make(chan struct{})
	defer func() {
		<-shutdownDone
		logger.Info("Server shutdown complete")
		logger.Debug("Closing the database connection")
		s.closer.Close()
	}()
	go func() {
		<-ctx.Done()
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()
		logger.Info("Gracefully shutting down the server...")
		//nolint:contextcheck
		if err := e.Shutdown(shutdownCtx); err != nil {
			fmt.Printf("Error shutting down the server: %v\n", err)
		}
		close(shutdownDone)
	}()
	return e.Start("")
}

func isIntKind(kind reflect.Kind) bool {
	return kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64
}

func isUIntKind(kind reflect.Kind) bool {
	return kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64
}

type HTTPError struct {
	Err string `json:"err"`
}

func httpResponseFromError(c echo.Context, e error) error {
	if e == nil {
		return c.String(http.StatusOK, "OK")
	}

	httpStatusCode := http.StatusInternalServerError

	var invalidParameterErr handler.InvalidParameterError
	if errors.As(e, &invalidParameterErr) {
		httpStatusCode = http.StatusBadRequest
		e = invalidParameterErr.Unwrap()
	}

	var notFoundErr handler.NotFoundError
	if errors.As(e, &notFoundErr) {
		httpStatusCode = http.StatusNotFound
		e = notFoundErr.Unwrap()
	}

	var duplicateRecordErr handler.DuplicateRecordError
	if errors.As(e, &duplicateRecordErr) {
		httpStatusCode = http.StatusConflict
		e = duplicateRecordErr.Unwrap()
	}

	return c.JSON(httpStatusCode, HTTPError{Err: e.Error()})
}
