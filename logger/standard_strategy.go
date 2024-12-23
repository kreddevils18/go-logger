package logger

import (
	"go.uber.org/zap"
)

type StandardLogStrategy struct {
	logger *zap.SugaredLogger
}

func (s *StandardLogStrategy) Init(cfg Config) error {
	if cfg.StandardLogConfig == nil {
		return nil
	}

	var zapConfig zap.Config
	if cfg.Environment == "production" {
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
	}

	level := getLogLevel(cfg.Level)
	zapConfig.Level = zap.NewAtomicLevelAt(level)

	zapConfig.OutputPaths = cfg.StandardLogConfig.OutputPaths
	zapConfig.ErrorOutputPaths = cfg.StandardLogConfig.ErrorOutputPaths

	log, err := zapConfig.Build()
	if err != nil {
		return nil
	}

	s.logger = log.Sugar()
	return nil
}

func (s *StandardLogStrategy) GetLogger() *zap.SugaredLogger {
	if s.logger == nil {
		panic("Logger is not initialized. Call init first.")
	}
	return s.logger
}

func (s *StandardLogStrategy) Sync() error {
	if s.logger != nil {
		return s.logger.Sync()
	}

	return nil
}

// Debug logs a debug level message
func (s *StandardLogStrategy) Debug(args ...interface{}) {
	s.logger.Debug(args...)
}

// Debugf logs a formatted debug level message
func (s *StandardLogStrategy) Debugf(format string, args ...interface{}) {
	s.logger.Debugf(format, args...)
}

// Info logs an info level message
func (s *StandardLogStrategy) Info(args ...interface{}) {
	s.logger.Info(args...)
}

// Infof logs a formatted info level message
func (s *StandardLogStrategy) Infof(format string, args ...interface{}) {
	s.logger.Infof(format, args...)
}

// Warn logs a warn level message
func (s *StandardLogStrategy) Warn(args ...interface{}) {
	s.logger.Warn(args...)
}

// Warnf logs a formatted warn level message
func (s *StandardLogStrategy) Warnf(format string, args ...interface{}) {
	s.logger.Warnf(format, args...)
}

// Error logs an error level message
func (s *StandardLogStrategy) Error(args ...interface{}) {
	s.logger.Error(args...)
}

// Errorf logs a formatted error level message
func (s *StandardLogStrategy) Errorf(format string, args ...interface{}) {
	s.logger.Errorf(format, args...)
}

// Fatal logs a fatal level message and exits
func (s *StandardLogStrategy) Fatal(args ...interface{}) {
	s.logger.Fatal(args...)
}

// Fatalf logs a formatted fatal level message and exits
func (s *StandardLogStrategy) Fatalf(format string, args ...interface{}) {
	s.logger.Fatalf(format, args...)
}
