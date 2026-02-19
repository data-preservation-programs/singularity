package pdptracker

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gotidy/ptr"
	"github.com/ipfs/go-cid"
	"gorm.io/gorm"
)

// processNewEvents reads rows from Shovel integration tables (the inbox),
// materializes state into singularity tables, and deletes processed rows.
// All state changes are idempotent so re-processing after reorgs is safe.
func processNewEvents(ctx context.Context, db *gorm.DB, rpcClient *ChainPDPClient) error {
	// process in dependency order
	if err := processDataSetCreated(ctx, db, rpcClient); err != nil {
		return errors.Wrap(err, "processing DataSetCreated")
	}
	if err := processSPChanged(ctx, db); err != nil {
		return errors.Wrap(err, "processing StorageProviderChanged")
	}
	if err := processPiecesChanged(ctx, db, rpcClient); err != nil {
		return errors.Wrap(err, "processing PiecesAdded/Removed")
	}
	if err := processNextProvingPeriod(ctx, db); err != nil {
		return errors.Wrap(err, "processing NextProvingPeriod")
	}
	if err := processPossessionProven(ctx, db); err != nil {
		return errors.Wrap(err, "processing PossessionProven")
	}
	if err := processDataSetDeleted(ctx, db); err != nil {
		return errors.Wrap(err, "processing DataSetDeleted")
	}
	return nil
}

func processDataSetCreated(ctx context.Context, db *gorm.DB, rpcClient *ChainPDPClient) error {
	type row struct {
		SetID           uint64 `gorm:"column:set_id"`
		StorageProvider []byte `gorm:"column:storage_provider"`
		BlockNum        int64  `gorm:"column:block_num"`
	}

	var rows []row
	if err := db.Raw("SELECT set_id, storage_provider, block_num FROM pdp_dataset_created").Scan(&rows).Error; err != nil {
		return err
	}
	if len(rows) == 0 {
		return nil
	}

	Logger.Infow("processing DataSetCreated events", "count", len(rows))

	for _, r := range rows {
		// get client address via RPC (not emitted in event)
		listener, err := rpcClient.GetDataSetListener(ctx, r.SetID)
		if err != nil {
			Logger.Warnw("failed to get dataset listener", "setId", r.SetID, "error", err)
			continue
		}

		clientAddr, err := commonToDelegatedAddress(listener)
		if err != nil {
			Logger.Warnw("failed to convert listener address", "setId", r.SetID, "error", err)
			continue
		}

		providerAddr, err := commonToDelegatedAddress(common.BytesToAddress(r.StorageProvider))
		if err != nil {
			Logger.Warnw("failed to convert provider address", "setId", r.SetID, "error", err)
			continue
		}

		ps := model.PDPProofSet{
			SetID:         r.SetID,
			ClientAddress: clientAddr.String(),
			Provider:      providerAddr.String(),
			CreatedBlock:  r.BlockNum,
		}

		err = database.DoRetry(ctx, func() error {
			return db.Where("set_id = ?", r.SetID).Attrs(ps).FirstOrCreate(&model.PDPProofSet{}).Error
		})
		if err != nil {
			Logger.Errorw("failed to upsert proof set", "setId", r.SetID, "error", err)
			continue
		}
		Logger.Infow("proof set created", "setId", r.SetID, "client", clientAddr)
	}

	return db.Exec("DELETE FROM pdp_dataset_created").Error
}

// processPiecesChanged handles both PiecesAdded and PiecesRemoved events.
// For each affected proof set, it fetches the current active pieces via RPC
// and reconciles against the local deal records.
func processPiecesChanged(ctx context.Context, db *gorm.DB, rpcClient *ChainPDPClient) error {
	// collect distinct set_ids from both tables
	setIDs := make(map[uint64]struct{})

	type row struct {
		SetID uint64 `gorm:"column:set_id"`
	}

	var addedRows []row
	if err := db.Raw("SELECT DISTINCT set_id FROM pdp_pieces_added").Scan(&addedRows).Error; err != nil {
		return err
	}
	for _, r := range addedRows {
		setIDs[r.SetID] = struct{}{}
	}

	var removedRows []row
	if err := db.Raw("SELECT DISTINCT set_id FROM pdp_pieces_removed").Scan(&removedRows).Error; err != nil {
		return err
	}
	for _, r := range removedRows {
		setIDs[r.SetID] = struct{}{}
	}

	if len(setIDs) == 0 {
		return nil
	}

	Logger.Infow("processing piece changes", "proofSets", len(setIDs))

	for setID := range setIDs {
		if err := reconcileProofSetPieces(ctx, db, rpcClient, setID); err != nil {
			Logger.Errorw("failed to reconcile pieces", "setId", setID, "error", err)
		}
	}

	// clean up both tables
	if err := db.Exec("DELETE FROM pdp_pieces_added").Error; err != nil {
		return err
	}
	return db.Exec("DELETE FROM pdp_pieces_removed").Error
}

