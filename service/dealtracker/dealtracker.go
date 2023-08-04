package dealtracker

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/data-preservation-programs/singularity/service/epochutil"
	"github.com/data-preservation-programs/singularity/service/healthcheck"
	"github.com/dustin/go-humanize"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"

	"github.com/bcicen/jstream"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-log/v2"
	"github.com/klauspost/compress/zstd"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Deal struct {
	Proposal DealProposal
	State    DealState
}

func (d Deal) Key() string {
	return fmt.Sprintf("%s-%s-%s-%d-%d", d.Proposal.Client, d.Proposal.Provider,
		d.Proposal.PieceCID.Root, d.Proposal.StartEpoch, d.Proposal.EndEpoch)
}

func (d Deal) GetState() model.DealState {
	if d.State.SlashEpoch > 0 {
		return model.DealSlashed
	}
	if d.State.SectorStartEpoch < 0 {
		if epochutil.EpochToTime(d.Proposal.StartEpoch).Before(time.Now().Add(-24 * time.Hour)) {
			return model.DealProposalExpired
		}
		return model.DealPublished
	}
	if epochutil.EpochToTime(d.Proposal.EndEpoch).Before(time.Now().Add(-24 * time.Hour)) {
		return model.DealExpired
	}
	return model.DealActive
}

type Cid struct {
	Root string `json:"/" mapstructure:"/"`
}

type DealProposal struct {
	PieceCID             Cid
	PieceSize            int64
	VerifiedDeal         bool
	Client               string
	Provider             string
	Label                string
	StartEpoch           int32
	EndEpoch             int32
	StoragePricePerEpoch string
}

type DealState struct {
	SectorStartEpoch int32
	LastUpdatedEpoch int32
	SlashEpoch       int32
}

type CloserFunc func() error

func (c CloserFunc) Close() error {
	return c()
}

var logger = log.Logger("dealtracker")

type DealTracker struct {
	workerID   uuid.UUID
	db         *gorm.DB
	interval   time.Duration
	dealZstURL string
	lotusURL   string
	lotusToken string
}

func NewDealTracker(
	db *gorm.DB,
	interval time.Duration,
	dealZstURL string,
	lotusURL string,
	lotusToken string) DealTracker {
	return DealTracker{
		workerID:   uuid.New(),
		db:         db,
		interval:   interval,
		dealZstURL: dealZstURL,
		lotusURL:   lotusURL,
		lotusToken: lotusToken,
	}
}

type threadSafeReadCloser struct {
	reader io.Reader
	closer func()
	closed bool
	mu     sync.Mutex
}

func (t *threadSafeReadCloser) Read(p []byte) (n int, err error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.closed {
		return 0, errors.New("closed")
	}
	return t.reader.Read(p)
}

func (t *threadSafeReadCloser) Close() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.closed = true
	t.closer()
}

func DealStateStreamFromHTTPRequest(request *http.Request, depth int, decompress bool) (chan *jstream.MetaValue, Counter, io.Closer, error) {
	//nolint: bodyclose
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to get deal state from lotus API")
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, nil, nil, errors.New("failed to get deal state: " + resp.Status)
	}
	var jsonDecoder *jstream.Decoder
	var closer io.Closer
	countingReader := NewCountingReader(resp.Body)
	if decompress {
		decompressor, err := zstd.NewReader(countingReader)
		if err != nil {
			resp.Body.Close()
			return nil, nil, nil, errors.Wrap(err, "failed to create zstd decompressor")
		}
		safeDecompressor := &threadSafeReadCloser{
			reader: decompressor,
			closer: decompressor.Close,
		}
		jsonDecoder = jstream.NewDecoder(safeDecompressor, depth).EmitKV()
		closer = CloserFunc(func() error {
			safeDecompressor.Close()
			return resp.Body.Close()
		})
	} else {
		jsonDecoder = jstream.NewDecoder(countingReader, depth).EmitKV()
		closer = resp.Body
	}

	return jsonDecoder.Stream(), countingReader, closer, nil
}

