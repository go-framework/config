package registry

// Config interface.
type Config interface {
	Update() error
}

// Callback interface.
type Callback interface {
	Callback() error
}
