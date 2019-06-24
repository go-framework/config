package influx

import (
	client "github.com/influxdata/influxdb1-client/v2"
)

// Convert influx client v2 config.
func (c Config) GetClientV2Config() *client.HTTPConfig {
	return &client.HTTPConfig{
		Addr:      c.Addr,
		Username:  c.Username,
		Password:  c.Password,
		UserAgent: c.UserAgent,
		Timeout:   c.Timeout,
	}
}

// New influx client v2.
func (c Config) NewClientV2() (client.Client, error) {
	return client.NewHTTPClient(*c.GetClientV2Config())
}
