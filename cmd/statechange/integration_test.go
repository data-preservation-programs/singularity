package statechange

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/statetracker"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli/v2"
)

type IntegrationTestSuite struct {
	testutil.TestSuite
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) TestCompleteStateManagementWorkflow() {
	// Setup: Create test deals and state changes
	deals := []model.Deal{
		{
			ID:            200,
			State:         "proposed",
			Provider:      "f01234",
			ClientActorID: "f1abcdef",
		},
		{
			ID:            201,
			State:         "published",
			Provider:      "f01234",
			ClientActorID: "f1abcdef",
		},
		{
			ID:            202,
			State:         "active",
			Provider:      "f05678",
			ClientActorID: "f1fedcba",
		},
		{
			ID:            203,
			State:         "error",
			Provider:      "f01234",
			ClientActorID: "f1abcdef",
		},
	}

	for _, deal := range deals {
		s.NoError(s.DB.Create(&deal).Error)
	}

	// Create state change history using the state tracker
	tracker := statetracker.NewStateChangeTracker(s.DB)

	// Deal 200: proposed -> published
	metadata := &statetracker.StateChangeMetadata{
		Reason: "Deal proposal accepted",
	}
	previousState := model.DealState("proposed")
	s.NoError(tracker.TrackStateChangeWithDetails(
		s.Context, 200, &previousState, "published", nil, nil, "f01234", "f1abcdef", metadata,
	))

	// Deal 201: published -> active
	metadata = &statetracker.StateChangeMetadata{
		Reason: "Deal activated",
		ActivationEpoch: func() *int32 { epoch := int32(123456); return &epoch }(),
	}
	previousState = model.DealState("published")
	s.NoError(tracker.TrackStateChangeWithDetails(
		s.Context, 201, &previousState, "active", nil, nil, "f01234", "f1abcdef", metadata,
	))

	// Deal 203: proposed -> error
	metadata = &statetracker.StateChangeMetadata{
		Reason: "Deal failed",
		Error: "Connection timeout",
	}
	previousState = model.DealState("proposed")
	s.NoError(tracker.TrackStateChangeWithDetails(
		s.Context, 203, &previousState, "error", nil, nil, "f01234", "f1abcdef", metadata,
	))

	// Test 1: List all state changes
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "state",
				Subcommands: []*cli.Command{
					ListCmd,
					GetCmd,
					StatsCmd,
					RepairCmd,
				},
			},
		},
	}

	err := app.Run([]string{"test", "state", "list"})
	s.NoError(err)

	// Test 2: Get state changes for specific deal
	err = app.Run([]string{"test", "state", "get", "200"})
	s.NoError(err)

	// Test 3: Export state changes to JSON
	tmpFile, err := os.CreateTemp("", "integration-test-*.json")
	s.NoError(err)
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	err = app.Run([]string{"test", "state", "list", "--export", "json", "--output", tmpFile.Name()})
	s.NoError(err)

	// Verify exported JSON
	file, err := os.Open(tmpFile.Name())
	s.NoError(err)
	defer file.Close()

	var exportData struct {
		Metadata struct {
			ExportTime string `json:"exportTime"`
			TotalCount int    `json:"totalCount"`
		} `json:"metadata"`
		StateChanges []model.DealStateChange `json:"stateChanges"`
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&exportData)
	s.NoError(err)
	s.Greater(exportData.Metadata.TotalCount, 0)
	s.NotEmpty(exportData.Metadata.ExportTime)

	// Test 4: Reset error deal
	err = app.Run([]string{"test", "state", "repair", "reset-error-deals", "--deal-id", "203"})
	s.NoError(err)

	// Verify deal was reset
	var resetDeal model.Deal
	s.NoError(s.DB.First(&resetDeal, 203).Error)
	s.Equal(model.DealState("proposed"), resetDeal.State)

	// Test 5: Force state transition
	err = app.Run([]string{"test", "state", "repair", "force-transition", "200", "active", "--reason", "Integration test"})
	s.NoError(err)

	// Verify forced transition
	var transitionedDeal model.Deal
	s.NoError(s.DB.First(&transitionedDeal, 200).Error)
	s.Equal(model.DealState("active"), transitionedDeal.State)

	// Test 6: Get statistics
	err = app.Run([]string{"test", "state", "stats"})
	s.NoError(err)
}

