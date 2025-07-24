package statechange

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/handler/statechange"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli/v2"
)

type ListCmdTestSuite struct {
	testutil.TestSuite
	mockHandler *statechange.MockStateChange
}

func TestListCmd(t *testing.T) {
	suite.Run(t, new(ListCmdTestSuite))
}

func (s *ListCmdTestSuite) SetupTest() {
	s.TestSuite.SetupTest()
	s.mockHandler = new(statechange.MockStateChange)
	statechange.Default = s.mockHandler
}

func (s *ListCmdTestSuite) TearDownTest() {
	statechange.Default = &statechange.DefaultHandler{}
	s.TestSuite.TearDownTest()
}

func (s *ListCmdTestSuite) TestListCmd_Success() {
	// Mock response
	now := time.Now()
	expectedResponse := statechange.StateChangeResponse{
		StateChanges: []model.DealStateChange{
			{
				ID:            1,
				DealID:        model.DealID(123),
				PreviousState: "proposed",
				NewState:      "published",
				Timestamp:     now,
				ProviderID:    "f01234",
				ClientAddress: "f1abcdef",
				Metadata:      "{}",
			},
		},
		Total:  1,
		Offset: nil,
		Limit:  nil,
	}

	s.mockHandler.On("ListStateChangesHandler", mock.Anything, mock.Anything, mock.MatchedBy(func(query model.DealStateChangeQuery) bool {
		return query.DealID == nil && query.State == nil
	})).Return(expectedResponse, nil)

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{ListCmd},
	}

	// Test successful list without filters
	err := app.Run([]string{"test", "list"})
	s.NoError(err)
	s.mockHandler.AssertExpectations(s.T())
}

func (s *ListCmdTestSuite) TestListCmd_WithFilters() {
	// Mock response
	dealID := model.DealID(123)
	state := model.DealState("published")
	provider := "f01234"
	client := "f1abcdef"
	startTime, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
	endTime, _ := time.Parse(time.RFC3339, "2023-12-31T23:59:59Z")

	expectedResponse := statechange.StateChangeResponse{
		StateChanges: []model.DealStateChange{},
		Total:        0,
		Offset:       nil,
		Limit:        nil,
	}

	s.mockHandler.On("ListStateChangesHandler", mock.Anything, mock.Anything, mock.MatchedBy(func(query model.DealStateChangeQuery) bool {
		return query.DealID != nil && *query.DealID == dealID &&
			query.State != nil && *query.State == state &&
			query.ProviderID != nil && *query.ProviderID == provider &&
			query.ClientAddress != nil && *query.ClientAddress == client &&
			query.StartTime != nil && query.StartTime.Equal(startTime) &&
			query.EndTime != nil && query.EndTime.Equal(endTime)
	})).Return(expectedResponse, nil)

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{ListCmd},
	}

	// Test with all filters
	err := app.Run([]string{"test", "list",
		"--deal-id", "123",
		"--state", "published",
		"--provider", "f01234",
		"--client", "f1abcdef",
		"--start-time", "2023-01-01T00:00:00Z",
		"--end-time", "2023-12-31T23:59:59Z",
	})
	s.NoError(err)
	s.mockHandler.AssertExpectations(s.T())
}

func (s *ListCmdTestSuite) TestListCmd_InvalidDealID() {
	// Create CLI context  
	app := &cli.App{
		Commands: []*cli.Command{ListCmd},
	}

	// Test with invalid deal ID
	err := app.Run([]string{"test", "list", "--deal-id", "invalid"})
	s.Error(err)
	s.Contains(err.Error(), "invalid deal ID format")
}

func (s *ListCmdTestSuite) TestListCmd_InvalidTimeFormat() {
	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{ListCmd},
	}

	// Test with invalid time format
	err := app.Run([]string{"test", "list", "--start-time", "invalid-time"})
	s.Error(err)
	s.Contains(err.Error(), "invalid start-time format")
}

func (s *ListCmdTestSuite) TestListCmd_ExportCSV() {
	// Mock response
	now := time.Now()
	expectedResponse := statechange.StateChangeResponse{
		StateChanges: []model.DealStateChange{
			{
				ID:            1,
				DealID:        model.DealID(123),
				PreviousState: "proposed",
				NewState:      "published",
				Timestamp:     now,
				ProviderID:    "f01234",
				ClientAddress: "f1abcdef",
				Metadata:      "{}",
			},
		},
		Total:  1,
		Offset: nil,
		Limit:  nil,
	}

	s.mockHandler.On("ListStateChangesHandler", mock.Anything, mock.Anything, mock.Anything).Return(expectedResponse, nil)

	// Create temporary file for export
	tmpFile, err := os.CreateTemp("", "test-export-*.csv")
	s.NoError(err)
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{ListCmd},
	}

	// Test CSV export
	err = app.Run([]string{"test", "list", "--export", "csv", "--output", tmpFile.Name()})
	s.NoError(err)

	// Verify file was created and has content
	stat, err := os.Stat(tmpFile.Name())
	s.NoError(err)
	s.Greater(stat.Size(), int64(0))

	s.mockHandler.AssertExpectations(s.T())
}

func (s *ListCmdTestSuite) TestListCmd_ExportJSON() {
	// Mock response
	now := time.Now()
	expectedResponse := statechange.StateChangeResponse{
		StateChanges: []model.DealStateChange{
			{
				ID:            1,
				DealID:        model.DealID(123),
				PreviousState: "proposed",
				NewState:      "published",
				Timestamp:     now,
				ProviderID:    "f01234",
				ClientAddress: "f1abcdef",
				Metadata:      "{}",
			},
		},
		Total:  1,
		Offset: nil,
		Limit:  nil,
	}

	s.mockHandler.On("ListStateChangesHandler", mock.Anything, mock.Anything, mock.Anything).Return(expectedResponse, nil)

	// Create temporary file for export
	tmpFile, err := os.CreateTemp("", "test-export-*.json")
	s.NoError(err)
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{ListCmd},
	}

	// Test JSON export
	err = app.Run([]string{"test", "list", "--export", "json", "--output", tmpFile.Name()})
	s.NoError(err)

	// Verify file was created and has content
	stat, err := os.Stat(tmpFile.Name())
	s.NoError(err)
	s.Greater(stat.Size(), int64(0))

	s.mockHandler.AssertExpectations(s.T())
}

func (s *ListCmdTestSuite) TestListCmd_UnsupportedExportFormat() {
	// Mock response
	expectedResponse := statechange.StateChangeResponse{
		StateChanges: []model.DealStateChange{},
		Total:        0,
		Offset:       nil,
		Limit:        nil,
	}

	s.mockHandler.On("ListStateChangesHandler", mock.Anything, mock.Anything, mock.Anything).Return(expectedResponse, nil)

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{ListCmd},
	}

	// Test unsupported export format
	err := app.Run([]string{"test", "list", "--export", "xml"})
	s.Error(err)
	s.Contains(err.Error(), "unsupported export format")
}