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
	Choose(ctx context.Context, actors []model.Actor) (model.Actor, error)
}

type RandomWalletChooser struct{}

var ErrNoWallet = errors.New("no actors to choose from")

var ErrNoDatacap = errors.New("no actors have enough datacap")

// randomly selects an actor using cryptographically secure random number generator
func (w RandomWalletChooser) Choose(ctx context.Context, actors []model.Actor) (model.Actor, error) {
	if len(actors) == 0 {
		return model.Actor{}, ErrNoWallet
	}

	randomPick, err := rand.Int(rand.Reader, big.NewInt(int64(len(actors))))
	if err != nil {
		return model.Actor{}, errors.WithStack(err)
	}
	chosen := actors[randomPick.Int64()]
	return chosen, nil
}

type DatacapWalletChooser struct {
	db          *gorm.DB
	cache       *ttlcache.Cache[string, int64]
	lotusClient jsonrpc.RPCClient
	min         uint64
}

func NewDatacapWalletChooser(db *gorm.DB, cacheTTL time.Duration,
	lotusAPI string, lotusToken string, min uint64, //nolint:predeclared // We're ok with using the same name as the predeclared identifier here
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

func (w DatacapWalletChooser) getDatacap(ctx context.Context, actor model.Actor) (int64, error) {
	var result string
	err := w.lotusClient.CallFor(ctx, &result, "Filecoin.StateMarketBalance", actor.Address, nil)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return strconv.ParseInt(result, 10, 64)
}

func (w DatacapWalletChooser) getDatacapCached(ctx context.Context, actor model.Actor) (int64, error) {
	file := w.cache.Get(actor.Address)
	if file != nil && !file.IsExpired() {
		return file.Value(), nil
	}
	datacap, err := w.getDatacap(ctx, actor)
	if err != nil {
		logger.Errorf("failed to get datacap for actor %s: %s", actor.Address, err)
		if file != nil {
			return file.Value(), nil
		}
		return 0, errors.WithStack(err)
	}
	w.cache.Set(actor.Address, datacap, ttlcache.DefaultTTL)
	return datacap, nil
}

func (w DatacapWalletChooser) getPendingDeals(ctx context.Context, actor model.Actor) (int64, error) {
	var totalPieceSize int64
	err := w.db.WithContext(ctx).Model(&model.Deal{}).
		Select("COALESCE(SUM(piece_size), 0)").
		Where("client_id = ? AND verified AND state = ?", actor.ID, model.DealProposed).
		Scan(&totalPieceSize).
		Error
	if err != nil {
		logger.Errorf("failed to get pending deals for actor %s: %s", actor.Address, err)
		return 0, errors.WithStack(err)
	}
	return totalPieceSize, nil
}

// selects random actor with sufficient datacap (datacap - pending deals >= min threshold)
func (w DatacapWalletChooser) Choose(ctx context.Context, actors []model.Actor) (model.Actor, error) {
	if len(actors) == 0 {
		return model.Actor{}, ErrNoWallet
	}

	var eligible []model.Actor
	for _, actor := range actors {
		datacap, err := w.getDatacapCached(ctx, actor)
		if err != nil {
			logger.Errorw("failed to get datacap for actor", "actor", actor.Address, "error", err)
			continue
		}
		pendingDeals, err := w.getPendingDeals(ctx, actor)
		if err != nil {
			logger.Errorw("failed to get pending deals for actor", "actor", actor.Address, "error", err)
			continue
		}
		if datacap-pendingDeals >= int64(w.min) {
			eligible = append(eligible, actor)
		}
	}

	if len(eligible) == 0 {
		return model.Actor{}, ErrNoDatacap
	}

	randomPick, err := rand.Int(rand.Reader, big.NewInt(int64(len(eligible))))
	if err != nil {
		return model.Actor{}, errors.WithStack(err)
	}
	chosen := eligible[randomPick.Int64()]
	return chosen, nil
}
