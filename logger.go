package gologger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger defines the interface for logging operations.
// This allows for dependency injection and easier testing.
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

	Sync() error // Add Sync method from SugaredLogger
}

// zapLoggerImpl is a concrete implementation of the Logger interface
// using zap.SugaredLogger.
// It embeds *zap.SugaredLogger to inherit its methods, but we also
// explicitly define them to make the implementation clear.
type zapLoggerImpl struct {
	*zap.SugaredLogger
}

// Config holds the configuration for the logger.
type Config struct {
	Level      string   // Log level: debug, info, warn, error, dpanic, panic, fatal
	Encoding   string   // Encoding: console, json
	Outputs    []string // Output paths: stdout, stderr, or file paths
	ErrOutputs []string // Error output paths: stdout, stderr, or file paths
}

// NewLogger creates a new logger implementing the Logger interface based on the provided configuration.
// It provides sensible defaults for development environments if no config is specified.
func NewLogger(cfg *Config) (Logger, error) { // Return Logger interface
	var zapCfg zap.Config

	if cfg == nil {
		// Default to development configuration
		zapCfg = zap.NewDevelopmentConfig()
		// Increase caller skip for zap's default dev config as well
		zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // Pretty console output
	} else {
		// Use provided configuration
		level := zap.NewAtomicLevel()
		if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
			// Default to info level if parsing fails
			level.SetLevel(zap.InfoLevel)
		}

		encoding := "console"
		if cfg.Encoding == "json" {
			encoding = "json"
		}

		outputs := []string{"stdout"}
		if len(cfg.Outputs) > 0 {
			outputs = cfg.Outputs
		}

		errOutputs := []string{"stderr"}
		if len(cfg.ErrOutputs) > 0 {
			errOutputs = cfg.ErrOutputs
		}

		encoderCfg := zap.NewProductionEncoderConfig()
		if encoding == "console" {
			encoderCfg = zap.NewDevelopmentEncoderConfig()
			// Customize console output for better readability
			encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
			encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		} else {
			// Customize json output if needed
			encoderCfg.TimeKey = "timestamp"
			encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		}

		zapCfg = zap.Config{
			Level:            level,
			Development:      encoding == "console", // Assume console is for development
			Encoding:         encoding,
			EncoderConfig:    encoderCfg,
			OutputPaths:      outputs,
			ErrorOutputPaths: errOutputs,
			// Disable sampling for now, can be configured if needed
			Sampling: nil,
			// Add caller skip to show the correct caller function
			// This might need adjustment depending on how the logger is wrapped/used
			// InitialCallerSkip: 0, // Consider zap.AddCallerSkip(1) if wrapping
		}
	}

	// Ensure caller information is included
	// Add AddCallerSkip(2) because we have two levels of wrappers:
	// 1. The zapLoggerImpl methods (like Debug, Info, etc.)
	// 2. The NewLogger function itself
	// Note: If only using embedding, AddCallerSkip(1) would suffice.
	loggerCore, err := zapCfg.Build(zap.AddCaller(), zap.AddCallerSkip(2))
	if err != nil {
		return nil, err
	}

	// Create the SugaredLogger
	sugar := loggerCore.Sugar()

	// Return our custom struct embedding the SugaredLogger
	// The pointer *zapLoggerImpl satisfies the Logger interface.
	return &zapLoggerImpl{SugaredLogger: sugar}, nil
}

// NewDefaultLogger creates a default logger implementing the Logger interface.
func NewDefaultLogger() (Logger, error) { // Return Logger interface
	return NewLogger(nil) // Uses development defaults
}

// --- Explicit implementation of Logger interface methods ---

func (l *zapLoggerImpl) Debug(args ...interface{}) {
	l.SugaredLogger.Debug(args...)
}

func (l *zapLoggerImpl) Info(args ...interface{}) {
	l.SugaredLogger.Info(args...)
}

func (l *zapLoggerImpl) Warn(args ...interface{}) {
	l.SugaredLogger.Warn(args...)
}

func (l *zapLoggerImpl) Error(args ...interface{}) {
	l.SugaredLogger.Error(args...)
}

func (l *zapLoggerImpl) DPanic(args ...interface{}) {
	l.SugaredLogger.DPanic(args...)
}

func (l *zapLoggerImpl) Panic(args ...interface{}) {
	l.SugaredLogger.Panic(args...)
}

func (l *zapLoggerImpl) Fatal(args ...interface{}) {
	l.SugaredLogger.Fatal(args...)
}

func (l *zapLoggerImpl) Debugf(template string, args ...interface{}) {
	l.SugaredLogger.Debugf(template, args...)
}

func (l *zapLoggerImpl) Infof(template string, args ...interface{}) {
	l.SugaredLogger.Infof(template, args...)
}

func (l *zapLoggerImpl) Warnf(template string, args ...interface{}) {
	l.SugaredLogger.Warnf(template, args...)
}

func (l *zapLoggerImpl) Errorf(template string, args ...interface{}) {
	l.SugaredLogger.Errorf(template, args...)
}

func (l *zapLoggerImpl) DPanicf(template string, args ...interface{}) {
	l.SugaredLogger.DPanicf(template, args...)
}

func (l *zapLoggerImpl) Panicf(template string, args ...interface{}) {
	l.SugaredLogger.Panicf(template, args...)
}

func (l *zapLoggerImpl) Fatalf(template string, args ...interface{}) {
	l.SugaredLogger.Fatalf(template, args...)
}

func (l *zapLoggerImpl) Debugw(msg string, keysAndValues ...interface{}) {
	l.SugaredLogger.Debugw(msg, keysAndValues...)
}

func (l *zapLoggerImpl) Infow(msg string, keysAndValues ...interface{}) {
	l.SugaredLogger.Infow(msg, keysAndValues...)
}

func (l *zapLoggerImpl) Warnw(msg string, keysAndValues ...interface{}) {
	l.SugaredLogger.Warnw(msg, keysAndValues...)
}

func (l *zapLoggerImpl) Errorw(msg string, keysAndValues ...interface{}) {
	l.SugaredLogger.Errorw(msg, keysAndValues...)
}

func (l *zapLoggerImpl) DPanicw(msg string, keysAndValues ...interface{}) {
	l.SugaredLogger.DPanicw(msg, keysAndValues...)
}

func (l *zapLoggerImpl) Panicw(msg string, keysAndValues ...interface{}) {
	l.SugaredLogger.Panicw(msg, keysAndValues...)
}

func (l *zapLoggerImpl) Fatalw(msg string, keysAndValues ...interface{}) {
	l.SugaredLogger.Fatalw(msg, keysAndValues...)
}

func (l *zapLoggerImpl) Sync() error {
	return l.SugaredLogger.Sync()
}
