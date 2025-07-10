package dealtracker

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBatchParser_shouldProcessDeal(t *testing.T) {
	walletIDs := map[string]struct{}{
		"t0100": {},
		"t0200": {},
	}

	parser := NewBatchParser(walletIDs, 10)

	// Test deal that should be processed (client in wallet set)
	dealJSON := `{
		"Proposal": {
			"Client": "t0100",
			"Provider": "t1provider1"
		}
	}`
	assert.True(t, parser.shouldProcessDeal([]byte(dealJSON)))

	// Test deal that should not be processed (client not in wallet set)
	dealJSON2 := `{
		"Proposal": {
			"Client": "t0300",
			"Provider": "t1provider1"
		}
	}`
	assert.False(t, parser.shouldProcessDeal([]byte(dealJSON2)))

	// Test invalid JSON
	assert.False(t, parser.shouldProcessDeal([]byte(`{"invalid":`)))
}

func TestBatchParser_parseDeal(t *testing.T) {
	walletIDs := map[string]struct{}{
		"t0100": {},
	}

	parser := NewBatchParser(walletIDs, 10)

	dealJSON := `{
		"DealID": "12345",
		"Proposal": {
			"Client": "t0100",
			"Provider": "t1provider1",
			"Label": "test-deal",
			"StartEpoch": 100,
			"EndEpoch": 200,
			"StoragePricePerEpoch": "1000",
			"ProviderCollateral": "5000",
			"ClientCollateral": "2000",
			"VerifiedDeal": true
		},
		"State": {
			"SectorStartEpoch": 105,
			"LastUpdatedEpoch": 150,
			"SlashEpoch": -1
		}
	}`

	result, err := parser.parseDeal([]byte(dealJSON))
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, uint64(12345), result.DealID)
	assert.Equal(t, "t0100", result.Deal.Proposal.Client)
	assert.Equal(t, "t1provider1", result.Deal.Proposal.Provider)
}

func TestBatchParser_ParseStream(t *testing.T) {
	walletIDs := map[string]struct{}{
		"t0100": {},
		"t0200": {},
	}

	parser := NewBatchParser(walletIDs, 2) // Small batch size for testing

	// Create a channel with test deals
	dealStream := make(chan []byte, 10)

	// Add deals - some matching, some not
	deals := []string{
		`{"DealID": "1", "Proposal": {"Client": "t0100", "Provider": "t1p1"}}`,
		`{"DealID": "2", "Proposal": {"Client": "t0300", "Provider": "t1p1"}}`, // Should be filtered
		`{"DealID": "3", "Proposal": {"Client": "t0200", "Provider": "t1p2"}}`,
		`{"DealID": "4", "Proposal": {"Client": "t0100", "Provider": "t1p3"}}`,
		`{"DealID": "5", "Proposal": {"Client": "t0400", "Provider": "t1p1"}}`, // Should be filtered
	}

	go func() {
		defer close(dealStream)
		for _, dealJSON := range deals {
			dealStream <- []byte(dealJSON)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	batches, err := parser.ParseStream(ctx, dealStream)
	require.NoError(t, err)

	var allDeals []ParsedDeal
	for batch := range batches {
		allDeals = append(allDeals, batch...)
	}

	// Should have 3 deals (IDs 1, 3, 4) after filtering
	require.Len(t, allDeals, 3)

	// Check the deal IDs are correct
	dealIDs := make([]uint64, len(allDeals))
	for i, deal := range allDeals {
		dealIDs[i] = deal.DealID
	}

	assert.Contains(t, dealIDs, uint64(1))
	assert.Contains(t, dealIDs, uint64(3))
	assert.Contains(t, dealIDs, uint64(4))
}
