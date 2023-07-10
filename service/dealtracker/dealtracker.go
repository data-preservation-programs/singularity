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
	"syscall"
	"time"

	"github.com/data-preservation-programs/singularity/service/healthcheck"
	"github.com/google/uuid"
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
		if model.EpochToTime(d.Proposal.StartEpoch).Before(time.Now()) {
			return model.DealProposalExpired
		}
		return model.DealPublished
	}
	if model.EpochToTime(d.Proposal.EndEpoch).Before(time.Now()) {
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

func DealStateStreamFromHTTPRequest(request *http.Request, depth int, decompress bool) (chan *jstream.MetaValue, io.Closer, error) {
	//nolint: bodyclose
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get deal state from lotus API")
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, nil, errors.New("failed to get deal state: " + resp.Status)
	}
	var jsonDecoder *jstream.Decoder
	var closer io.Closer
	if decompress {
		decompressor, err := zstd.NewReader(resp.Body)
		if err != nil {
			resp.Body.Close()
			return nil, nil, errors.Wrap(err, "failed to create zstd decompressor")
		}
		jsonDecoder = jstream.NewDecoder(decompressor, depth).EmitKV()
		closer = CloserFunc(func() error {
			decompressor.Close()
			return resp.Body.Close()
		})
	} else {
		jsonDecoder = jstream.NewDecoder(resp.Body, depth).EmitKV()
		closer = resp.Body
	}

	return jsonDecoder.Stream(), closer, nil
}

func (d *DealTracker) dealStateStream(ctx context.Context) (chan *jstream.MetaValue, io.Closer, error) {
	if d.dealZstURL != "" {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, d.dealZstURL, nil)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to create request to get deal state zst file")
		}
		return DealStateStreamFromHTTPRequest(req, 1, true)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, d.lotusURL, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create request to get deal state from lotus API")
	}
	if d.lotusToken != "" {
		req.Header.Set("Authorization", "Bearer "+d.lotusToken)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(strings.NewReader(`{"jsonrpc":"2.0","method":"Filecoin.StateMarketDeals","params":[null],"id":0}`))
	return DealStateStreamFromHTTPRequest(req, 2, false)
}

var staleThreshold = time.Minute * 5

