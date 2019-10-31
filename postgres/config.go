package postgres

// Postgres config.
type Config struct {
	// Connection URIs.
	// see: https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
	// format: postgresql://[user[:password]@][netloc][:port][,...][/dbname][?param1=value1&...]
	URI string `json:"uri" yaml:"uri"`
}


// New config
func NewConfig(uri string) *Config {
	return &Config{
		URI: uri,
	}
}