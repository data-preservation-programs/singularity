package dealtemplate

import (
	"context"
	"flag"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/dealtemplate"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

// TestValidateUpdateTemplateInputs tests the validation logic for update command
func TestValidateUpdateTemplateInputs(t *testing.T) {
	tests := []struct {
		name          string
		setup         func() *cli.Context
		expectError   bool
		errorContains string
	}{
		{
			name: "valid update with positive price",
			setup: func() *cli.Context {
				app := &cli.App{}
				set := flag.NewFlagSet("test", 0)

				// Add the required flags
				set.Float64("price-per-gb", 0.002, "")

				ctx := cli.NewContext(app, set, nil)
				ctx.Set("price-per-gb", "0.002")
				return ctx
			},
			expectError: false,
		},
		{
			name: "invalid update with negative price",
			setup: func() *cli.Context {
				app := &cli.App{}
				set := flag.NewFlagSet("test", 0)

				set.Float64("price-per-gb", -0.001, "")

				ctx := cli.NewContext(app, set, nil)
				ctx.Set("price-per-gb", "-0.001")
				return ctx
			},
			expectError:   true,
			errorContains: "deal price per GB must be non-negative",
		},
		{
			name: "invalid provider format",
			setup: func() *cli.Context {
				app := &cli.App{}
				set := flag.NewFlagSet("test", 0)

				set.String("provider", "invalid", "")

				ctx := cli.NewContext(app, set, nil)
				ctx.Set("provider", "invalid")
				return ctx
			},
			expectError:   true,
			errorContains: "deal provider must be a valid storage provider ID",
		},
		{
			name: "valid provider format",
			setup: func() *cli.Context {
				app := &cli.App{}
				set := flag.NewFlagSet("test", 0)

				set.String("provider", "f01234", "")

				ctx := cli.NewContext(app, set, nil)
				ctx.Set("provider", "f01234")
				return ctx
			},
			expectError: false,
		},
		{
			name: "invalid http header format",
			setup: func() *cli.Context {
				app := &cli.App{}
				set := flag.NewFlagSet("test", 0)

				set.StringSlice("http-header", []string{"invalid-header"}, "")

				ctx := cli.NewContext(app, set, nil)
				ctx.Set("http-header", "invalid-header")
				return ctx
			},
			expectError:   true,
			errorContains: "invalid HTTP header format",
		},
		{
			name: "replace-piece-cids without piece CIDs",
			setup: func() *cli.Context {
				app := &cli.App{}
				set := flag.NewFlagSet("test", 0)

				set.Bool("replace-piece-cids", true, "")

				ctx := cli.NewContext(app, set, nil)
				ctx.Set("replace-piece-cids", "true")
				return ctx
			},
			expectError:   true,
			errorContains: "--replace-piece-cids can only be used with --allowed-piece-cid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.setup()
			err := validateUpdateTemplateInputs(ctx)

			if tt.expectError {
				require.Error(t, err)
				if tt.errorContains != "" {
					require.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// TestUpdateHandlerIntegration tests the update handler integration
func TestUpdateHandlerIntegration(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		handler := dealtemplate.Default

		// Create a test template first
		createReq := dealtemplate.CreateRequest{
			Name:           "integration-test-template",
			Description:    "Test template for integration testing",
			DealPricePerGB: 0.001,
			DealProvider:   "f01000",
		}
		created, err := handler.CreateHandler(ctx, db, createReq)
		require.NoError(t, err)

		t.Run("basic update functionality", func(t *testing.T) {
			// Test updating description
			newDesc := "Updated integration test description"
			updateReq := dealtemplate.UpdateRequest{
				Description: &newDesc,
			}

			updated, err := handler.UpdateHandler(ctx, db, created.Name, updateReq)
			require.NoError(t, err)
			require.NotNil(t, updated)
			require.Equal(t, "Updated integration test description", updated.Description)
			require.Equal(t, 0.001, updated.DealConfig.DealPricePerGb) // unchanged
		})

		t.Run("update by ID", func(t *testing.T) {
			newPrice := 0.005
			updateReq := dealtemplate.UpdateRequest{
				DealPricePerGB: &newPrice,
			}

			// Convert ID to string for testing
			idStr := "1" // Assuming first template has ID 1
			updated, err := handler.UpdateHandler(ctx, db, idStr, updateReq)
			require.NoError(t, err)
			require.NotNil(t, updated)
			require.Equal(t, 0.005, updated.DealConfig.DealPricePerGb)
		})

		t.Run("partial update preserves other fields", func(t *testing.T) {
			// Only update provider, leave other fields unchanged
			newProvider := "f01001"
			updateReq := dealtemplate.UpdateRequest{
				DealProvider: &newProvider,
			}

			updated, err := handler.UpdateHandler(ctx, db, created.Name, updateReq)
			require.NoError(t, err)
			require.NotNil(t, updated)
			require.Equal(t, "f01001", updated.DealConfig.DealProvider)
			require.Equal(t, "Updated integration test description", updated.Description) // preserved from previous test
		})

		t.Run("nonexistent template error", func(t *testing.T) {
			newDesc := "This should fail"
			updateReq := dealtemplate.UpdateRequest{
				Description: &newDesc,
			}

			_, err := handler.UpdateHandler(ctx, db, "nonexistent-template-999", updateReq)
			require.Error(t, err)
		})

		t.Run("empty update request", func(t *testing.T) {
			// Empty update should succeed but not change anything
			updateReq := dealtemplate.UpdateRequest{}

			updated, err := handler.UpdateHandler(ctx, db, created.Name, updateReq)
			require.NoError(t, err)
			require.NotNil(t, updated)
			require.Equal(t, created.ID, updated.ID)
		})

		// Cleanup
		err = handler.DeleteHandler(ctx, db, created.Name)
		require.NoError(t, err)
	})
}

// TestUpdateCmdFlags tests that all expected flags are available
func TestUpdateCmdFlags(t *testing.T) {
	// Test that all expected flags exist in the UpdateCmd
	expectedFlags := []string{
		"name", "description", "provider", "price-per-gb", "price-per-gb-epoch",
		"price-per-deal", "duration", "start-delay", "verified", "keep-unsealed",
		"ipni", "url-template", "http-header", "notes", "force",
		"allowed-piece-cid", "allowed-piece-cid-file", "replace-piece-cids",
		"schedule-cron", "schedule-deal-number", "schedule-deal-size",
		"total-deal-number", "total-deal-size", "max-pending-deal-number",
		"max-pending-deal-size",
	}

	flagNames := make(map[string]bool)
	for _, flag := range UpdateCmd.Flags {
		switch f := flag.(type) {
		case *cli.StringFlag:
			flagNames[f.Name] = true
		case *cli.Float64Flag:
			flagNames[f.Name] = true
		case *cli.DurationFlag:
			flagNames[f.Name] = true
		case *cli.BoolFlag:
			flagNames[f.Name] = true
		case *cli.IntFlag:
			flagNames[f.Name] = true
		case *cli.StringSliceFlag:
			flagNames[f.Name] = true
		}
	}

	for _, expectedFlag := range expectedFlags {
		require.True(t, flagNames[expectedFlag], "Flag %s should be defined", expectedFlag)
	}
}

// TestUpdateCmdBasics tests basic command structure
func TestUpdateCmdBasics(t *testing.T) {
	require.NotNil(t, UpdateCmd)
	require.Equal(t, "update", UpdateCmd.Name)
	require.Equal(t, "Update an existing deal template", UpdateCmd.Usage)
	require.Equal(t, "Deal Template Management", UpdateCmd.Category)
	require.Equal(t, "<template_id_or_name>", UpdateCmd.ArgsUsage)
	require.NotEmpty(t, UpdateCmd.Description)
	require.NotNil(t, UpdateCmd.Action)
}
