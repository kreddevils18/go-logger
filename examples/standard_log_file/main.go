package main

import (
	"go-logger/logger"
	"log"
)

func main() {
	cfg := logger.Config{
		Environment: "development",
		Level:       "info",
		StandardLogConfig: &logger.StandardLogConfig{
			OutputPaths:      []string{"20241223.log"},
			ErrorOutputPaths: []string{"stderr"},
		},
	}

	standardLog := &logger.StandardLogStrategy{}

	err := standardLog.Init(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Standard Log: %v", err)
	}
	defer standardLog.Sync()

	standardLog.Info("Application started")
	standardLog.Warn("This is a warning message")
	standardLog.Error("This is an error message")
	standardLog.Debugf("Debugging value: %d", 42)
}
