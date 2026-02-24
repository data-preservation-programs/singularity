package pdptracker

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/indexsupply/shovel/shovel/config"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func createShovelTables(t *testing.T, db *gorm.DB) {
	t.Helper()
	conf := buildShovelConfig("postgres://unused", "https://unused", 314, common.Address{}, 0)
	require.NoError(t, config.ValidateFix(&conf))
	for _, ig := range conf.Integrations {
		for _, stmt := range ig.Table.DDL() {
			require.NoError(t, db.Exec(stmt).Error)
		}
	}
}

var testPieceCID cid.Cid

func init() {
	c, err := cid.Decode("baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq")
	if err != nil {
		panic(err)
	}
	testPieceCID = c
}

type pgTestEnv struct {
	ctx         context.Context
	db          *gorm.DB
	client      *ChainPDPClient
	mock        *mockContractCaller
	listenerEth common.Address
	providerEth common.Address
	listenerFil address.Address
	providerFil address.Address
}

func pgTest(t *testing.T, fn func(t *testing.T, e pgTestEnv)) {
	t.Helper()
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		if db.Dialector.Name() != "postgres" {
			t.Skip("PDP event processing requires Postgres")
		}
		createShovelTables(t, db)

		orig := address.CurrentNetwork
		t.Cleanup(func() { address.CurrentNetwork = orig })
		address.CurrentNetwork = address.Mainnet

		le := common.HexToAddress("0x1111111111111111111111111111111111111111")
		pe := common.HexToAddress("0x2222222222222222222222222222222222222222")
		lf, err := commonToDelegatedAddress(le)
		require.NoError(t, err)
		pf, err := commonToDelegatedAddress(pe)
		require.NoError(t, err)

		m := &mockContractCaller{
			listeners: map[uint64]common.Address{1: le},
			pieces:    map[uint64][]cid.Cid{1: {testPieceCID}},
		}

		fn(t, pgTestEnv{
			ctx: ctx, db: db,
			client:      &ChainPDPClient{contract: m, pageSize: 100},
			mock:        m,
			listenerEth: le, providerEth: pe,
			listenerFil: lf, providerFil: pf,
		})
	})
}

func (e pgTestEnv) setupFixtures(t *testing.T) {
	t.Helper()
	require.NoError(t, e.db.Create(&model.Wallet{
		ID: "f0100", Address: e.listenerFil.String(),
	}).Error)
	require.NoError(t, e.db.Create(&model.PDPProofSet{
		SetID: 1, ClientAddress: e.listenerFil.String(),
		Provider: e.providerFil.String(), CreatedBlock: 100,
	}).Error)
}

func (e pgTestEnv) insertDeal(t *testing.T, state model.DealState, opts ...func(*model.Deal)) model.Deal {
	t.Helper()
	setID := uint64(1)
	d := model.Deal{
		DealType:   model.DealTypePDP,
		State:      state,
		ClientID:   "f0100",
		ProofSetID: &setID,
	}
	for _, o := range opts {
		o(&d)
	}
	require.NoError(t, e.db.Create(&d).Error)
	return d
}

func (e pgTestEnv) shovelCount(t *testing.T, table string) int64 {
	t.Helper()
	var n int64
	require.NoError(t, e.db.Raw("SELECT COUNT(*) FROM "+table).Scan(&n).Error)
	return n
}

func TestProcessDataSetCreated(t *testing.T) {
	pgTest(t, func(t *testing.T, e pgTestEnv) {
		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_dataset_created (set_id, storage_provider, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?, ?)",
			1, e.providerEth.Bytes(), 100, 0, 0,
		).Error)

		require.NoError(t, processDataSetCreated(e.ctx, e.db, e.client))

		var ps model.PDPProofSet
		require.NoError(t, e.db.Where("set_id = ?", 1).First(&ps).Error)
		require.Equal(t, e.listenerFil.String(), ps.ClientAddress)
		require.Equal(t, e.providerFil.String(), ps.Provider)
		require.EqualValues(t, 100, ps.CreatedBlock)
		require.EqualValues(t, 0, e.shovelCount(t, "pdp_dataset_created"))
	})
}

