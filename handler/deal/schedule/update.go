package schedule

import (
	"context"
	"net/url"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/dustin/go-humanize"
	"github.com/ipfs/go-cid"
	"github.com/rjNemo/underscore"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

//nolint:lll
type UpdateRequest struct {
	HTTPHeaders           []string `json:"httpHeaders"`                                  // http headers to be passed with the request (i.e. key=value)
	URLTemplate           *string  `json:"urlTemplate"`                                  // URL template with PIECE_CID placeholder for boost to fetch the CAR file, i.e. http://127.0.0.1/piece/{PIECE_CID}.car
	PricePerGBEpoch       *float64 `default:"0"                  json:"pricePerGbEpoch"` // Price in FIL per GiB per epoch
	PricePerGB            *float64 `default:"0"                  json:"pricePerGb"`      // Price in FIL  per GiB
	PricePerDeal          *float64 `default:"0"                  json:"pricePerDeal"`    // Price in FIL per deal
	Verified              *bool    `default:"true"               json:"verified"`        // Whether the deal should be verified
	IPNI                  *bool    `default:"true"               json:"ipni"`            // Whether the deal should be IPNI
	KeepUnsealed          *bool    `default:"true"               json:"keepUnsealed"`    // Whether the deal should be kept unsealed
	StartDelay            *string  `default:"72h"                json:"startDelay"`      // Deal start delay in epoch or in duration format, i.e. 1000, 72h
	Duration              *string  `default:"12840h"             json:"duration"`        // Duration in epoch or in duration format, i.e. 1500000, 2400h
	ScheduleCron          *string  `json:"scheduleCron"`                                 // Schedule cron patter
	ScheduleCronPerpetual *bool    `json:"scheduleCronPerpetual"`                        // Whether a cron schedule should run in definitely
	ScheduleDealNumber    *int     `json:"scheduleDealNumber"`                           // Number of deals per scheduled time
	TotalDealNumber       *int     `json:"totalDealNumber"`                              // Total number of deals
	ScheduleDealSize      *string  `json:"scheduleDealSize"`                             // Size of deals per schedule trigger in human readable format, i.e. 100 TiB
	TotalDealSize         *string  `json:"totalDealSize"`                                // Total size of deals in human readable format, i.e. 100 TiB
	Notes                 *string  `json:"notes"`                                        // Notes
	MaxPendingDealSize    *string  `json:"maxPendingDealSize"`                           // Max pending deal size in human readable format, i.e. 100 TiB
	MaxPendingDealNumber  *int     `json:"maxPendingDealNumber"`                         // Max pending deal number
	//nolint:tagliatelle
	AllowedPieceCIDs []string `json:"allowedPieceCids"` // Allowed piece CIDs in this schedule
}

// UpdateHandler modifies an existing schedule record based on the provided update request.
//
// It looks for the schedule record by the given schedule ID. If found, it processes
// the provided UpdateRequest to determine which fields should be updated. Once the
// desired changes are captured, the function commits these updates to the database.
//
// Parameters:
//   - ctx: The context for managing timeouts and cancellation.
//   - db: The gorm.DB instance for database operations.
//   - id: The ID of the schedule to be updated.
//   - request: The UpdateRequest containing the desired changes.
//
// Returns:
//   - The updated model.Schedule if the operation is successful.
//   - An error if any issues occur during the operation, including invalid parameters or database errors.
func (DefaultHandler) UpdateHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
	request UpdateRequest,
) (*model.Schedule, error) {
	db = db.WithContext(ctx)

	var schedule model.Schedule
	err := db.First(&schedule, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "schedule %d not found", id)
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	updates := make(map[string]interface{})
	if request.HTTPHeaders != nil {
		headers := schedule.HTTPHeaders
		if headers == nil {
			headers = make(map[string]string)
		}
		for _, header := range request.HTTPHeaders {
			if header == "" {
				headers = make(map[string]string)
				continue
			}
			kv := strings.SplitN(header, "=", 2)
			if len(kv) != 2 {
				return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid http header: %s", header)
			}
			value, err := url.QueryUnescape(kv[1])
			if err != nil {
				return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid http header: %s", header)
			}
			if value == "" {
				delete(headers, kv[0])
			} else {
				headers[kv[0]] = value
			}
		}
		updates["http_headers"] = headers
	}

	if request.URLTemplate != nil {
		updates["url_template"] = *request.URLTemplate
	}

	if request.PricePerGBEpoch != nil {
		updates["price_per_gb_epoch"] = *request.PricePerGBEpoch
	}

	if request.PricePerGB != nil {
		updates["price_per_gb"] = *request.PricePerGB
	}

	if request.PricePerDeal != nil {
		updates["price_per_deal"] = *request.PricePerDeal
	}

	if request.Verified != nil {
		updates["verified"] = *request.Verified
	}

	if request.IPNI != nil {
		updates["announce_to_ipni"] = *request.IPNI
	}

	if request.KeepUnsealed != nil {
		updates["keep_unsealed"] = *request.KeepUnsealed
	}

	if request.StartDelay != nil {
		startDelay, err := argToDuration(*request.StartDelay)
		if err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid start delay: %s", *request.StartDelay)
		}
		updates["start_delay"] = startDelay
	}

	if request.Duration != nil {
		duration, err := argToDuration(*request.Duration)
		if err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid duration: %s", *request.Duration)
		}
		updates["duration"] = duration
	}

	if request.ScheduleCron != nil {
		if *request.ScheduleCron == "" && schedule.ScheduleCron != "" {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "Cannot switch from cron to non-cron schedule")
		}
		if *request.ScheduleCron != "" && schedule.ScheduleCron == "" {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "Cannot switch from non-cron to cron schedule")
		}
		cronParser := cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
		var scheduleCron string
		if *request.ScheduleCron != "" {
			_, err = cronParser.Parse(*request.ScheduleCron)
			if err != nil {
				return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid schedule cron %s", request.ScheduleCron)
			} else {
				scheduleCron = *request.ScheduleCron
			}
		}
		updates["schedule_cron"] = scheduleCron
	}

	if request.ScheduleCronPerpetual != nil {
		updates["schedule_cron_perpetual"] = *request.ScheduleCronPerpetual
	}

	if request.ScheduleDealNumber != nil {
		updates["schedule_deal_number"] = *request.ScheduleDealNumber
	}

	if request.TotalDealNumber != nil {
		updates["total_deal_number"] = *request.TotalDealNumber
	}

	if request.MaxPendingDealNumber != nil {
		updates["max_pending_deal_number"] = *request.MaxPendingDealNumber
	}

	if request.Notes != nil {
		updates["notes"] = *request.Notes
	}

	if len(request.AllowedPieceCIDs) > 0 {
		for _, pieceCID := range request.AllowedPieceCIDs {
			parsed, err := cid.Parse(pieceCID)
			if err != nil {
				return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid allowed piece CID %s", pieceCID)
			}
			if parsed.Type() != cid.FilCommitmentUnsealed {
				return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "allowed piece CID %s is not commp", pieceCID)
			}
		}
		updates["allowed_piece_cids"] = model.StringSlice(underscore.Unique(append(schedule.AllowedPieceCIDs, request.AllowedPieceCIDs...)))
	}

	if request.TotalDealSize != nil {
		totalDealSize := uint64(0)
		if *request.TotalDealSize != "" {
			totalDealSize, err = humanize.ParseBytes(*request.TotalDealSize)
			if err != nil {
				return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid total deal size: %s", *request.TotalDealSize)
			}
		}
		updates["total_deal_size"] = totalDealSize
	}

	if request.ScheduleDealSize != nil {
		scheduleDealSize := uint64(0)
		if *request.ScheduleDealSize != "" {
			scheduleDealSize, err = humanize.ParseBytes(*request.ScheduleDealSize)
			if err != nil {
				return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid schedule deal size: %s", *request.ScheduleDealSize)
			}
		}
		updates["schedule_deal_size"] = scheduleDealSize
	}

	if request.MaxPendingDealSize != nil {
		maxPendingDealSize := uint64(0)
		if *request.MaxPendingDealSize != "" {
			maxPendingDealSize, err = humanize.ParseBytes(*request.MaxPendingDealSize)
			if err != nil {
				return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid max pending deal size: %s", *request.MaxPendingDealSize)
			}
		}
		updates["max_pending_deal_size"] = maxPendingDealSize
	}

	err = db.Model(&schedule).Updates(updates).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &schedule, nil
}

// @ID UpdateSchedule
// @Summary Update a schedule
// @Description Update a schedule
// @Tags Deal Schedule
// @Accept json
// @Produce json
// @Param id path int true "Schedule ID"
// @Param body body UpdateRequest true "Update request"
// @Success 200 {object} model.Schedule
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /schedule/{id} [patch]
func _() {}
