package model

import (
	"fmt"
	"strconv"
	"time"
)

type DealState string

func EpochToTime(epoch int32) time.Time {
	//nolint:gomnd
	return time.Unix(int64(epoch)*30+1598306400, 0)
}
func EpochToUnix(epoch int32) int32 {
	return epoch*30 + 1598306400
}

func UnixToEpoch(unix int64) int32 {
	return (int32(unix) - 1598306400) / 30
}

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

const (
	ScheduleActive    ScheduleState = "active"
	SchedulePaused    ScheduleState = "paused"
	ScheduleCompleted ScheduleState = "completed"
)

func StoragePricePerEpochToPricePerDeal(price string, dealSize int64, durationEpoch int32) float64 {
	pricePerEpoch, _ := strconv.ParseFloat(price, 64)
	return pricePerEpoch / 1e18 / (float64(dealSize) / float64(1<<35)) * float64(durationEpoch)
}

type Deal struct {
	ID               uint64    `gorm:"primaryKey"                                         json:"id"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	DealID           *uint64   `gorm:"unique"                                             json:"dealId"`
	DatasetID        *uint32   `json:"datasetId"`
	Dataset          *Dataset  `gorm:"foreignKey:DatasetID;constraint:OnDelete:SET NULL"  json:"dataset,omitempty"  swaggerignore:"true"`
	State            DealState `gorm:"index:idx_stat;index:idx_pending;size:16"           json:"state"`
	ClientID         string    `gorm:"index:idx_pending;size:16"                          json:"clientId"`
	Wallet           *Wallet   `gorm:"foreignKey:ClientID;constraint:OnDelete:SET NULL"   json:"wallet,omitempty"   swaggerignore:"true"`
	Provider         string    `gorm:"index:idx_stat;size:16"                             json:"provider"`
	ProposalID       string    `json:"proposalId"`
	Label            string    `json:"label"`
	PieceCID         string    `gorm:"column:piece_cid;index"                             json:"pieceCid"`
	PieceSize        int64     `json:"pieceSize"`
	StartEpoch       int32     `json:"startEpoch"`
	EndEpoch         int32     `json:"endEpoch"`
	SectorStartEpoch int32     `json:"sectorStartEpoch"`
	Price            string    `json:"price"`
	Verified         bool      `gorm:"index:idx_pending"                                  json:"verified"`
	ErrorMessage     string    `json:"errorMessage"`
	ScheduleID       *uint32   `json:"scheduleId"`
	Schedule         *Schedule `gorm:"foreignKey:ScheduleID;constraint:OnDelete:SET NULL" json:"schedule,omitempty" swaggerignore:"true"`
}

func (d Deal) Key() string {
	return fmt.Sprintf("%s-%s-%s-%d-%d", d.ClientID, d.Provider, d.PieceCID, d.StartEpoch, d.EndEpoch)
}

type Schedule struct {
	ID                      uint32        `gorm:"primaryKey"                                                json:"id"`
	CreatedAt               time.Time     `json:"createdAt"`
	UpdatedAt               time.Time     `json:"updatedAt"`
	DatasetID               uint32        `json:"datasetId"`
	Dataset                 *Dataset      `gorm:"foreignKey:DatasetID;constraint:OnDelete:CASCADE"          json:"dataset,omitempty"        swaggerignore:"true"`
	URLTemplate             string        `json:"urlTemplate"`
	HTTPHeaders             []string      `gorm:"type:JSON"                                                 json:"httpHeaders"`
	Provider                string        `json:"provider"`
	Price                   float64       `json:"price"`
	TotalDealNumber         int           `json:"totalDealNumber"`
	TotalDealSize           int64         `json:"totalDealSize"`
	Verified                bool          `json:"verified"`
	KeepUnsealed            bool          `json:"keepUnsealed"`
	AnnounceToIPNI          bool          `json:"announceToIpni"`
	StartDelay              time.Duration `json:"startDelay"`
	Duration                time.Duration `json:"duration"`
	State                   ScheduleState `json:"state"`
	LastProcessedTimestamp  uint64        `json:"lastProcessedTimestamp"`
	ScheduleWorkerID        *string       `gorm:"index:schedule_cleanup"                                    json:"scheduleWorkerId"`
	ScheduleWorker          *Worker       `gorm:"foreignKey:ScheduleWorkerID;constraint:OnDelete:NO ACTION" json:"scheduleWorker,omitempty" swaggerignore:"true"`
	ScheduleIntervalSeconds uint64        `json:"scheduleIntervalSeconds"`
	ScheduleDealNumber      int           `json:"scheduleDealNumber"`
	ScheduleDealSize        int64         `json:"scheduleDealSize"`
	MaxPendingDealNumber    int           `json:"maxPendingDealNumber"`
	MaxPendingDealSize      int64         `json:"maxPendingDealSize"`
	Notes                   string        `json:"notes"`
	ErrorMessage            string        `json:"errorMessage"`
	AllowedPieceCIDs        StringSlice   `json:"allowedPieceCids"`
}

func (s Schedule) Equal(other Schedule) bool {
	return s.DatasetID == other.DatasetID &&
		s.URLTemplate == other.URLTemplate &&
		s.Provider == other.Provider &&
		s.Price == other.Price &&
		s.TotalDealNumber == other.TotalDealNumber &&
		s.TotalDealSize == other.TotalDealSize &&
		s.Verified == other.Verified &&
		s.StartDelay == other.StartDelay &&
		s.Duration == other.Duration &&
		s.ScheduleIntervalSeconds == other.ScheduleIntervalSeconds &&
		s.ScheduleDealNumber == other.ScheduleDealNumber &&
		s.ScheduleDealSize == other.ScheduleDealSize &&
		s.MaxPendingDealNumber == other.MaxPendingDealNumber &&
		s.MaxPendingDealSize == other.MaxPendingDealSize &&
		s.State == other.State
}

type Wallet struct {
	ID         string `gorm:"primaryKey;size:16"   json:"id"`      // ID is the short ID of the wallet
	Address    string `gorm:"unique;size:256"      json:"address"` // Address is the Filecoin full address of the wallet
	PrivateKey string `json:"privateKey,omitempty"`                // PrivateKey is the private key of the wallet
	RemotePeer string `json:"remotePeer,omitempty"`                // RemotePeer is the remote peer ID of the wallet, for remote signing purpose
}

type WalletAssignment struct {
	ID        uint32   `gorm:"primaryKey"                                       json:"id"`
	WalletID  string   `json:"walletId"`
	Wallet    *Wallet  `gorm:"foreignKey:WalletID;constraint:OnDelete:CASCADE"  json:"wallet,omitempty"  swaggerignore:"true"`
	DatasetID uint32   `json:"datasetId"`
	Dataset   *Dataset `gorm:"foreignKey:DatasetID;constraint:OnDelete:CASCADE" json:"dataset,omitempty" swaggerignore:"true"`
}
