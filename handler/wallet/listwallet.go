package wallet

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)


// ListWalletHandler godoc
// @Summary List all wallets of a dataset.
// @Tags Wallet
// @Produce json
// @Param name path string true "Dataset name"
// @Success 200 {array} model.Wallet
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /dataset/{datasetName}/wallets [get]
func ListWalletHandler(
	db *gorm.DB,
	datasetName string,
) ([]model.Wallet, *handler.Error) {
	if datasetName == "" {
		return nil, handler.NewBadRequestString("dataset name is required")
	}

	var dataset model.Dataset
	err := db.Preload("Wallets").Where("name = ?", datasetName).First(&dataset).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewBadRequestString("dataset not found")
	}
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return dataset.Wallets, nil
}
