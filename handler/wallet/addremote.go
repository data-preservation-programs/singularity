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
	LotusAPI   string `json:"lotusApi"   swaggerignore:"true"`
	LotusToken string `json:"lotusToken" swaggerignore:"true"`
}

// AddRemoteHandler godoc
// @Summary Add a remote wallet
// @Tags Wallet
// @Accept json
// @Param request body AddRemoteRequest true "Request body"
// @Success 204
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /wallet/remote [post]
func AddRemoteHandler(
	db *gorm.DB,
	request AddRemoteRequest,
) *handler.Error {
	address.CurrentNetwork = address.Mainnet
	addr, err := address.NewFromString(request.Address)
	if err != nil {
		return handler.NewBadRequestString("invalid address")
	}

	_, err = peer.Decode(request.RemotePeer)
	if err != nil {
		return handler.NewBadRequestString("invalid peer")
	}

	if addr.Protocol() == address.ID {
		return handler.NewBadRequestString("invalid address")
	}

	var lotusClient jsonrpc.RPCClient
	if request.LotusToken == "" {
		lotusClient = jsonrpc.NewClient(request.LotusAPI)
	} else {
		lotusClient = jsonrpc.NewClientWithOpts(request.LotusAPI, &jsonrpc.RPCClientOpts{
			CustomHeaders: map[string]string{
				"Authorization": "Bearer " + request.LotusToken,
			},
		})
	}

	var result string
	err = lotusClient.CallFor(context.Background(), &result, "Filecoin.StateLookupID", addr.String(), nil)
	if err != nil {
		return handler.NewBadRequestString("invalid wallet")
	}

	wallet := model.Wallet{
		ID:         result,
		Address:    addr.String(),
		RemotePeer: request.RemotePeer,
	}

	err = database.DoRetry(func() error { return db.Create(&wallet).Error })
	if err != nil {
		return handler.NewHandlerError(err)
	}

	return nil
}
