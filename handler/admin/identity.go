package admin

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SetIdentityRequest struct {
	Identity string `json:"identity"`
}

func (DefaultHandler) SetIdentityHandler(ctx context.Context, db *gorm.DB, request SetIdentityRequest) error {
	db = db.WithContext(ctx)
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(&model.Global{
		Key:   "identity",
		Value: request.Identity,
	}).Error

	return errors.WithStack(err)
}

// @ID SetIdentity
// @Summary Set the user identity for tracking purpose
// @Tags Admin
// @Accept json
// @Produce json
// @Param request body SetIdentityRequest true "Create Request"
// @Success 204
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /identity [post]
func _() {}
