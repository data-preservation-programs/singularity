package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"time"
)

type StringSlice []string
type Metadata map[string]interface{}

func (m Metadata) GetS3Metadata() (S3Metadata, error) {
	var out S3Metadata
	err := mapstructure.Decode(m, &out)
	if err != nil {
		return S3Metadata{}, errors.Wrap(err, "failed to decode metadata")
	}
	return out, nil
}

func (m Metadata) GetHTTPMetadata() (HTTPMetadata, error) {
	var out HTTPMetadata
	err := mapstructure.Decode(m, &out)
	if err != nil {
		return HTTPMetadata{}, errors.Wrap(err, "failed to decode metadata")
	}
	return out, nil
}

func (m S3Metadata) Encode() (Metadata, error) {
	var out Metadata
	err := mapstructure.Decode(m, &out)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode metadata")
	}
	return out, nil
}

func (m HTTPMetadata) Encode() (Metadata, error) {
	var out Metadata
	err := mapstructure.Decode(m, &out)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode metadata")
	}
	return out, nil
}

type HTTPMetadata struct {
	Headers map[string]string
}

type S3Metadata struct {
	Region          string
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
}

func (ss StringSlice) Value() (driver.Value, error) {
	return json.Marshal(ss)
}
func (m Metadata) Value() (driver.Value, error) {
	return json.Marshal(m)
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

func (m *Metadata) Scan(src interface{}) error {
	if src == nil {
		*m = nil
		return nil
	}

	source, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Map: %v", src)
	}

	return json.Unmarshal(source, m)
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

func (s SourceType) GetSupportedItemType() ItemType {
	switch s {
	case Dir:
		return File
	case Website:
		return URL
	case S3Path:
		return S3Object
	default:
		return ""
	}
}

const (
	File     ItemType = "file"
	URL      ItemType = "url"
	S3Object ItemType = "s3object"
)

var ItemTypes = []ItemType{File, URL, S3Object}

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
	WorkType      WorkType
	WorkingOn     string
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
	ID                   uint32 `gorm:"primaryKey"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DatasetID            uint32   `gorm:"index"`
	Dataset              *Dataset `gorm:"foreignKey:DatasetID;constraint:OnDelete:CASCADE" json:"Dataset,omitempty"`
	Type                 SourceType
	Path                 string
	Metadata             Metadata `gorm:"type:JSON"`
	PushOnly             bool
	ScanIntervalSeconds  uint64
	ScanningState        WorkState
	ScanningWorkerID     *string
	ScanningWorker       *Worker `gorm:"foreignKey:ScanningWorkerID;references:ID;constraint:OnDelete:SET NULL"`
	LastScannedTimestamp int64
	ErrorMessage         string
	RootDirectoryID      uint64
	RootDirectory        Directory `gorm:"foreignKey:RootDirectoryID;constraint:OnDelete:CASCADE"`
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
	Source          *Source `gorm:"foreignKey:SourceID;constraint:OnDelete:CASCADE"`
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
	ChunkID      *uint32 `gorm:"index"`
	Chunk        *Chunk  `gorm:"foreignKey:ChunkID;constraint:OnDelete:CASCADE" json:"Chunk,omitempty"`
	SourceID     uint32  `gorm:"index"`
	Source       *Source `gorm:"foreignKey:SourceID;constraint:OnDelete:CASCADE" json:"Source,omitempty"`
	Type         ItemType
	Path         string
	Size         uint64
	Offset       uint64
	Length       uint64
	LastModified *time.Time
	Version      uint32
	CID          string `gorm:"column:cid"`
	ErrorMessage string
	DirectoryID  *uint64    `gorm:"index"`
	Directory    *Directory `gorm:"foreignKey:DirectoryID;constraint:OnDelete:CASCADE" json:"Directory,omitempty"`
}

// Directory is a link between parent and child directories
type Directory struct {
	ID       uint64 `gorm:"primaryKey"`
	CID      string `gorm:"column:cid"`
	Name     string
	ParentID *uint64    `gorm:"index"`
	Parent   *Directory `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE" json:"Parent,omitempty"`
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
	CID     string `gorm:"index;column:cid"`
	ChunkID uint32
	Chunk   *Chunk `gorm:"foreignKey:ChunkID;constraint:OnDelete:CASCADE"`
	Block   []byte
	Length  uint32
}

// CarBlock tells us the CIDs of all blocks inside a CAR file
// and the offset of the block inside the CAR file. From this table
// we can determine how to get the block by CID from a CAR file.
// or we can determine how to assemble a CAR file from blocks from
// original file.
type CarBlock struct {
	CarID  uint32
	Car    *Car   `gorm:"foreignKey:CarID;constraint:OnDelete:CASCADE"`
	CID    string `gorm:"index;column:cid"`
	Offset uint64
	Length uint64
	Varint uint64
	// Raw block
	RawBlock []byte
	// If block is null, this block is a part of an item
	SourceID   *uint32
	Source     *Source `gorm:"foreignKey:SourceID;constraint:OnDelete:CASCADE"`
	ItemID     *uint64
	Item       *Item `gorm:"foreignKey:ItemID;constraint:OnDelete:CASCADE"`
	ItemOffset uint64
	// Common
	BlockLength uint64
}
