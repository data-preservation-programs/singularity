package deal

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type MockDealMaker struct {
	mock.Mock
}

func (m *MockDealMaker) MakeDeal(ctx context.Context, actorObj model.Actor, car model.Car, dealConfig replication.DealConfig, signer replication.ProposalSigner) (*model.Deal, error) {
	args := m.Called(ctx, actorObj, car, dealConfig)
	return args.Get(0).(*model.Deal), args.Error(1)
}

var proposal = Proposal{
	HTTPHeaders:   []string{"a=b"},
	URLTemplate:   "http://127.0.0.1/piece",
	RootCID:       "bafy2bzaceczlclcg4notjmrz4ayenf7fi4mngnqbgjs27r3resyhzwxjnviay",
	Verified:      true,
	IPNI:          true,
	KeepUnsealed:  true,
	StartDelay:    "24h",
	Duration:      "2400h",
	ClientAddress: "f01000",
	ProviderID:    "f01001",
	PieceCID:      "baga6ea4seaqdyupo27fj2fk2mtefzlxvrbf6kdi4twdpccdzbyqrbpsvfsh5ula",
	PieceSize:     "32GiB",
	FileSize:      1000,
}

// creates a wallet (and optionally an actor) that matches proposal.ClientAddress
func createTestWalletAndActor(t *testing.T, db *gorm.DB, withActor bool) {
	t.Helper()
	actorID := "f01000"
	if withActor {
		require.NoError(t, db.Create(&model.Actor{ID: actorID, Address: "f01000"}).Error)
	}
	w := model.Wallet{
		Address: "f01000", KeyPath: "/tmp/key-manual", KeyStore: "local",
	}
	if withActor {
		w.ActorID = &actorID
	}
	require.NoError(t, db.Create(&w).Error)
}

func TestSendManualHandler_WalletNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		mockDealMaker := new(MockDealMaker)
		mockDealMaker.On("MakeDeal", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
		_, err := Default.SendManualHandler(ctx, db, nil, nil, mockDealMaker, proposal)
		require.ErrorIs(t, err, handlererror.ErrNotFound)
		require.ErrorContains(t, err, "wallet")
	})
}

func TestSendManualHandler_InvalidPieceCID(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		createTestWalletAndActor(t, db, false)
		mockDealMaker := new(MockDealMaker)
		badProposal := proposal
		badProposal.PieceCID = "bad"
		_, err := Default.SendManualHandler(ctx, db, nil, nil, mockDealMaker, badProposal)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid piece CID")
	})
}

func TestSendManualHandler_InvalidPieceCID_NOTCOMMP(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		createTestWalletAndActor(t, db, false)
		mockDealMaker := new(MockDealMaker)
		badProposal := proposal
		badProposal.PieceCID = proposal.RootCID
		_, err := Default.SendManualHandler(ctx, db, nil, nil, mockDealMaker, badProposal)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "must be commp")
	})
}

func TestSendManualHandler_InvalidPieceSize(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		createTestWalletAndActor(t, db, false)
		mockDealMaker := new(MockDealMaker)
		badProposal := proposal
		badProposal.PieceSize = "aaa"
		_, err := Default.SendManualHandler(ctx, db, nil, nil, mockDealMaker, badProposal)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid piece size")
	})
}

func TestSendManualHandler_InvalidPieceSize_NotPowerOfTwo(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		createTestWalletAndActor(t, db, false)
		mockDealMaker := new(MockDealMaker)
		badProposal := proposal
		badProposal.PieceSize = "31GiB"
		_, err := Default.SendManualHandler(ctx, db, nil, nil, mockDealMaker, badProposal)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "must be a power of 2")
	})
}

func TestSendManualHandler_InvalidRootCID(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		createTestWalletAndActor(t, db, false)
		mockDealMaker := new(MockDealMaker)
		badProposal := proposal
		badProposal.RootCID = "xxxx"
		_, err := Default.SendManualHandler(ctx, db, nil, nil, mockDealMaker, badProposal)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid root CID")
	})
}

func TestSendManualHandler_InvalidDuration(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		createTestWalletAndActor(t, db, false)
		mockDealMaker := new(MockDealMaker)
		badProposal := proposal
		badProposal.Duration = "xxxx"
		_, err := Default.SendManualHandler(ctx, db, nil, nil, mockDealMaker, badProposal)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid duration")
	})
}

func TestSendManualHandler_InvalidStartDelay(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		createTestWalletAndActor(t, db, false)
		mockDealMaker := new(MockDealMaker)
		badProposal := proposal
		badProposal.StartDelay = "xxxx"
		_, err := Default.SendManualHandler(ctx, db, nil, nil, mockDealMaker, badProposal)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid start delay")
	})
}

func TestSendManualHandler(t *testing.T) {
	actorID := "f01000"
	actor := model.Actor{
		ID:      actorID,
		Address: "f01000",
	}

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		createTestWalletAndActor(t, db, true)

		mockDealMaker := new(MockDealMaker)
		mockDealMaker.On("MakeDeal", mock.Anything, actor, mock.Anything, replication.DealConfig{
			Provider:        proposal.ProviderID,
			StartDelay:      24 * time.Hour,
			Duration:        2400 * time.Hour,
			Verified:        proposal.Verified,
			HTTPHeaders:     map[string]string{"a": "b"},
			URLTemplate:     proposal.URLTemplate,
			KeepUnsealed:    proposal.KeepUnsealed,
			AnnounceToIPNI:  proposal.IPNI,
			PricePerDeal:    proposal.PricePerDeal,
			PricePerGB:      proposal.PricePerGB,
			PricePerGBEpoch: proposal.PricePerGBEpoch,
		}).Return(&model.Deal{}, nil)
		// lotusClient is nil — GetOrCreateActor won't call lotus because ActorID is already set
		resp, err := Default.SendManualHandler(ctx, db, nil, nil, mockDealMaker, proposal)
		mockDealMaker.AssertExpectations(t)
		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}
