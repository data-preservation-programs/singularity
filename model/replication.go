package model

import "time"

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

const (
	ScheduleStarted  ScheduleState = "started"
	ScheduleStopped  ScheduleState = "stopped"
	ScheduleFinished ScheduleState = "finished"
)

type Deal struct {
	ID            uint64 `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DealID        *uint64
	State         DealState `gorm:"index:idx_stat"`
	Client        string
	ClientAddress string
	Provider      string `gorm:"index:idx_stat"`
	ProposalID    string
	Label         string
	PieceCID      string `gorm:"column:piece_cid;index"`
	PieceSize     uint64
	Start         time.Time
	Duration      time.Duration
	End           time.Time
	SectorStart   time.Time `gorm:"index:idx_stat"`
	Price         float64
	Verified      bool
	ErrorMessage  string
	ScheduleID    *uint32
	Schedule      *Schedule
}

type Schedule struct {
	ID                   uint32 `gorm:"primaryKey"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DatasetID            uint32
	Dataset              Dataset `gorm:"foreignKey:DatasetID;constraint:OnDelete:CASCADE"`
	URLTemplate          string
	HTTPHeaders          []string `gorm:"type:JSON"`
	Provider             string
	Price                float64
	TotalDealNumber      uint64
	TotalDealSize        uint64
	Verified             bool
	KeepUnsealed         bool
	AnnounceToIPNI       bool
	StartDelay           time.Duration
	Duration             time.Duration
	State                ScheduleState
	SchedulePattern      string
	ScheduleDealNumber   uint64
	ScheduleDealSize     uint64
	MaxPendingDealNumber uint64
	MaxPendingDealSize   uint64
	Notes                string
	ErrorMessage         string
	PieceCIDListPath     string `gorm:"column:piece_cid_list_path"`
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
		s.SchedulePattern == other.SchedulePattern &&
		s.ScheduleDealNumber == other.ScheduleDealNumber &&
		s.ScheduleDealSize == other.ScheduleDealSize &&
		s.MaxPendingDealNumber == other.MaxPendingDealNumber &&
		s.MaxPendingDealSize == other.MaxPendingDealSize &&
		s.State == other.State
}

type Wallet struct {
	ID         string `gorm:"primaryKey"`
	ShortID    string
	PrivateKey string
}

type WalletAssignment struct {
	ID        uint32 `gorm:"primaryKey"`
	WalletID  string
	Wallet    Wallet `gorm:"foreignKey:WalletID;constraint:OnDelete:CASCADE"`
	DatasetID uint32
	Dataset   Dataset `gorm:"foreignKey:DatasetID;constraint:OnDelete:CASCADE"`
}
