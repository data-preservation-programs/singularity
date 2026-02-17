package pdptracker

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

type mockPDPClient struct {
	proofSets map[address.Address][]ProofSetInfo
}

func (m *mockPDPClient) GetProofSetsForClient(_ context.Context, clientAddress address.Address) ([]ProofSetInfo, error) {
	return m.proofSets[clientAddress], nil
}

func (m *mockPDPClient) GetProofSetInfo(_ context.Context, _ uint64) (*ProofSetInfo, error) {
	return nil, nil
}

func (m *mockPDPClient) IsProofSetLive(_ context.Context, _ uint64) (bool, error) {
	return false, nil
}

func (m *mockPDPClient) GetNextChallengeEpoch(_ context.Context, _ uint64) (int32, error) {
	return 0, nil
}

func TestPDPTracker_Name(t *testing.T) {
	tracker := NewPDPTracker(nil, time.Minute, "", nil, true)
	require.Equal(t, "PDPTracker", tracker.Name())
}

func TestPDPTracker_RunOnce_UpsertByParsedPieceCID(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		walletSubaddr := make([]byte, 20)
		walletSubaddr[19] = 1
		walletAddr, err := address.NewDelegatedAddress(10, walletSubaddr)
		require.NoError(t, err)

		providerSubaddr := make([]byte, 20)
		providerSubaddr[19] = 2
		providerAddr, err := address.NewDelegatedAddress(10, providerSubaddr)
		require.NoError(t, err)

		err := db.Create(&model.Wallet{
			ID:      "f0100",
			Address: walletAddr.String(),
		}).Error
		require.NoError(t, err)

		const pieceCID = "baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq"
		parsedPieceCID, err := cid.Decode(pieceCID)
		require.NoError(t, err)
		client := &mockPDPClient{
			proofSets: map[address.Address][]ProofSetInfo{
				walletAddr: {
					{
						ProofSetID:         7,
						ClientAddress:      walletAddr,
						ProviderAddress:    providerAddr,
						IsLive:             true,
						NextChallengeEpoch: 10,
						PieceCIDs:          []cid.Cid{parsedPieceCID},
					},
				},
			},
		}

		tracker := NewPDPTracker(db, time.Minute, "", client, true)
		require.NoError(t, tracker.runOnce(ctx))

		var first model.Deal
		err = db.Where("deal_type = ?", model.DealTypePDP).First(&first).Error
		require.NoError(t, err)
		require.Equal(t, model.DealTypePDP, first.DealType)
		require.Equal(t, pieceCID, first.PieceCID.String())
		require.NotNil(t, first.ProofSetID)
		require.EqualValues(t, 7, *first.ProofSetID)
		require.NotNil(t, first.ProofSetLive)
		require.True(t, *first.ProofSetLive)
		require.Equal(t, model.DealActive, first.State)
		require.NotNil(t, first.LastVerifiedAt)

		client.proofSets[walletAddr][0].IsLive = false
		client.proofSets[walletAddr][0].NextChallengeEpoch = 11
		require.NoError(t, tracker.runOnce(ctx))

		var deals []model.Deal
		err = db.Where("deal_type = ?", model.DealTypePDP).Find(&deals).Error
		require.NoError(t, err)
		require.Len(t, deals, 1)
		require.NotNil(t, deals[0].ProofSetLive)
		require.False(t, *deals[0].ProofSetLive)
		require.NotNil(t, deals[0].NextChallengeEpoch)
		require.EqualValues(t, 11, *deals[0].NextChallengeEpoch)
		require.Equal(t, model.DealPublished, deals[0].State)
		require.NotNil(t, deals[0].LastVerifiedAt)
	})
}

func TestPDPTracker_RunOnce_SkipsInvalidPieceCID(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		walletSubaddr := make([]byte, 20)
		walletSubaddr[19] = 3
		walletAddr, err := address.NewDelegatedAddress(10, walletSubaddr)
		require.NoError(t, err)

		providerSubaddr := make([]byte, 20)
		providerSubaddr[19] = 4
		providerAddr, err := address.NewDelegatedAddress(10, providerSubaddr)
		require.NoError(t, err)

		err := db.Create(&model.Wallet{
			ID:      "f0100",
			Address: walletAddr.String(),
		}).Error
		require.NoError(t, err)

		client := &mockPDPClient{
			proofSets: map[address.Address][]ProofSetInfo{
				walletAddr: {
					{
						ProofSetID:         7,
						ClientAddress:      walletAddr,
						ProviderAddress:    providerAddr,
						IsLive:             true,
						NextChallengeEpoch: 10,
						PieceCIDs:          []cid.Cid{cid.Undef},
					},
				},
			},
		}
		tracker := NewPDPTracker(db, time.Minute, "", client, true)
		require.NoError(t, tracker.runOnce(ctx))

		var count int64
		err = db.Model(&model.Deal{}).Where("deal_type = ?", model.DealTypePDP).Count(&count).Error
		require.NoError(t, err)
		require.EqualValues(t, 0, count)
	})
}
