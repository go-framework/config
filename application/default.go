package application

import (
	"os"
	"path/filepath"
)

var (
	// Default application config.
	DefaultConfig *Config = GetDefaultApplicationConfig()
)

// Get default application config.
func GetDefaultApplicationConfig() *Config {
	return &Config{
		Name: filepath.Base(os.Args[0]),
	}
}
