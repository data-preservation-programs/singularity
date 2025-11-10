// Package pdptracker provides a service for tracking PDP (Proof of Data Possession) deals
// using the f41 actor on Filecoin. This is distinct from legacy f05 market deals.
//
// PDP deals use proof sets managed through the PDPVerifier contract, where data is verified
// through cryptographic challenges rather than the traditional sector sealing process.
package pdptracker

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/healthcheck"
	"github.com/filecoin-project/go-address"
	"github.com/google/uuid"
	"github.com/gotidy/ptr"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

var ErrAlreadyRunning = errors.New("another PDP tracker worker already running")

const (
	healthRegisterRetryInterval = time.Minute
	cleanupTimeout              = 5 * time.Second
)

var Logger = log.Logger("pdptracker")

// ProofSetInfo contains information about a PDP proof set retrieved from on-chain state
type ProofSetInfo struct {
	ProofSetID         uint64
	ClientAddress      address.Address // f4 address of the client
	ProviderAddress    address.Address // Provider/record keeper address
	IsLive             bool            // Whether the proof set is actively being challenged
	NextChallengeEpoch int32           // Next epoch when a challenge is due
	PieceCIDs          []cid.Cid
}

// PDPClient is the interface for interacting with PDP on-chain state.
type PDPClient interface {
	// GetProofSetsForClient returns all proof sets associated with a client address
	GetProofSetsForClient(ctx context.Context, clientAddress address.Address) ([]ProofSetInfo, error)
	// GetProofSetInfo returns detailed information about a specific proof set
	GetProofSetInfo(ctx context.Context, proofSetID uint64) (*ProofSetInfo, error)
	// IsProofSetLive checks if a proof set is actively being challenged
	IsProofSetLive(ctx context.Context, proofSetID uint64) (bool, error)
	// GetNextChallengeEpoch returns the next challenge epoch for a proof set
	GetNextChallengeEpoch(ctx context.Context, proofSetID uint64) (int32, error)
}

// PDPBulkClient is an optional optimization interface for fetching all proof sets in one call.
type PDPBulkClient interface {
	GetProofSets(ctx context.Context) ([]ProofSetInfo, error)
}

// PDPTracker tracks PDP deals (f41 actor) on the Filecoin network.
// It monitors proof sets and updates deal status based on on-chain state.
type PDPTracker struct {
	workerID    uuid.UUID
	dbNoContext *gorm.DB
	config      PDPConfig
	pdpClient   PDPClient
	once        bool
}

// NewPDPTracker creates a new PDP deal tracker.
//
// Parameters:
//   - db: Database connection for storing deal information
//   - config: Tracker runtime configuration
//   - pdpClient: Client for interacting with PDP contracts
//   - once: If true, run only once instead of continuously
func NewPDPTracker(
	db *gorm.DB,
	config PDPConfig,
	pdpClient PDPClient,
	once bool,
) PDPTracker {
	return PDPTracker{
		workerID:    uuid.New(),
		dbNoContext: db,
		config:      config,
		pdpClient:   pdpClient,
		once:        once,
	}
}

func (*PDPTracker) Name() string {
	return "PDPTracker"
}

