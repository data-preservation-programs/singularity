package model

import (
	"context"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/ipfs/go-cid"
	"gorm.io/gorm"
)

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

// Preparation is a data preparation definition that can attach multiple source storages and up to one output storage.
type Preparation struct {
	ID                uint32    `gorm:"primaryKey"        json:"id"`
	CreatedAt         time.Time `json:"createdAt"         table:"verbose;format:2006-01-02 15:04:05"`
	UpdatedAt         time.Time `json:"updatedAt"         table:"verbose;format:2006-01-02 15:04:05"`
	DeleteAfterExport bool      `json:"deleteAfterExport"` // DeleteAfterExport is a flag that indicates whether the source files should be deleted after export.
	MaxSize           int64     `json:"maxSize"`
	PieceSize         int64     `json:"pieceSize"`

	// Associations
	Wallets        []Wallet  `gorm:"many2many:wallet_assignments"                             json:"wallets,omitempty"        swaggerignore:"true" table:"expand"`
	SourceStorages []Storage `gorm:"many2many:source_attachments;constraint:OnDelete:CASCADE" json:"sourceStorages,omitempty" swaggerignore:"true" table:"expand;header:Source Storages:"`
	OutputStorages []Storage `gorm:"many2many:output_attachments;constraint:OnDelete:CASCADE" json:"outputStorages,omitempty" swaggerignore:"true" table:"expand;header:Output Storages:"`
}

