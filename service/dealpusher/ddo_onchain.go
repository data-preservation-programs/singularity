package dealpusher

import (
	"context"
	"fmt"
	"math/big"
	"time"

	ddocontract "github.com/Eastore-project/ddo-client/pkg/contract/ddo"
	paymentscontract "github.com/Eastore-project/ddo-client/pkg/contract/payments"
	ddotypes "github.com/Eastore-project/ddo-client/pkg/types"
	ddoutils "github.com/Eastore-project/ddo-client/pkg/utils"
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/go-synapse/signer"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type ddoConfirmationClient interface {
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*ethtypes.Receipt, error)
	BlockNumber(ctx context.Context) (uint64, error)
}

// OnChainDDO implements the DDO scheduling interfaces using ddo-client.
// It keeps a read-only DDO client for queries and confirmation polling, and
// derives write-capable DDO/payments clients from the caller's EVMSigner.
//
// caveat: ddo-client (pinned at 2fff1a5b168a) overwrites auth.Context with
// context.Background() in its write methods, so ctx cancellation does not
// propagate to in-flight EnsurePayments/CreateAllocations calls.
type OnChainDDO struct {
	ddoClient            *ddocontract.Client
	confirmClient        ddoConfirmationClient
	chainID              *big.Int
	ddoContractAddr      common.Address
	paymentsContractAddr common.Address
	paymentToken         common.Address
}

func NewOnChainDDO(
	ctx context.Context,
	rpcURL string,
	ddoAddr, paymentsAddr, payToken string,
) (*OnChainDDO, error) {
	if rpcURL == "" {
		return nil, errors.New("eth rpc URL is required")
	}
	ddoAddress, err := parseHexAddress(ddoAddr)
	if err != nil {
		return nil, errors.Wrap(err, "invalid ddo contract address")
	}
	paymentsAddress, err := parseHexAddress(paymentsAddr)
	if err != nil {
		return nil, errors.Wrap(err, "invalid ddo payments contract address")
	}
	tokenAddress, err := parseHexAddress(payToken)
	if err != nil {
		return nil, errors.Wrap(err, "invalid ddo payment token address")
	}

	ddoClient, err := ddocontract.NewReadOnlyClientWithParams(rpcURL, ddoAddr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize DDO contract client")
	}

	chainID, err := ddoClient.GetEthClient().ChainID(ctx)
	if err != nil {
		ddoClient.Close()
		return nil, errors.Wrap(err, "failed to detect EVM chain ID")
	}

	return &OnChainDDO{
		ddoClient:            ddoClient,
		confirmClient:        ddoClient.GetEthClient(),
		chainID:              chainID,
		ddoContractAddr:      ddoAddress,
		paymentsContractAddr: paymentsAddress,
		paymentToken:         tokenAddress,
	}, nil
}

func (o *OnChainDDO) Close() {
	if o.ddoClient != nil {
		o.ddoClient.Close()
	}
}

func (o *OnChainDDO) ValidateSP(ctx context.Context, providerActorID uint64) (*DDOSPConfig, error) {
	cfg, err := o.ddoClient.GetSPConfig(providerActorID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch DDO SP config for provider %d", providerActorID)
	}
	if cfg == nil {
		return &DDOSPConfig{}, nil
	}
	if cfg.IsActive && !supportsPaymentToken(cfg.SupportedTokens, o.paymentToken) {
		return nil, fmt.Errorf("provider %d does not support payment token %s", providerActorID, o.paymentToken.Hex())
	}

	return &DDOSPConfig{
		IsActive:     cfg.IsActive,
		MinPieceSize: cfg.MinPieceSize,
		MaxPieceSize: cfg.MaxPieceSize,
		MinTermLen:   cfg.MinTermLength,
		MaxTermLen:   cfg.MaxTermLength,
	}, nil
}

