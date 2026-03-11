package dealpusher

import (
	"context"
	"fmt"
	"time"

	ddocontract "github.com/Eastore-project/ddo-client/pkg/contract/ddo"
	paymentscontract "github.com/Eastore-project/ddo-client/pkg/contract/payments"
	ddotypes "github.com/Eastore-project/ddo-client/pkg/types"
	ddoutils "github.com/Eastore-project/ddo-client/pkg/utils"
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/go-synapse/signer"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

type ddoConfirmationClient interface {
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*ethtypes.Receipt, error)
	BlockNumber(ctx context.Context) (uint64, error)
}

// OnChainDDO implements the DDO scheduling interfaces using ddo-client.
//
// The current upstream SDK still requires a raw private key for write clients.
// Until it accepts a signer/transactor directly, this adapter must be bound to
// a specific EVM key and rejects mismatched signers at call time.
type OnChainDDO struct {
	ddoClient       *ddocontract.Client
	paymentsClient  *paymentscontract.Client
	confirmClient   ddoConfirmationClient
	rpcURL          string
	privateKey      string
	ddoContractAddr common.Address
	paymentToken    common.Address
	signerAddr      common.Address
}

func NewOnChainDDO(
	ctx context.Context,
	rpcURL string,
	ddoAddr, paymentsAddr, payToken, privateKey string,
) (*OnChainDDO, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if rpcURL == "" {
		return nil, errors.New("eth rpc URL is required")
	}
	if ddoAddr == "" {
		return nil, errors.New("ddo contract address is required")
	}
	if paymentsAddr == "" {
		return nil, errors.New("ddo payments contract address is required")
	}
	if payToken == "" {
		return nil, errors.New("ddo payment token address is required")
	}
	if privateKey == "" {
		return nil, errors.New("ddo private key is required")
	}

	ecdsaKey, err := ethcrypto.HexToECDSA(trimHexPrefix(privateKey))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse DDO private key")
	}
	signerAddr := ethcrypto.PubkeyToAddress(ecdsaKey.PublicKey)

	ddoClient, err := ddocontract.NewClientWithParams(rpcURL, ddoAddr, privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize DDO contract client")
	}

	paymentsClient, err := paymentscontract.NewClientWithParams(rpcURL, paymentsAddr, privateKey)
	if err != nil {
		ddoClient.Close()
		return nil, errors.Wrap(err, "failed to initialize DDO payments client")
	}

	return &OnChainDDO{
		ddoClient:       ddoClient,
		paymentsClient:  paymentsClient,
		confirmClient:   ddoClient.GetEthClient(),
		rpcURL:          rpcURL,
		privateKey:      privateKey,
		ddoContractAddr: common.HexToAddress(ddoAddr),
		paymentToken:    common.HexToAddress(payToken),
		signerAddr:      signerAddr,
	}, nil
}

func (o *OnChainDDO) Close() {
	if o.ddoClient != nil {
		o.ddoClient.Close()
	}
	if o.paymentsClient != nil {
		o.paymentsClient.Close()
	}
}

func (o *OnChainDDO) ValidateSP(ctx context.Context, providerActorID uint64) (*DDOSPConfig, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

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
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := o.validateSigner(evmSigner); err != nil {
		return err
	}

	pieceInfos, err := o.buildPieceInfos(pieces, cfg)
	if err != nil {
		return err
	}

	if err := ddoutils.CheckAndSetupPayments(
		o.ddoClient.GetEthClient(),
		o.ddoClient,
		o.paymentsClient,
		pieceInfos,
		evmSigner.EVMAddress(),
		o.ddoContractAddr,
		o.rpcURL,
		o.privateKey,
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
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if err := o.validateSigner(evmSigner); err != nil {
		return nil, err
	}

	pieceInfos, err := o.buildPieceInfos(pieces, cfg)
	if err != nil {
		return nil, err
	}

	txHash, err := o.ddoClient.CreateAllocationRequests(pieceInfos)
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

	for {
		select {
		case <-ctx.Done():
			return nil, errors.Wrap(ctx.Err(), "context canceled while waiting for DDO transaction confirmation")
		case <-ticker.C:
			receipt, err := o.confirmClient.TransactionReceipt(ctx, hash)
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

			currentBlock, err := o.confirmClient.BlockNumber(ctx)
			if err != nil {
				return nil, errors.Wrap(err, "failed to fetch latest block number")
			}
			if receipt.BlockNumber != nil && receipt.BlockNumber.Uint64()+depth <= currentBlock {
				return out, nil
			}
		}
	}
}

func (o *OnChainDDO) ParseAllocationIDs(ctx context.Context, txHash string) ([]uint64, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

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
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	info, err := o.ddoClient.GetAllocationInfo(allocationID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch DDO allocation %d", allocationID)
	}
	return &DDOAllocationStatus{
		Activated:    info.Activated,
		SectorNumber: info.SectorNumber,
	}, nil
}

func (o *OnChainDDO) validateSigner(evmSigner signer.EVMSigner) error {
	if evmSigner == nil {
		return errors.New("evm signer is required")
	}
	if evmSigner.EVMAddress() != o.signerAddr {
		return fmt.Errorf(
			"configured DDO adapter signer %s does not match requested signer %s",
			o.signerAddr.Hex(),
			evmSigner.EVMAddress().Hex(),
		)
	}
	return nil
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

func trimHexPrefix(value string) string {
	if len(value) >= 2 && value[:2] == "0x" {
		return value[2:]
	}
	return value
}

var _ DDODealManager = (*OnChainDDO)(nil)
var _ DDOAllocationTracker = (*OnChainDDO)(nil)
