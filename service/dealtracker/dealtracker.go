package dealtracker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	// Add pprof imports

	"net/http/pprof"
	_ "net/http/pprof"

	"bytes"

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
	"github.com/tidwall/gjson"
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
	batchSize   int
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
		batchSize:   100, // Default batch size
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

// DealStateStreamFromHTTPRequest creates a custom streaming parser for efficient deal processing
func DealStateStreamFromHTTPRequest(request *http.Request, depth int, decompress bool, walletIDs map[string]struct{}) (chan *ParsedDeal, Counter, io.Closer, error) {
	//nolint: bodyclose
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, nil, errors.WithStack(err)
	}
	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, nil, nil, errors.Newf("failed to get deal state: %s", resp.Status)
	}

	var reader io.Reader
	var closer io.Closer
	countingReader := NewCountingReader(resp.Body)

	if decompress {
		decompressor, err := zstd.NewReader(countingReader)
		if err != nil {
			_ = resp.Body.Close()
			return nil, nil, nil, errors.WithStack(err)
		}
		safeDecompressor := &ThreadSafeReadCloser{
			reader: decompressor,
			closer: decompressor.Close,
		}
		reader = safeDecompressor
		closer = CloserFunc(func() error {
			safeDecompressor.Close()
			return resp.Body.Close()
		})
	} else {
		reader = countingReader
		closer = resp.Body
	}

	// Create channel for parsed deals
	dealChan := make(chan *ParsedDeal, 100)

	// Create channel for raw deals that need processing
	rawDealChan := make(chan struct {
		DealID  uint64
		RawDeal json.RawMessage
	}, 200)

	// Start worker goroutines for parallel deal processing
	const numWorkers = 4
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			for raw := range rawDealChan {
				// Extract deal fields directly with gjson
				deal := Deal{
					Proposal: DealProposal{
						PieceCID: Cid{
							Root: gjson.GetBytes(raw.RawDeal, "Proposal.PieceCID./").String(),
						},
						PieceSize:            gjson.GetBytes(raw.RawDeal, "Proposal.PieceSize").Int(),
						VerifiedDeal:         gjson.GetBytes(raw.RawDeal, "Proposal.VerifiedDeal").Bool(),
						Client:               gjson.GetBytes(raw.RawDeal, "Proposal.Client").String(),
						Provider:             gjson.GetBytes(raw.RawDeal, "Proposal.Provider").String(),
						Label:                gjson.GetBytes(raw.RawDeal, "Proposal.Label").String(),
						StartEpoch:           int32(gjson.GetBytes(raw.RawDeal, "Proposal.StartEpoch").Int()),
						EndEpoch:             int32(gjson.GetBytes(raw.RawDeal, "Proposal.EndEpoch").Int()),
						StoragePricePerEpoch: gjson.GetBytes(raw.RawDeal, "Proposal.StoragePricePerEpoch").String(),
					},
					State: DealState{
						SectorStartEpoch: int32(gjson.GetBytes(raw.RawDeal, "State.SectorStartEpoch").Int()),
						LastUpdatedEpoch: int32(gjson.GetBytes(raw.RawDeal, "State.LastUpdatedEpoch").Int()),
						SlashEpoch:       int32(gjson.GetBytes(raw.RawDeal, "State.SlashEpoch").Int()),
					},
				}

				// Send parsed deal
				dealChan <- &ParsedDeal{
					DealID: raw.DealID,
					Deal:   deal,
				}
			}
		}()
	}

	// Goroutine to close dealChan after all workers finish
	go func() {
		wg.Wait()
		close(dealChan)
	}()

	go func() {
		defer close(rawDealChan)

		// Read in reasonable chunks to avoid loading 40GB at once
		const chunkSize = 4 * 1024 * 1024 // 4MB chunks

		var accumulated []byte
		buffer := make([]byte, chunkSize)

		for {
			n, err := reader.Read(buffer)
			if n > 0 {
				accumulated = append(accumulated, buffer[:n]...)

				// Process complete deals from accumulated data
				processed := findAndProcessCompleteDeals(accumulated, rawDealChan, walletIDs, depth)

				// Keep unprocessed remainder
				if processed > 0 && processed < len(accumulated) {
					copy(accumulated, accumulated[processed:])
					accumulated = accumulated[:len(accumulated)-processed]
				} else if processed == len(accumulated) {
					accumulated = accumulated[:0] // Clear all
				}
			}

			if err == io.EOF {
				// Process any remaining complete data
				if len(accumulated) > 0 {
					findAndProcessCompleteDeals(accumulated, rawDealChan, walletIDs, depth)
				}
				break
			}
			if err != nil {
				Logger.Errorw("failed to read response data", "error", err)
				return
			}
		}
	}()

	return dealChan, countingReader, closer, nil
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

	// Start pprof server for performance profiling
	pprofMux := http.NewServeMux()
	pprofMux.HandleFunc("/debug/pprof/", pprof.Index)
	pprofMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	pprofMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	pprofMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	pprofMux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	pprofServer := &http.Server{Addr: ":6060", Handler: pprofMux}

	go func() {
		Logger.Info("Starting pprof server on :6060")
		if err := pprofServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Logger.Warnw("pprof server error", "error", err)
		}
	}()

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

		// Shutdown pprof server
		if err := pprofServer.Shutdown(ctx2); err != nil {
			Logger.Warnw("failed to shutdown pprof server", "error", err)
		}

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
	ID         model.DealID
	ClientID   *model.WalletID
	Provider   string
	PieceCID   model.CID
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
	// If no data sources are configured, skip processing
	if d.dealZstURL == "" && d.lotusURL == "" {
		Logger.Info("no data sources configured, skipping deal tracking")
		return nil
	}

	headTime, err := util.GetLotusHeadTime(ctx, d.lotusURL, d.lotusToken)
	if err != nil {
		return errors.Wrapf(err, "failed to get lotus head time from %s", d.lotusURL)
	}

	var lastEpoch int32

	db := d.dbNoContext.WithContext(ctx)
	var wallets []model.Wallet
	err = db.Find(&wallets).Error
	if err != nil {
		return errors.Wrap(err, "failed to get wallets from database")
	}

	walletIDs := make(map[string]struct{})
	for _, wallet := range wallets {
		Logger.Infof("tracking deals for wallet %s", wallet.ActorID)
		walletIDs[wallet.ActorID] = struct{}{}
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
		Select("id", "deal_id", "state", "client_id", "client_actor_id", "provider", "piece_cid",
			"start_epoch", "end_epoch").Rows()
	if err != nil {
		return errors.WithStack(err)
	}
	for rows.Next() {
		var deal model.Deal
		err = rows.Scan(&deal.ID, &deal.DealID, &deal.State, &deal.ClientID, &deal.ClientActorID, &deal.Provider, &deal.PieceCID, &deal.StartEpoch, &deal.EndEpoch)
		if err != nil {
			return errors.WithStack(err)
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
	err = d.trackDeal(ctx, walletIDs, func(dealID uint64, deal Deal) error {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if deal.State.LastUpdatedEpoch > lastEpoch {
			lastEpoch = deal.State.LastUpdatedEpoch
		}
		_, ok := walletIDs[deal.Proposal.Client]
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

		var wallet model.Wallet
		if err := db.Where("actor_id = ?", deal.Proposal.Client).First(&wallet).Error; err != nil {
			return errors.Wrapf(err, "failed to find wallet for client %s", deal.Proposal.Client)
		}

		err = database.DoRetry(ctx, func() error {
			return db.Create(&model.Deal{
				DealID:           &dealID,
				State:            newState,
				ClientID:         &wallet.ID,
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

func (d *DealTracker) trackDeal(ctx context.Context, walletIDs map[string]struct{}, callback func(dealID uint64, deal Deal) error) error {
	dealStream, counter, closer, err := d.dealStateStreamCustom(ctx, walletIDs)
	if err != nil {
		return errors.WithStack(err)
	}
	defer closer.Close()

	// Start the download stats logger
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

	// Process deals directly - no batching needed since we're only getting wanted deals
	for parsed := range dealStream {
		if err := callback(parsed.DealID, parsed.Deal); err != nil {
			return errors.WithStack(err)
		}

		// Check if context is cancelled
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}

	return ctx.Err()
}

func (d *DealTracker) dealStateStreamCustom(ctx context.Context, walletIDs map[string]struct{}) (chan *ParsedDeal, Counter, io.Closer, error) {
	if d.dealZstURL != "" {
		Logger.Infof("getting deal state from %s", d.dealZstURL)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, d.dealZstURL, nil)
		if err != nil {
			return nil, nil, nil, errors.Wrapf(err, "failed to create request to get deal state zst file %s", d.dealZstURL)
		}
		return DealStateStreamFromHTTPRequest(req, 1, true, walletIDs)
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
	return DealStateStreamFromHTTPRequest(req, 2, false, walletIDs)
}

// Helper function for max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Find and process complete deal objects, return bytes processed
func findAndProcessCompleteDeals(data []byte, rawDealChan chan struct {
	DealID  uint64
	RawDeal json.RawMessage
}, walletIDs map[string]struct{}, depth int) int {

	// If data is small enough, just parse with gjson
	if len(data) < 8*1024*1024 { // Less than 8MB
		if isCompleteJSON(data, depth) {
			parseWithGjson(data, rawDealChan, walletIDs, depth)
			return len(data)
		}
		return 0 // Wait for more data
	}

	// For larger data, find individual complete deals
	processed := 0
	start := 0

	for {
		// Find the next complete deal starting from 'start'
		dealEnd := findNextCompleteDeal(data[start:])
		if dealEnd == -1 {
			break // No complete deal found
		}

		actualEnd := start + dealEnd
		dealData := data[start:actualEnd]

		// Process this deal
		if len(dealData) > 10 {
			processDealChunk(dealData, rawDealChan, walletIDs)
		}

		processed = actualEnd
		start = actualEnd

		// Find start of next deal (skip comma/whitespace)
		for start < len(data) && (data[start] == ',' || data[start] == ' ' || data[start] == '\t' || data[start] == '\n' || data[start] == '\r') {
			start++
			processed = start
		}

		if start >= len(data) {
			break
		}
	}

	return processed
}

// Check if data contains complete JSON
func isCompleteJSON(data []byte, depth int) bool {
	if len(data) < 10 {
		return false
	}

	// Simple check: starts with { and ends with }
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 {
		return false
	}

	return trimmed[0] == '{' && trimmed[len(trimmed)-1] == '}'
}

// Parse complete JSON with gjson
func parseWithGjson(data []byte, rawDealChan chan struct {
	DealID  uint64
	RawDeal json.RawMessage
}, walletIDs map[string]struct{}, depth int) {

	parsed := gjson.ParseBytes(data)

	// Handle depth wrapping
	dealsObj := parsed
	if depth == 2 {
		dealsObj = parsed.Get("result")
		if !dealsObj.Exists() {
			return
		}
	}

	// Process all deals
	dealsObj.ForEach(func(dealIDResult, dealData gjson.Result) bool {
		dealIDStr := dealIDResult.String()
		dealID, err := strconv.ParseUint(dealIDStr, 10, 64)
		if err != nil {
			return true
		}

		// Fast client check
		client := dealData.Get("Proposal.Client").String()
		if client == "" {
			return true
		}

		if _, want := walletIDs[client]; !want {
			return true
		}

		// Send to workers
		rawDealChan <- struct {
			DealID  uint64
			RawDeal json.RawMessage
		}{
			DealID:  dealID,
			RawDeal: json.RawMessage(dealData.Raw),
		}

		return true
	})
}

// Find the end of the next complete deal (simple approach)
func findNextCompleteDeal(data []byte) int {
	if len(data) < 10 {
		return -1
	}

	// Look for pattern: "123":{ ... }
	braceDepth := 0
	inString := false
	escape := false
	foundColon := false

	for i := 0; i < len(data); i++ {
		b := data[i]

		if escape {
			escape = false
			continue
		}

		switch b {
		case '\\':
			if inString {
				escape = true
			}
		case '"':
			inString = !inString
		case ':':
			if !inString && braceDepth == 0 {
				foundColon = true
			}
		case '{':
			if !inString && foundColon {
				braceDepth++
			}
		case '}':
			if !inString && braceDepth > 0 {
				braceDepth--
				if braceDepth == 0 {
					return i + 1 // Found complete deal
				}
			}
		}
	}

	return -1 // No complete deal found
}

// Process a single deal chunk
func processDealChunk(dealData []byte, rawDealChan chan struct {
	DealID  uint64
	RawDeal json.RawMessage
}, walletIDs map[string]struct{}) {

	// Extract ID and JSON from "123":{...}
	colonPos := bytes.IndexByte(dealData, ':')
	if colonPos <= 0 {
		return
	}

	// Extract ID
	idPart := bytes.TrimSpace(dealData[:colonPos])
	idPart = bytes.Trim(idPart, `"`)
	dealID, err := strconv.ParseUint(string(idPart), 10, 64)
	if err != nil {
		return
	}

	// Extract deal JSON
	dealJSON := bytes.TrimSpace(dealData[colonPos+1:])

	// Fast client check
	client := gjson.GetBytes(dealJSON, "Proposal.Client").String()
	if client == "" {
		return
	}

	if _, want := walletIDs[client]; !want {
		return
	}

	// Send to workers
	rawDealChan <- struct {
		DealID  uint64
		RawDeal json.RawMessage
	}{
		DealID:  dealID,
		RawDeal: json.RawMessage(dealJSON),
	}
}