func (d *DealTracker) Run(ctx context.Context) error {
	var activeWorkerCount int64
	err := d.db.WithContext(ctx).Model(&model.Worker{}).Where("work_type = ? AND last_heartbeat > ?", model.DealTracking, time.Now().UTC().Add(-staleThreshold)).
		Count(&activeWorkerCount).Error
	if err != nil {
		return errors.Wrap(err, "failed to count active workers")
	}
	if activeWorkerCount > 0 {
		return errors.New("deal tracker already running")
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGTRAP)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	getState := func() healthcheck.State {
		return healthcheck.State{
			WorkType: model.DealTracking,
		}
	}

	healthcheck.StartHealthCheck(ctx, d.db, d.workerID, getState)
	go healthcheck.StartHealthCheck(ctx, d.db, d.workerID, getState)

	go func() {
		for {
			err := d.runOnce(ctx)
			if err != nil {
				logger.Errorw("failed to run deal maker", "error", err)
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

func (d *DealTracker) runOnce(ctx context.Context) error {
	should, err := d.shouldTrackDeal(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to check if should track deal")
	}

	if !should {
		return nil
	}

	var wallets []model.Wallet
	err = d.db.WithContext(ctx).Find(&wallets).Error
	if err != nil {
		return errors.Wrap(err, "failed to get wallets from database")
	}

	// Clean up the deals table before querying
	// We want to mark all deals that has expired as expired
	err = d.db.WithContext(ctx).Model(&model.Deal{}).
		Where("end_epoch < ?", model.UnixToEpoch(time.Now().Unix())).
		Update("state", model.DealExpired).Error
	if err != nil {
		return errors.Wrap(err, "failed to update deals to expired")
	}

	var deals []model.Deal
	err = d.db.WithContext(ctx).Find(&deals).Error
	if err != nil {
		return errors.Wrap(err, "failed to get deals from database")
	}

	dealsByDealID := make(map[uint64]model.Deal)
	unPublishedDeals := make(map[string][]model.Deal)
	for _, deal := range deals {
		if deal.DealID != nil {
			dealsByDealID[*deal.DealID] = deal
		} else {
			unPublishedDeals[deal.Key()] = append(unPublishedDeals[deal.Key()], deal)
		}
	}

	err = d.trackDeal(ctx, func(dealID uint64, deal Deal) error {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		current, ok := dealsByDealID[dealID]
		newState := deal.GetState()
		if ok && current.State != newState {
			logger.Debugw("Deal state changed", "dealID", dealID, "oldState", current.State, "newState", newState)
			err = database.DoRetry(func() error {
				return d.db.WithContext(ctx).Model(&model.Deal{}).Where("id = ?", current.ID).Update("state", newState).Error
			})
			if err != nil {
				return errors.Wrap(err, "failed to update deal")
			}
			return nil
		}
		if ok {
			return nil
		}
		dealKey := deal.Key()
		found, ok := unPublishedDeals[dealKey]
		if ok {
			f := found[0]
			logger.Debugw("Deal matched on-chain", "dealID", dealID, "state", newState)
			err = database.DoRetry(func() error {
				return d.db.WithContext(ctx).Model(&model.Deal{}).Where("id = ?", f.ID).Updates(map[string]interface{}{
					"deal_id":            dealID,
					"state":              newState,
					"sector_start_epoch": deal.State.SectorStartEpoch,
				}).Error
			})
			if err != nil {
				return errors.Wrap(err, "failed to update deal")
			}
			unPublishedDeals[dealKey] = unPublishedDeals[dealKey][1:]
			if len(unPublishedDeals[dealKey]) == 0 {
				delete(unPublishedDeals, dealKey)
			}
			return nil
		}
		logger.Debugw("Deal inserted from on-chain", "dealID", dealID, "state", newState)
		err = database.DoRetry(func() error {
			return d.db.WithContext(ctx).Create(&model.Deal{
				DealID:           &dealID,
				State:            newState,
				ClientID:         deal.Proposal.Client,
				Provider:         deal.Proposal.Provider,
				Label:            deal.Proposal.Label,
				PieceCID:         deal.Proposal.PieceCID.Root,
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
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to track deal")
	}
	return nil
}

func (d *DealTracker) trackDeal(ctx context.Context, callback func(dealID uint64, deal Deal) error) error {
	kvstream, closer, err := d.dealStateStream(ctx)
	if err != nil {
		return err
	}
	defer closer.Close()
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
	return nil
}

func (d *DealTracker) shouldTrackDeal(ctx context.Context) (bool, error) {
	now := time.Now()
	last := model.Global{
		Key:   "dealTrackingLastTimestampNano",
		Value: fmt.Sprintf("%d", now.UnixNano()),
	}
	shouldContinue := false
	err := database.DoRetry(func() error {
		return d.db.WithContext(ctx).Transaction(func(db *gorm.DB) error {
			var global model.Global
			err := db.Where("key = ?", "dealTrackingLastTimestampNano").First(&global).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				shouldContinue = true
				return db.Create(&last).Error
			}
			if err != nil {
				return err
			}
			ts, err := strconv.ParseInt(global.Value, 10, 64)
			if err != nil {
				return err
			}

			if now.UnixNano()-ts > int64(d.interval) {
				shouldContinue = true
				return db.Model(&model.Global{}).Where("key = ?", "dealTrackingLastTimestampNano").Update("value", last.Value).Error
			}
			return nil
		})
	})
	if err != nil {
		return false, errors.Wrap(err, "failed to get last dealtracking timestamp from database")
	}
	return shouldContinue, nil
}
