// Package pdptracker tracks PDP (Proof of Data Possession) deals on Filecoin
// using Shovel-based event indexing instead of linear chain state scanning.
package pdptracker

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/healthcheck"
	"github.com/google/uuid"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

var ErrAlreadyRunning = errors.New("another PDP tracker worker already running")

const (
	healthRegisterRetryInterval = time.Minute
	cleanupTimeout              = 5 * time.Second
)

var Logger = log.Logger("pdptracker")

// PDPTracker reads events from Shovel integration tables and materializes
// them into singularity's deal model. It replaces the previous approach of
// linearly scanning all proof sets via RPC every cycle.
type PDPTracker struct {
	workerID    uuid.UUID
	dbNoContext *gorm.DB
	config      PDPConfig
	rpcClient   *ChainPDPClient
	once        bool
}

// NewPDPTracker creates a new event-driven PDP deal tracker.
func NewPDPTracker(
	db *gorm.DB,
	config PDPConfig,
	rpcClient *ChainPDPClient,
	once bool,
) PDPTracker {
	return PDPTracker{
		workerID:    uuid.New(),
		dbNoContext: db,
		config:      config,
		rpcClient:   rpcClient,
		once:        once,
	}
}

func (*PDPTracker) Name() string {
	return "PDPTracker"
}

// Start begins the PDP tracker service.
func (p *PDPTracker) Start(ctx context.Context, exitErr chan<- error) error {
	Logger.Infow("PDP tracker starting", "pollInterval", p.config.PollingInterval)

	var regTimer *time.Timer
	for {
		alreadyRunning, err := healthcheck.Register(ctx, p.dbNoContext, p.workerID, model.PDPTracker, false)
		if err != nil {
			return errors.WithStack(err)
		}
		if !alreadyRunning {
			break
		}

		Logger.Warnw("another PDP tracker worker already running")
		if p.once {
			return ErrAlreadyRunning
		}
		Logger.Warn("retrying in 1 minute")
		if regTimer == nil {
			regTimer = time.NewTimer(healthRegisterRetryInterval)
			defer regTimer.Stop()
		} else {
			regTimer.Reset(healthRegisterRetryInterval)
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-regTimer.C:
		}
	}

	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)

	healthcheckDone := make(chan struct{})
	go func() {
		defer close(healthcheckDone)
		healthcheck.StartReportHealth(ctx, p.dbNoContext, p.workerID, model.PDPTracker)
		Logger.Info("PDP tracker health report stopped")
	}()

	go func() {
		var timer *time.Timer
		var runErr error
		for {
			runErr = p.runOnce(ctx)
			if runErr != nil {
				if ctx.Err() != nil {
					if errors.Is(runErr, context.Canceled) {
						runErr = nil
					}
					Logger.Info("PDP tracker run stopped")
					break
				}
				Logger.Errorw("failed to run PDP tracker once", "error", runErr)
			}
			if p.once {
				Logger.Info("PDP tracker run once done")
				break
			}
			if timer == nil {
				timer = time.NewTimer(p.config.PollingInterval)
				defer timer.Stop()
			} else {
				timer.Reset(p.config.PollingInterval)
			}

			var stopped bool
			select {
			case <-ctx.Done():
				stopped = true
			case <-timer.C:
			}
			if stopped {
				Logger.Info("PDP tracker run stopped")
				break
			}
		}

		cancel()

		ctx2, cancel2 := context.WithTimeout(context.Background(), cleanupTimeout)
		defer cancel2()
		//nolint:contextcheck
		err := p.cleanup(ctx2)
		if err != nil {
			Logger.Errorw("failed to cleanup PDP tracker", "error", err)
		} else {
			Logger.Info("PDP tracker cleanup done")
		}

		<-healthcheckDone

		if exitErr != nil {
			exitErr <- runErr
		}
		Logger.Info("PDP tracker stopped")
	}()

	return nil
}

func (p *PDPTracker) cleanup(ctx context.Context) error {
	return database.DoRetry(ctx, func() error {
		return p.dbNoContext.WithContext(ctx).Where("id = ?", p.workerID).Delete(&model.Worker{}).Error
	})
}

// runOnce drains the Shovel event inbox and materializes state changes.
func (p *PDPTracker) runOnce(ctx context.Context) error {
	db := p.dbNoContext.WithContext(ctx)
	return processNewEvents(ctx, db, p.rpcClient)
}
