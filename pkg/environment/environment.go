package environment

import (
	"fmt"
	"log"
)

type Environment struct {
	Config   *Config
	Logger   *Logger
	DataBase *DataBase
}

func NewEnvironment(configFile string) (*Environment, error) {
	log.Println("Setting environment")

	cfg, err := NewConfig(configFile)
	if err != nil {
		return nil, fmt.Errorf("can't create environment because of config: %s", err.Error())
	}

	logger, err := NewLogger(cfg.LoggerConfig)
	if err != nil {
		return nil, fmt.Errorf("can't create enviroment beacuse of logger: %s", err.Error())
	}

	dataBase, err := NewDataBase(cfg.DBConfig)
	if err != nil {
		return nil, fmt.Errorf("cant create environment bacause of database: %s", err.Error())
	}

	// change on logger
	log.Printf("Host is %s\n", cfg.Host)
	log.Printf("Port is %d\n", cfg.Port)
	log.Printf("Logger Config is %+v", cfg.LoggerConfig)
	log.Printf("Database config is %+v", cfg.DBConfig)

	return &Environment{
		Config:   cfg,
		Logger:   logger,
		DataBase: dataBase,
	}, nil
}
