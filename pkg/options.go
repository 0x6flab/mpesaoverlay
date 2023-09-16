package pkg

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

func WithMaxIdleConns(conns int) Options {
	return func(cfg *Config) {
		cfg.MaxIdleConns = conns
	}
}
