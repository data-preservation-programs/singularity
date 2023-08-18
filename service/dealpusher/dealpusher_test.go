package dealpusher

import (
	"bytes"
	"context"
	"crypto/rand"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/service/epochutil"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockDealMaker struct {
	mock.Mock
}

func (m *MockDealMaker) MakeDeal(ctx context.Context, walletObj model.Wallet, car model.Car, dealConfig replication.DealConfig) (*model.Deal, error) {
	args := m.Called(ctx, walletObj, car, dealConfig)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	deal := *args.Get(0).(*model.Deal)
	deal.ID = 0
	deal.PieceCID = car.PieceCID
	deal.PieceSize = car.PieceSize
	deal.ClientID = walletObj.ID
	deal.Provider = dealConfig.Provider
	deal.Verified = dealConfig.Verified
	deal.ProposalID = uuid.NewString()
	deal.State = model.DealProposed
	now := time.Now()
	startEpoch := epochutil.TimeToEpoch(now.Add(dealConfig.StartDelay))
	endEpoch := epochutil.TimeToEpoch(now.Add(dealConfig.StartDelay + dealConfig.Duration))
	if deal.StartEpoch == 0 {
		deal.StartEpoch = int32(startEpoch)
	}
	if deal.EndEpoch == 0 {
		deal.EndEpoch = int32(endEpoch)
	}
	err := args.Error(1)
	if err != nil {
		return &deal, errors.WithStack(err)
	}
	return &deal, nil
}

func TestDealMakerService_Start(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	service, err := NewDealPusher(db, "https://api.node.glif.io", "", 1)
	require.NoError(t, err)
	ctx, cancel := context.WithCancel(context.Background())
	dones, _, err := service.Start(ctx)
	require.NoError(t, err)
	time.Sleep(time.Second)
	cancel()
	for _, done := range dones {
		<-done
	}
}

func TestDealMakerService_MultipleInstances(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	service1, err := NewDealPusher(db, "https://api.node.glif.io", "", 1)
	require.NoError(t, err)
	service2, err := NewDealPusher(db, "https://api.node.glif.io", "", 1)
	require.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	dones1, _, err := service1.Start(ctx)
	require.NoError(t, err)
	dones2, _, err2 := service2.Start(ctx)
	require.ErrorIs(t, err2, context.DeadlineExceeded)
	for _, done := range dones1 {
		<-done
	}
	for _, done := range dones2 {
		<-done
	}
}

func TestDealMakerService_FailtoSend(t *testing.T) {
	waitPendingInterval = 100 * time.Millisecond
	defer func() {
		waitPendingInterval = time.Second
	}()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	service, err := NewDealPusher(db, "https://api.node.glif.io", "", 2)
	require.NoError(t, err)
	mockDealmaker := new(MockDealMaker)
	service.dealMaker = mockDealmaker
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	provider := "f0miner"
	client := "f0client"
	schedule := model.Schedule{
		Preparation: &model.Preparation{
			SourceStorages: []model.Storage{{}},
			Name:           "test", Wallets: []model.Wallet{
				{
					ID: client, Address: "f0xx",
				},
			}},
		State:                model.ScheduleActive,
		Provider:             provider,
		MaxPendingDealNumber: 2,
		MaxPendingDealSize:   2048,
		TotalDealNumber:      4,
	}
	err = db.Create(&schedule).Error
	require.NoError(t, err)
	mockDealmaker.On("MakeDeal", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("send deal error"))
	pieceCIDs := []model.CID{
		model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
	}
	err = db.Create([]model.Car{
		{
			AttachmentID: 1,
			PieceCID:     pieceCIDs[0],
			PieceSize:    1024,
			FilePath:     "0",
		},
	}).Error
	require.NoError(t, err)
	service.runOnce(ctx)
	time.Sleep(3 * time.Second)
	schedule = model.Schedule{}
	err = db.First(&schedule).Error
	require.NoError(t, err)
	require.Equal(t, model.ScheduleError, schedule.State)
	require.Contains(t, schedule.ErrorMessage, "#2: send deal error")
}

