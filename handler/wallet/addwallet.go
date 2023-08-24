package wallet

import (
	"context"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func AddWalletHandler(
	ctx context.Context,
	db *gorm.DB,
	datasetName string,
	wallet string,
) (*model.WalletAssignment, error) {
	return addWalletHandler(ctx, db.WithContext(ctx), datasetName, wallet)
}

// @Summary Associate a new wallet with a dataset
// @Tags Wallet Association
// @Produce json
// @Accept json
// @Param datasetName path string true "Preparation name"
// @Param wallet path string true "Wallet Address"
// @Success 200 {object} model.WalletAssignment
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /dataset/{datasetName}/wallet/{wallet} [post]
func addWalletHandler(
	ctx context.Context,
	db *gorm.DB,
	datasetName string,
	wallet string,
) (*model.WalletAssignment, error) {
	if datasetName == "" {
		return nil, handlererror.NewInvalidParameterErr("dataset name is required")
	}

	if wallet == "" {
		return nil, handlererror.NewInvalidParameterErr("wallet address is required")
	}

	dataset, err := database.FindDatasetByName(db, datasetName)
	if err != nil {
		return nil, handlererror.NewInvalidParameterErr("failed to find dataset: " + err.Error())
	}

	var w model.Wallet
	err = db.Where("address = ? OR id = ?", wallet, wallet).First(&w).Error
	if err != nil {
		return nil, handlererror.NewInvalidParameterErr("failed to find wallet: " + err.Error())
	}

	a := model.WalletAssignment{
		DatasetID: dataset.ID,
		WalletID:  w.ID,
	}

	err = database.DoRetry(ctx, func() error {
		return db.Create(&a).Error
	})

	if err != nil {
		return nil, handlererror.NewInvalidParameterErr("failed to create wallet assignment: " + err.Error())
	}

	return &a, nil
}
