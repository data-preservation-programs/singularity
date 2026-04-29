package dealpusher

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/go-synapse/signer"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// proofSetManagerMock satisfies PDPProofSetManager. It records every batch
// it receives so tests can assert on the pieces and provider that were
// pushed.
type proofSetManagerMock struct {
	dataSetID uint64
	calls     []proofSetManagerCall
	err       error
}

type proofSetManagerCall struct {
	provider string
	pieces   []PDPPieceInput
}

func (m *proofSetManagerMock) PullPiecesToFWSS(_ context.Context, _ signer.EVMSigner, provider string, pieces []PDPPieceInput, _ PDPSchedulingConfig) (PDPPullResult, error) {
	if m.err != nil {
		return PDPPullResult{}, m.err
	}
	m.calls = append(m.calls, proofSetManagerCall{
		provider: provider,
		pieces:   append([]PDPPieceInput(nil), pieces...),
	})
	return PDPPullResult{DataSetID: m.dataSetID}, nil
}

func TestDealPusher_ResolveScheduleDealType_EmptyReturnsEmpty(t *testing.T) {
	d := &DealPusher{}
	require.Equal(t, model.DealType(""), d.resolveScheduleDealType(&model.Schedule{}))
}

func TestDealPusher_ResolveScheduleDealType_UsesScheduleDealType(t *testing.T) {
	d := &DealPusher{}
	require.Equal(t, model.DealTypePDP, d.resolveScheduleDealType(&model.Schedule{DealType: model.DealTypePDP}))
}

func TestDealPusher_RunSchedule_UnknownDealTypeErrors(t *testing.T) {
	d := &DealPusher{}
	state, err := d.runSchedule(context.Background(), &model.Schedule{DealType: "bogus"})
	require.Error(t, err)
	require.Equal(t, model.ScheduleError, state)
	require.Contains(t, err.Error(), "unknown deal type")
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

func TestDealPusher_RunSchedule_PDPPushesBatchAndCreatesDeals(t *testing.T) {
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

		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)
		keyPath, _, err := ks.Put(testutil.TestPrivateKeyHex)
		require.NoError(t, err)

		actorID := "f01001"
		require.NoError(t, db.Create(&model.Actor{ID: actorID, Address: clientAddr.String()}).Error)
		require.NoError(t, db.Create(&model.Actor{ID: "f01002", Address: providerAddr.String()}).Error)

		walletObj := model.Wallet{
			Address:  clientAddr.String(),
			KeyPath:  keyPath,
			KeyStore: "local",
			ActorID:  &actorID,
		}
		require.NoError(t, db.Create(&walletObj).Error)
		require.NoError(t, db.Model(&prep).Update("wallet_id", walletObj.ID).Error)

		storageRow := model.Storage{Name: "src-storage"}
		require.NoError(t, db.Create(&storageRow).Error)
		attachment := model.SourceAttachment{PreparationID: prep.ID, StorageID: storageRow.ID}
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
		schedule.Preparation = &model.Preparation{Wallet: &walletObj}

		psm := &proofSetManagerMock{dataSetID: 42}
		d := &DealPusher{
			dbNoContext:              db,
			keyStore:                 ks,
			pdpProofSetManager:       psm,
			pdpSchedulingConfig:      defaultPDPSchedulingConfig(),
			scheduleDealTypeResolver: func(_ *model.Schedule) model.DealType { return model.DealTypePDP },
		}

		state, err := d.runSchedule(ctx, &schedule)
		require.NoError(t, err)
		require.Equal(t, model.ScheduleCompleted, state)

		require.Len(t, psm.calls, 1, "expected one batch push")
		require.Equal(t, providerAddr.String(), psm.calls[0].provider)
		require.Len(t, psm.calls[0].pieces, 1)
		require.Equal(t, cid.Cid(pieceCID), psm.calls[0].pieces[0].PieceCID)
		require.Equal(t, int64(1024), psm.calls[0].pieces[0].PieceSize)

		var deals []model.Deal
		require.NoError(t, db.Where("schedule_id = ?", schedule.ID).Find(&deals).Error)
		require.Len(t, deals, 1)
		require.Equal(t, model.DealTypePDP, deals[0].DealType)
		require.Equal(t, model.DealProposed, deals[0].State)
		require.Equal(t, pdpDealEpochSentinel, deals[0].StartEpoch)
		require.Equal(t, pdpDealEpochSentinel, deals[0].EndEpoch)
		require.NotNil(t, deals[0].ProofSetID)
		require.Equal(t, uint64(42), *deals[0].ProofSetID)
		require.NotNil(t, deals[0].WalletID)
		require.Equal(t, walletObj.ID, *deals[0].WalletID)
	})
}

