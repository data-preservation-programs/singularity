package dataset

import (
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)


// ListWalletHandler godoc
// @Summary List all wallets of a dataset.
// @Tags Dataset
// @Produce json
// @Param name path string true "Dataset name"
// @Success 200 {array} model.Wallet
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /dataset/{name}/wallets [get]
func ListWalletHandler(
	db *gorm.DB,
	name string,
) ([]model.Wallet, *handler.Error) {
	log.SetAllLoggers(log.LevelInfo)
	if name == "" {
		return nil, handler.NewBadRequestString("dataset name is required")
	}

	var dataset model.Dataset
	err := db.Preload("Wallets").Where("name = ?", name).First(&dataset).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewBadRequestString("dataset not found")
	}
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return dataset.Wallets, nil
}
