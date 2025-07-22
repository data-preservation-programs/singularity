package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	http2 "net/http"
	"strings"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/client/swagger/http"
	admin2 "github.com/data-preservation-programs/singularity/client/swagger/http/admin"
	deal2 "github.com/data-preservation-programs/singularity/client/swagger/http/deal"
	"github.com/data-preservation-programs/singularity/client/swagger/http/deal_schedule"
	file2 "github.com/data-preservation-programs/singularity/client/swagger/http/file"
	job2 "github.com/data-preservation-programs/singularity/client/swagger/http/job"
	"github.com/data-preservation-programs/singularity/client/swagger/http/piece"
	"github.com/data-preservation-programs/singularity/client/swagger/http/preparation"
	storage2 "github.com/data-preservation-programs/singularity/client/swagger/http/storage"
	wallet2 "github.com/data-preservation-programs/singularity/client/swagger/http/wallet"
	"github.com/data-preservation-programs/singularity/client/swagger/http/wallet_association"
	"github.com/data-preservation-programs/singularity/client/swagger/models"
	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/singularity/handler/dealtemplate"
	"github.com/data-preservation-programs/singularity/handler/errorlog"
	"github.com/data-preservation-programs/singularity/handler/file"
	"github.com/data-preservation-programs/singularity/handler/job"
	"github.com/data-preservation-programs/singularity/handler/statechange"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
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

func setupMockAdmin() admin.Handler {
	m := new(admin.MockAdmin)
	m.On("InitHandler", mock.Anything, mock.Anything).
		Return(nil)
	m.On("ResetHandler", mock.Anything, mock.Anything).
		Return(nil)
	m.On("SetIdentityHandler", mock.Anything, mock.Anything, mock.Anything).
		Return(nil)
	return m
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
	m.On("RenamePreparationHandler", mock.Anything, mock.Anything, "old", mock.Anything).
		Return(&model.Preparation{}, nil)
	m.On("RemovePreparationHandler", mock.Anything, mock.Anything, "old", mock.Anything).
		Return(nil)
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
	m.On("RemoveHandler", mock.Anything, mock.Anything, uint32(1)).
		Return(nil)
	return m
}

type nopCloser struct {
	io.ReadSeeker
}

func (nopCloser) Close() error { return nil }

