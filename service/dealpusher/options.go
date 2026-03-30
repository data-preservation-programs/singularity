package dealpusher

import "github.com/data-preservation-programs/singularity/model"

// Option customizes DealPusher initialization.
type Option func(*DealPusher)

func WithPDPProofSetManager(manager PDPProofSetManager) Option {
	return func(d *DealPusher) {
		d.pdpProofSetManager = manager
	}
}

func WithPDPTransactionConfirmer(confirmer PDPTransactionConfirmer) Option {
	return func(d *DealPusher) {
		d.pdpTxConfirmer = confirmer
	}
}

func WithPDPSchedulingConfig(cfg PDPSchedulingConfig) Option {
	return func(d *DealPusher) {
		d.pdpSchedulingConfig = cfg
	}
}

func WithScheduleDealTypeResolver(resolver func(schedule *model.Schedule) model.DealType) Option {
	return func(d *DealPusher) {
		d.scheduleDealTypeResolver = resolver
	}
}

func WithDDODealManager(manager DDODealManager) Option {
	return func(d *DealPusher) {
		d.ddoDealManager = manager
	}
}

func WithDDOSchedulingConfig(cfg DDOSchedulingConfig) Option {
	return func(d *DealPusher) {
		d.ddoSchedulingConfig = cfg
	}
}
