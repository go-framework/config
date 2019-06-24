package influx

import (
	"time"
)

// Influx config.
type Config struct {
	// Addr should be of the form "http://host:port"
	// or "http://[ipv6-host%zone]:port".
	Addr string `json:"addr" yaml:"addr"`

	// Username is the influxdb username, optional.
	Username string `json:"username" yaml:"username"`

	// Password is the influxdb password, optional.
	Password string `json:"password" yaml:"password"`

	// UserAgent is the http User Agent, defaults to "InfluxDBClient".
	UserAgent string `json:"user_agent" yaml:"user_agent"`

	// Timeout for influxdb writes, defaults to no timeout.
	Timeout time.Duration `json:"timeout" yaml:"timeout"`
}


