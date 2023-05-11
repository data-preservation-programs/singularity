package replication

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/data-preservation-programs/go-singularity/util"
	"github.com/google/uuid"
	"github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Protocol string

const (
	StorageMarket111 = "/fil/storage/mk/1.1.1"
	StorageMarket120 = "/fil/storage/mk/1.2.0"
)

type Worker struct {
	id        string
	db        *gorm.DB
	logger    *log.ZapEventLogger
	dealMaker *DealMaker
}

type WorkerThread struct {
	db            *gorm.DB
	logger        *zap.SugaredLogger
	scheduleJob   model.Schedule
	cancel        context.CancelFunc
	cronJob       *cron.Cron
	dealNumber    uint64
	dealSize      uint64
	dealMaker     *DealMaker
	walletChooser WalletChooser
}

func NewWorkerThread(db *gorm.DB, scheduleJob model.Schedule, cancel context.CancelFunc, dealMaker *DealMaker) *WorkerThread {
	return &WorkerThread{
		db:            db,
		logger:        log.Logger("replication-worker").With("schedule_id", scheduleJob.ID),
		scheduleJob:   scheduleJob,
		cancel:        cancel,
		dealMaker:     dealMaker,
		walletChooser: WalletChooser{},
	}
}

func (w *WorkerThread) Cancel() {
	if w.cronJob != nil {
		w.cronJob.Stop()
	}
	w.cancel()
}

func (w *WorkerThread) Run(ctx context.Context) {
	if w.scheduleJob.SchedulePattern == "" {
		w.cronJob = cron.New()
		_, err := w.cronJob.AddFunc(w.scheduleJob.SchedulePattern, func() {
			w.RunBatch(ctx)
			if w.dealNumber >= w.scheduleJob.TotalDealNumber || w.dealSize >= w.scheduleJob.TotalDealSize {
				w.logger.Infow("stopping cron since the total deal number is reached")
				w.Cancel()
			}
		})
		if err != nil {
			w.logger.Errorw("failed to add cron job", "error", err)
			return
		}
	} else {
		w.logger.Infow("no schedule pattern, running once")
		w.RunBatch(ctx)
	}
}

func (w *WorkerThread) RunBatch(ctx context.Context) {
	numErrors := 0
	dealNumber := uint64(0)
	dealSize := uint64(0)
	pendingDealNumber := int64(0)
	pendingDealSize := uint64(0)
	err := w.db.Model(&model.Deal{}).
		Where("schedule_id = ? AND state IN (?)", w.scheduleJob.ID, []model.DealState{
			model.DealProposed, model.DealPublished,
		}).Count(&pendingDealNumber).Error
	if err != nil {
		w.logger.Errorw("failed to count pending deals", "error", err)
		return
	}

	err = w.db.Model(&model.Deal{}).
		Where("schedule_id = ? AND state IN (?)", w.scheduleJob.ID, []model.DealState{
			model.DealProposed, model.DealPublished,
		}).Select("SUM(piece_cid)").Scan(&pendingDealSize).Error
	if err != nil {
		w.logger.Errorw("failed to count pending deals", "error", err)
		return
	}

	if uint64(pendingDealNumber) >= w.scheduleJob.MaxPendingDealNumber || pendingDealSize >= w.scheduleJob.MaxPendingDealSize {
		w.logger.Infow("too many pending deals")
		return
	}

	for {
		if numErrors > 5 {
			w.logger.Errorw("too many errors, exiting")
			return
		}

		if w.dealNumber >= w.scheduleJob.TotalDealNumber || w.dealSize >= w.scheduleJob.TotalDealSize {
			w.logger.Infow("finished making deals")
			err := w.db.Model(&w.scheduleJob).Update("state", model.ScheduleCompleted).Error
			if err != nil {
				w.logger.Errorw("failed to update schedule state", "err", err)
			}
			return
		}

		if w.scheduleJob.SchedulePattern != "" &&
			dealNumber >= w.scheduleJob.ScheduleDealNumber ||
			dealSize >= w.scheduleJob.TotalDealNumber {
			w.logger.Infow("finished making deals for this batch")
			return
		}

		if uint64(pendingDealNumber) >= w.scheduleJob.MaxPendingDealNumber || pendingDealSize >= w.scheduleJob.MaxPendingDealSize {
			w.logger.Infow("too many pending deals")
			return
		}

		car := model.Car{}
		// Get a piece that has not been proposed
		err = w.db.Where("dataset_id = ? AND piece_cid NOT IN (?)",
			w.scheduleJob.DatasetID,
			w.db.Table("deals").Select("piece_cid").
				Where("provider = ? AND state NOT IN (?)",
					w.scheduleJob.Provider,
					[]model.DealState{
						model.DealProposed, model.DealPublished, model.DealActive,
					})).First(&car).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.logger.Infow("no more pieces to send deal")
			return
		}

		if err != nil {
			w.logger.Errorw("failed to get another piece to send deal", "error", err)
			return
		}

		providerInfo, err := w.dealMaker.GetProviderInfo(ctx, w.scheduleJob.Provider)
		if err != nil {
			w.logger.Errorw("failed to get provider info", "error", err)
			return
		}

		walletObj := w.walletChooser.Choose(ctx, w.scheduleJob.Dataset.Wallets)
		now := time.Now().UTC()
		proposalID, err := w.dealMaker.MakeDeal(ctx, now, walletObj, car, w.scheduleJob, peer.AddrInfo{
			ID:    providerInfo.PeerID,
			Addrs: providerInfo.Multiaddrs,
		})
		if err != nil {
			w.logger.Errorw("failed to make deal", "error", err)
			numErrors++
			continue
		}

		err = w.db.Create(&model.Deal{
			State:         model.DealProposed,
			ClientID:      walletObj.ID,
			ClientAddress: walletObj.ID,
			ProposalID:    proposalID,
			Label:         car.RootCID,
			PieceCID:      car.PieceCID,
			PieceSize:     car.PieceSize,
			Start:         now.Add(w.scheduleJob.StartDelay),
			Duration:      w.scheduleJob.Duration,
			End:           now.Add(w.scheduleJob.StartDelay + w.scheduleJob.Duration),
			Price:         w.scheduleJob.Price,
			Verified:      w.scheduleJob.Verified,
			ScheduleID:    &w.scheduleJob.ID,
		}).Error
		if err != nil {
			w.logger.Errorw("failed to create deal", "error", err)
			return
		}
	}
}

