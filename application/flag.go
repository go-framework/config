package application

import (
	"flag"
	"os"
	"path/filepath"
)

// Register application flag.
func (c *Config) RegisterFlag() {
	flag.StringVar(&c.Name, "name", filepath.Base(os.Args[0]), "application name")
}
