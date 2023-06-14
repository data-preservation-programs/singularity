package wallet

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// RemoveWalletHandler godoc
// @Summary Remove an associated wallet from a dataset
// @Tags Wallet
// @Param name path string true "Dataset name"
// @Param wallet path string true "Wallet Address"
// @Success 204
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /dataset/{datasetName}/wallet/{wallet} [delete]
func RemoveWalletHandler(
	db *gorm.DB,
	datasetName string,
	wallet string,
) *handler.Error {
	if datasetName == "" {
		return handler.NewBadRequestString("dataset name is required")
	}

	if wallet == "" {
		return handler.NewBadRequestString("wallet address is required")
	}

	dataset, err := database.FindDatasetByName(db, datasetName)
	if err != nil {
		return handler.NewBadRequestString("failed to find dataset: " + err.Error())
	}

	var w model.Wallet
	err = db.Where("address = ?", wallet).First(&w).Error
	if err != nil {
		return handler.NewBadRequestString("failed to find wallet: " + err.Error())
	}

	err = database.DoRetry(func() error {
		return db.Where("dataset_id = ? AND wallet_id = ?", dataset.ID, w.ID).Delete(&model.WalletAssignment{}).Error
	})

	if err != nil {
		return handler.NewHandlerError(err)
	}

	return nil
}
