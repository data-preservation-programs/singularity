package wallet

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func ListWalletHandler(
	db *gorm.DB,
	datasetName string,
) ([]model.Wallet, error) {
	return listWalletHandler(db, datasetName)
}

// @Summary List all wallets of a dataset.
// @Tags Wallet
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Wallet
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /dataset/{datasetName}/wallet [get]
func listWalletHandler(
	db *gorm.DB,
	datasetName string,
) ([]model.Wallet, error) {
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
