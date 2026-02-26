package replication

import (
	"context"
	"crypto/rand"
	"math/big"
	"strconv"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	logging "github.com/ipfs/go-log/v2"
	"github.com/jellydator/ttlcache/v3"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

var logger = logging.Logger("replication")

type WalletChooser interface {
	Choose(ctx context.Context, wallets []model.Wallet) (model.Wallet, error)
}

type RandomWalletChooser struct{}

var ErrNoWallet = errors.New("no wallets to choose from")

var ErrNoDatacap = errors.New("no wallets have enough datacap")

func (w RandomWalletChooser) Choose(_ context.Context, wallets []model.Wallet) (model.Wallet, error) {
	if len(wallets) == 0 {
		return model.Wallet{}, ErrNoWallet
	}

	randomPick, err := rand.Int(rand.Reader, big.NewInt(int64(len(wallets))))
	if err != nil {
		return model.Wallet{}, errors.WithStack(err)
	}
	return wallets[randomPick.Int64()], nil
}

// DatacapWalletChooser selects a wallet whose linked actor has sufficient datacap.
// only meaningful for market deals where datacap matters.
type DatacapWalletChooser struct {
	db          *gorm.DB
	cache       *ttlcache.Cache[string, int64]
	lotusClient jsonrpc.RPCClient
	min         uint64
}

func NewDatacapWalletChooser(db *gorm.DB, cacheTTL time.Duration,
	lotusAPI string, lotusToken string, min uint64, //nolint:predeclared
) DatacapWalletChooser {
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
	if wallet.ActorID == nil {
		return 0, errors.Newf("wallet %s has no linked actor", wallet.Address)
	}
	var result string
	err := w.lotusClient.CallFor(ctx, &result, "Filecoin.StateMarketBalance", wallet.Address, nil)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return strconv.ParseInt(result, 10, 64)
}

func (w DatacapWalletChooser) getDatacapCached(ctx context.Context, wallet model.Wallet) (int64, error) {
	file := w.cache.Get(wallet.Address)
	if file != nil && !file.IsExpired() {
		return file.Value(), nil
	}
	datacap, err := w.getDatacap(ctx, wallet)
	if err != nil {
		logger.Errorf("failed to get datacap for wallet %s: %s", wallet.Address, err)
		if file != nil {
			return file.Value(), nil
		}
		return 0, errors.WithStack(err)
	}
	w.cache.Set(wallet.Address, datacap, ttlcache.DefaultTTL)
	return datacap, nil
}

func (w DatacapWalletChooser) getPendingDeals(ctx context.Context, wallet model.Wallet) (int64, error) {
	var totalPieceSize int64
	err := w.db.WithContext(ctx).Model(&model.Deal{}).
		Select("COALESCE(SUM(piece_size), 0)").
		Where("wallet_id = ? AND verified AND state = ?", wallet.ID, model.DealProposed).
		Scan(&totalPieceSize).
		Error
	if err != nil {
		logger.Errorf("failed to get pending deals for wallet %s: %s", wallet.Address, err)
		return 0, errors.WithStack(err)
	}
	return totalPieceSize, nil
}

func (w DatacapWalletChooser) Choose(ctx context.Context, wallets []model.Wallet) (model.Wallet, error) {
	if len(wallets) == 0 {
		return model.Wallet{}, ErrNoWallet
	}

	var eligible []model.Wallet
	for _, wallet := range wallets {
		datacap, err := w.getDatacapCached(ctx, wallet)
		if err != nil {
			logger.Errorw("failed to get datacap for wallet", "address", wallet.Address, "error", err)
			continue
		}
		pendingDeals, err := w.getPendingDeals(ctx, wallet)
		if err != nil {
			logger.Errorw("failed to get pending deals for wallet", "address", wallet.Address, "error", err)
			continue
		}
		if datacap-pendingDeals >= int64(w.min) {
			eligible = append(eligible, wallet)
		}
	}

	if len(eligible) == 0 {
		return model.Wallet{}, ErrNoDatacap
	}

	randomPick, err := rand.Int(rand.Reader, big.NewInt(int64(len(eligible))))
	if err != nil {
		return model.Wallet{}, errors.WithStack(err)
	}
	return eligible[randomPick.Int64()], nil
}
