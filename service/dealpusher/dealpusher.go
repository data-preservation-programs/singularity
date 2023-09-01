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

// DealPusher represents a struct that encapsulates the data and functionality related to pushing deals in a replication process.
type DealPusher struct {
	dbNoContext              *gorm.DB                      // Pointer to a gorm.DB object representing a database connection.
	walletChooser            replication.WalletChooser     // Object responsible for choosing a wallet for replication.
	dealMaker                replication.DealMaker         // Object responsible for making a deal in replication.
	workerID                 uuid.UUID                     // UUID identifying the associated worker.
	activeSchedule           map[uint32]*model.Schedule    // Map storing active schedules with schedule IDs as keys and pointers to model.Schedule objects as values.
	activeScheduleCancelFunc map[uint32]context.CancelFunc // Map storing cancel functions for active schedules with schedule IDs as keys and CancelFunc as values.
	cronEntries              map[uint32]cron.EntryID       // Map storing cron entries for the DealPusher with schedule IDs as keys and EntryID as values.
	cron                     *cron.Cron                    // Job scheduler for scheduling tasks at specified intervals.
	mutex                    sync.Mutex                    // Mutex for providing mutual exclusion to protect shared resources.
	sendDealAttempts         uint                          // Number of attempts for sending a deal.
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

// runScheduleAndUpdateState is a method of the DealPusher type.
// It runs the specified Schedule, assesses the outcome, and updates the Schedule's state
// accordingly in the database. If errors are encountered during the run, they are logged
// and potentially saved to the Schedule's record in the database, depending on the Schedule's Cron setting.
//
// The steps it takes are as follows:
// 1. Runs the Schedule using the runSchedule method, which attempts to make deals based on the Schedule's configuration.
// 2. If runSchedule returns an error, logs the error and saves it to the Schedule's record if ScheduleCron is not set.
// 3. If runSchedule returns a non-empty state (either ScheduleCompleted or ScheduleError), updates the Schedule's state in the database.
// 4. Logs the Schedule's completion or error state, if applicable, and removes the Schedule from the DealPusher's active schedules.
//
// Parameters:
//
//	ctx:      The context for managing the lifecycle of this Schedule run.
//	          If the context is Done, the function exits cleanly.
//	schedule: A pointer to the Schedule that this function is processing.
//
// This function does not return any values but updates the Schedule's state in the database
// based on the actions performed in the runSchedule function. It also handles errors and logs relevant information.
//
// Note: This function is designed to act as a controller that runs a Schedule,
// handles the outcome, updates the Schedule's state, and logs the results.
func (d *DealPusher) runScheduleAndUpdateState(ctx context.Context, schedule *model.Schedule) {
	db := d.dbNoContext.WithContext(ctx)
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
		err = db.Model(schedule).Updates(updates).Error
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

// runSchedule is a method of the DealPusher type. It processes a single Schedule,
// and continuously attempts to make deals based on the information and constraints specified in the Schedule.
//
// The steps it takes in each iteration are as follows:
// 1. Counts the number and size of pending and total active deals for the current schedule from the database.
// 2. Checks various conditions defined in the Schedule to decide whether to proceed with making a new deal.
// 3. Finds a car (Content Addressed Archive) that has not been sent to the provider for a deal.
// 4. Chooses a wallet from the dataset’s associated wallets.
// 5. Makes a deal using the details from the car and wallet, and the deal parameters defined in the Schedule.
// 6. Saves the newly created deal to the database.
// 7. Updates the counts of pending, total, and current deals based on the new deal.
// 8. If the context is done, returns immediately, otherwise waits for a specified interval before the next iteration.
//
// Parameters:
//
//	ctx:      The context for managing the lifecycle of this Schedule run.
//	          If the context is Done, the function exits cleanly.
//	schedule: A pointer to the Schedule that this function is processing.
//
// Returns:
//  1. A ScheduleState which represents the new state of the Schedule
//     based on the actions performed in this run.
//     Possible values: ScheduleCompleted, ScheduleError, or an empty string.
//  2. An error if any step of the process encounters an issue, otherwise nil.
func (d *DealPusher) runSchedule(ctx context.Context, schedule *model.Schedule) (model.ScheduleState, error) {
	db := d.dbNoContext.WithContext(ctx)
	for {
		var pending sumResult
		err := db.Model(&model.Deal{}).
			Where("schedule_id = ? AND state IN (?)", schedule.ID, []model.DealState{
				model.DealProposed, model.DealPublished,
			}).Select("COUNT(*) AS deal_number, SUM(piece_size) AS deal_size").Scan(&pending).Error
		if err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed to count pending deals")
		}
		var total sumResult
		err = db.Model(&model.Deal{}).
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

			err = db.Where("dataset_id = ? AND piece_cid NOT IN (?)",
				schedule.DatasetID,
				db.Table("deals").Select("piece_cid").
					Where("provider = ? AND state IN (?)",
						schedule.Provider,
						[]model.DealState{
							model.DealProposed, model.DealPublished, model.DealActive,
						})).First(&car).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				Logger.Infow("no more pieces to send deal", "schedule_id", schedule.ID)
				// we're out of deals to schedule, but if we're running a perpetual cron, we simply put things on hold till next cron
				if schedule.ScheduleCron != "" && schedule.ScheduleCronPerpetual {
					return "", nil
				}
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
			}, retry.Attempts(d.sendDealAttempts), retry.Delay(time.Second),
				retry.DelayType(retry.FixedDelay), retry.Context(ctx))
			if err != nil {
				return "", errors.Wrap(err, "failed to send deal")
			}

			dealModel.ScheduleID = &schedule.ID

			Logger.Debugw("save accepted deal", "deal", dealModel)
			err = database.DoRetry(ctx, func() error { return db.Create(dealModel).Error })
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
		dbNoContext:              db,
		activeScheduleCancelFunc: make(map[uint32]context.CancelFunc),
		activeSchedule:           make(map[uint32]*model.Schedule),
		cronEntries:              make(map[uint32]cron.EntryID),
		walletChooser:            &replication.RandomWalletChooser{},
		dealMaker:                dealMaker,
		workerID:                 uuid.New(),
		cron: cron.New(cron.WithLogger(&cronLogger{}), cron.WithLocation(time.UTC),
			cron.WithParser(cron.NewParser(cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor))),
		sendDealAttempts: numAttempts,
	}, nil
}

