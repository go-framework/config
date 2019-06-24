package consul

// Consul config.
type Config struct {
	// Consul host:port address.
	Addr string `json:"addr" yaml:"addr"`
	// Consul encoder, type: yaml|json|xml|toml|hcl.
	Encoder EncoderType `json:"encoder" yaml:"encoder"`
}
