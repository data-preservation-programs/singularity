package dataprep

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// ListHandler fetches and returns a list of all Preparation records from the database.
// It also preloads the associated source and output storages for each Preparation.
//
// Parameters:
// - ctx: The context for database transactions and other operations.
// - db: A pointer to the gorm.DB instance representing the database connection.
//
// Returns:
// - A slice containing all Preparation records from the database.
// - An error, if any occurred during the database query operation.
//
// Note:
// The function uses the Preload() method of gorm to automatically load the related source
// and output storage records for each returned Preparation, simplifying subsequent operations
// on these records.
func ListHandler(ctx context.Context, db *gorm.DB) ([]model.Preparation, error) {
	var preparations []model.Preparation
	err := db.WithContext(ctx).Preload("SourceStorages").Preload("OutputStorages").Find(&preparations).Error
	return preparations, errors.WithStack(err)
}
