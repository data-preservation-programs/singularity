package model

// private key stored in external keystore, can be linked to on-chain actor
// wallets can exist before actors are created on-chain
type Wallet struct {
	ID uint `gorm:"primaryKey" json:"id"`

	KeyPath  string `gorm:"uniqueIndex;not null" json:"keyPath"` // absolute path to key file
	KeyStore string `gorm:"default:'local';not null" json:"keyStore"` // local, yubikey, aws-kms, etc
	Address  string `gorm:"index;not null" json:"address"`       // filecoin address (f1.../f3...)
	Name     string `json:"name,omitempty"`                      // optional label

	ActorID *string `gorm:"index;size:15" json:"actorId,omitempty"` // nullable, links to on-chain actor f0...

	Actor *Actor `gorm:"foreignKey:ActorID;references:ID;constraint:OnDelete:SET NULL" json:"actor,omitempty" swaggerignore:"true" table:"expand"`
}

// GORM will rename "wallet_keys" table to "wallets" on AutoMigrate
func (Wallet) TableName() string {
	return "wallets"
}
