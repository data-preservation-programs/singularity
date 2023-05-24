package model

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/ipfs/go-cid"
	"time"

	"github.com/pkg/errors"
)

type StringSlice []string
type Metadata map[string]string

type CIDBytes []byte

func (c CIDBytes) MarshalJSON() ([]byte, error) {
	v, err := cid.Cast(c)
	if err != nil {
		return nil, err
	}
	return []byte(v.String()), nil
}

func (c *CIDBytes) UnmarshalJSON(b []byte) error {
	v, err := cid.Decode(string(b))
	if err != nil {
		return err
	}

	*c = v.Bytes()
	return nil
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

type SourceType = string
type WorkState string

type WorkType string

const (
	Scan       WorkType = "scan"
	DealMaking WorkType = "deal_making"
	Packing    WorkType = "packing"
)

const (
	Local  SourceType = "local"
	Upload SourceType = "upload"
)

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
	MinSize              int64       `json:"minSize"`
	MaxSize              int64       `json:"maxSize"`
	PieceSize            int64       `json:"pieceSize"`
	OutputDirs           StringSlice `gorm:"type:JSON" json:"outputDirs"`
	EncryptionRecipients StringSlice `gorm:"type:JSON" json:"encryptionRecipients"`
	EncryptionScript     string      `json:"encryptionScript"`
	Wallets              []Wallet    `gorm:"many2many:wallet_assignments" json:"wallets,omitempty" swaggerignore:"true"`
}

// Source represents a source of data, i.e. a local file system directory.
type Source struct {
	ID                   uint32     `gorm:"primaryKey" json:"id"`
	CreatedAt            time.Time  `json:"createdAt"`
	UpdatedAt            time.Time  `json:"updatedAt"`
	DatasetID            uint32     `gorm:"uniqueIndex:dataset_type_path" json:"datasetId"`
	Dataset              *Dataset   `gorm:"foreignKey:DatasetID;constraint:OnDelete:CASCADE" json:"dataset,omitempty" swaggerignore:"true"`
	Type                 SourceType `gorm:"uniqueIndex:dataset_type_path" json:"type"`
	Path                 string     `gorm:"uniqueIndex:dataset_type_path" json:"path"`
	Metadata             Metadata   `gorm:"type:JSON" json:"metadata"`
	PushOnly             bool       `json:"pushOnly"`
	ScanIntervalSeconds  uint64     `json:"scanIntervalSeconds"`
	ScanningState        WorkState  `json:"scanningState"`
	ScanningWorkerID     *string    `json:"scanningWorkerId,omitempty"`
	ScanningWorker       *Worker    `gorm:"foreignKey:ScanningWorkerID;references:ID;constraint:OnDelete:SET NULL" json:"scanningWorker,omitempty" swaggerignore:"true"`
	LastScannedTimestamp int64      `json:"lastScannedTimestamp"`
	ErrorMessage         string     `json:"errorMessage"`
	RootDirectoryID      uint64     `json:"rootDirectoryId"`
	RootDirectory        *Directory `gorm:"foreignKey:RootDirectoryID;constraint:OnDelete:CASCADE" json:"rootDirectory,omitempty" swaggerignore:"true"`
}

