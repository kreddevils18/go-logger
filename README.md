# Go-Logger

A flexible, structured logging package for Go applications built on top of [Uber's zap](https://github.com/uber-go/zap) logger.

## Features

- Simple, clean interface for structured logging
- Multiple log levels (Debug, Info, Warn, Error, DPanic, Panic, Fatal)
- Support for both console and JSON output formats
- Configurable output destinations (stdout, stderr, files)
- Three logging styles:
  - Basic (`logger.Info("message")`)
  - Template (`logger.Infof("Hello %s", "world")`)
  - Structured (`logger.Infow("message", "key", value)`)
- Built on the high-performance zap logging library

## Installation

```bash
go get github.com/kreddevils18/go-logger
```

## Quick Start

```go
package main

import (
	"github.com/kreddevils18/go-logger"
)

func main() {
	// Create a default logger (development mode with console output)
	logger, err := gologger.NewDefaultLogger()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logger.Sync() // Flush any buffered log entries

	// Basic logging
	logger.Info("This is an info message")
	logger.Debug("This is a debug message")

	// Template-style logging
	logger.Infof("Hello, %s!", "world")

	// Structured logging with key-value pairs
	logger.Infow("User logged in",
		"userId", 123,
		"username", "johndoe",
		"loginTime", "2023-04-26T11:45:00Z",
	)
}
```

## Configuration

You can customize the logger with various configuration options:

```go
package main

import (
	"github.com/kreddevils18/go-logger"
)

func main() {
	// Custom configuration
	config := &gologger.Config{
		Level:      "info",              // Minimum log level
		Encoding:   "json",              // "json" or "console"
		Outputs:    []string{"stdout", "/var/log/myapp.log"}, // Output destinations
		ErrOutputs: []string{"stderr"},  // Error output destinations
	}

	logger, err := gologger.NewLogger(config)
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logger.Sync()

	// Use the logger
	logger.Info("Application started")
}
```

## Log Levels

The logger supports the following log levels (from lowest to highest severity):

- **Debug**: Detailed information for debugging purposes
- **Info**: General operational information
- **Warn**: Warning events that might cause issues
- **Error**: Error events that might still allow the application to continue
- **DPanic**: Critical errors that will cause panic in development mode
- **Panic**: Critical errors that will cause panic in all environments
- **Fatal**: Critical errors that will terminate the application

## Interface

The `Logger` interface provides a consistent API for logging:

```go
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	DPanic(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})

	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	DPanicf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})

	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	DPanicw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})

	Sync() error
}
```

## Best Practices

1. **Use structured logging**: Prefer the `*w` methods (e.g., `Infow`) with key-value pairs for better searchability and analysis.

2. **Choose appropriate log levels**: Use Info for normal operations, Debug for detailed troubleshooting, Error for actual errors.

3. **Always defer Sync()**: Call `defer logger.Sync()` after initializing the logger to ensure all logs are flushed.

4. **Use dependency injection**: Pass the logger as a dependency to your components rather than using a global instance.

5. **Include context**: Add relevant context information to your logs (request IDs, user IDs, etc.).

## License

[MIT License](LICENSE)
