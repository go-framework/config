package redis

import "flag"

// Register config flag.
func (c *Config) RegisterFlag() {
	flag.StringVar(&c.Network, "redis-network", "tcp", "redis network: tcp/unix")
	flag.StringVar(&c.Addr, "redis-addr", "localhost:6379", "redis address host:port")
	flag.StringVar(&c.Password, "redis-password", "", "redis password")
	flag.IntVar(&c.DB, "redis-db", 0, "redis connect db")
}