func (s *IntegrationTestSuite) TestBulkOperations() {
	// Create multiple deals in error state
	errorDeals := make([]model.Deal, 10)
	for i := 0; i < 10; i++ {
		errorDeals[i] = model.Deal{
			ID:            model.DealID(300 + i),
			State:         "error",
			Provider:      "f01234",
			ClientActorID: "f1abcdef",
		}
		s.NoError(s.DB.Create(&errorDeals[i]).Error)
	}

	// Create some successful deals
	activeDeals := make([]model.Deal, 5)
	for i := 0; i < 5; i++ {
		activeDeals[i] = model.Deal{
			ID:            model.DealID(400 + i),
			State:         "active",
			Provider:      "f05678",
			ClientActorID: "f1fedcba",
		}
		s.NoError(s.DB.Create(&activeDeals[i]).Error)
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "state",
				Subcommands: []*cli.Command{
					RepairCmd,
				},
			},
		},
	}

	// Test bulk reset with limit
	err := app.Run([]string{"test", "state", "repair", "reset-error-deals", "--limit", "5"})
	s.NoError(err)

	// Verify only 5 deals were reset
	var resetCount int64
	s.DB.Model(&model.Deal{}).Where("state = ? AND id BETWEEN ? AND ?", "proposed", 300, 309).Count(&resetCount)
	s.Equal(int64(5), resetCount)

	// Verify remaining deals are still in error state
	var errorCount int64
	s.DB.Model(&model.Deal{}).Where("state = ? AND id BETWEEN ? AND ?", "error", 300, 309).Count(&errorCount)
	s.Equal(int64(5), errorCount)

	// Test bulk reset by provider
	err = app.Run([]string{"test", "state", "repair", "reset-error-deals", "--provider", "f01234"})
	s.NoError(err)

	// Verify all remaining error deals for provider f01234 were reset
	var finalErrorCount int64
	s.DB.Model(&model.Deal{}).Where("state = ? AND provider = ?", "error", "f01234").Count(&finalErrorCount)
	s.Equal(int64(0), finalErrorCount)

	// Verify active deals from other provider were not affected
	var activeCount int64
	s.DB.Model(&model.Deal{}).Where("state = ? AND provider = ?", "active", "f05678").Count(&activeCount)
	s.Equal(int64(5), activeCount)
}

func (s *IntegrationTestSuite) TestFilteringAndPagination() {
	// Create deals with various states and different timestamps
	testDeals := []struct {
		id       model.DealID
		state    model.DealState
		provider string
		client   string
		delay    time.Duration
	}{
		{500, "proposed", "f01111", "f1aaaa", 0},
		{501, "published", "f01111", "f1aaaa", time.Hour},
		{502, "active", "f02222", "f1bbbb", 2 * time.Hour},
		{503, "expired", "f02222", "f1bbbb", 3 * time.Hour},
		{504, "error", "f03333", "f1cccc", 4 * time.Hour},
	}

	baseTime := time.Now().Add(-5 * time.Hour)
	tracker := statetracker.NewStateChangeTracker(s.DB)

	for _, td := range testDeals {
		deal := model.Deal{
			ID:            td.id,
			State:         td.state,
			Provider:      td.provider,
			ClientActorID: td.client,
		}
		s.NoError(s.DB.Create(&deal).Error)

		// Create state change with specific timestamp
		stateChange := model.DealStateChange{
			DealID:        td.id,
			PreviousState: "",
			NewState:      td.state,
			Timestamp:     baseTime.Add(td.delay),
			ProviderID:    td.provider,
			ClientAddress: td.client,
			Metadata:      "{}",
		}
		s.NoError(s.DB.Create(&stateChange).Error)
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "state",
				Subcommands: []*cli.Command{
					ListCmd,
				},
			},
		},
	}

	// Test filtering by provider
	err := app.Run([]string{"test", "state", "list", "--provider", "f01111"})
	s.NoError(err)

	// Test filtering by state
	err = app.Run([]string{"test", "state", "list", "--state", "active"})
	s.NoError(err)

	// Test filtering by client
	err = app.Run([]string{"test", "state", "list", "--client", "f1bbbb"})
	s.NoError(err)

	// Test time range filtering
	startTime := baseTime.Add(time.Hour).Format(time.RFC3339)
	endTime := baseTime.Add(3 * time.Hour).Format(time.RFC3339)
	err = app.Run([]string{"test", "state", "list", "--start-time", startTime, "--end-time", endTime})
	s.NoError(err)

	// Test pagination
	err = app.Run([]string{"test", "state", "list", "--limit", "2", "--offset", "1"})
	s.NoError(err)

	// Test sorting
	err = app.Run([]string{"test", "state", "list", "--order-by", "dealId", "--order", "asc"})
	s.NoError(err)
}