func setupMockFile() file.Handler {
	m := new(file.MockFile)
	m.On("GetFileDealsHandler", mock.Anything, mock.Anything, uint64(1)).
		Return([]file.DealsForFileRange{{}}, nil)
	m.On("GetFileHandler", mock.Anything, mock.Anything, uint64(1)).
		Return(&model.File{}, nil)
	m.On("PrepareToPackFileHandler", mock.Anything, mock.Anything, uint64(1)).
		Return(int64(1), nil)
	m.On("PushFileHandler", mock.Anything, mock.Anything, "id", "name", mock.Anything).
		Return(&model.File{}, nil)
	m.On("RetrieveFileHandler", mock.Anything, mock.Anything, mock.Anything, uint64(1)).
		Return(io.ReadSeekCloser(nopCloser{strings.NewReader("hello world")}), "hello.txt", time.Date(1999, 12, 31, 11, 59, 59, 0, time.UTC), nil)
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

func setupMockStorage() storage.Handler {
	m := new(storage.MockStorage)
	m.On("CreateStorageHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(&model.Storage{}, nil)
	m.On("ExploreHandler", mock.Anything, mock.Anything, "name", "path").
		Return([]storage.DirEntry{{}}, nil)
	m.On("ListStoragesHandler", mock.Anything, mock.Anything).
		Return([]model.Storage{{}}, nil)
	m.On("RemoveHandler", mock.Anything, mock.Anything, "name").
		Return(nil)
	m.On("UpdateStorageHandler", mock.Anything, mock.Anything, "name", mock.Anything).
		Return(&model.Storage{}, nil)
	m.On("RenameStorageHandler", mock.Anything, mock.Anything, "old", mock.Anything).
		Return(&model.Storage{}, nil)
	return m
}

func setupMockWallet() wallet.Handler {
	m := new(wallet.MockWallet)
	m.On("AttachHandler", mock.Anything, mock.Anything, "id", "wallet").
		Return(&model.Preparation{}, nil)
	m.On("CreateHandler", mock.Anything, mock.Anything, mock.Anything).
		Return(&model.Wallet{}, nil)
	m.On("DetachHandler", mock.Anything, mock.Anything, "id", "wallet").
		Return(&model.Preparation{}, nil)
	m.On("ImportHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(&model.Wallet{}, nil)
	m.On("InitHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(&model.Wallet{}, nil)
	m.On("ListHandler", mock.Anything, mock.Anything).
		Return([]model.Wallet{{}}, nil)
	m.On("ListAttachedHandler", mock.Anything, mock.Anything, "id").
		Return([]model.Wallet{{}}, nil)
	m.On("RemoveHandler", mock.Anything, mock.Anything, "wallet").
		Return(nil)
	return m
}

func TestAllAPIs(t *testing.T) {
	mockAdmin := setupMockAdmin()
	mockDataPrep := setupMockDataPrep()
	mockDeal := setupMockDeal()
	mockStorage := setupMockStorage()
	mockWallet := setupMockWallet()
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
			db:                  db,
			listener:            listener,
			lotusClient:         util.NewLotusClient("", ""),
			dealMaker:           mockDealMaker,
			closer:              io.NopCloser(nil),
			host:                h,
			adminHandler:        mockAdmin,
			storageHandler:      mockStorage,
			dataprepHandler:     mockDataPrep,
			dealHandler:         mockDeal,
			walletHandler:       mockWallet,
			fileHandler:         mockFile,
			jobHandler:          mockJob,
			scheduleHandler:     mockSchedule,
			stateChangeHandler:  &statechange.DefaultHandler{},
			dealtemplateHandler: &dealtemplate.DefaultHandler{},
			errorlogHandler:     &errorlog.DefaultHandler{},
		}
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		go func() {
			_ = service.StartServers(ctx, log.Logger("test"), &s)
		}()

		var resp *http2.Response
		// try every 100ms for up to 5 seconds for server to come up
		for i := 0; i < 50; i++ {
			time.Sleep(100 * time.Millisecond)
			resp, _, _ = gorequest.New().
				Get(fmt.Sprintf("http://%s/health", apiBind)).End()
			if resp != nil && resp.StatusCode == http2.StatusOK {
				break
			}
		}
		require.NotNil(t, resp)
		require.Equal(t, http2.StatusOK, resp.StatusCode)

		client := http.NewHTTPClientWithConfig(nil, &http.TransportConfig{
			Host:     apiBind,
			BasePath: http.DefaultBasePath,
		})

		t.Run("admin", func(t *testing.T) {
			t.Run("SetIdentity", func(t *testing.T) {
				resp, err := client.Admin.SetIdentity(&admin2.SetIdentityParams{
					Context: ctx,
					Request: &models.AdminSetIdentityRequest{
						Identity: "test",
					},
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
			})
		})

		t.Run("wallet_association", func(t *testing.T) {
			t.Run("AttachWallet", func(t *testing.T) {
				resp, err := client.WalletAssociation.AttachWallet(&wallet_association.AttachWalletParams{
					ID:      "id",
					Wallet:  "wallet",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("DetachWallet", func(t *testing.T) {
				resp, err := client.WalletAssociation.DetachWallet(&wallet_association.DetachWalletParams{
					ID:      "id",
					Wallet:  "wallet",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("ListAttachedHandler", func(t *testing.T) {
				resp, err := client.WalletAssociation.ListAttachedWallets(&wallet_association.ListAttachedWalletsParams{
					ID:      "id",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.Len(t, resp.Payload, 1)
			})
		})

		t.Run("wallet", func(t *testing.T) {
			t.Run("CreateWallet", func(t *testing.T) {
				resp, err := client.Wallet.CreateWallet(&wallet2.CreateWalletParams{
					Request: &models.WalletCreateRequest{},
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("ImportWallet", func(t *testing.T) {
				resp, err := client.Wallet.ImportWallet(&wallet2.ImportWalletParams{
					Request: &models.WalletImportRequest{},
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("InitWallet", func(t *testing.T) {
				resp, err := client.Wallet.InitWallet(&wallet2.InitWalletParams{
					Context: ctx,
					Address: "wallet",
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("ListWallets", func(t *testing.T) {
				resp, err := client.Wallet.ListWallets(&wallet2.ListWalletsParams{
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.Len(t, resp.Payload, 1)
			})
			t.Run("RemoveWallet", func(t *testing.T) {
				resp, err := client.Wallet.RemoveWallet(&wallet2.RemoveWalletParams{
					Context: ctx,
					Address: "wallet",
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
			})
		})

		t.Run("storage", func(t *testing.T) {
			t.Run("RenameStorage", func(t *testing.T) {
				resp, err := client.Storage.RenameStorage(&storage2.RenameStorageParams{
					Name: "old",
					Request: &models.StorageRenameRequest{
						Name: ptr.Of("new"),
					},
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("CreateS3AWSStorage", func(t *testing.T) {
				resp, err := client.Storage.CreateS3AWSStorage(&storage2.CreateS3AWSStorageParams{
					Request: &models.StorageCreateS3AWSStorageRequest{},
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
			t.Run("ExploreStorage", func(t *testing.T) {
				resp, err := client.Storage.ExploreStorage(&storage2.ExploreStorageParams{
					Name:    "name",
					Path:    "path",
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.Len(t, resp.Payload, 1)
			})
			t.Run("ListStorages", func(t *testing.T) {
				resp, err := client.Storage.ListStorages(&storage2.ListStoragesParams{
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.Len(t, resp.Payload, 1)
			})
			t.Run("RemoveStorage", func(t *testing.T) {
				resp, err := client.Storage.RemoveStorage(&storage2.RemoveStorageParams{
					Context: ctx,
					Name:    "name",
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
			})
			t.Run("UpdateStorage", func(t *testing.T) {
				resp, err := client.Storage.UpdateStorage(&storage2.UpdateStorageParams{
					Context: ctx,
					Name:    "name",
					Config:  map[string]string{},
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
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
			t.Run("RemoveSchedule", func(t *testing.T) {
				resp, err := client.DealSchedule.RemoveSchedule(&deal_schedule.RemoveScheduleParams{
					ID:      1,
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
			})
		})

		t.Run("preparation", func(t *testing.T) {
			t.Run("RenamePreparation", func(t *testing.T) {
				resp, err := client.Preparation.RenamePreparation(&preparation.RenamePreparationParams{
					Name: "old",
					Request: &models.DataprepRenameRequest{
						Name: ptr.Of("new"),
					},
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
				require.NotNil(t, resp.Payload)
			})
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
			t.Run("RemovePreparation", func(t *testing.T) {
				resp, err := client.Preparation.RemovePreparation(&preparation.RemovePreparationParams{
					Name: "old",
					Request: &models.DataprepRemoveRequest{
						RemoveCars: true,
					},
					Context: ctx,
				})
				require.NoError(t, err)
				require.True(t, resp.IsSuccess())
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
			t.Run("RetrieveFile", func(t *testing.T) {
				buf := new(bytes.Buffer)
				resp, partial, err := client.File.RetrieveFile(&file2.RetrieveFileParams{
					ID:      1,
					Context: ctx,
				}, buf)
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.True(t, resp.IsSuccess())
				require.Nil(t, partial)
				require.Equal(t, "hello world", buf.String())
			})
		})
	})
}