// Start begins the PDP tracker service.
func (p *PDPTracker) Start(ctx context.Context, exitErr chan<- error) error {
	Logger.Infow("PDP tracker started",
		"pollInterval", p.config.PollingInterval,
	)
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

// runOnce performs a single cycle of PDP deal tracking.
// It queries actors, fetches their PDP proof sets, and updates deal status.
func (p *PDPTracker) runOnce(ctx context.Context) error {
	db := p.dbNoContext.WithContext(ctx)

	var actors []model.Actor
	err := db.Find(&actors).Error
	if err != nil {
		return errors.Wrap(err, "failed to get actors from database")
	}

	now := time.Now()
	var updated, inserted int64
	trackedProofSets := make(map[uint64]struct{})

	processProofSet := func(actor model.Actor, ps ProofSetInfo) {
		for _, pieceCID := range ps.PieceCIDs {
			if pieceCID == cid.Undef {
				Logger.Warnw("invalid piece CID from PDP proof set", "pieceCID", pieceCID.String(), "proofSetID", ps.ProofSetID)
				continue
			}
			modelPieceCID := model.CID(pieceCID)

			var existingDeal model.Deal
			err := db.Where("proof_set_id = ? AND piece_cid = ? AND deal_type = ?",
				ps.ProofSetID, modelPieceCID, model.DealTypePDP).First(&existingDeal).Error

			if err == nil {
				updates := map[string]any{
					"proof_set_live":       ps.IsLive,
					"next_challenge_epoch": ps.NextChallengeEpoch,
					"state":                p.getPDPDealState(ps),
					"last_verified_at":     now,
				}
				if existingDeal.ProofSetLive == nil || *existingDeal.ProofSetLive != ps.IsLive {
					Logger.Infow("PDP proof set status changed",
						"proofSetID", ps.ProofSetID,
						"previousLive", existingDeal.ProofSetLive,
						"currentLive", ps.IsLive,
					)
				}
				err = database.DoRetry(ctx, func() error {
					return db.Model(&model.Deal{}).Where("id = ?", existingDeal.ID).Updates(updates).Error
				})
				if err != nil {
					Logger.Errorw("failed to update PDP deal", "dealID", existingDeal.ID, "error", err)
					continue
				}
				Logger.Infow("PDP deal updated", "dealID", existingDeal.ID, "proofSetID", ps.ProofSetID)
				updated++
			} else if errors.Is(err, gorm.ErrRecordNotFound) {
				newState := p.getPDPDealState(ps)
				newDeal := model.Deal{
					DealType:           model.DealTypePDP,
					State:              newState,
					ClientID:           actor.ID,
					Provider:           ps.ProviderAddress.String(),
					PieceCID:           modelPieceCID,
					ProofSetID:         ptr.Of(ps.ProofSetID),
					ProofSetLive:       ptr.Of(ps.IsLive),
					NextChallengeEpoch: ptr.Of(ps.NextChallengeEpoch),
					LastVerifiedAt:     ptr.Of(now),
				}

				err = database.DoRetry(ctx, func() error {
					return db.Create(&newDeal).Error
				})
				if err != nil {
					Logger.Errorw("failed to insert PDP deal", "proofSetID", ps.ProofSetID, "error", err)
					continue
				}
				Logger.Infow("PDP deal inserted", "proofSetID", ps.ProofSetID, "state", newState)
				inserted++
			} else {
				Logger.Errorw("failed to query existing PDP deal", "error", err)
			}
		}
	}

	if bulkClient, ok := p.pdpClient.(PDPBulkClient); ok {
		actorsByAddress := make(map[string][]model.Actor, len(actors))
		for _, actor := range actors {
			actorAddr, err := address.NewFromString(actor.Address)
			if err != nil {
				Logger.Warnw("invalid actor address for PDP tracking", "actorID", actor.ID, "address", actor.Address, "error", err)
				continue
			}
			actorsByAddress[actorAddr.String()] = append(actorsByAddress[actorAddr.String()], actor)
		}

		proofSets, err := bulkClient.GetProofSets(ctx)
		if err != nil {
			return errors.Wrap(err, "failed to get PDP proof sets")
		}
		for _, ps := range proofSets {
			trackedProofSets[ps.ProofSetID] = struct{}{}
			for _, actor := range actorsByAddress[ps.ClientAddress.String()] {
				processProofSet(actor, ps)
			}
		}
	} else {
		for _, actor := range actors {
			Logger.Infof("tracking PDP deals for actor %s", actor.ID)

			actorAddr, err := address.NewFromString(actor.Address)
			if err != nil {
				Logger.Warnw("invalid actor address for PDP tracking", "actorID", actor.ID, "address", actor.Address, "error", err)
				continue
			}

			proofSets, err := p.pdpClient.GetProofSetsForClient(ctx, actorAddr)
			if err != nil {
				Logger.Warnw("failed to get proof sets for actor", "actorID", actor.ID, "error", err)
				continue
			}

			for _, ps := range proofSets {
				trackedProofSets[ps.ProofSetID] = struct{}{}
				processProofSet(actor, ps)
			}
		}
	}

	Logger.Infof("PDP tracker: updated %d deals, inserted %d deals", updated, inserted)
	Logger.Infof("PDP tracker: tracked %d proof sets", len(trackedProofSets))
	return nil
}

// getPDPDealState determines the deal state based on proof set status
func (p *PDPTracker) getPDPDealState(ps ProofSetInfo) model.DealState {
	if ps.IsLive {
		return model.DealActive
	}
	// If not live, it might be proposed (waiting for first challenge) or expired
	// This logic may need refinement based on actual PDP contract semantics
	return model.DealPublished
}
