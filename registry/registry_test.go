package registry_test

import (
	"testing"

	"github.com/go-framework/config/registry"
)

type Config0 struct {
	Int    int    `json:"int" yaml:"int"`
	String string `json:"string" yaml:"string"`
}

type Config1 struct {
	String string  `json:"string" yaml:"string"`
	C0     Config0 `json:"c0" yaml:"c0"`
}

func TestRegistry(t *testing.T) {
	c0 := new(Config0)

	c1 := new(Config1)

	registry.Register("config0", c0)
	registry.Register("config1", c1)

	filename := "config.yaml"

	err := registry.ParseFiles(filename)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("config0", c0)
	t.Log("config1", c1)
}
