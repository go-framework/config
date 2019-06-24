package mongo

// Mongo config.
type Config struct {
	// Mongo URI.
	// see: https://docs.mongodb.com/manual/reference/connection-string/
	URI string `json:"uri" yaml:"uri"`
}

// New config
func NewConfig(uri string) *Config {
	return &Config{
		URI: uri,
	}
}
