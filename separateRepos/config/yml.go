// Package config implements utility functions for loading configuration data from various sources.
package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// LoadYamlFile expects parameter obj to be a pointer.
// E.g: err := config.LoadYamlFile("config.yml", &env)
func LoadYamlFile(fileName string, obj interface{}) error {
	src, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(src, obj)
}

// MustLoadYamlFile expects parameter obj to be a pointer.
// E.g: config.MustLoadYamlFile("config.yml", &env)
func MustLoadYamlFile(fileName string, obj interface{}) {
	if err := LoadYamlFile(fileName, obj); err != nil {
		panic(err)
	}
}