// Storage is a storage system definition that can be used as either source or output of a Preparation.
type Storage struct {
	ID        uint32    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"unique"     json:"name"`
	CreatedAt time.Time `json:"createdAt"  table:"verbose;format:2006-01-02 15:04:05"`
	UpdatedAt time.Time `json:"updatedAt"  table:"verbose;format:2006-01-02 15:04:05"`
	Type      string    `json:"type"`
	Path      string    `json:"path"`                       // Path is the path to the storage root.
	Config    ConfigMap `gorm:"type:JSON"  table:"verbose"` // Config is a map of key-value pairs that can be used to store RClone options.

	// Associations
	PreparationsAsSource []Preparation `gorm:"many2many:source_attachments;constraint:OnDelete:CASCADE" json:"preparationsAsSource,omitempty" swaggerignore:"true" table:"expand;header:As Source: "`
	PreparationsAsOutput []Preparation `gorm:"many2many:output_attachments;constraint:OnDelete:CASCADE" json:"preparationsAsOutput,omitempty" swaggerignore:"true" table:"expand;header:As Output: "`
	// For Cbor marshalling
	_ struct{} `cbor:",toarray"                                                               json:"-"                          swaggerignore:"true"`
}

// SourceAttachment is a link between a Preparation and a Storage that is used as a source.
type SourceAttachment struct {
	ID uint32 `gorm:"primaryKey" json:"id"`

	// Associations
	PreparationID   uint32       `gorm:"uniqueIndex:prep_source"                              json:"preparationId"`
	Preparation     *Preparation `gorm:"foreignKey:PreparationID;constraint:OnDelete:CASCADE" json:"preparation,omitempty" swaggerignore:"true"`
	StorageID       uint32       `gorm:"uniqueIndex:prep_source"                              json:"storageId"`
	Storage         *Storage     `gorm:"foreignKey:StorageID;constraint:OnDelete:CASCADE"     json:"storage,omitempty"     swaggerignore:"true"`
	LastScannedPath string       `json:"lastScannedPath"`
}

func (s SourceAttachment) RootDirectoryCID(ctx context.Context, db *gorm.DB) (cid.Cid, error) {
	db = db.WithContext(ctx)
	var root Directory
	err := db.Select("cid").Where("attachment_id = ? AND parent_id is null", s.ID).First(&root).Error
	return cid.Cid(root.CID), errors.WithStack(err)
}

func (s SourceAttachment) RootDirectoryID(ctx context.Context, db *gorm.DB) (uint64, error) {
	db = db.WithContext(ctx)
	var root Directory
	err := db.Select("id").Where("attachment_id = ? AND parent_id is null", s.ID).First(&root).Error
	return root.ID, errors.WithStack(err)
}

// OutputAttachment is a link between a Preparation and a Storage that is used as an output.
type OutputAttachment struct {
	ID uint32 `gorm:"primaryKey" json:"id"`

	// Associations
	PreparationID uint32       `gorm:"uniqueIndex:prep_output"                              json:"preparationId"`
	Preparation   *Preparation `gorm:"foreignKey:PreparationID;constraint:OnDelete:CASCADE" json:"preparation,omitempty" swaggerignore:"true"`
	StorageID     uint32       `gorm:"uniqueIndex:prep_output"                              json:"storageId"`
	Storage       *Storage     `gorm:"foreignKey:StorageID;constraint:OnDelete:CASCADE"     json:"storage,omitempty"     swaggerignore:"true"`
}

// Job is a job that is executed by a worker.
// The composite index on Type and State is used to find jobs that are ready to be executed.
type Job struct {
	ID              uint64   `gorm:"primaryKey"           json:"id"`
	Type            JobType  `gorm:"index:job_type_state" json:"type"`
	State           JobState `gorm:"index:job_type_state" json:"state"`
	ErrorMessage    string   `json:"errorMessage"`
	ErrorStackTrace string   `json:"errorStackTrace"      table:"verbose"`

	// Associations
	WorkerID     *string           `gorm:"size:63"                                                        json:"workerId,omitempty"`
	Worker       *Worker           `gorm:"foreignKey:WorkerID;references:ID;constraint:OnDelete:SET NULL" json:"worker,omitempty"     swaggerignore:"true" table:"verbose;expand"`
	AttachmentID uint32            `json:"attachmentId"                                                   table:"verbose"`
	Attachment   *SourceAttachment `gorm:"foreignKey:AttachmentID;constraint:OnDelete:CASCADE"            json:"attachment,omitempty" swaggerignore:"true" table:"expand"`
	FileRanges   []FileRange       `gorm:"foreignKey:JobID;constraint:OnDelete:SET NULL"                  json:"fileRanges,omitempty" swaggerignore:"true" table:"-"`
}

// File makes a reference to the source storage file, e.g., a local file.
// The index on Path is used as part of scanning to find existing file and add new versions.
// The index on DirectoryID is used to find all files in a directory.
type File struct {
	ID               uint64 `gorm:"primaryKey"                     json:"id"`
	CID              CID    `gorm:"column:cid;type:bytes;size:255" json:"cid"`  // CID is the CID of the file.
	Path             string `gorm:"index"                          json:"path"` // Path is the relative path to the file inside the storage.
	Hash             string `json:"hash"`                                       // Hash is the hash of the file.
	Size             int64  `json:"size"`                                       // Size is the size of the file in bytes.
	LastModifiedNano int64  `json:"lastModifiedNano"`

	// Associations
	AttachmentID uint32            `json:"attachmentId"`
	Attachment   *SourceAttachment `gorm:"foreignKey:AttachmentID;constraint:OnDelete:CASCADE" json:"attachment,omitempty" swaggerignore:"true"`
	DirectoryID  *uint64           `gorm:"index"                                               json:"directoryId"`
	Directory    *Directory        `gorm:"foreignKey:DirectoryID;constraint:OnDelete:CASCADE"  json:"directory,omitempty"  swaggerignore:"true"`
	FileRanges   []FileRange       `gorm:"constraint:OnDelete:CASCADE"                         json:"fileRanges,omitempty" swaggerignore:"true"`

	// For Cbor marshalling
	_ struct{} `cbor:",toarray"                                                  json:"-"                    swaggerignore:"true"`
}

func (i File) FileName() string {
	return i.Path[strings.LastIndex(i.Path, "/")+1:]
}

// Directory is a link between parent and child directories.
// The index on AttachmentID and ParentID is used to find all root directories, as well as all directories in a directory.
type Directory struct {
	ID       uint64 `gorm:"primaryKey"            json:"id"`
	CID      CID    `gorm:"column:cid;type:bytes" json:"cid"`                      // CID is the CID of the directory.
	Data     []byte `gorm:"column:data"           json:"-"   swaggerignore:"true"` // Data is the serialized directory data.
	Name     string `json:"name"`                                                  // Name is the name of the directory.
	Exported bool   `json:"exported"`                                              // Exported is a flag that indicates whether the directory has been exported to the DAG.

	// Associations
	AttachmentID uint32            `gorm:"index:directory_source_parent"                       json:"attachmentId"`
	Attachment   *SourceAttachment `gorm:"foreignKey:AttachmentID;constraint:OnDelete:CASCADE" json:"attachment,omitempty" swaggerignore:"true"`
	ParentID     *uint64           `gorm:"index:directory_source_parent"                       json:"parentId"`
	Parent       *Directory        `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE"     json:"parent,omitempty"     swaggerignore:"true"`
}

// FileRange is a range of bytes inside File.
// The index on FileID is used to find all FileRange in a file.
// The index on JobID is used to find all FileRange in a job.
type FileRange struct {
	ID     uint64 `gorm:"primaryKey"            json:"id"`
	Offset int64  `json:"offset"`                           // Offset is the offset of the range inside the file.
	Length int64  `json:"length"`                           // Length is the length of the range in bytes.
	CID    CID    `gorm:"column:cid;type:bytes" json:"cid"` // CID is the CID of the range.

	// Associations
	JobID  *uint64 `gorm:"index"                                         json:"jobId"`
	Job    *Job    `gorm:"foreignKey:JobID;constraint:OnDelete:SET NULL" json:"job,omitempty"  swaggerignore:"true"`
	FileID uint64  `gorm:"index"                                         json:"fileId"`
	File   *File   `gorm:"foreignKey:FileID;constraint:OnDelete:CASCADE" json:"file,omitempty" swaggerignore:"true"`
}

