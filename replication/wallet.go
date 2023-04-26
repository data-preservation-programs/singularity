package replication

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/model"
)

type WalletChooser struct {
}

func (w WalletChooser) Choose(ctx context.Context, wallets []model.Wallet) model.Wallet {
	panic("implement me")
}
