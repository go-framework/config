package registry

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/json-iterator/go"
	"gopkg.in/yaml.v3"

	"github.com/go-framework/errors"
)

// Callback interface.
type ICallback interface {
	Callback(config interface{}) error
}

// list node.
type node struct {
	name      string
	config    interface{}
	callbacks []ICallback
}

// Defined Registry interface.
type Registry struct {
	l *list.List
}

// New Registry.
func NewRegistry() *Registry {
	return &Registry{
		l: list.New(),
	}
}

// Register Config interface.
// parse every register config base of file.
func (r *Registry) Register(name string, config interface{}, callbacks ...ICallback) {
	// replace one
	for e := r.l.Front(); e != nil; e = e.Next() {
		if n, ok := e.Value.(*node); ok {
			if n.name == name {
				n.config = config
				n.callbacks = append(n.callbacks, callbacks...)
				return
			}
		}
	}

	// push new
	r.l.PushBack(&node{
		name:      name,
		config:    config,
		callbacks: callbacks,
	})
}

// Register Config callback.
func (r *Registry) RegisterCallback(name string, callbacks ...ICallback) {
	// replace one
	for e := r.l.Front(); e != nil; e = e.Next() {
		if n, ok := e.Value.(*node); ok {
			if n.name == name {
				n.callbacks = append(n.callbacks, callbacks...)
				return
			}
		}
	}

	// push new
	r.l.PushBack(&node{
		name:      name,
		callbacks: callbacks,
	})
}

// Register Config interface, parsed name's config after after's config.
// parse every register config base of file.
func (r *Registry) RegisterAfter(name string, config interface{}, after string) {
	var selected *list.Element

	// find after
	for e := r.l.Front(); e != nil; e = e.Next() {
		if n, ok := e.Value.(*node); ok {
			if n.name == after {
				selected = e
				break
			}
		}
	}

	// not find than push new one
	if selected == nil {
		selected = r.l.PushBack(&node{
			name: after,
		})
	}

	// new one push after
	r.l.InsertAfter(&node{
		name:   name,
		config: config,
	}, selected)
}

// Get config with key, when not exist return nil.
func (r Registry) Get(name string) interface{} {
	// find after
	for e := r.l.Front(); e != nil; e = e.Next() {
		if n, ok := e.Value.(*node); ok {
			if n.name == name {
				return n.config
			}
		}
	}

	return nil
}

// Parse data.
// parse every register config base of file.
func (r *Registry) ParseData(data []byte, ext string) (errs error) {

	// range
	for e := r.l.Front(); e != nil; e = e.Next() {
		if n, ok := e.Value.(*node); ok {
			switch ext {
			case ".json":
				if err := jsoniter.Unmarshal(data, n.config); err != nil {
					errs = errors.Append(errs, err)
					continue
				}
			case ".yaml", ".yml":
				if err := yaml.Unmarshal(data, n.config); err != nil {
					errs = errors.Append(errs, err)
					continue
				}
			default:
				errs = errors.Append(errs, fmt.Errorf("unsupported config file extension: %s", ext))
				continue
			}

			// update
			if inter, ok := n.config.(Config); ok {
				if err := inter.Update(); err != nil {
					errs = errors.Append(errs, err)
				}
			}

			// exec callbacks
			for i := 0; i < len(n.callbacks); i++ {
				if err := n.callbacks[i].Callback(n); err != nil {
					errs = errors.Append(errs, err)
				}
			}

		}
	}
	return
}

// Parse files.
// parse every register config base of file.
func (r *Registry) ParseFiles(files ...string) (errs error) {

	for _, file := range files {

		// read file
		data, err := ioutil.ReadFile(file)
		if err != nil {
			errs = errors.Append(errs, err)
			continue
		}

		// get file name extension
		ext := strings.ToLower(path.Ext(file))

		if err := r.ParseData(data, ext); err != nil {
			errs = errors.Append(errs, err)
			continue
		}
	}

	return
}

// Implement YAML Unmarshaler interface.
func (r *Registry) UnmarshalYAML(unmarshal func(interface{}) error) (errs error) {

	// range
	for e := r.l.Front(); e != nil; e = e.Next() {
		if n, ok := e.Value.(*node); ok {
			if err := unmarshal(n.config); err != nil {
				errs = errors.Append(errs, err)
			}
		}
	}

	return
}

// Implement JSON Unmarshaler interface.
func (r *Registry) UnmarshalJSON(data []byte) (errs error) {
	// range
	for e := r.l.Front(); e != nil; e = e.Next() {
		if n, ok := e.Value.(*node); ok {
			if err := jsoniter.Unmarshal(data, n.config); err != nil {
				errs = errors.Append(errs, err)
			}
		}
	}

	return
}

//
// default
//
var (
	// global Registry.
	defaultRegistry *Registry = NewRegistry()
)

// Register Config interface.
func Register(name string, config interface{}) {
	defaultRegistry.Register(name, config)
}

// Register Callback in name.
func RegisterCallback(name string, callbacks ...ICallback) {
	defaultRegistry.RegisterCallback(name, callbacks...)
}

// Register Config interface.
func RegisterAfter(name string, config interface{}, after string) {
	defaultRegistry.RegisterAfter(name, config, after)
}

// Get config with key, when not exist return nil.
func Get(name string) interface{} {
	return defaultRegistry.Get(name)
}

// Parse data.
func ParseData(data []byte, ext string) error {
	return defaultRegistry.ParseData(data, ext)
}

// Parse files.
func ParseFiles(files ...string) error {
	return defaultRegistry.ParseFiles(files...)
}
