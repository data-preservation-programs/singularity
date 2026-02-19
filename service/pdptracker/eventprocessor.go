package pdptracker

import (
	"context"
	"math"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gotidy/ptr"
	"github.com/ipfs/go-cid"
	"gorm.io/gorm"
)

// idempotent — safe to re-process after reorgs
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

type inboxRow interface {
	setID() uint64
}

// failed rows are retained for retry on next cycle
func processInbox[R inboxRow](db *gorm.DB, query, table string, fn func(R) error) error {
	var rows []R
	if err := db.Raw(query).Scan(&rows).Error; err != nil {
		return err
	}
	if len(rows) == 0 {
		return nil
	}

	var failed []uint64
	for _, r := range rows {
		if err := fn(r); err != nil {
			Logger.Errorw("inbox processing failed", "table", table, "setId", r.setID(), "error", err)
			failed = append(failed, r.setID())
		}
	}

	return deleteProcessedRows(db, table, "set_id", failed)
}

func deleteProcessedRows(db *gorm.DB, table, keyCol string, failed []uint64) error {
	if len(failed) == 0 {
		return db.Exec("DELETE FROM " + table).Error
	}
	return db.Exec("DELETE FROM "+table+" WHERE "+keyCol+" NOT IN (?)", failed).Error
}

type dataSetCreatedRow struct {
	SetID_          uint64 `gorm:"column:set_id"`
	StorageProvider []byte `gorm:"column:storage_provider"`
	BlockNum        int64  `gorm:"column:block_num"`
}

func (r dataSetCreatedRow) setID() uint64 { return r.SetID_ }

func processDataSetCreated(ctx context.Context, db *gorm.DB, rpcClient *ChainPDPClient) error {
	return processInbox(db,
		"SELECT set_id, storage_provider, block_num FROM pdp_dataset_created",
		"pdp_dataset_created",
		func(r dataSetCreatedRow) error {
			listener, err := rpcClient.GetDataSetListener(ctx, r.SetID_)
			if err != nil {
				return errors.Wrapf(err, "getDataSetListener for set %d", r.SetID_)
			}

			clientAddr, err := commonToDelegatedAddress(listener)
			if err != nil {
				return errors.Wrap(err, "converting listener address")
			}

			providerAddr, err := commonToDelegatedAddress(common.BytesToAddress(r.StorageProvider))
			if err != nil {
				return errors.Wrap(err, "converting provider address")
			}

			ps := model.PDPProofSet{
				SetID:         r.SetID_,
				ClientAddress: clientAddr.String(),
				Provider:      providerAddr.String(),
				CreatedBlock:  r.BlockNum,
			}

			return database.DoRetry(ctx, func() error {
				return db.Where("set_id = ?", r.SetID_).Attrs(ps).FirstOrCreate(&model.PDPProofSet{}).Error
			})
		},
	)
}

func processPiecesChanged(ctx context.Context, db *gorm.DB, rpcClient *ChainPDPClient) error {
	type row struct {
		SetID uint64 `gorm:"column:set_id"`
	}

	setIDs := make(map[uint64]struct{})
	for _, q := range []string{
		"SELECT DISTINCT set_id FROM pdp_pieces_added",
		"SELECT DISTINCT set_id FROM pdp_pieces_removed",
	} {
		var rows []row
		if err := db.Raw(q).Scan(&rows).Error; err != nil {
			return err
		}
		for _, r := range rows {
			setIDs[r.SetID] = struct{}{}
		}
	}

	if len(setIDs) == 0 {
		return nil
	}

	var failed []uint64
	for id := range setIDs {
		if err := reconcileProofSetPieces(ctx, db, rpcClient, id); err != nil {
			Logger.Errorw("failed to reconcile pieces", "setId", id, "error", err)
			failed = append(failed, id)
		}
	}

	if err := deleteProcessedRows(db, "pdp_pieces_added", "set_id", failed); err != nil {
		return err
	}
	return deleteProcessedRows(db, "pdp_pieces_removed", "set_id", failed)
}

func reconcileProofSetPieces(ctx context.Context, db *gorm.DB, rpcClient *ChainPDPClient, setID uint64) error {
	var ps model.PDPProofSet
	if err := db.Where("set_id = ?", setID).First(&ps).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// proof set not yet materialized locally; DataSetCreated may
			// still be pending retry — retain inbox rows for next cycle
			return errors.Errorf("proof set %d not found, retaining piece events", setID)
		}
		return err
	}
	if ps.Deleted {
		Logger.Debugw("ignoring piece events for deleted proof set", "setId", setID)
		return nil
	}

	var wallet model.Wallet
	if err := db.Where("address = ?", ps.ClientAddress).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Logger.Debugw("pieces changed for untracked client", "setId", setID, "client", ps.ClientAddress)
			return nil
		}
		return err
	}

	pieces, err := rpcClient.GetActivePieces(ctx, setID)
	if err != nil {
		return errors.Wrapf(err, "getActivePieces for set %d", setID)
	}

	activeCIDs := make(map[string]cid.Cid, len(pieces))
	for _, c := range pieces {
		if c != cid.Undef {
			activeCIDs[c.String()] = c
		}
	}

	now := time.Now()
	initialState := model.DealPublished
	if ps.IsLive {
		initialState = model.DealActive
	}

	var hadErrors bool

	for _, pieceCID := range activeCIDs {
		modelCID := model.CID(pieceCID)
		err = database.DoRetry(ctx, func() error {
			var existing model.Deal
			result := db.Where("proof_set_id = ? AND piece_cid = ? AND deal_type = ?",
				setID, modelCID, model.DealTypePDP).First(&existing)
			if result.Error == nil {
				if existing.State == model.DealExpired {
					return db.Model(&model.Deal{}).Where("id = ?", existing.ID).
						Update("state", initialState).Error
				}
				return nil
			}
			if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return result.Error
			}
			return db.Create(&model.Deal{
				DealType:       model.DealTypePDP,
				State:          initialState,
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
			hadErrors = true
		}
	}

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
				hadErrors = true
			}
		}
	}

	if hadErrors {
		return errors.Errorf("partial reconciliation failure for proof set %d", setID)
	}
	return nil
}

