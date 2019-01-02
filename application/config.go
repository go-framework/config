package application

// Application config.
type Config struct {
	// application name.
	Name string `json:"name" yaml:"name"`
	// application label.
	Label []string `json:"label" yaml:"label"`
}
