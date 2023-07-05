package wallet

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// AddWalletHandler godoc
// @Summary Associate a new wallet with a dataset
// @Tags Wallet Association
// @Produce json
// @Accept json
// @Param datasetName path string true "Dataset name"
// @Param wallet path string true "Wallet Address"
// @Success 200 {object} model.WalletAssignment
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /dataset/{datasetName}/wallet/{wallet} [post]
func AddWalletHandler(
	db *gorm.DB,
	datasetName string,
	wallet string,
) (*model.WalletAssignment, *handler.Error) {
	if datasetName == "" {
		return nil, handler.NewBadRequestString("dataset name is required")
	}

	if wallet == "" {
		return nil, handler.NewBadRequestString("wallet address is required")
	}

	dataset, err := database.FindDatasetByName(db, datasetName)
	if err != nil {
		return nil, handler.NewBadRequestString("failed to find dataset: " + err.Error())
	}

	var w model.Wallet
	err = db.Where("address = ? OR id = ?", wallet, wallet).First(&w).Error
	if err != nil {
		return nil, handler.NewBadRequestString("failed to find wallet: " + err.Error())
	}

	a := model.WalletAssignment{
		DatasetID: dataset.ID,
		WalletID:  w.ID,
	}

	err = database.DoRetry(func() error {
		return db.Create(&a).Error
	})

	if err != nil {
		return nil, handler.NewBadRequestString("failed to create wallet assignment: " + err.Error())
	}

	return &a, nil
}
