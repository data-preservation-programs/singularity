package deal

import (
	"context"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/ipfs/go-cid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

//nolint:lll
type Proposal struct {
	HTTPHeaders     []string `json:"httpHeaders"`                          // http headers to be passed with the request (i.e. key=value)
	URLTemplate     string   `json:"urlTemplate"`                          // URL template with PIECE_CID placeholder for boost to fetch the CAR file, i.e. http://127.0.0.1/piece/{PIECE_CID}.car
	PricePerGBEpoch float64  `default:"0"          json:"pricePerGbEpoch"` // Price in FIL per GiB per epoch
	PricePerGB      float64  `default:"0"          json:"pricePerGb"`      // Price in FIL  per GiB
	PricePerDeal    float64  `default:"0"          json:"pricePerDeal"`    // Price in FIL per deal
	RootCID         string   `default:"bafkqaaa"   json:"rootCid"`         // Root CID that is required as part of the deal proposal, if empty, will be set to empty CID
	Verified        bool     `default:"true"       json:"verified"`        // Whether the deal should be verified
	IPNI            bool     `default:"true"       json:"ipni"`            // Whether the deal should be IPNI
	KeepUnsealed    bool     `default:"true"       json:"keepUnsealed"`    // Whether the deal should be kept unsealed
	StartDelay      string   `default:"72h"        json:"startDelay"`      // Deal start delay in epoch or in duration format, i.e. 1000, 72h
	Duration        string   `default:"12740h"     json:"duration"`        // Duration in epoch or in duration format, i.e. 1500000, 2400h
	ClientAddress   string   `json:"clientAddress"`                        // Client address
	ProviderID      string   `json:"providerId"`                           // Provider ID
	PieceCID        string   `json:"pieceCid"`                             // Piece CID
	PieceSize       string   `json:"pieceSize"`                            // Piece size
	FileSize        uint64   `json:"fileSize"`                             // File size in bytes for boost to fetch the CAR file
}

func argToDuration(s string) (time.Duration, error) {
	duration, err := time.ParseDuration(s)
	if err == nil {
		return duration, nil
	}
	epochs, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return time.Duration(epochs) * 30 * time.Second, nil
}

func SendManualHandler(
	db *gorm.DB,
	ctx context.Context,
	dealMaker replication.DealMaker,
	request Proposal,
) (*model.Deal, error) {
	return sendManualHandler(db, ctx, dealMaker, request)
}

// @Summary Send a manual deal proposal
// @Description Send a manual deal proposal
// @Tags Deal
// @Accept json
// @Produce json
// @Param proposal body Proposal true "Proposal"
// @Success 200 {object} model.Deal
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /send_deal [post]
func sendManualHandler(
	db *gorm.DB,
	ctx context.Context,
	dealMaker replication.DealMaker,
	request Proposal,
) (*model.Deal, error) {
	// Get the wallet object
	wallet := model.Wallet{}
	err := db.Where("id = ?", request.ClientAddress).First(&wallet).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr("client address not found")
	}
	if err != nil {
		return nil, err
	}

	pieceCID, err := cid.Parse(request.PieceCID)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid piece CID")
	}
	if pieceCID.Type() != cid.FilCommitmentUnsealed {
		return nil, handler.NewInvalidParameterErr("piece CID must be commp")
	}
	pieceSize, err := humanize.ParseBytes(request.PieceSize)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid piece size")
	}
	if (pieceSize & (pieceSize - 1)) != 0 {
		return nil, handler.NewInvalidParameterErr("piece size must be a power of 2")
	}
	rootCID, err := cid.Parse(request.RootCID)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid root CID")
	}
	car := model.Car{
		PieceCID:  model.CID(pieceCID),
		PieceSize: int64(pieceSize),
		RootCID:   model.CID(rootCID),
		FileSize:  int64(request.FileSize),
	}
	duration, err := argToDuration(request.Duration)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid duration")
	}
	startDelay, err := argToDuration(request.StartDelay)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid start delay")
	}

	dealConfig := replication.DealConfig{
		URLTemplate:    request.URLTemplate,
		HTTPHeaders:    request.HTTPHeaders,
		Provider:       request.ProviderID,
		Verified:       request.Verified,
		KeepUnsealed:   request.KeepUnsealed,
		AnnounceToIPNI: request.IPNI,
		StartDelay:     duration,
		Duration:       startDelay,
	}

	dealModel, err := dealMaker.MakeDeal(ctx, wallet, car, dealConfig)
	if err != nil {
		return nil, err
	}
	return dealModel, nil
}
