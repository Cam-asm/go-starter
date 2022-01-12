package config

import (
	"encoding/json"
	"os"
)

// GetEnvJson gets the environment variable `key` containing JSON data & unmarshalls it into `obj`.
func GetEnvJson(key string, obj interface{}) error {
	data := os.Getenv(key)
	if len(data) == 0 {
		return nil
	}

	return json.Unmarshal([]byte(data), obj)
}

// MustGetEnvJson calls os.GetEnvJson & panics if an error is returned.
func MustGetEnvJson(key string, obj interface{}) {
	if err := GetEnvJson(key, obj); err != nil {
		panic(err)
	}
}

// LoadJsonFile expects parameter obj to be a pointer.
// E.g: err := config.LoadJsonFile("config.json", &config)
func LoadJsonFile(fileName string, obj interface{}) error {
	src, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(src, obj)
}

// MustLoadJsonFile expects parameter obj to be a pointer.
// E.g: config.MustLoadJsonFile("config.json", &config)
func MustLoadJsonFile(fileName string, obj interface{}) {
	if err := LoadJsonFile(fileName, obj); err != nil {
		panic(err)
	}
}
