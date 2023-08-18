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

// Deal is the deal model for all deals made by deal pusher or tracked by the tracker.
// The index on PieceCID is used to track replication of the same piece CID.
// The index on State and ClientID is used to calculate number and size of pending deals.
type Deal struct {
	ID               uint64    `gorm:"primaryKey"                      json:"id"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	DealID           *uint64   `gorm:"unique"                          json:"dealId"`
	State            DealState `gorm:"index:idx_pending"               json:"state"`
	Provider         string    `json:"provider"`
	ProposalID       string    `json:"proposalId"`
	Label            string    `json:"label"`
	PieceCID         CID       `gorm:"column:piece_cid;index;size:255" json:"pieceCid"`
	PieceSize        int64     `json:"pieceSize"`
	StartEpoch       int32     `json:"startEpoch"`
	EndEpoch         int32     `json:"endEpoch"`
	SectorStartEpoch int32     `json:"sectorStartEpoch"`
	Price            string    `json:"price"`
	Verified         bool      `json:"verified"`
	ErrorMessage     string    `json:"errorMessage"`

	// Associations
	ScheduleID *uint32   `json:"scheduleId"`
	Schedule   *Schedule `gorm:"foreignKey:ScheduleID;constraint:OnDelete:SET NULL"    json:"schedule,omitempty"    swaggerignore:"true"`
	ClientID   string    `gorm:"index:idx_pending"                                     json:"clientId"`
	Wallet     *Wallet   `gorm:"foreignKey:ClientID;constraint:OnDelete:SET NULL"      json:"wallet,omitempty"      swaggerignore:"true"`
}

// Key returns a mostly unique key to match deal from locally proposed deals and deals from the chain.
func (d Deal) Key() string {
	return fmt.Sprintf("%s-%s-%s-%d-%d", d.ClientID, d.Provider, d.PieceCID.String(), d.StartEpoch, d.EndEpoch)
}

type Schedule struct {
	ID                   uint32        `gorm:"primaryKey"           json:"id"`
	CreatedAt            time.Time     `json:"createdAt"`
	UpdatedAt            time.Time     `json:"updatedAt"`
	URLTemplate          string        `json:"urlTemplate"`
	HTTPHeaders          StringSlice   `gorm:"type:JSON"            json:"httpHeaders"`
	Provider             string        `json:"provider"`
	PricePerGBEpoch      float64       `json:"pricePerGbEpoch"`
	PricePerGB           float64       `json:"pricePerGb"`
	PricePerDeal         float64       `json:"pricePerDeal"`
	TotalDealNumber      int           `json:"totalDealNumber"`
	TotalDealSize        int64         `json:"totalDealSize"`
	Verified             bool          `json:"verified"`
	KeepUnsealed         bool          `json:"keepUnsealed"`
	AnnounceToIPNI       bool          `json:"announceToIpni"`
	StartDelay           time.Duration `json:"startDelay"           swaggertype:"primitive,integer"`
	Duration             time.Duration `json:"duration"             swaggertype:"primitive,integer"`
	State                ScheduleState `json:"state"`
	ScheduleCron         string        `json:"scheduleCron"`
	ScheduleDealNumber   int           `json:"scheduleDealNumber"`
	ScheduleDealSize     int64         `json:"scheduleDealSize"`
	MaxPendingDealNumber int           `json:"maxPendingDealNumber"`
	MaxPendingDealSize   int64         `json:"maxPendingDealSize"`
	Notes                string        `json:"notes"`
	ErrorMessage         string        `json:"errorMessage"`
	AllowedPieceCIDs     StringSlice   `gorm:"type:JSON"            json:"allowedPieceCids"`

	// Associations
	PreparationID uint32       `json:"preparationId"`
	Preparation   *Preparation `gorm:"foreignKey:PreparationID;constraint:OnDelete:CASCADE" json:"preparation,omitempty" swaggerignore:"true"`
}

type Wallet struct {
	ID         string `gorm:"primaryKey;size:15"   json:"id"`      // ID is the short ID of the wallet
	Address    string `gorm:"index"                json:"address"` // Address is the Filecoin full address of the wallet
	PrivateKey string `json:"privateKey,omitempty"`                // PrivateKey is the private key of the wallet
}
