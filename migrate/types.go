package migrate

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	ScanningRequestStatus    string
	GenerationRequestStatus  string
	ReplicationRequestStatus string
)

const (
	ScanningStatusActive       ScanningRequestStatus    = "active"
	ScanningStatusCompleted    ScanningRequestStatus    = "completed"
	ScanningStatusError        ScanningRequestStatus    = "error"
	ScanningStatusPaused       ScanningRequestStatus    = "paused"
	GenerationStatusActive     GenerationRequestStatus  = "active"
	GenerationStatusError      GenerationRequestStatus  = "error"
	GenerationStatusPaused     GenerationRequestStatus  = "paused"
	GenerationStatusCompleted  GenerationRequestStatus  = "completed"
	GenerationStatusCreated    GenerationRequestStatus  = "created"
	GenerationStatusDAG        GenerationRequestStatus  = "dag"
	ReplicationStatusActive    ReplicationRequestStatus = "active"
	ReplicationStatusError     ReplicationRequestStatus = "error"
	ReplicationStatusPaused    ReplicationRequestStatus = "paused"
	ReplicationStatusCompleted ReplicationRequestStatus = "completed"
)

type ScanningRequest struct {
	ID                    primitive.ObjectID    `bson:"_id,omitempty"`
	Name                  string                `bson:"name"`
	Path                  string                `bson:"path"`
	OutDir                string                `bson:"outDir"`
	MinSize               uint64                `bson:"minSize"`
	MaxSize               uint64                `bson:"maxSize"`
	Status                ScanningRequestStatus `bson:"status"`
	ErrorMessage          string                `bson:"errorMessage"`
	TmpDir                string                `bson:"tmpDir"`
	SkipInaccessibleFiles bool                  `bson:"skipInaccessibleFiles"`
}

type GenerationRequest struct {
	ID                    primitive.ObjectID      `bson:"_id,omitempty"`
	DatasetID             string                  `bson:"datasetId"`
	DatasetName           string                  `bson:"datasetName"`
	Path                  string                  `bson:"path"`
	OutDir                string                  `bson:"outDir"`
	Index                 int64                   `bson:"index"`
	Status                GenerationRequestStatus `bson:"status"`
	ErrorMessage          string                  `bson:"errorMessage"`
	DataCID               string                  `bson:"dataCid"`
	CarSize               uint64                  `bson:"carSize"`
	PieceCID              string                  `bson:"pieceCid"`
	PieceSize             uint64                  `bson:"pieceSize"`
	FilenameOverride      string                  `bson:"filenameOverride"`
	TmpDir                string                  `bson:"tmpDir"`
	SkipInaccessibleFiles bool                    `bson:"skipInaccessibleFiles"`
	CreatedAt             time.Time               `bson:"createdAt"`
}

type OutputFileList struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	GenerationID      string             `bson:"generationId"`
	Index             int64              `bson:"index"`
	GeneratedFileList []GeneratedFile    `bson:"generatedFileList"`
}

type GeneratedFile struct {
	Path  string `bson:"path"`
	Dir   bool   `bson:"dir"`
	CID   string `bson:"cid"`
	Size  uint64 `bson:"size"`
	Start uint64 `bson:"start"`
	End   uint64 `bson:"end"`
}

func (g GeneratedFile) IsComplete() bool {
	return g.Start == 0 && (g.End == 0 || g.End == g.Size)
}

type ReplicationRequest struct {
	ID                  primitive.ObjectID       `bson:"_id"`
	CreatedAt           time.Time                `bson:"createdAt"`
	UpdatedAt           time.Time                `bson:"updatedAt"`
	DatasetID           string                   `bson:"datasetId"`
	MaxReplicas         int                      `bson:"maxReplicas"`      // targeted replica per piece
	StorageProviders    string                   `bson:"storageProviders"` // comma separated SP
	Client              string                   `bson:"client"`           // deal sent from client address
	URLPrefix           string                   `bson:"urlPrefix"`
	MaxPrice            float64                  `bson:"maxPrice"`         // unit in Fil
	MaxNumberOfDeals    uint64                   `bson:"maxNumberOfDeals"` // per SP, unlimited if 0
	IsVerified          bool                     `bson:"isVerified"`
	StartDelay          uint64                   `bson:"startDelay"` // in epoch
	Duration            uint64                   `bson:"duration"`   // in epoch
	IsOffline           bool                     `bson:"isOffline"`
	Status              ReplicationRequestStatus `bson:"status"`
	CronSchedule        string                   `bson:"cronSchedule"`
	CronMaxDeals        uint64                   `bson:"cronMaxDeals"`
	CronMaxPendingDeals uint64                   `bson:"cronMaxPendingDeals"`
	FileListPath        string                   `bson:"fileListPath"`
	Notes               string                   `bson:"notes"`
	ErrorMessage        string                   `bson:"errorMessage"`
}
