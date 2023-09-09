package dataprep

import (
	"context"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler"
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
	ID           uint64    `json:"id"`
	CID          string    `json:"cid"`
	Hash         string    `json:"hash"`
	Size         int64     `json:"size"`
	LastModified time.Time `json:"lastModified" table:"format:2006-01-02 15:04:05"`
}

type ExploreRequest struct {
	Path string `json:"path"`
}

// ExploreHandler manages the exploration of directory structures associated with a source attached to a preparation.
//
// Given a preparation ID and source name, and a path within the source's directory structure,
// this handler fetches the directory's content. The result includes directory entries, such as subdirectories
// and files, along with their respective CIDs (Content Identifiers) and other metadata.
//
// If the specified source is not found to be attached to the preparation, or if a directory does not exist
// at the provided path, the handler returns a not-found error. All database operations are handled with care
// to produce descriptive errors in case of any issues.
//
// Parameters:
//   - ctx: The context for managing timeouts and cancellation.
//   - request: A handler request containing the ExploreRequest payload.
//     The Params field should include the preparation ID followed by the source name.
//   - dep: Contains the handler's dependencies, such as the gorm.DB instance.
//
// Returns:
//   - An ExploreResult which represents the contents of the directory at the specified path.
//     This includes subdirectories, files, their CIDs, and other metadata.
//   - An error if any issues occur during the operation, such as database errors or not-found situations.
func (DefaultHandler) ExploreHandler(
	ctx context.Context,
	request handler.Request[ExploreRequest],
	dep handler.Dependency,
) (*ExploreResult, error) {
	db := dep.DB.WithContext(ctx)
	id := request.Params[0]
	name := request.Params[1]
	path := request.Payload.Path
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
// @Param request body ExploreRequest true "Explore Request"
// @Success 200 {object} ExploreResult
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/source/{name}/explore [post]
func _() {}
