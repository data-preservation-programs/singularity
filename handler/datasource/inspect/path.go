package inspect

import (
	"strconv"
	"strings"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

type DirDetail struct {
	Current model.Directory
	Dirs    []model.Directory
	Items   []model.Item
}

type GetPathRequest struct {
	Path string `json:"path"`
}

func GetPathHandler(
	db *gorm.DB,
	id string,
	request GetPathRequest,
) (*DirDetail, error) {
	return getPathHandler(db, id, request)
}

// @Summary Get all item details inside a data source path
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Param request body GetPathRequest true "GetPathRequest"
// @Success 200 {object} DirDetail
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /source/{id}/path [get]
func getPathHandler(
	db *gorm.DB,
	id string,
	request GetPathRequest,
) (*DirDetail, error) {
	path := request.Path
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid source id")
	}
	var source model.Source
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewBadRequestString("source not found")
	}
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	segments := underscore.Filter(strings.Split(path, "/"), func(s string) bool { return s != "" })
	err = source.LoadRootDirectory(db)
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}
	dirID := source.RootDirectory().ID
	var subdir model.Directory
	for i, segment := range segments {
		err = db.Where("parent_id = ? AND name = ?", dirID, segment).First(&subdir).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if i == len(segments)-1 {
				var items []model.Item
				err = db.Where("directory_id = ? AND path = ?", dirID, strings.Join(segments, "/")).Find(&items).Error
				if err != nil {
					return nil, handler.NewHandlerError(err)
				}
				if len(items) == 0 {
					return nil, handler.NewBadRequestString("entry not found with given path")
				}
				return &DirDetail{
					Current: subdir,
					Dirs:    nil,
					Items:   items,
				}, nil
			}
			return nil, handler.NewBadRequestString("entry not found with given path")
		}
		if err != nil {
			return nil, handler.NewHandlerError(err)
		}
		dirID = subdir.ID
	}

	var current model.Directory
	var dirs []model.Directory
	var items []model.Item
	err = db.Where("id = ?", dirID).First(&current).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}
	err = db.Where("parent_id = ?", dirID).Find(&dirs).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}
	err = db.Where("directory_id = ?", dirID).Find(&items).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return &DirDetail{
		Current: current,
		Dirs:    dirs,
		Items:   items,
	}, nil
}