func TestProcessDataSetCreated_Idempotent(t *testing.T) {
	pgTest(t, func(t *testing.T, e pgTestEnv) {
		for i := range 2 {
			require.NoError(t, e.db.Exec(
				"INSERT INTO pdp_dataset_created (set_id, storage_provider, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?, ?)",
				1, e.providerEth.Bytes(), int64(100+i), 0, 0,
			).Error)
			require.NoError(t, processDataSetCreated(e.ctx, e.db, e.client))
		}

		var count int64
		require.NoError(t, e.db.Model(&model.PDPProofSet{}).Count(&count).Error)
		require.EqualValues(t, 1, count)
	})
}

func TestProcessPiecesChanged_CreatesDeals(t *testing.T) {
	pgTest(t, func(t *testing.T, e pgTestEnv) {
		e.setupFixtures(t)

		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_pieces_added (set_id, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?)",
			1, 200, 0, 0,
		).Error)
		require.NoError(t, processPiecesChanged(e.ctx, e.db, e.client))

		var deals []model.Deal
		require.NoError(t, e.db.Where("deal_type = ?", model.DealTypePDP).Find(&deals).Error)
		require.Len(t, deals, 1)
		require.Equal(t, model.DealPublished, deals[0].State)
		require.Equal(t, testPieceCID.String(), deals[0].PieceCID.String())
		require.EqualValues(t, 1, *deals[0].ProofSetID)
		require.Equal(t, "f0100", deals[0].ClientID)
		require.EqualValues(t, 0, e.shovelCount(t, "pdp_pieces_added"))
	})
}

func TestProcessPiecesChanged_LiveProofSetCreatesActiveDeals(t *testing.T) {
	pgTest(t, func(t *testing.T, e pgTestEnv) {
		e.setupFixtures(t)
		require.NoError(t, e.db.Model(&model.PDPProofSet{}).Where("set_id = ?", 1).
			Update("is_live", true).Error)

		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_pieces_added (set_id, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?)",
			1, 200, 0, 0,
		).Error)
		require.NoError(t, processPiecesChanged(e.ctx, e.db, e.client))

		var deals []model.Deal
		require.NoError(t, e.db.Where("deal_type = ?", model.DealTypePDP).Find(&deals).Error)
		require.Len(t, deals, 1)
		require.Equal(t, model.DealActive, deals[0].State)
		require.True(t, *deals[0].ProofSetLive)
	})
}

func TestProcessPiecesChanged_ExpiresRemovedPieces(t *testing.T) {
	pgTest(t, func(t *testing.T, e pgTestEnv) {
		e.setupFixtures(t)
		e.insertDeal(t, model.DealActive, func(d *model.Deal) {
			d.Provider = e.providerFil.String()
			d.PieceCID = model.CID(testPieceCID)
		})

		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_pieces_removed (set_id, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?)",
			1, 200, 0, 0,
		).Error)

		// contract returns empty active pieces
		e.mock.pieces[1] = nil
		require.NoError(t, processPiecesChanged(e.ctx, e.db, e.client))

		var deal model.Deal
		require.NoError(t, e.db.Where("deal_type = ?", model.DealTypePDP).First(&deal).Error)
		require.Equal(t, model.DealExpired, deal.State)
	})
}

func TestProcessNextProvingPeriod(t *testing.T) {
	pgTest(t, func(t *testing.T, e pgTestEnv) {
		e.setupFixtures(t)
		e.insertDeal(t, model.DealPublished)

		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_next_proving_period (set_id, challenge_epoch, leaf_count, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?, ?, ?)",
			1, 500, 42, 200, 0, 0,
		).Error)

		require.NoError(t, processNextProvingPeriod(e.ctx, e.db))

		var ps model.PDPProofSet
		require.NoError(t, e.db.Where("set_id = ?", 1).First(&ps).Error)
		require.EqualValues(t, 500, *ps.ChallengeEpoch)

		var deal model.Deal
		require.NoError(t, e.db.Where("deal_type = ?", model.DealTypePDP).First(&deal).Error)
		require.EqualValues(t, 500, *deal.NextChallengeEpoch)
		require.EqualValues(t, 0, e.shovelCount(t, "pdp_next_proving_period"))
	})
}

