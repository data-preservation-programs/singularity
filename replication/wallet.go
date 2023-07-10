package replication

import (
	"context"
	"crypto/rand"
	"math/big"
	"strconv"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	logging "github.com/ipfs/go-log/v2"
	"github.com/jellydator/ttlcache/v3"
	"github.com/pkg/errors"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

var logger = logging.Logger("replication")

type WalletChooser interface {
	Choose(ctx context.Context, wallets []model.Wallet) (model.Wallet, error)
}

type DefaultWalletChooser struct{}

var ErrNoWallet = errors.New("no wallets to choose from")

var ErrNoDatacap = errors.New("no wallets have enough datacap")

func (w DefaultWalletChooser) Choose(ctx context.Context, wallets []model.Wallet) (model.Wallet, error) {
	// Check if the wallets slice is empty
	if len(wallets) == 0 {
		return model.Wallet{}, ErrNoWallet
	}

	randomPick, err := rand.Int(rand.Reader, big.NewInt(int64(len(wallets))))
	if err != nil {
		return model.Wallet{}, err
	}
	chosenWallet := wallets[randomPick.Int64()]
	return chosenWallet, nil
}

type DatacapWalletChooser struct {
	db          *gorm.DB
	cache       *ttlcache.Cache[string, int64]
	lotusClient jsonrpc.RPCClient
	min         uint64
}

func NewDatacapWalletChooser(db *gorm.DB, cacheTTL time.Duration,
	lotusAPI string, lotusToken string, min uint64) DatacapWalletChooser {
	cache := ttlcache.New[string, int64](
		ttlcache.WithTTL[string, int64](cacheTTL),
		ttlcache.WithDisableTouchOnHit[string, int64]())

	lotusClient := util.NewLotusClient(lotusAPI, lotusToken)
	return DatacapWalletChooser{
		db:          db,
		cache:       cache,
		lotusClient: lotusClient,
		min:         min,
	}
}

func (w DatacapWalletChooser) getDatacap(ctx context.Context, wallet model.Wallet) (int64, error) {
	var result string
	err := w.lotusClient.CallFor(ctx, &result, "Filecoin.StateMarketBalance", wallet.Address, nil)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(result, 10, 64)
}

func (w DatacapWalletChooser) getDatacapCached(ctx context.Context, wallet model.Wallet) (int64, error) {
	item := w.cache.Get(wallet.Address)
	if item != nil && !item.IsExpired() {
		return item.Value(), nil
	}
	datacap, err := w.getDatacap(ctx, wallet)
	if err != nil {
		logger.Errorf("failed to get datacap for wallet %s: %s", wallet.Address, err)
		if item != nil {
			return item.Value(), nil
		}
		return 0, err
	}
	w.cache.Set(wallet.Address, datacap, ttlcache.DefaultTTL)
	return datacap, nil
}

func (w DatacapWalletChooser) getPendingDeals(ctx context.Context, wallet model.Wallet) (int64, error) {
	var totalPieceSize int64
	err := w.db.WithContext(ctx).Model(&model.Deal{}).
		Select("COALESCE(SUM(piece_size), 0)").
		Where("client_id = ? AND verified AND state = ?", wallet.ID, model.DealProposed).
		Scan(&totalPieceSize).
		Error
	if err != nil {
		logger.Errorf("failed to get pending deals for wallet %s: %s", wallet.Address, err)
		return 0, err
	}
	return totalPieceSize, nil
}

func (w DatacapWalletChooser) Choose(ctx context.Context, wallets []model.Wallet) (model.Wallet, error) {
	if len(wallets) == 0 {
		return model.Wallet{}, ErrNoWallet
	}

	var eligibleWallets []model.Wallet
	for _, wallet := range wallets {
		datacap, err := w.getDatacapCached(ctx, wallet)
		if err != nil {
			logger.Errorw("failed to get datacap for wallet", "wallet", wallet.Address, "error", err)
			continue
		}
		pendingDeals, err := w.getPendingDeals(ctx, wallet)
		if err != nil {
			logger.Errorw("failed to get pending deals for wallet", "wallet", wallet.Address, "error", err)
			continue
		}
		if datacap-pendingDeals >= int64(w.min) {
			eligibleWallets = append(eligibleWallets, wallet)
		}
	}

	if len(eligibleWallets) == 0 {
		return model.Wallet{}, ErrNoDatacap
	}

	randomPick, err := rand.Int(rand.Reader, big.NewInt(int64(len(eligibleWallets))))
	if err != nil {
		return model.Wallet{}, err
	}
	chosenWallet := eligibleWallets[randomPick.Int64()]
	return chosenWallet, nil
}
