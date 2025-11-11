package testutil

import (
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"gorm.io/gorm"
)

// EnableDeadlockLogging enables comprehensive deadlock logging for MySQL/MariaDB tests.
// This should be called early in tests that may encounter deadlocks.
// It will:
// - Enable innodb_print_all_deadlocks (logs all deadlocks to error log, not just the last one)
// - Log the current state of deadlock logging
//
// Note: innodb_print_all_deadlocks requires SUPER privilege and persists until server restart.
func EnableDeadlockLogging(t *testing.T, db *gorm.DB) {
	// Try to enable it (may fail if already enabled or insufficient privileges)
	err := database.EnableDeadlockLogging(db)
	if err != nil {
		t.Logf("Note: Could not enable innodb_print_all_deadlocks: %v (may not have SUPER privilege)", err)
	}

	// Check if it's enabled
	enabled, err := database.CheckDeadlockLoggingEnabled(db)
	if err != nil {
		t.Logf("Note: Could not check innodb_print_all_deadlocks status: %v", err)
		return
	}

	if enabled {
		t.Logf("Deadlock logging enabled: all deadlocks will be logged to MySQL error log")
	} else {
		t.Logf("Deadlock logging not enabled: only the most recent deadlock will be available")
	}
}
