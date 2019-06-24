package redis

import (
	"time"
)

// Redis config copy from go-redis redis Options.
type Config struct {
	// The network type, either tcp or unix.
	// Default is tcp.
	Network string `json:"network" yaml:"network"`
	// host:port address.
	Addr string `json:"addr" yaml:"addr"`

	// Optional password. Must match the password specified in the
	// requirepass server configuration option.
	Password string `json:"password" yaml:"password"`
	// Database to be selected after connecting to the server.
	DB int `json:"db" yaml:"db"`

	// Maximum number of retries before giving up.
	// Default is to not retry failed commands.
	MaxRetries int `json:"max_retries" yaml:"max_retries"`
	// Minimum backoff between each retry.
	// Default is 8 milliseconds; -1 disables backoff.
	MinRetryBackoff time.Duration `json:"min_retry_backoff" yaml:"min_retry_backoff"`
	// Maximum backoff between each retry.
	// Default is 512 milliseconds; -1 disables backoff.
	MaxRetryBackoff time.Duration `json:"max_retry_backoff" yaml:"max_retry_backoff"`

	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout time.Duration `json:"dial_timeout" yaml:"dial_timeout"`
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
	// Default is 3 seconds.
	ReadTimeout time.Duration `json:"read_timeout" yaml:"read_timeout"`
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is ReadTimeout.
	WriteTimeout time.Duration `json:"write_timeout" yaml:"write_timeout"`

	// Maximum number of socket connections.
	// Default is 10 connections per every CPU as reported by runtime.NumCPU.
	PoolSize int `json:"pool_size" yaml:"pool_size"`
	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	MinIdleConns int `json:"min_idle_conns" yaml:"min_idle_conns"`
	// Connection age at which client retires (closes) the connection.
	// Default is to not close aged connections.
	MaxConnAge time.Duration `json:"max_conn_age" yaml:"max_conn_age"`
	// Amount of time client waits for connection if all connections
	// are busy before returning an error.
	// Default is ReadTimeout + 1 second.
	PoolTimeout time.Duration `json:"pool_timeout" yaml:"pool_timeout"`
	// Amount of time after which client closes idle connections.
	// Should be less than server's timeout.
	// Default is 5 minutes. -1 disables idle timeout check.
	IdleTimeout time.Duration `json:"idle_timeout" yaml:"idle_timeout"`
	// Frequency of idle checks made by idle connections reaper.
	// Default is 1 minute. -1 disables idle connections reaper,
	// but idle connections are still discarded by the client
	// if IdleTimeout is set.
	IdleCheckFrequency time.Duration `json:"idle_check_frequency" yaml:"idle_check_frequency"`

	// register update func when self updated.
	updateFunc []func(*Config)
}
