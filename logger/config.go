package logger

type Config struct {
	Environment string
	Level       string

	StandardLogConfig *StandardLogConfig
}

type StandardLogConfig struct {
	OutputPaths      []string
	ErrorOutputPaths []string
}
