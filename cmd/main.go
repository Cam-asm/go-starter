package main

import (
	"net/http"

	"go-starter/api"
	"go-starter/separateRepos/config"
	"go-starter/separateRepos/graceful"
	"go-starter/separateRepos/logger"
	"go-starter/separateRepos/utl"
)

func main() {
	// Initialise logging.
	logger.Init(api.Name)

	// Load environment config from config.yml and environment variables.
	var env api.Config
	config.MustLoadYamlFileEnv("config.yml", api.EnvPrefix, &env)
	env.Print()

	graceful.Server(&http.Server{
		Addr:              utl.HostPort("", env.Port),
		Handler:           api.Router(),
		ReadHeaderTimeout: env.HeaderTimeout,
	})
}
