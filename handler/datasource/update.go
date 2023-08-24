package datasource

import (
	"context"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/rclone/rclone/fs"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

type Config map[string]any

func UpdateSourceHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
	config Config,
) (*model.Source, error) {
	return updateSourceHandler(ctx, db.WithContext(ctx), id, config)
}

// @Summary Update the config options of a source
// @Tags Data Source
// @Produce json
// @Accept json
// @Param id path string true "Source ID"
// @Param config body AllConfig true "Config"
// @Success 200 {object} model.Source
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /source/{id} [patch]
func updateSourceHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
	config Config,
) (*model.Source, error) {
	var source model.Source
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handlererror.NewInvalidParameterErr("invalid source id")
	}
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handlererror.NewInvalidParameterErr("source not found")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t := source.Type
	reg, err := fs.Find(t)
	if err != nil {
		return nil, errors.New("invalid source type")
	}
	value, ok := config["deleteAfterExport"]
	if ok {
		v, ok := value.(bool)
		if !ok {
			return nil, handlererror.NewInvalidParameterErr("invalid deleteAfterExport value")
		}
		source.DeleteAfterExport = v
	}
	value, ok = config["rescanInterval"]
	if ok {
		v, ok := value.(string)
		if !ok {
			return nil, handlererror.NewInvalidParameterErr("invalid rescanInterval value")
		}
		d, err := time.ParseDuration(v)
		if err != nil {
			return nil, handlererror.NewInvalidParameterErr("invalid rescanInterval value")
		}
		source.ScanIntervalSeconds = uint64(d.Seconds())
	}
	delete(config, "deleteAfterExport")
	delete(config, "rescanInterval")
	delete(config, "id")
	for key, value := range config {
		v, ok := value.(string)
		if !ok {
			return nil, handlererror.NewInvalidParameterErr("invalid config value: " + key)
		}
		snake := lowerCamelToSnake(key)
		snake = strings.ReplaceAll(snake, "-", "_")
		splitted := strings.SplitN(snake, "_", 2)
		if len(splitted) != 2 {
			return nil, handlererror.NewInvalidParameterErr("invalid config key: " + key)
		}
		if splitted[0] != t {
			return nil, handlererror.NewInvalidParameterErr("invalid config key for this data source: " + key)
		}
		name := splitted[1]
		_, err := underscore.Find(reg.Options, func(option fs.Option) bool {
			return option.Name == name
		})
		if err != nil {
			return nil, handlererror.NewInvalidParameterErr("config key cannot be found for the data source: " + key)
		}
		source.Metadata[name] = v
	}

	h, err := storagesystem.DefaultHandlerResolver{}.Resolve(ctx, source)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	_, err = h.List(ctx, "")
	if err != nil {
		return nil, handlererror.InvalidParameterError{Err: err}
	}

	err = database.DoRetry(ctx, func() error {
		return db.Model(&source).Updates(map[string]any{
			"metadata":              source.Metadata,
			"scan_interval_seconds": source.ScanIntervalSeconds,
			"delete_after_export":   source.DeleteAfterExport,
		}).Error
	})
	if err != nil {
		return nil, errors.WithStack(err)
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
