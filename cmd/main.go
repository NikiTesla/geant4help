package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := getConfig(); err != nil {
		logrus.Fatalf("can't load configs, err: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("can't load env variables, err: %s", err.Error())
	}

}

func getConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("conf.debug")
	return viper.ReadInConfig()
}
