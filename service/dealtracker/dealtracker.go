package dealtracker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"strconv"
	"sync"
	"time"

	"runtime"
	"sync/atomic"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/epochutil"
	"github.com/data-preservation-programs/singularity/service/healthcheck"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/google/uuid"
	"github.com/gotidy/ptr"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	"github.com/klauspost/compress/zstd"
	"github.com/tidwall/gjson"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"

	// Add riba's Lotus client library for filexp-style approach
	"code.riba.cloud/go/toolbox-interplanetary/fil"

	filabi "github.com/filecoin-project/go-state-types/abi"
	filbuiltin "github.com/filecoin-project/go-state-types/builtin"
	lchadt "github.com/filecoin-project/lotus/chain/actors/adt"
	lchmarket "github.com/filecoin-project/lotus/chain/actors/builtin/market"
	lchtypes "github.com/filecoin-project/lotus/chain/types"
	blocks "github.com/ipfs/go-block-format"
	ipldcbor "github.com/ipfs/go-ipld-cbor"
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

// StreamingRPCDealsReader provides an efficient streaming JSON reader for Lotus RPC responses
// Based on the approach from https://github.com/aschmahmann/filexp/commit/d16c79dd98c5179739b559ac5841beb69fef784d
type StreamingRPCDealsReader struct {
	ctx       context.Context
	reader    io.Reader
	walletIDs map[string]struct{}
	dealChan  chan *ParsedDeal
	counter   Counter
	closer    io.Closer

	// Progress tracking
	startTime       time.Time
	lastProgressLog time.Time
	totalDeals      int64
	filteredDeals   int64
	importedDeals   int64
}

// JsonEntry represents a single deal entry as it appears in the RPC response
type JsonEntry struct {
	DealID   *uint64 `json:",omitempty"`
	Proposal DealProposal
	State    DealState
}

// DealStateStreamFromLotusRPC creates an efficient streaming parser for Lotus RPC responses
// Based on the approach from https://github.com/aschmahmann/filexp/commit/d16c79dd98c5179739b559ac5841beb69fef784d
func DealStateStreamFromLotusRPC(request *http.Request, walletIDs map[string]struct{}) (chan *ParsedDeal, Counter, io.Closer, error) {
	Logger.Info("Starting RPC request to Lotus")

	//nolint: bodyclose
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, nil, errors.WithStack(err)
	}
	Logger.Infof("Got RPC response with status: %s", resp.Status)
	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, nil, nil, errors.Newf("failed to get deal state: %s", resp.Status)
	}

	countingReader := NewCountingReader(resp.Body)

	// Create channels
	dealChan := make(chan *ParsedDeal, 100)

	// Create the streaming reader
	reader := &StreamingRPCDealsReader{
		ctx:             request.Context(),
		reader:          countingReader,
		walletIDs:       walletIDs,
		dealChan:        dealChan,
		counter:         countingReader,
		closer:          resp.Body,
		startTime:       time.Now(),
		lastProgressLog: time.Now(),
	}

	Logger.Info("Starting RPC streaming processor")
	// Start the streaming processor
	go reader.processStream()

	return dealChan, countingReader, resp.Body, nil
}

// processStream handles the main streaming logic with worker goroutines
func (r *StreamingRPCDealsReader) processStream() {
	defer close(r.dealChan)
	defer r.logFinalStats()

	// Setup worker pool similar to filexp approach
	wrkCnt := runtime.NumCPU()
	if wrkCnt < 3 {
		wrkCnt = 3 // one iterator, and at least two encoders
	} else if wrkCnt > 12 {
		wrkCnt = 12 // do not overwhelm the blockstore
	}

	// Create work channels
	rawDealChan := make(chan struct {
		DealID  uint64
		RawDeal json.RawMessage
	}, wrkCnt*4)

	// Start worker goroutines
	eg, ctx := errgroup.WithContext(r.ctx)

	// JSON processing workers
	for i := 0; i < wrkCnt-1; i++ {
		eg.Go(func() error {
			return r.dealWorker(ctx, rawDealChan)
		})
	}

	// Main streaming reader
	eg.Go(func() error {
		defer close(rawDealChan)
		return r.streamReader(ctx, rawDealChan)
	})

	// Wait for all workers to complete
	_ = eg.Wait()
}