// Chunk is a grouping of items that are packed into a single CAR.
type Chunk struct {
	ID              uint32    `gorm:"primaryKey" json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	SourceID        uint32    `json:"sourceId"`
	Source          *Source   `gorm:"foreignKey:SourceID;constraint:OnDelete:CASCADE" json:"source,omitempty" swaggerignore:"true"`
	PackingState    WorkState `json:"packingState"`
	PackingWorkerID *string   `json:"packingWorkerId,omitempty"`
	PackingWorker   *Worker   `gorm:"foreignKey:PackingWorkerID;references:ID;constraint:OnDelete:SET NULL" json:"packingWorker,omitempty" swaggerignore:"true"`
	ErrorMessage    string    `json:"errorMessage"`
	Items           []Item    `json:"items,omitempty" swaggerignore:"true"`
	Car             *Car      `json:"car,omitempty"`
}

// Item makes a reference to the data source item, i.e. a local file.
type Item struct {
	ID           uint64     `gorm:"primaryKey" json:"id"`
	ScannedAt    time.Time  `json:"scannedAt"`
	ChunkID      *uint32    `gorm:"index" json:"chunkId"`
	Chunk        *Chunk     `gorm:"foreignKey:ChunkID;constraint:OnDelete:CASCADE" json:"chunk,omitempty" swaggerignore:"true"`
	SourceID     uint32     `gorm:"index" json:"sourceId"`
	Source       *Source    `gorm:"foreignKey:SourceID;constraint:OnDelete:CASCADE" json:"source,omitempty" swaggerignore:"true"`
	Path         string     `json:"path"`
	Size         int64      `json:"size"`
	Offset       int64      `json:"offset"`
	Length       int64      `json:"length"`
	LastModified time.Time  `json:"lastModified"`
	CID          CIDBytes   `gorm:"column:cid" json:"cid"`
	DirectoryID  *uint64    `gorm:"index" json:"directoryId"`
	Directory    *Directory `gorm:"foreignKey:DirectoryID;constraint:OnDelete:CASCADE" json:"directory,omitempty" swaggerignore:"true"`
}

// Directory is a link between parent and child directories.
type Directory struct {
	ID       uint64     `gorm:"primaryKey" json:"id"`
	CID      string     `gorm:"column:cid" json:"cid"`
	Name     string     `json:"name"`
	ParentID *uint64    `gorm:"index" json:"parentId"`
	Parent   *Directory `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE" json:"parent,omitempty" swaggerignore:"true"`
}

// Car makes a reference to a CAR file that has been potentially exported to the disk.
// In the case of inline preparation, the path may be empty so the Car should be constructed
// on the fly using CarBlock, ItemBlock and RawBlock tables.
type Car struct {
	ID        uint32    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	PieceCID  CIDBytes  `gorm:"column:piece_cid;index" json:"pieceCid"`
	PieceSize int64     `json:"pieceSize"`
	RootCID   CIDBytes  `gorm:"column:root_cid" json:"rootCid"`
	FileSize  int64     `json:"fileSize"`
	FilePath  string    `json:"filePath"`
	DatasetID uint32    `gorm:"index" json:"datasetId"`
	Dataset   *Dataset  `gorm:"foreignKey:DatasetID;constraint:OnDelete:CASCADE" json:"dataset,omitempty" swaggerignore:"true"`
	SourceID  *uint32   `gorm:"index" json:"sourceId"`
	Source    *Source   `gorm:"foreignKey:SourceID;constraint:OnDelete:CASCADE" json:"source,omitempty" swaggerignore:"true"`
	ChunkID   *uint32   `json:"chunkId"`
	Chunk     *Chunk    `gorm:"foreignKey:ChunkID;constraint:OnDelete:CASCADE" json:"chunk,omitempty" swaggerignore:"true"`
	Header    []byte    `json:"header"`
}

// CarBlock tells us the CIDs of all blocks inside a CAR file
// and the offset of the block inside the CAR file. From this table
// we can determine how to get the block by CID from a CAR file.
// or we can determine how to assemble a CAR file from blocks from
// original file.
type CarBlock struct {
	CarID uint32   `json:"carId"`
	Car   *Car     `gorm:"foreignKey:CarID;constraint:OnDelete:CASCADE" json:"car,omitempty" swaggerignore:"true"`
	CID   CIDBytes `gorm:"index;column:cid" json:"cid"`
	// Offset of the varint inside the CAR
	CarOffset      int64 `json:"carOffset"`
	CarBlockLength int32 `json:"carBlockLength"`
	// Value of the varint
	Varint []byte `json:"varint"`
	// Raw block
	RawBlock []byte `json:"rawBlock"`
	// If block is null, this block is a part of an item
	ItemID     *uint64 `json:"itemId"`
	Item       *Item   `gorm:"foreignKey:ItemID;constraint:OnDelete:CASCADE" json:"item,omitempty" swaggerignore:"true"`
	ItemOffset int64   `json:"itemOffset"`
}

func (c CarBlock) BlockLength() int32 {
	if c.RawBlock != nil {
		return int32(len(c.RawBlock))
	}
	return c.CarBlockLength - int32(len(c.CID)) - int32(len(c.Varint))
}
