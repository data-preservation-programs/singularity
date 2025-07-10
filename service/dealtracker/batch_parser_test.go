package dealtracker

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bcicen/jstream"
	"github.com/stretchr/testify/require"
)

func TestBatchParser_ParseStream(t *testing.T) {
	// Create test data
	walletIDs := map[string]struct{}{
		"t0100": {},
		"t0200": {},
	}

	parser := NewBatchParser(walletIDs, 2) // Small batch size for testing

	// Create a mock stream
	kvstream := make(chan *jstream.MetaValue, 10)

	// Add test deals
	testDeals := []struct {
		key   string
		value map[string]interface{}
	}{
		{
			key: "1",
			value: map[string]interface{}{
				"Proposal": map[string]interface{}{
					"Client":   "t0100", // Should match
					"Provider": "f01000",
					"PieceCID": map[string]interface{}{"/": "bafy2bzaced..."},
				},
				"State": map[string]interface{}{
					"SectorStartEpoch": 100,
				},
			},
		},
		{
			key: "2",
			value: map[string]interface{}{
				"Proposal": map[string]interface{}{
					"Client":   "t0300", // Should not match
					"Provider": "f01000",
				},
				"State": map[string]interface{}{},
			},
		},
		{
			key: "3",
			value: map[string]interface{}{
				"Proposal": map[string]interface{}{
					"Client":   "t0200", // Should match
					"Provider": "f01000",
				},
				"State": map[string]interface{}{},
			},
		},
	}

	// Send test data
	go func() {
		for _, td := range testDeals {
			kvstream <- &jstream.MetaValue{
				Value: jstream.KV{
					Key:   td.key,
					Value: td.value,
				},
			}
		}
		close(kvstream)
	}()

	ctx := context.Background()
	batches, err := parser.ParseStream(ctx, kvstream)
	require.NoError(t, err)

	// Collect results
	var results []ParsedDeal
	for batch := range batches {
		results = append(results, batch...)
	}

	// Should only have 2 results (deals 1 and 3)
	require.Len(t, results, 2)
	require.Equal(t, uint64(1), results[0].DealID)
	require.Equal(t, "t0100", results[0].Deal.Proposal.Client)
	require.Equal(t, uint64(3), results[1].DealID)
	require.Equal(t, "t0200", results[1].Deal.Proposal.Client)
}

func BenchmarkBatchParser_ShouldProcessDeal(b *testing.B) {
	walletIDs := map[string]struct{}{
		"t0100": {},
		"t0200": {},
		"t0300": {},
	}

	parser := NewBatchParser(walletIDs, 100)

	// Create test deal as map
	dealMap := map[string]interface{}{
		"Proposal": map[string]interface{}{
			"Client":    "t0100",
			"Provider":  "f01000",
			"PieceCID":  map[string]interface{}{"/": "bafy2bzaced..."},
			"PieceSize": 34359738368,
			"Label":     "some-label",
		},
		"State": map[string]interface{}{
			"SectorStartEpoch": 100,
			"LastUpdatedEpoch": 200,
		},
	}

	// Also create as JSON bytes for comparison
	dealJSON, _ := json.Marshal(dealMap)

	b.Run("MapInterface", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = parser.shouldProcessDeal(dealMap)
		}
	})

	b.Run("JSONBytes", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = parser.shouldProcessDeal(dealJSON)
		}
	})
}

func BenchmarkBatchParser_ParseDeal(b *testing.B) {
	walletIDs := map[string]struct{}{"t0100": {}}
	parser := NewBatchParser(walletIDs, 100)

	// Create a realistic deal
	dealMap := map[string]interface{}{
		"Proposal": map[string]interface{}{
			"PieceCID": map[string]interface{}{
				"/": "baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq",
			},
			"PieceSize":            int64(34359738368),
			"VerifiedDeal":         true,
			"Client":               "t0100",
			"Provider":             "f01000",
			"Label":                "bagboea4b5abcatlxechwbp7kjpjguna6r6q7ejrhe6mdp3lf34pmswn27pkkiekz",
			"StartEpoch":           int32(100),
			"EndEpoch":             int32(999999999),
			"StoragePricePerEpoch": "0",
		},
		"State": map[string]interface{}{
			"SectorStartEpoch": int32(100),
			"LastUpdatedEpoch": int32(200),
			"SlashEpoch":       int32(-1),
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := parser.parseDeal(fmt.Sprintf("%d", i), dealMap)
		if err != nil {
			b.Fatal(err)
		}
	}
}