// logProgress logs processing statistics every 15 seconds
func (r *StreamingRPCDealsReader) logProgress() {
	now := time.Now()
	if now.Sub(r.lastProgressLog) >= 15*time.Second {
		elapsed := now.Sub(r.startTime)
		totalDeals := atomic.LoadInt64(&r.totalDeals)
		filteredDeals := atomic.LoadInt64(&r.filteredDeals)
		importedDeals := atomic.LoadInt64(&r.importedDeals)

		dealsPerSec := float64(totalDeals) / elapsed.Seconds()
		filteredPerSec := float64(filteredDeals) / elapsed.Seconds()
		importedPerSec := float64(importedDeals) / elapsed.Seconds()

		var matchRate float64
		if totalDeals > 0 {
			matchRate = float64(importedDeals) / float64(totalDeals) * 100
		}

		Logger.Infof("RPC Progress: %d total deals (%.1f/s), %d filtered (%.1f/s), %d imported (%.1f/s), %.2f%% match rate",
			totalDeals, dealsPerSec, filteredDeals, filteredPerSec, importedDeals, importedPerSec, matchRate)

		r.lastProgressLog = now
	}
}

// forceProgressLog forces a progress log (useful for debugging)
func (r *StreamingRPCDealsReader) forceProgressLog(message string) {
	elapsed := time.Since(r.startTime)
	totalDeals := atomic.LoadInt64(&r.totalDeals)
	filteredDeals := atomic.LoadInt64(&r.filteredDeals)
	importedDeals := atomic.LoadInt64(&r.importedDeals)

	Logger.Infof("RPC %s: %d total, %d filtered, %d imported (%.1fs elapsed)",
		message, totalDeals, filteredDeals, importedDeals, elapsed.Seconds())
}

// logFinalStats logs final processing statistics
func (r *StreamingRPCDealsReader) logFinalStats() {
	elapsed := time.Since(r.startTime)
	totalDeals := atomic.LoadInt64(&r.totalDeals)
	filteredDeals := atomic.LoadInt64(&r.filteredDeals)
	importedDeals := atomic.LoadInt64(&r.importedDeals)

	if elapsed > 0 {
		dealsPerSec := float64(totalDeals) / elapsed.Seconds()
		Logger.Infof("RPC Complete: %d total deals processed in %v (%.1f deals/s), %d filtered, %d imported (%.2f%% match rate)",
			totalDeals, elapsed.Truncate(time.Millisecond), dealsPerSec, filteredDeals, importedDeals,
			float64(importedDeals)/float64(totalDeals)*100)
	}
}

// dealWorker processes individual deals from the raw channel
func (r *StreamingRPCDealsReader) dealWorker(ctx context.Context, rawDealChan <-chan struct {
	DealID  uint64
	RawDeal json.RawMessage
}) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case raw, ok := <-rawDealChan:
			if !ok {
				return nil
			}

			// Parse deal using gjson for efficiency
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
			select {
			case <-ctx.Done():
				return ctx.Err()
			case r.dealChan <- &ParsedDeal{
				DealID: raw.DealID,
				Deal:   deal,
			}:
				atomic.AddInt64(&r.importedDeals, 1)
			}
		}
	}
}

