package dealpusher

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

func defaultPDPSchedulingConfig() PDPSchedulingConfig {
	return PDPSchedulingConfig{
		BatchSize:         128,
		GasLimit:          5_000_000,
		ConfirmationDepth: 5,
		PollingInterval:   30 * time.Second,
	}
}

func inferScheduleDealType(schedule *model.Schedule) model.DealType {
	if schedule == nil {
		return model.DealTypeMarket
	}
	providerAddr, err := address.NewFromString(schedule.Provider)
	if err != nil {
		return model.DealTypeMarket
	}
	if providerAddr.Protocol() == address.Delegated {
		return model.DealTypePDP
	}
	return model.DealTypeMarket
}

func (d *DealPusher) runPDPSchedule(ctx context.Context, schedule *model.Schedule) (model.ScheduleState, error) {
	if d.pdpProofSetManager == nil || d.pdpTxConfirmer == nil {
		return model.ScheduleError, errors.New("pdp scheduling dependencies are not configured")
	}
	if err := d.pdpSchedulingConfig.Validate(); err != nil {
		return model.ScheduleError, errors.Wrap(err, "invalid PDP scheduling configuration")
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
		if schedule.TotalDealNumber > 0 && total.DealNumber >= schedule.TotalDealNumber {
			return model.ScheduleCompleted, nil
		}
		if schedule.TotalDealSize > 0 && total.DealSize >= schedule.TotalDealSize {
			return model.ScheduleCompleted, nil
		}
		if schedule.ScheduleCron != "" && schedule.ScheduleDealNumber > 0 && current.DealNumber >= schedule.ScheduleDealNumber {
			return "", nil
		}
		if schedule.ScheduleCron != "" && schedule.ScheduleDealSize > 0 && current.DealSize >= schedule.ScheduleDealSize {
			return "", nil
		}

		cars, err := d.findPDPCars(ctx, schedule, attachments, allowedPieceCIDs, overReplicatedCIDs, d.pdpSchedulingConfig.BatchSize)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if schedule.ScheduleCron != "" && schedule.ScheduleCronPerpetual {
					return "", nil
				}
				return model.ScheduleCompleted, nil
			}
			return model.ScheduleError, err
		}
		if len(cars) == 0 {
			if schedule.ScheduleCron != "" && schedule.ScheduleCronPerpetual {
				return "", nil
			}
			return model.ScheduleCompleted, nil
		}

		walletObj, err := d.walletChooser.Choose(ctx, schedule.Preparation.Wallets)
		if err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed to choose wallet")
		}

		proofSetID, err := d.pdpProofSetManager.EnsureProofSet(ctx, walletObj, schedule.Provider)
		if err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed to ensure PDP proof set")
		}

		pieceCIDs := make([]cid.Cid, 0, len(cars))
		for _, car := range cars {
			pieceCIDs = append(pieceCIDs, cid.Cid(car.PieceCID))
		}
		queuedTx, err := d.pdpProofSetManager.QueueAddRoots(ctx, proofSetID, pieceCIDs, d.pdpSchedulingConfig)
		if err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed to queue PDP root addition transaction")
		}

		_, err = d.pdpTxConfirmer.WaitForConfirmations(ctx, queuedTx.Hash, d.pdpSchedulingConfig.ConfirmationDepth, d.pdpSchedulingConfig.PollingInterval)
		if err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed waiting for PDP transaction confirmation")
		}

		for _, car := range cars {
			proofSetIDCopy := proofSetID
			dealModel := &model.Deal{
				State:      model.DealProposed,
				DealType:   model.DealTypePDP,
				Provider:   schedule.Provider,
				PieceCID:   car.PieceCID,
				PieceSize:  car.PieceSize,
				Verified:   schedule.Verified,
				ScheduleID: &schedule.ID,
				ClientID:   walletObj.ID,
				ProofSetID: &proofSetIDCopy,
			}

			if err := database.DoRetry(ctx, func() error { return db.Create(dealModel).Error }); err != nil {
				return model.ScheduleError, errors.Wrap(err, "failed to create PDP deal")
			}
			current.DealNumber++
			current.DealSize += car.PieceSize
		}
		continue
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
