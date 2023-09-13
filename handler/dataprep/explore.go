package dataprep

import (
	"context"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

type DirEntry struct {
	Path         string    `json:"path"`
	IsDir        bool      `json:"isDir"`
	CID          string    `json:"cid"`
	FileVersions []Version `json:"fileVersions" table:"verbose;expand"`
}

type ExploreResult struct {
	Path       string     `json:"path"`
	CID        string     `json:"cid"`
	SubEntries []DirEntry `json:"subEntries" table:"expand"`
}

type Version struct {
	ID           model.FileID `json:"id"`
	CID          string       `json:"cid"`
	Hash         string       `json:"hash"`
	Size         int64        `json:"size"`
	LastModified time.Time    `json:"lastModified" table:"format:2006-01-02 15:04:05"`
}

// ExploreHandler fetches directory entries (files and sub-directories) associated with a specific preparation
// in a given storage system and directory path. The function retrieves information from a local database
// rather than directly exploring the remote storage, making use of the stored relationships between files,
// directories, and storage systems.
//
// This function starts by fetching the desired Storage record based on the provided name. It then fetches the
// associated SourceAttachment record which connects a preparation to a storage. Using the RootDirectoryID method
// of the source, it retrieves the root directory's ID and navigates to the desired directory by iterating
// through the path segments. Once at the desired directory, it fetches the contained directories and files,
// constructing a result list from the gathered data.
//
// Parameters:
// - ctx: The context for database transactions and other operations.
// - db: A pointer to the gorm.DB instance representing the database connection.
// - id: The ID or name of the preparation associated with the storage.
// - name: The ID or name of the desired Storage record.
// - path: The directory path in the storage system to explore.
//
// Returns:
// - ExploreResult struct representing the entries in the explored directory.
// - An error, if any occurred during the operation.
func (DefaultHandler) ExploreHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
	name string,
	path string,
) (*ExploreResult, error) {
	db = db.WithContext(ctx)

	var source model.SourceAttachment
	err := source.FindByPreparationAndSource(db, id, name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "source '%s' is not attached to preparation %s", name, id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dirID, err := source.RootDirectoryID(ctx, db)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	segments := underscore.Filter(strings.Split(path, "/"), func(p string) bool { return p != "" })
	path = strings.Join(segments, "/")
	for _, segment := range segments {
		var dir model.Directory
		err = db.Where("parent_id = ? AND name = ?", dirID, segment).First(&dir).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(handlererror.ErrNotFound, "directory '%s' does not exist", path)
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		dirID = dir.ID
	}

	var current model.Directory
	err = db.First(&current, dirID).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []DirEntry
	var dirs []model.Directory
	err = db.Where("parent_id = ?", dirID).Find(&dirs).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, dir := range dirs {
		result = append(result, DirEntry{
			Path:  path + "/" + dir.Name,
			IsDir: true,
			CID:   dir.CID.String(),
		})
	}

	var files []model.File
	err = db.Where("directory_id = ?", dirID).Find(&files).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	filesByPath := make(map[string][]model.File)
	for _, file := range files {
		filesByPath[file.Path] = append(filesByPath[file.Path], file)
	}

	for filePath, files := range filesByPath {
		entry := DirEntry{
			Path:  filePath,
			IsDir: false,
			CID:   files[0].CID.String(),
		}
		for _, file := range files {
			version := Version{
				ID:           file.ID,
				CID:          file.CID.String(),
				Hash:         file.Hash,
				Size:         file.Size,
				LastModified: time.Unix(0, file.LastModifiedNano),
			}
			entry.FileVersions = append(entry.FileVersions, version)
		}
		result = append(result, entry)
	}

	return &ExploreResult{
		Path:       path,
		CID:        current.CID.String(),
		SubEntries: result,
	}, nil
}

// @ID ExplorePreparation
// @Summary Explore a directory in a prepared source storage
// @Tags Preparation
// @Accept json
// @Produce json
// @Param id path string true "Preparation ID or name"
// @Param name path string true "Source storage ID or name"
// @Param path path string true "Directory path"
// @Success 200 {object} ExploreResult
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/source/{name}/explore/{path} [get]
func _() {}
