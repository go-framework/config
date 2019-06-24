package application

// Application config.
type Config struct {
	// Application name.
	Name string `json:"name" yaml:"name"`
	// Application label.
	Label []string `json:"label" yaml:"label"`
}
