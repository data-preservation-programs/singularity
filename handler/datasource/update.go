package datasource

import (
	"context"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"github.com/rclone/rclone/fs"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Config map[string]string

// UpdateSourceHandler godoc
// @Summary Update the config options of a source
// @Tags Data Source
// @Param id path string true "Source ID"
// @Param config body AllConfig true "Config"
// @Success 200 {object} model.Source
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /source/{id} [patch]
func UpdateSourceHandler(
	db *gorm.DB,
	ctx context.Context,
	id string,
	deleteAfterExport *bool,
	rescanInterval *time.Duration,
	config Config,
) (*model.Source, *handler.Error) {
	var source model.Source
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid source id")
	}
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewBadRequestString("source not found")
	}
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}
	t := source.Type
	reg, err := fs.Find(t)
	if err != nil {
		return nil, handler.NewHandlerError(errors.New("invalid source type"))
	}
	if deleteAfterExport != nil {
		source.DeleteAfterExport = *deleteAfterExport
	}
	if rescanInterval != nil {
		source.ScanIntervalSeconds = uint64(rescanInterval.Seconds())
	}
	for key, value := range config {
		snake := lowerCamelToSnake(key)
		snake = strings.Replace(snake, "-", "_", -1)
		splitted := strings.SplitN(snake, "_", 2)
		if len(splitted) != 2 {
			return nil, handler.NewBadRequestString("invalid config key: " + key)
		}
		if splitted[0] != t {
			return nil, handler.NewBadRequestString("invalid config key for this data source: " + key)
		}
		name := splitted[1]
		_, err := underscore.Find(reg.Options, func(option fs.Option) bool {
			return option.Name == name
		})
		if err != nil {
			return nil, handler.NewBadRequestString("config key cannot be found for the data source: " + key)
		}
		source.Metadata[name] = value
	}

	h, err := datasource.DefaultHandlerResolver{}.Resolve(ctx, source)
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	_, err = h.List(ctx, "")
	if err != nil {
		return nil, handler.NewBadRequestError(err)
	}

	err = database.DoRetry(func() error {
		return db.Model(&source).Updates(map[string]interface{}{
			"metadata":              source.Metadata,
			"scan_interval_seconds": source.ScanIntervalSeconds,
			"delete_after_export":   source.DeleteAfterExport,
		}).Error
	})
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return &source, nil
}

func lowerCamelToSnake(s string) string {
	var result string
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i != 0 {
				result += "_"
			}
			result += strings.ToLower(string(r))
		} else {
			result += string(r)
		}
	}
	return result
}
