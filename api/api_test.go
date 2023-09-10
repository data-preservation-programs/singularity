package api

import (
	"context"
	"fmt"
	"io"
	"net"
	http2 "net/http"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/client/swagger/http"
	deal2 "github.com/data-preservation-programs/singularity/client/swagger/http/deal"
	"github.com/data-preservation-programs/singularity/client/swagger/http/deal_schedule"
	file2 "github.com/data-preservation-programs/singularity/client/swagger/http/file"
	job2 "github.com/data-preservation-programs/singularity/client/swagger/http/job"
	"github.com/data-preservation-programs/singularity/client/swagger/http/piece"
	"github.com/data-preservation-programs/singularity/client/swagger/http/preparation"
	"github.com/data-preservation-programs/singularity/client/swagger/models"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/singularity/handler/file"
	"github.com/data-preservation-programs/singularity/handler/job"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ipfs/go-log/v2"
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

const apiBind = "127.0.0.1:9092"

type MockDealMaker struct {
	mock.Mock
}

func (m *MockDealMaker) MakeDeal(ctx context.Context, walletObj model.Wallet, car model.Car, dealConfig replication.DealConfig) (*model.Deal, error) {
	args := m.Called(ctx, walletObj, car, dealConfig)
	return args.Get(0).(*model.Deal), args.Error(1)
}

func setupMockDataPrep() dataprep.Handler {
	m := new(dataprep.MockDataPrep)
	m.On("ListSchedulesHandler", mock.Anything, mock.Anything, "id").
		Return([]model.Schedule{{}}, nil)
	m.On("CreatePreparationHandler", mock.Anything, mock.Anything, mock.Anything).
		Return(&model.Preparation{}, nil)
	m.On("ExploreHandler", mock.Anything, mock.Anything, "id", "name", "path").
		Return(&dataprep.ExploreResult{}, nil)
	m.On("ListHandler", mock.Anything, mock.Anything).
		Return([]model.Preparation{{}}, nil)
	m.On("AddOutputStorageHandler", mock.Anything, mock.Anything, "id", "name").
		Return(&model.Preparation{}, nil)
	m.On("RemoveOutputStorageHandler", mock.Anything, mock.Anything, "id", "name").
		Return(&model.Preparation{}, nil)
	m.On("ListPiecesHandler", mock.Anything, mock.Anything, "id").
		Return([]dataprep.PieceList{{}}, nil)
	m.On("AddPieceHandler", mock.Anything, mock.Anything, "id", mock.Anything).
		Return(&model.Car{}, nil)
	m.On("AddSourceStorageHandler", mock.Anything, mock.Anything, "id", "name").
		Return(&model.Preparation{}, nil)
	return m
}

func setupMockDeal() deal.Handler {
	m := new(deal.MockDeal)
	m.On("ListHandler", mock.Anything, mock.Anything, mock.Anything).
		Return([]model.Deal{{}}, nil)
	m.On("SendManualHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(&model.Deal{}, nil)
	return m
}

func setupMockSchedule() schedule.Handler {
	m := new(schedule.MockSchedule)
	m.On("CreateHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(&model.Schedule{}, nil)
	m.On("ListHandler", mock.Anything, mock.Anything).
		Return([]model.Schedule{{}}, nil)
	m.On("PauseHandler", mock.Anything, mock.Anything, uint32(1)).
		Return(&model.Schedule{}, nil)
	m.On("ResumeHandler", mock.Anything, mock.Anything, uint32(1)).
		Return(&model.Schedule{}, nil)
	m.On("UpdateHandler", mock.Anything, mock.Anything, uint32(1), mock.Anything).
		Return(&model.Schedule{}, nil)
	return m
}

func setupMockFile() file.Handler {
	m := new(file.MockFile)
	m.On("GetFileDealsHandler", mock.Anything, mock.Anything, uint64(1)).
		Return([]model.Deal{{}}, nil)
	m.On("GetFileHandler", mock.Anything, mock.Anything, uint64(1)).
		Return(&model.File{}, nil)
	m.On("PrepareToPackFileHandler", mock.Anything, mock.Anything, uint64(1)).
		Return(int64(1), nil)
	m.On("PushFileHandler", mock.Anything, mock.Anything, "id", "name", mock.Anything).
		Return(&model.File{}, nil)
	return m
}

func setupMockJob() job.Handler {
	m := new(job.MockJob)
	m.On("StartDagGenHandler", mock.Anything, mock.Anything, "id", "name").
		Return(&model.Job{}, nil)
	m.On("PauseDagGenHandler", mock.Anything, mock.Anything, "id", "name").
		Return(&model.Job{}, nil)
	m.On("StartPackHandler", mock.Anything, mock.Anything, "id", "name", mock.Anything).
		Return([]model.Job{{}}, nil)
	m.On("PausePackHandler", mock.Anything, mock.Anything, "id", "name", mock.Anything).
		Return([]model.Job{{}}, nil)
	m.On("PackHandler", mock.Anything, mock.Anything, uint64(1)).
		Return(&model.Car{}, nil)
	m.On("PrepareToPackSourceHandler", mock.Anything, mock.Anything, "id", "name").
		Return(nil)
	m.On("StartScanHandler", mock.Anything, mock.Anything, "id", "name").
		Return(&model.Job{}, nil)
	m.On("PauseScanHandler", mock.Anything, mock.Anything, "id", "name").
		Return(&model.Job{}, nil)
	m.On("GetStatusHandler", mock.Anything, mock.Anything, "id").
		Return([]job.SourceStatus{{}}, nil)
	return m
}

func TestAllAPIs(t *testing.T) {
	mockDataPrep := setupMockDataPrep()
	mockDeal := setupMockDeal()
	mockStorage := new(storage.MockStorage)
	mockWallet := new(wallet.MockWallet)
	mockFile := setupMockFile()
	mockJob := setupMockJob()
	mockSchedule := setupMockSchedule()
	mockDealMaker := new(MockDealMaker)

	listener, err := net.Listen("tcp", apiBind)
	require.NoError(t, err)

	h, err := util.InitHost(nil)
	require.NoError(t, err)

	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		s := Server{
			db:              db,
			listener:        listener,
			lotusClient:     util.NewLotusClient("", ""),
			dealMaker:       mockDealMaker,
			closer:          io.NopCloser(nil),
			host:            h,
			storageHandler:  mockStorage,
			dataprepHandler: mockDataPrep,
			dealHandler:     mockDeal,
			walletHandler:   mockWallet,
			fileHandler:     mockFile,
			jobHandler:      mockJob,
			scheduleHandler: mockSchedule,
		}
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		go func() {
			service.StartServers(ctx, log.Logger("test"), &s)
		}()

		var resp *http2.Response
		var body string
		// try every 100ms for up to 5 seconds for server to come up
		for i := 0; i < 50; i++ {
			time.Sleep(100 * time.Millisecond)
			resp, body, _ = gorequest.New().
				Get(fmt.Sprintf("http://%s/robots.txt", apiBind)).End()
			if resp != nil && resp.StatusCode == http2.StatusOK {
				break
			}
		}
		require.NotNil(t, resp)
		require.Equal(t, http2.StatusOK, resp.StatusCode)
		require.Contains(t, body, "robotstxt.org")

		client := http.NewHTTPClientWithConfig(nil, &http.TransportConfig{
			Host:     apiBind,
			BasePath: http.DefaultBasePath,
		})

		t.Run("job", func(t *testing.T) {
			t.Run("StartDagGen", func(t *testing.T) {
				resp, err := client.Job.StartDagGen(&job2.StartDagGenParams{
					ID:      "id",
					Name:    "name",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("PauseDagGen", func(t *testing.T) {
				resp, err := client.Job.PauseDagGen(&job2.PauseDagGenParams{
					ID:      "id",
					Name:    "name",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("StartPack", func(t *testing.T) {
				resp, err := client.Job.StartPack(&job2.StartPackParams{
					ID:      "id",
					Name:    "name",
					Context: ctx,
					JobID:   1,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.Len(t, resp.Payload, 1)
			})
			t.Run("PausePack", func(t *testing.T) {
				resp, err := client.Job.PausePack(&job2.PausePackParams{
					ID:      "id",
					Name:    "name",
					Context: ctx,
					JobID:   1,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.Len(t, resp.Payload, 1)
			})
			t.Run("Pack", func(t *testing.T) {
				resp, err := client.Job.Pack(&job2.PackParams{
					ID:      1,
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("PrepareToPackSource", func(t *testing.T) {
				resp, err := client.Job.PrepareToPackSource(&job2.PrepareToPackSourceParams{
					ID:      "id",
					Name:    "name",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
			})
			t.Run("StartScan", func(t *testing.T) {
				resp, err := client.Job.StartScan(&job2.StartScanParams{
					ID:      "id",
					Name:    "name",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("PauseScan", func(t *testing.T) {
				resp, err := client.Job.PauseScan(&job2.PauseScanParams{
					ID:      "id",
					Name:    "name",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
		})

		t.Run("deal_schedule", func(t *testing.T) {
			t.Run("AddOutputStorage", func(t *testing.T) {
				resp, err := client.DealSchedule.ListPreparationSchedules(&deal_schedule.ListPreparationSchedulesParams{
					ID:      "id",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.Len(t, resp.Payload, 1)
			})
			t.Run("CreateSchedule", func(t *testing.T) {
				resp, err := client.DealSchedule.CreateSchedule(&deal_schedule.CreateScheduleParams{
					Schedule: &models.ScheduleCreateRequest{},
					Context:  ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("ListSchedules", func(t *testing.T) {
				resp, err := client.DealSchedule.ListSchedules(&deal_schedule.ListSchedulesParams{
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.Len(t, resp.Payload, 1)
			})
			t.Run("PauseHandler", func(t *testing.T) {
				resp, err := client.DealSchedule.PauseSchedule(&deal_schedule.PauseScheduleParams{
					Context: ctx,
					ID:      1,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("ResumeSchedule", func(t *testing.T) {
				resp, err := client.DealSchedule.ResumeSchedule(&deal_schedule.ResumeScheduleParams{
					Context: ctx,
					ID:      1,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("UpdateSchedule", func(t *testing.T) {
				resp, err := client.DealSchedule.UpdateSchedule(&deal_schedule.UpdateScheduleParams{
					Body:    &models.ScheduleUpdateRequest{},
					ID:      1,
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
		})

		t.Run("preparation", func(t *testing.T) {
			t.Run("CreatePreparation", func(t *testing.T) {
				resp, err := client.Preparation.CreatePreparation(&preparation.CreatePreparationParams{
					Context: ctx,
					Request: &models.DataprepCreateRequest{},
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("ExplorePreparation", func(t *testing.T) {
				resp, err := client.Preparation.ExplorePreparation(&preparation.ExplorePreparationParams{
					ID:      "id",
					Name:    "name",
					Path:    "path",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("ListPreparations", func(t *testing.T) {
				resp, err := client.Preparation.ListPreparations(&preparation.ListPreparationsParams{
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.Len(t, resp.Payload, 1)
			})
			t.Run("AddOutputStorage", func(t *testing.T) {
				resp, err := client.Preparation.AddOutputStorage(&preparation.AddOutputStorageParams{
					ID:      "id",
					Name:    "name",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("RemoveOutputStorage", func(t *testing.T) {
				resp, err := client.Preparation.RemoveOutputStorage(&preparation.RemoveOutputStorageParams{
					ID:      "id",
					Name:    "name",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("AddSourceStorage", func(t *testing.T) {
				resp, err := client.Preparation.AddSourceStorage(&preparation.AddSourceStorageParams{
					ID:      "id",
					Name:    "name",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("GetPreparationStatus", func(t *testing.T) {
				resp, err := client.Preparation.GetPreparationStatus(&preparation.GetPreparationStatusParams{
					ID:      "id",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.Len(t, resp.Payload, 1)
			})
		})

		t.Run("piece", func(t *testing.T) {
			t.Run("ListPieces", func(t *testing.T) {
				resp, err := client.Piece.ListPieces(&piece.ListPiecesParams{
					ID:      "id",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.Len(t, resp.Payload, 1)
			})
			t.Run("AddPiece", func(t *testing.T) {
				resp, err := client.Piece.AddPiece(&piece.AddPieceParams{
					ID:      "id",
					Request: &models.DataprepAddPieceRequest{},
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
		})

		t.Run("deal", func(t *testing.T) {
			t.Run("ListDeals", func(t *testing.T) {
				resp, err := client.Deal.ListDeals(&deal2.ListDealsParams{
					Context: ctx,
					Request: &models.DealListDealRequest{},
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.Len(t, resp.Payload, 1)
			})
			t.Run("SendManual", func(t *testing.T) {
				resp, err := client.Deal.SendManual(&deal2.SendManualParams{
					Proposal: &models.DealProposal{},
					Context:  ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
		})

		t.Run("file", func(t *testing.T) {
			t.Run("GetFileDeals", func(t *testing.T) {
				resp, err := client.File.GetFileDeals(&file2.GetFileDealsParams{
					ID:      1,
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.Len(t, resp.Payload, 1)
			})
			t.Run("GetFile", func(t *testing.T) {
				resp, err := client.File.GetFile(&file2.GetFileParams{
					ID:      1,
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("PrepareToPackFile", func(t *testing.T) {
				resp, err := client.File.PrepareToPackFile(&file2.PrepareToPackFileParams{
					ID:      1,
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.EqualValues(t, 1, resp.Payload)
			})
			t.Run("PushFile", func(t *testing.T) {
				resp, err := client.File.PushFile(&file2.PushFileParams{
					File:    &models.FileInfo{},
					ID:      "id",
					Name:    "name",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
		})
	})
}
