package environment

import (
	"bytes"
	"errors"
	"fmt"
)

type Environment int8

const (
	Development   Environment = iota // development environment.
	Testing                          // testing environment.
	PreProduction                    // pre production environment.
	Production                       // production environment.
)

// String returns a lower-case ASCII representation of the env.
func (env Environment) String() string {
	switch env {
	case Development:
		return "development"
	case Testing:
		return "testing"
	case PreProduction:
		return "pre-production"
	case Production:
		return "production"
	default:
		return fmt.Sprintf("Env(%d)", env)
	}
}

// MarshalText marshals the Environment to text. Note that the text representation
// drops the -Level suffix (see example).
func (env Environment) MarshalText() ([]byte, error) {
	return []byte(env.String()), nil
}

// UnmarshalText unmarshals text to a env. Like MarshalText, UnmarshalText
// expects the text representation of a Environment to drop the -Level suffix (see
// example).
//
// In particular, this makes it easy to configure logging levels using YAML,
// TOML, or JSON files.
func (env *Environment) UnmarshalText(text []byte) error {
	if env == nil {
		return errors.New("can't unmarshal a nil *Level")
	}
	if !env.unmarshalText(text) && !env.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized level: %q", text)
	}
	return nil
}

func (env *Environment) unmarshalText(text []byte) bool {
	switch string(text) {
	case "development", "": // make the zero value useful
		*env = Development
	case "testing":
		*env = Testing
	case "pre-production":
		*env = PreProduction
	case "production":
		*env = Production
	default:
		return false
	}
	return true
}

// Set sets the env for the flag.Value interface.
func (env *Environment) Set(s string) error {
	return env.UnmarshalText([]byte(s))
}

// Get gets the env for the flag.Getter interface.
func (env *Environment) Get() interface{} {
	return *env
}