func TestProcessNextProvingPeriod_RetainedWhenProofSetMissing(t *testing.T) {
	pgTest(t, func(t *testing.T, e pgTestEnv) {
		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_next_proving_period (set_id, challenge_epoch, leaf_count, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?, ?, ?)",
			1, 500, 42, 200, 0, 0,
		).Error)

		// process — proof set missing, row must be retained
		require.NoError(t, processNextProvingPeriod(e.ctx, e.db))
		require.EqualValues(t, 1, e.shovelCount(t, "pdp_next_proving_period"))

		// create the proof set, retry
		e.setupFixtures(t)
		e.insertDeal(t, model.DealPublished)
		require.NoError(t, processNextProvingPeriod(e.ctx, e.db))
		require.EqualValues(t, 0, e.shovelCount(t, "pdp_next_proving_period"))

		var ps model.PDPProofSet
		require.NoError(t, e.db.Where("set_id = ?", 1).First(&ps).Error)
		require.EqualValues(t, 500, *ps.ChallengeEpoch)
	})
}

func TestProcessPossessionProven(t *testing.T) {
	pgTest(t, func(t *testing.T, e pgTestEnv) {
		e.setupFixtures(t)
		e.insertDeal(t, model.DealPublished)

		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_possession_proven (set_id, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?)",
			1, 200, 0, 0,
		).Error)
		require.NoError(t, processPossessionProven(e.ctx, e.db))

		var ps model.PDPProofSet
		require.NoError(t, e.db.Where("set_id = ?", 1).First(&ps).Error)
		require.True(t, ps.IsLive)

		var deal model.Deal
		require.NoError(t, e.db.Where("deal_type = ?", model.DealTypePDP).First(&deal).Error)
		require.Equal(t, model.DealActive, deal.State)
		require.True(t, *deal.ProofSetLive)
		require.NotNil(t, deal.LastVerifiedAt)
		require.EqualValues(t, 0, e.shovelCount(t, "pdp_possession_proven"))
	})
}

func TestProcessPossessionProven_RetainedWhenProofSetMissing(t *testing.T) {
	pgTest(t, func(t *testing.T, e pgTestEnv) {
		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_possession_proven (set_id, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?)",
			1, 200, 0, 0,
		).Error)

		// process — proof set missing, row must be retained
		require.NoError(t, processPossessionProven(e.ctx, e.db))
		require.EqualValues(t, 1, e.shovelCount(t, "pdp_possession_proven"))

		// create the proof set, retry
		e.setupFixtures(t)
		e.insertDeal(t, model.DealPublished)
		require.NoError(t, processPossessionProven(e.ctx, e.db))
		require.EqualValues(t, 0, e.shovelCount(t, "pdp_possession_proven"))

		var ps model.PDPProofSet
		require.NoError(t, e.db.Where("set_id = ?", 1).First(&ps).Error)
		require.True(t, ps.IsLive)
	})
}

func TestProcessPossessionProven_DoesNotResurrectExpired(t *testing.T) {
	pgTest(t, func(t *testing.T, e pgTestEnv) {
		e.setupFixtures(t)
		e.insertDeal(t, model.DealPublished)
		e.insertDeal(t, model.DealExpired)

		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_possession_proven (set_id, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?)",
			1, 200, 0, 0,
		).Error)
		require.NoError(t, processPossessionProven(e.ctx, e.db))

		var deals []model.Deal
		require.NoError(t, e.db.Where("deal_type = ?", model.DealTypePDP).Order("id").Find(&deals).Error)
		require.Len(t, deals, 2)
		require.Equal(t, model.DealActive, deals[0].State)
		require.Equal(t, model.DealExpired, deals[1].State)
	})
}

