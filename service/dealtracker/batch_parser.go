package dealtracker

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/bcicen/jstream"
	"github.com/cockroachdb/errors"
	"github.com/mitchellh/mapstructure"
	"github.com/tidwall/gjson"
)

// ParsedDeal represents a deal that has been parsed and matched against our wallets
type ParsedDeal struct {
	DealID uint64
	Deal   Deal
}

// BatchParser processes deals in batches while maintaining low memory usage
type BatchParser struct {
	walletIDs map[string]struct{}
	batchSize int
}

// NewBatchParser creates a new batch parser with the given wallet IDs and batch size
func NewBatchParser(walletIDs map[string]struct{}, batchSize int) *BatchParser {
	return &BatchParser{
		walletIDs: walletIDs,
		batchSize: batchSize,
	}
}

// ParseStream processes a stream of deals and returns batches of parsed deals that match our wallets
func (p *BatchParser) ParseStream(ctx context.Context, kvstream <-chan *jstream.MetaValue) (<-chan []ParsedDeal, error) {
	batches := make(chan []ParsedDeal, 2) // Small buffer for batches

	go func() {
		defer close(batches)

		currentBatch := make([]ParsedDeal, 0, p.batchSize)

		for {
			select {
			case <-ctx.Done():
				// Send any remaining deals
				if len(currentBatch) > 0 {
					select {
					case batches <- currentBatch:
					case <-ctx.Done():
					}
				}
				return

			case stream, ok := <-kvstream:
				if !ok {
					// Stream closed, send remaining batch
					if len(currentBatch) > 0 {
						batches <- currentBatch
					}
					return
				}

				keyValuePair, ok := stream.Value.(jstream.KV)
				if !ok {
					continue
				}

				// Fast path: check client ID without full parsing
				if !p.shouldProcessDeal(keyValuePair.Value) {
					continue
				}

				// Parse the full deal
				parsed, err := p.parseDeal(keyValuePair.Key, keyValuePair.Value)
				if err != nil {
					Logger.Warnw("failed to parse deal", "key", keyValuePair.Key, "error", err)
					continue
				}

				currentBatch = append(currentBatch, parsed)

				// Send batch if it's full
				if len(currentBatch) >= p.batchSize {
					select {
					case batches <- currentBatch:
						currentBatch = make([]ParsedDeal, 0, p.batchSize)
					case <-ctx.Done():
						return
					}
				}
			}
		}
	}()

	return batches, nil
}

// shouldProcessDeal uses gjson to quickly check if this deal's client is in our wallet set
func (p *BatchParser) shouldProcessDeal(value interface{}) bool {
	// First try the fast path with type assertion
	if obj, ok := value.(map[string]any); ok {
		if prop, ok := obj["Proposal"].(map[string]any); ok {
			if client, ok := prop["Client"].(string); ok {
				_, want := p.walletIDs[client]
				return want
			}
		}
	}

	// Fallback to gjson for other cases
	var jsonBytes []byte
	switch v := value.(type) {
	case []byte:
		jsonBytes = v
	case json.RawMessage:
		jsonBytes = []byte(v)
	default:
		// Convert to JSON for gjson
		var err error
		jsonBytes, err = json.Marshal(value)
		if err != nil {
			return false
		}
	}

	client := gjson.GetBytes(jsonBytes, "Proposal.Client").String()
	if client == "" {
		return false
	}

	_, want := p.walletIDs[client]
	return want
}

// parseDeal parses a full deal from the key-value pair
func (p *BatchParser) parseDeal(key string, value interface{}) (ParsedDeal, error) {
	dealID, err := strconv.ParseUint(key, 10, 64)
	if err != nil {
		return ParsedDeal{}, errors.Wrapf(err, "failed to parse deal ID %s", key)
	}

	var deal Deal

	// Use mapstructure if value is already a map
	if _, ok := value.(map[string]any); ok {
		err = mapstructure.Decode(value, &deal)
	} else {
		// Otherwise use json.Unmarshal
		var jsonBytes []byte
		switch v := value.(type) {
		case []byte:
			jsonBytes = v
		case json.RawMessage:
			jsonBytes = []byte(v)
		default:
			jsonBytes, err = json.Marshal(value)
			if err != nil {
				return ParsedDeal{}, errors.Wrap(err, "failed to marshal value")
			}
		}
		err = json.Unmarshal(jsonBytes, &deal)
	}

	if err != nil {
		return ParsedDeal{}, errors.Wrap(err, "failed to decode deal")
	}

	return ParsedDeal{
		DealID: dealID,
		Deal:   deal,
	}, nil
}
