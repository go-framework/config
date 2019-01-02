package redis

import "github.com/go-redis/redis"

// Default go-redis redis Options.
var DefaultGoRedisOptions *redis.Options = DefaultRedisConfig.GetGoRedisOptions()

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
