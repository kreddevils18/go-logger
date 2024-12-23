package logger

import "go.uber.org/zap/zapcore"

func getLogLevel(level string) zapcore.Level {
	logLevelMapping := map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
		"fatal": zapcore.FatalLevel,
	}

	if val, ok := logLevelMapping[level]; ok {
		return val
	} else {
		return zapcore.InfoLevel
	}
}
