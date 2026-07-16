package dealpusher

// Option customizes DealPusher initialization.
type Option func(*DealPusher)

func WithPDPProofSetManager(manager PDPProofSetManager) Option {
	return func(d *DealPusher) {
		d.pdpProofSetManager = manager
	}
}

func WithPDPSchedulingConfig(cfg PDPSchedulingConfig) Option {
	return func(d *DealPusher) {
		d.pdpSchedulingConfig = cfg
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
