package dataprep

import (
	"context"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/rjNemo/underscore"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

type DirEntry struct {
	Path         string    `json:"path"`
	IsDir        bool      `json:"isDir"`
	CID          string    `json:"cid"`
	FileVersions []Version `json:"fileVersions"`
}

type Version struct {
	ID           uint64    `json:"id"`
	CID          string    `json:"cid"`
	Hash         string    `json:"hash"`
	Size         int64     `json:"size"`
	LastModified time.Time `json:"lastModified"`
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
// - id: The ID of the preparation associated with the storage.
// - name: The name of the desired Storage record.
// - path: The directory path in the storage system to explore.
//
// Returns:
// - A slice of DirEntry structs representing the entries in the explored directory.
// - An error, if any occurred during the operation.
func ExploreHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
	name string,
	path string,
) ([]DirEntry, error) {
	db = db.WithContext(ctx)
	var storage model.Storage
	err := db.Where("name = ?", name).First(&storage).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "storage '%s' does not exist", name)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var source model.SourceAttachment
	err = db.Where("preparation_id = ? AND storage_id = ?", id, storage.ID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "source '%s' is not attached to preparation %d", name, id)
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

	for path, files := range filesByPath {
		slices.SortFunc(files, func(i, j model.File) bool {
			return i.LastModifiedNano > j.LastModifiedNano
		})
		entry := DirEntry{
			Path:  path,
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

	return result, nil
}

// @Summary Explore a directory in a prepared source storage
// @Tags Preparation
// @Accept json
// @Produce json
// @Param id path int true "Preparation ID"
// @Param name path string true "Source storage name"
// @Param path path string true "Directory path"
// @Success 200 {array} DirEntry
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /preparation/{id}/source/{name}/explore/{path} [get]
func _() {}
