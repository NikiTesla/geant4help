package environment

import (
	"database/sql"

	"go.uber.org/zap"
)

type Env struct {
	Config *Config
	Logger *zap.Logger
	DB     *sql.DB
}
