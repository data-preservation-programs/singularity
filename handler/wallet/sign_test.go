package wallet

import (
	"context"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type MockRPCClient struct {
	mock.Mock
}

func (m *MockRPCClient) Call(ctx context.Context, method string, params ...any) (*jsonrpc.RPCResponse, error) {
	panic("implement me")
}

func (m *MockRPCClient) CallRaw(ctx context.Context, request *jsonrpc.RPCRequest) (*jsonrpc.RPCResponse, error) {
	panic("implement me")
}

func (m *MockRPCClient) CallFor(ctx context.Context, out any, method string, params ...any) error {
	return m.Called(ctx, out, method, params).Error(0)
}

func (m *MockRPCClient) CallBatch(ctx context.Context, requests jsonrpc.RPCRequests) (jsonrpc.RPCResponses, error) {
	panic("implement me")
}

func (m *MockRPCClient) CallBatchRaw(ctx context.Context, requests jsonrpc.RPCRequests) (jsonrpc.RPCResponses, error) {
	panic("implement me")
}

func TestGetOrCreateActor_AlreadyLinked(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		actorID := "f01234"
		require.NoError(t, db.Create(&model.Actor{ID: actorID, Address: "f3abc"}).Error)
		w := model.Wallet{
			Address: "f3abc", KeyPath: "/tmp/key-linked", KeyStore: "local",
			ActorID: &actorID,
		}
		require.NoError(t, db.Create(&w).Error)

		// lotusClient is nil — must not be called
		actor, err := GetOrCreateActor(ctx, db, nil, &w)
		require.NoError(t, err)
		require.Equal(t, actorID, actor.ID)
		require.Equal(t, "f3abc", actor.Address)
	})
}

func TestGetOrCreateActor_LotusLookupFails(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		w := model.Wallet{
			Address: "f3unfunded", KeyPath: "/tmp/key-unfunded", KeyStore: "local",
		}
		require.NoError(t, db.Create(&w).Error)

		lotus := new(MockRPCClient)
		lotus.On("CallFor", mock.Anything, mock.AnythingOfType("*string"),
			"Filecoin.StateLookupID", []any{"f3unfunded", nil}).
			Return(errors.New("actor not found"))

		_, err := GetOrCreateActor(ctx, db, lotus, &w)
		require.Error(t, err)
		require.ErrorContains(t, err, "not found on-chain")
		require.ErrorContains(t, err, "may need funding")
		lotus.AssertExpectations(t)
	})
}

func TestGetOrCreateActor_CreateNewActor(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		w := model.Wallet{
			Address: "f3new", KeyPath: "/tmp/key-new", KeyStore: "local",
		}
		require.NoError(t, db.Create(&w).Error)

		lotus := new(MockRPCClient)
		lotus.On("CallFor", mock.Anything, mock.AnythingOfType("*string"),
			"Filecoin.StateLookupID", []any{"f3new", nil}).
			Return(nil).Run(func(args mock.Arguments) {
			*args.Get(1).(*string) = "f09999"
		})

		actor, err := GetOrCreateActor(ctx, db, lotus, &w)
		require.NoError(t, err)
		require.Equal(t, "f09999", actor.ID)
		require.Equal(t, "f3new", actor.Address)

		// wallet should now be linked
		var updated model.Wallet
		require.NoError(t, db.First(&updated, w.ID).Error)
		require.NotNil(t, updated.ActorID)
		require.Equal(t, "f09999", *updated.ActorID)

		// actor should exist in DB
		var dbActor model.Actor
		require.NoError(t, db.First(&dbActor, "id = ?", "f09999").Error)
		require.Equal(t, "f3new", dbActor.Address)

		lotus.AssertExpectations(t)
	})
}

func TestGetOrCreateActor_LinkExistingActor(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// actor exists in DB but wallet is not yet linked to it
		require.NoError(t, db.Create(&model.Actor{ID: "f05555", Address: "f3existing"}).Error)

		w := model.Wallet{
			Address: "f3existing", KeyPath: "/tmp/key-existing", KeyStore: "local",
		}
		require.NoError(t, db.Create(&w).Error)

		lotus := new(MockRPCClient)
		lotus.On("CallFor", mock.Anything, mock.AnythingOfType("*string"),
			"Filecoin.StateLookupID", []any{"f3existing", nil}).
			Return(nil).Run(func(args mock.Arguments) {
			*args.Get(1).(*string) = "f05555"
		})

		actor, err := GetOrCreateActor(ctx, db, lotus, &w)
		require.NoError(t, err)
		require.Equal(t, "f05555", actor.ID)

		// wallet should now be linked
		var updated model.Wallet
		require.NoError(t, db.First(&updated, w.ID).Error)
		require.NotNil(t, updated.ActorID)
		require.Equal(t, "f05555", *updated.ActorID)

		lotus.AssertExpectations(t)
	})
}

func TestGetOrCreateActor_ActorLinkedToDifferentWallet(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		actorID := "f07777"
		require.NoError(t, db.Create(&model.Actor{ID: actorID, Address: "f3other"}).Error)

		// first wallet already linked to this actor
		other := model.Wallet{
			Address: "f3other", KeyPath: "/tmp/key-other", KeyStore: "local",
			ActorID: &actorID,
		}
		require.NoError(t, db.Create(&other).Error)

		// second wallet tries to claim the same actor
		w := model.Wallet{
			Address: "f3conflict", KeyPath: "/tmp/key-conflict", KeyStore: "local",
		}
		require.NoError(t, db.Create(&w).Error)

		lotus := new(MockRPCClient)
		lotus.On("CallFor", mock.Anything, mock.AnythingOfType("*string"),
			"Filecoin.StateLookupID", []any{"f3conflict", nil}).
			Return(nil).Run(func(args mock.Arguments) {
			*args.Get(1).(*string) = actorID
		})

		_, err := GetOrCreateActor(ctx, db, lotus, &w)
		require.Error(t, err)
		require.ErrorContains(t, err, "already linked to wallet")

		// wallet must remain unlinked
		var updated model.Wallet
		require.NoError(t, db.First(&updated, w.ID).Error)
		require.Nil(t, updated.ActorID)

		lotus.AssertExpectations(t)
	})
}
