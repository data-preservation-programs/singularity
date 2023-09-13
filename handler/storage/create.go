package storage

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

type CreateRequest struct {
	Provider string            `json:"provider"`
	Name     string            `binding:"required" json:"name"`
	Path     string            `json:"path"`
	Config   map[string]string `json:"config"`
}

// CreateStorageHandler initializes a new storage using the provided configurations
// and attempts to create a connection to the storage to ensure it is valid. If successful,
// it creates a new storage entry in the database.
//
// Parameters:
// - ctx: The context for database transactions and other operations.
// - db: A pointer to the gorm.DB instance representing the database connection.
// - storageType: The type of storage system (e.g., S3, GCS, local).
// - provider: The provider for the storage system (e.g., AWS, Google, etc.).
// - name: A unique name to represent the storage in the database.
// - path: The path or endpoint to access the storage.
// - config: A map containing the configuration key-value pairs required by the storage backend.
//
// Returns:
// - A pointer to the newly created Storage model if successful.
// - An error, if any occurred during the operation.
func (DefaultHandler) CreateStorageHandler(
	ctx context.Context,
	db *gorm.DB,
	storageType string,
	request CreateRequest,
) (*model.Storage, error) {
	db = db.WithContext(ctx)
	provider := request.Provider
	name := request.Name
	path := request.Path
	config := request.Config

	if util.IsAllDigits(name) || name == "" {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "storage name %s cannot be all digits or empty", name)
	}

	backend, ok := storagesystem.BackendMap[storageType]
	if !ok {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "storage type %s is not supported", storageType)
	}

	if config == nil {
		config = make(map[string]string)
	}

	if provider != "" {
		config["provider"] = provider
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

	for k, v := range config {
		rcloneConfig[k] = v
	}

	storage := model.Storage{
		Name:   name,
		Type:   storageType,
		Path:   path,
		Config: rcloneConfig,
	}
	rclone, err := storagesystem.NewRCloneHandler(ctx, storage)
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrap(err, "creating rclone handler failed"))
	}

	_, err = rclone.List(ctx, "")
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrap(err, "listing the storage failed"))
	}

	err = database.DoRetry(ctx, func() error {
		return db.Create(&storage).Error
	})

	if util.IsDuplicateKeyError(err) {
		return nil, errors.Wrapf(handlererror.ErrDuplicateRecord, "storage with name %s already exists", name)
	}

	return &storage, err
}
