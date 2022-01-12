package config

import "github.com/kelseyhightower/envconfig"

// MustLoadYamlFileEnv attempts to load YAML configuration file in yamlFilePath and then environment variables into obj.
// If yamlFilePath is empty, then no YAML file is parsed.
// obj is expected to be a pointer.
func MustLoadYamlFileEnv(yamlFilePath, envPrefix string, obj interface{}) {
	if yamlFilePath != "" {
		// Load defined fields from a YAML config file.
		MustLoadYamlFile(yamlFilePath, obj)
	}

	// Any environment variables loaded are over-riding any values defined from the YAML config file.
	envconfig.MustProcess(envPrefix, obj)
}

// MustLoadEnvYamlFile attempts to load environment variables and then YAML configuration file in yamlFilePath into obj. MustLoadEnvYamlFile is the same as MustLoadYamlFileEnv except it is executed in reverse order.
// If yamlFilePath is empty, then no YAML file is parsed.
// obj is expected to be a pointer.
func MustLoadEnvYamlFile(envPrefix, yamlFilePath string, obj interface{}) {
	// Load available environment variables into obj.
	envconfig.MustProcess(envPrefix, obj)

	// Any fields defined in the YAML config file override environment variables.
	if yamlFilePath != "" {
		// Load defined fields from a YAML config file.
		MustLoadYamlFile(yamlFilePath, obj)
	}
}

// MustLoadJsonFileEnv attempts to load JSON configuration file in jsonFilePath and then environment variables into obj.
// If jsonFilePath is empty, then no JSON file is parsed.
// obj is expected to be a pointer.
func MustLoadJsonFileEnv(jsonFilePath, envPrefix string, obj interface{}) {
	if jsonFilePath != "" {
		// Load defined fields from a JSON config file.
		MustLoadJsonFile(jsonFilePath, obj)
	}

	// Any environment variables loaded are over-riding any values defined from the JSON config file.
	envconfig.MustProcess(envPrefix, obj)
}

// MustLoadEnvJsonFile attempts to load environment variables and then JSON configuration file in jsonFilePath into obj.
// MustLoadEnvJsonFile is the same as MustLoadJsonFileEnv except it is executed in reverse order.
// If jsonFilePath is empty, then no JSON file is parsed.
// obj is expected to be a pointer.
func MustLoadEnvJsonFile(envPrefix, jsonFilePath string, obj interface{}) {
	// Load available environment variables into obj.
	envconfig.MustProcess(envPrefix, obj)

	// Any fields defined in the JSON config file override environment variables.
	if jsonFilePath != "" {
		// Load defined fields from a JSON config file.
		MustLoadJsonFile(jsonFilePath, obj)
	}
}
