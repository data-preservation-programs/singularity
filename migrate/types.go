package migrate

import "time"

type ScanningRequest struct {
	ID                     string    `bson:"_id"`
	Name                   string    `bson:"name"`
	MinSize                uint64    `bson:"minSize"`
	MaxSize                uint64    `bson:"maxSize"`
	Path                   string    `bson:"path"`
	Status                 string    `bson:"status"`
	OutDir                 string    `bson:"outDir"`
	TmpDir                 string    `bson:"tmpDir"`
	Scanned                uint64    `bson:"scanned"`
	DagGenerationAttempted bool      `bson:"dagGenerationAttempted"`
	UpdatedAt              time.Time `bson:"updatedAt"`
}

type GenerationRequest struct {
	ID          string    `bson:"_id"`
	DatasetName string    `bson:"datasetName"`
	Path        string    `bson:"path"`
	Index       uint64    `bson:"index"`
	OutDir      string    `bson:"outDir"`
	Status      string    `bson:"status"`
	TmpDir      string    `bson:"tmpDir"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
	CarSize     uint64    `bson:"carSize"`
	DataCID     string    `bson:"dataCid"`
	PieceCID    string    `bson:"pieceCid"`
	PieceSize   uint64    `bson:"pieceSize"`
}

type OutputFileList struct {
	ID                string          `bson:"_id"`
	GeneratedFileList []GeneratedFile `bson:"generatedFileList"`
}

type GeneratedFile struct {
	Path  string `bson:"path"`
	Dir   bool   `bson:"dir"`
	CID   string `bson:"cid"`
	Size  uint64 `bson:"size"`
	Start uint64 `bson:"start"`
	End   uint64 `bson:"end"`
}

type ReplicationRequest struct {
	ID                  string    `bson:"_id"`
	DatasetID           string    `bson:"datasetId"`
	StorageProviders    string    `bson:"storageProviders"`
	URLPrefix           string    `bson:"urlPrefix"`
	MaxPrice            uint64    `bson:"maxPrice"`
	IsVerified          bool      `bson:"isVerified"`
	StartDelay          uint64    `bson:"startDelay"`
	Duration            uint64    `bson:"duration"`
	MaxNumberOfDeals    uint64    `bson:"maxNumberOfDeals"`
	Status              string    `bson:"status"`
	CreatedAt           time.Time `bson:"createdAt"`
	UpdatedAt           time.Time `bson:"updatedAt"`
	FileListPath        string    `bson:"fileListPath"`
	CronSchedule        string    `bson:"cronSchedule"`
	CronMaxDeals        uint64    `bson:"cronMaxDeals"`
	CronMaxPendingDeals uint64    `bson:"cronMaxPendingDeals"`
	Notes               string    `bson:"notes"`
}

type DealState struct {
	Client     string    `bson:"client"`
	Provider   string    `bson:"provider"`
	DealCID    string    `bson:"dealCid"`
	PieceCID   string    `bson:"pieceCid"`
	StartEpoch int64     `bson:"startEpoch"`
	Expiration int64     `bson:"expiration"`
	Duration   int64     `bson:"duration"`
	DealID     uint64    `bson:"dealId"`
	State      string    `bson:"state"`
	CreatedAt  time.Time `bson:"createdAt"`
	UpdatedAt  time.Time `bson:"updatedAt"`
}
