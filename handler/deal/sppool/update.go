package sppool

import (
	"context"
	"net/url"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/dustin/go-humanize"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

//nolint:lll
type UpdateRequest struct {
	Notes                 *string  `json:"notes"`
	HTTPHeaders           []string `json:"httpHeaders"`
	URLTemplate           *string  `json:"urlTemplate"`
	PricePerGBEpoch       *float64 `json:"pricePerGbEpoch"`
	PricePerGB            *float64 `json:"pricePerGb"`
	PricePerDeal          *float64 `json:"pricePerDeal"`
	Verified              *bool    `json:"verified"`
	IPNI                  *bool    `json:"ipni"`
	KeepUnsealed          *bool    `json:"keepUnsealed"`
	StartDelay            *string  `json:"startDelay"`
	Duration              *string  `json:"duration"`
	ScheduleCron          *string  `json:"scheduleCron"`
	ScheduleCronPerpetual *bool    `json:"scheduleCronPerpetual"`
	ScheduleDealNumber    *int     `json:"scheduleDealNumber"`
	ScheduleDealSize      *string  `json:"scheduleDealSize"`
	MaxPendingDealSize    *string  `json:"maxPendingDealSize"`
	MaxPendingDealNumber  *int     `json:"maxPendingDealNumber"`
	Force                 *bool    `json:"force"`
}

func (DefaultHandler) UpdateHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
	request UpdateRequest,
) (*model.SPPool, error) {
	db = db.WithContext(ctx)
	var pool model.SPPool
	err := db.First(&pool, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "sp pool %d not found", id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if request.Notes != nil {
		pool.Notes = *request.Notes
	}
	if request.URLTemplate != nil {
		pool.URLTemplate = *request.URLTemplate
	}
	if request.PricePerGBEpoch != nil {
		pool.PricePerGBEpoch = *request.PricePerGBEpoch
	}
	if request.PricePerGB != nil {
		pool.PricePerGB = *request.PricePerGB
	}
	if request.PricePerDeal != nil {
		pool.PricePerDeal = *request.PricePerDeal
	}
	if request.Verified != nil {
		pool.Verified = *request.Verified
	}
	if request.IPNI != nil {
		pool.AnnounceToIPNI = *request.IPNI
	}
	if request.KeepUnsealed != nil {
		pool.KeepUnsealed = *request.KeepUnsealed
	}
	if request.StartDelay != nil {
		d, err := argToDuration(*request.StartDelay)
		if err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid start delay %s", *request.StartDelay)
		}
		pool.StartDelay = d
	}
	if request.Duration != nil {
		d, err := argToDuration(*request.Duration)
		if err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid duration %s", *request.Duration)
		}
		pool.Duration = d
	}
	if request.ScheduleCron != nil {
		if *request.ScheduleCron != "" {
			cronParser := cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
			if _, err := cronParser.Parse(*request.ScheduleCron); err != nil {
				return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid schedule cron %s", *request.ScheduleCron)
			}
		}
		pool.ScheduleCron = *request.ScheduleCron
	}
	if request.ScheduleCronPerpetual != nil {
		pool.ScheduleCronPerpetual = *request.ScheduleCronPerpetual
	}
	if request.ScheduleDealNumber != nil {
		pool.ScheduleDealNumber = *request.ScheduleDealNumber
	}
	if request.ScheduleDealSize != nil {
		sz, err := humanize.ParseBytes(*request.ScheduleDealSize)
		if err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid schedule deal size %s", *request.ScheduleDealSize)
		}
		pool.ScheduleDealSize = int64(sz)
	}
	if request.MaxPendingDealSize != nil {
		sz, err := humanize.ParseBytes(*request.MaxPendingDealSize)
		if err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid max pending deal size %s", *request.MaxPendingDealSize)
		}
		pool.MaxPendingDealSize = int64(sz)
	}
	if request.MaxPendingDealNumber != nil {
		pool.MaxPendingDealNumber = *request.MaxPendingDealNumber
	}
	if request.Force != nil {
		pool.Force = *request.Force
	}
	if len(request.HTTPHeaders) > 0 {
		headers := make(map[string]string)
		for _, header := range request.HTTPHeaders {
			kv := strings.SplitN(header, "=", 2)
			if len(kv) != 2 {
				return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid http header: %s", header)
			}
			var err error
			headers[kv[0]], err = url.QueryUnescape(kv[1])
			if err != nil {
				return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid http header: %s", header)
			}
		}
		pool.HTTPHeaders = headers
	}

	if err := database.DoRetry(ctx, func() error {
		return db.Save(&pool).Error
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	// Reconcile to propagate updated defaults to generated schedules.
	if err := reconcile(ctx, db, &pool); err != nil {
		return nil, err
	}

	return &pool, nil
}

// @ID UpdateSPPool
// @Summary Update an SP Pool
// @Tags SP Pool
// @Accept json
// @Produce json
// @Param id path int true "SP Pool ID"
// @Param request body UpdateRequest true "UpdateRequest"
// @Success 200 {object} model.SPPool
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /sp-pool/{id} [patch]
func _() {}
