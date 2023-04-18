package environment

import (
	"fmt"
	"log"

	"go.uber.org/zap"
)

type Environment struct {
	Config   *Config
	Logger   *zap.Logger
	DataBase *DataBase
}

func NewEnvironment(configFile string) (*Environment, error) {
	log.Println("Setting environment")

	cfg, err := NewConfig(configFile)
	if err != nil {
		return nil, fmt.Errorf("can't create environment because of config: %s", err.Error())
	}

	logger, err := NewZapLogger(cfg.LoggerConfig)
	if err != nil {
		return nil, fmt.Errorf("can't create enviroment beacuse of logger: %s", err.Error())
	}

	dataBase, err := NewDataBase(cfg.DBConfig)
	if err != nil {
		return nil, fmt.Errorf("cant create environment bacause of database: %s", err.Error())
	}

	log.Printf("Host is %s\n", cfg.Host)
	log.Printf("Port is %d\n", cfg.Port)
	log.Println("Logger config is Zap Logger")
	log.Printf("Database config is %+v", cfg.DBConfig)

	return &Environment{
		Config:   cfg,
		Logger:   logger,
		DataBase: dataBase,
	}, nil
}
