package postgres

var (
	// Default postgres config.
	DefaultPostgresConfig *Config = GetDefaultPostgresConfig()
)

// Get default psostgres config.
func GetDefaultPostgresConfig() *Config {
	return &Config{
		URI: "postgresql://localhost:5432",
	}
}
