package dealmaker

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/service/healthcheck"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/google/uuid"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var logger = log.Logger("dealmaker")

type DealMakerService struct {
	db                       *gorm.DB
	walletChooser            replication.WalletChooser
	dealMaker                replication.DealMakerImpl
	workerID                 uuid.UUID
	activeSchedule           map[uint32]*model.Schedule
	activeScheduleCancelFunc map[uint32]context.CancelFunc
	cronEntries              map[uint32]cron.EntryID
	cron                     *cron.Cron
	mutex                    sync.Mutex
}

type sumResult struct {
	DealNumber int
	DealSize   int64
}

type cronLogger struct{}

func (c cronLogger) Info(msg string, keysAndValues ...interface{}) {
	logger.Infow(msg, keysAndValues...)
}

func (c cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, "err", err)
	logger.Errorw(msg, keysAndValues...)
}

func (d *DealMakerService) hasSchedule(scheduleID uint32) bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, ok := d.activeSchedule[scheduleID]
	return ok
}

func (d *DealMakerService) runScheduleAndUpdateState(ctx context.Context, schedule *model.Schedule) {
	state, err := d.runSchedule(ctx, schedule)
	updates := make(map[string]interface{})
	if state != "" {
		updates["state"] = state
	}
	if err != nil {
		updates["error_message"] = err.Error()
	}
	if len(updates) > 0 {
		err = d.db.Model(&model.Schedule{}).Updates(updates).Error
		if err != nil {
			logger.Errorw("failed to update schedule", "schedule", schedule.ID, "error", err)
		}
	}
	if state == model.ScheduleCompleted {
		d.removeSchedule(*schedule)
	}
}

func (d *DealMakerService) addSchedule(ctx context.Context, schedule model.Schedule) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	scheduleCtx, cancel := context.WithCancel(ctx)
	if schedule.ScheduleCron == "" {
		go d.runScheduleAndUpdateState(ctx, &schedule)
		d.activeSchedule[schedule.ID] = &schedule
		d.activeScheduleCancelFunc[schedule.ID] = cancel
		return nil
	}

	entryID, err := d.cron.AddFunc(schedule.ScheduleCron, func() {
		d.runScheduleAndUpdateState(scheduleCtx, &schedule)
	})
	if err != nil {
		cancel()
		return errors.Wrap(err, "failed to add cron job")
	}
	d.activeSchedule[schedule.ID] = &schedule
	d.cronEntries[schedule.ID] = entryID
	d.activeScheduleCancelFunc[schedule.ID] = cancel
	return nil
}

func (d *DealMakerService) removeSchedule(schedule model.Schedule) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if schedule.ScheduleCron == "" {
		d.activeScheduleCancelFunc[schedule.ID]()
		delete(d.activeSchedule, schedule.ID)
		delete(d.activeScheduleCancelFunc, schedule.ID)
	}

	d.cron.Remove(d.cronEntries[schedule.ID])
	d.activeScheduleCancelFunc[schedule.ID]()
	delete(d.activeSchedule, schedule.ID)
	delete(d.activeScheduleCancelFunc, schedule.ID)
}

func (d *DealMakerService) updateSchedule(ctx context.Context, schedule model.Schedule) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	existing := d.activeSchedule[schedule.ID]
	if existing.ScheduleCron == "" && schedule.ScheduleCron != "" {
		return errors.New("cannot update schedule - changed from oneoff to cron")
	}
	if existing.ScheduleCron != "" && schedule.ScheduleCron == "" {
		return errors.New("cannot update schedule - changed from cron to oneoff")
	}

	*d.activeSchedule[schedule.ID] = schedule
	if schedule.ScheduleCron == "" {
		return nil
	}

	if d.activeSchedule[schedule.ID].ScheduleCron != schedule.ScheduleCron {
		d.cron.Remove(d.cronEntries[schedule.ID])
		d.activeScheduleCancelFunc[schedule.ID]()
		scheduleCtx, cancel := context.WithCancel(ctx)
		entryID, err := d.cron.AddFunc(schedule.ScheduleCron, func() {
			d.runScheduleAndUpdateState(scheduleCtx, &schedule)
		})
		if err != nil {
			cancel()
			return errors.Wrap(err, "failed to add cron job")
		}
		d.cronEntries[schedule.ID] = entryID
		d.activeScheduleCancelFunc[schedule.ID] = cancel
	}

	return nil
}