func TestDealPusher_RunSchedule_PDPRejectsInvalidPieceSize(t *testing.T) {
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

		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)
		keyPath, _, err := ks.Put(testutil.TestPrivateKeyHex)
		require.NoError(t, err)

		actorID := "f01001"
		require.NoError(t, db.Create(&model.Actor{ID: actorID, Address: clientAddr.String()}).Error)
		require.NoError(t, db.Create(&model.Actor{ID: "f01002", Address: providerAddr.String()}).Error)
		walletObj := model.Wallet{
			Address:  clientAddr.String(),
			KeyPath:  keyPath,
			KeyStore: "local",
			ActorID:  &actorID,
		}
		require.NoError(t, db.Create(&walletObj).Error)
		require.NoError(t, db.Model(&prep).Update("wallet_id", walletObj.ID).Error)

		storageRow := model.Storage{Name: "src-storage"}
		require.NoError(t, db.Create(&storageRow).Error)
		attachment := model.SourceAttachment{PreparationID: prep.ID, StorageID: storageRow.ID}
		require.NoError(t, db.Create(&attachment).Error)

		pieceCID := model.CID(calculateCommp(t, generateRandomBytes(1000), 1024))
		car := model.Car{
			AttachmentID:  &attachment.ID,
			PreparationID: &prep.ID,
			PieceCID:      pieceCID,
			PieceSize:     1536, // not a power of two
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
		schedule.Preparation = &model.Preparation{Wallet: &walletObj}

		psm := &proofSetManagerMock{dataSetID: 42}
		d := &DealPusher{
			dbNoContext:              db,
			keyStore:                 ks,
			pdpProofSetManager:       psm,
			pdpSchedulingConfig:      defaultPDPSchedulingConfig(),
			scheduleDealTypeResolver: func(_ *model.Schedule) model.DealType { return model.DealTypePDP },
		}

		state, err := d.runSchedule(ctx, &schedule)
		require.Error(t, err)
		require.Equal(t, model.ScheduleError, state)
		require.ErrorContains(t, err, "invalid piece size for piece")
		require.ErrorContains(t, err, "must be a power of two")
		require.Empty(t, psm.calls)

		var deals []model.Deal
		require.NoError(t, db.Where("schedule_id = ?", schedule.ID).Find(&deals).Error)
		require.Empty(t, deals)
	})
}

func TestDealPusher_RunSchedule_PDPRejectsPreparationWithOversizedPiece(t *testing.T) {
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

		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)
		keyPath, _, err := ks.Put(testutil.TestPrivateKeyHex)
		require.NoError(t, err)

		actorID := "f01001"
		require.NoError(t, db.Create(&model.Actor{ID: actorID, Address: clientAddr.String()}).Error)
		require.NoError(t, db.Create(&model.Actor{ID: "f01002", Address: providerAddr.String()}).Error)
		walletObj := model.Wallet{
			Address:  clientAddr.String(),
			KeyPath:  keyPath,
			KeyStore: "local",
			ActorID:  &actorID,
		}
		require.NoError(t, db.Create(&walletObj).Error)
		require.NoError(t, db.Model(&prep).Update("wallet_id", walletObj.ID).Error)

		storageRow := model.Storage{Name: "src-storage"}
		require.NoError(t, db.Create(&storageRow).Error)
		attachment := model.SourceAttachment{PreparationID: prep.ID, StorageID: storageRow.ID}
		require.NoError(t, db.Create(&attachment).Error)

		pieceCID := model.CID(calculateCommp(t, generateRandomBytes(1000), 1024))
		car := model.Car{
			AttachmentID:  &attachment.ID,
			PreparationID: &prep.ID,
			PieceCID:      pieceCID,
			PieceSize:     1 << 31, // 2 GiB, above current 1 GiB minus FR32 overhead PDP limit
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
		schedule.Preparation = &model.Preparation{Wallet: &walletObj}

		psm := &proofSetManagerMock{dataSetID: 42}
		d := &DealPusher{
			dbNoContext:              db,
			keyStore:                 ks,
			pdpProofSetManager:       psm,
			pdpSchedulingConfig:      defaultPDPSchedulingConfig(),
			scheduleDealTypeResolver: func(_ *model.Schedule) model.DealType { return model.DealTypePDP },
		}

		state, err := d.runSchedule(ctx, &schedule)
		require.Error(t, err)
		require.Equal(t, model.ScheduleError, state)
		require.ErrorContains(t, err, "piece limit is 1 GiB minus FR32 overhead")
		require.Empty(t, psm.calls)

		var deals []model.Deal
		require.NoError(t, db.Where("schedule_id = ?", schedule.ID).Find(&deals).Error)
		require.Empty(t, deals)
	})
}
