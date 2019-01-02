package debugger

// Debugger config.
type Config struct {
	// debugger enabled.
	Enabled bool `json:"enabled" yaml:"enabled"`
	// debug items.
	Object map[string]bool `json:"object" yaml:"object"`
}
