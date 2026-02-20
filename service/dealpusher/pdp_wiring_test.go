package dealpusher

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type noopPDPProofSetManager struct{}

func (noopPDPProofSetManager) EnsureProofSet(_ context.Context, _ model.Wallet, _ string) (uint64, error) {
	return 1, nil
}

func (noopPDPProofSetManager) QueueAddRoots(_ context.Context, _ uint64, _ []cid.Cid, _ PDPSchedulingConfig) (*PDPQueuedTx, error) {
	return &PDPQueuedTx{Hash: "0x1"}, nil
}

type noopPDPTransactionConfirmer struct{}

func (noopPDPTransactionConfirmer) WaitForConfirmations(_ context.Context, txHash string, _ uint64, _ time.Duration) (*PDPTransactionReceipt, error) {
	return &PDPTransactionReceipt{Hash: txHash}, nil
}

type fixedWalletChooser struct {
	wallet model.Wallet
}

func (c fixedWalletChooser) Choose(_ context.Context, _ []model.Wallet) (model.Wallet, error) {
	return c.wallet, nil
}

type proofSetManagerMock struct {
	proofSetID uint64
	pieceCIDs  []cid.Cid
}

func (m *proofSetManagerMock) EnsureProofSet(_ context.Context, _ model.Wallet, _ string) (uint64, error) {
	return m.proofSetID, nil
}

func (m *proofSetManagerMock) QueueAddRoots(_ context.Context, _ uint64, pieceCIDs []cid.Cid, _ PDPSchedulingConfig) (*PDPQueuedTx, error) {
	m.pieceCIDs = append([]cid.Cid(nil), pieceCIDs...)
	return &PDPQueuedTx{Hash: "0xabc"}, nil
}

type txConfirmerMock struct {
	txHash string
}

func (m *txConfirmerMock) WaitForConfirmations(_ context.Context, txHash string, _ uint64, _ time.Duration) (*PDPTransactionReceipt, error) {
	m.txHash = txHash
	return &PDPTransactionReceipt{Hash: txHash}, nil
}

func TestDealPusher_ResolveScheduleDealType_DefaultsToMarket(t *testing.T) {
	d := &DealPusher{}
	require.Equal(t, model.DealTypeMarket, d.resolveScheduleDealType(&model.Schedule{}))
}

func TestDealPusher_ResolveScheduleDealType_DefaultsToPDPForDelegatedProvider(t *testing.T) {
	subaddr := make([]byte, 20)
	subaddr[19] = 1
	providerAddr, err := address.NewDelegatedAddress(10, subaddr)
	require.NoError(t, err)
	d := &DealPusher{}
	require.Equal(t, model.DealTypePDP, d.resolveScheduleDealType(&model.Schedule{Provider: providerAddr.String()}))
}

func TestDealPusher_RunSchedule_PDPWithoutDependenciesReturnsConfiguredError(t *testing.T) {
	d := &DealPusher{
		scheduleDealTypeResolver: func(_ *model.Schedule) model.DealType {
			return model.DealTypePDP
		},
	}

	state, err := d.runSchedule(context.Background(), &model.Schedule{})
	require.Error(t, err)
	require.Equal(t, model.ScheduleError, state)
	require.Contains(t, err.Error(), "pdp scheduling dependencies are not configured")
}

func TestDealPusher_RunSchedule_PDPWithDependenciesCreatesDealsAfterConfirmation(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		clientSubaddr := make([]byte, 20)
		clientSubaddr[19] = 10
		clientAddr, err := address.NewDelegatedAddress(10, clientSubaddr)
		require.NoError(t, err)
		providerSubaddr := make([]byte, 20)
		providerSubaddr[19] = 20
		providerAddr, err := address.NewDelegatedAddress(10, providerSubaddr)
		require.NoError(t, err)

		prep := model.Preparation{Name: "prep"}
		require.NoError(t, db.Create(&prep).Error)
		require.NotZero(t, prep.ID)
		wallet := model.Wallet{
			ID:         clientAddr.String(),
			Address:    clientAddr.String(),
			PrivateKey: testutil.TestPrivateKeyHex,
		}
		require.NoError(t, db.Create(&wallet).Error)
		require.NoError(t, db.Model(&prep).Association("Wallets").Append(&wallet))
		storage := model.Storage{Name: "src-storage"}
		require.NoError(t, db.Create(&storage).Error)
		require.NotZero(t, storage.ID)
		attachment := model.SourceAttachment{PreparationID: prep.ID, StorageID: storage.ID}
		require.NoError(t, db.Create(&attachment).Error)
		require.NotZero(t, attachment.ID)

		pieceCID := model.CID(calculateCommp(t, generateRandomBytes(1000), 1024))
		car := model.Car{
			AttachmentID:  &attachment.ID,
			PreparationID: &prep.ID,
			PieceCID:      pieceCID,
			PieceSize:     1024,
			StoragePath:   "car-1",
		}
		require.NoError(t, db.Create(&car).Error)

		schedule := model.Schedule{
			PreparationID:   prep.ID,
			State:           model.ScheduleActive,
			Provider:        providerAddr.String(),
			TotalDealNumber: 1,
		}
		require.NoError(t, db.Create(&schedule).Error)
		schedule.Preparation = &model.Preparation{Wallets: []model.Wallet{wallet}}

		psm := &proofSetManagerMock{proofSetID: 42}
		conf := &txConfirmerMock{}
		d := &DealPusher{
			dbNoContext:              db,
			walletChooser:            fixedWalletChooser{wallet: wallet},
			pdpProofSetManager:       psm,
			pdpTxConfirmer:           conf,
			pdpSchedulingConfig:      defaultPDPSchedulingConfig(),
			scheduleDealTypeResolver: func(_ *model.Schedule) model.DealType { return model.DealTypePDP },
		}
		var attachments []model.SourceAttachment
		require.NoError(t, db.Where("preparation_id = ?", schedule.PreparationID).Find(&attachments).Error)
		require.Len(t, attachments, 1)
		overReplicatedCIDs := db.
			Table("deals").
			Select("piece_cid").
			Where("state in ?", []model.DealState{model.DealProposed, model.DealPublished, model.DealActive}).
			Group("piece_cid").
			Having("count(*) >= ?", d.maxReplicas)
		cars, err := d.findPDPCars(ctx, &schedule, attachments, nil, overReplicatedCIDs, d.pdpSchedulingConfig.BatchSize)
		require.NoError(t, err)
		require.Len(t, cars, 1)

		state, err := d.runSchedule(ctx, &schedule)
		require.NoError(t, err)
		require.Equal(t, model.ScheduleCompleted, state)
		require.Equal(t, "0xabc", conf.txHash)
		require.Len(t, psm.pieceCIDs, 1)
		require.Equal(t, cid.Cid(pieceCID), psm.pieceCIDs[0])

		var deals []model.Deal
		require.NoError(t, db.Where("schedule_id = ?", schedule.ID).Find(&deals).Error)
		require.Len(t, deals, 1)
		require.Equal(t, model.DealTypePDP, deals[0].DealType)
		require.Equal(t, model.DealProposed, deals[0].State)
		require.NotNil(t, deals[0].ProofSetID)
		require.Equal(t, uint64(42), *deals[0].ProofSetID)
	})
}
