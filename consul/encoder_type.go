package consul

import (
	"bytes"
	"errors"
	"fmt"
)

// Consul encoder type.
type EncoderType int8

const (
	YAML EncoderType = iota
	JSON
	XML
	TOML
	HCL
)

// String returns a lower-case ASCII representation of the enc.
func (enc EncoderType) String() string {
	switch enc {
	case YAML:
		return "yaml"
	case JSON:
		return "json"
	case XML:
		return "xml"
	case TOML:
		return "toml"
	case HCL:
		return "hcl"
	default:
		return fmt.Sprintf("Encoder(%d)", enc)
	}
}

// MarshalText marshals the EncoderType to text. Note that the text representation
// drops the -Level suffix (see example).
func (enc EncoderType) MarshalText() ([]byte, error) {
	return []byte(enc.String()), nil
}

// UnmarshalText unmarshals text to a enc. Like MarshalText, UnmarshalText
// expects the text representation of a EncoderType to drop the -Level suffix (see
// example).
//
// In particular, this makes it easy to configure logging levels using YAML,
// TOML, or JSON files.
func (enc *EncoderType) UnmarshalText(text []byte) error {
	if enc == nil {
		return errors.New("can't unmarshal a nil *Level")
	}
	if !enc.unmarshalText(text) && !enc.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized level: %q", text)
	}
	return nil
}

func (enc *EncoderType) unmarshalText(text []byte) bool {
	switch string(text) {
	case "yaml", "": // make the zero value useful
		*enc = YAML
	case "json":
		*enc = JSON
	case "xml":
		*enc = XML
	case "toml":
		*enc = TOML
	case "hcl":
		*enc = HCL
	default:
		return false
	}
	return true
}

// Set sets the enc for the flag.Value interface.
func (enc *EncoderType) Set(s string) error {
	return enc.UnmarshalText([]byte(s))
}

// Get gets the enc for the flag.Getter interface.
func (enc *EncoderType) Get() interface{} {
	return *enc
}
