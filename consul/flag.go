package consul

import "flag"

// Register debugger flag.
func (c *Config) RegisterFlag() {
	flag.StringVar(&c.Addr, "consul-addr", "localhost:8500", "consul addr host:port address")
	flag.CommandLine.Var(&c.Encoder, "consul-encoder", `consul encoder yaml|json|xml|toml|hcl (default "yaml")`)
}