// streamReader reads and parses the JSON stream
func (r *StreamingRPCDealsReader) streamReader(ctx context.Context, rawDealChan chan<- struct {
	DealID  uint64
	RawDeal json.RawMessage
}) error {
	Logger.Info("RPC streamReader starting - will use simple approach and read all data first")

	// Read all data into memory first (simpler approach for now)
	data, err := io.ReadAll(r.reader)
	if err != nil {
		return errors.Wrap(err, "failed to read RPC response")
	}

	Logger.Infof("Read %d bytes from RPC response", len(data))
	Logger.Infof("Response preview (first 200 bytes): %s", string(data[:min(len(data), 200)]))

	// Parse the JSON-RPC response
	var response struct {
		JSONRPC string          `json:"jsonrpc"`
		Result  json.RawMessage `json:"result"`
		Error   json.RawMessage `json:"error"`
		ID      int             `json:"id"`
	}

	if err := json.Unmarshal(data, &response); err != nil {
		return errors.Wrap(err, "failed to parse JSON-RPC response")
	}

	if response.Error != nil {
		return errors.Errorf("Lotus RPC error: %s", string(response.Error))
	}

	if response.Result == nil {
		return errors.New("empty result in RPC response")
	}

	Logger.Infof("Successfully parsed RPC response, processing deals from result (%d bytes)", len(response.Result))

	// Process deals from the result using gjson
	parsed := gjson.ParseBytes(response.Result)
	if !parsed.IsObject() {
		return errors.New("result is not a JSON object")
	}

	var processedCount int64
	parsed.ForEach(func(dealIDResult, dealData gjson.Result) bool {
		select {
		case <-ctx.Done():
			return false
		default:
		}

		dealIDStr := dealIDResult.String()
		dealID, err := strconv.ParseUint(dealIDStr, 10, 64)
		if err != nil {
			return true // Skip invalid deal IDs
		}

		atomic.AddInt64(&r.totalDeals, 1)
		processedCount++

		// Log first deal
		if processedCount == 1 {
			r.forceProgressLog("first deal processed")
		}

		// Quick client filter
		client := dealData.Get("Proposal.Client").String()
		if client == "" {
			return true
		}

		if _, want := r.walletIDs[client]; !want {
			return true
		}

		atomic.AddInt64(&r.filteredDeals, 1)

		// Send to worker
		select {
		case <-ctx.Done():
			return false
		case rawDealChan <- struct {
			DealID  uint64
			RawDeal json.RawMessage
		}{
			DealID:  dealID,
			RawDeal: json.RawMessage(dealData.Raw),
		}:
		}

		// Log progress periodically
		if processedCount%10000 == 0 {
			r.logProgress()
		}

		return true
	})

	Logger.Infof("Finished processing RPC deals: %d total processed", processedCount)
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
				dealsProcessed := counter.N()
				dealsPerSecond := counter.Speed()
				Logger.Infof("Processed %d deals with average speed %.0f deals/s", dealsProcessed, dealsPerSecond)
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

	Logger.Infof("getting deal state from %s via filexp-style approach", d.lotusURL)

	// Try filexp-style approach first
	dealChan, counter, closer, err := DealStateStreamFromLotusFilexp(ctx, d.lotusURL, d.lotusToken, walletIDs)
	if err != nil {
		Logger.Warnf("filexp-style approach failed: %v", err)
		Logger.Info("falling back to StateMarketDeals RPC approach")

		// Fall back to original StateMarketDeals RPC approach
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, d.lotusURL, nil)
		if err != nil {
			return nil, nil, nil, errors.Wrapf(err, "failed to create request to Lotus RPC %s", d.lotusURL)
		}

		// Set up the RPC request for StateMarketDeals
		if d.lotusToken != "" {
			req.Header.Set("Authorization", "Bearer "+d.lotusToken)
		}
		req.Header.Set("Content-Type", "application/json")

		return DealStateStreamFromLotusRPC(req, walletIDs)
	}

	return dealChan, counter, closer, nil
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

// FilexpMarketDealState matches the filexp structure for deal states
type FilexpMarketDealState struct {
	SectorNumber     filabi.SectorNumber
	SectorStartEpoch filabi.ChainEpoch
	LastUpdatedEpoch filabi.ChainEpoch
	SlashEpoch       filabi.ChainEpoch
}

// FilexpJsonEntry matches the filexp structure for deal entries
type FilexpJsonEntry struct {
	DealID   *filabi.DealID `json:",omitempty"`
	Proposal lchmarket.DealProposal
	State    FilexpMarketDealState
}

// DealStateStreamFromLotusFilexp creates a deal stream using filexp approach with riba library
func DealStateStreamFromLotusFilexp(ctx context.Context, lotusURL, lotusToken string, walletIDs map[string]struct{}) (chan *ParsedDeal, Counter, io.Closer, error) {
	Logger.Info("Starting filexp-style deal streaming")

	// Create riba Lotus client with proper signature
	lApi, apiCloser, err := fil.NewLotusDaemonAPIClientV0(ctx, lotusURL, 30, lotusToken)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to create Lotus API client")
	}

	// Get current chain head
	ts, err := lApi.ChainHead(ctx)
	if err != nil {
		apiCloser()
		return nil, nil, nil, errors.Wrap(err, "failed to get chain head")
	}

	Logger.Infof("Got chain head at height %d", ts.Height())

	// Create channels and counter with larger buffer
	dealChan := make(chan *ParsedDeal, 1000)
	counter := &simpleCounter{start: time.Now()}

	// Create closer that cleans up API connection
	closer := CloserFunc(func() error {
		apiCloser()
		return nil
	})

	// Start processing in background
	go func() {
		defer close(dealChan)
		err := processStateMarketDealsFilexp(ctx, lApi, ts, dealChan, walletIDs, counter)
		if err != nil {
			Logger.Errorw("failed to process state market deals", "error", err)
		}
	}()

	return dealChan, counter, closer, nil
}

