package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"time"
)

type StringSlice []string

func (ss StringSlice) Value() (driver.Value, error) {
	return json.Marshal(ss)
}

func (ss *StringSlice) Scan(src interface{}) error {
	if src == nil {
		*ss = nil
		return nil
	}

	source, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan StringSlice: %v", src)
	}

	return json.Unmarshal(source, ss)
}

type SourceType string
type ItemType string
type WorkState string

type WorkType string

const (
	Scan       WorkType = "scan"
	DealMaking WorkType = "deal_making"
	Packing    WorkType = "packing"
)

const (
	Dir     SourceType = "dir"
	Website SourceType = "website"
	S3Path  SourceType = "s3path"
	Upload  SourceType = "upload"
)

const (
	File     ItemType = "file"
	URL      ItemType = "url"
	S3Object ItemType = "s3object"
)

const (
	// Created means the item has been created is not ready for processing
	Created WorkState = "created"
	// Ready means the item is ready to be processed
	Ready WorkState = "ready"
	// Processing means the work is currently being processed
	Processing WorkState = "processing"
	// Complete means the work is complete
	Complete WorkState = "complete"
	// Error means the work has some error
	Error WorkState = "error"
)

type Worker struct {
	ID            string `gorm:"primaryKey"`
	Type          WorkType
	LastHeartbeat time.Time
	Hostname      string
}

type Global struct {
	Key   string `gorm:"primaryKey"`
	Value string
}

// Dataset is the top level object that represents a set of data to be onboarded
type Dataset struct {
	ID                   uint32 `gorm:"primaryKey"`
	Name                 string `gorm:"unique"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	MinSize              uint64
	MaxSize              uint64
	PieceSize            uint64
	OutputDirs           StringSlice `gorm:"type:JSON"`
	EncryptionRecipients StringSlice `gorm:"type:JSON"`
	EncryptionScript     string
	Wallets              []Wallet `gorm:"many2many:wallet_assignments;" json:"Wallets,omitempty"`
}

// Source represents a source of data, i.e. a local file system directory
type Source struct {
	ID               uint32 `gorm:"primaryKey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DatasetID        uint32
	Dataset          *Dataset `gorm:"foreignKey:DatasetID;constraint:OnDelete:CASCADE" json:"Dataset,omitempty"`
	Type             SourceType
	Path             string
	ScanInterval     time.Duration
	ScanningState    WorkState
	ScanningWorkerID *string
	ScanningWorker   *Worker `gorm:"foreignKey:ScanningWorkerID;references:ID;constraint:OnDelete:SET NULL"`
	LastScanned      time.Time
	MaxWait          time.Duration
	ErrorMessage     string
	RootDirectoryID  uint64
	RootDirectory    Directory `gorm:"foreignKey:RootDirectoryID;constraint:OnDelete:CASCADE"`
}

func NewSource(source string) (*Source, error) { // Get the absolute path
	absPath, err := filepath.Abs(source)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get absolute path")
	}

	// Check if the path exists and is a directory
	fileInfo, err := os.Stat(absPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.Wrap(err, "path does not exist")
		}
		return nil, errors.Wrap(err, "failed to get file info")
	}

	if !fileInfo.IsDir() {
		return nil, errors.New("path is not a directory")
	}

	return &Source{
		Type: Dir,
		Path: absPath,
	}, nil
}

// Chunk is a grouping of items that are packed into a single CAR
type Chunk struct {
	ID              uint32 `gorm:"primaryKey"`
	CreatedAt       time.Time
	SourceID        uint32
	Source          Source `gorm:"foreignKey:SourceID;constraint:OnDelete:CASCADE"`
	PackingState    WorkState
	PackingWorkerID *string
	PackingWorker   *Worker `gorm:"foreignKey:PackingWorkerID;references:ID;constraint:OnDelete:SET NULL"`
	ErrorMessage    string
	Items           []Item
}

// Item makes a reference to the data source item, i.e. a local file
type Item struct {
	ID           uint64 `gorm:"primaryKey"`
	ScannedAt    time.Time
	ChunkID      uint32 `gorm:"index"`
	Chunk        *Chunk `gorm:"foreignKey:ChunkID;constraint:OnDelete:CASCADE" json:"Chunk,omitempty"`
	Type         ItemType
	Path         string
	Size         uint64
	Offset       uint64
	Length       uint64
	LastModified *time.Time
	Version      uint32
	CID          string `gorm:"column:cid"`
	ErrorMessage string
	DirectoryID  uint64
	Directory    *Directory `gorm:"foreignKey:DirectoryID;constraint:OnDelete:CASCADE" json:"Directory,omitempty"`
}

// Directory is a link between parent and child directories
type Directory struct {
	ID       uint64 `gorm:"primaryKey"`
	CID      string `gorm:"column:cid"`
	Name     string
	ParentID *uint64
	Parent   *Directory `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE"`
}

// Car makes a reference to a CAR file that has been potentially exported to the disk.
// In the case of inline preparation, the path may be empty so the Car should be constructed
// on the fly using CarBlock, ItemBlock and RawBlock tables.
type Car struct {
	ID        uint32 `gorm:"primaryKey"`
	CreatedAt time.Time
	PieceCID  string `gorm:"column:piece_cid;index"`
	PieceSize uint64
	RootCID   string `gorm:"column:root_cid"`
	FileSize  uint64
	FilePath  string
	DatasetID uint32   `gorm:"index"`
	Dataset   *Dataset `gorm:"foreignKey:DatasetID;constraint:OnDelete:CASCADE" json:"Dataset,omitempty"`
	ChunkID   uint32
	Chunk     *Chunk `gorm:"foreignKey:ChunkID;constraint:OnDelete:CASCADE" json:"Chunk,omitempty"`
	Header    []byte
}

// RawBlock tells us the CIDs of all blocks for constructing the unixfs dag.
// i.e. the blocks that are upper layer of a file, or the directory blocks.
type RawBlock struct {
	CID      string `gorm:"index;column:cid"`
	SourceID uint32
	Source   Source `gorm:"foreignKey:SourceID;constraint:OnDelete:CASCADE"`
	Block    []byte
	Length   uint64
}

// ItemBlock tells us the CIDs of all blocks inside an item
type ItemBlock struct {
	CID    string `gorm:"index;column:cid"`
	ItemID uint64
	Item   Item `gorm:"foreignKey:ItemID;constraint:OnDelete:CASCADE"`
	Offset uint64
	Length uint64
}

// CarBlock tells us the CIDs of all blocks inside a CAR file
// and the offset of the block inside the CAR file. From this table
// we can determine how to get the block by CID from a CAR file.
// or we can determine how to assemble a CAR file from blocks from
// original file.
type CarBlock struct {
	CarID  uint32
	Car    Car    `gorm:"foreignKey:CarID;constraint:OnDelete:CASCADE"`
	CID    string `gorm:"index;column:cid"`
	Offset uint64
	Length uint64
}
