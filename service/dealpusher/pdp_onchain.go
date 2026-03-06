package dealpusher

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/cockroachdb/errors"
	synapse "github.com/data-preservation-programs/go-synapse"
	"github.com/data-preservation-programs/go-synapse/constants"
	"github.com/data-preservation-programs/go-synapse/pdp"
	"github.com/data-preservation-programs/go-synapse/signer"
	commcid "github.com/filecoin-project/go-fil-commcid"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"gorm.io/gorm"
)

// defaultGasLimit is a fixed gas limit for FEVM transactions.
// FEVM gas estimation (NoSend=true) is unreliable, so we use a fixed value.
// EVM traces show ~200K gas but FEVM execution needs ~17M due to actor overhead.
const defaultGasLimit = 30_000_000

type confirmationClient interface {
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	BlockNumber(ctx context.Context) (uint64, error)
}

// OnChainPDP implements PDP scheduling interfaces using FEVM RPC + go-synapse contracts.
type OnChainPDP struct {
	dbNoContext   *gorm.DB
	ethClient     *ethclient.Client
	confirmClient confirmationClient
	network       constants.Network
	chainID       *big.Int
	contractAddr  common.Address
}

func NewOnChainPDP(ctx context.Context, db *gorm.DB, rpcURL string) (*OnChainPDP, error) {
	return NewOnChainPDPWithAddress(ctx, db, rpcURL, common.Address{})
}

// NewOnChainPDPWithAddress is like NewOnChainPDP but accepts an optional
// contract address override. When addr is non-zero it is used instead of
// the default for the detected network (e.g. for devnet testing).
func NewOnChainPDPWithAddress(ctx context.Context, db *gorm.DB, rpcURL string, addr common.Address) (*OnChainPDP, error) {
	if rpcURL == "" {
		return nil, errors.New("eth rpc URL is required")
	}

	ethClient, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to FEVM RPC")
	}

	network, chainIDInt64, err := synapse.DetectNetwork(ctx, ethClient)
	if err != nil {
		ethClient.Close()
		return nil, errors.Wrap(err, "failed to detect FEVM network")
	}

	contractAddr := addr
	if contractAddr == (common.Address{}) {
		contractAddr = constants.GetPDPVerifierAddress(network)
	}
	if contractAddr == (common.Address{}) {
		ethClient.Close()
		return nil, fmt.Errorf("no PDPVerifier contract for network %s", network)
	}

	Logger.Infow("initialized PDP on-chain adapter",
		"network", network,
		"chainId", chainIDInt64,
		"contract", contractAddr.Hex(),
	)

	return &OnChainPDP{
		dbNoContext:   db,
		ethClient:     ethClient,
		confirmClient: ethClient,
		network:       network,
		chainID:       big.NewInt(chainIDInt64),
		contractAddr:  contractAddr,
	}, nil
}

func (o *OnChainPDP) Close() error {
	if o.ethClient != nil {
		o.ethClient.Close()
	}
	return nil
}

func (o *OnChainPDP) newManager(ctx context.Context, evmSigner signer.EVMSigner) (*pdp.Manager, error) {
	cfg := pdp.DefaultManagerConfig()
	cfg.ContractAddress = o.contractAddr
	cfg.DefaultGasLimit = defaultGasLimit
	mgr, err := pdp.NewManagerWithConfig(ctx, o.ethClient, evmSigner, o.network, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize PDP proof set manager")
	}
	return mgr, nil
}

func (o *OnChainPDP) EnsureProofSet(ctx context.Context, evmSigner signer.EVMSigner, provider string) (uint64, error) {
	clientAddr, err := commonToDelegatedAddress(evmSigner.EVMAddress())
	if err != nil {
		return 0, errors.Wrap(err, "failed to derive delegated client address from signer")
	}

	var existing model.PDPProofSet
	err = o.dbNoContext.WithContext(ctx).
		Where("client_address = ? AND provider = ? AND deleted = FALSE", clientAddr.String(), provider).
		Order("set_id").
		First(&existing).Error
	if err == nil {
		return existing.SetID, nil
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errors.Wrap(err, "failed to query existing PDP proof set")
	}

	manager, err := o.newManager(ctx, evmSigner)
	if err != nil {
		return 0, err
	}

	result, err := manager.CreateProofSet(ctx, pdp.CreateProofSetOptions{
		Value: pdp.SybilFee,
	})
	if err != nil {
		return 0, errors.Wrap(err, "failed to create PDP proof set")
	}

	setID := result.ProofSetID.Uint64()
	row := model.PDPProofSet{
		SetID:         setID,
		ClientAddress: clientAddr.String(),
		Provider:      provider,
		IsLive:        true,
	}
	if result.Receipt != nil && result.Receipt.BlockNumber != nil {
		row.CreatedBlock = int64(result.Receipt.BlockNumber.Uint64())
	}

	if err := o.dbNoContext.WithContext(ctx).Where("set_id = ?", setID).Assign(row).FirstOrCreate(&model.PDPProofSet{}).Error; err != nil {
		return 0, errors.Wrap(err, "failed to persist created PDP proof set")
	}
	return setID, nil
}

