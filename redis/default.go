package redis

var (
	// Default redis config.
	DefaultRedisConfig *Config = GetDefaultRedisConfig()
)

// Get default redis config.
func GetDefaultRedisConfig() *Config {
	return &Config{
		Network: "tcp",
		Addr:    "localhost:6379",
		DB:      0,
	}
}
