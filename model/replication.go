package model

import (
	"fmt"
	"strconv"
	"time"
)

type DealState string

func EpochToTime(epoch int32) time.Time {
	//nolint:gomnd
	return time.Unix(int64(epoch*30+1598306400), 0)
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
	ID           uint64 `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DealID       *uint64 `gorm:"uniqueIndex"`
	DatasetID    *uint32
	Dataset      *Dataset  `gorm:"foreignKey:DatasetID;constraint:OnDelete:SET NULL"`
	State        DealState `gorm:"index:idx_stat"`
	ClientID     string
	Wallet       *Wallet `gorm:"foreignKey:ClientID;constraint:OnDelete:NO ACTION"`
	Provider     string  `gorm:"index:idx_stat"`
	ProposalID   string
	Label        string
	PieceCID     string `gorm:"column:piece_cid;index"`
	PieceSize    int64
	Start        time.Time
	Duration     time.Duration
	End          time.Time
	SectorStart  time.Time `gorm:"index:idx_stat"`
	Price        float64
	Verified     bool
	ErrorMessage string
	ScheduleID   *uint32
	Schedule     *Schedule
}

func (d Deal) Key() string {
	return fmt.Sprintf("%s-%s-%s-%d-%d", d.ClientID, d.Provider, d.PieceCID, d.Start.Unix(), d.End.Unix())
}

type Schedule struct {
	ID                      uint32        `gorm:"primaryKey" json:"id"`
	CreatedAt               time.Time     `json:"createdAt"`
	UpdatedAt               time.Time     `json:"updatedAt"`
	DatasetID               uint32        `json:"datasetId"`
	Dataset                 *Dataset      `gorm:"foreignKey:DatasetID;constraint:OnDelete:CASCADE" json:"dataset,omitempty" swaggerignore:"true"`
	URLTemplate             string        `json:"urlTemplate"`
	HTTPHeaders             []string      `gorm:"type:JSON" json:"httpHeaders"`
	Provider                string        `json:"provider"`
	Price                   float64       `json:"price"`
	TotalDealNumber         int           `json:"totalDealNumber"`
	TotalDealSize           int64         `json:"totalDealSize"`
	Verified                bool          `json:"verified"`
	KeepUnsealed            bool          `json:"keepUnsealed"`
	AnnounceToIPNI          bool          `json:"announceToIPNI"`
	StartDelay              time.Duration `json:"startDelay"`
	Duration                time.Duration `json:"duration"`
	State                   ScheduleState `json:"state"`
	LastProcessedTimestamp  uint64        `json:"lastProcessedTimestamp"`
	ScheduleWorkerID        *string       `json:"scheduleWorkerId"`
	ScheduleWorker          *Worker       `gorm:"foreignKey:ScheduleWorkerID;constraint:OnDelete:NO ACTION" json:"scheduleWorker,omitempty" swaggerignore:"true"`
	ScheduleIntervalSeconds uint64        `json:"scheduleIntervalSeconds"`
	ScheduleDealNumber      int           `json:"scheduleDealNumber"`
	ScheduleDealSize        int64         `json:"scheduleDealSize"`
	MaxPendingDealNumber    int           `json:"maxPendingDealNumber"`
	MaxPendingDealSize      int64         `json:"maxPendingDealSize"`
	Notes                   string        `json:"notes"`
	ErrorMessage            string        `json:"errorMessage"`
	AllowedPieceCIDs        StringSlice   `json:"AllowedPieceCIDs"`
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
	ID         string `gorm:"primaryKey" json:"id"`
	Address    string `gorm:"unique" json:"address"`
	PrivateKey string `json:"privateKey"`
	RemotePeer string `json:"remotePeer"`
}

func (w Wallet) GetExportedKey() (string, error) {
	decrypted, err := DecryptFromBase64String(w.PrivateKey)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

type WalletAssignment struct {
	ID        uint32 `gorm:"primaryKey"`
	WalletID  string
	Wallet    Wallet `gorm:"foreignKey:WalletID;constraint:OnDelete:CASCADE"`
	DatasetID uint32
	Dataset   Dataset `gorm:"foreignKey:DatasetID;constraint:OnDelete:CASCADE"`
}
