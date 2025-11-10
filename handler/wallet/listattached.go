package wallet

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// ListAttachedHandler fetches and returns a list of wallets associated with a given preparation,
// identified by either its ID or name.
//
// The function looks for the preparation with the specified ID or name in the database. If found,
// it retrieves all wallets associated with that preparation. If no such preparation is found,
// an error is returned.
//
// Parameters:
//   - ctx: The context in which the handler function is executed, used for controlling cancellation.
//   - db: A pointer to a gorm.DB object, which provides database access.
//   - preparationID: The ID or name of the preparation whose attached wallets need to be fetched.
//
// Returns:
//   - A slice of model.Wallet objects that are attached to the specified preparation.
//   - An error if any issues arise during the process or if the preparation is not found, otherwise nil.
func (DefaultHandler) ListAttachedHandler(
	ctx context.Context,
	db *gorm.DB,
	preparationID string,
) ([]model.Actor, error) {
	var preparation model.Preparation
	err := preparation.FindByIDOrName(db, preparationID, "Actors")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %d not found", preparationID)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return preparation.Actors, nil
}

// @ID ListAttachedWallets
// @Summary List all wallets of a preparation.
// @Tags Wallet Association
// @Produce json
// @Accept json
// @Param id path string true "Preparation ID or name"
// @Success 200 {array} model.Wallet
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/wallet [get]
func _() {}
