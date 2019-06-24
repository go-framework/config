package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// config context key.
type configContextKey struct{}

// New config context.
func NewConfigContext(ctx context.Context, config *Config) context.Context {
	return context.WithValue(ctx, configContextKey{}, config)
}

// Get config from context.
func ConfigFromContext(ctx context.Context) (*Config, bool) {
	config, ok := ctx.Value(configContextKey{}).(*Config)
	return config, ok
}

// client context key.
type clientContextKey struct{}

// New client context.
func NewClientContext(ctx context.Context, client *mongo.Client) context.Context {
	return context.WithValue(ctx, clientContextKey{}, client)
}

// Get client from context.
func ClientFromContext(ctx context.Context) (*mongo.Client, bool) {
	client, ok := ctx.Value(clientContextKey{}).(*mongo.Client)
	return client, ok
}
