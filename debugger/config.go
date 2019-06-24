package debugger

// Debugger config.
type Config struct {
	// Debugger enabled.
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Debug items.
	Object map[string]bool `json:"object" yaml:"object"`
}
