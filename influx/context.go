package influx

import (
	"context"
)

// config context key.
type configContextKey struct{}

// New config context.
func NewContext(ctx context.Context, config *Config) context.Context {
	return context.WithValue(ctx, configContextKey{}, config)
}

// Get config from context.
func FromContext(ctx context.Context) (*Config, bool) {
	config, ok := ctx.Value(configContextKey{}).(*Config)
	return config, ok
}
