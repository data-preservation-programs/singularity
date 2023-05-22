package wallet

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// RemoveHandler godoc
// @Summary Remove a wallet
// @Tags Wallet
// @Param address path string true "Address"
// @Success 204
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /wallet/:address [delete]
func RemoveHandler(
	db *gorm.DB,
	address string,
) *handler.Error {
	err := db.Where("address = ? OR id = ?", address, address).Delete(&model.Wallet{}).Error
	if err != nil {
		return handler.NewHandlerError(err)
	}
	return nil
}
