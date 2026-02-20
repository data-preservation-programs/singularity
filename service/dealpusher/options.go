package dealpusher

import "github.com/data-preservation-programs/singularity/model"

// Option customizes DealPusher initialization.
type Option func(*DealPusher)

// WithPDPProofSetManager sets the PDP proof set manager dependency.
func WithPDPProofSetManager(manager PDPProofSetManager) Option {
	return func(d *DealPusher) {
		d.pdpProofSetManager = manager
	}
}

// WithPDPTransactionConfirmer sets the PDP transaction confirmer dependency.
func WithPDPTransactionConfirmer(confirmer PDPTransactionConfirmer) Option {
	return func(d *DealPusher) {
		d.pdpTxConfirmer = confirmer
	}
}

// WithPDPSchedulingConfig overrides PDP scheduling configuration.
func WithPDPSchedulingConfig(cfg PDPSchedulingConfig) Option {
	return func(d *DealPusher) {
		d.pdpSchedulingConfig = cfg
	}
}

// WithScheduleDealTypeResolver overrides schedule deal type resolution logic.
func WithScheduleDealTypeResolver(resolver func(schedule *model.Schedule) model.DealType) Option {
	return func(d *DealPusher) {
		d.scheduleDealTypeResolver = resolver
	}
}
