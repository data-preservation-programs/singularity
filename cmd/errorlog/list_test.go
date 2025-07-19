package errorlog

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/errorlog"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

func TestListCmd_Basic(t *testing.T) {
	db := testutil.TestDB(t)

	// Create test data
	testTime := time.Now()
	testLogs := []model.ErrorLog{
		{
			CreatedAt:  testTime.Add(-2 * time.Hour),
			EntityType: "preparation",
			EntityID:   "1",
			EventType:  "creation",
			Level:      model.ErrorLevelInfo,
			Message:    "Preparation created",
			Component:  "onboard",
			Metadata:   model.ConfigMap{"key1": "value1"},
		},
		{
			CreatedAt:  testTime.Add(-1 * time.Hour),
			EntityType: "schedule",
			EntityID:   "2",
			EventType:  "error",
			Level:      model.ErrorLevelError,
			Message:    "Schedule failed",
			Component:  "deal_schedule",
			Metadata:   model.ConfigMap{"key2": "value2"},
		},
	}

	for _, log := range testLogs {
		err := db.Create(&log).Error
		require.NoError(t, err)
	}

	// Test JSON output
	var output bytes.Buffer
	app := &cli.App{
		Writer: &output,
		Commands: []*cli.Command{
			{
				Name: "error",
				Subcommands: []*cli.Command{
					{
						Name: "log",
						Subcommands: []*cli.Command{
							ListCmd,
						},
					},
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "database-connection-string",
				Value: testutil.GenerateTestConnectionString(t),
			},
			&cli.BoolFlag{
				Name: "json",
			},
		},
	}

	// Test basic list command with JSON output
	err := app.RunContext(context.Background(), []string{
		"app", "error", "log", "list",
		"--json",
		"--database-connection-string=" + testutil.GenerateTestConnectionString(t),
	})
	require.NoError(t, err)

	// Parse JSON output
	var result struct {
		ErrorLogs []model.ErrorLog `json:"errorLogs"`
		Total     int64            `json:"total"`
		Limit     int              `json:"limit"`
		Offset    int              `json:"offset"`
		HasMore   bool             `json:"hasMore"`
	}
	err = json.Unmarshal(output.Bytes(), &result)
	require.NoError(t, err)
	require.Equal(t, int64(2), result.Total)
	require.Len(t, result.ErrorLogs, 2)
	require.False(t, result.HasMore)
}

func TestListCmd_Filtering(t *testing.T) {
	db := testutil.TestDB(t)

	// Create test data
	testLogs := []model.ErrorLog{
		{
			CreatedAt:  time.Now(),
			EntityType: "preparation",
			EntityID:   "1",
			EventType:  "creation",
			Level:      model.ErrorLevelInfo,
			Message:    "Preparation created",
			Component:  "onboard",
		},
		{
			CreatedAt:  time.Now(),
			EntityType: "schedule",
			EntityID:   "2",
			EventType:  "error",
			Level:      model.ErrorLevelError,
			Message:    "Schedule failed",
			Component:  "deal_schedule",
		},
	}

	for _, log := range testLogs {
		err := db.Create(&log).Error
		require.NoError(t, err)
	}

	// Test filtering by entity type
	var output bytes.Buffer
	app := &cli.App{
		Writer: &output,
		Commands: []*cli.Command{
			{
				Name: "error",
				Subcommands: []*cli.Command{
					{
						Name: "log",
						Subcommands: []*cli.Command{
							ListCmd,
						},
					},
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "database-connection-string",
				Value: testutil.GenerateTestConnectionString(t),
			},
			&cli.BoolFlag{
				Name: "json",
			},
		},
	}

	err := app.RunContext(context.Background(), []string{
		"app", "error", "log", "list",
		"--entity-type=preparation",
		"--json",
		"--database-connection-string=" + testutil.GenerateTestConnectionString(t),
	})
	require.NoError(t, err)

	var result struct {
		ErrorLogs []model.ErrorLog `json:"errorLogs"`
		Total     int64            `json:"total"`
	}
	err = json.Unmarshal(output.Bytes(), &result)
	require.NoError(t, err)
	require.Equal(t, int64(1), result.Total)
	require.Len(t, result.ErrorLogs, 1)
	require.Equal(t, "preparation", result.ErrorLogs[0].EntityType)
}

func TestListCmd_Validation(t *testing.T) {
	// Test invalid error level
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "error",
				Subcommands: []*cli.Command{
					{
						Name: "log",
						Subcommands: []*cli.Command{
							ListCmd,
						},
					},
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "database-connection-string",
				Value: "sqlite::memory:",
			},
		},
	}

	err := app.RunContext(context.Background(), []string{
		"app", "error", "log", "list",
		"--level=invalid",
		"--database-connection-string=sqlite::memory:",
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid error level")

	// Test invalid time format
	err = app.RunContext(context.Background(), []string{
		"app", "error", "log", "list",
		"--start-time=invalid",
		"--database-connection-string=sqlite::memory:",
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid start-time format")
}
