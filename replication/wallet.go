package replication

import (
	"context"
	"github.com/data-preservation-programs/singularity/model"
)

type WalletChooser struct {
}

func (w WalletChooser) Choose(ctx context.Context, wallets []model.Wallet) model.Wallet {
	return wallets[0]
}
