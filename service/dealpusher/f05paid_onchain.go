package dealpusher

import (
	"context"
	"fmt"
	"math/big"

	synapse "github.com/data-preservation-programs/go-synapse"
	"github.com/data-preservation-programs/go-synapse/contracts"
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

type f05PaymentsContract interface {
	GetAccountInfoIfSettled(ctx context.Context, token, owner common.Address) (fundedUntilEpoch, currentFunds, availableFunds, currentLockupRate *big.Int, err error)
	GetOperatorApproval(ctx context.Context, token, client, operator common.Address) (isApproved bool, rateAllowance, lockupAllowance, rateUsed, lockupUsed, maxLockupPeriod *big.Int, err error)
}

// OnChainF05Paid provides experimental registry and wallet-balance preflight
// for paid-f05 schedules while the final execution path is still being built.
type OnChainF05Paid struct {
	dbNoContext      *gorm.DB
	keyStore         keystore.KeyStore
	ethClient        *ethclient.Client
	balanceClient    f05WalletBalanceClient
	providerRegistry f05ProviderRegistry
	paymentsContract f05PaymentsContract
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

	paymentsContract, err := contracts.NewPaymentsContract(paymentsAddr, ethClient)
	if err != nil {
		ethClient.Close()
		return nil, fmt.Errorf("failed to initialize payments contract client: %w", err)
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
		paymentsContract: paymentsContract,
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
	if provider.ServiceProvider == (common.Address{}) {
		return model.ScheduleError, fmt.Errorf("provider %s has no service provider address configured in SP Registry", schedule.Provider)
	}
	if provider.Payee == (common.Address{}) {
		return model.ScheduleError, fmt.Errorf("provider %s has no payee configured in SP Registry", schedule.Provider)
	}

	product, ok := provider.Products["PDP"]
	if !ok || product == nil {
		return model.ScheduleError, fmt.Errorf("provider %s has no PDP product configured in SP Registry", schedule.Provider)
	}
	if !product.IsActive {
		return model.ScheduleError, fmt.Errorf("provider %s PDP product is not active in SP Registry", schedule.Provider)
	}
	if product.Data == nil {
		return model.ScheduleError, fmt.Errorf("provider %s PDP product is missing capability data in SP Registry", schedule.Provider)
	}
	if product.Data.ServiceURL == "" {
		return model.ScheduleError, fmt.Errorf("provider %s PDP product has no service URL configured in SP Registry", schedule.Provider)
	}
	if product.Data.PaymentTokenAddress == (common.Address{}) {
		return model.ScheduleError, fmt.Errorf("provider %s PDP product has no payment token configured in SP Registry", schedule.Provider)
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

	if o.paymentsContract == nil {
		return model.ScheduleError, fmt.Errorf("payments contract client is not configured")
	}

	_, _, availableFunds, _, err := o.paymentsContract.GetAccountInfoIfSettled(ctx, product.Data.PaymentTokenAddress, evmSigner.EVMAddress())
	if err != nil {
		return model.ScheduleError, fmt.Errorf("failed to query payments account for wallet %s: %w", walletObj.Address, err)
	}
	if availableFunds == nil || availableFunds.Sign() <= 0 {
		return model.ScheduleError, fmt.Errorf(
			"wallet %s has no available funds in payments contract %s for token %s",
			walletObj.Address,
			o.paymentsAddr.Hex(),
			product.Data.PaymentTokenAddress.Hex(),
		)
	}

	isApproved, rateAllowance, lockupAllowance, rateUsed, lockupUsed, maxLockupPeriod, err := o.paymentsContract.GetOperatorApproval(
		ctx,
		product.Data.PaymentTokenAddress,
		evmSigner.EVMAddress(),
		provider.ServiceProvider,
	)
	if err != nil {
		return model.ScheduleError, fmt.Errorf(
			"failed to query operator approval for provider %s service %s: %w",
			schedule.Provider,
			provider.ServiceProvider.Hex(),
			err,
		)
	}
	if !isApproved {
		return model.ScheduleError, fmt.Errorf(
			"wallet %s has not approved provider %s service %s on payments contract %s for token %s",
			walletObj.Address,
			schedule.Provider,
			provider.ServiceProvider.Hex(),
			o.paymentsAddr.Hex(),
			product.Data.PaymentTokenAddress.Hex(),
		)
	}

	return model.ScheduleError, fmt.Errorf(
		"paid f05 schedule passed provider, wallet, and payments preflight (payee=%s, service=%s, serviceURL=%s, token=%s, availableFunds=%s, rateAllowance=%s, rateUsed=%s, lockupAllowance=%s, lockupUsed=%s, maxLockupPeriod=%s, payments=%s), but execution is not implemented yet",
		provider.Payee.Hex(),
		provider.ServiceProvider.Hex(),
		product.Data.ServiceURL,
		product.Data.PaymentTokenAddress.Hex(),
		availableFunds.String(),
		rateAllowance.String(),
		rateUsed.String(),
		lockupAllowance.String(),
		lockupUsed.String(),
		maxLockupPeriod.String(),
		o.paymentsAddr.Hex(),
	)
}