func (s *IntegrationTestSuite) TestExportFormats() {
	// Create test data
	deal := model.Deal{
		ID:            600,
		State:         "active",
		Provider:      "f01234",
		ClientActorID: "f1abcdef",
	}
	s.NoError(s.DB.Create(&deal).Error)

	stateChange := model.DealStateChange{
		DealID:        600,
		PreviousState: "published",
		NewState:      "active",
		Timestamp:     time.Now(),
		ProviderID:    "f01234",
		ClientAddress: "f1abcdef",
		Metadata:      `{"reason":"test export"}`,
	}
	s.NoError(s.DB.Create(&stateChange).Error)

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "state",
				Subcommands: []*cli.Command{
					ListCmd,
					GetCmd,
				},
			},
		},
	}

	// Test CSV export from list
	csvFile, err := os.CreateTemp("", "integration-list-*.csv")
	s.NoError(err)
	defer os.Remove(csvFile.Name())
	csvFile.Close()

	err = app.Run([]string{"test", "state", "list", "--export", "csv", "--output", csvFile.Name()})
	s.NoError(err)

	// Verify CSV file has content
	stat, err := os.Stat(csvFile.Name())
	s.NoError(err)
	s.Greater(stat.Size(), int64(0))

	// Test JSON export from get
	jsonFile, err := os.CreateTemp("", "integration-get-*.json")
	s.NoError(err)
	defer os.Remove(jsonFile.Name())
	jsonFile.Close()

	err = app.Run([]string{"test", "state", "get", "600", "--export", "json", "--output", jsonFile.Name()})
	s.NoError(err)

	// Verify JSON file has content
	stat, err = os.Stat(jsonFile.Name())
	s.NoError(err)
	s.Greater(stat.Size(), int64(0))
}

func (s *IntegrationTestSuite) TestErrorHandling() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "state",
				Subcommands: []*cli.Command{
					ListCmd,
					GetCmd,
					RepairCmd,
				},
			},
		},
	}

	// Test various error conditions
	testCases := []struct {
		name        string
		args        []string
		expectError bool
	}{
		{
			name:        "Invalid deal ID in get",
			args:        []string{"test", "state", "get", "invalid"},
			expectError: true,
		},
		{
			name:        "Nonexistent deal in get",
			args:        []string{"test", "state", "get", "99999"},
			expectError: true,
		},
		{
			name:        "Invalid time format in list",
			args:        []string{"test", "state", "list", "--start-time", "invalid-time"},
			expectError: true,
		},
		{
			name:        "Invalid export format",
			args:        []string{"test", "state", "list", "--export", "xml"},
			expectError: true,
		},
		{
			name:        "Invalid state in force-transition",
			args:        []string{"test", "state", "repair", "force-transition", "123", "invalid-state"},
			expectError: true,
		},
		{
			name:        "Missing args in force-transition",
			args:        []string{"test", "state", "repair", "force-transition", "123"},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			err := app.Run(tc.args)
			if tc.expectError {
				s.Error(err, "Expected error for test case: %s", tc.name)
			} else {
				s.NoError(err, "Unexpected error for test case: %s", tc.name)
			}
		})
	}
}