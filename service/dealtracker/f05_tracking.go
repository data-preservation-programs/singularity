package dealtracker

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type f05ReceiptClient interface {
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*ethtypes.Receipt, error)
}

type F05PaymentReceipt struct {
	Status      uint64
	BlockNumber uint64
	GasUsed     uint64
}

type F05PaymentTrackingClient struct {
	client        *ethclient.Client
	receiptClient f05ReceiptClient
}

func NewF05PaymentTrackingClient(ctx context.Context, rpcURL string) (*F05PaymentTrackingClient, error) {
	if rpcURL == "" {
		return nil, errors.New("eth rpc URL is required")
	}
	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize paid f05 tracking client")
	}
	return &F05PaymentTrackingClient{
		client:        client,
		receiptClient: client,
	}, nil
}

func (c *F05PaymentTrackingClient) GetTransactionReceipt(ctx context.Context, txHash string) (*F05PaymentReceipt, error) {
	hash, err := parseReceiptTxHash(txHash)
	if err != nil {
		return nil, err
	}

	receipt, err := c.receiptClient.TransactionReceipt(ctx, hash)
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "failed to fetch paid f05 transaction receipt %s", txHash)
	}
	return toF05PaymentReceipt(receipt), nil
}

func (c *F05PaymentTrackingClient) Close() {
	if c.client != nil {
		c.client.Close()
	}
}

func parseReceiptTxHash(txHash string) (common.Hash, error) {
	rawHash, err := hexutil.Decode(txHash)
	if err != nil || len(rawHash) != common.HashLength {
		return common.Hash{}, fmt.Errorf("invalid tx hash %q", txHash)
	}
	return common.HexToHash(txHash), nil
}

func toF05PaymentReceipt(receipt *ethtypes.Receipt) *F05PaymentReceipt {
	out := &F05PaymentReceipt{}
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
