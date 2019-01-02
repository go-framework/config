package redis

import "github.com/gomodule/redigo/redis"

// Default redisgo redis Pool.
var DefaultRedisGoPool *redis.Pool = DefaultRedisConfig.GetRedisGoPool()

// Convert to redisgo redis Pool.
func (c Config) GetRedisGoPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			// redis dial option.
			opts := []redis.DialOption{
				redis.DialDatabase(c.DB),
			}

			if len(c.Password) > 0 {
				opts = append(opts, redis.DialPassword(c.Password))
			}

			conn, err := redis.Dial(c.Network, c.Addr, opts...)
			if err != nil {
				return nil, err
			}

			return conn, nil
		},
		MaxIdle:         c.MinIdleConns,
		IdleTimeout:     c.IdleTimeout,
		MaxConnLifetime: c.MaxConnAge,
	}
}
