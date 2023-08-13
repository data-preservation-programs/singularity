package dealpusher

import (
	"context"
	"sync"
	"time"

	"github.com/avast/retry-go"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/robfig/cron/v3"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/service/healthcheck"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/google/uuid"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var Logger = log.Logger("dealpusher")

var waitPendingInterval = time.Minute

type DealPusher struct {
	db                       *gorm.DB
	walletChooser            replication.WalletChooser
	dealMaker                replication.DealMaker
	workerID                 uuid.UUID
	activeSchedule           map[uint32]*model.Schedule
	activeScheduleCancelFunc map[uint32]context.CancelFunc
	cronEntries              map[uint32]cron.EntryID
	cron                     *cron.Cron
	mutex                    sync.Mutex
	sendDealRetry            uint
}

func (*DealPusher) Name() string {
	return "DealPusher"
}

type sumResult struct {
	DealNumber int
	DealSize   int64
}

type cronLogger struct{}

func (c cronLogger) Info(msg string, keysAndValues ...any) {
	Logger.Infow(msg, keysAndValues...)
}

func (c cronLogger) Error(err error, msg string, keysAndValues ...any) {
	keysAndValues = append(keysAndValues, "err", err)
	Logger.Errorw(msg, keysAndValues...)
}

func (d *DealPusher) hasSchedule(scheduleID uint32) bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, ok := d.activeSchedule[scheduleID]
	return ok
}

func (d *DealPusher) runScheduleAndUpdateState(ctx context.Context, schedule *model.Schedule) {
	state, err := d.runSchedule(ctx, schedule)
	updates := make(map[string]any)
	if err != nil {
		updates["error_message"] = err.Error()
		if schedule.ScheduleCron == "" {
			state = model.ScheduleError
		}
	}
	if state != "" {
		updates["state"] = state
	}
	if len(updates) > 0 {
		Logger.Debugw("updating schedule", "schedule", schedule.ID, "updates", updates)
		err = d.db.Model(schedule).Updates(updates).Error
		if err != nil {
			Logger.Errorw("failed to update schedule", "schedule", schedule.ID, "error", err)
		}
	}
	if state == model.ScheduleCompleted {
		Logger.Infow("schedule completed", "schedule", schedule.ID)
		d.removeSchedule(*schedule)
	}
	if state == model.ScheduleError {
		Logger.Errorw("schedule error", "schedule", schedule.ID, "error", err)
		d.removeSchedule(*schedule)
	}
}

func (d *DealPusher) addSchedule(ctx context.Context, schedule model.Schedule) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	scheduleCtx, cancel := context.WithCancel(ctx)
	if schedule.ScheduleCron == "" {
		d.activeSchedule[schedule.ID] = &schedule
		d.activeScheduleCancelFunc[schedule.ID] = cancel
		go d.runScheduleAndUpdateState(scheduleCtx, &schedule)
		return nil
	}

	d.activeSchedule[schedule.ID] = &schedule
	d.activeScheduleCancelFunc[schedule.ID] = cancel
	entryID, err := d.cron.AddFunc(schedule.ScheduleCron, func() {
		d.runScheduleAndUpdateState(scheduleCtx, &schedule)
	})
	if err != nil {
		cancel()
		delete(d.activeSchedule, schedule.ID)
		delete(d.activeScheduleCancelFunc, schedule.ID)
		return errors.Wrap(err, "failed to add cron job")
	}
	d.cronEntries[schedule.ID] = entryID
	return nil
}

func (d *DealPusher) removeSchedule(schedule model.Schedule) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.removeScheduleUnsafe(schedule)
}

func (d *DealPusher) removeScheduleUnsafe(schedule model.Schedule) {
	if schedule.ScheduleCron == "" {
		d.activeScheduleCancelFunc[schedule.ID]()
		delete(d.activeSchedule, schedule.ID)
		delete(d.activeScheduleCancelFunc, schedule.ID)
		return
	}

	d.cron.Remove(d.cronEntries[schedule.ID])
	d.activeScheduleCancelFunc[schedule.ID]()
	delete(d.activeSchedule, schedule.ID)
	delete(d.activeScheduleCancelFunc, schedule.ID)
}