func (o *OnChainPDP) QueueAddRoots(
	ctx context.Context,
	evmSigner signer.EVMSigner,
	proofSetID uint64,
	pieceCIDs []cid.Cid,
	pieceSizes []int64,
	cfg PDPSchedulingConfig,
) (*PDPQueuedTx, error) {
	if len(pieceCIDs) == 0 {
		return nil, errors.New("no piece CIDs provided")
	}
	if len(pieceCIDs) != len(pieceSizes) {
		return nil, fmt.Errorf("pieceCIDs length %d != pieceSizes length %d", len(pieceCIDs), len(pieceSizes))
	}
	if cfg.BatchSize > 0 && len(pieceCIDs) > cfg.BatchSize {
		return nil, fmt.Errorf("piece CID count %d exceeds configured PDP batch size %d", len(pieceCIDs), cfg.BatchSize)
	}

	// convert commP v1 CIDs to CommPv2 format expected by PDPVerifier contract
	roots := make([]pdp.Root, len(pieceCIDs))
	for i, pieceCID := range pieceCIDs {
		payloadSize := uint64(pieceSizes[i]) * 127 / 128
		v2CID, err := commcid.PieceCidV2FromV1(pieceCID, payloadSize)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to convert piece CID %s to CommPv2", pieceCID)
		}
		roots[i] = pdp.Root{PieceCID: v2CID}
	}

	manager, err := o.newManager(ctx, evmSigner)
	if err != nil {
		return nil, err
	}

	setIDBig := new(big.Int).SetUint64(proofSetID)
	result, err := manager.AddRoots(ctx, setIDBig, roots)
	if err != nil {
		return nil, errors.Wrap(err, "failed to submit PDP add-roots transaction")
	}

	return &PDPQueuedTx{Hash: result.TransactionHash.Hex()}, nil
}

func (o *OnChainPDP) WaitForConfirmations(
	ctx context.Context,
	txHash string,
	depth uint64,
	pollInterval time.Duration,
) (*PDPTransactionReceipt, error) {
	rawHash, err := hexutil.Decode(txHash)
	if err != nil || len(rawHash) != common.HashLength {
		return nil, fmt.Errorf("invalid tx hash %q", txHash)
	}
	if pollInterval <= 0 {
		pollInterval = 2 * time.Second
	}

	hash := common.HexToHash(txHash)
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, errors.Wrap(ctx.Err(), "context canceled while waiting for PDP transaction confirmation")
		case <-ticker.C:
			receipt, err := o.confirmClient.TransactionReceipt(ctx, hash)
			if err != nil {
				if errors.Is(err, ethereum.NotFound) {
					continue
				}
				return nil, errors.Wrap(err, "failed to fetch transaction receipt")
			}

			out := toPDPReceipt(txHash, receipt)
			if receipt.Status != types.ReceiptStatusSuccessful {
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

func toPDPReceipt(txHash string, receipt *types.Receipt) *PDPTransactionReceipt {
	out := &PDPTransactionReceipt{
		Hash: txHash,
	}
	if receipt == nil {
		return out
	}
	out.Status = receipt.Status
	out.GasUsed = receipt.GasUsed
	if receipt.BlockNumber != nil {
		out.BlockNumber = receipt.BlockNumber.Uint64()
	}
	if receipt.EffectiveGasPrice != nil {
		out.CostAttoFIL = new(big.Int).Mul(new(big.Int).SetUint64(receipt.GasUsed), receipt.EffectiveGasPrice)
	} else {
		out.CostAttoFIL = big.NewInt(0)
	}
	return out
}

func commonToDelegatedAddress(subaddr common.Address) (address.Address, error) {
	addr, err := address.NewDelegatedAddress(10, subaddr.Bytes())
	if err != nil {
		return address.Undef, errors.Wrap(err, "failed to encode delegated address")
	}
	return addr, nil
}
