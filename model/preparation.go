package model

import (
	"context"
	"strconv"
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

type PreparationID uint32

// Preparation is a data preparation definition that can attach multiple source storages and up to one output storage.
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

	// Associations
	Wallets        []Wallet  `gorm:"many2many:wallet_assignments"                             json:"wallets,omitempty"        swaggerignore:"true"                   table:"expand"`
	SourceStorages []Storage `gorm:"many2many:source_attachments;constraint:OnDelete:CASCADE" json:"sourceStorages,omitempty" table:"expand;header:Source Storages:"`
	OutputStorages []Storage `gorm:"many2many:output_attachments;constraint:OnDelete:CASCADE" json:"outputStorages,omitempty" table:"expand;header:Output Storages:"`
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
	return attachments, errors.WithStack(err)
}

type StorageID uint32

// Storage is a storage system definition that can be used as either source or output of a Preparation.
type Storage struct {
	ID           StorageID    `cbor:"-"                    gorm:"primaryKey" json:"id"`
	Name         string       `cbor:"-"                    gorm:"unique"     json:"name"`
	CreatedAt    time.Time    `cbor:"-"                    json:"createdAt"  table:"verbose;format:2006-01-02 15:04:05"`
	UpdatedAt    time.Time    `cbor:"-"                    json:"updatedAt"  table:"verbose;format:2006-01-02 15:04:05"`
	Type         string       `cbor:"1,keyasint,omitempty" json:"type"`
	Path         string       `cbor:"2,keyasint,omitempty" json:"path"`                                                                  // Path is the path to the storage root.
	Config       ConfigMap    `cbor:"3,keyasint,omitempty" gorm:"type:JSON"  json:"config"                              table:"verbose"` // Config is a map of key-value pairs that can be used to store RClone options.
	ClientConfig ClientConfig `cbor:"4,keyasint,omitempty" gorm:"type:JSON"  json:"clientConfig"                        table:"verbose"` // ClientConfig is the HTTP configuration for the storage, if applicable.

	// Associations
	PreparationsAsSource []Preparation `cbor:"-" gorm:"many2many:source_attachments;constraint:OnDelete:CASCADE" json:"preparationsAsSource,omitempty" table:"expand;header:As Source: "`
	PreparationsAsOutput []Preparation `cbor:"-" gorm:"many2many:output_attachments;constraint:OnDelete:CASCADE" json:"preparationsAsOutput,omitempty" table:"expand;header:As Output: "`
}

