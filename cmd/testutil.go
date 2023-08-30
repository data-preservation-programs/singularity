package cmd

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/cockroachdb/errors/oserror"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/mattn/go-shellwords"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

func CompareWith(t *testing.T, actual string, path string) {
	path = filepath.Join("testdata", path)
	_, err := os.Stat(path)
	if errors.Is(err, oserror.ErrNotExist) {
		err = os.WriteFile(path, []byte(actual), 0644)
		require.NoError(t, err)
	}
	require.NoError(t, err)

	expected, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Equal(t, string(expected), actual)
}

func Run(ctx context.Context, args string) (string, string, error) {
	// Create a clone of the app so that we can run from different tests concurrently
	app := *App
	app.ExitErrHandler = func(c *cli.Context, err error) {}
	parser := shellwords.NewParser()
	parser.ParseEnv = true // Enable environment variable parsing
	parsedArgs, err := parser.Parse(args)
	if err != nil {
		return "", "", errors.WithStack(err)
	}

	outWriter := bytes.NewBuffer(nil)
	errWriter := bytes.NewBuffer(nil)

	// Overwrite the stdout and stderr
	app.Writer = outWriter
	app.ErrWriter = errWriter

	err = app.RunContext(ctx, parsedArgs)
	return outWriter.String(), errWriter.String(), err
}

type MockAdmin struct {
	mock.Mock
}

func (m *MockAdmin) InitHandler(ctx context.Context, db *gorm.DB) error {
	args := m.Called(ctx, db)
	return args.Error(0)
}

func (m *MockAdmin) ResetHandler(ctx context.Context, db *gorm.DB) error {
	args := m.Called(ctx, db)
	return args.Error(0)
}

type MockDataPrep struct {
	mock.Mock
}

func (m *MockDataPrep) CreatePreparationHandler(ctx context.Context, db *gorm.DB, request dataprep.CreateRequest) (*model.Preparation, error) {
	args := m.Called(ctx, db, request)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockDataPrep) StartDagGenHandler(ctx context.Context, db *gorm.DB, id uint32, name string) (*model.Job, error) {
	args := m.Called(ctx, db, id, name)
	return args.Get(0).(*model.Job), args.Error(1)
}

func (m *MockDataPrep) PauseDagGenHandler(ctx context.Context, db *gorm.DB, id uint32, name string) (*model.Job, error) {
	args := m.Called(ctx, db, id, name)
	return args.Get(0).(*model.Job), args.Error(1)
}

func (m *MockDataPrep) ExploreHandler(ctx context.Context, db *gorm.DB, id uint32, name string, path string) ([]dataprep.DirEntry, error) {
	args := m.Called(ctx, db, id, name, path)
	return args.Get(0).([]dataprep.DirEntry), args.Error(1)
}

func (m *MockDataPrep) ListHandler(ctx context.Context, db *gorm.DB) ([]model.Preparation, error) {
	args := m.Called(ctx, db)
	return args.Get(0).([]model.Preparation), args.Error(1)
}

func (m *MockDataPrep) AddOutputStorageHandler(ctx context.Context, db *gorm.DB, id uint32, output string) (*model.Preparation, error) {
	args := m.Called(ctx, db, id, output)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockDataPrep) RemoveOutputStorageHandler(ctx context.Context, db *gorm.DB, id uint32, output string) (*model.Preparation, error) {
	args := m.Called(ctx, db, id, output)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockDataPrep) StartPackHandler(ctx context.Context, db *gorm.DB, id uint32, name string, jobID int64) ([]model.Job, error) {
	args := m.Called(ctx, db, id, name, jobID)
	return args.Get(0).([]model.Job), args.Error(1)
}

func (m *MockDataPrep) PausePackHandler(ctx context.Context, db *gorm.DB, id uint32, name string, jobID int64) ([]model.Job, error) {
	args := m.Called(ctx, db, id, name, jobID)
	return args.Get(0).([]model.Job), args.Error(1)
}

func (m *MockDataPrep) ListPiecesHandler(ctx context.Context, db *gorm.DB, id uint32) ([]dataprep.PieceList, error) {
	args := m.Called(ctx, db, id)
	return args.Get(0).([]dataprep.PieceList), args.Error(1)
}

func (m *MockDataPrep) AddPieceHandler(ctx context.Context, db *gorm.DB, id uint32, request dataprep.AddPieceRequest) (*model.Car, error) {
	args := m.Called(ctx, db, id, request)
	return args.Get(0).(*model.Car), args.Error(1)
}

func (m *MockDataPrep) StartScanHandler(ctx context.Context, db *gorm.DB, id uint32, name string) (*model.Job, error) {
	args := m.Called(ctx, db, id, name)
	return args.Get(0).(*model.Job), args.Error(1)
}

