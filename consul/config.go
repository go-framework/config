package consul

// Consul config.
type Config struct {
	// host:port address.
	Addr string `json:"addr" yaml:"addr"`
	// consul encoder, type: yaml|json|xml|toml|hcl.
	Encoder EncoderType `json:"encoder" yaml:"encoder"`
}
