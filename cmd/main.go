package main

import (
	"log"
	"os"

	"github.com/NikiTesla/geant4help"
	"github.com/NikiTesla/geant4help/pkg/environment"
	"github.com/NikiTesla/geant4help/pkg/handler"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("can't load env variables, err: %s", err.Error())
	}

	configFile := os.Getenv("CONFIGFILE")
	env, err := environment.NewEnvironment(configFile)
	if err != nil {
		log.Fatalf("can't load environment, err: %s", err.Error())
	}

	rtr := handler.Handler{Env: env}.InitRouter()
	server := geant4help.Server{}

	if err := server.Run(env.Config.Port, rtr); err != nil {
		log.Fatalf("error occured during serving: %s", err.Error())
	}
}
