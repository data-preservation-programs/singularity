package wallet

import (
	"context"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func ListWalletHandler(
	ctx context.Context,
	db *gorm.DB,
	datasetName string,
) ([]model.Wallet, error) {
	return listWalletHandler(db.WithContext(ctx), datasetName)
}

// @Summary List all wallets of a dataset.
// @Tags Wallet
// @Produce json
// @Param datasetName path string true "Preparation name"
// @Success 200 {array} model.Wallet
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /dataset/{datasetName}/wallet [get]
func listWalletHandler(
	db *gorm.DB,
	datasetName string,
) ([]model.Wallet, error) {
	if datasetName == "" {
		return nil, handler.NewInvalidParameterErr("dataset name is required")
	}

	var dataset model.Preparation
	err := db.Preload("Wallets").Where("name = ?", datasetName).First(&dataset).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr("dataset not found")
	}
	if err != nil {
		return nil, err
	}

	return dataset.Wallets, nil
}