func (s *Storage) FindByIDOrName(db *gorm.DB, name string, preloads ...string) error {
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

type SourceAttachmentID uint32

// SourceAttachment is a link between a Preparation and a Storage that is used as a source.
type SourceAttachment struct {
	ID SourceAttachmentID `gorm:"primaryKey" json:"id"`

	// Associations
	PreparationID PreparationID `gorm:"uniqueIndex:prep_source"                              json:"preparationId"`
	Preparation   *Preparation  `gorm:"foreignKey:PreparationID;constraint:OnDelete:CASCADE" json:"preparation,omitempty" swaggerignore:"true"`
	StorageID     StorageID     `gorm:"uniqueIndex:prep_source"                              json:"storageId"`
	Storage       *Storage      `gorm:"foreignKey:StorageID;constraint:OnDelete:CASCADE"     json:"storage,omitempty"     swaggerignore:"true"`
}

func (s *SourceAttachment) FindByPreparationAndSource(db *gorm.DB, preparation string, source string) error {
	var prep Preparation
	err := prep.FindByIDOrName(db, preparation)
	if err != nil {
		return errors.WithStack(err)
	}

	var storage Storage
	err = storage.FindByIDOrName(db, source)
	if err != nil {
		return errors.WithStack(err)
	}

	err = db.Where("preparation_id = ? AND storage_id = ?", prep.ID, storage.ID).First(s).Error
	if err != nil {
		return errors.WithStack(err)
	}

	s.Preparation = &prep
	s.Storage = &storage
	return nil
}

func (s *SourceAttachment) RootDirectoryCID(ctx context.Context, db *gorm.DB) (cid.Cid, error) {
	db = db.WithContext(ctx)
	var root Directory
	err := db.Select("cid").Where("attachment_id = ? AND parent_id is null", s.ID).First(&root).Error
	return cid.Cid(root.CID), errors.WithStack(err)
}

func (s *SourceAttachment) RootDirectoryID(ctx context.Context, db *gorm.DB) (DirectoryID, error) {
	db = db.WithContext(ctx)
	var root Directory
	err := db.Select("id").Where("attachment_id = ? AND parent_id is null", s.ID).First(&root).Error
	return root.ID, errors.WithStack(err)
}

type OutputAttachmentID uint32

// OutputAttachment is a link between a Preparation and a Storage that is used as an output.
type OutputAttachment struct {
	ID OutputAttachmentID `gorm:"primaryKey" json:"id"`

	// Associations
	PreparationID PreparationID `gorm:"uniqueIndex:prep_output"                              json:"preparationId"`
	Preparation   *Preparation  `gorm:"foreignKey:PreparationID;constraint:OnDelete:CASCADE" json:"preparation,omitempty" swaggerignore:"true"`
	StorageID     StorageID     `gorm:"uniqueIndex:prep_output"                              json:"storageId"`
	Storage       *Storage      `gorm:"foreignKey:StorageID;constraint:OnDelete:CASCADE"     json:"storage,omitempty"     swaggerignore:"true"`
}

type JobID uint64

// Job is a job that is executed by a worker.
// The composite index on Type and State is used to find jobs that are ready to be executed.
type Job struct {
	ID              JobID    `gorm:"primaryKey"           json:"id"`
	Type            JobType  `gorm:"index:job_type_state" json:"type"`
	State           JobState `gorm:"index:job_type_state" json:"state"`
	ErrorMessage    string   `json:"errorMessage"`
	ErrorStackTrace string   `json:"errorStackTrace"      table:"verbose"`

	// Associations
	WorkerID     *string            `gorm:"size:63"                                                        json:"workerId,omitempty"`
	Worker       *Worker            `gorm:"foreignKey:WorkerID;references:ID;constraint:OnDelete:SET NULL" json:"worker,omitempty"     swaggerignore:"true" table:"verbose;expand"`
	AttachmentID SourceAttachmentID `json:"attachmentId"                                                   table:"verbose"`
	Attachment   *SourceAttachment  `gorm:"foreignKey:AttachmentID;constraint:OnDelete:CASCADE"            json:"attachment,omitempty" swaggerignore:"true" table:"expand"`
	FileRanges   []FileRange        `gorm:"foreignKey:JobID;constraint:OnDelete:SET NULL"                  json:"fileRanges,omitempty" swaggerignore:"true" table:"-"`
}

type FileID uint64

// File makes a reference to the source storage file, e.g., a local file.
// The index on Path is used as part of scanning to find existing file and add new versions.
// The index on DirectoryID is used to find all files in a directory.
type File struct {
	ID               FileID `cbor:"1,keyasint,omitempty" gorm:"primaryKey"                     json:"id"`
	CID              CID    `cbor:"-"                    gorm:"column:cid;type:bytes;size:255" json:"cid"  swaggertype:"string"` // CID is the CID of the file.
	Path             string `cbor:"2,keyasint,omitempty" gorm:"index"                          json:"path"`                      // Path is the relative path to the file inside the storage.
	Hash             string `cbor:"3,keyasint,omitempty" json:"hash"`                                                            // Hash is the hash of the file.
	Size             int64  `cbor:"4,keyasint,omitempty" json:"size"`                                                            // Size is the size of the file in bytes.
	LastModifiedNano int64  `cbor:"5,keyasint,omitempty" json:"lastModifiedNano"`

	// Associations
	AttachmentID SourceAttachmentID `cbor:"-" json:"attachmentId"`
	Attachment   *SourceAttachment  `cbor:"-" gorm:"foreignKey:AttachmentID;constraint:OnDelete:CASCADE" json:"attachment,omitempty" swaggerignore:"true"`
	DirectoryID  *DirectoryID       `cbor:"-" gorm:"index"                                               json:"directoryId"`
	Directory    *Directory         `cbor:"-" gorm:"foreignKey:DirectoryID;constraint:OnDelete:CASCADE"  json:"directory,omitempty"  swaggerignore:"true"`
	FileRanges   []FileRange        `cbor:"-" gorm:"constraint:OnDelete:CASCADE"                         json:"fileRanges,omitempty"`
}

func (i File) FileName() string {
	return i.Path[strings.LastIndex(i.Path, "/")+1:]
}

type DirectoryID uint64

// Directory is a link between parent and child directories.
// The index on AttachmentID and ParentID is used to find all root directories, as well as all directories in a directory.
type Directory struct {
	ID       DirectoryID `gorm:"primaryKey"            json:"id"`
	CID      CID         `gorm:"column:cid;type:bytes" json:"cid" swaggertype:"string"` // CID is the CID of the directory.
	Data     []byte      `gorm:"column:data"           json:"-"   swaggerignore:"true"` // Data is the serialized directory data.
	Name     string      `json:"name"`                                                  // Name is the name of the directory.
	Exported bool        `json:"exported"`                                              // Exported is a flag that indicates whether the directory has been exported to the DAG.

	// Associations
	AttachmentID SourceAttachmentID `gorm:"index:directory_source_parent"                       json:"attachmentId"`
	Attachment   *SourceAttachment  `gorm:"foreignKey:AttachmentID;constraint:OnDelete:CASCADE" json:"attachment,omitempty" swaggerignore:"true"`
	ParentID     *DirectoryID       `gorm:"index:directory_source_parent"                       json:"parentId"`
	Parent       *Directory         `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE"     json:"parent,omitempty"     swaggerignore:"true"`
}

type FileRangeID uint64

// FileRange is a range of bytes inside File.
// The index on FileID is used to find all FileRange in a file.
// The index on JobID is used to find all FileRange in a job.
type FileRange struct {
	ID     FileRangeID `gorm:"primaryKey"            json:"id"`
	Offset int64       `json:"offset"`                                                // Offset is the offset of the range inside the file.
	Length int64       `json:"length"`                                                // Length is the length of the range in bytes.
	CID    CID         `gorm:"column:cid;type:bytes" json:"cid" swaggertype:"string"` // CID is the CID of the range.

	// Associations
	JobID  *JobID `gorm:"index"                                         json:"jobId"`
	Job    *Job   `gorm:"foreignKey:JobID;constraint:OnDelete:SET NULL" json:"job,omitempty"  swaggerignore:"true"`
	FileID FileID `gorm:"index"                                         json:"fileId"`
	File   *File  `gorm:"foreignKey:FileID;constraint:OnDelete:CASCADE" json:"file,omitempty" swaggerignore:"true"`
}

type CarID uint32

// Car makes a reference to a CAR file that has been potentially exported to the disk.
// In the case of inline preparation, the path may be empty so the Car should be constructed
// on the fly using CarBlock.
// The index on PieceCID is to find all CARs that can matches the PieceCID
type Car struct {
	ID          CarID      `cbor:"-"                    gorm:"primaryKey"                                        json:"id"                                  table:"verbose"`
	CreatedAt   time.Time  `cbor:"-"                    json:"createdAt"                                         table:"verbose;format:2006-01-02 15:04:05"`
	PieceCID    CID        `cbor:"1,keyasint,omitempty" gorm:"column:piece_cid;index;type:bytes;size:255"        json:"pieceCid"                            swaggertype:"string"`
	PieceSize   int64      `cbor:"2,keyasint,omitempty" json:"pieceSize"`
	RootCID     CID        `cbor:"3,keyasint,omitempty" gorm:"column:root_cid;type:bytes"                        json:"rootCid"                             swaggertype:"string"`
	FileSize    int64      `cbor:"4,keyasint,omitempty" json:"fileSize"`
	StorageID   *StorageID `cbor:"-"                    json:"storageId"                                         table:"verbose"`
	Storage     *Storage   `cbor:"-"                    gorm:"foreignKey:StorageID;constraint:OnDelete:SET NULL" json:"storage,omitempty"                   swaggerignore:"true" table:"expand"`
	StoragePath string     `cbor:"-"                    json:"storagePath"` // StoragePath is the path to the CAR file inside the storage. If the StorageID is nil but StoragePath is not empty, it means the CAR file is stored at the local absolute path.
	NumOfFiles  int64      `cbor:"-"                    json:"numOfFiles"                                        table:"verbose"`

	// Association
	PreparationID PreparationID       `cbor:"-" json:"preparationId"                                        table:"-"`
	Preparation   *Preparation        `cbor:"-" gorm:"foreignKey:PreparationID;constraint:OnDelete:CASCADE" json:"preparation,omitempty" swaggerignore:"true" table:"-"`
	AttachmentID  *SourceAttachmentID `cbor:"-" json:"attachmentId"                                         table:"-"`
	Attachment    *SourceAttachment   `cbor:"-" gorm:"foreignKey:AttachmentID;constraint:OnDelete:CASCADE"  json:"attachment,omitempty"  swaggerignore:"true" table:"-"`
	JobID         *JobID              `cbor:"-" json:"jobId,omitempty"                                      table:"-"`
	Job           *Job                `cbor:"-" gorm:"foreignKey:JobID;constraint:OnDelete:SET NULL"        json:"job,omitempty"         swaggerignore:"true" table:"-"`
}

type CarBlockID uint64

// CarBlock tells us the CIDs of all blocks inside a Car and the offset of those blocks.
// From this table we can determine how to get the block by CID from a Car.
// Or we can determine how to assemble a CAR file from blocks from File.
// It is also possible to get all Car that are associated with File.
// The index on CarID is used to find all blocks in a Car.
// The index on CID is used to find a specific block with CID.
type CarBlock struct {
	ID             CarBlockID `cbor:"-"                    gorm:"primaryKey"                           json:"id"`
	CID            CID        `cbor:"1,keyasint,omitempty" gorm:"index;column:cid;type:bytes;size:255" json:"cid" swaggertype:"string"` // CID is the CID of the block.
	CarOffset      int64      `cbor:"2,keyasint,omitempty" json:"carOffset"`                                                            // Offset of the block in the Car
	CarBlockLength int32      `cbor:"3,keyasint,omitempty" json:"carBlockLength"`                                                       // Length of the block in the Car, including varint, CID and raw block
	Varint         []byte     `cbor:"4,keyasint,omitempty" json:"varint"`                                                               // Varint is the varint that represents the length of the block and the CID.
	RawBlock       []byte     `cbor:"5,keyasint,omitempty" json:"rawBlock"`                                                             // Raw block
	FileOffset     int64      `cbor:"6,keyasint,omitempty" json:"fileOffset"`                                                           // Offset of the block in the File

	// Internal Caching
	blockLength int32 // Block length in bytes

	// Associations
	CarID  CarID   `cbor:"-"                    gorm:"index"                                         json:"carId"`
	Car    *Car    `cbor:"-"                    gorm:"foreignKey:CarID;constraint:OnDelete:CASCADE"  json:"car,omitempty"  swaggerignore:"true"`
	FileID *FileID `cbor:"7,keyasint,omitempty" json:"fileId"`
	File   *File   `cbor:"-"                    gorm:"foreignKey:FileID;constraint:OnDelete:CASCADE" json:"file,omitempty" swaggerignore:"true"`
}

// BlockLength computes and returns the length of the block data in bytes.
// If the block length has already been calculated and stored, it returns the stored value.
// Otherwise, it calculates the block length based on the raw block data if it is available,
// or by subtracting the CID byte length and varint length from the total CarBlock length.
//
// Returns:
//   - An int32 representing the length of the block data in bytes.
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