func TestProcessDataSetDeleted(t *testing.T) {
	pgTest(t, func(t *testing.T, e pgTestEnv) {
		e.setupFixtures(t)
		e.insertDeal(t, model.DealActive)

		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_dataset_deleted (set_id, deleted_leaf_count, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?, ?)",
			1, 10, 200, 0, 0,
		).Error)

		require.NoError(t, processDataSetDeleted(e.ctx, e.db))

		var ps model.PDPProofSet
		require.NoError(t, e.db.Where("set_id = ?", 1).First(&ps).Error)
		require.True(t, ps.Deleted)

		var deal model.Deal
		require.NoError(t, e.db.Where("deal_type = ?", model.DealTypePDP).First(&deal).Error)
		require.Equal(t, model.DealExpired, deal.State)
		require.EqualValues(t, 0, e.shovelCount(t, "pdp_dataset_deleted"))
	})
}

func TestProcessSPChanged(t *testing.T) {
	pgTest(t, func(t *testing.T, e pgTestEnv) {
		e.setupFixtures(t)
		e.insertDeal(t, model.DealActive, func(d *model.Deal) {
			d.Provider = e.providerFil.String()
		})

		newSP := common.HexToAddress("0x3333333333333333333333333333333333333333")
		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_sp_changed (set_id, old_sp, new_sp, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?, ?, ?)",
			1, e.providerEth.Bytes(), newSP.Bytes(), 200, 0, 0,
		).Error)

		require.NoError(t, processSPChanged(e.ctx, e.db))

		expectedNewSP, _ := commonToDelegatedAddress(newSP)

		var ps model.PDPProofSet
		require.NoError(t, e.db.Where("set_id = ?", 1).First(&ps).Error)
		require.Equal(t, expectedNewSP.String(), ps.Provider)

		var deal model.Deal
		require.NoError(t, e.db.Where("deal_type = ?", model.DealTypePDP).First(&deal).Error)
		require.Equal(t, expectedNewSP.String(), deal.Provider)
		require.EqualValues(t, 0, e.shovelCount(t, "pdp_sp_changed"))
	})
}

func TestDeleteByKeys(t *testing.T) {
	pgTest(t, func(t *testing.T, e pgTestEnv) {
		insert := func(setID int, blockNum int64, txIdx, logIdx int) {
			require.NoError(t, e.db.Exec(
				"INSERT INTO pdp_dataset_created (set_id, storage_provider, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?, ?)",
				setID, []byte{0x01}, blockNum, txIdx, logIdx,
			).Error)
		}
		remaining := func() []uint64 {
			type r struct {
				SetID uint64 `gorm:"column:set_id"`
			}
			var rows []r
			require.NoError(t, e.db.Raw("SELECT set_id FROM pdp_dataset_created ORDER BY set_id").Scan(&rows).Error)
			out := make([]uint64, len(rows))
			for i, row := range rows {
				out[i] = row.SetID
			}
			return out
		}
		clear := func() { require.NoError(t, e.db.Exec("DELETE FROM pdp_dataset_created").Error) }

		// delete specific keys, leave others
		insert(1, 100, 0, 0)
		insert(2, 100, 0, 1)
		insert(3, 101, 0, 0)
		require.NoError(t, deleteByKeys(e.db, "pdp_dataset_created", []eventKey{
			{100, 0, 0}, {101, 0, 0},
		}))
		require.Equal(t, []uint64{2}, remaining())
		clear()

		// empty keys → no-op
		insert(1, 100, 0, 0)
		require.NoError(t, deleteByKeys(e.db, "pdp_dataset_created", nil))
		require.Equal(t, []uint64{1}, remaining())
		clear()

		// delete all claimed
		insert(10, 200, 0, 0)
		insert(20, 200, 1, 0)
		insert(30, 201, 0, 0)
		require.NoError(t, deleteByKeys(e.db, "pdp_dataset_created", []eventKey{
			{200, 0, 0}, {200, 1, 0}, {201, 0, 0},
		}))
		require.Empty(t, remaining())
	})
}

