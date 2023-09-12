package metrics

import (
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
)

func TestCollector(t *testing.T) {
	Enabled = true
	defer func() {
		Enabled = false
	}()
	collector := Collector{}
	collector.QueuePushJobEvent(PackJobEvent{
		Timestamp:  time.Now().Unix(),
		SourceType: "test",
		OutputType: "test",
		PieceSize:  128,
		PieceCID:   testutil.TestCid.String(),
		CarSize:    100,
		NumOfFiles: 100,
	})
	collector.QueueDealEvent(DealProposalEvent{
		Timestamp: time.Now().Unix(),
		PieceSize: 128,
		PieceCID:  testutil.TestCid.String(),
		DataCID:   testutil.TestCid.String(),
		Provider:  "f099",
		Client:    "f099",
		Verified:  true,
	})
	err := collector.Flush()
	require.NoError(t, err)
}
