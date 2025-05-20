package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// Create migration for initial database schema
func _202505010840_wallet_actor_id() *gormigrate.Migration {
	// Table name
	const WALLET_TABLE_NAME = "wallets"

	// Temporary struct for old schema
	type OldWallet struct {
		ID         string `gorm:"primaryKey;size:15"   json:"id"`      // ID is the short ID of the wallet
		Address    string `gorm:"index"                json:"address"` // Address is the Filecoin full address of the wallet
		PrivateKey string `json:"privateKey,omitempty" table:"-"`      // PrivateKey is the private key of the wallet
	}

	type WalletType string
	const (
		UserWallet WalletType = "UserWallet"
		SPWallet   WalletType = "SPWallet"
	)

	// Temporary struct for new schema
	type NewWallet struct {
		ID          uint       `gorm:"primaryKey"    json:"id"`
		ActorID     string     `gorm:"index,size:15" json:"actorId"`   // ActorID is the short ID of the wallet
		ActorName   string     `json:"actorName"`                      // ActorName is readable label for the wallet
		Address     string     `gorm:"index"         json:"address"`   // Address is the Filecoin full address of the wallet
		Balance     float64    `json:"balance"`                        // Balance is in Fil cached from chain
		BalancePlus float64    `json:"balancePlus"`                    // BalancePlus is in Fil+ cached from chain
		ContactInfo string     `json:"contactInfo"`                    // ContactInfo is optional email for SP wallets
		Location    string     `json:"location"`                       // Location is optional region, country for SP wallets
		PrivateKey  string     `json:"privateKey,omitempty" table:"-"` // PrivateKey is the private key of the wallet
		Type        WalletType `json:"type"`                           // Type determines user or SP wallets
	}

	return &gormigrate.Migration{
		ID: "202505010840",
		Migrate: func(tx *gorm.DB) error {
			// Create new table
			err := tx.Migrator().CreateTable(&NewWallet{})
			if err != nil {
				return err
			}

			// Copy data from old to new table
			var oldWallets []OldWallet
			if err := tx.Table(WALLET_TABLE_NAME).Find(&oldWallets).Error; err != nil {
				return err
			}

			for _, oldWallet := range oldWallets {
				newWallet := NewWallet{
					ActorID:    oldWallet.ID,
					Address:    oldWallet.Address,
					PrivateKey: oldWallet.PrivateKey,
				}
				if err := tx.Create(&newWallet).Error; err != nil {
					return err
				}
			}

			// Drop old table and rename new table
			if err := tx.Migrator().DropTable(WALLET_TABLE_NAME); err != nil {
				return err
			}
			return tx.Migrator().RenameTable(&NewWallet{}, WALLET_TABLE_NAME)
		},
		Rollback: func(tx *gorm.DB) error {
			// Create old table
			err := tx.Migrator().CreateTable(&OldWallet{})
			if err != nil {
				return err
			}

			// Copy data from new to old table
			var newWallets []NewWallet
			if err := tx.Table(WALLET_TABLE_NAME).Find(&newWallets).Error; err != nil {
				return err
			}

			for _, newWallet := range newWallets {
				oldWallet := OldWallet{
					ID:         newWallet.ActorID,
					Address:    newWallet.Address,
					PrivateKey: newWallet.PrivateKey,
				}
				if err := tx.Create(&oldWallet).Error; err != nil {
					return err
				}
			}

			// Drop new table and rename old table
			if err := tx.Migrator().DropTable(WALLET_TABLE_NAME); err != nil {
				return err
			}
			return tx.Migrator().RenameTable(&OldWallet{}, WALLET_TABLE_NAME)
		},
	}
}