func (o *OnChainDDO) EnsurePayments(
	ctx context.Context,
	evmSigner signer.EVMSigner,
	pieces []DDOPieceSubmission,
	cfg DDOSchedulingConfig,
) error {
	auth, err := o.newTransactor(ctx, evmSigner)
	if err != nil {
		return err
	}

	pieceInfos, err := o.buildPieceInfos(pieces, cfg)
	if err != nil {
		return err
	}

	ddoWriteClient, err := o.newDDOWriteClient(auth)
	if err != nil {
		return err
	}
	paymentsWriteClient, err := o.newPaymentsWriteClient(auth)
	if err != nil {
		return err
	}

	if err := ddoutils.CheckAndSetupPayments(
		o.ddoClient.GetEthClient(),
		ddoWriteClient,
		paymentsWriteClient,
		pieceInfos,
		evmSigner.EVMAddress(),
		o.ddoContractAddr,
		auth,
	); err != nil {
		return errors.Wrap(err, "failed to ensure DDO payments")
	}
	return nil
}

func (o *OnChainDDO) CreateAllocations(
	ctx context.Context,
	evmSigner signer.EVMSigner,
	pieces []DDOPieceSubmission,
	cfg DDOSchedulingConfig,
) (*DDOQueuedTx, error) {
	auth, err := o.newTransactor(ctx, evmSigner)
	if err != nil {
		return nil, err
	}

	pieceInfos, err := o.buildPieceInfos(pieces, cfg)
	if err != nil {
		return nil, err
	}

	ddoWriteClient, err := o.newDDOWriteClient(auth)
	if err != nil {
		return nil, err
	}

	txHash, err := ddoWriteClient.CreateAllocationRequests(pieceInfos)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create DDO allocation requests")
	}
	return &DDOQueuedTx{Hash: txHash}, nil
}

func (o *OnChainDDO) WaitForConfirmations(
	ctx context.Context,
	txHash string,
	depth uint64,
	pollInterval time.Duration,
) (*DDOTransactionReceipt, error) {
	hash, err := parseTxHash(txHash)
	if err != nil {
		return nil, err
	}
	if pollInterval <= 0 {
		pollInterval = 2 * time.Second
	}

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	var receipt *ethtypes.Receipt

	for {
		select {
		case <-ctx.Done():
			return nil, errors.Wrap(ctx.Err(), "context canceled while waiting for DDO transaction confirmation")
		case <-ticker.C:
			if receipt == nil {
				receipt, err = o.confirmClient.TransactionReceipt(ctx, hash)
				if err != nil {
					if errors.Is(err, ethereum.NotFound) {
						continue
					}
					return nil, errors.Wrap(err, "failed to fetch DDO transaction receipt")
				}

				out := toDDOReceipt(txHash, receipt)
				if receipt.Status != ethtypes.ReceiptStatusSuccessful {
					return out, fmt.Errorf("transaction %s failed with status %d", txHash, receipt.Status)
				}
				if depth == 0 {
					return out, nil
				}
			}

			currentBlock, err := o.confirmClient.BlockNumber(ctx)
			if err != nil {
				return nil, errors.Wrap(err, "failed to fetch latest block number")
			}
			if receipt.BlockNumber != nil && receipt.BlockNumber.Uint64()+depth <= currentBlock {
				return toDDOReceipt(txHash, receipt), nil
			}
		}
	}
}

func (o *OnChainDDO) ParseAllocationIDs(ctx context.Context, txHash string) ([]uint64, error) {
	hash, err := parseTxHash(txHash)
	if err != nil {
		return nil, err
	}

	receipt, err := o.confirmClient.TransactionReceipt(ctx, hash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch DDO transaction receipt for event parsing")
	}

	allocationIDs, err := ddocontract.ParseAllocationCreatedEvents(receipt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse DDO allocation IDs from receipt")
	}
	return allocationIDs, nil
}

func (o *OnChainDDO) GetAllocationInfo(ctx context.Context, allocationID uint64) (*DDOAllocationStatus, error) {
	info, err := o.ddoClient.GetAllocationInfo(allocationID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch DDO allocation %d", allocationID)
	}
	return &DDOAllocationStatus{
		Activated:    info.Activated,
		SectorNumber: info.SectorNumber,
	}, nil
}

