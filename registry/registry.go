package registry

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
	"sync"

	"github.com/json-iterator/go"
	"gopkg.in/yaml.v3"

	"github.com/go-framework/errors"
)

// Defined Registry interface.
type Registry struct {
	configs sync.Map
}

// New Registry.
func NewRegistry() *Registry {
	return &Registry{}
}

// Register Config interface.
// parse every register config base of file.
func (r *Registry) Register(key string, config interface{}) {
	r.configs.Store(key, config)
}

// Get config with key, when not exist return nil.
func (r Registry) Get(key string) interface{} {
	config, _ := r.configs.Load(key)

	return config
}

// Parse data.
// parse every register config base of file.
func (r *Registry) ParseData(data []byte, ext string) (errs error) {

	// range map
	r.configs.Range(func(key, value interface{}) bool {

		switch ext {
		case ".json":
			if err := jsoniter.Unmarshal(data, value); err != nil {
				errs = errors.Append(errs, err)
				return true
			}
		case ".yaml", ".yml":
			if err := yaml.Unmarshal(data, value); err != nil {
				errs = errors.Append(errs, err)
				return true
			}
		default:
			errs = errors.Append(errs, fmt.Errorf("unsupported config file extension: %s", ext))
			return true
		}

		if inter, ok := value.(Config); ok {
			if err := inter.Update(); err != nil {
				errs = errors.Append(errs, err)
				return true
			}
		}

		return true
	})

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

	// range map
	r.configs.Range(func(key, value interface{}) bool {
		if err := unmarshal(value); err != nil {
			errs = errors.Append(errs, err)
		}

		return true
	})

	return
}

// Implement JSON Unmarshaler interface.
func (r *Registry) UnmarshalJSON(data []byte) (errs error) {

	// range map
	r.configs.Range(func(key, value interface{}) bool {
		if err := jsoniter.Unmarshal(data, value); err != nil {
			errs = errors.Append(errs, err)
		}
		return true
	})

	return
}

var (
	// global Registry.
	defaultRegistry *Registry = NewRegistry()
)

// Register Config interface.
func Register(key string, config interface{}) {
	defaultRegistry.Register(key, config)
}

// Get config with key, when not exist return nil.
func Get(key string) interface{} {
	return defaultRegistry.Get(key)
}

// Parse data.
func ParseData(data []byte, ext string) error {
	return defaultRegistry.ParseData(data, ext)
}

// Parse files.
func ParseFiles(files ...string) error {
	return defaultRegistry.ParseFiles(files...)
}
