package endpointfinder

import "time"

const (
	defaultLruSize         = 128
	defaultLruTimeout      = 2 * time.Hour
	defaultErrorLruSize    = 128
	defaultErrorLruTimeout = 5 * time.Minute
)

type config struct {
	LruSize         int
	LruTimeout      time.Duration
	ErrorLruSize    int
	ErrorLruTimeout time.Duration
}

type Option func(*config)

func WithLruSize(size int) Option {
	return func(cfg *config) {
		cfg.LruSize = size
	}
}

func WithLruTimeout(timeout time.Duration) Option {
	return func(cfg *config) {
		cfg.LruTimeout = timeout
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
