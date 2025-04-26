package gologger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func TestNewLogger_NilConfig(t *testing.T) {
	logger, err := NewLogger(nil)
	require.NoError(t, err)
	require.NotNil(t, logger)

	// Type assert to access underlying zap logger for testing internals
	impl, ok := logger.(*zapLoggerImpl)
	require.True(t, ok)
	assert.True(t, impl.SugaredLogger.Desugar().Core().Enabled(zapcore.DebugLevel))
}

func TestNewLogger_ConsoleConfig(t *testing.T) {
	cfg := &Config{
		Level:      "warn",
		Encoding:   "console",
		Outputs:    []string{"stdout"},
		ErrOutputs: []string{"stderr"},
	}

	logger, err := NewLogger(cfg)
	require.NoError(t, err)
	require.NotNil(t, logger)

	// Type assert
	impl, ok := logger.(*zapLoggerImpl)
	require.True(t, ok)

	// Check configured level
	assert.True(t, impl.SugaredLogger.Desugar().Core().Enabled(zapcore.WarnLevel))
	assert.False(t, impl.SugaredLogger.Desugar().Core().Enabled(zapcore.InfoLevel))
}

func TestNewLogger_JsonConfig(t *testing.T) {
	cfg := &Config{
		Level:      "debug",
		Encoding:   "json",
		Outputs:    []string{"stdout"},
		ErrOutputs: []string{"stderr"},
	}

	logger, err := NewLogger(cfg)
	require.NoError(t, err)
	require.NotNil(t, logger)

	// Type assert
	impl, ok := logger.(*zapLoggerImpl)
	require.True(t, ok)

	// Check configured level
	assert.True(t, impl.SugaredLogger.Desugar().Core().Enabled(zapcore.DebugLevel))

	// Basic check for json encoding (implementation detail, might be brittle)
	// We can check if the Development flag is false as per our logic
	// zapCfg.Development is false for json encoding in our NewLogger func
}

func TestNewLogger_InvalidLevel(t *testing.T) {
	cfg := &Config{
		Level:    "invalid-level",
		Encoding: "console",
	}

	logger, err := NewLogger(cfg)
	require.NoError(t, err)
	require.NotNil(t, logger)

	// Type assert
	impl, ok := logger.(*zapLoggerImpl)
	require.True(t, ok)

	// Check if it defaults to InfoLevel when level is invalid
	assert.True(t, impl.SugaredLogger.Desugar().Core().Enabled(zapcore.InfoLevel))
	assert.False(t, impl.SugaredLogger.Desugar().Core().Enabled(zapcore.DebugLevel))
}

func TestNewDefaultLogger(t *testing.T) {
	logger, err := NewDefaultLogger()
	require.NoError(t, err)
	require.NotNil(t, logger)

	// Type assert
	impl, ok := logger.(*zapLoggerImpl)
	require.True(t, ok)

	// Should behave the same as NewLogger(nil)
	assert.True(t, impl.SugaredLogger.Desugar().Core().Enabled(zapcore.DebugLevel))
}

func TestLoggingExecution(t *testing.T) {
	logger, err := NewDefaultLogger()
	require.NoError(t, err)

	assert.NotPanics(t, func() {
		logger.Debugw("Debug message", "key", "value")
		logger.Info("Info message")
		logger.Warnf("Warn message: %d", 123)
		logger.Error("Error message")
	})
}
