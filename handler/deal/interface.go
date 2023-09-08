package deal

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"gorm.io/gorm"
)

type Handler interface {
	ListHandler(ctx context.Context, db *gorm.DB, request ListDealRequest) ([]model.Deal, error)
	SendManualHandler(
		ctx context.Context,
		db *gorm.DB,
		dealMaker replication.DealMaker,
		request Proposal,
	) (*model.Deal, error)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}
