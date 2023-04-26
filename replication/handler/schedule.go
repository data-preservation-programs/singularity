package handler

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/dustin/go-humanize"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"time"
)

func validateProvider(provider string) error {
	return nil
}

func CreateHandler(c *cli.Context) error {
	logger := log.Logger("cli")
	log.SetAllLoggers(log.LevelInfo)

	if c.NArg() != 2 {
		return cli.Exit("DATASET_ID and PROVIDER are required", 1)
	}

	datasetName := c.Args().Get(0)
	dataset := model.Dataset{}
	provider := c.Args().Get(1)

	db := database.MustOpenFromCLI(c)
	err := db.Where("name = ?", datasetName).First(&dataset).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return cli.Exit("dataset not found", 1)
	}

	if err != nil {
		return cli.Exit("failed to query dataset: "+err.Error(), 1)
	}

	err = validateProvider(provider)
	if err != nil {
		return cli.Exit("invalid provider: "+err.Error(), 1)
	}

	urlTemplate := c.String("url-template")
	price := c.Float64("price")
	verified := c.Bool("verified")
	cronSchedule := c.String("cron-schedule")
	_, err = cron.ParseStandard(cronSchedule)
	if err != nil {
		return cli.Exit("invalid cron schedule: "+err.Error(), 1)
	}
	startDelay := time.Duration(c.Float64("start-delay") * float64(24*time.Hour))
	duration := time.Duration(c.Float64("duration") * float64(24*time.Hour))
	notes := c.String("notes")
	scheduleDealNumber := c.Uint64("schedule-deal-number")
	totalDealNumber := c.Uint64("total-deal-number")
	maxPendingDealNumber := c.Uint64("max-pending-deal-number")
	totalDealSize, err := humanize.ParseBytes(c.String("total-deal-size"))
	if err != nil {
		return cli.Exit("invalid value for total-deal-size: "+err.Error(), 1)
	}
	scheduleDealSize, err := humanize.ParseBytes(c.String("schedule-deal-size"))
	if err != nil {
		return cli.Exit("invalid value for schedule-deal-size: "+err.Error(), 1)
	}
	maxPendingDealSize, err := humanize.ParseBytes(c.String("max-pending-deal-size"))
	if err != nil {
		return cli.Exit("invalid value for max-pending-deal-size: "+err.Error(), 1)
	}
	ipni := c.Bool("ipni")
	unsealed := c.Bool("keep-unsealed")
	httpHeaders := c.StringSlice("http-header")

	schedule := model.Schedule{
		DatasetID:            dataset.ID,
		UrlTemplate:          urlTemplate,
		Provider:             provider,
		Price:                price,
		TotalDealNumber:      totalDealNumber,
		TotalDealSize:        totalDealSize,
		Verified:             verified,
		StartDelay:           startDelay,
		Duration:             duration,
		State:                model.ScheduleStarted,
		SchedulePattern:      cronSchedule,
		ScheduleDealNumber:   scheduleDealNumber,
		ScheduleDealSize:     scheduleDealSize,
		MaxPendingDealNumber: maxPendingDealNumber,
		MaxPendingDealSize:   maxPendingDealSize,
		Notes:                notes,
		AnnounceToIPNI:       ipni,
		KeepUnsealed:         unsealed,
		HttpHeaders:          httpHeaders,
	}

	err = db.Create(&schedule).Error
	if err != nil {
		return cli.Exit("failed to create schedule: "+err.Error(), 1)
	}

	logger.Infof("created schedule %d", schedule.ID)
	return nil
}