func TestDealMakerService_Cron(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	service, err := NewDealPusher(db, "https://api.node.glif.io", "", 1)
	require.NoError(t, err)
	mockDealmaker := new(MockDealMaker)
	service.dealMaker = mockDealmaker
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// All deal proposal will be accepted
	// Create test schedule
	provider := "f0miner"
	client := "f0client"
	schedule := model.Schedule{
		Preparation: &model.Preparation{
			SourceStorages: []model.Storage{{}},
			Name:           "test", Wallets: []model.Wallet{
				{
					ID: client, Address: "f0xx",
				},
			}},
		State:            model.ScheduleActive,
		ScheduleCron:     "0 0 1 1 1",
		ScheduleDealSize: 1,
		Provider:         provider,
	}
	err = db.Create(&schedule).Error
	require.NoError(t, err)

	mockDealmaker.On("MakeDeal", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Deal{
		ScheduleID: &schedule.ID,
	}, nil)

	err = db.Create([]model.Car{
		{
			AttachmentID: 1,
			PieceCID:     model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
			PieceSize:    1024,
		},
		{
			AttachmentID: 1,
			PieceCID:     model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
			PieceSize:    1024,
		},
		{
			AttachmentID: 1,
			PieceCID:     model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
			PieceSize:    1024,
		},
	}).Error

	require.NoError(t, err)
	service.cron.Start()
	defer service.cron.Stop()
	service.runOnce(ctx)

	// Update to a new cron schedule
	err = db.Model(&schedule).Update("schedule_cron", "* * * * * *").Error
	require.NoError(t, err)
	service.runOnce(ctx)

	time.Sleep(1 * time.Second)
	var deals []model.Deal
	err = db.Find(&deals).Error
	require.NoError(t, err)
	ndeals := len(deals)
	require.True(t, ndeals == 1 || ndeals == 2)

	// Pause the cron schedule
	err = db.Model(&schedule).Update("state", model.SchedulePaused).Error
	require.NoError(t, err)
	service.runOnce(ctx)
	time.Sleep(2 * time.Second)
	err = db.Find(&deals).Error
	require.NoError(t, err)
	require.Len(t, deals, ndeals)

	// Resume the cron schedule
	db.Model(&schedule).Update("state", model.ScheduleActive)
	service.runOnce(ctx)
	time.Sleep(3 * time.Second)
	err = db.Find(&deals).Error
	require.NoError(t, err)
	require.Greater(t, len(deals), ndeals)
}

func TestDealMakerService_ScheduleWithConstraints(t *testing.T) {
	waitPendingInterval = 100 * time.Millisecond
	defer func() {
		waitPendingInterval = time.Second
	}()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	service, err := NewDealPusher(db, "https://api.node.glif.io", "", 1)
	require.NoError(t, err)
	mockDealmaker := new(MockDealMaker)
	service.dealMaker = mockDealmaker
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	provider := "f0miner"
	client := "f0client"
	schedule := model.Schedule{
		Preparation: &model.Preparation{
			SourceStorages: []model.Storage{{}},
			Name:           "test", Wallets: []model.Wallet{
				{
					ID: client, Address: "f0xx",
				},
			}},
		State:                model.ScheduleActive,
		Provider:             provider,
		MaxPendingDealNumber: 2,
		MaxPendingDealSize:   2048,
		TotalDealNumber:      4,
	}
	err = db.Create(&schedule).Error
	require.NoError(t, err)
	mockDealmaker.On("MakeDeal", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Deal{
		ScheduleID: &schedule.ID,
	}, nil)
	pieceCIDs := []model.CID{
		model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
		model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
		model.CID(calculateCommp(t, generateRandomBytes(1000), 2048)),
		model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
		model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
	}
	err = db.Create([]model.Car{
		{
			AttachmentID: 1,
			PieceCID:     pieceCIDs[0],
			PieceSize:    1024,
			FilePath:     "0",
		},
		{
			AttachmentID: 1,
			PieceCID:     pieceCIDs[1],
			PieceSize:    1024,
			FilePath:     "1",
		},
		{
			AttachmentID: 1,
			PieceCID:     pieceCIDs[2],
			PieceSize:    2048,
			FilePath:     "2",
		},
		{
			AttachmentID: 1,
			PieceCID:     pieceCIDs[3],
			PieceSize:    1024,
			FilePath:     "3",
		},
		{
			AttachmentID: 1,
			PieceCID:     pieceCIDs[4],
			PieceSize:    1024,
			FilePath:     "4",
		},
	}).Error
	require.NoError(t, err)

	service.runOnce(ctx)
	time.Sleep(time.Second)
	var deals []model.Deal
	err = db.Find(&deals).Error
	require.NoError(t, err)
	require.Len(t, deals, 2)

	err = db.Model(&deals).Update("state", model.DealActive).Error
	require.NoError(t, err)
	time.Sleep(time.Second)
	err = db.Find(&deals).Error
	require.NoError(t, err)
	require.Len(t, deals, 3)

	err = db.Model(&deals).Update("state", model.DealActive).Error
	require.NoError(t, err)
	time.Sleep(time.Second)
	err = db.Find(&deals).Error
	require.NoError(t, err)
	require.Len(t, deals, 4)

	err = db.Model(&schedule).Update("state", model.ScheduleActive).
		Update("total_deal_size", 4096).
		Update("total_deal_number", 5).
		Error
	require.NoError(t, err)
	service.runOnce(ctx)
	time.Sleep(time.Second)
	err = db.Find(&deals).Error
	require.NoError(t, err)
	require.Len(t, deals, 4)
}

