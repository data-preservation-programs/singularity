package state

import (
	"fmt"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/statetracker"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

// forceTransitionAction handles the force-transition subcommand
func forceTransitionAction(c *cli.Context) error {
	if c.NArg() != 2 {
		return errors.New("deal ID and new state are required")
	}

	dealIDStr := c.Args().Get(0)
	newStateStr := c.Args().Get(1)

	dealID, err := strconv.ParseUint(dealIDStr, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid deal ID format")
	}

	newState, err := validateDealState(newStateStr)
	if err != nil {
		return err
	}

	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() { _ = closer.Close() }()

	// Get current deal information
	var deal model.Deal
	err = db.First(&deal, dealID).Error
	if err != nil {
		return errors.Wrap(err, "failed to find deal")
	}

	if c.Bool("dry-run") {
		cliutil.Print(c, map[string]interface{}{
			"message":       "DRY RUN: Would force state transition",
			"dealId":        dealIDStr,
			"currentState":  deal.State,
			"newState":      newState,
			"provider":      deal.Provider,
			"clientAddress": deal.ClientActorID,
			"reason":        c.String("reason"),
		})
		return nil
	}

	// Parse optional parameters
	epochHeight, sectorID, err := parseOptionalTransitionParams(c)
	if err != nil {
		return err
	}

	// Create state tracker and record the forced transition
	tracker := statetracker.NewStateChangeTracker(db)
	metadata := &statetracker.StateChangeMetadata{
		Reason: c.String("reason"),
		AdditionalFields: map[string]interface{}{
			"operationType": "manual_force_transition",
			"operator":      "cli",
		},
	}

	previousState := &deal.State
	err = tracker.TrackStateChangeWithDetails(
		c.Context,
		model.DealID(dealID),
		previousState,
		newState,
		epochHeight,
		sectorID,
		deal.Provider,
		deal.ClientActorID,
		metadata,
	)
	if err != nil {
		return errors.Wrap(err, "failed to record state change")
	}

	// Update the deal state in the database
	err = db.Model(&deal).Update("state", newState).Error
	if err != nil {
		return errors.Wrap(err, "failed to update deal state")
	}

	cliutil.Print(c, map[string]interface{}{
		"message":       "Deal state transition forced successfully",
		"dealId":        dealIDStr,
		"previousState": *previousState,
		"newState":      newState,
		"reason":        c.String("reason"),
	})
	return nil
}

// resetErrorDealsAction handles the reset-error-deals subcommand
func resetErrorDealsAction(c *cli.Context) error {
	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() { _ = closer.Close() }()

	resetToState := model.DealState(c.String("reset-to-state"))

	// Build query for error deals
	query, err := buildErrorDealsQuery(c, db)
	if err != nil {
		return err
	}

	// Get deals to reset
	var deals []model.Deal
	err = query.Find(&deals).Error
	if err != nil {
		return errors.Wrap(err, "failed to find error deals")
	}

	if len(deals) == 0 {
		cliutil.Print(c, map[string]interface{}{
			"message": "No error deals found matching the criteria",
		})
		return nil
	}

	if c.Bool("dry-run") {
		cliutil.Print(c, map[string]interface{}{
			"message":      "DRY RUN: Would reset the following deals",
			"dealCount":    len(deals),
			"resetToState": resetToState,
			"deals":        deals,
		})
		return nil
	}

	// Reset deals
	resetCount := performDealResets(c, db, deals, resetToState)

	cliutil.Print(c, map[string]interface{}{
		"message":         "Error deals reset successfully",
		"totalFound":      len(deals),
		"successfulReset": resetCount,
		"resetToState":    resetToState,
	})
	return nil
}

// cleanupOrphanedChangesAction handles the cleanup-orphaned-changes subcommand
func cleanupOrphanedChangesAction(c *cli.Context) error {
	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() { _ = closer.Close() }()

	// Find orphaned state changes
	orphanedChanges, err := findOrphanedStateChanges(db)
	if err != nil {
		return err
	}

	if len(orphanedChanges) == 0 {
		cliutil.Print(c, map[string]interface{}{
			"message": "No orphaned state changes found",
		})
		return nil
	}

	if c.Bool("dry-run") {
		cliutil.Print(c, map[string]interface{}{
			"message":     "DRY RUN: Would delete orphaned state changes",
			"orphanCount": len(orphanedChanges),
			"orphanedIds": extractStateChangeIds(orphanedChanges),
		})
		return nil
	}

	// Delete orphaned state changes
	err = deleteOrphanedChanges(db, orphanedChanges)
	if err != nil {
		return err
	}

	cliutil.Print(c, map[string]interface{}{
		"message":      "Orphaned state changes cleaned up successfully",
		"deletedCount": len(orphanedChanges),
	})
	return nil
}

// Helper functions

// validateDealState validates that the provided state is valid
func validateDealState(stateStr string) (model.DealState, error) {
	newState := model.DealState(stateStr)
	validStates := []model.DealState{
		"proposed", "published", "active", "expired",
		"proposal_expired", "rejected", "slashed", "error",
	}

	for _, validState := range validStates {
		if newState == validState {
			return newState, nil
		}
	}
	return "", errors.Errorf("invalid state: %s. Valid states: %v", stateStr, validStates)
}

// parseOptionalTransitionParams parses optional epoch and sector ID parameters
func parseOptionalTransitionParams(c *cli.Context) (*int32, *string, error) {
	var epochHeight *int32
	if epochStr := c.String("epoch"); epochStr != "" {
		epoch, err := strconv.ParseInt(epochStr, 10, 32)
		if err != nil {
			return nil, nil, errors.Wrap(err, "invalid epoch format")
		}
		epochInt32 := int32(epoch)
		epochHeight = &epochInt32
	}

	var sectorID *string
	if sector := c.String("sector-id"); sector != "" {
		sectorID = &sector
	}

	return epochHeight, sectorID, nil
}

// buildErrorDealsQuery builds the query for finding error deals
func buildErrorDealsQuery(c *cli.Context, db *gorm.DB) (*gorm.DB, error) {
	query := db.Where("state = ?", "error")

	// Filter by specific deal IDs if provided
	dealIDs := c.StringSlice("deal-id")
	if len(dealIDs) > 0 {
		var dealIDValues []uint64
		for _, idStr := range dealIDs {
			id, err := strconv.ParseUint(idStr, 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "invalid deal ID: %s", idStr)
			}
			dealIDValues = append(dealIDValues, id)
		}
		query = query.Where("id IN ?", dealIDValues)
	}

	// Filter by provider if specified
	if provider := c.String("provider"); provider != "" {
		query = query.Where("provider = ?", provider)
	}

	// Apply limit
	query = query.Limit(c.Int("limit"))

	return query, nil
}

