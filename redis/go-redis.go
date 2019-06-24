package redis

import "github.com/go-redis/redis"

// Type redis Client.
type Client = redis.Client

// Global variable.
var (
	// Defined go-redis nil.
	GoRedisNil = redis.Nil

	// Default go-redis redis Options.
	DefaultGoRedisOptions *redis.Options = DefaultRedisConfig.GetGoRedisOptions()
)

// Convert to go-redis redis Options.
func (c Config) GetGoRedisOptions() *redis.Options {
	return &redis.Options{
		Network:            c.Network,
		Addr:               c.Addr,
		Password:           c.Password,
		DB:                 c.DB,
		MaxRetries:         c.MaxRetries,
		MinRetryBackoff:    c.MinRetryBackoff,
		MaxRetryBackoff:    c.MaxRetryBackoff,
		DialTimeout:        c.DialTimeout,
		ReadTimeout:        c.ReadTimeout,
		WriteTimeout:       c.WriteTimeout,
		PoolSize:           c.PoolSize,
		MinIdleConns:       c.MinIdleConns,
		MaxConnAge:         c.MaxConnAge,
		PoolTimeout:        c.PoolTimeout,
		IdleTimeout:        c.IdleTimeout,
		IdleCheckFrequency: c.IdleCheckFrequency,
	}
}

// New go-redis client.
func (c Config) NewGoClient() *Client {
	return redis.NewClient(c.GetGoRedisOptions())
}
