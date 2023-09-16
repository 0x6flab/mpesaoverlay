package pkg

import (
	"context"
	"net/http"
)

type Options func(*Config)

func WithBaseURL(url string) Options {
	return func(cfg *Config) {
		cfg.BaseURL = url
	}
}

func WithAppKey(key string) Options {
	return func(cfg *Config) {
		cfg.AppKey = key
	}
}

func WithAppSecret(secret string) Options {
	return func(cfg *Config) {
		cfg.AppSecret = secret
	}
}

func WithCertFile(certFile string) Options {
	return func(cfg *Config) {
		cfg.CertFile = certFile
	}
}

func WithContext(ctx context.Context) Options {
	return func(cfg *Config) {
		cfg.Context = ctx
	}
}

func WithHTTPClient(client *http.Client) Options {
	return func(cfg *Config) {
		cfg.HTTPClient = client
	}
}