func (o *OnChainDDO) newTransactor(ctx context.Context, evmSigner signer.EVMSigner) (*bind.TransactOpts, error) {
	if evmSigner == nil {
		return nil, errors.New("evm signer is required")
	}
	if o.chainID == nil {
		return nil, errors.New("ddo adapter chain ID is not initialized")
	}

	auth, err := evmSigner.Transactor(o.chainID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create DDO transactor from signer")
	}
	auth.Context = ctx
	return auth, nil
}

func (o *OnChainDDO) newDDOWriteClient(auth *bind.TransactOpts) (*ddocontract.Client, error) {
	if o.ddoClient == nil || o.ddoClient.GetEthClient() == nil {
		return nil, errors.New("ddo eth client is not initialized")
	}
	client, err := ddocontract.NewClientWithTransactor(o.ddoClient.GetEthClient(), o.ddoContractAddr.Hex(), auth)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize write-capable DDO client")
	}
	return client, nil
}

func (o *OnChainDDO) newPaymentsWriteClient(auth *bind.TransactOpts) (*paymentscontract.Client, error) {
	if o.ddoClient == nil || o.ddoClient.GetEthClient() == nil {
		return nil, errors.New("ddo eth client is not initialized")
	}
	client, err := paymentscontract.NewClientWithTransactor(o.ddoClient.GetEthClient(), o.paymentsContractAddr.Hex(), auth)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize write-capable DDO payments client")
	}
	return client, nil
}

func (o *OnChainDDO) buildPieceInfos(pieces []DDOPieceSubmission, cfg DDOSchedulingConfig) ([]ddotypes.PieceInfo, error) {
	if len(pieces) == 0 {
		return nil, errors.New("no DDO pieces provided")
	}
	if err := cfg.Validate(); err != nil {
		return nil, errors.Wrap(err, "invalid DDO scheduling configuration")
	}

	pieceInfos := make([]ddotypes.PieceInfo, len(pieces))
	for i, piece := range pieces {
		if !piece.PieceCID.Defined() {
			return nil, fmt.Errorf("piece %d has undefined piece CID", i)
		}
		if piece.PieceSize == 0 {
			return nil, fmt.Errorf("piece %d has invalid piece size 0", i)
		}
		if piece.ProviderID == 0 {
			return nil, fmt.Errorf("piece %d has invalid provider ID 0", i)
		}
		pieceInfos[i] = ddotypes.PieceInfo{
			PieceCid:            piece.PieceCID.Bytes(),
			Size:                piece.PieceSize,
			Provider:            piece.ProviderID,
			TermMin:             cfg.TermMin,
			TermMax:             cfg.TermMax,
			ExpirationOffset:    cfg.ExpirationOffset,
			DownloadURL:         piece.DownloadURL,
			PaymentTokenAddress: o.paymentToken,
		}
	}
	return pieceInfos, nil
}

// parseHexAddress validates and converts a hex string to common.Address.
// common.HexToAddress silently maps malformed input to 0x0; this fails fast.
func parseHexAddress(s string) (common.Address, error) {
	if !common.IsHexAddress(s) {
		return common.Address{}, fmt.Errorf("invalid hex address %q", s)
	}
	return common.HexToAddress(s), nil
}

func parseTxHash(txHash string) (common.Hash, error) {
	rawHash, err := hexutil.Decode(txHash)
	if err != nil || len(rawHash) != common.HashLength {
		return common.Hash{}, fmt.Errorf("invalid tx hash %q", txHash)
	}
	return common.HexToHash(txHash), nil
}

func toDDOReceipt(txHash string, receipt *ethtypes.Receipt) *DDOTransactionReceipt {
	out := &DDOTransactionReceipt{Hash: txHash}
	if receipt == nil {
		return out
	}
	out.Status = receipt.Status
	out.GasUsed = receipt.GasUsed
	if receipt.BlockNumber != nil {
		out.BlockNumber = receipt.BlockNumber.Uint64()
	}
	return out
}

func supportsPaymentToken(tokens []ddotypes.TokenConfig, token common.Address) bool {
	if len(tokens) == 0 {
		return false
	}
	for _, cfg := range tokens {
		if cfg.Token == token && cfg.IsActive {
			return true
		}
	}
	return false
}

var _ DDODealManager = (*OnChainDDO)(nil)
var _ DDOAllocationTracker = (*OnChainDDO)(nil)