// processStateMarketDealsFilexp processes deals using filexp approach with worker goroutines
func processStateMarketDealsFilexp(ctx context.Context, lApi fil.LotusDaemonAPIClientV0, ts *lchtypes.TipSet, dealChan chan<- *ParsedDeal, walletIDs map[string]struct{}, counter *simpleCounter) error {
	// Setup worker pool similar to filexp approach
	wrkCnt := runtime.NumCPU()
	if wrkCnt < 3 {
		wrkCnt = 3 // one iterator, and at least two encoders
	} else if wrkCnt > 12 {
		wrkCnt = 12 // do not overwhelm the block provider
	}

	Logger.Infof("Starting filexp-style processing with %d workers", wrkCnt)

	// Create work channels with large buffer like filexp (8<<10 = 8192)
	rawDealChan := make(chan FilexpJsonEntry, 8<<10)

	// Start worker goroutines
	eg, ctx := errgroup.WithContext(ctx)

	// JSON processing workers
	for i := 0; i < wrkCnt-1; i++ {
		eg.Go(func() error {
			return filexpDealWorker(ctx, rawDealChan, dealChan, walletIDs)
		})
	}

	// Main deal iterator (adapted from filexp)
	eg.Go(func() error {
		defer close(rawDealChan)
		return filexpDealIterator(ctx, lApi, ts, rawDealChan, walletIDs, counter)
	})

	// Wait for all workers to complete
	return eg.Wait()
}

// filexpDealIterator iterates through deals like filexp does
func filexpDealIterator(ctx context.Context, lApi fil.LotusDaemonAPIClientV0, ts *lchtypes.TipSet, rawDealChan chan<- FilexpJsonEntry, walletIDs map[string]struct{}, counter *simpleCounter) error {
	Logger.Info("Starting filexp deal iterator")

	// Get the storage market actor state (like filexp does)
	marketActor, err := lApi.StateGetActor(ctx, filbuiltin.StorageMarketActorAddr, ts.Key())
	if err != nil {
		return errors.Wrap(err, "failed to get storage market actor")
	}

	Logger.Infof("Market actor info: Code=%s, Head=%s, Nonce=%d, Balance=%s",
		marketActor.Code.String(), marketActor.Head.String(), marketActor.Nonce, marketActor.Balance.String())

	// Create a simple blockstore adapter for the Lotus API
	// Since riba doesn't export NewAPIBlockstore, we'll create a simple adapter
	bs := &lotusAPIBlockstore{api: lApi}
	cbs := ipldcbor.NewCborStore(bs)

	// Load market state
	marketState, err := lchmarket.Load(lchadt.WrapStore(ctx, cbs), marketActor)
	if err != nil {
		Logger.Warnf("Failed to load market state with actor code %s: %v", marketActor.Code.String(), err)
		Logger.Warn("This usually means the Lotus version doesn't recognize the current network's market actor version")
		Logger.Warn("Falling back to original StateMarketDeals RPC approach...")
		return errors.Wrap(err, "failed to load market state - consider using StateMarketDeals URL instead of direct Lotus API")
	}

	// Get proposals and states
	proposals, err := marketState.Proposals()
	if err != nil {
		return errors.Wrap(err, "failed to get proposals")
	}

	states, err := marketState.States()
	if err != nil {
		return errors.Wrap(err, "failed to get states")
	}

	Logger.Info("Got market state, starting deal iteration")

	// EXACTLY like filexp - create inner errgroup with limits
	egInner, ctx := errgroup.WithContext(ctx)
	wrkCnt := runtime.NumCPU()
	if wrkCnt < 3 {
		wrkCnt = 3 // one iterator, and at least two encoders
	} else if wrkCnt > 12 {
		wrkCnt = 12 // do not overwhelm the block provider
	}
	egInner.SetLimit(wrkCnt)

	var processedCount atomic.Int64
	var filteredCount atomic.Int64

	// EXACTLY like filexp - one goroutine for the iterator
	egInner.Go(func() error {
		// Iterate through proposals (like filexp does)
		return proposals.ForEach(func(did filabi.DealID, dp lchmarket.DealProposal) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			processed := processedCount.Add(1)
			atomic.AddInt64(&counter.count, 1)

			// Log progress periodically
			if processed%10000 == 0 {
				Logger.Infof("Processed %d deals, filtered %d for our wallets", processed, filteredCount.Load())
			}

			// MODIFICATION 1: Bail out if deal doesn't match our wallets (early filtering)
			clientAddr := dp.Client.String()
			if _, want := walletIDs[clientAddr]; !want {
				return nil // Skip deals not for our wallets
			}

			filteredCount.Add(1)

			// EXACTLY like filexp - spawn goroutine for each matching deal
			egInner.Go(func() error {
				// Get deal state (like filexp does)
				mds := FilexpMarketDealState{
					SectorNumber:     0,
					SectorStartEpoch: -1,
					LastUpdatedEpoch: -1,
					SlashEpoch:       -1,
				}

				s, found, err := states.Get(did)
				if err != nil {
					// Don't fail the whole process for one bad deal
					Logger.Warnw("failed to get deal state", "dealID", did, "error", err)
					return nil
				}
				if found {
					mds.SectorNumber = s.SectorNumber()
					mds.SectorStartEpoch = s.SectorStartEpoch()
					mds.LastUpdatedEpoch = s.LastUpdatedEpoch()
					mds.SlashEpoch = s.SlashEpoch()
				}

				// MODIFICATION 2: Feed matched deals to normal deal tracking logic instead of JSON output
				entry := FilexpJsonEntry{
					DealID:   &did,
					Proposal: dp,
					State:    mds,
				}

				select {
				case <-ctx.Done():
					return ctx.Err()
				case rawDealChan <- entry:
				}

				return nil
			})
			return nil
		})
	})

	// Wait for inner errgroup to complete
	return egInner.Wait()
}

