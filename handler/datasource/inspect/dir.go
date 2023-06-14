package inspect

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type DirDetail struct {
	Current model.Directory
	Dirs    []model.Directory
	Items   []model.Item
}

// GetDirectoryHandler godoc
// @Summary Get all item details of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {array} model.Item
// @Failure 500 {object} handler.HTTPError
// @Router /source/{id}/items [get]
func GetDirectoryHandler(
	db *gorm.DB,
	id string,
	path string,
) (*DirDetail, *handler.Error) {
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
	dirId := source.RootDirectory().ID
	for _, segment := range segments {
		var subdir model.Directory
		err = db.Where("parent_id = ? AND name = ?", dirId, segment).First(&subdir).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, handler.NewBadRequestString("dir not found with given path")
		}
		if err != nil {
			return nil, handler.NewHandlerError(err)
		}
		dirId = subdir.ID
	}

	var dirs []model.Directory
	var items []model.Item
	err = db.Where("parent_id = ?", dirId).Find(&dirs).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}
	err = db.Where("directory_id = ?", dirId).Find(&items).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return &DirDetail{
		Current: *source.RootDirectory(),
		Dirs:    dirs,
		Items:   items,
	}, nil
}
