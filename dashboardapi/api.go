package dashboardapi

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type DealStats struct {
	Provider string
	State    model.DealState
	Day      string
	DealSize int64
}

type DashboardAPI struct {
	db   *gorm.DB
	bind string
	port int
}

func Run(c *cli.Context) error {
	db := database.MustOpenFromCLI(c)
	port := c.Int("port")
	bind := c.String("bind")
	return DashboardAPI{db: db, bind: bind, port: port}.Run(c)
}

func (d DashboardAPI) Run(c *cli.Context) error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
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

func (d DashboardAPI) GetOverallDealStats(c echo.Context) error {
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

func (d DashboardAPI) GetDealStats(c echo.Context) error {
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

func (d DashboardAPI) GetDatasets(c echo.Context) error {
	var datasets []model.Dataset
	err := d.db.Find(&datasets).Error
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, datasets)
}

func (d DashboardAPI) GetSources(c echo.Context) error {
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

func (d DashboardAPI) GetCars(c echo.Context) error {
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

func (d DashboardAPI) GetItems(c echo.Context) error {
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

func (d DashboardAPI) GetDealsForCar(c echo.Context) error {
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

func (d DashboardAPI) GetDealsForItem(c echo.Context) error {
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

func (d DashboardAPI) GetDirectoryEntries(c echo.Context) error {
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