func (d *DealTracker) dealStateStream(ctx context.Context) (chan *jstream.MetaValue, Counter, io.Closer, error) {
	if d.dealZstURL != "" {
		logger.Infof("getting deal state from %s", d.dealZstURL)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, d.dealZstURL, nil)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "failed to create request to get deal state zst file")
		}
		return DealStateStreamFromHTTPRequest(req, 1, true)
	}

	logger.Infof("getting deal state from %s", d.lotusURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, d.lotusURL, nil)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to create request to get deal state from lotus API")
	}
	if d.lotusToken != "" {
		req.Header.Set("Authorization", "Bearer "+d.lotusToken)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(strings.NewReader(`{"jsonrpc":"2.0","method":"Filecoin.StateMarketDeals","params":[null],"id":0}`))
	return DealStateStreamFromHTTPRequest(req, 2, false)
}

func (d *DealTracker) Run(ctx context.Context, once bool) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	getState := func() healthcheck.State {
		return healthcheck.State{
			WorkType: model.DealTracking,
		}
	}

	for {
		alreadyRunning, err := healthcheck.Register(ctx, d.db, d.workerID, getState, false)
		if err == nil && !alreadyRunning {
			break
		}
		if err != nil {
			logger.Errorw("failed to register worker", "error", err)
			if once {
				return err
			}
		}
		if alreadyRunning {
			logger.Warnw("another worker already running")
			if once {
				return nil
			}
		}
		logger.Warn("retrying in 1 minute")
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Minute):
		}
	}

	go healthcheck.StartReportHealth(ctx, d.db, d.workerID, getState)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGTRAP)

	go func() {
		for {
			err := d.runOnce(ctx)
			if err != nil {
				logger.Errorw("failed to run deal maker", "error", err)
			}
			if once {
				cancel()
				return
			}
			select {
			case <-ctx.Done():
				return
			case <-time.After(d.interval):
			}
		}
	}()

	select {
	case <-ctx.Done():
		//nolint:errcheck
		d.cleanup()
		return ctx.Err()
	case <-signalChan:
		//nolint:errcheck
		d.cleanup()
		return cli.Exit("received signal", 130)
	}
}

func (d *DealTracker) cleanup() error {
	return d.db.Where("id = ?", d.workerID).Delete(&model.Worker{}).Error
}

type KnownDeal struct {
	State model.DealState
}
type UnknownDeal struct {
	ID         uint64
	ClientID   string
	Provider   string
	PieceCID   model.CID
	StartEpoch int32
	EndEpoch   int32
}