// performDealResets resets the provided deals to the specified state
func performDealResets(c *cli.Context, db *gorm.DB, deals []model.Deal, resetToState model.DealState) int {
	tracker := statetracker.NewStateChangeTracker(db)
	resetCount := 0

	for _, deal := range deals {
		metadata := &statetracker.StateChangeMetadata{
			Reason: "Manual error state reset",
			AdditionalFields: map[string]interface{}{
				"operationType": "error_state_reset",
				"operator":      "cli",
			},
		}

		previousState := &deal.State
		err := tracker.TrackStateChangeWithDetails(
			c.Context,
			deal.ID,
			previousState,
			resetToState,
			nil,
			nil,
			deal.Provider,
			deal.ClientActorID,
			metadata,
		)
		if err != nil {
			cliutil.Print(c, map[string]interface{}{
				"warning": fmt.Sprintf("Failed to track state change for deal %d: %v", deal.ID, err),
			})
			continue
		}

		// Update the deal state
		err = db.Model(&deal).Update("state", resetToState).Error
		if err != nil {
			cliutil.Print(c, map[string]interface{}{
				"warning": fmt.Sprintf("Failed to update deal %d state: %v", deal.ID, err),
			})
			continue
		}

		resetCount++
	}

	return resetCount
}

// findOrphanedStateChanges finds state changes without corresponding deals
func findOrphanedStateChanges(db *gorm.DB) ([]model.DealStateChange, error) {
	var orphanedChanges []model.DealStateChange
	err := db.Table("deal_state_changes").
		Select("deal_state_changes.*").
		Joins("LEFT JOIN deals ON deals.id = deal_state_changes.deal_id").
		Where("deals.id IS NULL").
		Find(&orphanedChanges).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to find orphaned state changes")
	}
	return orphanedChanges, nil
}

