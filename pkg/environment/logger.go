package environment

import (
	"fmt"

	"go.uber.org/zap"
)

func NewZapLogger(cfg LoggerConfig) (*zap.Logger, error) {
	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("can't create logger with current config, err: %s", err.Error())
	}

	return logger, nil
}
