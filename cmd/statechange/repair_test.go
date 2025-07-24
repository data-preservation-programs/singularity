package statechange

import (
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

type RepairCmdTestSuite struct {
	testutil.TestSuite
}

func TestRepairCmd(t *testing.T) {
	suite.Run(t, new(RepairCmdTestSuite))
}

func (s *RepairCmdTestSuite) TestForceTransition_Success() {
	// Create test deal
	deal := model.Deal{
		ID:            123,
		State:         "proposed",
		Provider:      "f01234",
		ClientActorID: "f1abcdef",
	}
	s.NoError(s.DB.Create(&deal).Error)

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{RepairCmd},
	}

	// Test force transition
	err := app.Run([]string{"test", "repair", "force-transition", "123", "published", "--reason", "test transition"})
	s.NoError(err)

	// Verify deal state was updated
	var updatedDeal model.Deal
	s.NoError(s.DB.First(&updatedDeal, 123).Error)
	s.Equal(model.DealState("published"), updatedDeal.State)

	// Verify state change was recorded
	var stateChange model.DealStateChange
	s.NoError(s.DB.Where("deal_id = ?", 123).First(&stateChange).Error)
	s.Equal(model.DealID(123), stateChange.DealID)
	s.Equal(model.DealState("proposed"), stateChange.PreviousState)
	s.Equal(model.DealState("published"), stateChange.NewState)
	s.Contains(stateChange.Metadata, "test transition")
}

func (s *RepairCmdTestSuite) TestForceTransition_DryRun() {
	// Create test deal
	deal := model.Deal{
		ID:            124,
		State:         "proposed",
		Provider:      "f01234",
		ClientActorID: "f1abcdef",
	}
	s.NoError(s.DB.Create(&deal).Error)

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{RepairCmd},
	}

	// Test dry run
	err := app.Run([]string{"test", "repair", "force-transition", "124", "published", "--dry-run"})
	s.NoError(err)

	// Verify deal state was NOT updated
	var unchangedDeal model.Deal
	s.NoError(s.DB.First(&unchangedDeal, 124).Error)
	s.Equal(model.DealState("proposed"), unchangedDeal.State)

	// Verify no state change was recorded
	var count int64
	s.DB.Model(&model.DealStateChange{}).Where("deal_id = ?", 124).Count(&count)
	s.Equal(int64(0), count)
}

func (s *RepairCmdTestSuite) TestForceTransition_InvalidState() {
	// Create test deal
	deal := model.Deal{
		ID:            125,
		State:         "proposed",
		Provider:      "f01234",
		ClientActorID: "f1abcdef",
	}
	s.NoError(s.DB.Create(&deal).Error)

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{RepairCmd},
	}

	// Test invalid state
	err := app.Run([]string{"test", "repair", "force-transition", "125", "invalid-state"})
	s.Error(err)
	s.Contains(err.Error(), "invalid state")
}

func (s *RepairCmdTestSuite) TestForceTransition_MissingArgs() {
	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{RepairCmd},
	}

	// Test missing arguments
	err := app.Run([]string{"test", "repair", "force-transition", "123"})
	s.Error(err)
	s.Contains(err.Error(), "deal ID and new state are required")
}

func (s *RepairCmdTestSuite) TestForceTransition_InvalidDealID() {
	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{RepairCmd},
	}

	// Test invalid deal ID
	err := app.Run([]string{"test", "repair", "force-transition", "invalid", "published"})
	s.Error(err)
	s.Contains(err.Error(), "invalid deal ID format")
}

func (s *RepairCmdTestSuite) TestForceTransition_NonexistentDeal() {
	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{RepairCmd},
	}

	// Test nonexistent deal
	err := app.Run([]string{"test", "repair", "force-transition", "99999", "published"})
	s.Error(err)
	s.Contains(err.Error(), "failed to find deal")
}

func (s *RepairCmdTestSuite) TestResetErrorDeals_Success() {
	// Create test deals in error state
	deals := []model.Deal{
		{
			ID:            130,
			State:         "error",
			Provider:      "f01234",
			ClientActorID: "f1abcdef",
		},
		{
			ID:            131,
			State:         "error",
			Provider:      "f05678",
			ClientActorID: "f1fedcba",
		},
		{
			ID:            132,
			State:         "active", // Should not be affected
			Provider:      "f01234",
			ClientActorID: "f1abcdef",
		},
	}
	for _, deal := range deals {
		s.NoError(s.DB.Create(&deal).Error)
	}

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{RepairCmd},
	}

	// Test reset error deals
	err := app.Run([]string{"test", "repair", "reset-error-deals", "--reset-to-state", "proposed"})
	s.NoError(err)

	// Verify error deals were reset
	var resetDeals []model.Deal
	s.NoError(s.DB.Where("id IN ?", []uint64{130, 131}).Find(&resetDeals).Error)
	for _, deal := range resetDeals {
		s.Equal(model.DealState("proposed"), deal.State)
	}

	// Verify active deal was not affected
	var activeDeal model.Deal
	s.NoError(s.DB.First(&activeDeal, 132).Error)
	s.Equal(model.DealState("active"), activeDeal.State)

	// Verify state changes were recorded
	var stateChangeCount int64
	s.DB.Model(&model.DealStateChange{}).Where("deal_id IN ?", []uint64{130, 131}).Count(&stateChangeCount)
	s.Equal(int64(2), stateChangeCount)
}

