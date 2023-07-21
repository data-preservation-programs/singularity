package schedule

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/ybbus/jsonrpc/v3"
)

type MockRPCClient struct {
	mock.Mock
}

func (m *MockRPCClient) Call(ctx context.Context, method string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRPCClient) CallRaw(ctx context.Context, request *jsonrpc.RPCRequest) (*jsonrpc.RPCResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRPCClient) CallFor(ctx context.Context, out interface{}, method string, params ...interface{}) error {
	return m.Called(ctx, out, method, params).Error(0)
}

func (m *MockRPCClient) CallBatch(ctx context.Context, requests jsonrpc.RPCRequests) (jsonrpc.RPCResponses, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRPCClient) CallBatchRaw(ctx context.Context, requests jsonrpc.RPCRequests) (jsonrpc.RPCResponses, error) {
	//TODO implement me
	panic("implement me")
}

func getMockLotusClient() jsonrpc.RPCClient {
	lotusClient := new(MockRPCClient)
	// Set up expectations for the lotusClient mock
	lotusClient.On("CallFor", mock.Anything, mock.Anything, "Filecoin.StateLookupID", mock.Anything).
		Return(nil)
	return lotusClient
}

var createRequest = CreateRequest{
	DatasetName:          "test",
	Provider:             "f01000",
	HTTPHeaders:          []string{"a=b"},
	URLTemplate:          "http://127.0.0.1",
	PricePerGBEpoch:      0,
	PricePerGB:           0,
	PricePerDeal:         0,
	Verified:             true,
	IPNI:                 true,
	KeepUnsealed:         true,
	StartDelay:           "24h",
	Duration:             "2400h",
	ScheduleCron:         "",
	ScheduleDealNumber:   100,
	TotalDealNumber:      100,
	ScheduleDealSize:     "1TiB",
	TotalDealSize:        "1PiB",
	Notes:                "notes",
	MaxPendingDealSize:   "10TiB",
	MaxPendingDealNumber: 100,
	AllowedPieceCIDs:     []string{"baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq"},
}

func TestCreateHandler_DatasetNotFound(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
	_, err = CreateHandler(db, context.Background(), getMockLotusClient(), createRequest)
	require.ErrorContains(t, err, "dataset not found")
}

func TestCreateHandler_InvalidStartDelay(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
	require.NoError(t, db.Create(&model.Dataset{Name: "test"}).Error)
	badRequest := createRequest
	badRequest.StartDelay = "1year"
	_, err = CreateHandler(db, context.Background(), getMockLotusClient(), badRequest)
	require.ErrorContains(t, err, "invalid start delay")
}

func TestCreateHandler_InvalidDuration(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
	require.NoError(t, db.Create(&model.Dataset{Name: "test"}).Error)
	badRequest := createRequest
	badRequest.Duration = "1year"
	_, err = CreateHandler(db, context.Background(), getMockLotusClient(), badRequest)
	require.ErrorContains(t, err, "invalid duration")
}

func TestCreateHandler_InvalidScheduleInterval(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
	require.NoError(t, db.Create(&model.Dataset{Name: "test"}).Error)
	badRequest := createRequest
	badRequest.ScheduleCron = "1year"
	_, err = CreateHandler(db, context.Background(), getMockLotusClient(), badRequest)
	require.ErrorContains(t, err, "invalid schedule cron")
}

func TestCreateHandler_InvalidScheduleDealSize(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&model.Dataset{Name: "test"}).Error
	require.NoError(t, err)
	badRequest := createRequest
	badRequest.ScheduleDealSize = "One PB"
	_, err = CreateHandler(db, context.Background(), getMockLotusClient(), badRequest)
	require.ErrorContains(t, err, "invalid schedule deal size")
}

func TestCreateHandler_InvalidTotalDealSize(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&model.Dataset{Name: "test"}).Error
	require.NoError(t, err)
	badRequest := createRequest
	badRequest.TotalDealSize = "One PB"
	_, err = CreateHandler(db, context.Background(), getMockLotusClient(), badRequest)
	require.ErrorContains(t, err, "invalid total deal size")
}

func TestCreateHandler_InvalidPendingDealSize(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&model.Dataset{Name: "test"}).Error
	require.NoError(t, err)
	badRequest := createRequest
	badRequest.MaxPendingDealSize = "One PB"
	_, err = CreateHandler(db, context.Background(), getMockLotusClient(), badRequest)
	require.ErrorContains(t, err, "invalid pending deal size")
}

func TestCreateHandler_InvalidAllowedPieceCID_NotCID(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&model.Dataset{Name: "test"}).Error
	require.NoError(t, err)
	badRequest := createRequest
	badRequest.AllowedPieceCIDs = []string{"not a cid"}
	_, err = CreateHandler(db, context.Background(), getMockLotusClient(), badRequest)
	require.ErrorContains(t, err, "it's not a CID")
}

func TestCreateHandler_InvalidAllowedPieceCID_NotCommp(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&model.Dataset{Name: "test"}).Error
	require.NoError(t, err)
	badRequest := createRequest
	badRequest.AllowedPieceCIDs = []string{"bafybeiejlvvmfokp5c6q2eqgbfjeaokz3nqho5c7yy3ov527vsatgsqfma"}
	_, err = CreateHandler(db, context.Background(), getMockLotusClient(), badRequest)
	require.ErrorContains(t, err, "it's not a commp")
}

func TestCreateHandler_NoAssociatedWallet(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&model.Dataset{Name: "test"}).Error
	require.NoError(t, err)
	_, err = CreateHandler(db, context.Background(), getMockLotusClient(), createRequest)
	require.ErrorContains(t, err, "no wallet")
}

func TestCreateHandler_InvalidProvider(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&model.Dataset{Name: "test"}).Error
	require.NoError(t, err)
	err = db.Create(&model.Wallet{ID: "f01"}).Error
	require.NoError(t, err)
	err = db.Create(&model.WalletAssignment{WalletID: "f01", DatasetID: 1}).Error
	require.NoError(t, err)
	lotusClient := new(MockRPCClient)
	lotusClient.On("CallFor", mock.Anything, mock.Anything, "Filecoin.StateLookupID", mock.Anything).
		Return(errors.New("Some provider error"))
	_, err = CreateHandler(db, context.Background(), lotusClient, createRequest)
	require.ErrorContains(t, err, "Some provider error")
}

func TestCreateHandler_Success(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&model.Dataset{Name: "test"}).Error
	require.NoError(t, err)
	err = db.Create(&model.Wallet{ID: "f01"}).Error
	require.NoError(t, err)
	err = db.Create(&model.WalletAssignment{WalletID: "f01", DatasetID: 1}).Error
	require.NoError(t, err)
	schedule, err := CreateHandler(db, context.Background(), getMockLotusClient(), createRequest)
	require.NoError(t, err)
	require.NotNil(t, schedule)
}