func (d *DealMakerService) runSchedule(ctx context.Context, schedule *model.Schedule) (model.ScheduleState, error) {
	var pending sumResult
	err := d.db.WithContext(ctx).Model(&model.Deal{}).
		Where("schedule_id = ? AND state IN (?)", schedule.ID, []model.DealState{
			model.DealProposed, model.DealPublished,
		}).Select("COUNT(*) AS deal_number, SUM(piece_cid) AS deal_size").Scan(&pending).Error
	if err != nil {
		return model.ScheduleError, errors.Wrap(err, "failed to count pending deals")
	}
	var total sumResult
	err = d.db.WithContext(ctx).Model(&model.Deal{}).
		Where("schedule_id = ? AND state IN (?)", schedule.ID, []model.DealState{
			model.DealActive}).Select("COUNT(*) AS deal_number, SUM(piece_cid) AS deal_size").Scan(&total).Error
	if err != nil {
		return model.ScheduleError, errors.Wrap(err, "failed to count total active deals")
	}

	var current sumResult

	for {
		var car model.Car
		var dealModel *model.Deal
		var walletObj model.Wallet
		if schedule.MaxPendingDealNumber > 0 && pending.DealNumber >= schedule.MaxPendingDealNumber {
			logger.Infow("skipping this time since the max pending deal is reached", "schedule_id", schedule.ID)
			goto waitAndNext
		}
		if schedule.MaxPendingDealSize > 0 && pending.DealSize >= schedule.MaxPendingDealSize {
			logger.Infow("skipping this time since the max pending deal size is reached", "schedule_id", schedule.ID)
			goto waitAndNext
		}
		if schedule.TotalDealNumber > 0 && total.DealNumber >= schedule.TotalDealNumber {
			logger.Infow("completing since the total deal number is reached", "schedule_id", schedule.ID)
			return model.ScheduleCompleted, nil
		}
		if schedule.TotalDealSize > 0 && total.DealSize >= schedule.TotalDealSize {
			logger.Infow("completing since the total deal size is reached", "schedule_id", schedule.ID)
			return model.ScheduleCompleted, nil
		}
		if schedule.ScheduleCron != "" && schedule.ScheduleDealNumber > 0 && current.DealNumber >= schedule.ScheduleDealNumber {
			logger.Infow("completing this batch since the schedule deal number is reached", "schedule_id", schedule.ID)
			return "", nil
		}
		if schedule.ScheduleCron != "" && schedule.ScheduleDealSize > 0 && current.DealSize >= schedule.ScheduleDealSize {
			logger.Infow("completing this batch since the schedule deal size is reached", "schedule_id", schedule.ID)
			return "", nil
		}

		err = d.db.WithContext(ctx).Where("dataset_id = ? AND piece_cid NOT IN (?)",
			schedule.DatasetID,
			d.db.Table("deals").Select("piece_cid").
				Where("provider = ? AND state IN (?)",
					schedule.Provider,
					[]model.DealState{
						model.DealProposed, model.DealPublished, model.DealActive,
					})).First(&car).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Infow("no more pieces to send deal", "schedule_id", schedule.ID)
			return model.ScheduleCompleted, nil
		}
		if err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed to find car")
		}

		walletObj, err = d.walletChooser.Choose(ctx, schedule.Dataset.Wallets)
		if err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed to choose wallet")
		}

		dealModel, err = d.dealMaker.MakeDeal(
			ctx,
			walletObj,
			car,
			replication.DealConfig{
				Provider:        schedule.Provider,
				StartDelay:      schedule.StartDelay,
				Duration:        schedule.Duration,
				Verified:        schedule.Verified,
				HTTPHeaders:     schedule.HTTPHeaders,
				URLTemplate:     schedule.URLTemplate,
				KeepUnsealed:    schedule.KeepUnsealed,
				AnnounceToIPNI:  schedule.AnnounceToIPNI,
				PricePerDeal:    schedule.PricePerDeal,
				PricePerGB:      schedule.PricePerGB,
				PricePerGBEpoch: schedule.PricePerGBEpoch,
			})
		if err != nil {
			logger.Errorw("failed to make deal", "error", err)
			goto waitAndNext
		}
		dealModel.ScheduleID = &schedule.ID

		err = d.db.Create(dealModel).Error
		if err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed to create deal")
		}

		current.DealSize += car.PieceSize
		current.DealNumber += 1
		total.DealSize += car.PieceSize
		total.DealNumber += 1
		pending.DealSize += car.PieceSize
		pending.DealNumber += 1

	waitAndNext:
		select {
		case <-ctx.Done():
			return "", nil
		case <-time.After(time.Minute):
		}
	}
}