func (d *DealPusher) updateSchedule(ctx context.Context, schedule model.Schedule) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	existing, ok := d.activeSchedule[schedule.ID]
	if !ok {
		return nil
	}
	if existing.ScheduleCron == "" && schedule.ScheduleCron != "" {
		Logger.Warnw("schedule changed from oneoff to cron", "schedule_id", schedule.ID)
		d.removeScheduleUnsafe(*existing)
	}
	if existing.ScheduleCron != "" && schedule.ScheduleCron == "" {
		Logger.Warnw("schedule changed from cron to oneoff", "schedule_id", schedule.ID)
		d.removeScheduleUnsafe(*existing)
	}

	if schedule.ScheduleCron == "" {
		*d.activeSchedule[schedule.ID] = schedule
		return nil
	}

	if d.activeSchedule[schedule.ID].ScheduleCron != schedule.ScheduleCron {
		Logger.Info("cron schedule has changed", "old", d.activeSchedule[schedule.ID].ScheduleCron, "new", schedule.ScheduleCron)
		*d.activeSchedule[schedule.ID] = schedule
		d.cron.Remove(d.cronEntries[schedule.ID])
		d.activeScheduleCancelFunc[schedule.ID]()
		scheduleCtx, cancel := context.WithCancel(ctx)
		d.activeScheduleCancelFunc[schedule.ID] = cancel
		entryID, err := d.cron.AddFunc(schedule.ScheduleCron, func() {
			d.runScheduleAndUpdateState(scheduleCtx, &schedule)
		})
		if err != nil {
			cancel()
			delete(d.activeSchedule, schedule.ID)
			delete(d.activeScheduleCancelFunc, schedule.ID)
			return errors.Wrap(err, "failed to add cron job")
		}
		d.cronEntries[schedule.ID] = entryID
	}

	return nil
}

func (d *DealPusher) runSchedule(ctx context.Context, schedule *model.Schedule) (model.ScheduleState, error) {
	for {
		var pending sumResult
		err := d.db.WithContext(ctx).Model(&model.Deal{}).
			Where("schedule_id = ? AND state IN (?)", schedule.ID, []model.DealState{
				model.DealProposed, model.DealPublished,
			}).Select("COUNT(*) AS deal_number, SUM(piece_size) AS deal_size").Scan(&pending).Error
		if err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed to count pending deals")
		}
		var total sumResult
		err = d.db.WithContext(ctx).Model(&model.Deal{}).
			Where("schedule_id = ? AND state IN (?)", schedule.ID, []model.DealState{
				model.DealActive, model.DealProposed, model.DealPublished}).Select("COUNT(*) AS deal_number, SUM(piece_size) AS deal_size").Scan(&total).Error
		if err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed to count total active and pending deals")
		}

		var current sumResult

		Logger.Infow("current stats for schedule", "schedule_id", schedule.ID, "pending", pending, "total", total, "current", current)
		for {
			if ctx.Err() != nil {
				//nolint:nilerr
				return "", nil
			}
			var car model.Car
			var dealModel *model.Deal
			var walletObj model.Wallet
			if schedule.MaxPendingDealNumber > 0 && pending.DealNumber >= schedule.MaxPendingDealNumber {
				Logger.Infow("skipping this time since the max pending deal is reached", "schedule_id", schedule.ID)
				goto waitForPending
			}
			if schedule.MaxPendingDealSize > 0 && pending.DealSize >= schedule.MaxPendingDealSize {
				Logger.Infow("skipping this time since the max pending deal size is reached", "schedule_id", schedule.ID)
				goto waitForPending
			}
			if schedule.TotalDealNumber > 0 && total.DealNumber >= schedule.TotalDealNumber {
				Logger.Infow("completing since the total deal number is reached", "schedule_id", schedule.ID)
				return model.ScheduleCompleted, nil
			}
			if schedule.TotalDealSize > 0 && total.DealSize >= schedule.TotalDealSize {
				Logger.Infow("completing since the total deal size is reached", "schedule_id", schedule.ID)
				return model.ScheduleCompleted, nil
			}
			if schedule.ScheduleCron != "" && schedule.ScheduleDealNumber > 0 && current.DealNumber >= schedule.ScheduleDealNumber {
				Logger.Infow("completing this batch since the schedule deal number is reached", "schedule_id", schedule.ID)
				return "", nil
			}
			if schedule.ScheduleCron != "" && schedule.ScheduleDealSize > 0 && current.DealSize >= schedule.ScheduleDealSize {
				Logger.Infow("completing this batch since the schedule deal size is reached", "schedule_id", schedule.ID)
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
				Logger.Infow("no more pieces to send deal", "schedule_id", schedule.ID)
				return model.ScheduleCompleted, nil
			}
			if err != nil {
				return model.ScheduleError, errors.Wrap(err, "failed to find car")
			}

			walletObj, err = d.walletChooser.Choose(ctx, schedule.Dataset.Wallets)
			if err != nil {
				return model.ScheduleError, errors.Wrap(err, "failed to choose wallet")
			}

			err = retry.Do(func() error {
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
				return err
			}, retry.Attempts(d.sendDealRetry), retry.Delay(time.Second),
				retry.DelayType(retry.FixedDelay), retry.Context(ctx))
			if err != nil {
				return "", errors.Wrap(err, "failed to send deal")
			}

			dealModel.ScheduleID = &schedule.ID

			Logger.Debugw("save accepted deal", "deal", dealModel)
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
			continue
		}
	waitForPending:
		select {
		case <-ctx.Done():
			return "", nil
		case <-time.After(waitPendingInterval):
		}
	}
}

