package wallet

import (
	"context"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/filecoin-project/go-address"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type AddRemoteRequest struct {
	Address    string `json:"address"`    // Address is the Filecoin full address of the wallet
	RemotePeer string `json:"remotePeer"` // RemotePeer is the remote peer ID of the wallet, for remote signing purpose
}

func AddRemoteHandler(
	db *gorm.DB,
	ctx context.Context,
	lotusClient jsonrpc.RPCClient,
	request AddRemoteRequest,
) (*model.Wallet, error) {
	return addRemoteHandler(db, ctx, lotusClient, request)
}

// @Summary Add a remote wallet
// @Tags Wallet
// @Accept json
// @Produce json
// @Param request body AddRemoteRequest true "Request body"
// @Success 200 {object} model.Wallet
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /wallet/remote [post]
func addRemoteHandler(
	db *gorm.DB,
	ctx context.Context,
	lotusClient jsonrpc.RPCClient,
	request AddRemoteRequest,
) (*model.Wallet, error) {
	addr, err := address.NewFromString(request.Address)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid address")
	}

	_, err = peer.Decode(request.RemotePeer)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid peer")
	}

	if addr.Protocol() == address.ID {
		return nil, handler.NewBadRequestString("invalid address")
	}

	var result string
	err = lotusClient.CallFor(ctx, &result, "Filecoin.StateLookupID", addr.String(), nil)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid wallet")
	}

	wallet := model.Wallet{
		ID:         result,
		Address:    request.Address,
		RemotePeer: request.RemotePeer,
	}

	err = database.DoRetry(func() error { return db.Create(&wallet).Error })
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return &wallet, nil
}
