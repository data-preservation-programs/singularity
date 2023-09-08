package deal

import (
	"context"
	"strconv"
	"time"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/dustin/go-humanize"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/ipfs/go-cid"
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
		return 0, errors.WithStack(err)
	}
	return time.Duration(epochs) * 30 * time.Second, nil
}

// SendManualHandler creates a deal proposal manually based on the information provided in the Proposal.
//
// The function searches for the client's wallet using the provided address, validates various input fields such as the
// pieceCID, rootCID, piece size, etc., and then uses the dealMaker to create a deal. The result is a model.Deal that
// represents the proposal. Any issues during these operations result in an appropriate error response.
//
// Parameters:
// - ctx:       The context for the operation which can be used for timeouts and cancellations.
// - db:        The database connection for accessing and storing related data.
// - dealMaker: An interface responsible for creating deals based on the given configuration.
// - request:   The request object containing all the necessary information for creating a deal proposal.
//
// Returns:
// - A pointer to a model.Deal object representing the created deal.
// - An error indicating any issues that occurred during the process.
func (DefaultHandler) SendManualHandler(
	ctx context.Context,
	db *gorm.DB,
	dealMaker replication.DealMaker,
	request Proposal,
) (*model.Deal, error) {
	db = db.WithContext(ctx)
	// Get the wallet object
	wallet := model.Wallet{}
	err := db.Where("id = ? OR address = ?", request.ClientAddress, request.ClientAddress).First(&wallet).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "client address %s not found", request.ClientAddress)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pieceCID, err := cid.Parse(request.PieceCID)
	if err != nil {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid piece CID: %s", request.PieceCID)
	}
	if pieceCID.Type() != cid.FilCommitmentUnsealed {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "piece CID %s must be commp", request.PieceCID)
	}
	pieceSize, err := humanize.ParseBytes(request.PieceSize)
	if err != nil {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid piece size: %s", request.PieceSize)
	}
	if (pieceSize & (pieceSize - 1)) != 0 {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "piece size %d must be a power of 2", pieceSize)
	}
	rootCID, err := cid.Parse(request.RootCID)
	if err != nil {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid root CID: %s", request.RootCID)
	}
	car := model.Car{
		PieceCID:  model.CID(pieceCID),
		PieceSize: int64(pieceSize),
		RootCID:   model.CID(rootCID),
		FileSize:  int64(request.FileSize),
	}
	duration, err := argToDuration(request.Duration)
	if err != nil {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid duration: %s", request.Duration)
	}
	startDelay, err := argToDuration(request.StartDelay)
	if err != nil {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid start delay: %s", request.StartDelay)
	}

	dealConfig := replication.DealConfig{
		URLTemplate:    request.URLTemplate,
		HTTPHeaders:    request.HTTPHeaders,
		Provider:       request.ProviderID,
		Verified:       request.Verified,
		KeepUnsealed:   request.KeepUnsealed,
		AnnounceToIPNI: request.IPNI,
		StartDelay:     startDelay,
		Duration:       duration,
	}

	dealModel, err := dealMaker.MakeDeal(ctx, wallet, car, dealConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return dealModel, nil
}

// @ID SendManual
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
func _() {}