func (d *DealTracker) runOnce(ctx context.Context) error {
	var wallets []model.Wallet
	err := d.db.WithContext(ctx).Find(&wallets).Error
	if err != nil {
		return errors.Wrap(err, "failed to get wallets from database")
	}

	walletIDs := make(map[string]struct{})
	for _, wallet := range wallets {
		logger.Infof("tracking deals for wallet %s", wallet.ID)
		walletIDs[wallet.ID] = struct{}{}
	}

	rows, err := d.db.WithContext(ctx).Model(&model.Deal{}).
		Select("id", "deal_id", "state", "client_id", "provider", "piece_cid",
			"start_epoch", "end_epoch").Rows()
	if err != nil {
		return errors.Wrap(err, "failed to get deals from database")
	}

	knownDeals := make(map[uint64]KnownDeal)
	unknownDeals := make(map[string][]UnknownDeal)
	for rows.Next() {
		var deal model.Deal
		err = rows.Scan(&deal.ID, &deal.DealID, &deal.State, &deal.ClientID, &deal.Provider, &deal.PieceCID, &deal.StartEpoch, &deal.EndEpoch)
		if err != nil {
			return errors.Wrap(err, "failed to scan row")
		}
		if deal.DealID != nil {
			knownDeals[*deal.DealID] = KnownDeal{
				State: deal.State,
			}
		} else {
			key := deal.Key()
			unknownDeals[key] = append(unknownDeals[key], UnknownDeal{
				ID:         deal.ID,
				ClientID:   deal.ClientID,
				Provider:   deal.Provider,
				PieceCID:   deal.PieceCID,
				StartEpoch: deal.StartEpoch,
				EndEpoch:   deal.EndEpoch,
			})
		}
	}

	var updated int64
	var inserted int64
	defer func() {
		logger.Infof("updated %d deals and inserted %d deals", updated, inserted)
	}()
	err = d.trackDeal(ctx, func(dealID uint64, deal Deal) error {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		_, ok := walletIDs[deal.Proposal.Client]
		if !ok {
			return nil
		}
		newState := deal.GetState()
		current, ok := knownDeals[dealID]
		if ok {
			if current.State == newState {
				return nil
			}

			logger.Infow("Deal state changed", "dealID", dealID, "oldState", current.State, "newState", newState)
			err = database.DoRetry(func() error {
				return d.db.WithContext(ctx).Model(&model.Deal{}).Where("deal_id = ?", dealID).Updates(
					map[string]any{
						"state":              newState,
						"sector_start_epoch": deal.State.SectorStartEpoch,
					}).Error
			})
			if err != nil {
				return errors.Wrap(err, "failed to update deal")
			}
			updated++
			return nil
		}
		dealKey := deal.Key()
		found, ok := unknownDeals[dealKey]
		if ok {
			f := found[0]
			logger.Infow("Deal matched on-chain", "dealID", dealID, "state", newState)
			err = database.DoRetry(func() error {
				return d.db.WithContext(ctx).Model(&model.Deal{}).Where("id = ?", f.ID).Updates(map[string]any{
					"deal_id":            dealID,
					"state":              newState,
					"sector_start_epoch": deal.State.SectorStartEpoch,
				}).Error
			})
			if err != nil {
				return errors.Wrap(err, "failed to update deal")
			}
			updated++
			unknownDeals[dealKey] = unknownDeals[dealKey][1:]
			if len(unknownDeals[dealKey]) == 0 {
				delete(unknownDeals, dealKey)
			}
			return nil
		}
		logger.Infow("Deal external inserted from on-chain", "dealID", dealID, "state", newState)
		root, err := cid.Parse(deal.Proposal.PieceCID.Root)
		if err != nil {
			return errors.Wrap(err, "failed to parse piece CID")
		}
		err = database.DoRetry(func() error {
			return d.db.WithContext(ctx).Create(&model.Deal{
				DealID:           &dealID,
				State:            newState,
				ClientID:         deal.Proposal.Client,
				Provider:         deal.Proposal.Provider,
				Label:            deal.Proposal.Label,
				PieceCID:         model.CID(root),
				PieceSize:        deal.Proposal.PieceSize,
				StartEpoch:       deal.Proposal.StartEpoch,
				EndEpoch:         deal.Proposal.EndEpoch,
				SectorStartEpoch: deal.State.SectorStartEpoch,
				Price:            deal.Proposal.StoragePricePerEpoch,
				Verified:         deal.Proposal.VerifiedDeal,
			}).Error
		})
		if err != nil {
			return errors.Wrap(err, "failed to insert deal")
		}
		inserted++
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to track deal")
	}

	// Mark all expired active deals as expired
	result := d.db.WithContext(ctx).Model(&model.Deal{}).
		Where("end_epoch < ? AND state = 'active'", epochutil.UnixToEpoch(time.Now().Add(-24*time.Hour).Unix())).
		Update("state", model.DealExpired)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to update deals to expired")
	}
	logger.Infof("marked %d deals as expired", result.RowsAffected)

	// Mark all expired deal proposals
	result = d.db.WithContext(ctx).Model(&model.Deal{}).
		Where("state in ('proposed', 'published') AND start_epoch < ?", epochutil.UnixToEpoch(time.Now().Add(-24*time.Hour).Unix())).
		Update("state", model.DealProposalExpired)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to update deal proposals to expired")
	}
	logger.Infof("marked %d deal as proposal_expired", result.RowsAffected)

	return nil
}

func (d *DealTracker) trackDeal(ctx context.Context, callback func(dealID uint64, deal Deal) error) error {
	kvstream, counter, closer, err := d.dealStateStream(ctx)
	if err != nil {
		return err
	}
	defer closer.Close()
	countingCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		for {
			select {
			case <-countingCtx.Done():
				return
			case <-time.After(15 * time.Second):
				downloaded := humanize.Bytes(uint64(counter.N()))
				speed := humanize.Bytes(uint64(counter.Speed()))
				logger.Infof("Downloaded %s with average speed %s / s", downloaded, speed)
			}
		}
	}()
	for stream := range kvstream {
		keyValuePair, ok := stream.Value.(jstream.KV)

		if !ok {
			return errors.New("failed to get key value pair")
		}

		var deal Deal
		err = mapstructure.Decode(keyValuePair.Value, &deal)
		if err != nil {
			return errors.Wrap(err, "failed to decode deal")
		}

		dealID, err := strconv.ParseUint(keyValuePair.Key, 10, 64)
		if err != nil {
			return errors.Wrap(err, "failed to convert deal id to int")
		}

		err = callback(dealID, deal)
		if err != nil {
			return errors.Wrap(err, "failed to callback")
		}
	}

	return ctx.Err()
}
