package dealpusher

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"github.com/ipfs/go-cid"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

func defaultDDOSchedulingConfig() DDOSchedulingConfig {
	return DDOSchedulingConfig{
		BatchSize:         10,
		ConfirmationDepth: 5,
		PollingInterval:   30 * time.Second,
		TermMin:           518400,
		TermMax:           5256000,
		ExpirationOffset:  172800,
	}
}

// parseProviderActorID extracts the numeric actor ID from an f0 provider string.
func parseProviderActorID(provider string) (uint64, error) {
	s := strings.TrimPrefix(provider, "f0")
	s = strings.TrimPrefix(s, "t0")
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("cannot parse provider actor ID from %q: %w", provider, err)
	}
	return id, nil
}

func (d *DealPusher) runDDOSchedule(ctx context.Context, schedule *model.Schedule) (model.ScheduleState, error) {
	if d.ddoDealManager == nil {
		return model.ScheduleError, errors.New("ddo scheduling dependencies are not configured")
	}
	cfg := d.ddoSchedulingConfig

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

	providerActorID, err := parseProviderActorID(schedule.Provider)
	if err != nil {
		return model.ScheduleError, errors.Wrap(err, "failed to parse provider actor ID")
	}

	// validate SP registration upfront
	spConfig, err := d.ddoDealManager.ValidateSP(ctx, providerActorID)
	if err != nil {
		return model.ScheduleError, errors.Wrap(err, "failed to validate SP")
	}
	if !spConfig.IsActive {
		return model.ScheduleError, fmt.Errorf("provider %s (actor %d) is not active in the DDO contract", schedule.Provider, providerActorID)
	}

	var timer *time.Timer
	current := sumResult{}
	for {
		if ctx.Err() != nil {
			return "", nil
		}

		pending, total, err := d.getDDOScheduleCounts(ctx, schedule)
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
			if schedule.ScheduleCron != "" {
				return "", nil
			}
			return model.ScheduleCompleted, nil
		}

		cars, err := d.findDDOCars(ctx, schedule, attachments, allowedPieceCIDs, overReplicatedCIDs, cfg.BatchSize)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if schedule.ScheduleCron != "" && schedule.ScheduleCronPerpetual {
					return "", nil
				}
				return model.ScheduleCompleted, nil
			}
			return model.ScheduleError, err
		}

		// validate piece sizes against SP config
		for _, car := range cars {
			if uint64(car.PieceSize) < spConfig.MinPieceSize {
				return model.ScheduleError, fmt.Errorf("piece %s size %d below SP min %d", car.PieceCID.String(), car.PieceSize, spConfig.MinPieceSize)
			}
			if spConfig.MaxPieceSize > 0 && uint64(car.PieceSize) > spConfig.MaxPieceSize {
				return model.ScheduleError, fmt.Errorf("piece %s size %d exceeds SP max %d", car.PieceCID.String(), car.PieceSize, spConfig.MaxPieceSize)
			}
		}

		// build piece submissions
		pieces := make([]DDOPieceSubmission, len(cars))
		for i, car := range cars {
			downloadURL := strings.ReplaceAll(schedule.URLTemplate, "{PIECE_CID}", cid.Cid(car.PieceCID).String())
			pieces[i] = DDOPieceSubmission{
				PieceCID:    cid.Cid(car.PieceCID),
				PieceSize:   uint64(car.PieceSize),
				ProviderID:  providerActorID,
				DownloadURL: downloadURL,
			}
		}

		if err := d.ddoDealManager.EnsurePayments(ctx, evmSigner, pieces, cfg); err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed to ensure DDO payments")
		}

		queuedTx, err := d.ddoDealManager.CreateAllocations(ctx, evmSigner, pieces, cfg)
		if err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed to create DDO allocations")
		}

		_, err = d.ddoDealManager.WaitForConfirmations(ctx, queuedTx.Hash, cfg.ConfirmationDepth, cfg.PollingInterval)
		if err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed waiting for DDO transaction confirmation")
		}

		allocationIDs, err := d.ddoDealManager.ParseAllocationIDs(ctx, queuedTx.Hash)
		if err != nil {
			return model.ScheduleError, errors.Wrap(err, "failed to parse DDO allocation IDs")
		}
		if len(allocationIDs) != len(cars) {
			return model.ScheduleError, fmt.Errorf("allocation count mismatch: got %d allocations for %d pieces", len(allocationIDs), len(cars))
		}

		for i, car := range cars {
			allocID := allocationIDs[i]
			dealModel := &model.Deal{
				State:           model.DealProposed,
				DealType:        model.DealTypeDDO,
				Provider:        schedule.Provider,
				PieceCID:        car.PieceCID,
				PieceSize:       car.PieceSize,
				Verified:        schedule.Verified,
				ScheduleID:      &schedule.ID,
				ClientID:        clientID,
				WalletID:        &walletObj.ID,
				DDOAllocationID: &allocID,
			}
			if err := database.DoRetry(ctx, func() error { return db.Create(dealModel).Error }); err != nil {
				return model.ScheduleError, errors.Wrap(err, "failed to create DDO deal")
			}
			current.DealNumber++
			current.DealSize += car.PieceSize
		}
	}
}

func (d *DealPusher) getDDOScheduleCounts(ctx context.Context, schedule *model.Schedule) (sumResult, sumResult, error) {
	db := d.dbNoContext.WithContext(ctx)
	var pending sumResult
	err := db.Model(&model.Deal{}).
		Where("schedule_id = ? AND deal_type = ? AND state IN (?)", schedule.ID, model.DealTypeDDO, []model.DealState{
			model.DealProposed, model.DealPublished,
		}).
		Select("COUNT(*) AS deal_number, SUM(piece_size) AS deal_size").
		Scan(&pending).Error
	if err != nil {
		return sumResult{}, sumResult{}, errors.Wrap(err, "failed to count pending DDO deals")
	}

	var total sumResult
	err = db.Model(&model.Deal{}).
		Where("schedule_id = ? AND deal_type = ? AND state IN (?)", schedule.ID, model.DealTypeDDO, []model.DealState{
			model.DealActive, model.DealProposed, model.DealPublished,
		}).
		Select("COUNT(*) AS deal_number, SUM(piece_size) AS deal_size").
		Scan(&total).Error
	if err != nil {
		return sumResult{}, sumResult{}, errors.Wrap(err, "failed to count total DDO deals")
	}

	return pending, total, nil
}

func (d *DealPusher) findDDOCars(
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
			model.DealTypeDDO,
			[]model.DealState{
				model.DealProposed, model.DealPublished, model.DealActive,
			}).
		Where("piece_cid IS NOT NULL")
	if schedule.Force {
		existingPieceCIDQuery = db.Table("deals").Select("piece_cid").
			Where("schedule_id = ? AND deal_type = ?", schedule.ID, model.DealTypeDDO).
			Where("piece_cid IS NOT NULL")
	}

	var existingPieceCIDs []model.CID
	if err := existingPieceCIDQuery.Find(&existingPieceCIDs).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query existing DDO piece CIDs")
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
			return nil, errors.Wrap(err, "failed to find DDO cars")
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
			return nil, errors.Wrap(err, "failed to find DDO cars by allowed piece CID")
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
