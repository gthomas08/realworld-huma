package logger

import (
	"log"
	"log/slog"
	"os"
)

type Logger struct {
	logger *slog.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}

func (l *Logger) NewLogLogger() *log.Logger {
	return slog.NewLogLogger(l.logger.Handler(), slog.LevelError)
}

func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}
