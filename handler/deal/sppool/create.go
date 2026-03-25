package sppool

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/dustin/go-humanize"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

//nolint:lll
type CreateRequest struct {
	Name                  string   `json:"name"                  validation:"required"` // Pool name (must be unique)
	Notes                 string   `json:"notes"`                                       // Notes
	HTTPHeaders           []string `json:"httpHeaders"`                                 // http headers to be passed with the request (i.e. key=value)
	URLTemplate           string   `json:"urlTemplate"`                                 // URL template with PIECE_CID placeholder
	PricePerGBEpoch       float64  `default:"0"                  json:"pricePerGbEpoch"`
	PricePerGB            float64  `default:"0"                  json:"pricePerGb"`
	PricePerDeal          float64  `default:"0"                  json:"pricePerDeal"`
	Verified              bool     `default:"true"               json:"verified"`
	IPNI                  bool     `default:"true"               json:"ipni"`
	KeepUnsealed          bool     `default:"true"               json:"keepUnsealed"`
	StartDelay            string   `default:"72h"                json:"startDelay"`
	Duration              string   `default:"12840h"             json:"duration"`
	ScheduleCron          string   `json:"scheduleCron"`
	ScheduleCronPerpetual bool     `json:"scheduleCronPerpetual"`
	ScheduleDealNumber    int      `json:"scheduleDealNumber"`
	ScheduleDealSize      string   `json:"scheduleDealSize"`
	MaxPendingDealSize    string   `json:"maxPendingDealSize"`
	MaxPendingDealNumber  int      `json:"maxPendingDealNumber"`
	Force                 bool     `json:"force"`
}

func (DefaultHandler) CreateHandler(
	ctx context.Context,
	db *gorm.DB,
	request CreateRequest,
) (*model.SPPool, error) {
	db = db.WithContext(ctx)

	if request.Name == "" {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "name is required")
	}
	if request.ScheduleDealSize == "" {
		request.ScheduleDealSize = "0"
	}
	if request.MaxPendingDealSize == "" {
		request.MaxPendingDealSize = "0"
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
		if _, err := cronParser.Parse(request.ScheduleCron); err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid schedule cron %s", request.ScheduleCron)
		}
		scheduleCron = request.ScheduleCron
	}
	scheduleDealSize, err := humanize.ParseBytes(request.ScheduleDealSize)
	if err != nil {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid schedule deal size %s", request.ScheduleDealSize)
	}
	pendingDealSize, err := humanize.ParseBytes(request.MaxPendingDealSize)
	if err != nil {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid max pending deal size %s", request.MaxPendingDealSize)
	}
	if scheduleCron != "" && scheduleDealSize == 0 && request.ScheduleDealNumber == 0 {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "schedule deal number or size must be set when using cron schedule")
	}
	if scheduleCron == "" && (scheduleDealSize > 0 || request.ScheduleDealNumber > 0) {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "schedule cron must be set when using schedule deal number or size")
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

	pool := model.SPPool{
		Name:                  request.Name,
		State:                 model.SPPoolActive,
		Notes:                 request.Notes,
		URLTemplate:           request.URLTemplate,
		HTTPHeaders:           headers,
		PricePerGBEpoch:       request.PricePerGBEpoch,
		PricePerGB:            request.PricePerGB,
		PricePerDeal:          request.PricePerDeal,
		Verified:              request.Verified,
		KeepUnsealed:          request.KeepUnsealed,
		AnnounceToIPNI:        request.IPNI,
		StartDelay:            startDelay,
		Duration:              duration,
		ScheduleCron:          scheduleCron,
		ScheduleCronPerpetual: request.ScheduleCronPerpetual,
		ScheduleDealNumber:    request.ScheduleDealNumber,
		ScheduleDealSize:      int64(scheduleDealSize),
		MaxPendingDealNumber:  request.MaxPendingDealNumber,
		MaxPendingDealSize:    int64(pendingDealSize),
		Force:                 request.Force,
	}

	if err := database.DoRetry(ctx, func() error {
		return db.Create(&pool).Error
	}); err != nil {
		if util.IsDuplicateKeyError(err) {
			return nil, errors.Wrapf(handlererror.ErrDuplicateRecord, "pool with name %q already exists", request.Name)
		}
		return nil, errors.WithStack(err)
	}
	return &pool, nil
}

func argToDuration(s string) (time.Duration, error) {
	if s == "" {
		return 0, nil
	}
	d, err := time.ParseDuration(s)
	if err == nil {
		return d, nil
	}
	epochs, err2 := parseEpochs(s)
	if err2 != nil {
		return 0, errors.WithStack(err)
	}
	return time.Duration(epochs) * 30 * time.Second, nil
}

func parseEpochs(s string) (int64, error) {
	var v int64
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, errors.New("not a number")
		}
		v = v*10 + int64(c-'0')
	}
	return v, nil
}

// @ID CreateSPPool
// @Summary Create a new SP Pool
// @Description Create a new SP Pool with default deal parameters
// @Tags SP Pool
// @Accept json
// @Produce json
// @Param request body CreateRequest true "CreateRequest"
// @Success 200 {object} model.SPPool
// @Failure 400 {object} api.HTTPError
// @Failure 409 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /sp-pool [post]
func _() {}
