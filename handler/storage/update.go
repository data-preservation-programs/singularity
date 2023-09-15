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

type UpdateRequest struct {
	Config       map[string]string  `json:"config"`
	ClientConfig model.ClientConfig `json:"clientConfig"`
}

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
	request UpdateRequest,
) (*model.Storage, error) {
	db = db.WithContext(ctx)

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
	if p, ok := request.Config["provider"]; ok {
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

	for key, value := range request.Config {
		rcloneConfig[key] = value
	}

	storage.Config = rcloneConfig
	overrideStorageWithClientConfig(&storage, request.ClientConfig)
	rclone, err := storagesystem.NewRCloneHandler(ctx, storage)
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrap(err, "creating rclone handler failed"))
	}

	_, err = rclone.List(ctx, "")
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrap(err, "listing the storage failed"))
	}

	err = database.DoRetry(ctx, func() error {
		return db.Model(&model.Storage{}).Where("id = ?", storage.ID).Updates(map[string]any{
			"config":        storage.Config,
			"client_config": storage.ClientConfig,
		}).Error
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &storage, err
}

func overrideStorageWithClientConfig(storage *model.Storage, config model.ClientConfig) {
	if config.ConnectTimeout != nil {
		storage.ClientConfig.ConnectTimeout = config.ConnectTimeout
	}
	if config.Timeout != nil {
		storage.ClientConfig.Timeout = config.Timeout
	}
	if config.ExpectContinueTimeout != nil {
		storage.ClientConfig.ExpectContinueTimeout = config.ExpectContinueTimeout
	}
	if config.InsecureSkipVerify != nil {
		storage.ClientConfig.InsecureSkipVerify = config.InsecureSkipVerify
	}
	if config.NoGzip != nil {
		storage.ClientConfig.NoGzip = config.NoGzip
	}
	if config.UserAgent != nil {
		storage.ClientConfig.UserAgent = config.UserAgent
		if *storage.ClientConfig.UserAgent == "" {
			storage.ClientConfig.UserAgent = nil
		}
	}
	if len(config.CaCert) > 0 {
		storage.ClientConfig.CaCert = config.CaCert
		if storage.ClientConfig.CaCert[0] == "" {
			storage.ClientConfig.CaCert = nil
		}
	}
	if config.ClientCert != nil {
		storage.ClientConfig.ClientCert = config.ClientCert
		if *storage.ClientConfig.ClientCert == "" {
			storage.ClientConfig.ClientCert = nil
		}
	}
	if config.ClientKey != nil {
		storage.ClientConfig.ClientKey = config.ClientKey
		if *storage.ClientConfig.ClientKey == "" {
			storage.ClientConfig.ClientKey = nil
		}
	}
	if config.DisableHTTP2 != nil {
		storage.ClientConfig.DisableHTTP2 = config.DisableHTTP2
	}
	if config.DisableHTTPKeepAlives != nil {
		storage.ClientConfig.DisableHTTPKeepAlives = config.DisableHTTPKeepAlives
	}
	if config.Headers != nil {
		for key, value := range config.Headers {
			if key == "" && value == "" {
				storage.ClientConfig.Headers = make(map[string]string)
				break
			}
			if value == "" {
				delete(storage.ClientConfig.Headers, key)
			} else {
				storage.ClientConfig.Headers[key] = value
			}
		}
	}
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
