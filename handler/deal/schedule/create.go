package schedule

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type CreateRequest struct {
	DatasetName          string        `json:"datasetName"`
	Provider             string        `json:"provider"`
	HTTPHeaders          []string      `json:"httpHeaders"`
	URLTemplate          string        `json:"urlTemplate"`
	Price                float64       `json:"price"`
	Verified             bool          `json:"verified"`
	IPNI                 bool          `json:"ipni"`
	KeepUnsealed         bool          `json:"keepUnsealed"`
	ScheduleInterval     time.Duration `json:"scheduleInterval"`
	StartDelayDays       float64       `json:"startDelayDays"`
	DurationDays         float64       `json:"durationDays"`
	ScheduleDealNumber   int           `json:"scheduleDealNumber"`
	TotalDealNumber      int           `json:"totalDealNumber"`
	ScheduleDealSize     string        `json:"scheduleDealSize"`
	TotalDealSize        string        `json:"totalDealSize"`
	Notes                string        `json:"notes"`
	MaxPendingDealSize   string        `json:"maxPendingDealSize"`
	MaxPendingDealNumber int           `json:"maxPendingDealNumber"`
	AllowedPieceCIDs     []string      `json:"allowedPieceCIDs"`
}

// CreateHandler godoc
// @Summary Create a new schedule
// @Description Create a new schedule
// @Tags Deal Schedule
// @Accept json
// @Produce json
// @Param schedule body CreateRequest true "CreateRequest"
// @Success 200 {object} model.Schedule
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /deal/schedule [post]
func CreateHandler(
	db *gorm.DB,
	request CreateRequest,
) (*model.Schedule, *handler.Error) {
	dataset, err := database.FindDatasetByName(db, request.DatasetName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewBadRequestString("dataset not found")
	}
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}
	totalDealSize, err := humanize.ParseBytes(request.TotalDealSize)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid total deal size")
	}
	scheduleDealSize, err := humanize.ParseBytes(request.ScheduleDealSize)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid schedule deal size")
	}
	pendingDealSize, err := humanize.ParseBytes(request.MaxPendingDealSize)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid max pending deal size")
	}
	intervalSeconds := uint64(request.ScheduleInterval.Seconds())
	startDelay := time.Duration(float64(time.Hour*24) * request.StartDelayDays)
	duration := time.Duration(float64(time.Hour*24) * request.DurationDays)
	schedule := model.Schedule{
		DatasetID:               dataset.ID,
		URLTemplate:             request.URLTemplate,
		HTTPHeaders:             request.HTTPHeaders,
		Provider:                request.Provider,
		Price:                   request.Price,
		TotalDealNumber:         request.TotalDealNumber,
		TotalDealSize:           int64(totalDealSize),
		Verified:                request.Verified,
		KeepUnsealed:            request.KeepUnsealed,
		AnnounceToIPNI:          request.IPNI,
		StartDelay:              startDelay,
		Duration:                duration,
		State:                   model.ScheduleActive,
		ScheduleIntervalSeconds: intervalSeconds,
		ScheduleDealNumber:      request.ScheduleDealNumber,
		ScheduleDealSize:        int64(scheduleDealSize),
		MaxPendingDealNumber:    request.MaxPendingDealNumber,
		MaxPendingDealSize:      int64(pendingDealSize),
		Notes:                   request.Notes,
		AllowedPieceCIDs:        request.AllowedPieceCIDs,
	}
	if err := db.Create(&schedule).Error; err != nil {
		return nil, handler.NewHandlerError(err)
	}
	return &schedule, nil
}
