package Core

import (
	"context"
	"log"
)

type LoggerFactory interface {
	Logger(ctx context.Context) Logger
}

type Logger interface {
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
}

type noopLoggerFactory struct{}

func (*noopLoggerFactory) Logger(_ context.Context) Logger {
	return noopLogger{}
}

type noopLogger struct{}

func (noopLogger) Debugf(format string, args ...any) {
	log.Printf(" [DEBUG] "+format, args...)
}
func (noopLogger) Infof(format string, args ...any) {
	log.Printf(" [INFO ] "+format, args...)
}
func (noopLogger) Warnf(format string, args ...any) {
	log.Printf(" [WARN ] "+format, args...)
}
func (noopLogger) Errorf(format string, args ...any) {
	log.Printf(" [ERROR ] "+format, args...)
}