func TestPiecesRetainedWhenProofSetMissing(t *testing.T) {
	pgTest(t, func(t *testing.T, e pgTestEnv) {
		// only create wallet, no proof set yet
		require.NoError(t, e.db.Create(&model.Wallet{
			ID: "f0100", Address: e.listenerFil.String(),
		}).Error)

		// PiecesAdded arrives before DataSetCreated is processed
		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_pieces_added (set_id, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?)",
			1, 200, 0, 0,
		).Error)

		// process — proof set missing, rows must be retained
		require.NoError(t, processPiecesChanged(e.ctx, e.db, e.client))
		require.EqualValues(t, 1, e.shovelCount(t, "pdp_pieces_added"))

		var dealCount int64
		require.NoError(t, e.db.Model(&model.Deal{}).Where("deal_type = ?", model.DealTypePDP).Count(&dealCount).Error)
		require.EqualValues(t, 0, dealCount)

		// DataSetCreated succeeds
		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_dataset_created (set_id, storage_provider, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?, ?)",
			1, e.providerEth.Bytes(), 100, 0, 0,
		).Error)
		require.NoError(t, processDataSetCreated(e.ctx, e.db, e.client))

		// retry — proof set exists now, pieces processed
		require.NoError(t, processPiecesChanged(e.ctx, e.db, e.client))
		require.EqualValues(t, 0, e.shovelCount(t, "pdp_pieces_added"))

		require.NoError(t, e.db.Model(&model.Deal{}).Where("deal_type = ?", model.DealTypePDP).Count(&dealCount).Error)
		require.EqualValues(t, 1, dealCount)
	})
}

func TestProcessNewEvents_EmptyTables(t *testing.T) {
	pgTest(t, func(t *testing.T, e pgTestEnv) {
		require.NoError(t, processNewEvents(e.ctx, e.db, e.client))
	})
}

func TestProcessNewEvents_FullLifecycle(t *testing.T) {
	pgTest(t, func(t *testing.T, e pgTestEnv) {
		require.NoError(t, e.db.Create(&model.Wallet{
			ID: "f0100", Address: e.listenerFil.String(),
		}).Error)

		// step 1: DataSetCreated
		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_dataset_created (set_id, storage_provider, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?, ?)",
			1, e.providerEth.Bytes(), 100, 0, 0,
		).Error)
		require.NoError(t, processNewEvents(e.ctx, e.db, e.client))

		var psCount int64
		require.NoError(t, e.db.Model(&model.PDPProofSet{}).Count(&psCount).Error)
		require.EqualValues(t, 1, psCount)

		// step 2: PiecesAdded
		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_pieces_added (set_id, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?)",
			1, 200, 0, 0,
		).Error)
		require.NoError(t, processNewEvents(e.ctx, e.db, e.client))

		var dealCount int64
		require.NoError(t, e.db.Model(&model.Deal{}).Where("deal_type = ?", model.DealTypePDP).Count(&dealCount).Error)
		require.EqualValues(t, 1, dealCount)

		// step 3: PossessionProven → active
		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_possession_proven (set_id, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?)",
			1, 300, 0, 0,
		).Error)
		require.NoError(t, processNewEvents(e.ctx, e.db, e.client))

		var deal model.Deal
		require.NoError(t, e.db.Where("deal_type = ?", model.DealTypePDP).First(&deal).Error)
		require.Equal(t, model.DealActive, deal.State)
		require.True(t, *deal.ProofSetLive)

		// step 4: DataSetDeleted → expired
		require.NoError(t, e.db.Exec(
			"INSERT INTO pdp_dataset_deleted (set_id, deleted_leaf_count, block_num, tx_idx, log_idx) VALUES (?, ?, ?, ?, ?)",
			1, 1, 400, 0, 0,
		).Error)
		require.NoError(t, processNewEvents(e.ctx, e.db, e.client))

		require.NoError(t, e.db.Where("deal_type = ?", model.DealTypePDP).First(&deal).Error)
		require.Equal(t, model.DealExpired, deal.State)

		var ps model.PDPProofSet
		require.NoError(t, e.db.Where("set_id = ?", 1).First(&ps).Error)
		require.True(t, ps.Deleted)
	})
}