func reconcileProofSetPieces(ctx context.Context, db *gorm.DB, rpcClient *ChainPDPClient, setID uint64) error {
	var ps model.PDPProofSet
	if err := db.Where("set_id = ? AND deleted = false", setID).First(&ps).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Logger.Debugw("pieces changed for unknown proof set", "setId", setID)
			return nil
		}
		return err
	}

	// check if this proof set's client is a tracked wallet
	var wallet model.Wallet
	if err := db.Where("address = ?", ps.ClientAddress).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Logger.Debugw("pieces changed for untracked client", "setId", setID, "client", ps.ClientAddress)
			return nil
		}
		return err
	}

	// fetch current active pieces from chain
	pieces, err := rpcClient.GetActivePieces(ctx, setID)
	if err != nil {
		return errors.Wrapf(err, "getActivePieces for set %d", setID)
	}

	// build set of active CIDs
	activeCIDs := make(map[string]cid.Cid, len(pieces))
	for _, c := range pieces {
		if c == cid.Undef {
			continue
		}
		activeCIDs[c.String()] = c
	}

	now := time.Now()

	// create deals for newly active pieces
	for _, pieceCID := range activeCIDs {
		modelCID := model.CID(pieceCID)
		err = database.DoRetry(ctx, func() error {
			var existing model.Deal
			result := db.Where("proof_set_id = ? AND piece_cid = ? AND deal_type = ?",
				setID, modelCID, model.DealTypePDP).First(&existing)
			if result.Error == nil {
				// already tracked, ensure not expired
				if existing.State == model.DealExpired {
					return db.Model(&model.Deal{}).Where("id = ?", existing.ID).
						Update("state", model.DealPublished).Error
				}
				return nil
			}
			if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return result.Error
			}
			return db.Create(&model.Deal{
				DealType:       model.DealTypePDP,
				State:          model.DealPublished,
				ClientID:       wallet.ID,
				Provider:       ps.Provider,
				PieceCID:       modelCID,
				ProofSetID:     ptr.Of(setID),
				ProofSetLive:   ptr.Of(ps.IsLive),
				LastVerifiedAt: ptr.Of(now),
			}).Error
		})
		if err != nil {
			Logger.Errorw("failed to upsert deal", "setId", setID, "pieceCid", pieceCID, "error", err)
		}
	}

	// expire deals for pieces no longer active
	var existingDeals []model.Deal
	if err := db.Where("proof_set_id = ? AND deal_type = ? AND state != ?",
		setID, model.DealTypePDP, model.DealExpired).Find(&existingDeals).Error; err != nil {
		return err
	}

	for _, deal := range existingDeals {
		if _, ok := activeCIDs[deal.PieceCID.String()]; !ok {
			err = database.DoRetry(ctx, func() error {
				return db.Model(&model.Deal{}).Where("id = ?", deal.ID).
					Update("state", model.DealExpired).Error
			})
			if err != nil {
				Logger.Errorw("failed to expire removed deal", "dealId", deal.ID, "error", err)
			}
		}
	}

	return nil
}

func processNextProvingPeriod(ctx context.Context, db *gorm.DB) error {
	type row struct {
		SetID          uint64 `gorm:"column:set_id"`
		ChallengeEpoch int64  `gorm:"column:challenge_epoch"`
	}

	var rows []row
	if err := db.Raw("SELECT set_id, challenge_epoch FROM pdp_next_proving_period").Scan(&rows).Error; err != nil {
		return err
	}
	if len(rows) == 0 {
		return nil
	}

	Logger.Infow("processing NextProvingPeriod events", "count", len(rows))

	for _, r := range rows {
		epoch32 := int32(r.ChallengeEpoch)

		err := database.DoRetry(ctx, func() error {
			return db.Model(&model.PDPProofSet{}).Where("set_id = ?", r.SetID).
				Update("challenge_epoch", r.ChallengeEpoch).Error
		})
		if err != nil {
			Logger.Errorw("failed to update challenge epoch", "setId", r.SetID, "error", err)
			continue
		}

		err = database.DoRetry(ctx, func() error {
			return db.Model(&model.Deal{}).Where("proof_set_id = ? AND deal_type = ?",
				r.SetID, model.DealTypePDP).
				Update("next_challenge_epoch", epoch32).Error
		})
		if err != nil {
			Logger.Errorw("failed to update deal challenge epochs", "setId", r.SetID, "error", err)
		}
	}

	return db.Exec("DELETE FROM pdp_next_proving_period").Error
}

