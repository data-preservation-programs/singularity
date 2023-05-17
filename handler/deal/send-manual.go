package deal

import (
	"context"
	"strconv"
	"time"

	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/data-preservation-programs/go-singularity/replication"
	"github.com/data-preservation-programs/go-singularity/util"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Proposal struct {
	HTTPHeaders    []string `json:"httpHeaders"`
	URLTemplate    string   `json:"urlTemplate"`
	Price          float64  `json:"price"`
	RootCID        string   `json:"rootCID"`
	Verified       bool     `json:"verified"`
	IPNI           bool     `json:"ipni"`
	KeepUnsealed   bool     `json:"keepUnsealed"`
	StartDelayDays float64  `json:"startDelayDays"`
	DurationDays   float64  `json:"durationDays"`
	ClientAddress  string   `json:"clientAddress"`
	ProviderID     string   `json:"providerID"`
	PieceCID       string   `json:"pieceCID"`
	PieceSize      string   `json:"pieceSize"`
	FileSize       uint64   `json:"fileSize"`
	LotusAPI       string   `swaggerignore:"true"`
	LotusToken     string   `swaggerignore:"true"`
}

// SendManualHandler godoc
// @Summary Send a manual deal proposal
// @Description Send a manual deal proposal
// @Tags Deal
// @Accept json
// @Produce json
// @Param proposal body Proposal true "Proposal"
// @Success 200 {string} string
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /deal/send_manual [post]
func SendManualHandler(
	db *gorm.DB,
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
	dealMaker, err := replication.NewDealMaker(request.LotusAPI, request.LotusToken, host)
	if err != nil {
		return "", handler.NewHandlerError(err)
	}
	providerInfo, err := dealMaker.GetProviderInfo(ctx, request.ProviderID)
	if err != nil {
		return "", handler.NewHandlerError(err)
	}
	pieceCID, err := cid.Parse(request.PieceCID)
	if err != nil {
		return "", handler.NewBadRequestString("invalid piece CID")
	}
	if pieceCID.Type() != cid.FilCommitmentUnsealed {
		return "", handler.NewBadRequestString("piece CID must be commp")
	}
	pieceSize, err := strconv.ParseInt(request.PieceSize, 10, 64)
	if err != nil {
		return "", handler.NewBadRequestString("invalid piece size")
	}
	if (pieceSize & (pieceSize - 1)) != 0 {
		return "", handler.NewBadRequestString("piece size must be a power of 2")
	}
	_, err = cid.Parse(request.RootCID)
	if err != nil {
		return "", handler.NewBadRequestString("invalid root CID")
	}
	car := model.Car{
		PieceCID:  request.PieceCID,
		PieceSize: uint64(pieceSize),
		RootCID:   request.RootCID,
		FileSize:  request.FileSize,
	}
	schedule := model.Schedule{
		URLTemplate:    request.URLTemplate,
		HTTPHeaders:    request.HTTPHeaders,
		Provider:       request.ProviderID,
		Price:          request.Price,
		Verified:       request.Verified,
		KeepUnsealed:   request.KeepUnsealed,
		AnnounceToIPNI: request.IPNI,
		StartDelay:     time.Duration(request.StartDelayDays * 24 * float64(time.Hour) / float64(time.Nanosecond)),
		Duration:       time.Duration(request.DurationDays * 24 * float64(time.Hour) / float64(time.Nanosecond)),
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
