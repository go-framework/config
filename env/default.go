package env

var (
	// Default environment, value is Development.
	DefaultEnvironment Environment = Development
)

// Get Development environment
func GetDevelopmentEnvironment() Environment {
	return Development
}

// Get Production environment
func GetProductionEnvironment() Environment {
	return Production
}