func processPossessionProven(ctx context.Context, db *gorm.DB) error {
	type row struct {
		SetID uint64 `gorm:"column:set_id"`
	}

	var rows []row
	if err := db.Raw("SELECT DISTINCT set_id FROM pdp_possession_proven").Scan(&rows).Error; err != nil {
		return err
	}
	if len(rows) == 0 {
		return nil
	}

	Logger.Infow("processing PossessionProven events", "count", len(rows))

	now := time.Now()
	for _, r := range rows {
		err := database.DoRetry(ctx, func() error {
			return db.Model(&model.PDPProofSet{}).Where("set_id = ?", r.SetID).
				Update("is_live", true).Error
		})
		if err != nil {
			Logger.Errorw("failed to update proof set liveness", "setId", r.SetID, "error", err)
			continue
		}

		err = database.DoRetry(ctx, func() error {
			return db.Model(&model.Deal{}).Where("proof_set_id = ? AND deal_type = ?",
				r.SetID, model.DealTypePDP).
				Updates(map[string]any{
					"proof_set_live":   true,
					"state":            model.DealActive,
					"last_verified_at": now,
				}).Error
		})
		if err != nil {
			Logger.Errorw("failed to update deal liveness", "setId", r.SetID, "error", err)
		}
	}

	return db.Exec("DELETE FROM pdp_possession_proven").Error
}

func processDataSetDeleted(ctx context.Context, db *gorm.DB) error {
	type row struct {
		SetID uint64 `gorm:"column:set_id"`
	}

	var rows []row
	if err := db.Raw("SELECT set_id FROM pdp_dataset_deleted").Scan(&rows).Error; err != nil {
		return err
	}
	if len(rows) == 0 {
		return nil
	}

	Logger.Infow("processing DataSetDeleted events", "count", len(rows))

	for _, r := range rows {
		err := database.DoRetry(ctx, func() error {
			return db.Model(&model.PDPProofSet{}).Where("set_id = ?", r.SetID).
				Update("deleted", true).Error
		})
		if err != nil {
			Logger.Errorw("failed to mark proof set deleted", "setId", r.SetID, "error", err)
			continue
		}

		err = database.DoRetry(ctx, func() error {
			return db.Model(&model.Deal{}).Where("proof_set_id = ? AND deal_type = ?",
				r.SetID, model.DealTypePDP).
				Update("state", model.DealExpired).Error
		})
		if err != nil {
			Logger.Errorw("failed to expire deals for deleted set", "setId", r.SetID, "error", err)
		}
	}

	return db.Exec("DELETE FROM pdp_dataset_deleted").Error
}

func processSPChanged(ctx context.Context, db *gorm.DB) error {
	type row struct {
		SetID uint64 `gorm:"column:set_id"`
		NewSP []byte `gorm:"column:new_sp"`
	}

	var rows []row
	if err := db.Raw("SELECT set_id, new_sp FROM pdp_sp_changed").Scan(&rows).Error; err != nil {
		return err
	}
	if len(rows) == 0 {
		return nil
	}

	Logger.Infow("processing StorageProviderChanged events", "count", len(rows))

	for _, r := range rows {
		newAddr, err := commonToDelegatedAddress(common.BytesToAddress(r.NewSP))
		if err != nil {
			Logger.Warnw("failed to convert SP address", "setId", r.SetID, "error", err)
			continue
		}

		err = database.DoRetry(ctx, func() error {
			return db.Model(&model.PDPProofSet{}).Where("set_id = ?", r.SetID).
				Update("provider", newAddr.String()).Error
		})
		if err != nil {
			Logger.Errorw("failed to update proof set provider", "setId", r.SetID, "error", err)
			continue
		}

		err = database.DoRetry(ctx, func() error {
			return db.Model(&model.Deal{}).Where("proof_set_id = ? AND deal_type = ?",
				r.SetID, model.DealTypePDP).
				Update("provider", newAddr.String()).Error
		})
		if err != nil {
			Logger.Errorw("failed to update deal provider", "setId", r.SetID, "error", err)
		}
	}

	return db.Exec("DELETE FROM pdp_sp_changed").Error
}
