package statechange

import (
	"errors"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/statechange"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli/v2"
)

type StatsCmdTestSuite struct {
	testutil.TestSuite
	mockHandler *statechange.MockStateChange
}

func TestStatsCmd(t *testing.T) {
	suite.Run(t, new(StatsCmdTestSuite))
}

func (s *StatsCmdTestSuite) SetupTest() {
	s.TestSuite.SetupTest()
	s.mockHandler = new(statechange.MockStateChange)
	statechange.Default = s.mockHandler
}

func (s *StatsCmdTestSuite) TearDownTest() {
	statechange.Default = &statechange.DefaultHandler{}
	s.TestSuite.TearDownTest()
}

func (s *StatsCmdTestSuite) TestStatsCmd_Success() {
	// Mock stats response
	expectedStats := map[string]interface{}{
		"totalStateChanges": 1250,
		"stateDistribution": map[string]interface{}{
			"proposed":          300,
			"published":         250,
			"active":           500,
			"expired":          150,
			"proposal_expired":  30,
			"rejected":         15,
			"slashed":          3,
			"error":            2,
		},
		"recentActivity": map[string]interface{}{
			"last24Hours": 45,
			"last7Days":   320,
			"last30Days":  890,
		},
		"providerStats": map[string]interface{}{
			"totalProviders": 25,
			"topProviders": []map[string]interface{}{
				{"providerId": "f01234", "stateChanges": 125},
				{"providerId": "f05678", "stateChanges": 98},
				{"providerId": "f09999", "stateChanges": 87},
			},
		},
		"clientStats": map[string]interface{}{
			"totalClients": 15,
			"topClients": []map[string]interface{}{
				{"clientAddress": "f1abcdef", "stateChanges": 200},
				{"clientAddress": "f1fedcba", "stateChanges": 150},
			},
		},
	}

	s.mockHandler.On("GetStateChangeStatsHandler", mock.Anything, mock.Anything).Return(expectedStats, nil)

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{StatsCmd},
	}

	// Test successful stats retrieval
	err := app.Run([]string{"test", "stats"})
	s.NoError(err)
	s.mockHandler.AssertExpectations(s.T())
}

func (s *StatsCmdTestSuite) TestStatsCmd_EmptyStats() {
	// Mock empty stats response
	expectedStats := map[string]interface{}{
		"totalStateChanges": 0,
		"stateDistribution": map[string]interface{}{},
		"recentActivity": map[string]interface{}{
			"last24Hours": 0,
			"last7Days":   0,
			"last30Days":  0,
		},
		"providerStats": map[string]interface{}{
			"totalProviders": 0,
			"topProviders":   []map[string]interface{}{},
		},
		"clientStats": map[string]interface{}{
			"totalClients": 0,
			"topClients":   []map[string]interface{}{},
		},
	}

	s.mockHandler.On("GetStateChangeStatsHandler", mock.Anything, mock.Anything).Return(expectedStats, nil)

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{StatsCmd},
	}

	// Test with empty stats
	err := app.Run([]string{"test", "stats"})
	s.NoError(err)
	s.mockHandler.AssertExpectations(s.T())
}

func (s *StatsCmdTestSuite) TestStatsCmd_MinimalStats() {
	// Mock minimal stats response (only essential fields)
	expectedStats := map[string]interface{}{
		"totalStateChanges": 10,
		"stateDistribution": map[string]interface{}{
			"proposed":  5,
			"published": 3,
			"active":    2,
		},
	}

	s.mockHandler.On("GetStateChangeStatsHandler", mock.Anything, mock.Anything).Return(expectedStats, nil)

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{StatsCmd},
	}

	// Test with minimal stats
	err := app.Run([]string{"test", "stats"})
	s.NoError(err)
	s.mockHandler.AssertExpectations(s.T())
}

func (s *StatsCmdTestSuite) TestStatsCmd_DatabaseError() {
	// Mock database error
	s.mockHandler.On("GetStateChangeStatsHandler", mock.Anything, mock.Anything).Return(map[string]interface{}{}, errors.New("database error"))

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{StatsCmd},
	}

	// Test database error handling
	err := app.Run([]string{"test", "stats"})
	s.Error(err)
	s.mockHandler.AssertExpectations(s.T())
}