package environment

import (
	"encoding/json"
	"log"
	"os"
)

type DBConfig struct {
	Port     int    `json:"port"`
	Host     string `json:"host"`
	Username string `json:"username"`
	DBname   string `json:"dbname"`
	SSlmode  string `json:"sslmode"`
}

type LoggerConfig struct {
	LogFile string `json:"log-file"`
}

type Config struct {
	Port         int          `json:"port"`
	Host         string       `json:"host"`
	StaticDir    string       `json:"static-dir"`
	LoggerConfig LoggerConfig `json:"logger-config"`
	DBConfig     DBConfig     `json:"db-config"`
}

func NewConfig(configFile string) (*Config, error) {
	rawJSON, err := os.ReadFile(configFile)
	if err != nil {
		log.Printf("Can't read config file, err: %s", err.Error())
		return nil, err
	}

	var config Config
	err = json.Unmarshal(rawJSON, &config)
	if err != nil {
		log.Printf("Can't unmarshall config json, err: %s", err.Error())
		return nil, err
	}

	return &config, nil
}
