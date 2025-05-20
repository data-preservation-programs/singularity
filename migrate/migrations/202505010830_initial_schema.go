package migrations

import (
	"strconv"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/ipfs/go-cid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// NOTE: This recreates original models at time of transition from AutoMigrate
// to versioned migrations so that future modifications to the actual models
// don't change this initial schema definition.
type StringSlice []string
type ConfigMap map[string]string
type CID cid.Cid
type ClientConfig struct {
	ConnectTimeout          *time.Duration    `cbor:"1,keyasint,omitempty"  json:"connectTimeout,omitempty"          swaggertype:"primitive,integer"` // HTTP Client Connect timeout
	Timeout                 *time.Duration    `cbor:"2,keyasint,omitempty"  json:"timeout,omitempty"                 swaggertype:"primitive,integer"` // IO idle timeout
	ExpectContinueTimeout   *time.Duration    `cbor:"3,keyasint,omitempty"  json:"expectContinueTimeout,omitempty"   swaggertype:"primitive,integer"` // Timeout when using expect / 100-continue in HTTP
	InsecureSkipVerify      *bool             `cbor:"4,keyasint,omitempty"  json:"insecureSkipVerify,omitempty"`                                      // Do not verify the server SSL certificate (insecure)
	NoGzip                  *bool             `cbor:"5,keyasint,omitempty"  json:"noGzip,omitempty"`                                                  // Don't set Accept-Encoding: gzip
	UserAgent               *string           `cbor:"6,keyasint,omitempty"  json:"userAgent,omitempty"`                                               // Set the user-agent to a specified string
	CaCert                  []string          `cbor:"7,keyasint,omitempty"  json:"caCert,omitempty"`                                                  // Paths to CA certificate used to verify servers
	ClientCert              *string           `cbor:"8,keyasint,omitempty"  json:"clientCert,omitempty"`                                              // Path to Client SSL certificate (PEM) for mutual TLS auth
	ClientKey               *string           `cbor:"9,keyasint,omitempty"  json:"clientKey,omitempty"`                                               // Path to Client SSL private key (PEM) for mutual TLS auth
	Headers                 map[string]string `cbor:"10,keyasint,omitempty" json:"headers,omitempty"`                                                 // Set HTTP header for all transactions
	DisableHTTP2            *bool             `cbor:"11,keyasint,omitempty" json:"disableHttp2,omitempty"`                                            // Disable HTTP/2 in the transport
	DisableHTTPKeepAlives   *bool             `cbor:"12,keyasint,omitempty" json:"disableHttpKeepAlives,omitempty"`                                   // Disable HTTP keep-alives and use each connection once.
	RetryMaxCount           *int              `cbor:"13,keyasint,omitempty" json:"retryMaxCount,omitempty"`                                           // Maximum number of retries. Default is 10 retries.
	RetryDelay              *time.Duration    `cbor:"14,keyasint,omitempty" json:"retryDelay,omitempty"              swaggertype:"primitive,integer"` // Delay between retries. Default is 1s.
	RetryBackoff            *time.Duration    `cbor:"15,keyasint,omitempty" json:"retryBackoff,omitempty"            swaggertype:"primitive,integer"` // Constant backoff between retries. Default is 1s.
	RetryBackoffExponential *float64          `cbor:"16,keyasint,omitempty" json:"retryBackoffExponential,omitempty"`                                 // Exponential backoff between retries. Default is 1.0.
	SkipInaccessibleFile    *bool             `cbor:"17,keyasint,omitempty" json:"skipInaccessibleFile,omitempty"`                                    // Skip inaccessible files. Default is false.
	UseServerModTime        *bool             `cbor:"18,keyasint,omitempty" json:"useServerModTime,omitempty"`                                        // Use server modified time instead of object metadata
	LowLevelRetries         *int              `cbor:"19,keyasint,omitempty" json:"lowlevelRetries,omitempty"`                                         // Maximum number of retries for low-level client errors. Default is 10 retries.
	ScanConcurrency         *int              `cbor:"20,keyasint,omitempty" json:"scanConcurrency,omitempty"`                                         // Maximum number of concurrent scan requests. Default is 1.
}
type WorkerType string
type Worker struct {
	ID            string     `gorm:"primaryKey"    json:"id"`
	LastHeartbeat time.Time  `json:"lastHeartbeat"`
	Hostname      string     `json:"hostname"`
	Type          WorkerType `json:"type"`
}
type Global struct {
	Key   string `gorm:"primaryKey" json:"key"`
	Value string `json:"value"`
}
type Wallet struct {
	ID         string `gorm:"primaryKey;size:15"   json:"id"`      // ID is the short ID of the wallet
	Address    string `gorm:"index"                json:"address"` // Address is the Filecoin full address of the wallet
	PrivateKey string `json:"privateKey,omitempty" table:"-"`      // PrivateKey is the private key of the wallet
}
type PreparationID uint32
type Preparation struct {
	ID                PreparationID `gorm:"primaryKey"        json:"id"`
	Name              string        `gorm:"unique"            json:"name"`
	CreatedAt         time.Time     `json:"createdAt"         table:"verbose;format:2006-01-02 15:04:05"`
	UpdatedAt         time.Time     `json:"updatedAt"         table:"verbose;format:2006-01-02 15:04:05"`
	DeleteAfterExport bool          `json:"deleteAfterExport"` // DeleteAfterExport is a flag that indicates whether the source files should be deleted after export.
	MaxSize           int64         `json:"maxSize"`
	PieceSize         int64         `json:"pieceSize"`
	NoInline          bool          `json:"noInline"`
	NoDag             bool          `json:"noDag"`
	Wallets           []Wallet      `gorm:"many2many:wallet_assignments"                             json:"wallets,omitempty"        swaggerignore:"true"                   table:"expand"`
	SourceStorages    []Storage     `gorm:"many2many:source_attachments;constraint:OnDelete:CASCADE" json:"sourceStorages,omitempty" table:"expand;header:Source Storages:"`
	OutputStorages    []Storage     `gorm:"many2many:output_attachments;constraint:OnDelete:CASCADE" json:"outputStorages,omitempty" table:"expand;header:Output Storages:"`
}

func (s *Preparation) FindByIDOrName(db *gorm.DB, name string, preloads ...string) error {
	id, err := strconv.ParseUint(name, 10, 32)
	if err == nil {
		for _, preload := range preloads {
			db = db.Preload(preload)
		}
		return db.First(s, id).Error
	} else {
		for _, preload := range preloads {
			db = db.Preload(preload)
		}
		return db.Where("name = ?", name).First(s).Error
	}
}
func (s *Preparation) SourceAttachments(db *gorm.DB, preloads ...string) ([]SourceAttachment, error) {
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	var attachments []SourceAttachment
	err := db.Where("preparation_id = ?", s.ID).Find(&attachments).Error
	return attachments, errors.Wrap(err, "failed to find source attachments")
}

type StorageID uint32
type Storage struct {
	ID                   StorageID     `cbor:"-"                    gorm:"primaryKey" json:"id"`
	Name                 string        `cbor:"-"                    gorm:"unique"     json:"name"`
	CreatedAt            time.Time     `cbor:"-"                    json:"createdAt"  table:"verbose;format:2006-01-02 15:04:05"`
	UpdatedAt            time.Time     `cbor:"-"                    json:"updatedAt"  table:"verbose;format:2006-01-02 15:04:05"`
	Type                 string        `cbor:"1,keyasint,omitempty" json:"type"`
	Path                 string        `cbor:"2,keyasint,omitempty" json:"path"`                                                                  // Path is the path to the storage root.
	Config               ConfigMap     `cbor:"3,keyasint,omitempty" gorm:"type:JSON"  json:"config"                              table:"verbose"` // Config is a map of key-value pairs that can be used to store RClone options.
	ClientConfig         ClientConfig  `cbor:"4,keyasint,omitempty" gorm:"type:JSON"  json:"clientConfig"                        table:"verbose"` // ClientConfig is the HTTP configuration for the storage, if applicable.
	PreparationsAsSource []Preparation `cbor:"-" gorm:"many2many:source_attachments;constraint:OnDelete:CASCADE" json:"preparationsAsSource,omitempty" table:"expand;header:As Source: "`
	PreparationsAsOutput []Preparation `cbor:"-" gorm:"many2many:output_attachments;constraint:OnDelete:CASCADE" json:"preparationsAsOutput,omitempty" table:"expand;header:As Output: "`
}
type ScheduleID uint32
type ScheduleState string
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
	PreparationID         PreparationID `json:"preparationId"`
	Preparation           *Preparation  `gorm:"foreignKey:PreparationID;constraint:OnDelete:CASCADE" json:"preparation,omitempty" swaggerignore:"true" table:"expand"`
}
type DealState string
type DealID uint64
type Deal struct {
	ID               DealID      `gorm:"primaryKey"                      json:"id"                                  table:"verbose"`
	CreatedAt        time.Time   `json:"createdAt"                       table:"verbose;format:2006-01-02 15:04:05"`
	UpdatedAt        time.Time   `json:"updatedAt"                       table:"verbose;format:2006-01-02 15:04:05"`
	LastVerifiedAt   *time.Time  `json:"lastVerifiedAt"                  table:"verbose;format:2006-01-02 15:04:05"` // LastVerifiedAt is the last time the deal was verified as active by the tracker
	DealID           *uint64     `gorm:"unique"                          json:"dealId"`
	State            DealState   `gorm:"index:idx_pending"               json:"state"`
	Provider         string      `json:"provider"`
	ProposalID       string      `json:"proposalId"                      table:"verbose"`
	Label            string      `json:"label"                           table:"verbose"`
	PieceCID         CID         `gorm:"column:piece_cid;index;size:255" json:"pieceCid"                            swaggertype:"string"`
	PieceSize        int64       `json:"pieceSize"`
	StartEpoch       int32       `json:"startEpoch"`
	EndEpoch         int32       `json:"endEpoch"                        table:"verbose"`
	SectorStartEpoch int32       `json:"sectorStartEpoch"                table:"verbose"`
	Price            string      `json:"price"`
	Verified         bool        `json:"verified"`
	ErrorMessage     string      `json:"errorMessage"                    table:"verbose"`
	ScheduleID       *ScheduleID `json:"scheduleId"                                         table:"verbose"`
	Schedule         *Schedule   `gorm:"foreignKey:ScheduleID;constraint:OnDelete:SET NULL" json:"schedule,omitempty" swaggerignore:"true" table:"expand"`
	ClientID         string      `gorm:"index:idx_pending"                                  json:"clientId"`
	Wallet           *Wallet     `gorm:"foreignKey:ClientID;constraint:OnDelete:SET NULL"   json:"wallet,omitempty"   swaggerignore:"true" table:"expand"`
}
type OutputAttachment struct {
	ID            uint32 `gorm:"primaryKey"`
	PreparationID PreparationID
	StorageID     StorageID
}
type SourceAttachment struct {
	ID            uint32 `gorm:"primaryKey"`
	PreparationID PreparationID
	StorageID     StorageID
}
type Job struct {
	ID            uint32 `gorm:"primaryKey"`
	PreparationID PreparationID
	Status        string
	CreatedAt     time.Time
}
type File struct {
	ID         uint32 `gorm:"primaryKey"`
	Path       string
	Size       int64
	ModifiedAt time.Time
}
type FileRange struct {
	ID     uint32 `gorm:"primaryKey"`
	FileID uint32
	Offset int64
	Length int64
}
type Directory struct {
	ID   uint32 `gorm:"primaryKey"`
	Path string
	Size int64
}
type Car struct {
	ID      uint32 `gorm:"primaryKey"`
	RootCID CID
	Size    int64
}
type CarBlock struct {
	ID    uint32 `gorm:"primaryKey"`
	CarID uint32
	CID   CID
	Size  int64
}

// Create migration for initial database schema
func _202505010830_initial_schema() *gormigrate.Migration {
	var InitTables = []any{
		&Worker{},
		&Global{},
		&Preparation{},
		&Storage{},
		&OutputAttachment{},
		&SourceAttachment{},
		&Job{},
		&File{},
		&FileRange{},
		&Directory{},
		&Car{},
		&CarBlock{},
		&Deal{},
		&Schedule{},
		&Wallet{},
	}

	return &gormigrate.Migration{
		ID: "202505010830",
		Migrate: func(tx *gorm.DB) error {
			// NOTE: this should match any existing database at the time of transition
			// to versioned migration strategy
			return tx.AutoMigrate(InitTables...)
		},
		Rollback: func(tx *gorm.DB) error {
			for _, table := range InitTables {
				err := tx.Migrator().DropTable(table)
				if err != nil {
					return errors.Wrap(err, "failed to drop table")
				}
			}
			return nil
		},
	}
}
