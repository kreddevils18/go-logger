package logger

import "go.uber.org/zap"

type LogStrategy interface {
	Init(cfg Config) error
	GetLogger() *zap.SugaredLogger
	Sync() error
}

type LogContext struct {
	strategy LogStrategy
}

func (lc *LogContext) SetStrategy(strategy LogStrategy) {
	lc.strategy = strategy
}

func (lc *LogContext) Init(cfg Config) error {
	if lc.strategy == nil {
		return nil
	}

	return lc.strategy.Init(cfg)
}

func (lc *LogContext) GetLogger() *zap.SugaredLogger {
	if lc.strategy == nil {
		panic("Logger strategy is not set")
	}

	return lc.strategy.GetLogger()
}

func (lc *LogContext) Sync() error {
	if lc.strategy != nil {
		return lc.strategy.Sync()
	}

	return nil
}
