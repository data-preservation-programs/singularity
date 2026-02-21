package pdptracker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPDPTracker_Name(t *testing.T) {
	tracker := NewPDPTracker(nil, PDPConfig{PollingInterval: time.Minute}, nil, true)
	require.Equal(t, "PDPTracker", tracker.Name())
}
