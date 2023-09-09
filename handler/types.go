package handler

import (
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type NoContent struct {
}

type Request[T any] struct {
	Params  []string
	Payload T
}

type Dependency struct {
	DB          *gorm.DB
	LotusClient jsonrpc.RPCClient
	DealMaker   replication.DealMaker
}