type nextProvingPeriodRow struct {
	SetID_         uint64 `gorm:"column:set_id"`
	ChallengeEpoch int64  `gorm:"column:challenge_epoch"`
}

func (r nextProvingPeriodRow) setID() uint64 { return r.SetID_ }

func processNextProvingPeriod(ctx context.Context, db *gorm.DB) error {
	return processInbox(db,
		"SELECT set_id, challenge_epoch FROM pdp_next_proving_period",
		"pdp_next_proving_period",
		func(r nextProvingPeriodRow) error {
			if r.ChallengeEpoch > math.MaxInt32 || r.ChallengeEpoch < math.MinInt32 {
				return errors.Errorf("challenge epoch %d overflows int32", r.ChallengeEpoch)
			}
			epoch32 := int32(r.ChallengeEpoch)

			if err := database.DoRetry(ctx, func() error {
				return db.Model(&model.PDPProofSet{}).Where("set_id = ?", r.SetID_).
					Update("challenge_epoch", r.ChallengeEpoch).Error
			}); err != nil {
				return err
			}

			return database.DoRetry(ctx, func() error {
				return db.Model(&model.Deal{}).Where("proof_set_id = ? AND deal_type = ?",
					r.SetID_, model.DealTypePDP).
					Update("next_challenge_epoch", epoch32).Error
			})
		},
	)
}

type possessionProvenRow struct {
	SetID_ uint64 `gorm:"column:set_id"`
}

func (r possessionProvenRow) setID() uint64 { return r.SetID_ }

func processPossessionProven(ctx context.Context, db *gorm.DB) error {
	now := time.Now()
	return processInbox(db,
		"SELECT DISTINCT set_id FROM pdp_possession_proven",
		"pdp_possession_proven",
		func(r possessionProvenRow) error {
			if err := database.DoRetry(ctx, func() error {
				return db.Model(&model.PDPProofSet{}).Where("set_id = ?", r.SetID_).
					Update("is_live", true).Error
			}); err != nil {
				return err
			}

			// only activate non-expired deals; expired deals (from piece removal
			// or dataset deletion) must not be resurrected by a later proof
			return database.DoRetry(ctx, func() error {
				return db.Model(&model.Deal{}).
					Where("proof_set_id = ? AND deal_type = ? AND state != ?",
						r.SetID_, model.DealTypePDP, model.DealExpired).
					Updates(map[string]any{
						"proof_set_live":   true,
						"state":            model.DealActive,
						"last_verified_at": now,
					}).Error
			})
		},
	)
}

type dataSetDeletedRow struct {
	SetID_ uint64 `gorm:"column:set_id"`
}

func (r dataSetDeletedRow) setID() uint64 { return r.SetID_ }

func processDataSetDeleted(ctx context.Context, db *gorm.DB) error {
	return processInbox(db,
		"SELECT set_id FROM pdp_dataset_deleted",
		"pdp_dataset_deleted",
		func(r dataSetDeletedRow) error {
			if err := database.DoRetry(ctx, func() error {
				return db.Model(&model.PDPProofSet{}).Where("set_id = ?", r.SetID_).
					Update("deleted", true).Error
			}); err != nil {
				return err
			}

			return database.DoRetry(ctx, func() error {
				return db.Model(&model.Deal{}).Where("proof_set_id = ? AND deal_type = ?",
					r.SetID_, model.DealTypePDP).
					Update("state", model.DealExpired).Error
			})
		},
	)
}

type spChangedRow struct {
	SetID_ uint64 `gorm:"column:set_id"`
	NewSP  []byte `gorm:"column:new_sp"`
}

func (r spChangedRow) setID() uint64 { return r.SetID_ }

func processSPChanged(ctx context.Context, db *gorm.DB) error {
	return processInbox(db,
		"SELECT set_id, new_sp FROM pdp_sp_changed",
		"pdp_sp_changed",
		func(r spChangedRow) error {
			newAddr, err := commonToDelegatedAddress(common.BytesToAddress(r.NewSP))
			if err != nil {
				return errors.Wrap(err, "converting SP address")
			}

			if err := database.DoRetry(ctx, func() error {
				return db.Model(&model.PDPProofSet{}).Where("set_id = ?", r.SetID_).
					Update("provider", newAddr.String()).Error
			}); err != nil {
				return err
			}

			return database.DoRetry(ctx, func() error {
				return db.Model(&model.Deal{}).Where("proof_set_id = ? AND deal_type = ?",
					r.SetID_, model.DealTypePDP).
					Update("provider", newAddr.String()).Error
			})
		},
	)
}
