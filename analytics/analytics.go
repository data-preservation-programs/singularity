package analytics

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/fxamacker/cbor/v2"
	"github.com/ipfs/go-log/v2"
	"github.com/klauspost/compress/zstd"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const flushInterval = time.Hour

var Enabled = true

var logger = log.Logger("analytics")

// Init initializes the global variables 'Instance' and 'Enabled' based on values stored in the database
// and environment variables respectively. The function uses the 'instance_id' key to fetch the 'Instance'
// value from the database. If the 'Instance' value is already set, the function returns early. If the environment
// variable 'SINGULARITY_ANALYTICS' is set to "0", 'Enabled' is set to false.
//
// Parameters:
//   - ctx: The context for managing timeouts and cancellation.
//   - db: The gorm.DB instance for making database queries.
//
// Returns:
//   - An error if there are issues fetching the instance id from the database or if the database appears empty.
func Init(ctx context.Context, db *gorm.DB) error {
	if Instance != "" {
		return nil
	}
	var global model.Global
	where := clause.Where{Exprs: []clause.Expression{
		clause.Eq{Column: clause.Column{Name: "key"}, Value: "instance_id"},
	}}
	err := db.WithContext(ctx).Clauses(where).First(&global).Error
	if err != nil {
		return errors.Wrapf(err, "failed to get instance id, is the database empty?")
	}
	Instance = global.Value

	var identity model.Global
	where = clause.Where{Exprs: []clause.Expression{
		clause.Eq{Column: clause.Column{Name: "key"}, Value: "identity"},
	}}
	db.WithContext(ctx).Clauses(where).First(&identity)
	Identity = identity.Value

	if os.Getenv("SINGULARITY_ANALYTICS") != "1" {
		Enabled = false
	}
	return nil
}

var Instance string
var Identity string

type Collector struct {
	mu            sync.Mutex
	packJobEvents []PackJobEvent
	dealEvents    []DealProposalEvent
}

func (c *Collector) QueuePushJobEvent(event PackJobEvent) {
	if !Enabled {
		return
	}
	event.Timestamp = time.Now().Unix()
	event.Instance = Instance
	event.Identity = Identity
	c.mu.Lock()
	defer c.mu.Unlock()
	c.packJobEvents = append(c.packJobEvents, event)
}

func (c *Collector) QueueDealEvent(event DealProposalEvent) {
	if !Enabled {
		return
	}
	event.Timestamp = time.Now().Unix()
	event.Instance = Instance
	event.Identity = Identity
	c.mu.Lock()
	defer c.mu.Unlock()
	c.dealEvents = append(c.dealEvents, event)
}

var zstdEncoder, _ = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedDefault))

// Flush ensures that all buffered events in the Collector are sent to the
// remote metrics service. This function uses batch processing to send events
// in chunks. Each chunk can have a maximum of 100 events. If the global
// 'Enabled' flag is false, the function exits early without doing anything.
//
// Process:
//  1. The function checks if there are any events buffered in the Collector.
//  2. If there are more than 100 events of a specific type, it picks the
//     first 100 and sends them. Otherwise, it sends all available events.
//  3. It then encodes the selected events using CBOR and compresses them
//     using the zstandard compression algorithm.
//  4. The compressed data is encoded using Base64 encoding and sent to the
//     remote metrics service using an HTTP POST request.
//  5. If the HTTP response is not '200 OK', the function reads the response
//     body and returns an error.
//
// Returns:
//   - nil if all buffered events are successfully sent.
//   - An error if there are issues during any stage of the process.
func (c *Collector) Flush() error {
	if !Enabled {
		return nil
	}
	for {
		c.mu.Lock()
		if len(c.packJobEvents) == 0 && len(c.dealEvents) == 0 {
			c.mu.Unlock()
			return nil
		}
		var packJobs []PackJobEvent
		var dealEvents []DealProposalEvent
		if len(c.packJobEvents) > 100 {
			packJobs = c.packJobEvents[:100]
			c.packJobEvents = c.packJobEvents[100:]
		} else {
			packJobs = c.packJobEvents
			c.packJobEvents = nil
		}
		if len(c.dealEvents) > 100 {
			dealEvents = c.dealEvents[:100]
			c.dealEvents = c.dealEvents[100:]
		} else {
			dealEvents = c.dealEvents
			c.dealEvents = nil
		}
		c.mu.Unlock()
		events := Events{
			PackJobEvents: packJobs,
			DealEvents:    dealEvents,
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Hour*15)
		defer cancel()

		body := bytes.NewBuffer(nil)
		encoder := cbor.NewEncoder(body)
		err := encoder.Encode(events)
		if err != nil {
			logger.Error("failed to encode events", err)
			continue
		}

		compressed := zstdEncoder.EncodeAll(body.Bytes(), make([]byte, 0, body.Len()))

		request, err := http.NewRequestWithContext(ctx, http.MethodPost,
			"https://singularity-metrics.dataprogram.io/api",
			bytes.NewBufferString(base64.StdEncoding.EncodeToString(compressed)))
		if err != nil {
			return errors.WithStack(err)
		}

		request.Header.Set("Content-Type", "text/plain")
		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			return errors.WithStack(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			responseBody, err := io.ReadAll(resp.Body)
			if err != nil {
				return errors.WithStack(err)
			}
			return errors.Errorf("failed to send events: %s", responseBody)
		}
	}
}

func (c *Collector) Start(ctx context.Context) {
	timer := time.NewTimer(flushInterval)
	defer timer.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			timer.Reset(flushInterval)
		}
		//nolint:contextcheck
		c.Flush()
	}
}

var Default Collector = Collector{}