func (s *RepairCmdTestSuite) TestResetErrorDeals_DryRun() {
	// Create test deal in error state
	deal := model.Deal{
		ID:            133,
		State:         "error",
		Provider:      "f01234",
		ClientActorID: "f1abcdef",
	}
	s.NoError(s.DB.Create(&deal).Error)

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{RepairCmd},
	}

	// Test dry run
	err := app.Run([]string{"test", "repair", "reset-error-deals", "--dry-run"})
	s.NoError(err)

	// Verify deal state was NOT changed
	var unchangedDeal model.Deal
	s.NoError(s.DB.First(&unchangedDeal, 133).Error)
	s.Equal(model.DealState("error"), unchangedDeal.State)
}

func (s *RepairCmdTestSuite) TestResetErrorDeals_SpecificDeals() {
	// Create test deals in error state
	deals := []model.Deal{
		{
			ID:            134,
			State:         "error",
			Provider:      "f01234",
			ClientActorID: "f1abcdef",
		},
		{
			ID:            135,
			State:         "error",
			Provider:      "f05678",
			ClientActorID: "f1fedcba",
		},
	}
	for _, deal := range deals {
		s.NoError(s.DB.Create(&deal).Error)
	}

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{RepairCmd},
	}

	// Test reset specific deal
	err := app.Run([]string{"test", "repair", "reset-error-deals", "--deal-id", "134"})
	s.NoError(err)

	// Verify only specified deal was reset
	var resetDeal model.Deal
	s.NoError(s.DB.First(&resetDeal, 134).Error)
	s.Equal(model.DealState("proposed"), resetDeal.State)

	// Verify other deal was not affected
	var untouchedDeal model.Deal
	s.NoError(s.DB.First(&untouchedDeal, 135).Error)
	s.Equal(model.DealState("error"), untouchedDeal.State)
}

func (s *RepairCmdTestSuite) TestResetErrorDeals_ByProvider() {
	// Create test deals in error state for different providers
	deals := []model.Deal{
		{
			ID:            136,
			State:         "error",
			Provider:      "f01234",
			ClientActorID: "f1abcdef",
		},
		{
			ID:            137,
			State:         "error",
			Provider:      "f05678",
			ClientActorID: "f1fedcba",
		},
	}
	for _, deal := range deals {
		s.NoError(s.DB.Create(&deal).Error)
	}

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{RepairCmd},
	}

	// Test reset by provider
	err := app.Run([]string{"test", "repair", "reset-error-deals", "--provider", "f01234"})
	s.NoError(err)

	// Verify only deals from specified provider were reset
	var resetDeal model.Deal
	s.NoError(s.DB.First(&resetDeal, 136).Error)
	s.Equal(model.DealState("proposed"), resetDeal.State)

	// Verify other provider's deal was not affected
	var untouchedDeal model.Deal
	s.NoError(s.DB.First(&untouchedDeal, 137).Error)
	s.Equal(model.DealState("error"), untouchedDeal.State)
}

func (s *RepairCmdTestSuite) TestCleanupOrphanedChanges_Success() {
	// Create a valid deal
	deal := model.Deal{
		ID:            140,
		State:         "active",
		Provider:      "f01234",
		ClientActorID: "f1abcdef",
	}
	s.NoError(s.DB.Create(&deal).Error)

	// Create state changes - one valid, one orphaned
	validChange := model.DealStateChange{
		DealID:        140,
		PreviousState: "published",
		NewState:      "active",
		Timestamp:     time.Now(),
		ProviderID:    "f01234",
		ClientAddress: "f1abcdef",
		Metadata:      "{}",
	}
	s.NoError(s.DB.Create(&validChange).Error)

	orphanedChange := model.DealStateChange{
		DealID:        99999, // Non-existent deal
		PreviousState: "proposed",
		NewState:      "error",
		Timestamp:     time.Now(),
		ProviderID:    "f05678",
		ClientAddress: "f1fedcba",
		Metadata:      "{}",
	}
	s.NoError(s.DB.Create(&orphanedChange).Error)

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{RepairCmd},
	}

	// Test cleanup
	err := app.Run([]string{"test", "repair", "cleanup-orphaned-changes"})
	s.NoError(err)

	// Verify valid change still exists
	var remainingChanges []model.DealStateChange
	s.NoError(s.DB.Find(&remainingChanges).Error)
	s.Len(remainingChanges, 1)
	s.Equal(model.DealID(140), remainingChanges[0].DealID)
}

func (s *RepairCmdTestSuite) TestCleanupOrphanedChanges_DryRun() {
	// Create orphaned state change
	orphanedChange := model.DealStateChange{
		DealID:        99998, // Non-existent deal
		PreviousState: "proposed",
		NewState:      "error",
		Timestamp:     time.Now(),
		ProviderID:    "f05678",
		ClientAddress: "f1fedcba",
		Metadata:      "{}",
	}
	s.NoError(s.DB.Create(&orphanedChange).Error)

	// Create CLI context
	app := &cli.App{
		Commands: []*cli.Command{RepairCmd},
	}

	// Test dry run
	err := app.Run([]string{"test", "repair", "cleanup-orphaned-changes", "--dry-run"})
	s.NoError(err)

	// Verify orphaned change still exists
	var count int64
	s.DB.Model(&model.DealStateChange{}).Count(&count)
	s.Equal(int64(1), count)
}

func (s *RepairCmdTestSuite) TestCleanupOrphanedChanges_NoOrphaned() {
	// Create CLI context without any orphaned changes
	app := &cli.App{
		Commands: []*cli.Command{RepairCmd},
	}

	// Test with no orphaned changes
	err := app.Run([]string{"test", "repair", "cleanup-orphaned-changes"})
	s.NoError(err)

	// Should complete without error
}