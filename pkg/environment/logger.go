package environment

type Logger struct {
	level string `jsin`
}

func NewLogger(cfg LoggerConfig) (*Logger, error) {
	return &Logger{}, nil
}
