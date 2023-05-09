package api

import (
	_ "github.com/data-preservation-programs/go-singularity/api/docs"
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/handler/dataset"
	"github.com/data-preservation-programs/go-singularity/model"
	logging "github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type DealStats struct {
	Provider string
	State    model.DealState
	Day      string
	DealSize int64
}

type Server struct {
	db   *gorm.DB
	bind string
	port int
}

func Run(c *cli.Context) error {
	db := database.MustOpenFromCLI(c)
	port := c.Int("port")
	bind := c.String("bind")
	return Server{db: db, bind: bind, port: port}.Run(c)
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
		inputParams := []reflect.Value{reflect.ValueOf(d.db)}

		// Get path parameters
		for i := 1; i < handlerFuncType.NumIn(); i++ {
			paramType := handlerFuncType.In(i)
			if paramType.Kind() == reflect.String {
				if len(c.ParamValues()) < i {
					logger.Error("Invalid handler function signature.")
					return echo.NewHTTPError(http.StatusInternalServerError, "Invalid handler function signature.")
				}
				paramValue := c.ParamValues()[i-1]
				inputParams = append(inputParams, reflect.ValueOf(paramValue))
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
	e.POST("/admin/api/init", d.toEchoHandler(handler.InitHandler))

	e.POST("/admin/api/dataset", d.toEchoHandler(dataset.CreateHandler))

	e.GET("/admin/api/datasets", d.toEchoHandler(dataset.ListHandler))

	e.DELETE("/admin/api/dataset/:name", func(c echo.Context) error {
		name := c.Param("name")
		err := dataset.RemoveHandler(d.db, name)
		if err != nil {
			return err.HttpResponse(c)
		}

		return c.String(http.StatusNoContent, "")
	})

	e.POST("/admin/api/dataset/:name/source", func(c echo.Context) error {
		name := c.Param("name")
		request := dataset.AddSourceRequest{}
		err := c.Bind(&request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"err": err.Error()})
		}
		request.DatasetName = name

		response, err2 := dataset.AddSourceHandler(d.db, request)
		if err != nil {
			return err2.HttpResponse(c)
		}

		return c.JSON(http.StatusOK, response)
	})

	e.GET("/admin/api/dataset/:name/sources", func(c echo.Context) error {
		name := c.Param("name")
		response, err := dataset.ListSourceHandler(d.db, name)
		if err != nil {
			return err.HttpResponse(c)
		}

		return c.JSON(http.StatusOK, response)
	})

	e.DELETE("/admin/api/dataset/:name/sourcepath/:source", func(c echo.Context) error {
		name := c.Param("name")
		source := c.Param("source")
		err := dataset.RemoveSourceHandler(d.db, name, source)
		if err != nil {
			return err.HttpResponse(c)
		}

		return c.String(http.StatusNoContent, "")
	})
}

var logger = logging.Logger("api")

func (d Server) Run(c *cli.Context) error {
	e := echo.New()
	current := logging.GetConfig().Level
	if logging.LevelInfo < current {
		logging.SetAllLoggers(logging.LevelInfo)
	}
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			uri := v.URI
			status := v.Status
			latency := time.Now().Sub(v.StartTime)
			err := v.Error
			method := v.Method
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
	return e.Start(d.bind + ":" + c.String("port"))
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
