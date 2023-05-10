package deal

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/data-preservation-programs/go-singularity/replication"
	"github.com/data-preservation-programs/go-singularity/util"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type Proposal struct {
	HTTPHeaders    []string `json:"httpHeaders"`
	URLTemplate    string   `json:"urlTemplate"`
	Price          float64  `json:"price"`
	Label          string   `json:"label"`
	Verified       bool     `json:"verified"`
	IPNI           bool     `json:"ipni"`
	KeepUnsealed   bool     `json:"keepUnsealed"`
	StartDelayDays float64  `json:"startDelayDays"`
	DurationDays   float64  `json:"durationDays"`
	ClientAddress  string   `json:"clientAddress"`
	ProviderID     string   `json:"providerID"`
	PieceCID       string   `json:"pieceCID"`
	PieceSize      string   `json:"pieceSize"`
}

func SendManualHandler(
	db *gorm.DB,
	lotusAPI string,
	lotusToken string,
	request Proposal,
) (string, *handler.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// Get the wallet object
	wallet := model.Wallet{}
	err := db.Where("id = ?", request.ClientAddress).First(&wallet).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", handler.NewBadRequestString("client address not found")
	}
	if err != nil {
		return "", handler.NewHandlerError(err)
	}

	host, err := util.InitHost(ctx, nil)
	if err != nil {
		return "", handler.NewHandlerError(err)
	}
	dealMaker, err := replication.NewDealMaker(lotusAPI, lotusToken, host)
	if err != nil {
		return "", handler.NewHandlerError(err)
	}
	providerInfo, err := dealMaker.GetProviderInfo(ctx, request.ProviderID)
	if err != nil {
		return "", handler.NewHandlerError(err)
	}
	car := model.Car{
		PieceCID:  "",
		PieceSize: 0,
		RootCID:   "",
		FileSize:  0,
		FilePath:  "",
	}
	schedule := model.Schedule{
		URLTemplate:          request.URLTemplate,
		HTTPHeaders:          request.HTTPHeaders,
		Provider:             request.ProviderID,
		Price:                request.Price,
		Verified:             request.Verified,
		KeepUnsealed:         request.KeepUnsealed,
		AnnounceToIPNI:       request.IPNI,
		StartDelay:           time.Duration(request.StartDelayDays * 24 * float64(time.Hour) / float64(time.Nanosecond)),
		Duration:             time.Duration(request.DurationDays * 24 * float64(time.Hour) / float64(time.Nanosecond)),
	}
	addrInfo := peer.AddrInfo{
		ID:    providerInfo.PeerID,
		Addrs: providerInfo.Multiaddrs,
	}
	proposalID, err := dealMaker.MakeDeal(ctx, time.Now(), wallet, car, schedule, addrInfo)
	if err != nil {
		return "", handler.NewHandlerError(err)
	}
	return proposalID, nil
}
