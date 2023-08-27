package pkg

import "context"

type SDKOption func(*Config)

func WithContext(ctx context.Context) SDKOption {
	return func(cfg *Config) {
		cfg.CTX = ctx
	}
}

func WithBaseURL(url string) SDKOption {
	return func(cfg *Config) {
		cfg.BaseURL = url
	}
}

func WithAppKey(key string) SDKOption {
	return func(cfg *Config) {
		cfg.AppKey = key
	}
}

func WithAppSecret(secret string) SDKOption {
	return func(cfg *Config) {
		cfg.AppSecret = secret
	}
}

func WithMaxIdleConns(conns int) SDKOption {
	return func(cfg *Config) {
		cfg.MaxIdleConns = conns
	}
}
