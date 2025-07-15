package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func _202507091100_add_tracked_wallet_type() *gormigrate.Migration {
	type WalletType string
	const (
		UserWallet    WalletType = "UserWallet"
		SPWallet      WalletType = "SPWallet"
		TrackedWallet WalletType = "TrackedWallet"
	)

	type Wallet struct {
		ID               uint       `gorm:"primaryKey"           json:"id"`
		ActorID          string     `gorm:"index;size:15"        json:"actorId"`
		ActorName        string     `json:"actorName"`
		Address          string     `gorm:"uniqueIndex;size:86"  json:"address"`
		Balance          float64    `json:"balance"`
		BalancePlus      float64    `json:"balancePlus"`
		BalanceUpdatedAt *time.Time `json:"balanceUpdatedAt"`
		ContactInfo      string     `json:"contactInfo"`
		Location         string     `json:"location"`
		PrivateKey       string     `json:"privateKey,omitempty" table:"-"`
		WalletType       WalletType `gorm:"default:'UserWallet'" json:"walletType"`
	}

	return &gormigrate.Migration{
		ID: "202507091100",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&Wallet{})
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	}
}
