package main

import (
	"log"

	"github.com/amalmadhu06/simple-api/pkg/config"
	"github.com/amalmadhu06/simple-api/pkg/di"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	server, err := di.InitializeAPI(cfg)
	if err != nil {
		log.Fatal("cannot initialize API: ", err)
	}
	server.Start()
}
