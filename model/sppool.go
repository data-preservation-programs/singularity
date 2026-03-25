package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/cockroachdb/errors"
)

var ErrInvalidReplicationPolicyEntry = errors.New("invalid replication policy entry in the database")

// ReplicationPolicy maps deal types to the number of schedules of that type
// to create per provider per preparation. Example: {"market": 1, "pdp": 1}.
type ReplicationPolicy map[DealType]int

func (rp ReplicationPolicy) Value() (driver.Value, error) {
	return json.Marshal(rp)
}

func (rp *ReplicationPolicy) Scan(src any) error {
	if src == nil {
		*rp = nil
		return nil
	}

	source, ok := src.([]byte)
	if !ok {
		return ErrInvalidReplicationPolicyEntry
	}

	return json.Unmarshal(source, rp)
}

type SPPoolID uint32

type SPPoolState string

const (
	SPPoolActive SPPoolState = "active"
	SPPoolPaused SPPoolState = "paused"
)

// SPPool groups a set of preparations and storage providers with a replication
// policy, automatically generating and managing Schedules to match the desired state.
type SPPool struct {
	ID        SPPoolID    `gorm:"primaryKey"       json:"id"`
	CreatedAt time.Time   `json:"createdAt"        table:"verbose;format:2006-01-02 15:04:05"`
	UpdatedAt time.Time   `json:"updatedAt"        table:"verbose;format:2006-01-02 15:04:05"`
	Name      string      `gorm:"uniqueIndex"      json:"name"`
	State     SPPoolState `gorm:"default:'active'" json:"state"`
	Notes     string      `json:"notes"            table:"verbose"`

	// Default deal parameters inherited by all generated schedules.
	URLTemplate           string        `json:"urlTemplate"           table:"verbose"`
	HTTPHeaders           ConfigMap     `gorm:"type:JSON"             json:"httpHeaders"           table:"verbose"`
	PricePerGBEpoch       float64       `json:"pricePerGbEpoch"       table:"verbose"`
	PricePerGB            float64       `json:"pricePerGb"            table:"verbose"`
	PricePerDeal          float64       `json:"pricePerDeal"          table:"verbose"`
	Verified              bool          `json:"verified"`
	KeepUnsealed          bool          `json:"keepUnsealed"          table:"verbose"`
	AnnounceToIPNI        bool          `gorm:"column:announce_to_ipni" json:"announceToIpni"     table:"verbose"`
	StartDelay            time.Duration `json:"startDelay"            swaggertype:"primitive,integer"`
	Duration              time.Duration `json:"duration"              swaggertype:"primitive,integer"`
	ScheduleCron          string        `json:"scheduleCron"          table:"verbose"`
	ScheduleCronPerpetual bool          `json:"scheduleCronPerpetual" table:"verbose"`
	ScheduleDealNumber    int           `json:"scheduleDealNumber"    table:"verbose"`
	ScheduleDealSize      int64         `json:"scheduleDealSize"      table:"verbose"`
	MaxPendingDealNumber  int           `json:"maxPendingDealNumber"  table:"verbose"`
	MaxPendingDealSize    int64         `json:"maxPendingDealSize"    table:"verbose"`
	Force                 bool          `json:"force"                 table:"verbose"`

	// Associations
	Providers    []SPPoolProvider    `gorm:"foreignKey:PoolID;constraint:OnDelete:CASCADE" json:"providers,omitempty"    swaggerignore:"true"`
	Preparations []SPPoolPreparation `gorm:"foreignKey:PoolID;constraint:OnDelete:CASCADE" json:"preparations,omitempty" swaggerignore:"true"`
}

func (SPPool) TableName() string {
	return "sp_pools"
}

type SPPoolProviderID uint32

// SPPoolProvider links a storage provider to a pool with a per-provider replication policy.
type SPPoolProvider struct {
	ID       SPPoolProviderID `gorm:"primaryKey"                     json:"id"`
	PoolID   SPPoolID         `gorm:"uniqueIndex:idx_pool_provider"  json:"poolId"`
	Pool     *SPPool          `gorm:"foreignKey:PoolID;constraint:OnDelete:CASCADE" json:"pool,omitempty" swaggerignore:"true"`
	Provider string           `gorm:"uniqueIndex:idx_pool_provider"  json:"provider"`

	// Policy: how many schedules of each deal type per preparation.
	// Example: {"market": 1, "pdp": 1} means 1 market + 1 PDP schedule per preparation.
	Policy ReplicationPolicy `gorm:"type:JSON" json:"policy"`

	// Optional per-provider URL template override (nil = inherit from pool).
	URLTemplate *string `json:"urlTemplate,omitempty"`
}

func (SPPoolProvider) TableName() string {
	return "sp_pool_providers"
}

type SPPoolPreparationID uint32

// SPPoolPreparation links a preparation to a pool.
type SPPoolPreparation struct {
	ID            SPPoolPreparationID `gorm:"primaryKey"                         json:"id"`
	PoolID        SPPoolID            `gorm:"uniqueIndex:idx_pool_preparation"   json:"poolId"`
	Pool          *SPPool             `gorm:"foreignKey:PoolID;constraint:OnDelete:CASCADE"           json:"pool,omitempty"        swaggerignore:"true"`
	PreparationID PreparationID       `gorm:"uniqueIndex:idx_pool_preparation"   json:"preparationId"`
	Preparation   *Preparation        `gorm:"foreignKey:PreparationID;constraint:OnDelete:CASCADE"    json:"preparation,omitempty" swaggerignore:"true"`
}

func (SPPoolPreparation) TableName() string {
	return "sp_pool_preparations"
}
