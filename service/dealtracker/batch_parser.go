package dealtracker

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/tidwall/gjson"
)

// ParsedDeal represents a fully parsed deal with metadata
type ParsedDeal struct {
	DealID uint64
	Deal   Deal
}

// BatchParser handles batching and filtering of deals to reduce database contention
type BatchParser struct {
	walletIDs map[string]struct{}
	batchSize int
}

// NewBatchParser creates a new batch parser with the given wallet IDs and batch size
func NewBatchParser(walletIDs map[string]struct{}, batchSize int) *BatchParser {
	if batchSize <= 0 {
		batchSize = 100 // Default batch size
	}
	return &BatchParser{
		walletIDs: walletIDs,
		batchSize: batchSize,
	}
}

// ParseStream processes a stream of raw JSON deal data and returns batches of parsed deals that match our wallets
func (p *BatchParser) ParseStream(ctx context.Context, dealStream <-chan []byte) (<-chan []ParsedDeal, error) {
	batches := make(chan []ParsedDeal, 2) // Small buffer for batches

	go func() {
		defer close(batches)

		var currentBatch []ParsedDeal

		for {
			select {
			case <-ctx.Done():
				// Send any remaining batch before exiting
				if len(currentBatch) > 0 {
					select {
					case batches <- currentBatch:
					case <-ctx.Done():
					}
				}
				return

			case dealData, ok := <-dealStream:
				if !ok {
					// Stream closed, send final batch if any
					if len(currentBatch) > 0 {
						select {
						case batches <- currentBatch:
						case <-ctx.Done():
						}
					}
					return
				}

				// Quick check if we should process this deal
				if !p.shouldProcessDeal(dealData) {
					continue // Skip deals for wallets we don't track
				}

				// Parse the full deal only if it matches our criteria
				parsedDeal, err := p.parseDeal(dealData)
				if err != nil {
					Logger.Warnw("failed to parse deal", "error", err)
					continue
				}

				currentBatch = append(currentBatch, *parsedDeal)

				// Send batch when it reaches the configured size
				if len(currentBatch) >= p.batchSize {
					select {
					case batches <- currentBatch:
						currentBatch = nil // Reset for next batch
					case <-ctx.Done():
						return
					}
				}
			}
		}
	}()

	return batches, nil
}

// shouldProcessDeal quickly checks if a deal should be processed based on client ID
func (p *BatchParser) shouldProcessDeal(data []byte) bool {
	// Fast client check using gjson
	result := gjson.GetBytes(data, "Proposal.Client")
	if !result.Exists() {
		return false
	}

	client := result.String()
	_, want := p.walletIDs[client]
	return want
}

// parseDeal parses a full deal from raw JSON data
func (p *BatchParser) parseDeal(data []byte) (*ParsedDeal, error) {
	dealID, err := strconv.ParseUint(gjson.GetBytes(data, "DealID").String(), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse deal ID from raw data")
	}

	var deal Deal
	err = json.Unmarshal(data, &deal)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode deal")
	}

	return &ParsedDeal{
		DealID: dealID,
		Deal:   deal,
	}, nil
}
