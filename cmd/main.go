package main

import (
	"log"
	"os"

	"github.com/NikiTesla/geant4help"
	"github.com/NikiTesla/geant4help/pkg/environment"
	"github.com/NikiTesla/geant4help/pkg/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("can't load env variables, err: %s", err.Error())
	}

	configFile := os.Getenv("CONFIGFILE")
	cfg, err := environment.NewConfig(configFile)
	if err != nil {
		log.Fatalf("can't load config, err: %s", err.Error())
	}

	rtr := routes.InitRouter()
	server := geant4help.Server{}

	if err := server.Run(cfg.Port, rtr); err != nil {
		log.Fatalf("error occured during serving: %s", err.Error())
	}

}
