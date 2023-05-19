package service

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/data-preservation-programs/go-singularity/replication"
	"github.com/data-preservation-programs/go-singularity/util"
	"github.com/google/uuid"
	"github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type DealMakerService struct {
	db            *gorm.DB
	logger        *log.ZapEventLogger
	jobs          map[uint32]context.CancelFunc
	walletChooser *replication.WalletChooser
	dealMaker     *replication.DealMaker
	workerID      uuid.UUID
}

type DealMakerWorker struct {
	db            *gorm.DB
	logger        *zap.SugaredLogger
	dealMaker     *replication.DealMaker
	walletChooser *replication.WalletChooser
	workerID      uuid.UUID
}

type sumResult struct {
	DealNumber int
	DealSize   int64
}

func NewDealMakerWorker(db *gorm.DB,
	dealMaker *replication.DealMaker,
	walletChooser *replication.WalletChooser,
	workerID uuid.UUID) *DealMakerWorker {
	return &DealMakerWorker{
		db:            db,
		dealMaker:     dealMaker,
		walletChooser: walletChooser,
		logger:        log.Logger("dealmaker").With("worker_id", uuid.NewString()),
		workerID:      workerID,
	}
}

func (w *DealMakerWorker) Run(ctx context.Context, schedule model.Schedule) {
	// First, set the worker ID if it's not set
	result := w.db.Model(&model.Schedule{}).Where("id = ? AND schedule_worker_id IS NULL", schedule.ID).Update("schedule_worker_id", w.workerID.String())
	if result.Error != nil {
		w.logger.Errorw("failed to set worker ID", "error", result.Error)
		return
	}
	if result.RowsAffected == 0 {
		w.logger.Warnw("schedule already has a worker", "schedule_id", schedule.ID)
		return
	}

	done, err := w.runOnce(ctx, schedule)
	var updates map[string]interface{}
	switch {
	case err == nil && done:
		updates = map[string]interface{}{"state": model.ScheduleCompleted, "last_processed_timestamp": time.Now(), "schedule_worker_id": nil, "error_message": ""}
	case err == nil && !done:
		updates = map[string]interface{}{"last_processed_timestamp": time.Now(), "schedule_worker_id": nil, "error_message": ""}
	case err != nil && schedule.ScheduleIntervalSeconds > 0:
		updates = map[string]interface{}{"error_message": err.Error(), "last_processed_timestamp": time.Now(), "schedule_worker_id": nil}
	case err != nil && schedule.ScheduleIntervalSeconds == 0:
		updates = map[string]interface{}{"state": model.SchedulePaused, "error_message": err.Error(), "last_processed_timestamp": time.Now(), "schedule_worker_id": nil}
	}
	err = w.db.WithContext(ctx).Model(&model.Schedule{}).Where("id = ?", schedule.ID).Updates(updates).Error
	if err != nil {
		w.logger.Errorw("failed to update schedule", "schedule_id", schedule.ID, "error", err)
	}
	return
}

