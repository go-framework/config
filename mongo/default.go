package mongo

var (
	// Default mongo config.
	DefaultMongoConfig *Config = GetDefaultMongoConfig()
)

// Get default mongo config.
func GetDefaultMongoConfig() *Config {
	return &Config{
		URI: "mongodb://localhost:27017",
	}
}
