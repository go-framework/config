package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// New mongo client.
func (c Config) NewClient() (*mongo.Client, error) {
	return mongo.NewClient(options.Client().ApplyURI(c.URI))
}

// New mongo client.
func (c Config) Connect(ctx context.Context) (*mongo.Client, error) {
	return mongo.Connect(ctx, options.Client().ApplyURI(c.URI))
}