// lotusAPIBlockstore is a simple adapter that implements blockstore interface for Lotus API
type lotusAPIBlockstore struct {
	api fil.LotusDaemonAPIClientV0
}

func (bs *lotusAPIBlockstore) Get(ctx context.Context, c cid.Cid) (blocks.Block, error) {
	blkData, err := bs.api.ChainReadObj(ctx, c)
	if err != nil {
		return nil, err
	}
	return blocks.NewBlockWithCid(blkData, c)
}

func (bs *lotusAPIBlockstore) Put(ctx context.Context, blk blocks.Block) error {
	return errors.New("put not supported")
}

// filexpDealWorker processes deals in parallel like filexp
func filexpDealWorker(ctx context.Context, rawDealChan <-chan FilexpJsonEntry, dealChan chan<- *ParsedDeal, walletIDs map[string]struct{}) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case entry, ok := <-rawDealChan:
			if !ok {
				return nil
			}

			// Convert DealLabel to string properly
			labelStr := ""
			if entry.Proposal.Label.IsString() {
				if str, err := entry.Proposal.Label.ToString(); err == nil {
					labelStr = str
				}
			} else if entry.Proposal.Label.IsBytes() {
				if bytes, err := entry.Proposal.Label.ToBytes(); err == nil {
					labelStr = string(bytes)
				}
			}

			// Convert filexp entry to our Deal structure
			deal := Deal{
				Proposal: DealProposal{
					PieceCID: Cid{
						Root: entry.Proposal.PieceCID.String(),
					},
					PieceSize:            int64(entry.Proposal.PieceSize),
					VerifiedDeal:         entry.Proposal.VerifiedDeal,
					Client:               entry.Proposal.Client.String(),
					Provider:             entry.Proposal.Provider.String(),
					Label:                labelStr,
					StartEpoch:           int32(entry.Proposal.StartEpoch),
					EndEpoch:             int32(entry.Proposal.EndEpoch),
					StoragePricePerEpoch: entry.Proposal.StoragePricePerEpoch.String(),
				},
				State: DealState{
					SectorStartEpoch: int32(entry.State.SectorStartEpoch),
					LastUpdatedEpoch: int32(entry.State.LastUpdatedEpoch),
					SlashEpoch:       int32(entry.State.SlashEpoch),
				},
			}

			// Send parsed deal
			select {
			case <-ctx.Done():
				return ctx.Err()
			case dealChan <- &ParsedDeal{
				DealID: uint64(*entry.DealID),
				Deal:   deal,
			}:
			}
		}
	}
}

// simpleCounter implements Counter interface for filexp compatibility
type simpleCounter struct {
	count int64
	start time.Time
}

func (c *simpleCounter) N() int64 {
	return atomic.LoadInt64(&c.count)
}

func (c *simpleCounter) Speed() float64 {
	if c.start.IsZero() {
		c.start = time.Now()
	}
	elapsed := time.Since(c.start)
	if elapsed == 0 {
		return 0
	}
	return float64(atomic.LoadInt64(&c.count)) / elapsed.Seconds()
}
