package statechange

import (
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

type GetCmdTestSuite struct {
	testutil.TestSuite
	mockHandler *statechange.MockStateChange
}

func TestGetCmd(t *testing.T) {
	suite.Run(t, new(GetCmdTestSuite))
}

func (s *GetCmdTestSuite) SetupTest() {
	s.TestSuite.SetupTest()
	s.mockHandler = new(statechange.MockStateChange)
	statechange.Default = s.mockHandler
}

func (s *GetCmdTestSuite) TearDownTest() {
	statechange.Default = &statechange.DefaultHandler{}
	s.TestSuite.TearDownTest()
}

func (s *GetCmdTestSuite) TestGetCmd_Success() {
	// Mock response
	now := time.Now()
	expectedStateChanges := []model.DealStateChange{
		{
			ID:            1,
			DealID:        model.DealID(123),
			PreviousState: "",
			NewState:      "proposed",
			Timestamp:     now.Add(-2 * time.Hour),
			ProviderID:    "f01234",
			ClientAddress: "f1abcdef",
			Metadata:      "{}",
		},
		{
			ID:            2,
			DealID:        model.DealID(123),
			PreviousState: "proposed",
			NewState:      "published",
			Timestamp:     now.Add(-1 * time.Hour),
			ProviderID:    "f01234",
			ClientAddress: "f1abcdef",
			Metadata:      "{}",
		},
		{
			ID:            3,
			DealID:        model.DealID(123),
			PreviousState: "published",
			NewState:      "active",
			Timestamp:     now,
			ProviderID:    "f01234",
			ClientAddress: "f1abcdef",
			Metadata:      "{}",
		},
	}

	s.mockHandler.On("GetDealStateChangesHandler", mock.Anything, mock.Anything, model.DealID(123)).Return(expectedStateChanges, nil)

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{GetCmd},
	}

	// Test successful get
	err := app.Run([]string{"test", "get", "123"})
	s.NoError(err)
	s.mockHandler.AssertExpectations(s.T())
}

func (s *GetCmdTestSuite) TestGetCmd_NoDealID() {
	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{GetCmd},
	}

	// Test without deal ID
	err := app.Run([]string{"test", "get"})
	s.Error(err)
	s.Contains(err.Error(), "deal ID is required")
}

func (s *GetCmdTestSuite) TestGetCmd_InvalidDealID() {
	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{GetCmd},
	}

	// Test with invalid deal ID
	err := app.Run([]string{"test", "get", "invalid"})
	s.Error(err)
	s.Contains(err.Error(), "invalid deal ID format")
}

func (s *GetCmdTestSuite) TestGetCmd_NoStateChanges() {
	// Mock empty response
	s.mockHandler.On("GetDealStateChangesHandler", mock.Anything, mock.Anything, model.DealID(123)).Return([]model.DealStateChange{}, nil)

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{GetCmd},
	}

	// Test with deal that has no state changes
	err := app.Run([]string{"test", "get", "123"})
	s.NoError(err)
	s.mockHandler.AssertExpectations(s.T())
}

func (s *GetCmdTestSuite) TestGetCmd_ExportCSV() {
	// Mock response
	now := time.Now()
	expectedStateChanges := []model.DealStateChange{
		{
			ID:            1,
			DealID:        model.DealID(123),
			PreviousState: "",
			NewState:      "proposed",
			Timestamp:     now,
			ProviderID:    "f01234",
			ClientAddress: "f1abcdef",
			Metadata:      "{}",
		},
	}

	s.mockHandler.On("GetDealStateChangesHandler", mock.Anything, mock.Anything, model.DealID(123)).Return(expectedStateChanges, nil)

	// Create temporary file for export
	tmpFile, err := os.CreateTemp("", "test-deal-export-*.csv")
	s.NoError(err)
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{GetCmd},
	}

	// Test CSV export
	err = app.Run([]string{"test", "get", "123", "--export", "csv", "--output", tmpFile.Name()})
	s.NoError(err)

	// Verify file was created and has content
	stat, err := os.Stat(tmpFile.Name())
	s.NoError(err)
	s.Greater(stat.Size(), int64(0))

	s.mockHandler.AssertExpectations(s.T())
}

func (s *GetCmdTestSuite) TestGetCmd_ExportJSON() {
	// Mock response
	now := time.Now()
	expectedStateChanges := []model.DealStateChange{
		{
			ID:            1,
			DealID:        model.DealID(123),
			PreviousState: "",
			NewState:      "proposed",
			Timestamp:     now,
			ProviderID:    "f01234",
			ClientAddress: "f1abcdef",
			Metadata:      "{}",
		},
	}

	s.mockHandler.On("GetDealStateChangesHandler", mock.Anything, mock.Anything, model.DealID(123)).Return(expectedStateChanges, nil)

	// Create temporary file for export
	tmpFile, err := os.CreateTemp("", "test-deal-export-*.json")
	s.NoError(err)
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{GetCmd},
	}

	// Test JSON export
	err = app.Run([]string{"test", "get", "123", "--export", "json", "--output", tmpFile.Name()})
	s.NoError(err)

	// Verify file was created and has content
	stat, err := os.Stat(tmpFile.Name())
	s.NoError(err)
	s.Greater(stat.Size(), int64(0))

	s.mockHandler.AssertExpectations(s.T())
}

func (s *GetCmdTestSuite) TestGetCmd_UnsupportedExportFormat() {
	// Mock response
	expectedStateChanges := []model.DealStateChange{
		{
			ID:            1,
			DealID:        model.DealID(123),
			PreviousState: "",
			NewState:      "proposed",
			Timestamp:     time.Now(),
			ProviderID:    "f01234",
			ClientAddress: "f1abcdef",
			Metadata:      "{}",
		},
	}

	s.mockHandler.On("GetDealStateChangesHandler", mock.Anything, mock.Anything, model.DealID(123)).Return(expectedStateChanges, nil)

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{GetCmd},
	}

	// Test unsupported export format
	err := app.Run([]string{"test", "get", "123", "--export", "xml"})
	s.Error(err)
	s.Contains(err.Error(), "unsupported export format")

	s.mockHandler.AssertExpectations(s.T())
}

func (s *GetCmdTestSuite) TestGetCmd_ExportWithEmptyData() {
	// Mock empty response but export should still work
	s.mockHandler.On("GetDealStateChangesHandler", mock.Anything, mock.Anything, model.DealID(123)).Return([]model.DealStateChange{}, nil)

	// Create temporary file for export
	tmpFile, err := os.CreateTemp("", "test-deal-empty-export-*.csv")
	s.NoError(err)
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{GetCmd},
	}

	// Test CSV export with empty data
	err = app.Run([]string{"test", "get", "123", "--export", "csv", "--output", tmpFile.Name()})
	s.NoError(err)

	// Verify file was created (should have header even with no data)
	stat, err := os.Stat(tmpFile.Name())
	s.NoError(err)
	s.Greater(stat.Size(), int64(0))

	s.mockHandler.AssertExpectations(s.T())
}