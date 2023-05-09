package model

import (
	"database/sql/driver"
	"encoding/json"
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
	if out.SecretAccessKey != "" {
		secret, err := DecryptFromBase64String(out.SecretAccessKey)
		if err != nil {
			return S3Metadata{}, errors.Wrap(err, "failed to decrypt secret access key")
		}
		out.SecretAccessKey = string(secret)
	}
	return out, nil
}

func (m Metadata) GetHTTPMetadata() (HTTPMetadata, error) {
	var out HTTPMetadata
	err := mapstructure.Decode(m, &out)
	if err != nil {
		return HTTPMetadata{}, errors.Wrap(err, "failed to decode metadata")
	}
	for k, v := range out.Headers {
		decrypted, err := DecryptFromBase64String(v)
		if err != nil {
			return HTTPMetadata{}, errors.Wrap(err, "failed to decrypt header")
		}
		out.Headers[k] = string(decrypted)
	}
	return out, nil
}

func (m S3Metadata) Encode() (Metadata, error) {
	if m.SecretAccessKey != "" {
		secret, err := EncryptToBase64String([]byte(m.SecretAccessKey))
		if err != nil {
			return nil, errors.Wrap(err, "failed to encrypt secret access key")
		}
		m.SecretAccessKey = secret
	}
	var out Metadata
	err := mapstructure.Decode(m, &out)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode metadata")
	}
	return out, nil
}

func (m HTTPMetadata) Encode() (Metadata, error) {
	for k, v := range m.Headers {
		encrypted, err := EncryptToBase64String([]byte(v))
		if err != nil {
			return nil, errors.Wrap(err, "failed to encrypt header")
		}
		m.Headers[k] = encrypted
	}
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
		return errors.New("failed to scan StringSlice")
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
		return errors.New("failed to scan Metadata")
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
	case Dir, Upload:
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
	// Created means the item has been created is not ready for processing.
	Created WorkState = "created"
	// Ready means the item is ready to be processed.
	Ready WorkState = "ready"
	// Processing means the work is currently being processed.
	Processing WorkState = "processing"
	// Complete means the work is complete.
	Complete WorkState = "complete"
	// Error means the work has some error.
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

// Dataset is the top level object that represents a set of data to be onboarded.
type Dataset struct {
	ID                   uint32      `gorm:"primaryKey" json:"id"`
	Name                 string      `gorm:"unique" json:"name"`
	CreatedAt            time.Time   `json:"createdAt"`
	UpdatedAt            time.Time   `json:"updatedAt"`
	MinSize              uint64      `json:"minSize"`
	MaxSize              uint64      `json:"maxSize"`
	PieceSize            uint64      `json:"pieceSize"`
	OutputDirs           StringSlice `gorm:"type:JSON" json:"outputDirs"`
	EncryptionRecipients StringSlice `gorm:"type:JSON" json:"encryptionRecipients"`
	EncryptionScript     string      `json:"encryptionScript"`
}

// Source represents a source of data, i.e. a local file system directory.
type Source struct {
	ID                   uint32 `gorm:"primaryKey"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DatasetID            uint32   `gorm:"uniqueIndex:dataset_path"`
	Dataset              *Dataset `gorm:"foreignKey:DatasetID;constraint:OnDelete:CASCADE" json:"Dataset,omitempty"`
	Type                 SourceType
	Path                 string   `gorm:"uniqueIndex:dataset_path"`
	Metadata             Metadata `gorm:"type:JSON"`
	PushOnly             bool
	ScanIntervalSeconds  uint64
	ScanningState        WorkState
	ScanningWorkerID     *string `json:"ScanningWorkerID,omitempty"`
	ScanningWorker       *Worker `gorm:"foreignKey:ScanningWorkerID;references:ID;constraint:OnDelete:SET NULL" json:"ScanningWorker,omitempty"`
	LastScannedTimestamp int64
	ErrorMessage         string
	RootDirectoryID      uint64
	RootDirectory        *Directory `gorm:"foreignKey:RootDirectoryID;constraint:OnDelete:CASCADE" json:"RootDirectory,omitempty"`
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

// Chunk is a grouping of items that are packed into a single CAR.
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

// Item makes a reference to the data source item, i.e. a local file.
type Item struct {
	ID           uint64 `gorm:"primaryKey"`
	ScannedAt    time.Time
	ChunkID      *uint32 `gorm:"index"`
	Chunk        *Chunk  `gorm:"foreignKey:ChunkID;constraint:OnDelete:CASCADE" json:"omitempty"`
	SourceID     uint32  `gorm:"index"`
	Source       *Source `gorm:"foreignKey:SourceID;constraint:OnDelete:CASCADE" json:"omitempty"`
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
	Directory    *Directory `gorm:"foreignKey:DirectoryID;constraint:OnDelete:CASCADE" json:"omitempty"`
}

// Directory is a link between parent and child directories.
type Directory struct {
	ID       uint64 `gorm:"primaryKey"`
	CID      string `gorm:"column:cid"`
	Name     string
	ParentID *uint64    `gorm:"index"`
	Parent   *Directory `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE" json:"omitempty"`
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
	Dataset   *Dataset `gorm:"foreignKey:DatasetID;constraint:OnDelete:CASCADE" json:"omitempty"`
	ChunkID   uint32
	Chunk     *Chunk `gorm:"foreignKey:ChunkID;constraint:OnDelete:CASCADE" json:"omitempty"`
	Header    []byte
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
