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

type RandomWalletChooser struct{}

var ErrNoWallet = errors.New("no wallets to choose from")

var ErrNoDatacap = errors.New("no wallets have enough datacap")

// Choose selects a random Wallet from the provided slice of Wallets.
//
// The Choose function of the RandomWalletChooser type randomly selects
// a Wallet from a given slice of Wallets. If the slice is empty, the function
// returns an error. It uses a cryptographically secure random number generator
// to make the selection.
//
// Parameters:
//   - ctx context.Context: The context to use for cancellation and deadlines,
//     although it is not used in this implementation.
//   - wallets []model.Wallet: A slice of Wallet objects from which a random Wallet
//     will be chosen.
//
// Returns:
//   - model.Wallet: The randomly chosen Wallet object from the provided slice.
//   - error: An error that will be returned if any issues were encountered while trying
//     to choose a Wallet. This includes the case when the input slice is empty,
//     in which case ErrNoWallet will be returned, or if there is an issue generating
//     a random number.
func (w RandomWalletChooser) Choose(ctx context.Context, wallets []model.Wallet) (model.Wallet, error) {
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

// Choose selects a random Wallet from the provided slice of Wallets based on certain criteria.
//
// The Choose function of the DatacapWalletChooser type filters the given slice of Wallets
// based on a specific criterion, which is whether the datacap for the wallet minus
// the pending deals for the wallet is greater or equal to a minimum threshold (w.min).
// From the filtered eligible Wallets, the function then randomly selects one Wallet.
// It uses a cryptographically secure random number generator to make the selection.
// If the initial slice of Wallets is empty, or if no Wallets meet the criteria,
// the function returns an error.
//
// Parameters:
//   - ctx context.Context: The context to use for cancellation and deadlines, used
//     in the datacap and pending deals fetching operations.
//   - wallets []model.Wallet: A slice of Wallet objects from which a random Wallet
//     will be chosen based on the criteria.
//
// Returns:
//   - model.Wallet: The randomly chosen Wallet object from the filtered eligible Wallets.
//   - error: An error that will be returned if any issues were encountered while trying
//     to choose a Wallet. This includes the case when the input slice is empty,
//     in which case ErrNoWallet will be returned, when no Wallets meet the criteria,
//     in which case ErrNoDatacap will be returned, or if there is an issue generating
//     a random number.
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