// runOnce is a method of the DealPusher type that runs a single iteration of the deal pushing logic.
//
// In each iteration, the method performs the following actions:
//  1. Fetches all the active schedules from the database.
//  2. Constructs a map of these schedules for quick lookup.
//  3. Cancels all the jobs in the DealPusher that are no longer active (based on the latest fetched schedules).
//  4. For each schedule in the fetched active schedules:
//     a. If the schedule is already being processed, it updates that schedule's processing logic.
//     b. If the schedule is new, it starts processing that schedule.
//
// Parameters:
//
//	ctx : The context for managing the lifecycle of this iteration. If Done, the function exits cleanly.
//
// This function is designed to be idempotent, meaning it can be run multiple times with the same effect.
// It is called repeatedly by the main deal processing loop in DealPusher.Start.
//
// Note: Errors encountered during this process are logged but do not stop the function's execution.
func (d *DealPusher) runOnce(ctx context.Context) {
	var schedules []model.Schedule
	scheduleMap := map[uint32]model.Schedule{}
	Logger.Debugw("getting schedules")
	db := d.dbNoContext.WithContext(ctx)
	err := db.Preload("Dataset.Wallets").Where("state = ?",
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
			d.removeScheduleUnsafe(*active)
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

// Start initializes and starts the DealPusher service.
//
// It first attempts to register the worker with the health check system.
// If another worker is already running, it waits and retries until it can register or the context is cancelled.
// Once registered, it launches three main activities in separate goroutines:
// 1. Reporting its health status.
// 2. Running the deal processing loop.
// 3. Handling cleanup when the service is stopped.
//
// Parameters:
//
//	ctx : The context for managing the lifecycle of the Start function. If Done, the function exits cleanly.
//
// Returns:
//   - A slice of channels that the caller can select on to wait for the service to stop.
//   - A channel for errors that may occur while the service is running.
//   - An error if there was a problem starting the service.
//
// This function is intended to be called once at the start of the service lifecycle.
func (d *DealPusher) Start(ctx context.Context) ([]service.Done, service.Fail, error) {
	getState := func() healthcheck.State {
		return healthcheck.State{
			WorkType: model.DealMaking,
		}
	}

	for {
		alreadyRunning, err := healthcheck.Register(ctx, d.dbNoContext, d.workerID, getState, false)
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
		healthcheck.StartReportHealth(ctx, d.dbNoContext, d.workerID, getState)
		Logger.Info("healthcheck stopped")
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
				Logger.Info("cron stopped")
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
		ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		//nolint:contextcheck
		err := d.cleanup(ctx2)
		if err != nil {
			Logger.Errorw("failed to cleanup", "error", err)
		} else {
			Logger.Info("cleanup done")
		}
	}()

	return []service.Done{runDone, cleanupDone, healthcheckDone}, fail, nil
}

func (d *DealPusher) cleanup(ctx context.Context) error {
	d.cron.Stop()
	return database.DoRetry(ctx, func() error {
		return d.dbNoContext.WithContext(ctx).Where("id = ?", d.workerID).Delete(&model.Worker{}).Error
	})
}
