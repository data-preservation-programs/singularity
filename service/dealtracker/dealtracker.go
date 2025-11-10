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
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/epochutil"
	"github.com/data-preservation-programs/singularity/service/healthcheck"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/dustin/go-humanize"
	"github.com/google/uuid"
	"github.com/gotidy/ptr"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	"github.com/klauspost/compress/zstd"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

var ErrAlreadyRunning = errors.New("another worker already running")

const (
	healthRegisterRetryInterval = time.Minute
	cleanupTimeout              = 5 * time.Second
	logStatsInterval            = 15 * time.Second
)

type Deal struct {
	Proposal DealProposal
	State    DealState
}

func (d Deal) Key() string {
	return fmt.Sprintf("%s-%s-%s-%d-%d", d.Proposal.Client, d.Proposal.Provider,
		d.Proposal.PieceCID.Root, d.Proposal.StartEpoch, d.Proposal.EndEpoch)
}

func (d Deal) GetState(headTime time.Time) model.DealState {
	if d.State.SlashEpoch > 0 {
		return model.DealSlashed
	}
	if d.State.SectorStartEpoch < 0 {
		if epochutil.EpochToTime(d.Proposal.StartEpoch).Before(headTime) {
			return model.DealProposalExpired
		}
		return model.DealPublished
	}
	if epochutil.EpochToTime(d.Proposal.EndEpoch).Before(headTime) {
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
	workerID    uuid.UUID
	dbNoContext *gorm.DB
	interval    time.Duration
	dealZstURL  string
	lotusURL    string
	lotusToken  string
	once        bool
}

func NewDealTracker(
	db *gorm.DB,
	interval time.Duration,
	dealZstURL string,
	lotusURL string,
	lotusToken string,
	once bool,
) DealTracker {
	return DealTracker{
		workerID:    uuid.New(),
		dbNoContext: db,
		interval:    interval,
		dealZstURL:  dealZstURL,
		lotusURL:    lotusURL,
		lotusToken:  lotusToken,
		once:        once,
	}
}

// ThreadSafeReadCloser is a thread-safe implementation of the io.ReadCloser interface.
//
// The ThreadSafeReadCloser struct has the following fields:
//   - reader: The underlying io.Reader.
//   - closer: The function to close the reader.
//   - closed: A boolean indicating whether the reader is closed.
//   - mu: A mutex used to synchronize access to the closed field.
//
// The ThreadSafeReadCloser struct implements the io.ReadCloser interface and provides the following methods:
//   - Read: Reads data from the underlying reader. It acquires a lock on the mutex to ensure thread safety.
//   - Close: Closes the reader. It acquires a lock on the mutex to ensure thread safety and sets the closed field to true before calling the closer function.
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
//   - request: The HTTP request to retrieve the deal state.
//   - depth: The depth of the JSON decoding.
//   - decompress: A boolean flag indicating whether to decompress the response body.
//
// The function performs the following steps:
//
//  1. Sends an HTTP request using http.DefaultClient.Do.
//
//  2. If an error occurs during the request, it returns nil for the channel, Counter, io.Closer, and the error wrapped with an appropriate message.
//
//  3. If the response status code is not http.StatusOK, it closes the response body and returns nil for the channel, Counter, io.Closer, and an error indicating the failure.
//
//  4. Creates a countingReader using NewCountingReader to count the number of bytes read from the response body.
//
//  5. If decompress is true, creates a zstd decompressor using zstd.NewReader and wraps it in a ThreadSafeReadCloser.
//     - If an error occurs during decompression, it closes the response body and returns nil for the channel, Counter, io.Closer, and the error wrapped with an appropriate message.
//     - Creates a jstream.Decoder using jstream.NewDecoder with the decompressor and specified depth, and sets it to emit key-value pairs.
//     - Creates a CloserFunc that closes the decompressor and response body.
//
//  6. If decompress is false, creates a jstream.Decoder using jstream.NewDecoder with the countingReader and specified depth, and sets it to emit key-value pairs.
//     - Sets the response body as the closer.
//
//  7. Returns the jstream.MetaValue stream from jsonDecoder.Stream(), the countingReader, closer, and nil for the error.
func DealStateStreamFromHTTPRequest(request *http.Request, depth int, decompress bool) (chan *jstream.MetaValue, Counter, io.Closer, error) {
	//nolint: bodyclose
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, nil, errors.WithStack(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, nil, nil, errors.Newf("failed to get deal state: %s", resp.Status)
	}
	var jsonDecoder *jstream.Decoder
	var closer io.Closer
	countingReader := NewCountingReader(resp.Body)
	if decompress {
		decompressor, err := zstd.NewReader(countingReader)
		if err != nil {
			resp.Body.Close()
			return nil, nil, nil, errors.WithStack(err)
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
			return nil, nil, nil, errors.Wrapf(err, "failed to create request to get deal state zst file %s", d.dealZstURL)
		}
		return DealStateStreamFromHTTPRequest(req, 1, true)
	}

	Logger.Infof("getting deal state from %s", d.lotusURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, d.lotusURL, nil)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "failed to create request to get deal state from lotus API %s", d.lotusURL)
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
//
//  1. Defines a getState function that returns a healthcheck.State with JobType set to model.DealTracking.
//
//  2. Registers the worker using healthcheck.Register with the provided context, dbNoContext, workerID, getState function, and false for the force flag.
//     - If an error occurs during registration, it returns nil for the service.Done channels, nil for the service.Fail channel, and the error wrapped with an appropriate message.
//     - If another worker is already running, it logs a warning and checks if d.once is true. If d.once is true, it returns nil for the service.Done channels,
//     nil for the service.Fail channel, and an error indicating that another worker is already running.
//
//  3. Logs a warning message and waits for 1 minute before retrying.
//     - If the context is done during the wait, it returns nil for the service.Done channels, nil for the service.Fail channel, and the context error.
//
//  4. Starts reporting health using healthcheck.StartReportHealth with the provided context, dbNoContext, workerID, and getState function in a separate goroutine.
//
//  5. Runs the main loop in a separate goroutine.
//     - Calls d.runOnce to execute the main logic of the DealTracker.
//     - If an error occurs during execution, it logs an error message.
//     - If d.once is true, it returns from the goroutine.
//     - Waits for the specified interval before running the next iteration.
//     - If the context is done during the wait, it returns from the goroutine.
//
//  6. Cleans up resources when the context is done.
//     - Calls d.cleanup to perform cleanup operations.
//     - If an error occurs during cleanup, it logs an error message.
//
//  7. Returns a list of service.Done channels containing healthcheckDone, runDone, and cleanupDone, the service.Fail channel fail, and nil for the error.
func (d *DealTracker) Start(ctx context.Context, exitErr chan<- error) error {
	var regTimer *time.Timer
	for {
		alreadyRunning, err := healthcheck.Register(ctx, d.dbNoContext, d.workerID, model.DealTracker, false)
		if err != nil {
			return errors.WithStack(err)
		}
		if !alreadyRunning {
			break
		}

		Logger.Warnw("another worker already running")
		if d.once {
			return ErrAlreadyRunning
		}
		Logger.Warn("retrying in 1 minute")
		if regTimer == nil {
			regTimer = time.NewTimer(healthRegisterRetryInterval)
			defer regTimer.Stop()
		} else {
			regTimer.Reset(healthRegisterRetryInterval)
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-regTimer.C:
		}
	}

	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)

	healthcheckDone := make(chan struct{})
	go func() {
		defer close(healthcheckDone)
		healthcheck.StartReportHealth(ctx, d.dbNoContext, d.workerID, model.DealTracker)
		Logger.Info("health report stopped")
	}()

	go func() {
		var timer *time.Timer
		var runErr error
		for {
			runErr = d.runOnce(ctx)
			if runErr != nil {
				if ctx.Err() != nil {
					if errors.Is(runErr, context.Canceled) {
						runErr = nil
					}
					Logger.Info("run stopped")
					break
				}
				Logger.Errorw("failed to run once", "error", runErr)
			}
			if d.once {
				Logger.Info("run once done")
				break
			}
			if timer == nil {
				timer = time.NewTimer(d.interval)
				defer timer.Stop()
			} else {
				timer.Reset(d.interval)
			}

			var stopped bool
			select {
			case <-ctx.Done():
				stopped = true
			case <-timer.C:
			}
			if stopped {
				Logger.Info("run stopped")
				break
			}
		}

		cancel()

		ctx2, cancel2 := context.WithTimeout(context.Background(), cleanupTimeout)
		defer cancel2()
		//nolint:contextcheck
		err := d.cleanup(ctx2)
		if err != nil {
			Logger.Errorw("failed to cleanup", "error", err)
		} else {
			Logger.Info("cleanup done")
		}

		<-healthcheckDone

		if exitErr != nil {
			exitErr <- runErr
		}
	}()

	return nil
}

