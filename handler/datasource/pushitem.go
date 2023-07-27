package datasource

import (
	"context"
	"errors"
	"fmt"

	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/datasetworker"
	fs2 "github.com/rclone/rclone/fs"
	"gorm.io/gorm"
)

type ItemInfo struct {
	Path string `json:"path"` // Path to the new item, relative to the source
}

func PushItemHandler(
	db *gorm.DB,
	ctx context.Context,
	datasourceHandlerResolver datasource.HandlerResolver,
	sourceID uint32,
	itemInfo ItemInfo,
) (*model.Item, error) {
	return pushItemHandler(db, ctx, datasourceHandlerResolver, sourceID, itemInfo)
}

// @Summary Push an item to be queued
// @Description Tells Singularity that something is ready to be grabbed for data preparation
// @Tags Data Source
// @Accept json
// @Produce json
// @Param item body ItemInfo true "Item"
// @Success 201 {object} model.Item
// @Failure 400 {string} string "Bad Request"
// @Failure 409 {string} string "Item already exists"
// @Failure 500 {string} string "Internal Server Error"
// @Router /source/{id}/push [post]
func pushItemHandler(
	db *gorm.DB,
	ctx context.Context,
	datasourceHandlerResolver datasource.HandlerResolver,
	sourceID uint32,
	itemInfo ItemInfo,
) (*model.Item, error) {
	var source model.Source
	err := db.Preload("Dataset").Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr(fmt.Sprintf("source %d not found.", sourceID))
	}
	if err != nil {
		return nil, err
	}

	dsHandler, err := datasourceHandlerResolver.Resolve(ctx, source)
	if err != nil {
		return nil, err
	}

	entry, err := dsHandler.Check(ctx, itemInfo.Path)
	if err != nil {
		return nil, handler.InvalidParameterError{Err: err}
	}

	obj, ok := entry.(fs2.ObjectInfo)
	if !ok {
		return nil, handler.NewInvalidParameterErr(fmt.Sprintf("%s is not an object", itemInfo.Path))
	}

	item, itemParts, err := datasetworker.PushItem(ctx, db, obj, source, *source.Dataset, map[string]uint64{})

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, handler.NewDuplicateRecordError(fmt.Sprintf("%s already exists", obj.Remote()))
	}

	item.ItemParts = itemParts

	return item, nil
}
