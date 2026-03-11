package dealpusher

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

const pdpDealEpochSentinel = int32(math.MaxInt32)

func defaultPDPSchedulingConfig() PDPSchedulingConfig {
	return PDPSchedulingConfig{
		BatchSize:            128,
		MaxPiecesPerProofSet: 1024,
		ConfirmationDepth:    5,
		PollingInterval:      30 * time.Second,
	}
}

func validatePDPProofSetPieceSize(pieceSize int64) error {
	if pieceSize <= 0 {
		return fmt.Errorf("piece size must be greater than 0, got %d", pieceSize)
	}
	if !util.IsPowerOfTwo(uint64(pieceSize)) {
		return fmt.Errorf("piece size must be a power of two, got %d", pieceSize)
	}
	if pieceSize > model.PDPProofSetMaxPieceSize {
		return fmt.Errorf("piece size %d exceeds max allowed %d (1 GiB minus FR32 overhead)", pieceSize, model.PDPProofSetMaxPieceSize)
	}
	return nil
}

func (d *DealPusher) validatePDPPreparationPieceSizes(ctx context.Context, schedule *model.Schedule) error {
	db := d.dbNoContext.WithContext(ctx)

	var oversized model.Car
	err := db.Model(&model.Car{}).
		Select("piece_cid", "piece_size").
		Where("preparation_id = ? AND piece_size > ?", schedule.PreparationID, model.PDPProofSetMaxPieceSize).
		Order("piece_size DESC").
		First(&oversized).Error
	if err == nil {
		return fmt.Errorf(
			"current PDP proofset piece limit is 1 GiB minus FR32 overhead; preparation %d has oversized piece %s (%d bytes)",
			schedule.PreparationID,
			oversized.PieceCID.String(),
			oversized.PieceSize,
		)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.Wrap(err, "failed to validate preparation piece sizes for PDP")
}

// resolveProviderEVMAddress looks up the provider's Actor record and derives
// the EVM address from its delegated (f410) filecoin address.
func (d *DealPusher) resolveProviderEVMAddress(ctx context.Context, provider string) (common.Address, error) {
	db := d.dbNoContext.WithContext(ctx)

	var actor model.Actor
	err := db.Where("address = ? OR id = ?", provider, provider).First(&actor).Error
	if err != nil {
		return common.Address{}, errors.Wrapf(err, "failed to resolve actor for provider %s", provider)
	}

	addr, err := address.NewFromString(actor.Address)
	if err != nil {
		return common.Address{}, errors.Wrapf(err, "failed to parse actor address %s", actor.Address)
	}
	if addr.Protocol() != address.Delegated {
		return common.Address{}, fmt.Errorf("provider actor address %s is not a delegated (f410) address", actor.Address)
	}

	payload := addr.Payload()
	// delegated address payload: first varint byte(s) for namespace, then 20 bytes for EVM address
	// for f410 (namespace 10), payload[0] is the namespace varint, rest is the subaddress
	if len(payload) < 21 {
		return common.Address{}, fmt.Errorf("provider delegated address payload too short: %d bytes", len(payload))
	}
	// skip namespace varint (1 byte for namespace 10)
	return common.BytesToAddress(payload[1:21]), nil
}

func (d *DealPusher) runPDPSchedule(ctx context.Context, schedule *model.Schedule) (model.ScheduleState, error) {
	if d.pdpProofSetManager == nil || d.pdpTxConfirmer == nil {
		return model.ScheduleError, errors.New("pdp scheduling dependencies are not configured")
	}
	cfg := d.pdpSchedulingConfig
	if err := cfg.Validate(); err != nil {
		return model.ScheduleError, errors.Wrap(err, "invalid PDP scheduling configuration")
	}
	if err := d.validatePDPPreparationPieceSizes(ctx, schedule); err != nil {
		return model.ScheduleError, err
	}

	// resolve SP EVM address upfront -- needed for handoff after filling
	spEVMAddr, err := d.resolveProviderEVMAddress(ctx, schedule.Provider)
	if err != nil {
		return model.ScheduleError, errors.Wrap(err, "failed to resolve provider EVM address for PDP handoff")
	}

	db := d.dbNoContext.WithContext(ctx)
	var attachments []model.SourceAttachment
	if err := db.Model(&model.SourceAttachment{}).
		Where("preparation_id = ?", schedule.PreparationID).
		Find(&attachments).Error; err != nil {
		return model.ScheduleError, errors.Wrap(err, "failed to find attachments")
	}

	allowedPieceCIDs := make([]model.CID, 0, len(schedule.AllowedPieceCIDs))
	for _, c := range schedule.AllowedPieceCIDs {
		parsed, err := cid.Parse(c)
		if err != nil {
			return model.ScheduleError, errors.Wrapf(err, "failed to parse CID %s", c)
		}
		allowedPieceCIDs = append(allowedPieceCIDs, model.CID(parsed))
	}

	overReplicatedCIDs := db.
		Table("deals").
		Select("piece_cid").
		Where("state in ?", []model.DealState{model.DealProposed, model.DealPublished, model.DealActive}).
		Group("piece_cid").
		Having("count(*) >= ?", d.maxReplicas)

	if schedule.Preparation == nil || schedule.Preparation.Wallet == nil {
		return model.ScheduleError, errors.New("schedule has no wallet configured")
	}
	walletObj := *schedule.Preparation.Wallet

	evmSigner, err := keystore.EVMSigner(d.keyStore, walletObj)
	if err != nil {
		return model.ScheduleError, errors.Wrap(err, "failed to load EVM signer for wallet")
	}

	clientID := ""
	if walletObj.ActorID != nil {
		clientID = *walletObj.ActorID
	}

	// track the current proof set so we can propose transfer when it fills
	var currentProofSetID uint64
	var currentProofSetPieceCount int

	var timer *time.Timer
	current := sumResult{}
	for {
		if ctx.Err() != nil {
			return "", nil
		}

		pending, total, err := d.getPDPScheduleCounts(ctx, schedule)
		if err != nil {
			return model.ScheduleError, err
		}

		shouldWait := false
		if schedule.MaxPendingDealNumber > 0 && pending.DealNumber >= schedule.MaxPendingDealNumber {
			shouldWait = true
		}
		if schedule.MaxPendingDealSize > 0 && pending.DealSize >= schedule.MaxPendingDealSize {
			shouldWait = true
		}
		if shouldWait {
			if timer == nil {
				timer = time.NewTimer(waitPendingInterval)
				defer timer.Stop()
			} else {
				timer.Reset(waitPendingInterval)
			}
			select {
			case <-ctx.Done():
				return "", nil
			case <-timer.C:
			}
			continue
		}

		scheduleComplete := false
		if schedule.TotalDealNumber > 0 && total.DealNumber >= schedule.TotalDealNumber {
			scheduleComplete = true
		}
		if schedule.TotalDealSize > 0 && total.DealSize >= schedule.TotalDealSize {
			scheduleComplete = true
		}
		if schedule.ScheduleCron != "" && schedule.ScheduleDealNumber > 0 && current.DealNumber >= schedule.ScheduleDealNumber {
			scheduleComplete = true
		}
		if schedule.ScheduleCron != "" && schedule.ScheduleDealSize > 0 && current.DealSize >= schedule.ScheduleDealSize {
			scheduleComplete = true
		}

		if scheduleComplete {
			// propose transfer for any partially-filled proof set before exiting
			if currentProofSetID != 0 && currentProofSetPieceCount > 0 {
				if err := d.pdpProofSetManager.ProposeTransfer(ctx, evmSigner, currentProofSetID, spEVMAddr); err != nil {
					return model.ScheduleError, errors.Wrap(err, "failed to propose transfer for final proof set")
				}
			}
			if schedule.ScheduleCron != "" {
				return "", nil
			}
			return model.ScheduleCompleted, nil
		}

		// cap batch to remaining room in current proof set
		batchLimit := cfg.BatchSize
		if currentProofSetID != 0 {
			remaining := cfg.MaxPiecesPerProofSet - currentProofSetPieceCount
			if remaining <= 0 {
				// current proof set is full -- propose transfer and reset
				if err := d.pdpProofSetManager.ProposeTransfer(ctx, evmSigner, currentProofSetID, spEVMAddr); err != nil {
					return model.ScheduleError, errors.Wrap(err, "failed to propose transfer for full proof set")
				}
				currentProofSetID = 0
				currentProofSetPieceCount = 0
				continue
			}
			if remaining < batchLimit {
				batchLimit = remaining
			}
		}

		cars, err := d.findPDPCars(ctx, schedule, attachments, allowedPieceCIDs, overReplicatedCIDs, batchLimit)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// no more cars -- propose transfer for any partial proof set
				if currentProofSetID != 0 && currentProofSetPieceCount > 0 {
					if err := d.pdpProofSetManager.ProposeTransfer(ctx, evmSigner, currentProofSetID, spEVMAddr); err != nil {
						return model.ScheduleError, errors.Wrap(err, "failed to propose transfer for final proof set")
					}
				}
				if schedule.ScheduleCron != "" && schedule.ScheduleCronPerpetual {
					return "", nil
				}
				return model.ScheduleCompleted, nil
			}
			return model.ScheduleError, err
		}
		if len(cars) == 0 {
			if currentProofSetID != 0 && currentProofSetPieceCount > 0 {
				if err := d.pdpProofSetManager.ProposeTransfer(ctx, evmSigner, currentProofSetID, spEVMAddr); err != nil {
					return model.ScheduleError, errors.Wrap(err, "failed to propose transfer for final proof set")
				}
			}
			if schedule.ScheduleCron != "" && schedule.ScheduleCronPerpetual {
				return "", nil
			}
			return model.ScheduleCompleted, nil
		}
		for _, car := range cars {
			if err := validatePDPProofSetPieceSize(car.PieceSize); err != nil {
				return model.ScheduleError, errors.Wrapf(err, "invalid piece size for piece %s", car.PieceCID.String())
			}
		}

		proofSetID, err := d.pdpProofSetManager.EnsureProofSet(ctx, evmSigner, schedule.Provider, cfg)
		if err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed to ensure PDP proof set")
		}
		if currentProofSetID == 0 {
			currentProofSetID = proofSetID
			// load existing piece count from DB in case we're resuming
			var ps model.PDPProofSet
			if err := db.Where("set_id = ?", proofSetID).First(&ps).Error; err == nil {
				currentProofSetPieceCount = ps.PieceCount
			}
		}

		pieceCIDs := make([]cid.Cid, 0, len(cars))
		pieceSizes := make([]int64, 0, len(cars))
		for _, car := range cars {
			pieceCIDs = append(pieceCIDs, cid.Cid(car.PieceCID))
			pieceSizes = append(pieceSizes, car.PieceSize)
		}
		queuedTx, err := d.pdpProofSetManager.QueueAddRoots(ctx, evmSigner, proofSetID, pieceCIDs, pieceSizes, cfg)
		if err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed to queue PDP root addition transaction")
		}

		_, err = d.pdpTxConfirmer.WaitForConfirmations(ctx, queuedTx.Hash, cfg.ConfirmationDepth, cfg.PollingInterval)
		if err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed waiting for PDP transaction confirmation")
		}

		// update durable piece count after confirmed on-chain add
		if err := d.pdpProofSetManager.IncrementPieceCount(ctx, proofSetID, len(cars)); err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed to update proof set piece count")
		}
		currentProofSetPieceCount += len(cars)

		for _, car := range cars {
			proofSetIDCopy := proofSetID
			dealModel := &model.Deal{
				State:      model.DealProposed,
				DealType:   model.DealTypePDP,
				Provider:   schedule.Provider,
				PieceCID:   car.PieceCID,
				PieceSize:  car.PieceSize,
				StartEpoch: pdpDealEpochSentinel,
				EndEpoch:   pdpDealEpochSentinel,
				Verified:   schedule.Verified,
				ScheduleID: &schedule.ID,
				ClientID:   clientID,
				WalletID:   &walletObj.ID,
				ProofSetID: &proofSetIDCopy,
			}

			if err := database.DoRetry(ctx, func() error { return db.Create(dealModel).Error }); err != nil {
				return model.ScheduleError, errors.Wrap(err, "failed to create PDP deal")
			}
			current.DealNumber++
			current.DealSize += car.PieceSize
		}

		// check if proof set is now full
		if currentProofSetPieceCount >= cfg.MaxPiecesPerProofSet {
			if err := d.pdpProofSetManager.ProposeTransfer(ctx, evmSigner, currentProofSetID, spEVMAddr); err != nil {
				return model.ScheduleError, errors.Wrap(err, "failed to propose transfer for full proof set")
			}
			currentProofSetID = 0
			currentProofSetPieceCount = 0
		}
	}
}

