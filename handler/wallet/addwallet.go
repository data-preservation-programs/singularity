package wallet

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func AddWalletHandler(
	db *gorm.DB,
	datasetName string,
	wallet string,
) (*model.WalletAssignment, error) {
	return addWalletHandler(db, datasetName, wallet)
}

// @Summary Associate a new wallet with a dataset
// @Tags Wallet Association
// @Produce json
// @Accept json
// @Param datasetName path string true "Dataset name"
// @Param wallet path string true "Wallet Address"
// @Success 200 {object} model.WalletAssignment
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /dataset/{datasetName}/wallet/{wallet} [post]
func addWalletHandler(
	db *gorm.DB,
	datasetName string,
	wallet string,
) (*model.WalletAssignment, error) {
	if datasetName == "" {
		return nil, handler.NewInvalidParameterErr("dataset name is required")
	}

	if wallet == "" {
		return nil, handler.NewInvalidParameterErr("wallet address is required")
	}

	dataset, err := database.FindDatasetByName(db, datasetName)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("failed to find dataset: " + err.Error())
	}

	var w model.Wallet
	err = db.Where("address = ? OR id = ?", wallet, wallet).First(&w).Error
	if err != nil {
		return nil, handler.NewInvalidParameterErr("failed to find wallet: " + err.Error())
	}

	a := model.WalletAssignment{
		DatasetID: dataset.ID,
		WalletID:  w.ID,
	}

	err = database.DoRetry(func() error {
		return db.Create(&a).Error
	})

	if err != nil {
		return nil, handler.NewInvalidParameterErr("failed to create wallet assignment: " + err.Error())
	}

	return &a, nil
}
