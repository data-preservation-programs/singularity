package wallet

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func RemoveHandler(
	db *gorm.DB,
	address string,
) error {
	return removeHandler(db, address)
}

// @Summary Remove a wallet
// @Tags Wallet
// @Param address path string true "Address"
// @Success 204
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /wallet/{address} [delete]
func removeHandler(
	db *gorm.DB,
	address string,
) error {
	err := database.DoRetry(func() error { return db.Where("address = ? OR id = ?", address, address).Delete(&model.Wallet{}).Error })
	if err != nil {
		return err
	}
	return nil
}
