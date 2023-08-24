package wallet

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func ListAttachedHandler(
	ctx context.Context,
	db *gorm.DB,
	preparationID uint32,
) (*model.Preparation, error) {
	var preparation model.Preparation
	err := db.Preload("SourceStorages").Preload("OutputStorages").Preload("Wallets").Where("id = ?", preparationID).First(&preparation).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %d not found", preparationID)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &preparation, nil
}

// @Summary List all wallets of a dataset.
// @Tags Wallet
// @Produce json
// @Param datasetName path string true "Preparation name"
// @Success 200 {array} model.Wallet
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /dataset/{datasetName}/wallet [get]
func _() {}
