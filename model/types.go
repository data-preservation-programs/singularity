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

type SourceType int
type ItemType int
type WorkState int

const (
	Dir     SourceType = 1
	Website SourceType = 2
	S3Path  SourceType = 3
	Upload  SourceType = 4
)

const (
	File     ItemType = 1
	URL      ItemType = 2
	S3Object ItemType = 3
)

const (
	// Created means the item has been created is not ready for processing
	Created WorkState = 1
	// Ready means the item is ready to be processed
	Ready WorkState = 2
	// Processing means the work is currently being processed
	Processing WorkState = 3
	// Complete means the work is complete
	Complete WorkState = 4
	// Error means the work has some error
	Error WorkState = 5
)

type Worker struct {
	ID            string `gorm:"primaryKey"`
	LastHeartbeat time.Time
}

type Dataset struct {
	ID                   uint32 `gorm:"primaryKey"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Name                 string `gorm:"unique"`
	MinSize              uint64
	MaxSize              uint64
	PieceSize            uint64
	OutputDirs           []string `gorm:"type:JSON"`
	EncryptionRecipients []string `gorm:"type:JSON"`
	EncryptionScript     string
	MaxWait              time.Duration
}

type Source struct {
	ID               uint32 `gorm:"primaryKey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DatasetID        uint32
	Type             SourceType
	Path             string
	ScanInterval     time.Duration
	ScanningState    WorkState
	ScanningWorkerID string
	ScanningWorker   Worker `gorm:"foreignKey:ScanningWorkerID;references:ID;constraint:OnDelete:CASCADE"`
	LastScanned      time.Time
	RootCID          string
	ErrorMessage     string
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

type Chunk struct {
	ID           int32 `gorm:"primaryKey"`
	CreatedAt    time.Time
	DatasetID    uint32
	PieceSize    uint64
	PieceCID     string
	PackingState WorkState
	ErrorMessage *string
}

type Item struct {
	ID           uint64 `gorm:"primaryKey"`
	ScannedAt    time.Time
	DatasetID    uint32
	SourceID     uint32
	ChunkID      *uint32
	Type         ItemType
	Path         string
	Size         uint64
	Offset       *uint64
	Length       *uint64
	LastModified *time.Time
	Version      uint32
	CID          string
	ErrorMessage string
}

type BlockRaw struct {
	CID   string `gorm:"primaryKey"`
	Block []byte
}

type BlockReference struct {
	CID          string `gorm:"index"`
	Type         ItemType
	Path         string
	Offset       *uint64
	Length       *uint64
	LastModified time.Time
}
