package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// _202506240816_create_deal_templates creates the deal_templates table
// with embedded deal config fields prefixed with "template_"
func _202506240816_create_deal_templates() *gormigrate.Migration {
	type DealTemplate struct {
		ID          uint   `gorm:"primaryKey"`
		Name        string `gorm:"unique"`
		Description string
		CreatedAt   time.Time
		UpdatedAt   time.Time

		// DealConfig fields (embedded with prefix "template_")
		TemplateAutoCreateDeals     bool    `gorm:"column:template_auto_create_deals;default:false"`
		TemplateDealProvider        string  `gorm:"column:template_deal_provider;type:varchar(255)"`
		TemplateDealTemplate        string  `gorm:"column:template_deal_template;type:varchar(255)"`
		TemplateDealVerified        bool    `gorm:"column:template_deal_verified;default:false"`
		TemplateDealKeepUnsealed    bool    `gorm:"column:template_deal_keep_unsealed;default:false"`
		TemplateDealAnnounceToIpni  bool    `gorm:"column:template_deal_announce_to_ipni;default:true"`
		TemplateDealDuration        int64   `gorm:"column:template_deal_duration;default:15552000000000000"` // ~180 days in nanoseconds
		TemplateDealStartDelay      int64   `gorm:"column:template_deal_start_delay;default:86400000000000"` // ~1 day in nanoseconds
		TemplateDealPricePerDeal    float64 `gorm:"column:template_deal_price_per_deal;default:0"`
		TemplateDealPricePerGb      float64 `gorm:"column:template_deal_price_per_gb;default:0"`
		TemplateDealPricePerGbEpoch float64 `gorm:"column:template_deal_price_per_gb_epoch;default:0"`
		TemplateDealHTTPHeaders     string  `gorm:"column:template_deal_http_headers;type:text"`
		TemplateDealURLTemplate     string  `gorm:"column:template_deal_url_template;type:text"`
	}

	return &gormigrate.Migration{
		ID: "202506240816",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().CreateTable(&DealTemplate{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("deal_templates")
		},
	}
}
