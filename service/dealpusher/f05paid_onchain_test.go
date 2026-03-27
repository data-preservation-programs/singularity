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

type paymentsContractMock struct {
	availableFunds *big.Int
	approved       bool
	accountErr     error
	approvalErr    error
}

func (m *paymentsContractMock) GetAccountInfoIfSettled(context.Context, common.Address, common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	if m.accountErr != nil {
		return nil, nil, nil, nil, m.accountErr
	}
	availableFunds := big.NewInt(0)
	if m.availableFunds != nil {
		availableFunds = new(big.Int).Set(m.availableFunds)
	}
	return big.NewInt(0), big.NewInt(0), availableFunds, big.NewInt(0), nil
}

func (m *paymentsContractMock) GetOperatorApproval(context.Context, common.Address, common.Address, common.Address) (bool, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	if m.approvalErr != nil {
		return false, nil, nil, nil, nil, nil, m.approvalErr
	}
	return m.approved, big.NewInt(10), big.NewInt(20), big.NewInt(1), big.NewInt(2), big.NewInt(30), nil
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

func newPaidProviderInfo() *spregistry.ProviderInfo {
	return &spregistry.ProviderInfo{
		Active:          true,
		ServiceProvider: common.HexToAddress("0x00000000000000000000000000000000000000cc"),
		Payee:           common.HexToAddress("0x00000000000000000000000000000000000000aa"),
		Products: map[string]*spregistry.ServiceProduct{
			"PDP": {
				IsActive: true,
				Data: &spregistry.PDPOffering{
					ServiceURL:          "https://provider.example",
					PaymentTokenAddress: common.HexToAddress("0x00000000000000000000000000000000000000dd"),
				},
			},
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
		paymentsContract: &paymentsContractMock{availableFunds: big.NewInt(1), approved: true},
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
		paymentsContract: &paymentsContractMock{availableFunds: big.NewInt(1), approved: true},
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
		balanceClient:    &balanceClientMock{balance: big.NewInt(1)},
		providerRegistry: &providerRegistryMock{provider: newPaidProviderInfo()},
		paymentsContract: &paymentsContractMock{availableFunds: big.NewInt(1), approved: true},
	}

	state, err := manager.RunSchedule(context.Background(), schedule)
	require.ErrorContains(t, err, "below the configured minimum")
	require.Equal(t, model.ScheduleError, state)
}

func TestOnChainF05PaidRunScheduleRejectsMissingPaymentsFunds(t *testing.T) {
	ks, schedule := newF05PaidTestSchedule(t)
	manager := &OnChainF05Paid{
		keyStore:         ks,
		cfg:              defaultF05PaidSchedulingConfig(),
		balanceClient:    &balanceClientMock{balance: big.NewInt(1)},
		paymentsAddr:     common.HexToAddress("0x00000000000000000000000000000000000000bb"),
		providerRegistry: &providerRegistryMock{provider: newPaidProviderInfo()},
		paymentsContract: &paymentsContractMock{availableFunds: big.NewInt(0), approved: true},
	}

	state, err := manager.RunSchedule(context.Background(), schedule)
	require.ErrorContains(t, err, "has no available funds in payments contract")
	require.Equal(t, model.ScheduleError, state)
}

func TestOnChainF05PaidRunScheduleRejectsMissingOperatorApproval(t *testing.T) {
	ks, schedule := newF05PaidTestSchedule(t)
	manager := &OnChainF05Paid{
		keyStore:         ks,
		cfg:              defaultF05PaidSchedulingConfig(),
		balanceClient:    &balanceClientMock{balance: big.NewInt(1)},
		paymentsAddr:     common.HexToAddress("0x00000000000000000000000000000000000000bb"),
		providerRegistry: &providerRegistryMock{provider: newPaidProviderInfo()},
		paymentsContract: &paymentsContractMock{availableFunds: big.NewInt(100), approved: false},
	}

	state, err := manager.RunSchedule(context.Background(), schedule)
	require.ErrorContains(t, err, "has not approved provider")
	require.Equal(t, model.ScheduleError, state)
}

func TestOnChainF05PaidRunScheduleReturnsNotImplementedAfterPreflight(t *testing.T) {
	ks, schedule := newF05PaidTestSchedule(t)
	manager := &OnChainF05Paid{
		keyStore:         ks,
		cfg:              defaultF05PaidSchedulingConfig(),
		balanceClient:    &balanceClientMock{balance: big.NewInt(1)},
		paymentsAddr:     common.HexToAddress("0x00000000000000000000000000000000000000bb"),
		providerRegistry: &providerRegistryMock{provider: newPaidProviderInfo()},
		paymentsContract: &paymentsContractMock{availableFunds: big.NewInt(100), approved: true},
	}

	state, err := manager.RunSchedule(context.Background(), schedule)
	require.ErrorContains(t, err, "execution is not implemented yet")
	require.Equal(t, model.ScheduleError, state)
}
