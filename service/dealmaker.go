package service

import (
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

type DealMaker struct {
	db *gorm.DB
	logger *log.ZapEventLogger
}

func NewDealMaker(db *gorm.DB) DealMaker {
	return DealMaker{db: db, logger: log.Logger("deal-maker")}
}

func (d DealMaker) Run() {
	d.logger.Info("Starting deal maker...")
}
