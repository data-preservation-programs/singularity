package storage

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/rclone/rclone/fs"
	"gorm.io/gorm"
)

type DirEntry struct {
	Path         string    `json:"path"`
	LastModified time.Time `json:"lastModified" table:"format:2006-01-02 15:04:05"`
	Size         int64     `json:"size"`
	IsDir        bool      `json:"isDir"`
	DirID        string    `json:"dirId"        table:"verbose"`
	NumItems     int64     `json:"numItems"     table:"verbose"`
	Hash         string    `json:"hash"         table:"verbose"`
}

// ExploreHandler fetches directory entries (files and sub-directories) for a given storage system
// and directory path using rclone. It returns a list of these entries, allowing you to view
// the contents of a directory in a remote storage system.
//
// This function starts by fetching the desired Storage record based on the provided name.
// It then initializes an RCloneHandler for this storage. Using rclone, it lists the contents of
// the given directory path. Each entry is processed to determine if it is a file or a directory
// and the relevant information for each entry is returned.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - name: The name of the desired Storage record.
//   - path: The directory path in the storage system to explore.
//
// Returns:
//   - A slice of DirEntry structs representing the entries in the explored directory.
//   - An error, if any occurred during the operation.
func (DefaultHandler) ExploreHandler(
	ctx context.Context,
	db *gorm.DB,
	name string,
	path string,
) ([]DirEntry, error) {
	db = db.WithContext(ctx)
	var storage model.Storage
	err := storage.FindByIDOrName(db, name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "storage %s not found", name)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rclone, err := storagesystem.NewRCloneHandler(ctx, storage)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	entries, err := rclone.List(ctx, path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list %s", path)
	}

	var result []DirEntry
	for _, entry := range entries {
		switch v := entry.(type) {
		case fs.Directory:
			result = append(result, DirEntry{
				Path:         v.Remote(),
				LastModified: v.ModTime(ctx),
				Size:         -1,
				IsDir:        true,
				DirID:        v.ID(),
				NumItems:     v.Items(),
			})
		case fs.ObjectInfo:
			hashValue, err := storagesystem.GetHash(ctx, v)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			result = append(result, DirEntry{
				Path:         v.Remote(),
				LastModified: v.ModTime(ctx),
				Size:         v.Size(),
				IsDir:        false,
				Hash:         hashValue,
			})
		}
	}

	return result, nil
}

// @ID ExploreStorage
// @Summary Explore directory entries in a storage system
// @Tags Storage
// @Accept json
// @Produce json
// @Param name path string true "Storage ID or name"
// @Param path path string true "Path in the storage system to explore"
// @Success 200 {array} DirEntry
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /storage/{name}/explore/{path} [get]
func _() {}
