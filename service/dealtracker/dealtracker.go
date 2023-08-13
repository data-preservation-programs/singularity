package dealtracker

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bcicen/jstream"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/service/epochutil"
	"github.com/data-preservation-programs/singularity/service/healthcheck"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/dustin/go-humanize"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	"github.com/klauspost/compress/zstd"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var ErrAlreadyRunning = errors.New("another worker already running")

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

var Logger = log.Logger("dealtracker")

type DealTracker struct {
	workerID   uuid.UUID
	db         *gorm.DB
	interval   time.Duration
	dealZstURL string
	lotusURL   string
	lotusToken string
	once       bool
}

func NewDealTracker(
	db *gorm.DB,
	interval time.Duration,
	dealZstURL string,
	lotusURL string,
	lotusToken string,
	once bool) DealTracker {
	return DealTracker{
		workerID:   uuid.New(),
		db:         db,
		interval:   interval,
		dealZstURL: dealZstURL,
		lotusURL:   lotusURL,
		lotusToken: lotusToken,
		once:       once,
	}
}

// ThreadSafeReadCloser is a thread-safe implementation of the io.ReadCloser interface.
//
// The ThreadSafeReadCloser struct has the following fields:
// - reader: The underlying io.Reader.
// - closer: The function to close the reader.
// - closed: A boolean indicating whether the reader is closed.
// - mu: A mutex used to synchronize access to the closed field.
//
// The ThreadSafeReadCloser struct implements the io.ReadCloser interface and provides the following methods:
// - Read: Reads data from the underlying reader. It acquires a lock on the mutex to ensure thread safety.
// - Close: Closes the reader. It acquires a lock on the mutex to ensure thread safety and sets the closed field to true before calling the closer function.
type ThreadSafeReadCloser struct {
	reader io.Reader
	closer func()
	closed bool
	mu     sync.Mutex
}

func (t *ThreadSafeReadCloser) Read(p []byte) (n int, err error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.closed {
		return 0, errors.New("closed")
	}
	return t.reader.Read(p)
}

func (t *ThreadSafeReadCloser) Close() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.closed = true
	t.closer()
}

// DealStateStreamFromHTTPRequest retrieves the deal state from an HTTP request and returns a stream of jstream.MetaValue,
// along with a Counter, io.Closer, and any error encountered.
//
// The function takes the following parameters:
// - request: The HTTP request to retrieve the deal state.
// - depth: The depth of the JSON decoding.
// - decompress: A boolean flag indicating whether to decompress the response body.
//
// The function performs the following steps:
// 1. Sends an HTTP request using http.DefaultClient.Do.
// 2. If an error occurs during the request, it returns nil for the channel, Counter, io.Closer, and the error wrapped with an appropriate message.
// 3. If the response status code is not http.StatusOK, it closes the response body and returns nil for the channel, Counter, io.Closer, and an error indicating the failure.
// 4. Creates a countingReader using NewCountingReader to count the number of bytes read from the response body.
// 5. If decompress is true, creates a zstd decompressor using zstd.NewReader and wraps it in a ThreadSafeReadCloser.
//   - If an error occurs during decompression, it closes the response body and returns nil for the channel, Counter, io.Closer, and the error wrapped with an appropriate message.
//   - Creates a jstream.Decoder using jstream.NewDecoder with the decompressor and specified depth, and sets it to emit key-value pairs.
//   - Creates a CloserFunc that closes the decompressor and response body.
//
// 6. If decompress is false, creates a jstream.Decoder using jstream.NewDecoder with the countingReader and specified depth, and sets it to emit key-value pairs.
//   - Sets the response body as the closer.
//
// 7. Returns the jstream.MetaValue stream from jsonDecoder.Stream(), the countingReader, closer, and nil for the error.
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
		safeDecompressor := &ThreadSafeReadCloser{
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
		Logger.Infof("getting deal state from %s", d.dealZstURL)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, d.dealZstURL, nil)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "failed to create request to get deal state zst file")
		}
		return DealStateStreamFromHTTPRequest(req, 1, true)
	}

	Logger.Infof("getting deal state from %s", d.lotusURL)
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

func (*DealTracker) Name() string {
	return "DealTracker"
}

