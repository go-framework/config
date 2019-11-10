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

const config = `
int: 0
string: c

c0:
  int: 10
  string: c0
`

func TestRegistry(t *testing.T) {
	c0 := new(Config0)
	c1 := new(Config1)
	c2 := new(Config1)

	registry.RegisterAfter("config2", c2, "config0")
	registry.Register("config0", c0)
	registry.Register("config1", c1)

	err := registry.ParseData([]byte(config), ".yaml")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("config0", c0)
	t.Log("config1", c1)
	t.Log("config2", c2)

}
