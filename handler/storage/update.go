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

// UpdateStorageHandler updates the configuration of a given storage system.
//
// This method performs the following operations:
//   - Retrieves the storage system with the specified name.
//   - Merges the new configuration with the current configuration.
//   - Ensures the provided configuration is valid by checking the storage type's backend compatibility.
//   - Initializes an RCloneHandler with the merged configuration to validate the config against the actual storage backend.
//   - Updates the storage system's configuration in the database.
//
// Parameters:
//   - ctx: A context.Context for request-scoped values, cancellation signals, and deadlines.
//   - db: A *gorm.DB instance for database interactions.
//   - name: The name of the storage system to be updated.
//   - config: A map representing the new configuration values to be merged with the existing configuration.
//
// Returns:
//   - A pointer to the updated storage model.
//   - An error if there's an issue during the update process, otherwise nil.
//
// Note:
//
//	If the specified storage system does not exist, an error will be returned.
//	The method also ensures that the merged configuration is valid and the storage backend is accessible
//	using RCloneHandler. If the new configuration is invalid, or if the backend is inaccessible, an error is returned.
func (DefaultHandler) UpdateStorageHandler(
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
	err := storage.FindByIDOrName(db, name)
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

// @ID UpdateStorage
// @Summary Update a storage connection
// @Tags Storage
// @Param name path string true "Storage ID or name"
// @Param config body map[string]string true "Configuration"
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /storage/{name} [patch]
func _() {}
