package storage

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

func UpdateStorageHandler(
	ctx context.Context,
	db *gorm.DB,
	name string,
	config map[string]string,
) (*model.Storage, error) {
	db = db.WithContext(ctx)

	if config == nil {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "nothing to update")
	}

	var storage model.Storage
	err := db.Where("name = ?", name).First(&storage).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "storage %s does not exist", name)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	backend, ok := storagesystem.BackendMap[storage.Type]
	if !ok {
		return nil, errors.Newf("storage type %s is not supported", storage.Type)
	}

	provider := storage.Config["provider"]
	if p, ok := config["provider"]; ok {
		provider = p
	}

	rcloneConfig := make(map[string]string)
	providerOptions, err := underscore.Find(backend.ProviderOptions, func(providerOption storagesystem.ProviderOptions) bool {
		return providerOption.Provider == provider
	})
	if err != nil {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "provider '%s' is not supported", provider)
	}

	for _, option := range providerOptions.Options {
		if option.Default != nil {
			rcloneConfig[option.Name] = fmt.Sprintf("%v", option.Default)
		}
	}

	for key, value := range storage.Config {
		rcloneConfig[key] = value
	}

	for key, value := range config {
		rcloneConfig[key] = value
	}

	storage.Config = rcloneConfig
	rclone, err := storagesystem.NewRCloneHandler(ctx, storage)
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrap(err, "creating rclone handler failed"))
	}

	_, err = rclone.List(ctx, "")
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrap(err, "listing the storage failed"))
	}

	err = database.DoRetry(ctx, func() error {
		return db.Model(&model.Storage{}).Where("id = ?", storage.ID).Update("config", storage.Config).Error
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &storage, err
}

// @Summary Update a storage connection
// @Tags Storage
// @Param name path string true "Name"
// @Param config body object true "Configuration"
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /storage/{name} [patch]
func _() {}
