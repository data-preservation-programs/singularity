package service

import (
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pkg/errors"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
	"time"
)

type SpadeAPI struct {
	db            *gorm.DB
	logger        *log.ZapEventLogger
	dealMaker     *replication.DealMaker
	walletChooser *replication.WalletChooser
	bind          string
}

func NewSpadeAPIService(db *gorm.DB, dealMaker *replication.DealMaker, walletChooser *replication.WalletChooser, bind string) *SpadeAPI {
	return &SpadeAPI{
		db:            db,
		logger:        log.Logger("spade-api"),
		dealMaker:     dealMaker,
		walletChooser: walletChooser,
		bind:          bind,
	}
}

func (s SpadeAPI) Start() error {
	e := echo.New()
	current := log.GetConfig().Level
	if log.LevelInfo < current {
		log.SetAllLoggers(log.LevelInfo)
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
				s.logger.With("status", status, "latency_ms", latency.Milliseconds(), "err", err).Error(method + " " + uri)
			} else {
				s.logger.With("status", status, "latency_ms", latency.Milliseconds()).Info(method + " " + uri)
			}
			return nil
		},
	}))
	e.GET("/eligible_pieces", s.GetEligiblePieces)
	e.GET("/request_piece/:piece_cid", s.RequestPiece)
	return e.Start(s.bind)
}

func (s SpadeAPI) RequestPiece(c echo.Context) error {
	provider := c.QueryParam("provider")
	datasetName := c.QueryParam("dataset")
	pieceCID := c.Param("piece_cid")
	var dataset model.Dataset
	err := s.db.Preload("Wallets").Where("name = ?", datasetName).First(&dataset).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(404, "dataset not found")
	}
	if err != nil {
		return err
	}
	var car model.Car
	err = s.db.Where("dataset_id = ? AND piece_cid = ?", dataset.ID, pieceCID).First(&car).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(404, "no eligible pieces")
	}
	if err != nil {
		return err
	}
	wallet := s.walletChooser.Choose(c.Request().Context(), dataset.Wallets)

	providerInfo, err := s.dealMaker.GetProviderInfo(c.Request().Context(), provider)
	if err != nil {
		return err
	}
	schedule := model.Schedule{
		DatasetID:               dataset.ID,
		URLTemplate:             "",
		HTTPHeaders:             nil,
		Provider:                provider,
		Price:                   0,
		TotalDealNumber:         0,
		TotalDealSize:           0,
		Verified:                true,
		KeepUnsealed:            true,
		AnnounceToIPNI:          true,
		StartDelay:              72 * time.Hour,
		Duration:                530 * 24 * time.Hour,
		State:                   model.ScheduleActive,
	}
	addrInfo := peer.AddrInfo{
		ID:    providerInfo.PeerID,
		Addrs: providerInfo.Multiaddrs,
	}
	proposal, err := s.dealMaker.MakeDeal(c.Request().Context(), time.Now(), wallet, car, schedule, addrInfo)
	if err != nil {
		return err
	}
	return c.JSON(200, proposal)
}

func (s SpadeAPI) GetEligiblePieces(c echo.Context) error {
	provider := c.QueryParam("provider")
	datasetName := c.QueryParam("dataset")
	var dataset model.Dataset
	err := s.db.Where("name = ?", datasetName).First(&dataset).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(404, "dataset not found")
	}
	if err != nil {
		return err
	}
	var cars []model.Car
	err = s.db.Where("dataset_id = ? AND piece_cid NOT IN (?)", dataset.ID,
		s.db.Table("deals").Select("piece_cid").
			Where("provider = ? AND state IN (?)",
				provider,
				[]model.DealState{
					model.DealProposed, model.DealPublished, model.DealActive,
				})).Find(&cars).Error
	if err != nil {
		return err
	}
	pieceCIDs := underscore.Map(cars, func(car model.Car) model.CID { return car.PieceCID })
	return c.JSON(200, pieceCIDs)
}
