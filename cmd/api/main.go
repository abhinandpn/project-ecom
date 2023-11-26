package main

import (
	"log"

	config "github.com/abhinandpn/project-ecom/pkg/config"

	di "github.com/abhinandpn/project-ecom/pkg/di"
)

func main() {

	cfg, configErr := config.LoadConfig()

	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(cfg)

	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}

}

