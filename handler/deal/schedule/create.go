package schedule

import (
	"context"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/rjNemo/underscore"
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
	Preparation           string   `json:"preparation"           validation:"required"`  // Preparation ID or name
	Provider              string   `json:"provider"              validation:"required"`  // Provider
	HTTPHeaders           []string `json:"httpHeaders"`                                  // http headers to be passed with the request (i.e. key=value)
	URLTemplate           string   `json:"urlTemplate"`                                  // URL template with PIECE_CID placeholder for boost to fetch the CAR file, i.e. http://127.0.0.1/piece/{PIECE_CID}.car
	PricePerGBEpoch       float64  `default:"0"                  json:"pricePerGbEpoch"` // Price in FIL per GiB per epoch
	PricePerGB            float64  `default:"0"                  json:"pricePerGb"`      // Price in FIL  per GiB
	PricePerDeal          float64  `default:"0"                  json:"pricePerDeal"`    // Price in FIL per deal
	Verified              bool     `default:"true"               json:"verified"`        // Whether the deal should be verified
	IPNI                  bool     `default:"true"               json:"ipni"`            // Whether the deal should be IPNI
	KeepUnsealed          bool     `default:"true"               json:"keepUnsealed"`    // Whether the deal should be kept unsealed
	StartDelay            string   `default:"72h"                json:"startDelay"`      // Deal start delay in epoch or in duration format, i.e. 1000, 72h
	Duration              string   `default:"12840h"             json:"duration"`        // Duration in epoch or in duration format, i.e. 1500000, 2400h
	ScheduleCron          string   `json:"scheduleCron"`                                 // Schedule cron patter
	ScheduleCronPerpetual bool     `json:"scheduleCronPerpetual"`                        // Whether a cron schedule should run in definitely
	ScheduleDealNumber    int      `json:"scheduleDealNumber"`                           // Number of deals per scheduled time
	TotalDealNumber       int      `json:"totalDealNumber"`                              // Total number of deals
	ScheduleDealSize      string   `json:"scheduleDealSize"`                             // Size of deals per schedule trigger in human readable format, i.e. 100 TiB
	TotalDealSize         string   `json:"totalDealSize"`                                // Total size of deals in human readable format, i.e. 100 TiB
	Notes                 string   `json:"notes"`                                        // Notes
	MaxPendingDealSize    string   `json:"maxPendingDealSize"`                           // Max pending deal size in human readable format, i.e. 100 TiB
	MaxPendingDealNumber  int      `json:"maxPendingDealNumber"`                         // Max pending deal number
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

// CreateHandler creates a new schedule based on the provided CreateRequest.
//
// The function performs the following steps:
//  1. Associates the provided context with the database connection.
//  2. Retrieves the preparation from the database using the ID from the request.
//  3. Parses the provided start delay and duration to ensure valid durations.
//  4. If a ScheduleCron string is provided, it validates its correctness.
//  5. Parses and validates the provided sizes: TotalDealSize, ScheduleDealSize,
//     and MaxPendingDealSize.
//  6. Verifies all provided piece CIDs in AllowedPieceCIDs to ensure their correctness.
//  7. Checks for the presence of wallets attached to the preparation.
//  8. Uses the lotusClient to retrieve the provider actor.
//  9. Constructs a new model.Schedule instance from the provided and parsed data.
//  10. Inserts the newly created schedule into the database.
//  11. Returns the newly created schedule.
//
// Parameters:
// - ctx: The context for the operation, used for timeouts and cancellation.
// - db: The database connection, used for CRUD operations related to schedules.
// - lotusClient: The Lotus client, used for Filecoin RPC calls.
// - request: The request object containing the data for the new schedule.
//
// Returns:
// - A pointer to the created model.Schedule if successful.
// - An error indicating the reason for any failure during the operation.
func (DefaultHandler) CreateHandler(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	request CreateRequest,
) (*model.Schedule, error) {
	db = db.WithContext(ctx)
	if request.ScheduleDealSize == "" {
		request.ScheduleDealSize = "0"
	}
	if request.TotalDealSize == "" {
		request.TotalDealSize = "0"
	}
	if request.MaxPendingDealSize == "" {
		request.MaxPendingDealSize = "0"
	}
	var preparation model.Preparation
	err := preparation.FindByIDOrName(db, request.Preparation, "Wallets")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %d not found", request.Preparation)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	startDelay, err := argToDuration(request.StartDelay)
	if err != nil {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid start delay %s", request.StartDelay)
	}

	duration, err := argToDuration(request.Duration)
	if err != nil {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid duration %s", request.Duration)
	}

	var scheduleCron string
	if request.ScheduleCron != "" {
		cronParser := cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
		_, err = cronParser.Parse(request.ScheduleCron)
		if err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid schedule cron %s", request.ScheduleCron)
		} else {
			scheduleCron = request.ScheduleCron
		}
	}
	var totalDealSize, scheduleDealSize, pendingDealSize uint64
	if request.TotalDealSize != "" {
		totalDealSize, err = humanize.ParseBytes(request.TotalDealSize)
		if err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid total deal size %s", request.TotalDealSize)
		}
	}
	if request.ScheduleDealSize != "" {
		scheduleDealSize, err = humanize.ParseBytes(request.ScheduleDealSize)
		if err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid schedule deal size %s", request.ScheduleDealSize)
		}
	}
	if request.MaxPendingDealSize != "" {
		pendingDealSize, err = humanize.ParseBytes(request.MaxPendingDealSize)
		if err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid max pending deal size %s", request.MaxPendingDealSize)
		}
	}
	if scheduleCron != "" && scheduleDealSize == 0 && request.ScheduleDealNumber == 0 {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "schedule deal number or size must be set when using cron schedule")
	}
	for _, pieceCID := range request.AllowedPieceCIDs {
		parsed, err := cid.Parse(pieceCID)
		if err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid allowed piece CID %s", pieceCID)
		}
		if parsed.Type() != cid.FilCommitmentUnsealed {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "allowed piece CID %s is not commp", pieceCID)
		}
	}

	if len(preparation.Wallets) == 0 {
		return nil, errors.Wrap(handlererror.ErrNotFound, "no wallet attached to preparation")
	}

	var providerActor string
	err = lotusClient.CallFor(ctx, &providerActor, "Filecoin.StateLookupID", request.Provider, nil)
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrapf(err, "provider %s cannot be resolved", request.Provider))
	}

	headers := make(map[string]string)
	for _, header := range request.HTTPHeaders {
		kv := strings.SplitN(header, "=", 2)
		if len(kv) != 2 {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid http header: %s", header)
		}
		headers[kv[0]], err = url.QueryUnescape(kv[1])
		if err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid http header: %s", header)
		}
	}

	schedule := model.Schedule{
		PreparationID:         preparation.ID,
		URLTemplate:           request.URLTemplate,
		HTTPHeaders:           headers,
		Provider:              request.Provider,
		TotalDealNumber:       request.TotalDealNumber,
		TotalDealSize:         int64(totalDealSize),
		Verified:              request.Verified,
		KeepUnsealed:          request.KeepUnsealed,
		AnnounceToIPNI:        request.IPNI,
		StartDelay:            startDelay,
		Duration:              duration,
		State:                 model.ScheduleActive,
		ScheduleDealNumber:    request.ScheduleDealNumber,
		ScheduleDealSize:      int64(scheduleDealSize),
		MaxPendingDealNumber:  request.MaxPendingDealNumber,
		MaxPendingDealSize:    int64(pendingDealSize),
		Notes:                 request.Notes,
		AllowedPieceCIDs:      underscore.Unique(request.AllowedPieceCIDs),
		ScheduleCron:          scheduleCron,
		PricePerGBEpoch:       request.PricePerGBEpoch,
		PricePerGB:            request.PricePerGB,
		PricePerDeal:          request.PricePerDeal,
		ScheduleCronPerpetual: request.ScheduleCronPerpetual,
	}

	if err := database.DoRetry(ctx, func() error {
		return db.Create(&schedule).Error
	}); err != nil {
		return nil, errors.WithStack(err)
	}
	return &schedule, nil
}

// @ID CreateSchedule
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
func _() {}
