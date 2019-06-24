package influx

import (
	"time"
)

// Config option interface.
type Option interface {
	Apply(c *Config)
}

type optionFunc func(c *Config)

func (f optionFunc) Apply(c *Config) {
	f(c)
}

// With timeout option.
func WithTimeout(duration time.Duration) Option {
	return optionFunc(func(c *Config) {
		c.Timeout = duration
	})
}