func (d *DealPusher) getPDPScheduleCounts(ctx context.Context, schedule *model.Schedule) (sumResult, sumResult, error) {
	db := d.dbNoContext.WithContext(ctx)
	var pending sumResult
	err := db.Model(&model.Deal{}).
		Where("schedule_id = ? AND deal_type = ? AND state IN (?)", schedule.ID, model.DealTypePDP, []model.DealState{
			model.DealProposed, model.DealPublished,
		}).
		Select("COUNT(*) AS deal_number, SUM(piece_size) AS deal_size").
		Scan(&pending).Error
	if err != nil {
		return sumResult{}, sumResult{}, errors.Wrap(err, "failed to count pending PDP deals")
	}

	var total sumResult
	err = db.Model(&model.Deal{}).
		Where("schedule_id = ? AND deal_type = ? AND state IN (?)", schedule.ID, model.DealTypePDP, []model.DealState{
			model.DealActive, model.DealProposed, model.DealPublished,
		}).
		Select("COUNT(*) AS deal_number, SUM(piece_size) AS deal_size").
		Scan(&total).Error
	if err != nil {
		return sumResult{}, sumResult{}, errors.Wrap(err, "failed to count total PDP deals")
	}

	return pending, total, nil
}

func (d *DealPusher) findPDPCars(
	ctx context.Context,
	schedule *model.Schedule,
	attachments []model.SourceAttachment,
	allowedPieceCIDs []model.CID,
	overReplicatedCIDs *gorm.DB,
	limit int,
) ([]model.Car, error) {
	db := d.dbNoContext.WithContext(ctx)
	attachmentIDs := underscore.Map(attachments, func(a model.SourceAttachment) uint32 { return uint32(a.ID) })
	existingPieceCIDQuery := db.Table("deals").Select("piece_cid").
		Where("provider = ? AND deal_type = ? AND state IN (?)",
			schedule.Provider,
			model.DealTypePDP,
			[]model.DealState{
				model.DealProposed, model.DealPublished, model.DealActive,
			}).
		Where("piece_cid IS NOT NULL")
	if schedule.Force {
		existingPieceCIDQuery = db.Table("deals").Select("piece_cid").
			Where("schedule_id = ? AND deal_type = ?", schedule.ID, model.DealTypePDP).
			Where("piece_cid IS NOT NULL")
	}

	var existingPieceCIDs []model.CID
	if err := existingPieceCIDQuery.Find(&existingPieceCIDs).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query existing PDP piece CIDs")
	}
	existingSet := make(map[string]struct{}, len(existingPieceCIDs))
	for _, existing := range existingPieceCIDs {
		existingSet[cid.Cid(existing).String()] = struct{}{}
	}

	baseQuery := func() *gorm.DB {
		query := db.Where("attachment_id IN ?", attachmentIDs)
		if d.maxReplicas > 0 && !schedule.Force {
			query = query.Where("piece_cid NOT IN (?)", overReplicatedCIDs)
		}
		return query
	}

	if len(allowedPieceCIDs) == 0 {
		var cars []model.Car
		if err := baseQuery().Find(&cars).Error; err != nil {
			return nil, errors.Wrap(err, "failed to find PDP cars")
		}
		filtered := make([]model.Car, 0, limit)
		for _, car := range cars {
			if _, exists := existingSet[cid.Cid(car.PieceCID).String()]; exists {
				continue
			}
			filtered = append(filtered, car)
			if len(filtered) >= limit {
				break
			}
		}
		if len(filtered) == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		return filtered, nil
	}

	cars := make([]model.Car, 0, limit)
	pieceCIDChunks := util.ChunkSlice(allowedPieceCIDs, util.BatchSize)
	for _, pieceCIDChunk := range pieceCIDChunks {
		if len(cars) >= limit {
			break
		}
		var chunkCars []model.Car
		if err := baseQuery().
			Where("piece_cid IN ?", pieceCIDChunk).
			Find(&chunkCars).Error; err != nil {
			return nil, errors.Wrap(err, "failed to find PDP cars by allowed piece CID")
		}
		for _, car := range chunkCars {
			if _, exists := existingSet[cid.Cid(car.PieceCID).String()]; exists {
				continue
			}
			cars = append(cars, car)
			if len(cars) >= limit {
				break
			}
		}
	}
	if len(cars) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return cars, nil
}