// deleteOrphanedChanges deletes the provided orphaned state changes
func deleteOrphanedChanges(db *gorm.DB, orphanedChanges []model.DealStateChange) error {
	var orphanedIds []uint64
	for _, change := range orphanedChanges {
		orphanedIds = append(orphanedIds, change.ID)
	}

	err := db.Where("id IN ?", orphanedIds).Delete(&model.DealStateChange{}).Error
	if err != nil {
		return errors.Wrap(err, "failed to delete orphaned state changes")
	}
	return nil
}

// extractStateChangeIds extracts state change IDs from a slice of state changes
func extractStateChangeIds(changes []model.DealStateChange) []uint64 {
	ids := make([]uint64, len(changes))
	for i, change := range changes {
		ids[i] = change.ID
	}
	return ids
}

var RepairCmd = &cli.Command{
	Name:  "repair",
	Usage: "Manual recovery and repair commands for deal state management",
	Description: `Provides manual recovery and repair capabilities for deal state management:
- Force state transitions for stuck deals
- Reset deal states to allow retry
- Repair corrupted state transitions
- Bulk operations for multiple deals`,
	Subcommands: []*cli.Command{
		{
			Name:      "force-transition",
			Usage:     "Force a state transition for a specific deal",
			ArgsUsage: "<deal-id> <new-state>",
			Description: `Force a deal to transition to a new state. Use with caution!
Valid states: proposed, published, active, expired, proposal_expired, rejected, slashed, error
Note: deal-id refers to the internal database ID, not a content CID.`,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "reason",
					Usage: "Reason for the forced state transition",
					Value: "Manual repair operation",
				},
				&cli.StringFlag{
					Name:  "epoch",
					Usage: "Filecoin epoch height for the state change",
				},
				&cli.StringFlag{
					Name:  "sector-id",
					Usage: "Storage provider sector ID",
				},
				&cli.BoolFlag{
					Name:  "dry-run",
					Usage: "Show what would be done without making changes",
				},
			},
			Action: forceTransitionAction,
		},
		{
			Name:  "reset-error-deals",
			Usage: "Reset deals in error state to allow retry",
			Description: `Reset deals that are in error state back to their previous valid state.
This allows the system to retry operations that may have failed temporarily.`,
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:  "deal-id",
					Usage: "Specific deal IDs to reset (internal database IDs, not CIDs, can be specified multiple times)",
				},
				&cli.StringFlag{
					Name:  "provider",
					Usage: "Reset error deals for a specific provider",
				},
				&cli.StringFlag{
					Name:  "reset-to-state",
					Usage: "State to reset deals to (default: proposed)",
					Value: "proposed",
				},
				&cli.IntFlag{
					Name:  "limit",
					Usage: "Maximum number of deals to reset",
					Value: 100,
				},
				&cli.BoolFlag{
					Name:  "dry-run",
					Usage: "Show what would be done without making changes",
				},
			},
			Action: resetErrorDealsAction,
		},
		{
			Name:  "cleanup-orphaned-changes",
			Usage: "Clean up orphaned state changes without corresponding deals",
			Description: `Remove state change records that reference deals that no longer exist.
This helps maintain database consistency and reduce storage usage.`,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "dry-run",
					Usage: "Show what would be deleted without making changes",
				},
			},
			Action: cleanupOrphanedChangesAction,
		},
	},
}
