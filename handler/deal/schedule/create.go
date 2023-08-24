package schedule

import (
	"context"
	"strconv"
	"time"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/robfig/cron/v3"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/dustin/go-humanize"
	"github.com/ipfs/go-cid"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

//nolint:lll
type CreateRequest struct {
	DatasetName          string   `json:"datasetName"          validation:"required"`  // Preparation name
	Provider             string   `json:"provider"             validation:"required"`  // Provider
	HTTPHeaders          []string `json:"httpHeaders"`                                 // http headers to be passed with the request (i.e. key=value)
	URLTemplate          string   `json:"urlTemplate"`                                 // URL template with PIECE_CID placeholder for boost to fetch the CAR file, i.e. http://127.0.0.1/piece/{PIECE_CID}.car
	PricePerGBEpoch      float64  `default:"0"                 json:"pricePerGbEpoch"` // Price in FIL per GiB per epoch
	PricePerGB           float64  `default:"0"                 json:"pricePerGb"`      // Price in FIL  per GiB
	PricePerDeal         float64  `default:"0"                 json:"pricePerDeal"`    // Price in FIL per deal
	Verified             bool     `default:"true"              json:"verified"`        // Whether the deal should be verified
	IPNI                 bool     `default:"true"              json:"ipni"`            // Whether the deal should be IPNI
	KeepUnsealed         bool     `default:"true"              json:"keepUnsealed"`    // Whether the deal should be kept unsealed
	StartDelay           string   `default:"72h"               json:"startDelay"`      // Deal start delay in epoch or in duration format, i.e. 1000, 72h
	Duration             string   `default:"12840h"            json:"duration"`        // Duration in epoch or in duration format, i.e. 1500000, 2400h
	ScheduleCron         string   `json:"scheduleCron"`                                // Schedule cron patter
	ScheduleDealNumber   int      `json:"scheduleDealNumber"`                          // Number of deals per scheduled time
	TotalDealNumber      int      `json:"totalDealNumber"`                             // Total number of deals
	ScheduleDealSize     string   `json:"scheduleDealSize"`                            // Size of deals per schedule trigger in human readable format, i.e. 100 TiB
	TotalDealSize        string   `json:"totalDealSize"`                               // Total size of deals in human readable format, i.e. 100 TiB
	Notes                string   `json:"notes"`                                       // Notes
	MaxPendingDealSize   string   `json:"maxPendingDealSize"`                          // Max pending deal size in human readable format, i.e. 100 TiB
	MaxPendingDealNumber int      `json:"maxPendingDealNumber"`                        // Max pending deal number
	//nolint:tagliatelle
	AllowedPieceCIDs []string `json:"allowedPieceCids"` // Allowed piece CIDs in this schedule
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

func CreateHandler(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	request CreateRequest,
) (*model.Schedule, error) {
	return createHandler(ctx, db.WithContext(ctx), lotusClient, request)
}

// @Summary Create a new schedule
// @Description Create a new schedule
// @Tags Deal Schedule
// @Accept json
// @Produce json
// @Param schedule body CreateRequest true "CreateRequest"
// @Success 200 {object} model.Schedule
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /schedule [post]
func createHandler(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	request CreateRequest,
) (*model.Schedule, error) {
	dataset, err := database.FindDatasetByName(db, request.DatasetName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handlererror.NewInvalidParameterErr("dataset not found")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	startDelay, err := argToDuration(request.StartDelay)
	if err != nil {
		return nil, handlererror.NewInvalidParameterErr("invalid start delay")
	}

	duration, err := argToDuration(request.Duration)
	if err != nil {
		return nil, handlererror.NewInvalidParameterErr("invalid duration")
	}

	var scheduleCron string
	if request.ScheduleCron != "" {
		cronParser := cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
		_, err = cronParser.Parse(request.ScheduleCron)
		if err != nil {
			return nil, handlererror.NewInvalidParameterErr("invalid schedule cron")
		} else {
			scheduleCron = request.ScheduleCron
		}
	}

	totalDealSize, err := humanize.ParseBytes(request.TotalDealSize)
	if err != nil {
		return nil, handlererror.NewInvalidParameterErr("invalid total deal size")
	}
	scheduleDealSize, err := humanize.ParseBytes(request.ScheduleDealSize)
	if err != nil {
		return nil, handlererror.NewInvalidParameterErr("invalid schedule deal size")
	}
	pendingDealSize, err := humanize.ParseBytes(request.MaxPendingDealSize)
	if err != nil {
		return nil, handlererror.NewInvalidParameterErr("invalid pending deal size")
	}
	if scheduleCron != "" && scheduleDealSize == 0 && request.ScheduleDealNumber == 0 {
		return nil, handlererror.NewInvalidParameterErr("schedule deal number or size must be set when using cron schedule")
	}
	for _, pieceCID := range request.AllowedPieceCIDs {
		parsed, err := cid.Parse(pieceCID)
		if err != nil {
			return nil, handlererror.NewInvalidParameterErr("invalid allowed piece CID, it's not a CID")
		}
		if parsed.Type() != cid.FilCommitmentUnsealed {
			return nil, handlererror.NewInvalidParameterErr("invalid allowed piece CID, it's not a commp")
		}
	}

	var walletCount int64
	err = db.Model(&model.WalletAssignment{}).Where("dataset_id = ?", dataset.ID).Count(&walletCount).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if walletCount == 0 {
		return nil, handlererror.NewInvalidParameterErr("no wallet assigned to this dataset")
	}

	var providerActor string
	err = lotusClient.CallFor(ctx, &providerActor, "Filecoin.StateLookupID", request.Provider, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	schedule := model.Schedule{
		DatasetID:            dataset.ID,
		URLTemplate:          request.URLTemplate,
		HTTPHeaders:          request.HTTPHeaders,
		Provider:             request.Provider,
		TotalDealNumber:      request.TotalDealNumber,
		TotalDealSize:        int64(totalDealSize),
		Verified:             request.Verified,
		KeepUnsealed:         request.KeepUnsealed,
		AnnounceToIPNI:       request.IPNI,
		StartDelay:           startDelay,
		Duration:             duration,
		State:                model.ScheduleActive,
		ScheduleDealNumber:   request.ScheduleDealNumber,
		ScheduleDealSize:     int64(scheduleDealSize),
		MaxPendingDealNumber: request.MaxPendingDealNumber,
		MaxPendingDealSize:   int64(pendingDealSize),
		Notes:                request.Notes,
		AllowedPieceCIDs:     request.AllowedPieceCIDs,
		ScheduleCron:         scheduleCron,
		PricePerGBEpoch:      request.PricePerGBEpoch,
		PricePerGB:           request.PricePerGB,
		PricePerDeal:         request.PricePerDeal,
	}

	if err := database.DoRetry(ctx, func() error {
		return db.Create(&schedule).Error
	}); err != nil {
		return nil, errors.WithStack(err)
	}
	return &schedule, nil
}
