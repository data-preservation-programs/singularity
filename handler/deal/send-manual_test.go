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

func (m *MockDealMaker) MakeDeal(ctx context.Context, walletObj model.Wallet, car model.Car, dealConfig replication.DealConfig) (*model.Deal, error) {
	args := m.Called(ctx, walletObj, car, dealConfig)
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

func TestSendManualHandler_WalletNotFound(t *testing.T) {
	wallet := model.Wallet{
		ActorID: "f09999",
		Address: "f10000",
	}

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&wallet).Error
		require.NoError(t, err)

		mockDealMaker := new(MockDealMaker)
		mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
		_, err = Default.SendManualHandler(ctx, db, mockDealMaker, proposal)
		require.ErrorIs(t, err, handlererror.ErrNotFound)
		require.ErrorContains(t, err, "client address")
	})
}

func TestSendManualHandler_InvalidPieceCID(t *testing.T) {
	wallet := model.Wallet{
		ActorID: "f01000",
		Address: "f10000",
	}

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&wallet).Error
		require.NoError(t, err)

		mockDealMaker := new(MockDealMaker)
		mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
		badProposal := proposal
		badProposal.PieceCID = "bad"
		_, err = Default.SendManualHandler(ctx, db, mockDealMaker, badProposal)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid piece CID")
	})
}

func TestSendManualHandler_InvalidPieceCID_NOTCOMMP(t *testing.T) {
	wallet := model.Wallet{
		ActorID: "f01000",
		Address: "f10000",
	}

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&wallet).Error
		require.NoError(t, err)

		mockDealMaker := new(MockDealMaker)
		mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
		badProposal := proposal
		badProposal.PieceCID = proposal.RootCID
		_, err = Default.SendManualHandler(ctx, db, mockDealMaker, badProposal)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "must be commp")
	})
}

func TestSendManualHandler_InvalidPieceSize(t *testing.T) {
	wallet := model.Wallet{
		ActorID: "f01000",
		Address: "f10000",
	}

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&wallet).Error
		require.NoError(t, err)

		mockDealMaker := new(MockDealMaker)
		mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
		badProposal := proposal
		badProposal.PieceSize = "aaa"
		_, err = Default.SendManualHandler(ctx, db, mockDealMaker, badProposal)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid piece size")
	})
}

func TestSendManualHandler_InvalidPieceSize_NotPowerOfTwo(t *testing.T) {
	wallet := model.Wallet{
		ActorID: "f01000",
		Address: "f10000",
	}

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&wallet).Error
		require.NoError(t, err)

		mockDealMaker := new(MockDealMaker)
		mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
		badProposal := proposal
		badProposal.PieceSize = "31GiB"
		_, err = Default.SendManualHandler(ctx, db, mockDealMaker, badProposal)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "must be a power of 2")
	})
}

func TestSendManualHandler_InvalidRootCID(t *testing.T) {
	wallet := model.Wallet{
		ActorID: "f01000",
		Address: "f10000",
	}

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&wallet).Error
		require.NoError(t, err)

		mockDealMaker := new(MockDealMaker)
		mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
		badProposal := proposal
		badProposal.RootCID = "xxxx"
		_, err = Default.SendManualHandler(ctx, db, mockDealMaker, badProposal)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid root CID")
	})
}

func TestSendManualHandler_InvalidDuration(t *testing.T) {
	wallet := model.Wallet{
		ActorID: "f01000",
		Address: "f10000",
	}

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&wallet).Error
		require.NoError(t, err)

		mockDealMaker := new(MockDealMaker)
		mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
		badProposal := proposal
		badProposal.Duration = "xxxx"
		_, err = Default.SendManualHandler(ctx, db, mockDealMaker, badProposal)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid duration")
	})
}

func TestSendManualHandler_InvalidStartDelay(t *testing.T) {
	wallet := model.Wallet{
		ActorID: "f01000",
		Address: "f10000",
	}

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&wallet).Error
		require.NoError(t, err)

		mockDealMaker := new(MockDealMaker)
		mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
		badProposal := proposal
		badProposal.StartDelay = "xxxx"
		_, err = Default.SendManualHandler(ctx, db, mockDealMaker, badProposal)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid start delay")
	})
}

func TestSendManualHandler(t *testing.T) {
	wallet := model.Wallet{
		ActorID: "f01000",
		Address: "f10000",
	}

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&wallet).Error
		require.NoError(t, err)

		mockDealMaker := new(MockDealMaker)
		mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, replication.DealConfig{
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
		resp, err := Default.SendManualHandler(ctx, db, mockDealMaker, proposal)
		mockDealMaker.AssertExpectations(t)
		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}
