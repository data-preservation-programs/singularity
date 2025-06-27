package migrations

import (
	"fmt"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// Create migration for initial database schema
func _202505010840WalletActorID() *gormigrate.Migration {
	// Table names
	const WALLET_TABLE = "wallets"
	const DEAL_TABLE = "deals"

	// Temporary struct for old Wallet schema
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

	type WalletID uint

	// Temporary struct for new Wallet schema
	type NewWallet struct {
		ID               WalletID   `gorm:"primaryKey"           json:"id"`
		ActorID          string     `gorm:"index;size:15"        json:"actorId"`                             // ActorID is the short ID of the wallet
		ActorName        string     `json:"actorName"`                                                       // ActorName is readable label for the wallet
		Address          string     `gorm:"uniqueIndex;size:86"  json:"address"`                             // Address is the Filecoin full address of the wallet
		Balance          float64    `json:"balance"`                                                         // Balance is in Fil cached from chain
		BalancePlus      float64    `json:"balancePlus"`                                                     // BalancePlus is in Fil+ cached from chain
		BalanceUpdatedAt *time.Time `json:"balanceUpdatedAt"     table:"verbose;format:2006-01-02 15:04:05"` // BalanceUpdatedAt is a timestamp when balance info was last pulled from chain
		ContactInfo      string     `json:"contactInfo"`                                                     // ContactInfo is optional email for SP wallets
		Location         string     `json:"location"`                                                        // Location is optional region, country for SP wallets
		PrivateKey       string     `json:"privateKey,omitempty" table:"-"`                                  // PrivateKey is the private key of the wallet
		WalletType       WalletType `gorm:"default:'UserWallet'" json:"walletType"`
	}

	type NewDeal struct {
		ID            uint64    `gorm:"column:id"`
		ClientActorID string    `json:"clientActorId"`
		ClientID      *WalletID `gorm:"index:idx_pending"                                json:"clientId"`
		Wallet        *Wallet   `gorm:"foreignKey:ClientID;constraint:OnDelete:SET NULL" json:"wallet,omitempty" swaggerignore:"true" table:"expand"`
	}
	type OldDeal struct {
		ID       uint64  `gorm:"column:id"`
		ClientID string  `gorm:"index:idx_pending"                                json:"clientId"`
		Wallet   *Wallet `gorm:"foreignKey:ClientID;constraint:OnDelete:SET NULL" json:"wallet,omitempty" swaggerignore:"true" table:"expand"`
	}

	return &gormigrate.Migration{
		ID: "202505010840",
		Migrate: func(tx *gorm.DB) error {
			// Create new table
			if err := tx.Migrator().AutoMigrate(&NewWallet{}); err != nil {
				return fmt.Errorf("failed to create new wallets table: %w", err)
			}

			// Copy data from old to new table
			var oldWallets []OldWallet
			if err := tx.Table(WALLET_TABLE).Find(&oldWallets).Error; err != nil {
				return err
			}
			// Create map to store old ID => new ID of wallet
			idMap := make(map[string]WalletID)
			for _, oldWallet := range oldWallets {
				newWallet := NewWallet{
					ActorID:    oldWallet.ID,
					Address:    oldWallet.Address,
					PrivateKey: oldWallet.PrivateKey,
					WalletType: UserWallet,
				}
				if err := tx.Create(&newWallet).Error; err != nil {
					return err
				}
				idMap[newWallet.ActorID] = newWallet.ID
			}

			// Modify Deals table to replace ActorID FK with new ID column
			// Drop old FK constraint since Wallet ID type changed
			if err := tx.Migrator().DropConstraint(DEAL_TABLE, "fk_deals_wallet"); err != nil {
				// constraint might not exist or have different name, so continue on
				fmt.Printf("Warning: could not drop foreign key constraint: %v\n", err)
			}
			// Rename old column to make it clear it's not the FK
			if err := tx.Migrator().RenameColumn(DEAL_TABLE, "client_id", "client_actor_id"); err != nil {
				return fmt.Errorf("failed to rename ClientID to ClientActorID: %w", err)
			}
			// Add new column for updated type
			if err := tx.Table(DEAL_TABLE).Migrator().AddColumn(&NewDeal{}, "ClientID"); err != nil {
				return fmt.Errorf("failed to create new client_id column: %w", err)
			}
			// Update data using ID map
			var dealRows []NewDeal
			if err := tx.Table(DEAL_TABLE).Select("id, client_actor_id, client_id").Find(&dealRows).Error; err != nil {
				return fmt.Errorf("failed to fetch deal rows: %w", err)
			}
			for _, deal := range dealRows {
				if err := tx.Table(DEAL_TABLE).Where("id = ?", deal.ID).Update("client_id", idMap[deal.ClientActorID]).Error; err != nil {
					return fmt.Errorf("failed to update deal %d with new ClientID: %w", deal.ID, err)
				}
			}

			// Add new FK constraint on deals table
			if err := tx.Table(DEAL_TABLE).Migrator().CreateConstraint(&NewDeal{}, "Wallet"); err != nil {
				return fmt.Errorf("failed to add foreign key constraint: %w", err)
			}

			// Drop old wallets table and rename new wallets table
			if err := tx.Migrator().DropTable(WALLET_TABLE); err != nil {
				return err
			}
			return tx.Migrator().RenameTable(&NewWallet{}, WALLET_TABLE)
		},
		Rollback: func(tx *gorm.DB) error {
			// Create old table
			err := tx.Migrator().CreateTable(&OldWallet{})
			if err != nil {
				return err
			}

			// Copy data from new to old table
			var newWallets []NewWallet
			if err := tx.Table(WALLET_TABLE).Find(&newWallets).Error; err != nil {
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

			// Modify Deal table back to original FK
			// Drop old FK constraint since Wallet ID type changed
			if err := tx.Migrator().DropConstraint(DEAL_TABLE, "fk_deals_wallet"); err != nil {
				// constraint might not exist or have different name, so continue on
				fmt.Printf("Warning: could not drop foreign key constraint: %v\n", err)
			}
			// Drop new column
			if err := tx.Table(DEAL_TABLE).Migrator().DropColumn(&NewDeal{}, "ClientID"); err != nil {
				return fmt.Errorf("failed to drop ClientID column: %w", err)
			}
			// Rename old column back to FK
			if err := tx.Migrator().RenameColumn(DEAL_TABLE, "client_actor_id", "client_id"); err != nil {
				return fmt.Errorf("failed to rename ClientID to ClientActorID: %w", err)
			}
			// Add original constraint back
			if err := tx.Table(DEAL_TABLE).Migrator().CreateConstraint(&OldDeal{}, "Wallet"); err != nil {
				return fmt.Errorf("failed to add foreign key constraint: %w", err)
			}

			// Drop new table and rename old table
			if err := tx.Migrator().DropTable(WALLET_TABLE); err != nil {
				return err
			}
			return tx.Migrator().RenameTable(&OldWallet{}, WALLET_TABLE)
		},
	}
}
