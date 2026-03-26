package dealpusher

import (
	"context"
	"fmt"
	"math/big"

	synapse "github.com/data-preservation-programs/go-synapse"
	synpayments "github.com/data-preservation-programs/go-synapse/payments"
	"github.com/data-preservation-programs/go-synapse/spregistry"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type f05WalletBalanceClient interface {
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
}

type f05ProviderRegistry interface {
	GetProvider(ctx context.Context, providerID int) (*spregistry.ProviderInfo, error)
}

// OnChainF05Paid provides experimental registry and wallet-balance preflight
// for paid-f05 schedules while the final execution path is still being built.
type OnChainF05Paid struct {
	dbNoContext      *gorm.DB
	keyStore         keystore.KeyStore
	ethClient        *ethclient.Client
	balanceClient    f05WalletBalanceClient
	providerRegistry f05ProviderRegistry
	cfg              F05PaidSchedulingConfig
	network          synapse.Network
	chainID          *big.Int
	spRegistryAddr   common.Address
	paymentsAddr     common.Address
}

func NewOnChainF05Paid(
	ctx context.Context,
	db *gorm.DB,
	keyStore keystore.KeyStore,
	rpcURL string,
	cfg F05PaidSchedulingConfig,
) (*OnChainF05Paid, error) {
	if rpcURL == "" {
		return nil, fmt.Errorf("eth rpc URL is required")
	}
	if keyStore == nil {
		return nil, fmt.Errorf("keystore is required")
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	ethClient, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to FEVM RPC: %w", err)
	}

	network, chainIDInt64, err := synapse.DetectNetwork(ctx, ethClient)
	if err != nil {
		ethClient.Close()
		return nil, fmt.Errorf("failed to detect FEVM network: %w", err)
	}

	registryAddr := synapse.GetSPRegistryAddress(network)
	if cfg.SPRegistryAddress != "" {
		registryAddr, err = parseHexAddress(cfg.SPRegistryAddress)
		if err != nil {
			ethClient.Close()
			return nil, fmt.Errorf("invalid f05 SP registry contract address: %w", err)
		}
	}
	if registryAddr == (common.Address{}) {
		ethClient.Close()
		return nil, fmt.Errorf("no SP registry contract configured for network %s", network)
	}

	paymentsAddr, ok := synpayments.PaymentsAddresses[chainIDInt64]
	if cfg.PaymentsAddress != "" {
		paymentsAddr, err = parseHexAddress(cfg.PaymentsAddress)
		if err != nil {
			ethClient.Close()
			return nil, fmt.Errorf("invalid f05 payments contract address: %w", err)
		}
		ok = true
	}
	if !ok || paymentsAddr == (common.Address{}) {
		ethClient.Close()
		return nil, fmt.Errorf("no payments contract configured for chain ID %d", chainIDInt64)
	}

	providerRegistry, err := spregistry.NewService(ethClient, registryAddr, nil, big.NewInt(chainIDInt64))
	if err != nil {
		ethClient.Close()
		return nil, fmt.Errorf("failed to initialize SP registry client: %w", err)
	}

	Logger.Infow("initialized experimental paid f05 adapter",
		"network", network,
		"chainId", chainIDInt64,
		"spRegistry", registryAddr.Hex(),
		"payments", paymentsAddr.Hex(),
	)

	return &OnChainF05Paid{
		dbNoContext:      db,
		keyStore:         keyStore,
		ethClient:        ethClient,
		balanceClient:    ethClient,
		providerRegistry: providerRegistry,
		cfg:              cfg,
		network:          network,
		chainID:          big.NewInt(chainIDInt64),
		spRegistryAddr:   registryAddr,
		paymentsAddr:     paymentsAddr,
	}, nil
}

func (o *OnChainF05Paid) Close() error {
	if o.ethClient != nil {
		o.ethClient.Close()
	}
	return nil
}

func (o *OnChainF05Paid) RunSchedule(ctx context.Context, schedule *model.Schedule) (model.ScheduleState, error) {
	if err := o.cfg.Validate(); err != nil {
		return model.ScheduleError, fmt.Errorf("invalid paid f05 scheduling configuration: %w", err)
	}
	if schedule == nil {
		return model.ScheduleError, fmt.Errorf("schedule is required")
	}
	if schedule.Preparation == nil || schedule.Preparation.Wallet == nil {
		return model.ScheduleError, fmt.Errorf("schedule has no wallet configured")
	}

	walletObj := *schedule.Preparation.Wallet
	evmSigner, err := keystore.EVMSigner(o.keyStore, walletObj)
	if err != nil {
		return model.ScheduleError, fmt.Errorf("failed to load EVM signer for wallet: %w", err)
	}

	providerActorID, err := parseProviderActorID(schedule.Provider)
	if err != nil {
		return model.ScheduleError, fmt.Errorf("failed to parse provider actor ID: %w", err)
	}

	provider, err := o.providerRegistry.GetProvider(ctx, int(providerActorID))
	if err != nil {
		return model.ScheduleError, fmt.Errorf("failed to query provider %s in SP Registry: %w", schedule.Provider, err)
	}
	if provider == nil {
		return model.ScheduleError, fmt.Errorf("provider %s is not registered in SP Registry", schedule.Provider)
	}
	if !provider.Active {
		return model.ScheduleError, fmt.Errorf("provider %s is not active in SP Registry", schedule.Provider)
	}
	if provider.Payee == (common.Address{}) {
		return model.ScheduleError, fmt.Errorf("provider %s has no payee configured in SP Registry", schedule.Provider)
	}

	walletBalance, err := o.balanceClient.BalanceAt(ctx, evmSigner.EVMAddress(), nil)
	if err != nil {
		return model.ScheduleError, fmt.Errorf("failed to query FIL balance for wallet %s: %w", walletObj.Address, err)
	}
	if walletBalance.Sign() <= 0 {
		return model.ScheduleError, fmt.Errorf("wallet %s has no FIL balance available for paid f05 gas", walletObj.Address)
	}
	if walletBalance.Cmp(o.cfg.MinWalletBalanceAttoFIL) < 0 {
		return model.ScheduleError, fmt.Errorf(
			"wallet %s FIL balance %s is below the configured minimum %s",
			walletObj.Address,
			formatAttoFIL(walletBalance),
			formatAttoFIL(o.cfg.MinWalletBalanceAttoFIL),
		)
	}

	serviceURL := ""
	if product, ok := provider.Products["PDP"]; ok && product != nil && product.Data != nil {
		serviceURL = product.Data.ServiceURL
	}

	return model.ScheduleError, fmt.Errorf(
		"paid f05 schedule passed provider and FIL balance preflight (payee=%s, serviceURL=%s, payments=%s), but execution is not implemented yet",
		provider.Payee.Hex(),
		serviceURL,
		o.paymentsAddr.Hex(),
	)
}
