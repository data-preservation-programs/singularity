package dataprep

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/gotidy/ptr"
	"gorm.io/gorm"
)

type SourceStatus struct {
	AttachmentID    *uint32        `json:"attachmentId"`
	SourceStorageID *uint32        `json:"storageId"`
	SourceStorage   *model.Storage `json:"source"       table:"expand"`
	Jobs            []model.Job    `json:"jobs"         table:"expand"`
}

// GetStatusHandler fetches and returns the current status of a specific Preparation.
// The status includes the Preparation record and associated jobs for each source attachment.
//
// Parameters:
// - ctx: The context for database transactions and other operations.
// - db: A pointer to the gorm.DB instance representing the database connection.
// - id: The unique identifier for the desired Preparation record.
//
// Returns:
//   - A pointer to a Status structure that encapsulates the Preparation record and
//     the associated jobs for each source attachment.
//   - An error, if any occurred during the database query operation or if the Preparation record
//     with the specified ID does not exist.
//
// Note:
// The function fetches not only the Preparation record but also all associated SourceAttachment
// records with their associated Job records, providing a comprehensive status of a specific preparation.
func (DefaultHandler) GetStatusHandler(ctx context.Context, db *gorm.DB, id uint32) ([]SourceStatus, error) {
	db = db.WithContext(ctx)
	var preparation model.Preparation
	err := db.First(&preparation, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %d cannot be found", id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var sources []model.SourceAttachment
	err = db.Preload("Storage").Where("preparation_id = ?", id).Find(&sources).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var allStatuses []SourceStatus
	for _, source := range sources {
		var jobs []model.Job
		err = db.Where("attachment_id = ?", source.ID).Find(&jobs).Error
		if err != nil {
			return nil, errors.WithStack(err)
		}
		status := SourceStatus{
			AttachmentID:    ptr.Of(source.ID),
			SourceStorageID: ptr.Of(source.StorageID),
			SourceStorage:   source.Storage,
			Jobs:            jobs,
		}
		allStatuses = append(allStatuses, status)
	}

	return allStatuses, nil
}

// @Summary Get the status of a preparation
// @Tags Preparation
// @Param id path integer true "ID"
// @Produce json
// @Success 200 {object} Status
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /preparation/{id} [get]
func _() {}