func NewDealMakerService(db *gorm.DB, lotusURL string,
	lotusToken string) (*DealMakerService, error) {
	h, err := util.InitHost(nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init host")
	}
	lotusClient := util.NewLotusClient(lotusURL, lotusToken)
	dealMaker := replication.NewDealMaker(lotusClient, h, time.Hour, time.Minute)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init deal maker")
	}
	return &DealMakerService{
		db:                       db,
		activeScheduleCancelFunc: make(map[uint32]context.CancelFunc),
		activeSchedule:           make(map[uint32]*model.Schedule),
		walletChooser:            &replication.DefaultWalletChooser{},
		dealMaker:                dealMaker,
		workerID:                 uuid.New(),
		cron:                     cron.New(cron.WithLogger(&cronLogger{}), cron.WithLocation(time.UTC)),
	}, nil
}

var staleThreshold = time.Minute * 5

func (d *DealMakerService) Run(ctx context.Context) error {
	var activeWorkerCount int64
	err := d.db.Where("work_type = ? AND last_heartbeat > ?", model.DealMaking, time.Now().UTC().Add(-staleThreshold)).
		Count(&activeWorkerCount).Error
	if err != nil {
		return errors.Wrap(err, "failed to count active workers")
	}
	if activeWorkerCount > 0 {
		return errors.New("deal maker already running")
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGTRAP)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	getState := func() healthcheck.State {
		return healthcheck.State{
			WorkType: model.DealMaking,
		}
	}
	healthcheck.HealthCheck(d.db, d.workerID, getState)
	go healthcheck.StartHealthCheck(ctx, d.db, d.workerID, getState)

	for {
		var schedules []model.Schedule
		scheduleMap := map[uint32]model.Schedule{}
		err := d.db.WithContext(ctx).Preload("Dataset.Wallets").Where("state = ?",
			model.ScheduleActive).Find(&schedules).Error
		if err != nil {
			logger.Errorw("failed to get schedules", "error", err)
			goto nextloop
		}
		for _, schedule := range schedules {
			scheduleMap[schedule.ID] = schedule
		}
		// Cancel all jobs that are no longer active
		for id := range d.activeSchedule {
			if s, ok := scheduleMap[id]; !ok {
				d.removeSchedule(s)
			}
		}

		for _, schedule := range schedules {
			if d.hasSchedule(schedule.ID) {
				err = d.updateSchedule(ctx, schedule)
				if err != nil {
					logger.Errorw("failed to update schedule", "error", err)
				}
			} else {
				err = d.addSchedule(ctx, schedule)
				if err != nil {
					logger.Errorw("failed to add schedule", "error", err)
				}
			}
		}
	nextloop:
		select {
		case <-signalChan:
			logger.Infow("received signal, stopping")
			for _, cancel := range d.activeScheduleCancelFunc {
				cancel()
			}
			//nolint:errcheck
			d.cleanup()
			return cli.Exit("received signal", 130)
		case <-ctx.Done():
			//nolint:errcheck
			d.cleanup()
			return ctx.Err()
		case <-time.After(5 * time.Second):
		}
	}
}

func (d *DealMakerService) cleanup() error {
	return d.db.Where("id = ?", d.workerID).Delete(&model.Worker{}).Error
}
