package database

import (
	"strings"

	"gorm.io/gorm"
)

// PrintDeadlockInfo prints detailed deadlock information from MySQL/MariaDB InnoDB status.
// This should be called when a deadlock error is detected to help diagnose the issue.
// Returns the deadlock information as a string, or empty string if not available.
func PrintDeadlockInfo(db *gorm.DB) string {
	if db.Dialector.Name() != "mysql" {
		return ""
	}

	// Get InnoDB status
	var results []map[string]interface{}
	err := db.Raw("SHOW ENGINE INNODB STATUS").Scan(&results).Error
	if err != nil || len(results) == 0 {
		return ""
	}

	// Extract status from result
	status, ok := results[0]["Status"].(string)
	if !ok {
		return ""
	}

	// Extract just the deadlock section
	if idx := strings.Index(status, "LATEST DETECTED DEADLOCK"); idx >= 0 {
		endIdx := strings.Index(status[idx:], "--------\nTRANSACTIONS")
		if endIdx > 0 {
			return status[idx : idx+endIdx]
		}
		// If no TRANSACTIONS section found, just return everything after deadlock
		return status[idx:]
	}

	return ""
}

// EnableDeadlockLogging enables logging of all deadlocks to the MySQL error log.
// By default, MySQL/MariaDB only logs the most recent deadlock.
// This setting persists until the server is restarted.
func EnableDeadlockLogging(db *gorm.DB) error {
	return db.Exec("SET GLOBAL innodb_print_all_deadlocks = ON").Error
}

// CheckDeadlockLoggingEnabled checks if innodb_print_all_deadlocks is enabled.
func CheckDeadlockLoggingEnabled(db *gorm.DB) (bool, error) {
	var value string
	err := db.Raw("SHOW VARIABLES LIKE 'innodb_print_all_deadlocks'").Scan(&value).Error
	if err != nil {
		return false, err
	}
	return strings.ToLower(value) == "on", nil
}

// GetDataLockWaits returns current lock wait information from performance_schema.
// This requires MySQL 8.0.30+ or MariaDB 10.5+.
func GetDataLockWaits(db *gorm.DB) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := db.Raw("SELECT * FROM performance_schema.data_lock_waits").Scan(&results).Error
	return results, err
}

// GetLockWaitTransactions returns transactions currently waiting for locks.
// This requires MySQL 8.0.30+ or MariaDB 10.5+.
func GetLockWaitTransactions(db *gorm.DB) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := db.Raw(`
		SELECT * FROM performance_schema.events_transactions_current
		WHERE STATE = 'ACTIVE'
		AND AUTOCOMMIT = 'NO'
	`).Scan(&results).Error
	return results, err
}
