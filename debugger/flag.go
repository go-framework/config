package debugger

import "flag"

// Register debugger flag.
func (c *Config) RegisterFlag() {
	flag.BoolVar(&c.Enabled, "debugger-enable", false, "debugger enable")
}
