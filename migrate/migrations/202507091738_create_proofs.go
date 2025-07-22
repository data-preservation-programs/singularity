package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// ProofType represents the type of proof
type ProofType string

const (
	ProofOfReplication ProofType = "replication"
	ProofOfSpacetime   ProofType = "spacetime"
)

// Proof represents a proof record at the time of migration
type Proof struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// Link to existing deal
	DealID *uint64 `gorm:"index" json:"dealId"`

	// Proof details
	ProofType ProofType `gorm:"index" json:"proofType"`
	MessageID string    `gorm:"index" json:"messageId"` // Message CID
	BlockCID  string    `gorm:"index" json:"blockCid"`  // Block CID where proof was included
	Height    int64     `gorm:"index" json:"height"`    // Block height
	Method    string    `json:"method"`                 // Proof method/algorithm
	Verified  bool      `json:"verified"`               // Whether proof was successfully verified

	// Metadata
	SectorID *uint64 `json:"sectorId"`
	Provider string  `gorm:"index" json:"provider"` // Storage Provider ID
	ErrorMsg string  `json:"errorMessage,omitempty"`
}

// Create migration for proofs table
func _202507091738_create_proofs() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202507091738",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&Proof{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&Proof{})
		},
	}
}
