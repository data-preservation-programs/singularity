package storage

import (
	"context"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"gorm.io/gorm"
)

type RenameRequest struct {
	Name string `binding:"required" json:"name"`
}

func (DefaultHandler) RenameStorageHandler(
	ctx context.Context,
	db *gorm.DB,
	name string,
	request RenameRequest,
) (*model.Storage, error) {
	db = db.WithContext(ctx)

	// Validate input
	if util.IsAllDigits(request.Name) || request.Name == "" {
		return nil, errors.Wrapf(
			handlererror.ErrInvalidParameter,
			"storage name cannot be all digits or empty",
		)
	}

	// Find existing storage
	var storage model.Storage
	if err := storage.FindByIDOrName(db, name); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(
				handlererror.ErrNotFound,
				"storage %s does not exist",
				name,
			)
		}
		return nil, errors.WithStack(err)
	}

	// Skip if no name change
	if storage.Name == request.Name {
		return &storage, nil
	}

	// Check for existing name first
	var existing model.Storage
	if err := db.Where("name = ?", request.Name).First(&existing).Error; err == nil {
		return nil, errors.New("duplicated key not allowed")
	}

	// Perform the update
	if err := db.Model(&storage).Update("name", request.Name).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return nil, errors.New("duplicated key not allowed")
		}
		return nil, errors.WithStack(err)
	}

	storage.Name = request.Name
	return &storage, nil
}