func TestDealMakerService_NewScheduleOneOff(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	service, err := NewDealPusher(db, "https://api.node.glif.io", "", 1)
	require.NoError(t, err)
	mockDealmaker := new(MockDealMaker)
	service.dealMaker = mockDealmaker
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// All deal proposal will be accepted
	// Create test schedule
	provider := "f0miner"
	client := "f0client"
	schedule := model.Schedule{
		Preparation: &model.Preparation{
			Name: "test", Wallets: []model.Wallet{
				{
					ID: client, Address: "f0xx",
				},
			},
			SourceStorages: []model.Storage{{}},
		},
		State:    model.ScheduleActive,
		Provider: provider,
	}
	err = db.Create(&schedule).Error
	require.NoError(t, err)

	mockDealmaker.On("MakeDeal", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Deal{
		ScheduleID: &schedule.ID,
	}, nil)

	pieceCIDs := []model.CID{
		model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
		model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
		model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
		model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
		model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
	}
	err = db.Create([]model.Car{
		{
			AttachmentID: 1,
			PieceCID:     pieceCIDs[0],
			PieceSize:    1024,
			FilePath:     "0",
		},
		{
			AttachmentID: 1,
			PieceCID:     pieceCIDs[1],
			PieceSize:    1024,
			FilePath:     "1",
		},
		{
			AttachmentID: 1,
			PieceCID:     pieceCIDs[2],
			PieceSize:    1024,
			FilePath:     "2",
		},
		{
			AttachmentID: 1,
			PieceCID:     pieceCIDs[3],
			PieceSize:    1024,
			FilePath:     "3",
		},
		{
			AttachmentID: 1,
			PieceCID:     pieceCIDs[4],
			PieceSize:    1024,
			FilePath:     "4",
		},
	}).Error
	require.NoError(t, err)

	// Test1 is already proposed
	// Test2 is expired proposal
	// Test3 is active
	// Test4 is proposed by other schedules
	// Test5 is not proposed
	err = db.Create([]model.Deal{
		{
			ScheduleID: &schedule.ID,
			Provider:   provider,
			ClientID:   client,
			PieceCID:   pieceCIDs[0],
			PieceSize:  1024,
			State:      model.DealProposed,
		},
		{
			ScheduleID: &schedule.ID,
			Provider:   provider,
			ClientID:   client,
			PieceCID:   pieceCIDs[1],
			PieceSize:  1024,
			State:      model.DealProposalExpired,
		},
		{
			ScheduleID: &schedule.ID,
			Provider:   provider,
			ClientID:   client,
			PieceCID:   pieceCIDs[2],
			PieceSize:  1024,
			State:      model.DealActive,
		},
		{
			Provider:  provider,
			ClientID:  client,
			PieceCID:  pieceCIDs[3],
			PieceSize: 1024,
			State:     model.DealProposed,
		},
	}).Error
	require.NoError(t, err)
	service.runOnce(ctx)
	time.Sleep(time.Second)
	var deals []model.Deal
	err = db.Find(&deals).Error
	require.NoError(t, err)
	require.Len(t, deals, 6)
	require.Equal(t, pieceCIDs[1], deals[4].PieceCID)
	require.Equal(t, pieceCIDs[4], deals[5].PieceCID)
}

func calculateCommp(t *testing.T, content []byte, targetPieceSize uint64) cid.Cid {
	calc := &commp.Calc{}
	_, err := bytes.NewBuffer(content).WriteTo(calc)
	require.NoError(t, err)
	c, _, err := pack.GetCommp(calc, targetPieceSize)
	require.NoError(t, err)
	return c
}

func generateRandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}
