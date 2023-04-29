package main

import (
	"log"

	config "github.com/abhinandpn/project-ecom/pkg/config"
	di "github.com/abhinandpn/project-ecom/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()go mod vendor

	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