func (d *DealTracker) cleanup(ctx context.Context) error {
	return database.DoRetry(ctx, func() error {
		return d.dbNoContext.WithContext(ctx).Where("id = ?", d.workerID).Delete(&model.Worker{}).Error
	})
}

type KnownDeal struct {
	State model.DealState
}
type UnknownDeal struct {
	ID            model.DealID
	ClientID string
	Provider      string
	PieceCID      model.CID
	StartEpoch int32
	EndEpoch   int32
}

// runOnce is a method of the DealTracker type. It is responsible for performing a single cycle
// of deal tracking. It queries the local database for known deals and wallets, compares the
// local data with on-chain data, updates the local data to reflect any changes, inserts new deals
// found on-chain but not in the local data, and marks expired deals and deal proposals as such.
//
// The steps it takes are as follows:
//  1. Calculate the delay time based on Lotus head time if dealZstURL is empty, or default to 1 hour.
//  2. Retrieve the wallets from the local database.
//  3. Create a set of wallet IDs for lookup purposes.
//  4. Retrieve the known deals from the local database.
//  5. Retrieve the unknown deals from the local database.
//  6. Invoke trackDeal function to compare and update the local deals with on-chain data.
//  7. In trackDeal's callback, update existing deals if the state has changed.
//  8. In trackDeal's callback, match unknown deals in the local database to known deals on-chain.
//  9. In trackDeal's callback, insert new deals found on-chain that don't exist in the local database.
//  10. Mark all expired active deals as 'expired' in the local database.
//  11. Mark all expired deal proposals as 'proposal_expired' in the local database.
//
// Parameters:
//
//   - ctx: The context to control the lifecycle of the run. If the context is done,
//     the function exits cleanly.
//
// Returns:
//
//   - error: An error that represents the failure of the operation, or nil if the operation was successful.
func (d *DealTracker) runOnce(ctx context.Context) error {
	headTime, err := util.GetLotusHeadTime(ctx, d.lotusURL, d.lotusToken)
	if err != nil {
		return errors.Wrapf(err, "failed to get lotus head time from %s", d.lotusURL)
	}

	var lastEpoch int32

	db := d.dbNoContext.WithContext(ctx)
	var actors []model.Actor
	err = db.Find(&actors).Error
	if err != nil {
		return errors.Wrap(err, "failed to get actors from database")
	}

	actorIDs := make(map[string]struct{})
	for _, actor := range actors {
		Logger.Infof("tracking deals for actor %s", actor.ID)
		actorIDs[actor.ID] = struct{}{}
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
		return errors.WithStack(err)
	}
	for rows.Next() {
		var deal model.Deal
		err = rows.Scan(&deal.ID, &deal.DealID, &deal.State, &deal.ClientID, &deal.Provider, &deal.PieceCID, &deal.StartEpoch, &deal.EndEpoch)
		if err != nil {
			return errors.WithStack(err)
		}
		key := deal.Key()
		unknownDeals[key] = append(unknownDeals[key], UnknownDeal{
			ID:            deal.ID,
			ClientID: deal.ClientID,
			Provider:      deal.Provider,
			PieceCID:      deal.PieceCID,
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
		if deal.State.LastUpdatedEpoch > lastEpoch {
			lastEpoch = deal.State.LastUpdatedEpoch
		}
		_, ok := actorIDs[deal.Proposal.Client]
		if !ok {
			return nil
		}
		newState := deal.GetState(headTime)
		var lastVerifiedAt *time.Time
		if newState == model.DealActive {
			lastVerifiedAt = ptr.Of(headTime)
		}
		current, ok := knownDeals[dealID]
		if ok {
			if current == newState {
				return nil
			}

			if newState == model.DealExpired || newState == model.DealProposalExpired {
				return nil
			}
			Logger.Infow("Deal state changed", "dealID", dealID, "oldState", current, "newState", newState)
			err = database.DoRetry(ctx, func() error {
				return db.Model(&model.Deal{}).Where("deal_id = ?", dealID).Updates(
					map[string]any{
						"state":              newState,
						"sector_start_epoch": deal.State.SectorStartEpoch,
						"last_verified_at":   lastVerifiedAt,
					}).Error
			})
			if err != nil {
				return errors.WithStack(err)
			}
			updated++
			return nil
		}
		dealKey := deal.Key()
		found, ok := unknownDeals[dealKey]
		if ok {
			if newState == model.DealExpired || newState == model.DealProposalExpired {
				return nil
			}
			f := found[0]
			Logger.Infow("Deal matched on-chain", "dealID", dealID, "state", newState)
			err = database.DoRetry(ctx, func() error {
				return db.Model(&model.Deal{}).Where("id = ?", f.ID).Updates(map[string]any{
					"deal_id":            dealID,
					"state":              newState,
					"sector_start_epoch": deal.State.SectorStartEpoch,
					"last_verified_at":   lastVerifiedAt,
				}).Error
			})
			if err != nil {
				return errors.WithStack(err)
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
			return errors.Wrapf(err, "failed to parse piece CID %s", deal.Proposal.PieceCID.Root)
		}
		err = database.DoRetry(ctx, func() error {
			return db.Create(&model.Deal{
				DealID:           &dealID,
				State:            newState,
				DealType:         model.DealTypeMarket,
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
				LastVerifiedAt:   lastVerifiedAt,
			}).Error
		})
		if err != nil {
			return errors.WithStack(err)
		}
		inserted++
		return nil
	})
	if err != nil {
		return errors.WithStack(err)
	}

	// Mark all expired active deals as expired
	result := db.Model(&model.Deal{}).
		Where("end_epoch < ? AND state = 'active'", lastEpoch).
		Update("state", model.DealExpired)
	if result.Error != nil {
		return errors.WithStack(err)
	}
	Logger.Infof("marked %d deals as expired", result.RowsAffected)

	// Mark all expired deal proposals
	result = db.Model(&model.Deal{}).
		Where("state in ('proposed', 'published') AND start_epoch < ?", lastEpoch).
		Update("state", model.DealProposalExpired)
	if result.Error != nil {
		return errors.WithStack(err)
	}
	Logger.Infof("marked %d deal as proposal_expired", result.RowsAffected)

	return nil
}

func (d *DealTracker) trackDeal(ctx context.Context, callback func(dealID uint64, deal Deal) error) error {
	kvstream, counter, closer, err := d.dealStateStream(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	defer closer.Close()
	countingCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		timer := time.NewTimer(logStatsInterval)
		defer timer.Stop()
		for {
			select {
			case <-countingCtx.Done():
				return
			case <-timer.C:
				downloaded := humanize.Bytes(uint64(counter.N()))
				speed := humanize.Bytes(uint64(counter.Speed()))
				Logger.Infof("Downloaded %s with average speed %s / s", downloaded, speed)
				timer.Reset(logStatsInterval)
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
			return errors.Wrapf(err, "failed to decode deal %s", keyValuePair.Value)
		}

		dealID, err := strconv.ParseUint(keyValuePair.Key, 10, 64)
		if err != nil {
			return errors.Wrapf(err, "failed to convert deal id %s to int", keyValuePair.Key)
		}

		err = callback(dealID, deal)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return ctx.Err()
}