func NewWorker(db *gorm.DB, lotusURL string, lotusToken string) (*Worker, error) {
	libp2p, err := util.InitHost(context.Background(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init libp2p host")
	}
	dealMaker, err := NewDealMaker(lotusURL, lotusToken, libp2p)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create deal maker")
	}
	return &Worker{
		id:        uuid.NewString(),
		db:        db,
		logger:    log.Logger("replication-worker"),
		dealMaker: dealMaker,
	}, nil
}

func (w *Worker) Cleanup() error {
	return w.db.Where("id = ?", w.id).Delete(&model.Worker{}).Error
}

func (w *Worker) Run(parent context.Context) error {
	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	w.db = w.db.WithContext(ctx)
	go w.run(ctx)

	<-signalChan
	w.logger.Info("received signal, cleaning up")
	err := w.Cleanup()
	if err != nil {
		w.logger.Errorw("failed to cleanup", "error", err)
	}
	return cli.Exit("received signal", 130)
}

func (w *Worker) run(ctx context.Context) {
	host, err := util.InitHost(ctx, nil)
	if err != nil {
		panic(err)
	}
	running := map[uint32]*WorkerThread{}
	dealMaker, err := NewDealMaker("", "", host)
	if err != nil {
		panic(err)
	}
	for {
		schedules := make([]model.Schedule, 0)
		err := w.db.Where("state = ?", model.ScheduleActive).Find(&schedules)
		if err != nil {
			w.logger.Error("failed to fetch schedules", err)
			time.Sleep(time.Minute)
			continue
		}

		// convert to map for easier lookup
		scheduleMap := map[uint32]model.Schedule{}
		for _, schedule := range schedules {
			scheduleMap[schedule.ID] = schedule
		}

		// For schedules that no longer exists, cancel the worker thread
		for id, worker := range running {
			if _, ok := scheduleMap[id]; !ok {
				worker.Cancel()
				delete(running, id)
			}
		}

		// For schedules that has changed, cancel the worker thread and start a new one
		for id, schedule := range scheduleMap {
			if worker, ok := running[id]; ok {
				if !worker.scheduleJob.Equal(schedule) {
					worker.Cancel()
					delete(running, id)
					workerCtx, cancel := context.WithCancel(ctx)
					if schedule.State == model.ScheduleActive {
						worker = NewWorkerThread(w.db, schedule, cancel, dealMaker)
						running[id] = worker
						go worker.Run(workerCtx)
					}
				}
			} else if schedule.State == model.ScheduleActive {
				workerCtx, cancel := context.WithCancel(ctx)
				worker = NewWorkerThread(w.db, schedule, cancel, dealMaker)
				running[id] = worker
				go worker.Run(workerCtx)
			}
		}
	}
}