func (m *MockDataPrep) PauseScanHandler(ctx context.Context, db *gorm.DB, id uint32, name string) (*model.Job, error) {
	args := m.Called(ctx, db, id, name)
	return args.Get(0).(*model.Job), args.Error(1)
}

func (m *MockDataPrep) AddSourceStorageHandler(ctx context.Context, db *gorm.DB, id uint32, source string) (*model.Preparation, error) {
	args := m.Called(ctx, db, id, source)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockDataPrep) GetStatusHandler(ctx context.Context, db *gorm.DB, id uint32) (*dataprep.Status, error) {
	args := m.Called(ctx, db, id)
	return args.Get(0).(*dataprep.Status), args.Error(1)
}

type MockSchedule struct {
	mock.Mock
}

func (m *MockSchedule) CreateHandler(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, request schedule.CreateRequest) (*model.Schedule, error) {
	args := m.Called(ctx, db, lotusClient, request)
	return args.Get(0).(*model.Schedule), args.Error(1)
}

func (m *MockSchedule) ListHandler(ctx context.Context, db *gorm.DB) ([]model.Schedule, error) {
	args := m.Called(ctx, db)
	return args.Get(0).([]model.Schedule), args.Error(1)
}

func (m *MockSchedule) PauseHandler(ctx context.Context, db *gorm.DB, scheduleID uint32) (*model.Schedule, error) {
	args := m.Called(ctx, db, scheduleID)
	return args.Get(0).(*model.Schedule), args.Error(1)
}

func (m *MockSchedule) ResumeHandler(ctx context.Context, db *gorm.DB, scheduleID uint32) (*model.Schedule, error) {
	args := m.Called(ctx, db, scheduleID)
	return args.Get(0).(*model.Schedule), args.Error(1)
}

type MockDeal struct {
	mock.Mock
}

func (m *MockDeal) ListHandler(ctx context.Context, db *gorm.DB, request deal.ListDealRequest) ([]model.Deal, error) {
	args := m.Called(ctx, db, request)
	return args.Get(0).([]model.Deal), args.Error(1)
}

func (m *MockDeal) SendManualHandler(ctx context.Context, db *gorm.DB, dealMaker replication.DealMaker, request deal.Proposal) (*model.Deal, error) {
	args := m.Called(ctx, db, dealMaker, request)
	return args.Get(0).(*model.Deal), args.Error(1)
}

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) CreateStorageHandler(ctx context.Context, db *gorm.DB, storageType string, request storage.CreateRequest) (*model.Storage, error) {
	args := m.Called(ctx, db, storageType, request)
	return args.Get(0).(*model.Storage), args.Error(1)
}

func (m *MockStorage) ExploreHandler(ctx context.Context, db *gorm.DB, name string, path string) ([]storage.DirEntry, error) {
	args := m.Called(ctx, db, name, path)
	return args.Get(0).([]storage.DirEntry), args.Error(1)
}

func (m *MockStorage) ListStoragesHandler(ctx context.Context, db *gorm.DB) ([]model.Storage, error) {
	args := m.Called(ctx, db)
	return args.Get(0).([]model.Storage), args.Error(1)
}

func (m *MockStorage) RemoveHandler(ctx context.Context, db *gorm.DB, name string) error {
	args := m.Called(ctx, db, name)
	return args.Error(0)
}

func (m *MockStorage) UpdateStorageHandler(ctx context.Context, db *gorm.DB, name string, config map[string]string) (*model.Storage, error) {
	args := m.Called(ctx, db, name, config)
	return args.Get(0).(*model.Storage), args.Error(1)
}

type MockWallet struct {
	mock.Mock
}

func (m *MockWallet) AttachHandler(ctx context.Context, db *gorm.DB, preparationID uint32, wallet string) (*model.Preparation, error) {
	args := m.Called(ctx, db, preparationID, wallet)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockWallet) DetachHandler(ctx context.Context, db *gorm.DB, preparationID uint32, wallet string) (*model.Preparation, error) {
	args := m.Called(ctx, db, preparationID, wallet)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockWallet) ImportHandler(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, request wallet.ImportRequest) (*model.Wallet, error) {
	args := m.Called(ctx, db, lotusClient, request)
	return args.Get(0).(*model.Wallet), args.Error(1)
}

func (m *MockWallet) ListHandler(ctx context.Context, db *gorm.DB) ([]model.Wallet, error) {
	args := m.Called(ctx, db)
	return args.Get(0).([]model.Wallet), args.Error(1)
}

func (m *MockWallet) ListAttachedHandler(ctx context.Context, db *gorm.DB, preparationID uint32) ([]model.Wallet, error) {
	args := m.Called(ctx, db, preparationID)
	return args.Get(0).([]model.Wallet), args.Error(1)
}

func (m *MockWallet) RemoveHandler(ctx context.Context, db *gorm.DB, address string) error {
	args := m.Called(ctx, db, address)
	return args.Error(0)
}
