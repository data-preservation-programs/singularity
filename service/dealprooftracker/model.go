package dealprooftracker

import (
	"time"
)


// DealProofTracking stores proof tracking info for a deal/sector
type DealProofTracking struct {
	DealID                uint64    `gorm:"primaryKey;column:deal_id"`
	Provider              string    `gorm:"column:provider;index"`
	SectorID              int64     `gorm:"column:sector_id;index"`
	SectorStartEpoch      int32     `gorm:"column:sector_start_epoch"`
	CurrentDeadlineIndex  int32     `gorm:"column:current_deadline_index"`
	PeriodStartEpoch      int32     `gorm:"column:period_start_epoch"`
	EstimatedNextProofTime time.Time `gorm:"column:estimated_next_proof_time"`
	Faults                int       `gorm:"column:faults"`
	Recoveries            int       `gorm:"column:recoveries"`
	LastUpdatedAt         time.Time `gorm:"column:last_updated_at;autoUpdateTime"`
}

// TableName sets the table name for GORM
func (DealProofTracking) TableName() string {
	return "deal_proof_trackings"
}

// AutoMigrateDealProofTracking migrates the table
func AutoMigrateDealProofTracking(db interface{ AutoMigrate(...interface{}) error }) error {
	return db.AutoMigrate(&DealProofTracking{})
}