// Car makes a reference to a CAR file that has been potentially exported to the disk.
// In the case of inline preparation, the path may be empty so the Car should be constructed
// on the fly using CarBlock.
// The index on PieceCID is to find all CARs that can matches the PieceCID
type Car struct {
	ID          uint32    `gorm:"primaryKey"                                        json:"id"                                  table:"verbose"`
	CreatedAt   time.Time `json:"createdAt"                                         table:"verbose;format:2006-01-02 15:04:05"`
	PieceCID    CID       `gorm:"column:piece_cid;index;type:bytes;size:255"        json:"pieceCid"`
	PieceSize   int64     `json:"pieceSize"`
	RootCID     CID       `gorm:"column:root_cid;type:bytes"                        json:"rootCid"`
	FileSize    int64     `json:"fileSize"`
	StorageID   *uint32   `json:"storageId"                                         table:"verbose"`
	Storage     *Storage  `gorm:"foreignKey:StorageID;constraint:OnDelete:SET NULL" json:"storage,omitempty"                   swaggerignore:"true" table:"expand"`
	StoragePath string    `json:"storagePath"` // StoragePath is the path to the CAR file inside the storage. If the StorageID is nil but StoragePath is not empty, it means the CAR file is stored at the local absolute path.

	// Association
	PreparationID uint32            `json:"preparationId"                                        table:"-"`
	Preparation   *Preparation      `gorm:"foreignKey:PreparationID;constraint:OnDelete:CASCADE" json:"preparation,omitempty" swaggerignore:"true" table:"-"`
	AttachmentID  *uint32           `json:"attachmentId"                                         table:"-"`
	Attachment    *SourceAttachment `gorm:"foreignKey:AttachmentID;constraint:OnDelete:CASCADE"  json:"attachment,omitempty"  swaggerignore:"true" table:"-"`
	JobID         *uint64           `json:"jobId,omitempty"                                      table:"-"`
	Job           *Job              `gorm:"foreignKey:JobID;constraint:OnDelete:SET NULL"        json:"job,omitempty"         swaggerignore:"true" table:"-"`

	// For Cbor marshalling
	_ struct{} `cbor:",toarray"                                                  json:"-"                    swaggerignore:"true"`
}

// CarBlock tells us the CIDs of all blocks inside a Car and the offset of those blocks.
// From this table we can determine how to get the block by CID from a Car.
// Or we can determine how to assemble a CAR file from blocks from File.
// It is also possible to get all Car that are associated with File.
// The index on CarID is used to find all blocks in a Car.
// The index on CID is used to find a specific block with CID.
type CarBlock struct {
	ID             uint64 `gorm:"primaryKey"                           json:"id"`
	CID            CID    `gorm:"index;column:cid;type:bytes;size:255" json:"cid"` // CID is the CID of the block.
	CarOffset      int64  `json:"carOffset"`                                       // Offset of the block in the Car
	CarBlockLength int32  `json:"carBlockLength"`                                  // Length of the block in the Car, including varint, CID and raw block
	Varint         []byte `json:"varint"`                                          // Varint is the varint that represents the length of the block and the CID.
	RawBlock       []byte `json:"rawBlock"`                                        // Raw block
	FileOffset     int64  `json:"fileOffset"`                                      // Offset of the block in the File

	// Internal Caching
	blockLength int32 // Block length in bytes

	// Associations
	CarID  uint32  `gorm:"index"                                         json:"carId"`
	Car    *Car    `gorm:"foreignKey:CarID;constraint:OnDelete:CASCADE"  json:"car,omitempty"  swaggerignore:"true"`
	FileID *uint64 `json:"fileId"`
	File   *File   `gorm:"foreignKey:FileID;constraint:OnDelete:CASCADE" json:"file,omitempty" swaggerignore:"true"`

	// For Cbor marshalling
	_ struct{} `cbor:",toarray"                                                  json:"-"                    swaggerignore:"true"`
}

// BlockLength computes and returns the length of the block data in bytes.
// If the block length has already been calculated and stored, it returns the stored value.
// Otherwise, it calculates the block length based on the raw block data if it is available,
// or by subtracting the CID byte length and varint length from the total CarBlock length.
//
// Returns:
// - An int32 representing the length of the block data in bytes.
func (c CarBlock) BlockLength() int32 {
	if c.blockLength != 0 {
		return c.blockLength
	}

	if c.RawBlock != nil {
		c.blockLength = int32(len(c.RawBlock))
	} else {
		c.blockLength = c.CarBlockLength - int32(cid.Cid(c.CID).ByteLen()) - int32(len(c.Varint))
	}

	return c.blockLength
}
