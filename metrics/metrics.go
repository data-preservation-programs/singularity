package metrics

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/fxamacker/cbor/v2"
	"github.com/ipfs/go-log/v2"
	"github.com/klauspost/compress/zstd"
	"gorm.io/gorm"
)

var Enabled = true

var logger = log.Logger("metrics")

func Init(ctx context.Context, db *gorm.DB) error {
	if Instance != "" {
		return nil
	}
	var global model.Global
	err := db.WithContext(ctx).Where("key = ?", "instance_id").First(&global).Error
	if err != nil {
		return errors.Wrapf(err, "failed to get instance id, is the database empty?")
	}
	Instance = global.Value

	return nil
}

var Instance string

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
	c.mu.Lock()
	defer c.mu.Unlock()
	c.dealEvents = append(c.dealEvents, event)
}

var zstdEncoder, _ = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedDefault))

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
			logger.Error("failed to encode metrics", err)
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
			return errors.Errorf("failed to send metrics: %s", responseBody)
		}
	}
}

func (c *Collector) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Hour):
		}
		//nolint:contextcheck
		c.Flush()
	}
}

var Default Collector = Collector{}