func NewDealPusher(db *gorm.DB, lotusURL string,
	lotusToken string, numAttempts uint) (*DealPusher, error) {
	if numAttempts <= 1 {
		numAttempts = 1
	}
	h, err := util.InitHost(nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init host")
	}
	lotusClient := util.NewLotusClient(lotusURL, lotusToken)
	dealMaker := replication.NewDealMaker(lotusClient, h, time.Hour, time.Minute)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init deal maker")
	}
	return &DealPusher{
		db:                       db,
		activeScheduleCancelFunc: make(map[uint32]context.CancelFunc),
		activeSchedule:           make(map[uint32]*model.Schedule),
		cronEntries:              make(map[uint32]cron.EntryID),
		walletChooser:            &replication.DefaultWalletChooser{},
		dealMaker:                dealMaker,
		workerID:                 uuid.New(),
		cron: cron.New(cron.WithLogger(&cronLogger{}), cron.WithLocation(time.UTC),
			cron.WithParser(cron.NewParser(cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor))),
		sendDealRetry: numAttempts,
	}, nil
}

func (d *DealPusher) runOnce(ctx context.Context) {
	var schedules []model.Schedule
	scheduleMap := map[uint32]model.Schedule{}
	Logger.Debugw("getting schedules")
	err := d.db.WithContext(ctx).Preload("Dataset.Wallets").Where("state = ?",
		model.ScheduleActive).Find(&schedules).Error
	if err != nil {
		Logger.Errorw("failed to get schedules", "error", err)
		return
	}
	for _, schedule := range schedules {
		scheduleMap[schedule.ID] = schedule
	}
	// Cancel all jobs that are no longer active
	d.mutex.Lock()
	for id, active := range d.activeSchedule {
		if _, ok := scheduleMap[id]; !ok {
			Logger.Infow("removing inactive schedule", "schedule_id", id)
			d.removeSchedule(*active)
		}
	}
	d.mutex.Unlock()

	for _, schedule := range schedules {
		if d.hasSchedule(schedule.ID) {
			err = d.updateSchedule(ctx, schedule)
			if err != nil {
				Logger.Errorw("failed to update schedule", "error", err)
			}
		} else {
			Logger.Infow("adding new schedule", "schedule_id", schedule.ID)
			err = d.addSchedule(ctx, schedule)
			if err != nil {
				Logger.Errorw("failed to add schedule", "error", err)
			}
		}
	}
}

func (d *DealPusher) Start(ctx context.Context) ([]service.Done, service.Fail, error) {
	getState := func() healthcheck.State {
		return healthcheck.State{
			WorkType: model.DealMaking,
		}
	}

	for {
		alreadyRunning, err := healthcheck.Register(ctx, d.db, d.workerID, getState, false)
		if err == nil && !alreadyRunning {
			break
		}
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to register worker")
		}
		if alreadyRunning {
			Logger.Warnw("another worker already running")
		}
		Logger.Warn("retrying in 1 minute")
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		case <-time.After(time.Minute):
		}
	}

	healthcheckDone := make(chan struct{})
	go func() {
		defer close(healthcheckDone)
		healthcheck.StartReportHealth(ctx, d.db, d.workerID, getState)
	}()

	runDone := make(chan struct{})
	go func() {
		defer close(runDone)
		d.cron.Start()
		for {
			d.runOnce(ctx)
			Logger.Debug("waiting for deal schedule check in 15 secs")
			select {
			case <-ctx.Done():
				return
			case <-time.After(15 * time.Second):
			}
		}
	}()

	fail := make(chan error)

	cleanupDone := make(chan struct{})
	go func() {
		defer close(cleanupDone)
		<-ctx.Done()
		err := d.cleanup()
		if err != nil {
			Logger.Errorw("failed to cleanup", "error", err)
		}
	}()

	return []service.Done{runDone, cleanupDone, healthcheckDone}, fail, nil
}

func (d *DealPusher) cleanup() error {
	d.cron.Stop()
	return database.DoRetry(func() error {
		return d.db.Where("id = ?", d.workerID).Delete(&model.Worker{}).Error
	})
}
