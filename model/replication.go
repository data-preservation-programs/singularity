package model

import (
	"fmt"
	"strconv"
	"time"
)

type DealState string

type ScheduleState string

const (
	DealProposed        DealState = "proposed"
	DealPublished       DealState = "published"
	DealActive          DealState = "active"
	DealExpired         DealState = "expired"
	DealProposalExpired DealState = "proposal_expired"
	DealRejected        DealState = "rejected"
	DealSlashed         DealState = "slashed"
	DealErrored         DealState = "error"
)

var DealStateStrings = []string{
	string(DealProposed),
	string(DealPublished),
	string(DealActive),
	string(DealExpired),
	string(DealProposalExpired),
	string(DealRejected),
	string(DealSlashed),
	string(DealErrored),
}

var DealStates = []DealState{
	DealProposed,
	DealPublished,
	DealActive,
	DealExpired,
	DealProposalExpired,
	DealRejected,
	DealSlashed,
	DealErrored,
}

const (
	ScheduleActive    ScheduleState = "active"
	SchedulePaused    ScheduleState = "paused"
	ScheduleError     ScheduleState = "error"
	ScheduleCompleted ScheduleState = "completed"
)

var ScheduleStates = []ScheduleState{
	ScheduleActive,
	SchedulePaused,
	ScheduleError,
	ScheduleCompleted,
}

var ScheduleStateStrings = []string{
	string(ScheduleActive),
	string(SchedulePaused),
	string(ScheduleError),
	string(ScheduleCompleted),
}

func StoragePricePerEpochToPricePerDeal(price string, dealSize int64, durationEpoch int32) float64 {
	pricePerEpoch, _ := strconv.ParseFloat(price, 64)
	return pricePerEpoch / 1e18 / (float64(dealSize) / float64(1<<35)) * float64(durationEpoch)
}

type DealID uint64

// Deal is the deal model for all deals made by deal pusher or tracked by the tracker.
// The index on PieceCID is used to track replication of the same piece CID.
// The index on State and ClientID is used to calculate number and size of pending deals.
type Deal struct {
	ID               DealID     `gorm:"primaryKey"                      json:"id"                                  table:"verbose"`
	CreatedAt        time.Time  `json:"createdAt"                       table:"verbose;format:2006-01-02 15:04:05"`
	UpdatedAt        time.Time  `json:"updatedAt"                       table:"verbose;format:2006-01-02 15:04:05"`
	LastVerifiedAt   *time.Time `json:"lastVerifiedAt"                  table:"verbose;format:2006-01-02 15:04:05"` // LastVerifiedAt is the last time the deal was verified as active by the tracker
	DealID           *uint64    `gorm:"unique"                          json:"dealId"`
	State            DealState  `gorm:"index:idx_pending"               json:"state"`
	Provider         string     `json:"provider"`
	ProposalID       string     `json:"proposalId"                      table:"verbose"`
	Label            string     `json:"label"                           table:"verbose"`
	PieceCID         CID        `gorm:"column:piece_cid;index;size:255" json:"pieceCid"                            swaggertype:"string"`
	PieceSize        int64      `json:"pieceSize"`
	StartEpoch       int32      `json:"startEpoch"`
	EndEpoch         int32      `json:"endEpoch"                        table:"verbose"`
	SectorStartEpoch int32      `json:"sectorStartEpoch"                table:"verbose"`
	Price            string     `json:"price"`
	Verified         bool       `json:"verified"`
	ErrorMessage     string     `json:"errorMessage"                    table:"verbose"`

	// Associations
	ScheduleID *ScheduleID `json:"scheduleId"                                         table:"verbose"`
	Schedule   *Schedule   `gorm:"foreignKey:ScheduleID;constraint:OnDelete:SET NULL" json:"schedule,omitempty" swaggerignore:"true" table:"expand"`
	ClientID   string      `gorm:"index:idx_pending"                                  json:"clientId"`
	Wallet     *Wallet     `gorm:"foreignKey:ClientID;constraint:OnDelete:SET NULL"   json:"wallet,omitempty"   swaggerignore:"true" table:"expand"`
}

// Key returns a mostly unique key to match deal from locally proposed deals and deals from the chain.
func (d Deal) Key() string {
	return fmt.Sprintf("%s-%s-%s-%d-%d", d.ClientID, d.Provider, d.PieceCID.String(), d.StartEpoch, d.EndEpoch)
}

type ScheduleID uint32

type Schedule struct {
	ID                    ScheduleID    `gorm:"primaryKey"                          json:"id"`
	CreatedAt             time.Time     `json:"createdAt"                           table:"verbose;format:2006-01-02 15:04:05"`
	UpdatedAt             time.Time     `json:"updatedAt"                           table:"verbose;format:2006-01-02 15:04:05"`
	URLTemplate           string        `json:"urlTemplate"                         table:"verbose"`
	HTTPHeaders           ConfigMap     `gorm:"type:JSON"                           json:"httpHeaders"                         table:"verbose"`
	Provider              string        `json:"provider"`
	PricePerGBEpoch       float64       `json:"pricePerGbEpoch"                     table:"verbose"`
	PricePerGB            float64       `json:"pricePerGb"                          table:"verbose"`
	PricePerDeal          float64       `json:"pricePerDeal"                        table:"verbose"`
	TotalDealNumber       int           `json:"totalDealNumber"                     table:"verbose"`
	TotalDealSize         int64         `json:"totalDealSize"`
	Verified              bool          `json:"verified"`
	KeepUnsealed          bool          `json:"keepUnsealed"                        table:"verbose"`
	AnnounceToIPNI        bool          `gorm:"column:announce_to_ipni"             json:"announceToIpni"                      table:"verbose"`
	StartDelay            time.Duration `json:"startDelay"                          swaggertype:"primitive,integer"`
	Duration              time.Duration `json:"duration"                            swaggertype:"primitive,integer"`
	State                 ScheduleState `json:"state"`
	ScheduleCron          string        `json:"scheduleCron"`
	ScheduleCronPerpetual bool          `json:"scheduleCronPerpetual"`
	ScheduleDealNumber    int           `json:"scheduleDealNumber"`
	ScheduleDealSize      int64         `json:"scheduleDealSize"`
	MaxPendingDealNumber  int           `json:"maxPendingDealNumber"`
	MaxPendingDealSize    int64         `json:"maxPendingDealSize"`
	Notes                 string        `json:"notes"`
	ErrorMessage          string        `json:"errorMessage"                        table:"verbose"`
	AllowedPieceCIDs      StringSlice   `gorm:"type:JSON;column:allowed_piece_cids" json:"allowedPieceCids"                    table:"verbose"`
	Force                 bool          `json:"force"`

	// Associations
	PreparationID PreparationID `json:"preparationId"`
	Preparation   *Preparation  `gorm:"foreignKey:PreparationID;constraint:OnDelete:CASCADE" json:"preparation,omitempty" swaggerignore:"true" table:"expand"`
}

type WalletType string

const (
	UserWallet WalletType = "UserWallet"
	SPWallet   WalletType = "SPWallet"
)

type Wallet struct {
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
