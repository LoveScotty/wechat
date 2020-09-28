package mp

// Setter changes config.
type Setter func(config *Config) error

// WithAppId changes a config appId.
func WithAppId(appId string) Setter {
	return func(config *Config) error {
		config.AppId = appId
		return nil
	}
}

// WithAppSecret changes a config appSecret.
func WithAppSecret(appSecret string) Setter {
	return func(config *Config) error {
		config.AppSecret = appSecret
		return nil
	}
}

// WithToken changes a config token.
func WithToken(token string) Setter {
	return func(config *Config) error {
		config.Token = token
		return nil
	}
}

// WithEncodingAesKey changes a config encodingAesKey.
func WithEncodingAesKey(encodingAesKey string) Setter {
	return func(config *Config) error {
		config.EncodingAesKey = encodingAesKey
		return nil
	}
}
