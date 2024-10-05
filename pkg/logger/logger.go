package logger

import (
	"log"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// LogWriter wraps zerolog.Logger to implement io.Writer.
type LogWriter struct {
	logger zerolog.Logger
}

// Write implements the io.Writer interface for LogWriter.
func (lw *LogWriter) Write(p []byte) (n int, err error) {
	lw.logger.Error().Msg(string(p))
	return len(p), nil
}

type Logger struct {
	logger zerolog.Logger
}

// NewLogger initializes a new Logger instance.
func NewLogger() *Logger {
	// Configure zerolog
	zerolog.TimeFieldFormat = time.RFC3339Nano
	output := zerolog.ConsoleWriter{Out: os.Stderr}
	logger := zerolog.New(output).With().Timestamp().Logger()

	return &Logger{logger: logger}
}

// NewLogLogger returns a standard log.Logger that logs errors using zerolog.
func (l *Logger) NewLogLogger() *log.Logger {
	return log.New(&LogWriter{logger: l.logger}, "", 0)
}

// Info logs an info message with optional key-value pairs.
func (l *Logger) Info(msg string, keysAndValues ...any) {
	fields := pairs(keysAndValues)
	event := l.logger.Info()
	for i := 0; i < len(fields); i += 2 {
		event = event.Interface(fields[i].(string), fields[i+1])
	}
	event.Msg(msg)
}

// Error logs an error message with optional key-value pairs.
func (l *Logger) Error(msg string, keysAndValues ...any) {
	fields := pairs(keysAndValues)
	event := l.logger.Error()
	for i := 0; i < len(fields); i += 2 {
		event = event.Interface(fields[i].(string), fields[i+1])
	}
	event.Msg(msg)
}

// Debug logs a debug message with optional key-value pairs.
func (l *Logger) Debug(msg string, keysAndValues ...any) {
	fields := pairs(keysAndValues)
	event := l.logger.Debug()
	for i := 0; i < len(fields); i += 2 {
		event = event.Interface(fields[i].(string), fields[i+1])
	}
	event.Msg(msg)
}

// Warn logs a warning message with optional key-value pairs.
func (l *Logger) Warn(msg string, keysAndValues ...any) {
	fields := pairs(keysAndValues)
	event := l.logger.Warn()
	for i := 0; i < len(fields); i += 2 {
		event = event.Interface(fields[i].(string), fields[i+1])
	}
	event.Msg(msg)
}

// Fatal logs a fatal message with optional key-value pairs.
// It will terminate the program after logging.
func (l *Logger) Fatal(msg string, keysAndValues ...any) {
	fields := pairs(keysAndValues)
	event := l.logger.Fatal()
	for i := 0; i < len(fields); i += 2 {
		event = event.Interface(fields[i].(string), fields[i+1])
	}
	event.Msg(msg)
}

// Panic logs a panic message with optional key-value pairs.
// It will cause a panic after logging.
func (l *Logger) Panic(msg string, keysAndValues ...any) {
	fields := pairs(keysAndValues)
	event := l.logger.Panic()
	for i := 0; i < len(fields); i += 2 {
		event = event.Interface(fields[i].(string), fields[i+1])
	}
	event.Msg(msg)
}

// pairs ensures the key-values are in pairs.
// It drops the last entry if there's an odd number of entries.
func pairs(keysAndValues []any) []any {
	fields := keysAndValues
	num := len(fields)
	if num%2 == 1 {
		// Odd number of key-values, drop the last one
		num--
		fields = fields[:num]
	}
	return fields
}
