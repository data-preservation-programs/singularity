package service

import (
	"context"
	"fmt"
	"github.com/bcicen/jstream"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-log/v2"
	"github.com/klauspost/compress/zstd"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Deal struct {
	Proposal DealProposal
	State    DealState
}

func (d Deal) Key() string {
	return fmt.Sprintf("%s-%s-%s-%d-%d", d.Proposal.Client, d.Proposal.Provider,
		d.Proposal.PieceCID.Root, model.EpochToUnix(d.Proposal.StartEpoch), model.EpochToUnix(d.Proposal.EndEpoch))
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
	PieceSize            uint64
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

type DealTracker struct {
	db         *gorm.DB
	interval   time.Duration
	dealZstURL string
	lotusURL   string
	lotusToken string
	logger     *log.ZapEventLogger
}

func NewDealTracker(db *gorm.DB, interval time.Duration,
	dealZstURL string,
	lotusURL string,
	lotusToken string) DealTracker {
	return DealTracker{
		db:         db,
		interval:   interval,
		dealZstURL: dealZstURL,
		lotusURL:   lotusURL,
		lotusToken: lotusToken,
		logger:     log.Logger("dealtracker"),
	}
}

func (d *DealTracker) dealStateStreamFromHttpRequest(request *http.Request, depth int) (chan *jstream.MetaValue, io.Closer, error) {
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get deal state from lotus API")
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, nil, fmt.Errorf("failed to get deal state zst file: %s", resp.Status)
	}
	decompressor, err := zstd.NewReader(resp.Body)
	if err != nil {
		resp.Body.Close()
		return nil, nil, errors.Wrap(err, "failed to create zstd decompressor")
	}

	jsonDecoder := jstream.NewDecoder(decompressor, depth).EmitKV()

	return jsonDecoder.Stream(), CloserFunc(func() error {
		decompressor.Close()
		return resp.Body.Close()
	}), nil
}

func (d *DealTracker) dealStateStream(ctx context.Context) (chan *jstream.MetaValue, io.Closer, error) {
	if d.dealZstURL != "" {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, d.dealZstURL, nil)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to create request to get deal state zst file")
		}
		return d.dealStateStreamFromHttpRequest(req, 1)
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
	return d.dealStateStreamFromHttpRequest(req, 2)
}

func (d *DealTracker) Run(ctx context.Context) {
	for {
		err := d.runOnce(ctx)
		if err != nil {
			d.logger.Errorw("failed to run deal maker", "error", err)
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(d.interval + time.Minute):
		}
	}
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
	for i, _ := range wallets {
		decrypted, err := model.DecryptFromBase64String(wallets[i].PrivateKey)
		if err != nil {
			return errors.Wrap(err, "failed to decrypt private key")
		}
		wallets[i].PrivateKey = string(decrypted)
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
		current, ok := dealsByDealID[dealID]
		newState := deal.GetState()
		if ok && current.State != newState {
			d.logger.Debugw("Deal state changed", "dealID", dealID, "oldState", current.State, "newState", newState)
			d.db.WithContext(ctx).Where("deal_id = ?", dealID).Update("state", newState)
			return nil
		}
		if ok {
			return nil
		}
		dealKey := deal.Key()
		found, ok := unPublishedDeals[dealKey]
		if ok {
			f := found[0]
			d.logger.Debugw("Deal matched on-chain", "dealID", dealID, "state", newState)
			err = d.db.WithContext(ctx).Where("id = ?", f.ID).Updates(map[string]interface{}{
				"deal_id":      dealID,
				"state":        newState,
				"sector_start": model.EpochToTime(deal.State.SectorStartEpoch),
			}).Error
			if err != nil {
				return errors.Wrap(err, "failed to update deal")
			}
			unPublishedDeals[dealKey] = unPublishedDeals[dealKey][1:]
			if len(unPublishedDeals[dealKey]) == 0 {
				delete(unPublishedDeals, dealKey)
			}
			return nil
		}
		d.logger.Debugw("Deal inserted from on-chain", "dealID", dealID, "state", newState)
		sectorStart := time.Time{}
		if deal.State.SectorStartEpoch > 0 {
			sectorStart = model.EpochToTime(deal.State.SectorStartEpoch)
		}
		err = d.db.WithContext(ctx).Create(&model.Deal{
			DealID:      &dealID,
			State:       newState,
			ClientID:    deal.Proposal.Client,
			Provider:    deal.Proposal.Provider,
			Label:       deal.Proposal.Label,
			PieceCID:    deal.Proposal.PieceCID.Root,
			PieceSize:   deal.Proposal.PieceSize,
			Start:       model.EpochToTime(deal.Proposal.StartEpoch),
			Duration:    model.EpochToTime(deal.Proposal.EndEpoch).Sub(model.EpochToTime(deal.Proposal.StartEpoch)),
			End:         model.EpochToTime(deal.Proposal.EndEpoch),
			SectorStart: sectorStart,
			Price: model.StoragePricePerEpochToPricePerDeal(
				deal.Proposal.StoragePricePerEpoch,
				deal.Proposal.PieceSize,
				deal.Proposal.EndEpoch-deal.Proposal.StartEpoch),
			Verified: deal.Proposal.VerifiedDeal,
		}).Error
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
		Key:   "dealTrackingLastTimestamp",
		Value: fmt.Sprintf("%d", now.Unix()),
	}
	shouldContinue := false
	err := d.db.WithContext(ctx).Transaction(func(db *gorm.DB) error {
		var global model.Global
		err := db.Where("key = ?", "dealTrackingLastTimestamp").First(&global).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			shouldContinue = true
			return db.Create(&last).Error
		}
		if err != nil {
			return err
		}
		ts, err := strconv.Atoi(global.Value)
		if err != nil {
			return err
		}

		unix := time.Unix(int64(ts), 0)
		if now.Sub(unix) > d.interval {
			db.Where("key = ?", "dealTrackingLastTimestamp").Update("value", last.Value)
			shouldContinue = true
		}
		return nil
	})
	if err != nil {
		return false, errors.Wrap(err, "failed to get last dealtracking timestamp from database")
	}
	return shouldContinue, nil
}
