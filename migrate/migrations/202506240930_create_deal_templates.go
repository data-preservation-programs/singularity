package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// _202506240930_create_deal_templates creates the deal_templates table
// with embedded deal config fields prefixed with "template_"
func _202506240930_create_deal_templates() *gormigrate.Migration {
	type DealTemplate struct {
		ID          uint   `gorm:"primaryKey"`
		Name        string `gorm:"unique"`
		Description string
		CreatedAt   time.Time
		UpdatedAt   time.Time

		// DealConfig fields (embedded with prefix)
		AutoCreateDeals     bool    `gorm:"column:template_auto_create_deals;default:false"`
		DealProvider        string  `gorm:"column:template_deal_provider;type:varchar(255)"`
		DealTemplate        string  `gorm:"column:template_deal_template;type:varchar(255)"`
		DealVerified        bool    `gorm:"column:template_deal_verified;default:false"`
		DealKeepUnsealed    bool    `gorm:"column:template_deal_keep_unsealed;default:false"`
		DealAnnounceToIpni  bool    `gorm:"column:template_deal_announce_to_ipni;default:true"`
		DealDuration        int64   `gorm:"column:template_deal_duration;default:15552000000000000"` // ~180 days
		DealStartDelay      int64   `gorm:"column:template_deal_start_delay;default:86400000000000"` // ~1 day
		DealPricePerDeal    float64 `gorm:"column:template_deal_price_per_deal;default:0"`
		DealPricePerGb      float64 `gorm:"column:template_deal_price_per_gb;default:0"`
		DealPricePerGbEpoch float64 `gorm:"column:template_deal_price_per_gb_epoch;default:0"`
		DealHTTPHeaders     string  `gorm:"column:template_deal_http_headers;type:text"`
		DealURLTemplate     string  `gorm:"column:template_deal_url_template;type:text"`
	}

	return &gormigrate.Migration{
		ID: "202506240930",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().AutoMigrate(&DealTemplate{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("deal_templates")
		},
	}
}
