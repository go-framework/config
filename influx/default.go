package influx

var (
	// Default influx config.
	DefaultInfluxConfig *Config = GetDefaultInfluxConfig()
)

// Get default influx config.
func GetDefaultInfluxConfig() *Config {
	return &Config{
		Addr: "http://localhost:8086",
	}
}
