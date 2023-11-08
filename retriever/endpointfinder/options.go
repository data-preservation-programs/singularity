package endpointfinder

import "time"

const (
	defaultLruSize         = 128
	defaultErrorLruSize    = 128
	defaultErrorLruTimeout = 5 * time.Minute
)

type config struct {
	LruSize         int
	ErrorLruSize    int
	ErrorLruTimeout time.Duration
}

func applyOptions(opts ...Option) *config {
	cfg := &config{
		LruSize:         defaultLruSize,
		ErrorLruSize:    defaultErrorLruSize,
		ErrorLruTimeout: defaultErrorLruTimeout,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

type Option func(*config)

func WithLruSize(size int) Option {
	return func(cfg *config) {
		cfg.LruSize = size
	}
}

func WithErrorLruSize(size int) Option {
	return func(cfg *config) {
		cfg.ErrorLruSize = size
	}
}

func WithErrorLruTimeout(timeout time.Duration) Option {
	return func(cfg *config) {
		cfg.ErrorLruTimeout = timeout
	}
}
