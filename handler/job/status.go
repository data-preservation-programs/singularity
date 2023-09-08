package job

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
// - id: The ID or name for the desired Preparation record.
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
func (DefaultHandler) GetStatusHandler(ctx context.Context, db *gorm.DB, id string) ([]SourceStatus, error) {
	db = db.WithContext(ctx)
	var preparation model.Preparation
	err := preparation.FindByIDOrName(db, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %d cannot be found", id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var sourceAttachments []model.SourceAttachment
	err = db.Preload("Storage").Where("preparation_id = ?", preparation.ID).Find(&sourceAttachments).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var allStatuses []SourceStatus
	for _, sourceAttachment := range sourceAttachments {
		var jobs []model.Job
		err = db.Where("attachment_id = ?", sourceAttachment.ID).Find(&jobs).Error
		if err != nil {
			return nil, errors.WithStack(err)
		}
		status := SourceStatus{
			AttachmentID:    ptr.Of(sourceAttachment.ID),
			SourceStorageID: ptr.Of(sourceAttachment.StorageID),
			SourceStorage:   sourceAttachment.Storage,
			Jobs:            jobs,
		}
		allStatuses = append(allStatuses, status)
	}

	return allStatuses, nil
}

// @Summary Get the status of a preparation
// @Tags Preparation
// @Param id path string true "Preparation ID or name"
// @Produce json
// @Success 200 {array} SourceStatus
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id} [get]
func _() {}