func (w *DealMakerWorker) runOnce(ctx context.Context, schedule model.Schedule) (bool, error) {
	if schedule.ScheduleDealNumber == 0 {
		schedule.ScheduleDealNumber = math.MaxInt
	}
	if schedule.ScheduleDealSize == 0 {
		schedule.ScheduleDealSize = math.MaxInt64
	}
	if schedule.TotalDealNumber == 0 {
		schedule.TotalDealNumber = math.MaxInt
	}
	if schedule.TotalDealSize == 0 {
		schedule.TotalDealSize = math.MaxInt64
	}
	if schedule.MaxPendingDealNumber == 0 {
		schedule.MaxPendingDealNumber = math.MaxInt
	}
	if schedule.MaxPendingDealSize == 0 {
		schedule.MaxPendingDealSize = math.MaxInt64
	}

	var pendingResult sumResult
	err := w.db.WithContext(ctx).Model(&model.Deal{}).
		Where("schedule_id = ? AND state IN (?)", schedule.ID, []model.DealState{
			model.DealProposed, model.DealPublished,
		}).Select("COUNT(*) AS deal_number, SUM(piece_cid) AS deal_size").Scan(&pendingResult).Error
	if err != nil {
		return false, errors.Wrap(err, "failed to count pending deals")
	}
	if pendingResult.DealNumber >= schedule.MaxPendingDealNumber || pendingResult.DealSize >= schedule.MaxPendingDealSize {
		w.logger.Infow("stopping since the max pending deal number is reached", "schedule_id", schedule.ID)
		return false, nil
	}

	var totalActiveResult sumResult
	err = w.db.WithContext(ctx).Model(&model.Deal{}).
		Where("schedule_id = ? AND state IN (?)", schedule.ID, []model.DealState{
			model.DealActive}).Select("COUNT(*) AS deal_number, SUM(piece_cid) AS deal_size").Scan(&totalActiveResult).Error
	if err != nil {
		return false, errors.Wrap(err, "failed to count total active deals")
	}
	if totalActiveResult.DealNumber >= schedule.TotalDealNumber || totalActiveResult.DealSize >= schedule.TotalDealSize {
		w.logger.Infow("completing since the total deal number is reached", "schedule_id", schedule.ID)
		return true, nil
	}

	var totalResult sumResult
	err = w.db.WithContext(ctx).Model(&model.Deal{}).
		Where("schedule_id = ? AND state IN (?)", schedule.ID, []model.DealState{
			model.DealProposed, model.DealPublished, model.DealActive}).
		Select("COUNT(*) AS deal_number, SUM(piece_cid) AS deal_size").Scan(&totalResult).Error
	if err != nil {
		return false, errors.Wrap(err, "failed to count total deals")
	}
	if totalResult.DealNumber >= schedule.TotalDealNumber || totalResult.DealSize >= schedule.TotalDealSize {
		w.logger.Infow("waiting for some deals to become active before sending more deals", "schedule_id", schedule.ID)
		return false, nil
	}

	number := schedule.ScheduleDealNumber
	size := schedule.ScheduleDealSize
	if schedule.ScheduleIntervalSeconds == 0 {
		number = schedule.TotalDealNumber
		size = schedule.TotalDealSize
	}

	for number > 0 && size > 0 &&
		totalResult.DealNumber < schedule.TotalDealNumber &&
		totalResult.DealSize < schedule.TotalDealSize {
		var car model.Car
		err = w.db.WithContext(ctx).Where("dataset_id = ? AND piece_cid NOT IN (?)",
			schedule.DatasetID,
			w.db.Table("deals").Select("piece_cid").
				Where("provider = ? AND state IN (?)",
					schedule.Provider,
					[]model.DealState{
						model.DealProposed, model.DealPublished, model.DealActive,
					})).First(&car).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.logger.Infow("no more pieces to send deal")
			return true, nil
		}
		if err != nil {
			return false, errors.Wrap(err, "failed to get piece")
		}
		number -= 1
		size -= car.PieceSize
		totalResult.DealSize += car.PieceSize
		totalResult.DealNumber += 1

		providerInfo, err := w.dealMaker.GetProviderInfo(ctx, schedule.Provider)
		if err != nil {
			return false, errors.Wrap(err, "failed to get provider info")
		}

		walletObj := w.walletChooser.Choose(ctx, schedule.Dataset.Wallets)
		now := time.Now().UTC()
		proposalID, err := w.dealMaker.MakeDeal(ctx, now, walletObj, car, schedule, peer.AddrInfo{
			ID:    providerInfo.PeerID,
			Addrs: providerInfo.Multiaddrs,
		})
		if err != nil {
			w.logger.Errorw("failed to make deal", "error", err)
			continue
		}

		err = w.db.Create(&model.Deal{
			State:      model.DealProposed,
			DatasetID:  &schedule.DatasetID,
			ClientID:   walletObj.ID,
			Provider:   schedule.Provider,
			ProposalID: proposalID,
			Label:      car.RootCID,
			PieceCID:   car.PieceCID,
			PieceSize:  car.PieceSize,
			Start:      now.Add(schedule.StartDelay),
			Duration:   schedule.Duration,
			End:        now.Add(schedule.StartDelay + schedule.Duration),
			Price:      schedule.Price,
			Verified:   schedule.Verified,
			ScheduleID: &schedule.ID,
		}).Error
		if err != nil {
			return false, errors.Wrap(err, "failed to create deal")
		}
	}
	return true, nil
}

func NewDealMakerService(db *gorm.DB, lotusURL string, lotusToken string) (*DealMakerService, error) {
	h, err := util.InitHost(context.Background(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init host")
	}
	dealMaker, err := replication.NewDealMaker(lotusURL, lotusToken, h)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init deal maker")
	}
	return &DealMakerService{
		db:            db,
		logger:        log.Logger("deal-maker"),
		jobs:          make(map[uint32]context.CancelFunc),
		walletChooser: &replication.WalletChooser{},
		dealMaker:     dealMaker,
		workerID:      uuid.New(),
	}, nil
}

func (d DealMakerService) Run(parent context.Context) error {
	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	getState := func() State {
		return State{
			WorkType:  model.DealMaking,
			WorkingOn: "deals",
		}
	}
	HealthCheck(d.db, d.workerID, getState)
	go StartHealthCheck(ctx, d.db, d.workerID, getState)
	go StartHealthCheckCleanup(ctx, d.db)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGTRAP, os.Kill, os.Interrupt)

	for {
		var schedules []model.Schedule
		err := d.db.WithContext(ctx).Preload("Dataset.Wallets").Where("state = ? AND (schedule_interval_seconds = 0 OR (schedule_interval_seconds + last_processed_timestamp > ?))",
			model.ScheduleActive, time.Now()).Find(&schedules).Error
		if err != nil {
			d.logger.Errorw("failed to get schedules", "error", err)
			goto nextloop
		}
		// Cancel all jobs that are no longer active
		for id, cancel := range d.jobs {
			active := slices.ContainsFunc(schedules, func(i model.Schedule) bool {
				return i.ID == id && (i.ScheduleWorkerID != nil && *i.ScheduleWorkerID == d.workerID.String())
			})
			if !active {
				cancel()
				delete(d.jobs, id)
			}
		}

		// Kick off new jobs
		for _, schedule := range schedules {
			if _, ok := d.jobs[schedule.ID]; ok {
				continue
			}
			ctx, cancel := context.WithCancel(ctx)
			d.jobs[schedule.ID] = cancel
			go func(schedule model.Schedule) {
				worker := NewDealMakerWorker(d.db, d.dealMaker, d.walletChooser, d.workerID)
				worker.Run(ctx, schedule)
			}(schedule)
		}
	nextloop:
		select {
		case <-signalChan:
			d.logger.Infow("received signal, stopping")
			for _, cancel := range d.jobs {
				cancel()
			}
			d.cleanup()
			return cli.Exit("received signal", 130)
		case <-ctx.Done():
			d.cleanup()
			return nil
		case <-time.After(1 * time.Minute):
		}
	}
}

func (d DealMakerService) cleanup() error {
	return d.db.Where("id = ?", d.workerID).Delete(&model.Worker{}).Error
}