// Start starts the DealTracker and returns a list of service.Done channels, a service.Fail channel, and an error.
//
// The Start method takes a context.Context as input and performs the following steps:
// 1. Defines a getState function that returns a healthcheck.State with WorkType set to model.DealTracking.
// 2. Registers the worker using healthcheck.Register with the provided context, db, workerID, getState function, and false for the force flag.
//   - If an error occurs during registration, it returns nil for the service.Done channels, nil for the service.Fail channel, and the error wrapped with an appropriate message.
//   - If another worker is already running, it logs a warning and checks if d.once is true. If d.once is true, it returns nil for the service.Done channels,
//     nil for the service.Fail channel, and an error indicating that another worker is already running.
//
// 3. Logs a warning message and waits for 1 minute before retrying.
//   - If the context is done during the wait, it returns nil for the service.Done channels, nil for the service.Fail channel, and the context error.
//
// 4. Starts reporting health using healthcheck.StartReportHealth with the provided context, db, workerID, and getState function in a separate goroutine.
// 5. Runs the main loop in a separate goroutine.
//   - Calls d.runOnce to execute the main logic of the DealTracker.
//   - If an error occurs during execution, it logs an error message.
//   - If d.once is true, it returns from the goroutine.
//   - Waits for the specified interval before running the next iteration.
//   - If the context is done during the wait, it returns from the goroutine.
//
// 6. Cleans up resources when the context is done.
//   - Calls d.cleanup to perform cleanup operations.
//   - If an error occurs during cleanup, it logs an error message.
//
// 7. Returns a list of service.Done channels containing healthcheckDone, runDone, and cleanupDone, the service.Fail channel fail, and nil for the error.
func (d *DealTracker) Start(ctx context.Context) ([]service.Done, service.Fail, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
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
			cancel()
			return nil, nil, errors.Wrap(err, "failed to register worker")
		}
		if alreadyRunning {
			Logger.Warnw("another worker already running")
			if d.once {
				cancel()
				return nil, nil, ErrAlreadyRunning
			}
		}
		Logger.Warn("retrying in 1 minute")
		select {
		case <-ctx.Done():
			cancel()
			return nil, nil, ctx.Err()
		case <-time.After(time.Minute):
		}
	}

	healthcheckDone := make(chan struct{})
	go func() {
		defer close(healthcheckDone)
		healthcheck.StartReportHealth(ctx, d.db, d.workerID, getState)
		Logger.Info("health report stopped")
	}()

	runDone := make(chan struct{})
	fail := make(chan error)
	go func() {
		defer cancel()
		defer close(runDone)
		for {
			err := d.runOnce(ctx)
			if err != nil {
				Logger.Errorw("failed to run once", "error", err)
			}
			if d.once {
				cancel()
				Logger.Info("run once done")
				return
			}
			select {
			case <-ctx.Done():
				Logger.Info("run stopped")
				return
			case <-time.After(d.interval):
			}
		}
	}()

	cleanupDone := make(chan struct{})
	go func() {
		defer close(cleanupDone)
		<-ctx.Done()
		err := d.cleanup()
		if err != nil {
			Logger.Errorw("failed to cleanup", "error", err)
		} else {
			Logger.Info("cleanup done")
		}
	}()

	return []service.Done{healthcheckDone, runDone, cleanupDone}, fail, nil
}

func (d *DealTracker) cleanup() error {
	return database.DoRetry(func() error {
		return d.db.Where("id = ?", d.workerID).Delete(&model.Worker{}).Error
	})
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
	delay := time.Hour
	var err error
	if d.dealZstURL == "" {
		lotusTime, err := util.GetLotusHeadTime(ctx, d.lotusURL, d.lotusToken)
		if err != nil {
			return errors.Wrap(err, "failed to get lotus head time")
		}
		delay = time.Since(lotusTime)
	}

	db := d.db.WithContext(ctx)
	var wallets []model.Wallet
	err = db.Find(&wallets).Error
	if err != nil {
		return errors.Wrap(err, "failed to get wallets from database")
	}

	walletIDs := make(map[string]struct{})
	for _, wallet := range wallets {
		Logger.Infof("tracking deals for wallet %s", wallet.ID)
		walletIDs[wallet.ID] = struct{}{}
	}

	knownDeals := make(map[uint64]model.DealState)
	rows, err := db.Model(&model.Deal{}).Where("deal_id IS NOT NULL").
		Select("deal_id", "state").Rows()
	if err != nil {
		return errors.Wrap(err, "failed to get known deals from database")
	}
	for rows.Next() {
		var dealID uint64
		var state model.DealState
		err = rows.Scan(&dealID, &state)
		if err != nil {
			return errors.Wrap(err, "failed to scan row")
		}
		knownDeals[dealID] = state
	}

	unknownDeals := make(map[string][]UnknownDeal)
	rows, err = db.Model(&model.Deal{}).Where("deal_id IS NULL AND state NOT IN ?", []model.DealState{model.DealExpired, model.DealProposalExpired}).
		Select("id", "deal_id", "state", "client_id", "provider", "piece_cid",
			"start_epoch", "end_epoch").Rows()
	if err != nil {
		return errors.Wrap(err, "failed to get unknown deals from database")
	}
	for rows.Next() {
		var deal model.Deal
		err = rows.Scan(&deal.ID, &deal.DealID, &deal.State, &deal.ClientID, &deal.Provider, &deal.PieceCID, &deal.StartEpoch, &deal.EndEpoch)
		if err != nil {
			return errors.Wrap(err, "failed to scan row")
		}
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

	var updated int64
	var inserted int64
	defer func() {
		Logger.Infof("updated %d deals and inserted %d deals", updated, inserted)
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
			if current == newState {
				return nil
			}

			Logger.Infow("Deal state changed", "dealID", dealID, "oldState", current, "newState", newState)
			err = database.DoRetry(func() error {
				return db.Model(&model.Deal{}).Where("deal_id = ?", dealID).Updates(
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
			Logger.Infow("Deal matched on-chain", "dealID", dealID, "state", newState)
			err = database.DoRetry(func() error {
				return db.Model(&model.Deal{}).Where("id = ?", f.ID).Updates(map[string]any{
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
		Logger.Infow("Deal external inserted from on-chain", "dealID", dealID, "state", newState)
		root, err := cid.Parse(deal.Proposal.PieceCID.Root)
		if err != nil {
			return errors.Wrap(err, "failed to parse piece CID")
		}
		err = database.DoRetry(func() error {
			return db.Create(&model.Deal{
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
	result := db.Model(&model.Deal{}).
		Where("end_epoch < ? AND state = 'active'", epochutil.UnixToEpoch(time.Now().Add(-delay).Unix())).
		Update("state", model.DealExpired)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to update deals to expired")
	}
	Logger.Infof("marked %d deals as expired", result.RowsAffected)

	// Mark all expired deal proposals
	result = db.Model(&model.Deal{}).
		Where("state in ('proposed', 'published') AND start_epoch < ?", epochutil.UnixToEpoch(time.Now().Add(-delay).Unix())).
		Update("state", model.DealProposalExpired)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to update deal proposals to expired")
	}
	Logger.Infof("marked %d deal as proposal_expired", result.RowsAffected)

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
				Logger.Infof("Downloaded %s with average speed %s / s", downloaded, speed)
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
