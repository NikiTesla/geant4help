package environment

import (
	"encoding/json"
	"fmt"
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
}

type Config struct {
	StaticDir string   `json:"static-dir"`
	DBConfig  DBConfig `json:"db-config"`
	Port      int      `json:"port"`
	Host      string   `json:"host"`
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

	fmt.Print(config)

	return &config, nil
}
