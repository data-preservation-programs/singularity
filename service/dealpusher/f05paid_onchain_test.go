package dealpusher

import (
	"context"
	"math/big"
	"testing"

	"github.com/data-preservation-programs/go-synapse/spregistry"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

type balanceClientMock struct {
	balance *big.Int
	err     error
}

func (m *balanceClientMock) BalanceAt(context.Context, common.Address, *big.Int) (*big.Int, error) {
	if m.err != nil {
		return nil, m.err
	}
	return new(big.Int).Set(m.balance), nil
}

type providerRegistryMock struct {
	provider *spregistry.ProviderInfo
	err      error
}

func (m *providerRegistryMock) GetProvider(context.Context, int) (*spregistry.ProviderInfo, error) {
	return m.provider, m.err
}

func newF05PaidTestSchedule(t *testing.T) (*keystore.LocalKeyStore, *model.Schedule) {
	t.Helper()

	ks, err := keystore.NewLocalKeyStore(t.TempDir())
	require.NoError(t, err)

	keyPath, addr, err := ks.Put(testutil.TestPrivateKeyHex)
	require.NoError(t, err)

	wallet := &model.Wallet{
		ID:       1,
		Address:  addr.String(),
		KeyPath:  keyPath,
		KeyStore: "local",
	}

	return ks, &model.Schedule{
		ID:       1,
		Provider: "f01000",
		Preparation: &model.Preparation{
			Wallet: wallet,
		},
	}
}

func TestOnChainF05PaidRunScheduleRequiresWallet(t *testing.T) {
	manager := &OnChainF05Paid{cfg: defaultF05PaidSchedulingConfig()}
	state, err := manager.RunSchedule(context.Background(), &model.Schedule{})
	require.ErrorContains(t, err, "no wallet configured")
	require.Equal(t, model.ScheduleError, state)
}

func TestOnChainF05PaidRunScheduleRejectsMissingProvider(t *testing.T) {
	ks, schedule := newF05PaidTestSchedule(t)
	manager := &OnChainF05Paid{
		keyStore:         ks,
		cfg:              defaultF05PaidSchedulingConfig(),
		balanceClient:    &balanceClientMock{balance: big.NewInt(1)},
		providerRegistry: &providerRegistryMock{},
	}

	state, err := manager.RunSchedule(context.Background(), schedule)
	require.ErrorContains(t, err, "not registered in SP Registry")
	require.Equal(t, model.ScheduleError, state)
}

func TestOnChainF05PaidRunScheduleRejectsInactiveProvider(t *testing.T) {
	ks, schedule := newF05PaidTestSchedule(t)
	manager := &OnChainF05Paid{
		keyStore:      ks,
		cfg:           defaultF05PaidSchedulingConfig(),
		balanceClient: &balanceClientMock{balance: big.NewInt(1)},
		providerRegistry: &providerRegistryMock{provider: &spregistry.ProviderInfo{
			Active: false,
		}},
	}

	state, err := manager.RunSchedule(context.Background(), schedule)
	require.ErrorContains(t, err, "not active in SP Registry")
	require.Equal(t, model.ScheduleError, state)
}

func TestOnChainF05PaidRunScheduleRejectsLowBalance(t *testing.T) {
	ks, schedule := newF05PaidTestSchedule(t)
	manager := &OnChainF05Paid{
		keyStore: ks,
		cfg: F05PaidSchedulingConfig{
			MinWalletBalanceAttoFIL: big.NewInt(2),
		},
		balanceClient: &balanceClientMock{balance: big.NewInt(1)},
		providerRegistry: &providerRegistryMock{provider: &spregistry.ProviderInfo{
			Active: true,
			Payee:  common.HexToAddress("0x00000000000000000000000000000000000000aa"),
		}},
	}

	state, err := manager.RunSchedule(context.Background(), schedule)
	require.ErrorContains(t, err, "below the configured minimum")
	require.Equal(t, model.ScheduleError, state)
}

func TestOnChainF05PaidRunScheduleReturnsNotImplementedAfterPreflight(t *testing.T) {
	ks, schedule := newF05PaidTestSchedule(t)
	manager := &OnChainF05Paid{
		keyStore:      ks,
		cfg:           defaultF05PaidSchedulingConfig(),
		balanceClient: &balanceClientMock{balance: big.NewInt(1)},
		paymentsAddr:  common.HexToAddress("0x00000000000000000000000000000000000000bb"),
		providerRegistry: &providerRegistryMock{provider: &spregistry.ProviderInfo{
			Active: true,
			Payee:  common.HexToAddress("0x00000000000000000000000000000000000000aa"),
			Products: map[string]*spregistry.ServiceProduct{
				"PDP": {
					Data: &spregistry.PDPOffering{ServiceURL: "https://provider.example"},
				},
			},
		}},
	}

	state, err := manager.RunSchedule(context.Background(), schedule)
	require.ErrorContains(t, err, "execution is not implemented yet")
	require.Equal(t, model.ScheduleError, state)
}
